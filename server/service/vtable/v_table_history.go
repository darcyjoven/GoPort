
package vtable

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/vtable"
    vtableReq "github.com/flipped-aurora/gin-vue-admin/server/model/vtable/request"
    "gorm.io/gorm"
)

type VTableHistoryService struct {}
// CreateVTableHistory 创建表格设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (VTHistoryService *VTableHistoryService) CreateVTableHistory(ctx context.Context, VTHistory *vtable.VTableHistory) (err error) {
	err = global.GVA_DB.Create(VTHistory).Error
	return err
}

// DeleteVTableHistory 删除表格设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (VTHistoryService *VTableHistoryService)DeleteVTableHistory(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&vtable.VTableHistory{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&vtable.VTableHistory{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteVTableHistoryByIds 批量删除表格设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (VTHistoryService *VTableHistoryService)DeleteVTableHistoryByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&vtable.VTableHistory{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&vtable.VTableHistory{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateVTableHistory 更新表格设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (VTHistoryService *VTableHistoryService)UpdateVTableHistory(ctx context.Context, VTHistory vtable.VTableHistory) (err error) {
	err = global.GVA_DB.Model(&vtable.VTableHistory{}).Where("id = ?",VTHistory.ID).Updates(&VTHistory).Error
	return err
}

// GetVTableHistory 根据ID获取表格设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (VTHistoryService *VTableHistoryService)GetVTableHistory(ctx context.Context, ID string) (VTHistory vtable.VTableHistory, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&VTHistory).Error
	return
}
// GetVTableHistoryInfoList 分页获取表格设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (VTHistoryService *VTableHistoryService)GetVTableHistoryInfoList(ctx context.Context, info vtableReq.VTableHistorySearch) (list []vtable.VTableHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&vtable.VTableHistory{})
    var VTHistorys []vtable.VTableHistory
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Version != nil && *info.Version != "" {
        db = db.Where("version = ?", *info.Version)
    }
    if info.SourceID != nil {
        db = db.Where("source_i_d = ?", *info.SourceID)
    }
    if info.Remark != nil && *info.Remark != "" {
        db = db.Where("remark LIKE ?", "%"+ *info.Remark+"%")
    }
    if info.Active != nil {
        db = db.Where("active = ?", *info.Active)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&VTHistorys).Error
	return  VTHistorys, total, err
}
func (VTHistoryService *VTableHistoryService)GetVTableHistoryDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   sourceID := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("vtable_degisn").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&sourceID)
	   res["sourceID"] = sourceID
	return
}
func (VTHistoryService *VTableHistoryService)GetVTableHistoryPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
