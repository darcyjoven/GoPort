package router

import (
	apiPlugin "github.com/flipped-aurora/gin-vue-admin/server/plugin/sso/api"
)

var (
	Router = new(router)
	api    = apiPlugin.ApiGroupApp
)

type router struct{ SSOLogin ssoLogin }

