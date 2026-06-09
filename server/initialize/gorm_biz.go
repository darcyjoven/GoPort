package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pdfme"
	"github.com/flipped-aurora/gin-vue-admin/server/model/vtable"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(datasource.DataSourceField{}, vtable.VTableDesign{}, vtable.VTableHistory{}, pdfme.PdfmeDesign{}, pdfme.PdfmeHistory{}, datasource.APISource{}, datasource.FileSource{}, datasource.FileField{}, datasource.SearchTemplate{})
	if err != nil {
		return err
	}
	return nil
}
