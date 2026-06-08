
package datasource

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
    datasourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/datasource/request"
    "gorm.io/gorm"
)

type APISourceService struct {}
// CreateAPISource 创建api配置记录
// Author [yourname](https://github.com/yourname)
func (ASourceService *APISourceService) CreateAPISource(ctx context.Context, ASource *datasource.APISource) (err error) {
	err = global.GVA_DB.Create(ASource).Error
	return err
}

// DeleteAPISource 删除api配置记录
// Author [yourname](https://github.com/yourname)
func (ASourceService *APISourceService)DeleteAPISource(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&datasource.APISource{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&datasource.APISource{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAPISourceByIds 批量删除api配置记录
// Author [yourname](https://github.com/yourname)
func (ASourceService *APISourceService)DeleteAPISourceByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&datasource.APISource{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&datasource.APISource{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAPISource 更新api配置记录
// Author [yourname](https://github.com/yourname)
func (ASourceService *APISourceService)UpdateAPISource(ctx context.Context, ASource datasource.APISource) (err error) {
	err = global.GVA_DB.Model(&datasource.APISource{}).Where("id = ?",ASource.ID).Updates(&ASource).Error
	return err
}

// GetAPISource 根据ID获取api配置记录
// Author [yourname](https://github.com/yourname)
func (ASourceService *APISourceService)GetAPISource(ctx context.Context, ID string) (ASource datasource.APISource, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ASource).Error
	return
}
// GetAPISourceInfoList 分页获取api配置记录
// Author [yourname](https://github.com/yourname)
func (ASourceService *APISourceService)GetAPISourceInfoList(ctx context.Context, info datasourceReq.APISourceSearch) (list []datasource.APISource, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&datasource.APISource{})
    var ASources []datasource.APISource
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Path != nil && *info.Path != "" {
        db = db.Where("path LIKE ?", "%"+ *info.Path+"%")
    }
    if info.Remark != nil && *info.Remark != "" {
        db = db.Where("remark LIKE ?", "%"+ *info.Remark+"%")
    }
    if info.Argv1 != nil && *info.Argv1 != "" {
        db = db.Where("argv1 LIKE ?", "%"+ *info.Argv1+"%")
    }
    if info.Argv2 != nil && *info.Argv2 != "" {
        db = db.Where("argv2 LIKE ?", "%"+ *info.Argv2+"%")
    }
    if info.Argv3 != nil && *info.Argv3 != "" {
        db = db.Where("argv3 LIKE ?", "%"+ *info.Argv3+"%")
    }
    if info.Argv4 != nil && *info.Argv4 != "" {
        db = db.Where("argv4 LIKE ?", "%"+ *info.Argv4+"%")
    }
    if info.Argv5 != nil && *info.Argv5 != "" {
        db = db.Where("argv5 LIKE ?", "%"+ *info.Argv5+"%")
    }
    if info.Argv6 != nil && *info.Argv6 != "" {
        db = db.Where("argv6 LIKE ?", "%"+ *info.Argv6+"%")
    }
    if info.Argv7 != nil && *info.Argv7 != "" {
        db = db.Where("argv7 LIKE ?", "%"+ *info.Argv7+"%")
    }
    if info.Argv8 != nil && *info.Argv8 != "" {
        db = db.Where("argv8 LIKE ?", "%"+ *info.Argv8+"%")
    }
    if info.Argv9 != nil && *info.Argv9 != "" {
        db = db.Where("argv9 LIKE ?", "%"+ *info.Argv9+"%")
    }
    if info.Argv10 != nil && *info.Argv10 != "" {
        db = db.Where("argv10 LIKE ?", "%"+ *info.Argv10+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&ASources).Error
	return  ASources, total, err
}
func (ASourceService *APISourceService)GetAPISourcePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
