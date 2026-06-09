package datasource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SearchTemplateRouter struct{}

func (s *SearchTemplateRouter) InitSearchTemplateRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	searchTempRouter := Router.Group("searchTemp").Use(middleware.OperationRecord())
	searchTempRouterWithoutRecord := Router.Group("searchTemp")
	searchTempRouterWithoutAuth := PublicRouter.Group("searchTemp")
	{
		searchTempRouter.POST("createSearchTemplate", searchTempApi.CreateSearchTemplate)
		searchTempRouter.DELETE("deleteSearchTemplate", searchTempApi.DeleteSearchTemplate)
		searchTempRouter.DELETE("deleteSearchTemplateByIds", searchTempApi.DeleteSearchTemplateByIds)
		searchTempRouter.PUT("updateSearchTemplate", searchTempApi.UpdateSearchTemplate)
	}
	{
		searchTempRouterWithoutRecord.GET("findSearchTemplate", searchTempApi.FindSearchTemplate)
		searchTempRouterWithoutRecord.GET("getSearchTemplateList", searchTempApi.GetSearchTemplateList)
	}
	{
		searchTempRouterWithoutAuth.GET("getSearchTemplateDataSource", searchTempApi.GetSearchTemplateDataSource)
		searchTempRouterWithoutAuth.GET("getSearchTemplatePublic", searchTempApi.GetSearchTemplatePublic)
		searchTempRouterWithoutAuth.GET("searchIma", searchTempApi.SearchIma)
	}
}
