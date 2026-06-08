package vtable

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VTableHistoryRouter struct {}

// InitVTableHistoryRouter 初始化 表格设计历史资料 路由信息
func (s *VTableHistoryRouter) InitVTableHistoryRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	VTHistoryRouter := Router.Group("VTHistory").Use(middleware.OperationRecord())
	VTHistoryRouterWithoutRecord := Router.Group("VTHistory")
	VTHistoryRouterWithoutAuth := PublicRouter.Group("VTHistory")
	{
		VTHistoryRouter.POST("createVTableHistory", VTHistoryApi.CreateVTableHistory)   // 新建表格设计历史资料
		VTHistoryRouter.DELETE("deleteVTableHistory", VTHistoryApi.DeleteVTableHistory) // 删除表格设计历史资料
		VTHistoryRouter.DELETE("deleteVTableHistoryByIds", VTHistoryApi.DeleteVTableHistoryByIds) // 批量删除表格设计历史资料
		VTHistoryRouter.PUT("updateVTableHistory", VTHistoryApi.UpdateVTableHistory)    // 更新表格设计历史资料
	}
	{
		VTHistoryRouterWithoutRecord.GET("findVTableHistory", VTHistoryApi.FindVTableHistory)        // 根据ID获取表格设计历史资料
		VTHistoryRouterWithoutRecord.GET("getVTableHistoryList", VTHistoryApi.GetVTableHistoryList)  // 获取表格设计历史资料列表
	}
	{
	    VTHistoryRouterWithoutAuth.GET("getVTableHistoryDataSource", VTHistoryApi.GetVTableHistoryDataSource)  // 获取表格设计历史资料数据源
	    VTHistoryRouterWithoutAuth.GET("getVTableHistoryPublic", VTHistoryApi.GetVTableHistoryPublic)  // 表格设计历史资料开放接口
	}
}
