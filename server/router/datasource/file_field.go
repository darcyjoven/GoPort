package datasource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileFieldRouter struct {}

// InitFileFieldRouter 初始化 文件字段定义 路由信息
func (s *FileFieldRouter) InitFileFieldRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	FFieldRouter := Router.Group("FField").Use(middleware.OperationRecord())
	FFieldRouterWithoutRecord := Router.Group("FField")
	FFieldRouterWithoutAuth := PublicRouter.Group("FField")
	{
		FFieldRouter.POST("createFileField", FFieldApi.CreateFileField)   // 新建文件字段定义
		FFieldRouter.DELETE("deleteFileField", FFieldApi.DeleteFileField) // 删除文件字段定义
		FFieldRouter.DELETE("deleteFileFieldByIds", FFieldApi.DeleteFileFieldByIds) // 批量删除文件字段定义
		FFieldRouter.PUT("updateFileField", FFieldApi.UpdateFileField)    // 更新文件字段定义
	}
	{
		FFieldRouterWithoutRecord.GET("findFileField", FFieldApi.FindFileField)        // 根据ID获取文件字段定义
		FFieldRouterWithoutRecord.GET("getFileFieldList", FFieldApi.GetFileFieldList)  // 获取文件字段定义列表
	}
	{
	    FFieldRouterWithoutAuth.GET("getFileFieldDataSource", FFieldApi.GetFileFieldDataSource)  // 获取文件字段定义数据源
	    FFieldRouterWithoutAuth.GET("getFileFieldPublic", FFieldApi.GetFileFieldPublic)  // 文件字段定义开放接口
	}
}
