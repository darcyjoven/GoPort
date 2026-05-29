package datasource

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	DatabaseConfigRouter
	SearchTemplateRouter
}

var (
	dbConfigApi   = api.ApiGroupApp.DatasourceApiGroup.DatabaseConfigApi
	searchTempApi = api.ApiGroupApp.DatasourceApiGroup.SearchTemplateApi
)
