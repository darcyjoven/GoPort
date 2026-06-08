package datasource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type APISourceRouter struct {}

// InitAPISourceRouter 初始化 api配置 路由信息
func (s *APISourceRouter) InitAPISourceRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	ASourceRouter := Router.Group("ASource").Use(middleware.OperationRecord())
	ASourceRouterWithoutRecord := Router.Group("ASource")
	ASourceRouterWithoutAuth := PublicRouter.Group("ASource")
	{
		ASourceRouter.POST("createAPISource", ASourceApi.CreateAPISource)   // 新建api配置
		ASourceRouter.DELETE("deleteAPISource", ASourceApi.DeleteAPISource) // 删除api配置
		ASourceRouter.DELETE("deleteAPISourceByIds", ASourceApi.DeleteAPISourceByIds) // 批量删除api配置
		ASourceRouter.PUT("updateAPISource", ASourceApi.UpdateAPISource)    // 更新api配置
	}
	{
		ASourceRouterWithoutRecord.GET("findAPISource", ASourceApi.FindAPISource)        // 根据ID获取api配置
		ASourceRouterWithoutRecord.GET("getAPISourceList", ASourceApi.GetAPISourceList)  // 获取api配置列表
	}
	{
	    ASourceRouterWithoutAuth.GET("getAPISourcePublic", ASourceApi.GetAPISourcePublic)  // api配置开放接口
	}
}
