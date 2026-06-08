package datasource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileSourceRouter struct {}

// InitFileSourceRouter 初始化 文件源 路由信息
func (s *FileSourceRouter) InitFileSourceRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	FSourceRouter := Router.Group("FSource").Use(middleware.OperationRecord())
	FSourceRouterWithoutRecord := Router.Group("FSource")
	FSourceRouterWithoutAuth := PublicRouter.Group("FSource")
	{
		FSourceRouter.POST("createFileSource", FSourceApi.CreateFileSource)   // 新建文件源
		FSourceRouter.DELETE("deleteFileSource", FSourceApi.DeleteFileSource) // 删除文件源
		FSourceRouter.DELETE("deleteFileSourceByIds", FSourceApi.DeleteFileSourceByIds) // 批量删除文件源
		FSourceRouter.PUT("updateFileSource", FSourceApi.UpdateFileSource)    // 更新文件源
	}
	{
		FSourceRouterWithoutRecord.GET("findFileSource", FSourceApi.FindFileSource)        // 根据ID获取文件源
		FSourceRouterWithoutRecord.GET("getFileSourceList", FSourceApi.GetFileSourceList)  // 获取文件源列表
	}
	{
	    FSourceRouterWithoutAuth.GET("getFileSourcePublic", FSourceApi.GetFileSourcePublic)  // 文件源开放接口
	}
}
