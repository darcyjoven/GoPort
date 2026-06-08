package pdfme

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	PdfmeDesignRouter
	PdfmeHistoryRouter
}

var (
	PMDesignApi  = api.ApiGroupApp.PdfmeApiGroup.PdfmeDesignApi
	PMHistoryApi = api.ApiGroupApp.PdfmeApiGroup.PdfmeHistoryApi
)
