package pdfme

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PdfmeDesignRouter struct {}

// InitPdfmeDesignRouter 初始化 打印设计 路由信息
func (s *PdfmeDesignRouter) InitPdfmeDesignRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	PMDesignRouter := Router.Group("PMDesign").Use(middleware.OperationRecord())
	PMDesignRouterWithoutRecord := Router.Group("PMDesign")
	PMDesignRouterWithoutAuth := PublicRouter.Group("PMDesign")
	{
		PMDesignRouter.POST("createPdfmeDesign", PMDesignApi.CreatePdfmeDesign)   // 新建打印设计
		PMDesignRouter.DELETE("deletePdfmeDesign", PMDesignApi.DeletePdfmeDesign) // 删除打印设计
		PMDesignRouter.DELETE("deletePdfmeDesignByIds", PMDesignApi.DeletePdfmeDesignByIds) // 批量删除打印设计
		PMDesignRouter.PUT("updatePdfmeDesign", PMDesignApi.UpdatePdfmeDesign)    // 更新打印设计
	}
	{
		PMDesignRouterWithoutRecord.GET("findPdfmeDesign", PMDesignApi.FindPdfmeDesign)        // 根据ID获取打印设计
		PMDesignRouterWithoutRecord.GET("getPdfmeDesignList", PMDesignApi.GetPdfmeDesignList)  // 获取打印设计列表
	}
	{
	    PMDesignRouterWithoutAuth.GET("getPdfmeDesignPublic", PMDesignApi.GetPdfmeDesignPublic)  // 打印设计开放接口
	}
}
