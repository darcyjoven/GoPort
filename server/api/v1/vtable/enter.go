package vtable

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	VTableDesignApi
	VTableHistoryApi
}

var (
	VTDesignService  = service.ServiceGroupApp.VtableServiceGroup.VTableDesignService
	VTHistoryService = service.ServiceGroupApp.VtableServiceGroup.VTableHistoryService
)
