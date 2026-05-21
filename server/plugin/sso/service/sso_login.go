package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SSOLoginService struct{}

var Config *config.SSO

// InitConfig 初始化配置
func InitConfig(cfg *config.SSO) {
	Config = cfg
}

// GetProviderByName 获取提供商配置
func (t *SSOLoginService) GetProviderByName(name string) (*config.Provider, bool) {
	for _, p := range Config.Providers {
		if p.Name == name && p.Enabled {
			return &p, true
		}
	}
	return nil, false
}

// VerifySignature 验证签名
func (t *SSOLoginService) VerifySignature(req *request.SSOLoginReq, secretKey string) bool {
	// 构建签名字符串：provider=xxx&username=xxx&realName=xxx&department=xxx×tamp=xxx
	signStr := fmt.Sprintf("provider=%s&username=%s&realName=%s&department=%s×tamp=%d",
		req.Provider, req.Username, req.RealName, req.Department, req.Timestamp)

	// 使用 HMAC-SHA256 签名
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(signStr))
	expectedSign := hex.EncodeToString(h.Sum(nil))

	return expectedSign == req.Sign
}

// VerifyTimestamp 验证时间戳是否在5分钟内
func (t *SSOLoginService) VerifyTimestamp(timestamp int64) bool {
	now := time.Now().Unix()
	diff := now - timestamp
	return diff >= -300 && diff <= 300 // 5分钟内
}

// VerifySSOToken 验证 SSO 服务器
func (t *SSOLoginService) VerifySSOToken(token string, provider *config.Provider) (bool, error) {
	// 优先检查本地缓存中是否有未过期的token
	cachedToken := &model.SSOToken{}
	err := global.GVA_DB.Where("token = ? AND expires_at > ?", token, time.Now()).First(cachedToken).Error
	if err == nil {
		return true, nil
	}

	// 本地没有，调用 SSO 验证接口
	verifyURL := provider.VerifyURL
	if !strings.Contains(verifyURL, "?") {
		verifyURL += "?"
	}
	verifyURL += fmt.Sprintf("%s=%s", provider.VerifyText, url.QueryEscape(token))

	resp, err := http.Get(verifyURL)
	if err != nil {
		global.GVA_LOG.Error("调用 SSO 验证接口失败", zap.Error(err))
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	success := strings.Contains(string(body), provider.SuccessText)
	if success {
		// 验证成功，保存token
		expireTime, err := utils.ParseDuration(provider.TokenExpireTime)
		if err != nil {
			expireTime = time.Hour * 24 * 7 // 默认7天
		}

		newToken := &model.SSOToken{
			ProviderName: provider.Name,
			Token:        token,
			ExpiresAt:    time.Now().Add(expireTime),
		}
		global.GVA_DB.Create(newToken)
	}

	return success, nil
}

// FindOrCreateUser 查找或创建用户
func (t *SSOLoginService) FindOrCreateUser(req *request.SSOLoginReq, provider *config.Provider) (*system.SysUser, error) {
	// 先查找用户
	var user system.SysUser
	err := global.GVA_DB.Where("username = ?", req.Username).First(&user).Error

	if err == nil {
		// 用户已存在，返回
		return &user, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 数据库查询错误
		return nil, err
	}

	// 用户不存在，检查是否允许自动注册
	if !provider.AllowAutoRegister {
		return nil, errors.New("用户不存在且不允许自动注册")
	}

	// 自动创建用户
	nickName := req.RealName
	if nickName == "" {
		nickName = req.Username
	}

	defaultAuthorityId := provider.DefaultAuthorityId
	if defaultAuthorityId == 0 {
		defaultAuthorityId = 888 // 默认普通用户角色
	}

	newUser := &system.SysUser{
		UUID:        uuid.New(),
		Username:    req.Username,
		NickName:    nickName,
		Password:    utils.BcryptHash(uuid.New().String()), // 随机密码
		AuthorityId: defaultAuthorityId,
		Enable:      1,
	}

	err = global.GVA_DB.Create(newUser).Error
	if err != nil {
		return nil, err
	}

	// 重新查询完整用户信息（包括角色）
	var createdUser system.SysUser
	err = global.GVA_DB.Preload("Authority").Where("id = ?", newUser.ID).First(&createdUser).Error

	return &createdUser, nil
}

// SSOLogin SSO 登录主流程
func (t *SSOLoginService) SSOLogin(req *request.SSOLoginReq, remoteIP string) (*system.SysUser, string, time.Time, error) {
	// 1. 获取提供商配置
	provider, ok := t.GetProviderByName(req.Provider)
	if !ok {
		return nil, "", time.Time{}, errors.New("无效的提供商")
	}

	// 2. 验证签名
	if !t.VerifySignature(req, provider.SecretKey) {
		return nil, "", time.Time{}, errors.New("签名验证失败")
	}

	// 3. 验证时间戳
	if !t.VerifyTimestamp(req.Timestamp) {
		return nil, "", time.Time{}, errors.New("请求已过期")
	}

	// 4. 如果提供了token，验证 SSO token
	if req.Token != "" {
		valid, err := t.VerifySSOToken(req.Token, provider)
		if err != nil {
			return nil, "", time.Time{}, err
		}
		if !valid {
			return nil, "", time.Time{}, errors.New("SSO token验证失败")
		}
	}

	// 5. 查找或创建用户
	user, err := t.FindOrCreateUser(req, provider)
	if err != nil {
		return nil, "", time.Time{}, err
	}

	if user.Enable != 1 {
		return nil, "", time.Time{}, errors.New("用户已被禁用")
	}

	// 6. 生成JWT token
	j := utils.NewJWT()
	claims := j.CreateClaims(systemReq.BaseClaims{
		ID:          user.ID,
		UUID:        user.UUID,
		Username:    user.Username,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		return nil, "", time.Time{}, err
	}

	// 7. 更新token缓存（如果启用多点登录）
	if global.GVA_CONFIG.System.UseMultipoint {
		if err := utils.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		}
	}

	return user, token, claims.ExpiresAt.Time, nil
}

// GenerateSign 生成测试签名（用于调试）
func (t *SSOLoginService) GenerateSign(req *request.SSOLoginReq, secretKey string) string {
	signStr := fmt.Sprintf("provider=%s&username=%s&realName=%s&department=%s×tamp=%d",
		req.Provider, req.Username, req.RealName, req.Department, req.Timestamp)

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}
