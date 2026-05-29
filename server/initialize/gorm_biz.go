package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(datasource.DatabaseConfig{}, datasource.SearchTemplate{})
	if err != nil {
		return err
	}
	return nil
}
