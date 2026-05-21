package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/model"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.AutoMigrate(&model.SSOToken{})
	if err != nil {
		global.GVA_LOG.Error("SSO 数据表迁移失败!", zap.Error(err))
	}
}
