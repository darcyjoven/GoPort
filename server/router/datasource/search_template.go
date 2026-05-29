package datasource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SearchTemplateRouter struct {}

// InitSearchTemplateRouter 初始化 查询SQL模板 路由信息
func (s *SearchTemplateRouter) InitSearchTemplateRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	searchTempRouter := Router.Group("searchTemp").Use(middleware.OperationRecord())
	searchTempRouterWithoutRecord := Router.Group("searchTemp")
	searchTempRouterWithoutAuth := PublicRouter.Group("searchTemp")
	{
		searchTempRouter.POST("createSearchTemplate", searchTempApi.CreateSearchTemplate)   // 新建查询SQL模板
		searchTempRouter.DELETE("deleteSearchTemplate", searchTempApi.DeleteSearchTemplate) // 删除查询SQL模板
		searchTempRouter.DELETE("deleteSearchTemplateByIds", searchTempApi.DeleteSearchTemplateByIds) // 批量删除查询SQL模板
		searchTempRouter.PUT("updateSearchTemplate", searchTempApi.UpdateSearchTemplate)    // 更新查询SQL模板
	}
	{
		searchTempRouterWithoutRecord.GET("findSearchTemplate", searchTempApi.FindSearchTemplate)        // 根据ID获取查询SQL模板
		searchTempRouterWithoutRecord.GET("getSearchTemplateList", searchTempApi.GetSearchTemplateList)  // 获取查询SQL模板列表
	}
	{
	    searchTempRouterWithoutAuth.GET("getSearchTemplateDataSource", searchTempApi.GetSearchTemplateDataSource)  // 获取查询SQL模板数据源
	    searchTempRouterWithoutAuth.GET("getSearchTemplatePublic", searchTempApi.GetSearchTemplatePublic)  // 查询SQL模板开放接口
	}
}
