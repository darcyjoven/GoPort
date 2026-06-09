package datasource

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DataSourceFieldApi
	APISourceApi
	FileSourceApi
	FileFieldApi
	SearchTemplateApi
}

var (
	DSFieldService    = service.ServiceGroupApp.DatasourceServiceGroup.DataSourceFieldService
	ASourceService    = service.ServiceGroupApp.DatasourceServiceGroup.APISourceService
	FSourceService    = service.ServiceGroupApp.DatasourceServiceGroup.FileSourceService
	FFieldService     = service.ServiceGroupApp.DatasourceServiceGroup.FileFieldService
	searchTempService = service.ServiceGroupApp.DatasourceServiceGroup.SearchTemplateService
)
