package datasource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DatabaseConfigRouter struct {}

// InitDatabaseConfigRouter 初始化 数据库连接配置 路由信息
func (s *DatabaseConfigRouter) InitDatabaseConfigRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	dbConfigRouter := Router.Group("dbConfig").Use(middleware.OperationRecord())
	dbConfigRouterWithoutRecord := Router.Group("dbConfig")
	dbConfigRouterWithoutAuth := PublicRouter.Group("dbConfig")
	{
		dbConfigRouter.POST("createDatabaseConfig", dbConfigApi.CreateDatabaseConfig)   // 新建数据库连接配置
		dbConfigRouter.DELETE("deleteDatabaseConfig", dbConfigApi.DeleteDatabaseConfig) // 删除数据库连接配置
		dbConfigRouter.DELETE("deleteDatabaseConfigByIds", dbConfigApi.DeleteDatabaseConfigByIds) // 批量删除数据库连接配置
		dbConfigRouter.PUT("updateDatabaseConfig", dbConfigApi.UpdateDatabaseConfig)    // 更新数据库连接配置
	}
	{
		dbConfigRouterWithoutRecord.GET("findDatabaseConfig", dbConfigApi.FindDatabaseConfig)        // 根据ID获取数据库连接配置
		dbConfigRouterWithoutRecord.GET("getDatabaseConfigList", dbConfigApi.GetDatabaseConfigList)  // 获取数据库连接配置列表
	}
	{
	    dbConfigRouterWithoutAuth.GET("getDatabaseConfigPublic", dbConfigApi.GetDatabaseConfigPublic)  // 数据库连接配置开放接口
	}
}
