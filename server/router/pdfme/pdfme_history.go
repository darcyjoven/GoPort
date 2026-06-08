package pdfme

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PdfmeHistoryRouter struct {}

// InitPdfmeHistoryRouter 初始化 打印设计历史资料 路由信息
func (s *PdfmeHistoryRouter) InitPdfmeHistoryRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	PMHistoryRouter := Router.Group("PMHistory").Use(middleware.OperationRecord())
	PMHistoryRouterWithoutRecord := Router.Group("PMHistory")
	PMHistoryRouterWithoutAuth := PublicRouter.Group("PMHistory")
	{
		PMHistoryRouter.POST("createPdfmeHistory", PMHistoryApi.CreatePdfmeHistory)   // 新建打印设计历史资料
		PMHistoryRouter.DELETE("deletePdfmeHistory", PMHistoryApi.DeletePdfmeHistory) // 删除打印设计历史资料
		PMHistoryRouter.DELETE("deletePdfmeHistoryByIds", PMHistoryApi.DeletePdfmeHistoryByIds) // 批量删除打印设计历史资料
		PMHistoryRouter.PUT("updatePdfmeHistory", PMHistoryApi.UpdatePdfmeHistory)    // 更新打印设计历史资料
	}
	{
		PMHistoryRouterWithoutRecord.GET("findPdfmeHistory", PMHistoryApi.FindPdfmeHistory)        // 根据ID获取打印设计历史资料
		PMHistoryRouterWithoutRecord.GET("getPdfmeHistoryList", PMHistoryApi.GetPdfmeHistoryList)  // 获取打印设计历史资料列表
	}
	{
	    PMHistoryRouterWithoutAuth.GET("getPdfmeHistoryDataSource", PMHistoryApi.GetPdfmeHistoryDataSource)  // 获取打印设计历史资料数据源
	    PMHistoryRouterWithoutAuth.GET("getPdfmeHistoryPublic", PMHistoryApi.GetPdfmeHistoryPublic)  // 打印设计历史资料开放接口
	}
}
