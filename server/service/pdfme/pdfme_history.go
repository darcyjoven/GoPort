
package pdfme

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pdfme"
    pdfmeReq "github.com/flipped-aurora/gin-vue-admin/server/model/pdfme/request"
    "gorm.io/gorm"
)

type PdfmeHistoryService struct {}
// CreatePdfmeHistory 创建打印设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (PMHistoryService *PdfmeHistoryService) CreatePdfmeHistory(ctx context.Context, PMHistory *pdfme.PdfmeHistory) (err error) {
	err = global.GVA_DB.Create(PMHistory).Error
	return err
}

// DeletePdfmeHistory 删除打印设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (PMHistoryService *PdfmeHistoryService)DeletePdfmeHistory(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&pdfme.PdfmeHistory{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&pdfme.PdfmeHistory{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeletePdfmeHistoryByIds 批量删除打印设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (PMHistoryService *PdfmeHistoryService)DeletePdfmeHistoryByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&pdfme.PdfmeHistory{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&pdfme.PdfmeHistory{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdatePdfmeHistory 更新打印设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (PMHistoryService *PdfmeHistoryService)UpdatePdfmeHistory(ctx context.Context, PMHistory pdfme.PdfmeHistory) (err error) {
	err = global.GVA_DB.Model(&pdfme.PdfmeHistory{}).Where("id = ?",PMHistory.ID).Updates(&PMHistory).Error
	return err
}

// GetPdfmeHistory 根据ID获取打印设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (PMHistoryService *PdfmeHistoryService)GetPdfmeHistory(ctx context.Context, ID string) (PMHistory pdfme.PdfmeHistory, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&PMHistory).Error
	return
}
// GetPdfmeHistoryInfoList 分页获取打印设计历史资料记录
// Author [yourname](https://github.com/yourname)
func (PMHistoryService *PdfmeHistoryService)GetPdfmeHistoryInfoList(ctx context.Context, info pdfmeReq.PdfmeHistorySearch) (list []pdfme.PdfmeHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pdfme.PdfmeHistory{})
    var PMHistorys []pdfme.PdfmeHistory
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

	err = db.Find(&PMHistorys).Error
	return  PMHistorys, total, err
}
func (PMHistoryService *PdfmeHistoryService)GetPdfmeHistoryDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   sourceID := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("vtable_degisn").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&sourceID)
	   res["sourceID"] = sourceID
	return
}
func (PMHistoryService *PdfmeHistoryService)GetPdfmeHistoryPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
