
package pdfme

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pdfme"
    pdfmeReq "github.com/flipped-aurora/gin-vue-admin/server/model/pdfme/request"
)

type PdfmeDesignService struct {}
// CreatePdfmeDesign 创建打印设计记录
// Author [yourname](https://github.com/yourname)
func (PMDesignService *PdfmeDesignService) CreatePdfmeDesign(ctx context.Context, PMDesign *pdfme.PdfmeDesign) (err error) {
	err = global.GVA_DB.Create(PMDesign).Error
	return err
}

// DeletePdfmeDesign 删除打印设计记录
// Author [yourname](https://github.com/yourname)
func (PMDesignService *PdfmeDesignService)DeletePdfmeDesign(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&pdfme.PdfmeDesign{},"id = ?",ID).Error
	return err
}

// DeletePdfmeDesignByIds 批量删除打印设计记录
// Author [yourname](https://github.com/yourname)
func (PMDesignService *PdfmeDesignService)DeletePdfmeDesignByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]pdfme.PdfmeDesign{},"id in ?",IDs).Error
	return err
}

// UpdatePdfmeDesign 更新打印设计记录
// Author [yourname](https://github.com/yourname)
func (PMDesignService *PdfmeDesignService)UpdatePdfmeDesign(ctx context.Context, PMDesign pdfme.PdfmeDesign) (err error) {
	err = global.GVA_DB.Model(&pdfme.PdfmeDesign{}).Where("id = ?",PMDesign.ID).Updates(&PMDesign).Error
	return err
}

// GetPdfmeDesign 根据ID获取打印设计记录
// Author [yourname](https://github.com/yourname)
func (PMDesignService *PdfmeDesignService)GetPdfmeDesign(ctx context.Context, ID string) (PMDesign pdfme.PdfmeDesign, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&PMDesign).Error
	return
}
// GetPdfmeDesignInfoList 分页获取打印设计记录
// Author [yourname](https://github.com/yourname)
func (PMDesignService *PdfmeDesignService)GetPdfmeDesignInfoList(ctx context.Context, info pdfmeReq.PdfmeDesignSearch) (list []pdfme.PdfmeDesign, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pdfme.PdfmeDesign{})
    var PMDesigns []pdfme.PdfmeDesign
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

	err = db.Find(&PMDesigns).Error
	return  PMDesigns, total, err
}
func (PMDesignService *PdfmeDesignService)GetPdfmeDesignPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
