
package datasource

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
    datasourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/datasource/request"
    "gorm.io/gorm"
)

type FileFieldService struct {}
// CreateFileField 创建文件字段定义记录
// Author [yourname](https://github.com/yourname)
func (FFieldService *FileFieldService) CreateFileField(ctx context.Context, FField *datasource.FileField) (err error) {
	err = global.GVA_DB.Create(FField).Error
	return err
}

// DeleteFileField 删除文件字段定义记录
// Author [yourname](https://github.com/yourname)
func (FFieldService *FileFieldService)DeleteFileField(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&datasource.FileField{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&datasource.FileField{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteFileFieldByIds 批量删除文件字段定义记录
// Author [yourname](https://github.com/yourname)
func (FFieldService *FileFieldService)DeleteFileFieldByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&datasource.FileField{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&datasource.FileField{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateFileField 更新文件字段定义记录
// Author [yourname](https://github.com/yourname)
func (FFieldService *FileFieldService)UpdateFileField(ctx context.Context, FField datasource.FileField) (err error) {
	err = global.GVA_DB.Model(&datasource.FileField{}).Where("id = ?",FField.ID).Updates(&FField).Error
	return err
}

// GetFileField 根据ID获取文件字段定义记录
// Author [yourname](https://github.com/yourname)
func (FFieldService *FileFieldService)GetFileField(ctx context.Context, ID string) (FField datasource.FileField, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&FField).Error
	return
}
// GetFileFieldInfoList 分页获取文件字段定义记录
// Author [yourname](https://github.com/yourname)
func (FFieldService *FileFieldService)GetFileFieldInfoList(ctx context.Context, info datasourceReq.FileFieldSearch) (list []datasource.FileField, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&datasource.FileField{})
    var FFields []datasource.FileField
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.SourceID != nil {
        db = db.Where("source_id = ?", *info.SourceID)
    }
    if info.Index != nil {
        db = db.Where("index = ?", *info.Index)
    }
    if info.Key != nil && *info.Key != "" {
        db = db.Where("key LIKE ?", "%"+ *info.Key+"%")
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?", *info.Name)
    }
    if info.Description != nil && *info.Description != "" {
        db = db.Where("description LIKE ?", "%"+ *info.Description+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&FFields).Error
	return  FFields, total, err
}
func (FFieldService *FileFieldService)GetFileFieldDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   sourceID := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("file_source").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&sourceID)
	   res["sourceID"] = sourceID
	return
}
func (FFieldService *FileFieldService)GetFileFieldPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
