package datasource

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	DatabaseConfigRouter
	SearchTemplateRouter
	DataSourceFieldRouter
	APISourceRouter
	FileSourceRouter
	FileFieldRouter
}

var (
	dbConfigApi   = api.ApiGroupApp.DatasourceApiGroup.DatabaseConfigApi
	searchTempApi = api.ApiGroupApp.DatasourceApiGroup.SearchTemplateApi
	DSFieldApi    = api.ApiGroupApp.DatasourceApiGroup.DataSourceFieldApi
	ASourceApi    = api.ApiGroupApp.DatasourceApiGroup.APISourceApi
	FSourceApi    = api.ApiGroupApp.DatasourceApiGroup.FileSourceApi
	FFieldApi     = api.ApiGroupApp.DatasourceApiGroup.FileFieldApi
)
