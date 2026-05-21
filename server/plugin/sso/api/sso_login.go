package api

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SSOLoginApi struct{}

// SSOLogin SSO 登录
// @Tags SSOLogin
// @Summary SSO 系统免密登录
// @Accept json
// @Produce json
// @Param provider query string true "提供商名称"
// @Param username query string true "用户名"
// @Param realName query string true "真实姓名"
// @Param department query string false "部门"
// @Param targetPage query string false "目标页面"
// @Param timestamp query int64 true "时间戳"
// @Param sign query string true "签名"
// @Param token query string false "SSO token"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /sso/login [get]
func (t *SSOLoginApi) SSOLogin(c *gin.Context) {
	var req request.SSOLoginReq
	
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	if req.Provider == "" || req.Username == "" || req.Sign == "" || req.Timestamp == 0 {
		response.FailWithMessage("缺少必要参数", c)
		return
	}
	
	// 调用服务层处理登录
	user, token, expiresAt, err := service.ServiceGroupApp.SSOLoginService.SSOLogin(&req, c.ClientIP())
	if err != nil {
		global.GVA_LOG.Error("SSO 登录失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 设置cookie中的token
	utils.SetToken(c, token, int(expiresAt.Unix()-time.Now().Unix()))
	
	// 如果提供了目标页面，进行重定向
	if req.TargetPage != "" {
		// 构建重定向URL，带上token信息
		redirectURL := req.TargetPage
		if !strings.Contains(redirectURL, "?") {
			redirectURL += "?"
		} else {
			redirectURL += "&"
		}
		redirectURL += fmt.Sprintf("token=%s&expiresAt=%d", url.QueryEscape(token), expiresAt.Unix()*1000)
		c.Redirect(302, redirectURL)
		return
	}
	
	// 否则返回JSON响应
	response.OkWithDetailed(systemRes.LoginResponse{
		User:      *user,
		Token:     token,
		ExpiresAt: expiresAt.Unix() * 1000,
	}, "登录成功", c)
}

// GenerateTestSign 生成测试签名（调试用）
// @Tags SSOLogin
// @Summary 生成测试签名
// @Accept json
// @Produce json
// @Param data body request.SSOLoginReq true "请求参数"
// @Success 200 {object} response.Response{data=string,msg=string} "返回签名"
// @Router /sso/generate-sign [post]
func (t *SSOLoginApi) GenerateTestSign(c *gin.Context) {
	var req request.SSOLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	provider, ok := service.ServiceGroupApp.SSOLoginService.GetProviderByName(req.Provider)
	if !ok {
		response.FailWithMessage("无效的提供商", c)
		return
	}
	
	sign := service.ServiceGroupApp.SSOLoginService.GenerateSign(&req, provider.SecretKey)
	response.OkWithDetailed(sign, "生成成功", c)
}
