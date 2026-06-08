package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/datasource"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/pdfme"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/vtable"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Datasource datasource.RouterGroup
	Vtable     vtable.RouterGroup
	Pdfme      pdfme.RouterGroup
}
