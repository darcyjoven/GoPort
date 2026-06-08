package datasource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataSourceFieldRouter struct {}

// InitDataSourceFieldRouter 初始化 数据源字段信息 路由信息
func (s *DataSourceFieldRouter) InitDataSourceFieldRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	DSFieldRouter := Router.Group("DSField").Use(middleware.OperationRecord())
	DSFieldRouterWithoutRecord := Router.Group("DSField")
	DSFieldRouterWithoutAuth := PublicRouter.Group("DSField")
	{
		DSFieldRouter.POST("createDataSourceField", DSFieldApi.CreateDataSourceField)   // 新建数据源字段信息
		DSFieldRouter.DELETE("deleteDataSourceField", DSFieldApi.DeleteDataSourceField) // 删除数据源字段信息
		DSFieldRouter.DELETE("deleteDataSourceFieldByIds", DSFieldApi.DeleteDataSourceFieldByIds) // 批量删除数据源字段信息
		DSFieldRouter.PUT("updateDataSourceField", DSFieldApi.UpdateDataSourceField)    // 更新数据源字段信息
	}
	{
		DSFieldRouterWithoutRecord.GET("findDataSourceField", DSFieldApi.FindDataSourceField)        // 根据ID获取数据源字段信息
		DSFieldRouterWithoutRecord.GET("getDataSourceFieldList", DSFieldApi.GetDataSourceFieldList)  // 获取数据源字段信息列表
	}
	{
	    DSFieldRouterWithoutAuth.GET("getDataSourceFieldDataSource", DSFieldApi.GetDataSourceFieldDataSource)  // 获取数据源字段信息数据源
	    DSFieldRouterWithoutAuth.GET("getDataSourceFieldPublic", DSFieldApi.GetDataSourceFieldPublic)  // 数据源字段信息开放接口
	}
}
