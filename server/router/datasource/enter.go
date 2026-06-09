package datasource

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	DataSourceFieldRouter
	APISourceRouter
	FileSourceRouter
	FileFieldRouter
	SearchTemplateRouter
}

var (
	DSFieldApi    = api.ApiGroupApp.DatasourceApiGroup.DataSourceFieldApi
	ASourceApi    = api.ApiGroupApp.DatasourceApiGroup.APISourceApi
	FSourceApi    = api.ApiGroupApp.DatasourceApiGroup.FileSourceApi
	FFieldApi     = api.ApiGroupApp.DatasourceApiGroup.FileFieldApi
	searchTempApi = api.ApiGroupApp.DatasourceApiGroup.SearchTemplateApi
)
