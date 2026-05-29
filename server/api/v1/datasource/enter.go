package datasource

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DatabaseConfigApi
	SearchTemplateApi
}

var (
	dbConfigService   = service.ServiceGroupApp.DatasourceServiceGroup.DatabaseConfigService
	searchTempService = service.ServiceGroupApp.DatasourceServiceGroup.SearchTemplateService
)
