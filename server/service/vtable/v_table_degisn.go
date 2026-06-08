
package vtable

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/vtable"
    vtableReq "github.com/flipped-aurora/gin-vue-admin/server/model/vtable/request"
)

type VTableDesignService struct {}
// CreateVTableDesign 创建表格设计记录
// Author [yourname](https://github.com/yourname)
func (VTDesignService *VTableDesignService) CreateVTableDesign(ctx context.Context, VTDesign *vtable.VTableDesign) (err error) {
	err = global.GVA_DB.Create(VTDesign).Error
	return err
}

// DeleteVTableDesign 删除表格设计记录
// Author [yourname](https://github.com/yourname)
func (VTDesignService *VTableDesignService)DeleteVTableDesign(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&vtable.VTableDesign{},"id = ?",ID).Error
	return err
}

// DeleteVTableDesignByIds 批量删除表格设计记录
// Author [yourname](https://github.com/yourname)
func (VTDesignService *VTableDesignService)DeleteVTableDesignByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]vtable.VTableDesign{},"id in ?",IDs).Error
	return err
}

// UpdateVTableDesign 更新表格设计记录
// Author [yourname](https://github.com/yourname)
func (VTDesignService *VTableDesignService)UpdateVTableDesign(ctx context.Context, VTDesign vtable.VTableDesign) (err error) {
	err = global.GVA_DB.Model(&vtable.VTableDesign{}).Where("id = ?",VTDesign.ID).Updates(&VTDesign).Error
	return err
}

// GetVTableDesign 根据ID获取表格设计记录
// Author [yourname](https://github.com/yourname)
func (VTDesignService *VTableDesignService)GetVTableDesign(ctx context.Context, ID string) (VTDesign vtable.VTableDesign, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&VTDesign).Error
	return
}
// GetVTableDesignInfoList 分页获取表格设计记录
// Author [yourname](https://github.com/yourname)
func (VTDesignService *VTableDesignService)GetVTableDesignInfoList(ctx context.Context, info vtableReq.VTableDesignSearch) (list []vtable.VTableDesign, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&vtable.VTableDesign{})
    var VTDesigns []vtable.VTableDesign
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Desciption != nil && *info.Desciption != "" {
        db = db.Where("desciption LIKE ?", "%"+ *info.Desciption+"%")
    }
    if info.Module != nil && *info.Module != "" {
        db = db.Where("module = ?", *info.Module)
    }
    if info.CurrentVersion != nil && *info.CurrentVersion != "" {
        db = db.Where("current_version = ?", *info.CurrentVersion)
    }
    if info.Checkout != nil {
        db = db.Where("checkout = ?", *info.Checkout)
    }
    if info.CheckoutVersion != nil && *info.CheckoutVersion != "" {
        db = db.Where("checkout_version = ?", *info.CheckoutVersion)
    }
			if len(info.LastCheckOutTimeRange) == 2 {
				db = db.Where("last_check_out_time BETWEEN ? AND ? ", info.LastCheckOutTimeRange[0], info.LastCheckOutTimeRange[1])
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
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["id"] = true
           orderMap["created_at"] = true
         	orderMap["name"] = true
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

	err = db.Find(&VTDesigns).Error
	return  VTDesigns, total, err
}
func (VTDesignService *VTableDesignService)GetVTableDesignPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
