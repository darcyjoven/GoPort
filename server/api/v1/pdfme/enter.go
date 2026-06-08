package pdfme

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	PdfmeDesignApi
	PdfmeHistoryApi
}

var (
	PMDesignService  = service.ServiceGroupApp.PdfmeServiceGroup.PdfmeDesignService
	PMHistoryService = service.ServiceGroupApp.PdfmeServiceGroup.PdfmeHistoryService
)
