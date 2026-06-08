package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/datasource"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/pdfme"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/vtable"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	DatasourceApiGroup datasource.ApiGroup
	VtableApiGroup     vtable.ApiGroup
	PdfmeApiGroup      pdfme.ApiGroup
}
