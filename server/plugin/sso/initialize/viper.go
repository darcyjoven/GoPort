package initialize

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/service"
	"go.uber.org/zap"
)

func Viper() *config.SSO {
	cfg := &config.SSO{}
	err := global.GVA_VP.UnmarshalKey("sso", cfg)
	if err != nil {
		global.GVA_LOG.Error("读取SSO配置失败!", zap.Error(err))
	} else {
		global.GVA_LOG.Info(fmt.Sprintf("成功读取 %d 个SSO提供商配置", len(cfg.Providers)))
	}
	service.InitConfig(cfg)
	return cfg
}
