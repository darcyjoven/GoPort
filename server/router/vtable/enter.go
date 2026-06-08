package vtable

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	VTableDesignRouter
	VTableHistoryRouter
}

var (
	VTDesignApi  = api.ApiGroupApp.VtableApiGroup.VTableDesignApi
	VTHistoryApi = api.ApiGroupApp.VtableApiGroup.VTableHistoryApi
)
