
package datasource

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
    datasourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/datasource/request"
    "gorm.io/gorm"
)

type DataSourceFieldService struct {}
// CreateDataSourceField 创建数据源字段信息记录
// Author [yourname](https://github.com/yourname)
func (DSFieldService *DataSourceFieldService) CreateDataSourceField(ctx context.Context, DSField *datasource.DataSourceField) (err error) {
	err = global.GVA_DB.Create(DSField).Error
	return err
}

// DeleteDataSourceField 删除数据源字段信息记录
// Author [yourname](https://github.com/yourname)
func (DSFieldService *DataSourceFieldService)DeleteDataSourceField(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&datasource.DataSourceField{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&datasource.DataSourceField{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteDataSourceFieldByIds 批量删除数据源字段信息记录
// Author [yourname](https://github.com/yourname)
func (DSFieldService *DataSourceFieldService)DeleteDataSourceFieldByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&datasource.DataSourceField{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&datasource.DataSourceField{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateDataSourceField 更新数据源字段信息记录
// Author [yourname](https://github.com/yourname)
func (DSFieldService *DataSourceFieldService)UpdateDataSourceField(ctx context.Context, DSField datasource.DataSourceField) (err error) {
	err = global.GVA_DB.Model(&datasource.DataSourceField{}).Where("id = ?",DSField.ID).Updates(&DSField).Error
	return err
}

// GetDataSourceField 根据ID获取数据源字段信息记录
// Author [yourname](https://github.com/yourname)
func (DSFieldService *DataSourceFieldService)GetDataSourceField(ctx context.Context, ID string) (DSField datasource.DataSourceField, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&DSField).Error
	return
}
// GetDataSourceFieldInfoList 分页获取数据源字段信息记录
// Author [yourname](https://github.com/yourname)
func (DSFieldService *DataSourceFieldService)GetDataSourceFieldInfoList(ctx context.Context, info datasourceReq.DataSourceFieldSearch) (list []datasource.DataSourceField, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&datasource.DataSourceField{})
    var DSFields []datasource.DataSourceField
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.SourceType != nil && *info.SourceType != "" {
        db = db.Where("source_type = ?", *info.SourceType)
    }
    if info.FieldIndex != nil {
        db = db.Where("field_index = ?", *info.FieldIndex)
    }
    if info.FiledType != nil && *info.FiledType != "" {
        db = db.Where("filed_type = ?", *info.FiledType)
    }
    if info.FieldName != nil && *info.FieldName != "" {
        db = db.Where("field_name = ?", *info.FieldName)
    }
    if info.Description != nil && *info.Description != "" {
        db = db.Where("description LIKE ?", "%"+ *info.Description+"%")
    }
    if info.Sortable != nil && *info.Sortable != "" {
        db = db.Where("sortable = ?", *info.Sortable)
    }
    if info.Width != nil {
        db = db.Where("width = ?", *info.Width)
    }
    if info.Format != nil && *info.Format != "" {
        db = db.Where("format LIKE ?", "%"+ *info.Format+"%")
    }
    if info.Warp != nil {
        db = db.Where("warp = ?", *info.Warp)
    }
    if info.Align != nil && *info.Align != "" {
        db = db.Where("align = ?", *info.Align)
    }
    if info.Extra != nil && *info.Extra != "" {
        db = db.Where("extra LIKE ?", "%"+ *info.Extra+"%")
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
         	orderMap["source_type"] = true
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

	err = db.Find(&DSFields).Error
	return  DSFields, total, err
}
func (DSFieldService *DataSourceFieldService)GetDataSourceFieldDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   sourceID := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("search_template").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&sourceID)
	   res["sourceID"] = sourceID
	return
}
func (DSFieldService *DataSourceFieldService)GetDataSourceFieldPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
