
package datasource

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
    datasourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/datasource/request"
)

type DatabaseConfigService struct {}
// CreateDatabaseConfig 创建数据库连接配置记录
// Author [yourname](https://github.com/yourname)
func (dbConfigService *DatabaseConfigService) CreateDatabaseConfig(ctx context.Context, dbConfig *datasource.DatabaseConfig) (err error) {
	err = global.GVA_DB.Create(dbConfig).Error
	return err
}

// DeleteDatabaseConfig 删除数据库连接配置记录
// Author [yourname](https://github.com/yourname)
func (dbConfigService *DatabaseConfigService)DeleteDatabaseConfig(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&datasource.DatabaseConfig{},"id = ?",ID).Error
	return err
}

// DeleteDatabaseConfigByIds 批量删除数据库连接配置记录
// Author [yourname](https://github.com/yourname)
func (dbConfigService *DatabaseConfigService)DeleteDatabaseConfigByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]datasource.DatabaseConfig{},"id in ?",IDs).Error
	return err
}

// UpdateDatabaseConfig 更新数据库连接配置记录
// Author [yourname](https://github.com/yourname)
func (dbConfigService *DatabaseConfigService)UpdateDatabaseConfig(ctx context.Context, dbConfig datasource.DatabaseConfig) (err error) {
	err = global.GVA_DB.Model(&datasource.DatabaseConfig{}).Where("id = ?",dbConfig.ID).Updates(&dbConfig).Error
	return err
}

// GetDatabaseConfig 根据ID获取数据库连接配置记录
// Author [yourname](https://github.com/yourname)
func (dbConfigService *DatabaseConfigService)GetDatabaseConfig(ctx context.Context, ID string) (dbConfig datasource.DatabaseConfig, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dbConfig).Error
	return
}
// GetDatabaseConfigInfoList 分页获取数据库连接配置记录
// Author [yourname](https://github.com/yourname)
func (dbConfigService *DatabaseConfigService)GetDatabaseConfigInfoList(ctx context.Context, info datasourceReq.DatabaseConfigSearch) (list []datasource.DatabaseConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&datasource.DatabaseConfig{})
    var dbConfigs []datasource.DatabaseConfig
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.DbType != nil && *info.DbType != "" {
        db = db.Where("db_type = ?", *info.DbType)
    }
    if info.Host != nil && *info.Host != "" {
        db = db.Where("host LIKE ?", "%"+ *info.Host+"%")
    }
    if info.Port != nil {
        db = db.Where("port = ?", *info.Port)
    }
    if info.Server != nil && *info.Server != "" {
        db = db.Where("server = ?", *info.Server)
    }
    if info.Username != nil && *info.Username != "" {
        db = db.Where("username = ?", *info.Username)
    }
    if info.Password != nil && *info.Password != "" {
        db = db.Where("password = ?", *info.Password)
    }
    if info.Remark != nil && *info.Remark != "" {
        db = db.Where("remark LIKE ?", "%"+ *info.Remark+"%")
    }
    if info.Enable != nil {
        db = db.Where("enable = ?", *info.Enable)
    }
    if info.LastTestTime != nil {
        db = db.Where("last_test_time = ?", *info.LastTestTime)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["id"] = true
           orderMap["created_at"] = true
         	orderMap["name"] = true
         	orderMap["db_type"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&dbConfigs).Error
	return  dbConfigs, total, err
}
func (dbConfigService *DatabaseConfigService)GetDatabaseConfigPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
