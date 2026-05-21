package router

import (
	apiPlugin "github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/api"
	"github.com/gin-gonic/gin"
)

var SSOLogin = new(ssoLogin)

type ssoLogin struct{}

// Init 初始化 SSO 登录路由信息
func (r *ssoLogin) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := public.Group("sso")
		group.GET("login", apiPlugin.ApiGroupApp.SSOLoginApi.SSOLogin)
		group.POST("generate-sign", apiPlugin.ApiGroupApp.SSOLoginApi.GenerateTestSign)
	}
}
