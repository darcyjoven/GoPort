package vtable

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VTableDesignRouter struct {}

// InitVTableDesignRouter 初始化 表格设计 路由信息
func (s *VTableDesignRouter) InitVTableDesignRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	VTDesignRouter := Router.Group("VTDesign").Use(middleware.OperationRecord())
	VTDesignRouterWithoutRecord := Router.Group("VTDesign")
	VTDesignRouterWithoutAuth := PublicRouter.Group("VTDesign")
	{
		VTDesignRouter.POST("createVTableDesign", VTDesignApi.CreateVTableDesign)   // 新建表格设计
		VTDesignRouter.DELETE("deleteVTableDesign", VTDesignApi.DeleteVTableDesign) // 删除表格设计
		VTDesignRouter.DELETE("deleteVTableDesignByIds", VTDesignApi.DeleteVTableDesignByIds) // 批量删除表格设计
		VTDesignRouter.PUT("updateVTableDesign", VTDesignApi.UpdateVTableDesign)    // 更新表格设计
	}
	{
		VTDesignRouterWithoutRecord.GET("findVTableDesign", VTDesignApi.FindVTableDesign)        // 根据ID获取表格设计
		VTDesignRouterWithoutRecord.GET("getVTableDesignList", VTDesignApi.GetVTableDesignList)  // 获取表格设计列表
	}
	{
	    VTDesignRouterWithoutAuth.GET("getVTableDesignPublic", VTDesignApi.GetVTableDesignPublic)  // 表格设计开放接口
	}
}
