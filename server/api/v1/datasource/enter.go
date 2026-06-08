package datasource

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DatabaseConfigApi
	SearchTemplateApi
	DataSourceFieldApi
	APISourceApi
	FileSourceApi
	FileFieldApi
}

var (
	dbConfigService   = service.ServiceGroupApp.DatasourceServiceGroup.DatabaseConfigService
	searchTempService = service.ServiceGroupApp.DatasourceServiceGroup.SearchTemplateService
	DSFieldService    = service.ServiceGroupApp.DatasourceServiceGroup.DataSourceFieldService
	ASourceService    = service.ServiceGroupApp.DatasourceServiceGroup.APISourceService
	FSourceService    = service.ServiceGroupApp.DatasourceServiceGroup.FileSourceService
	FFieldService     = service.ServiceGroupApp.DatasourceServiceGroup.FileFieldService
)
