
package datasource

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
    datasourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/datasource/request"
    "gorm.io/gorm"
)

type FileSourceService struct {}
// CreateFileSource 创建文件源记录
// Author [yourname](https://github.com/yourname)
func (FSourceService *FileSourceService) CreateFileSource(ctx context.Context, FSource *datasource.FileSource) (err error) {
	err = global.GVA_DB.Create(FSource).Error
	return err
}

// DeleteFileSource 删除文件源记录
// Author [yourname](https://github.com/yourname)
func (FSourceService *FileSourceService)DeleteFileSource(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&datasource.FileSource{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&datasource.FileSource{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteFileSourceByIds 批量删除文件源记录
// Author [yourname](https://github.com/yourname)
func (FSourceService *FileSourceService)DeleteFileSourceByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&datasource.FileSource{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&datasource.FileSource{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateFileSource 更新文件源记录
// Author [yourname](https://github.com/yourname)
func (FSourceService *FileSourceService)UpdateFileSource(ctx context.Context, FSource datasource.FileSource) (err error) {
	err = global.GVA_DB.Model(&datasource.FileSource{}).Where("id = ?",FSource.ID).Updates(&FSource).Error
	return err
}

// GetFileSource 根据ID获取文件源记录
// Author [yourname](https://github.com/yourname)
func (FSourceService *FileSourceService)GetFileSource(ctx context.Context, ID string) (FSource datasource.FileSource, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&FSource).Error
	return
}
// GetFileSourceInfoList 分页获取文件源记录
// Author [yourname](https://github.com/yourname)
func (FSourceService *FileSourceService)GetFileSourceInfoList(ctx context.Context, info datasourceReq.FileSourceSearch) (list []datasource.FileSource, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&datasource.FileSource{})
    var FSources []datasource.FileSource
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Description != nil && *info.Description != "" {
        db = db.Where("description LIKE ?", "%"+ *info.Description+"%")
    }
    if info.Remark != nil && *info.Remark != "" {
        db = db.Where("remark LIKE ?", "%"+ *info.Remark+"%")
    }
    if info.FileType != nil && *info.FileType != "" {
        db = db.Where("file_type = ?", *info.FileType)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&FSources).Error
	return  FSources, total, err
}
func (FSourceService *FileSourceService)GetFileSourcePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
