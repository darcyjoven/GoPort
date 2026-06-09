package datasource

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
	datasourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/datasource/request"
)

type SearchTemplateService struct{}

// CreateSearchTemplate 创建查询SQL模板记录
// Author [yourname](https://github.com/yourname)
func (searchTempService *SearchTemplateService) CreateSearchTemplate(ctx context.Context, searchTemp *datasource.SearchTemplate) (err error) {
	err = global.GVA_DB.Create(searchTemp).Error
	return err
}

// DeleteSearchTemplate 删除查询SQL模板记录
// Author [yourname](https://github.com/yourname)
func (searchTempService *SearchTemplateService) DeleteSearchTemplate(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&datasource.SearchTemplate{}, "id = ?", ID).Error
	return err
}

// DeleteSearchTemplateByIds 批量删除查询SQL模板记录
// Author [yourname](https://github.com/yourname)
func (searchTempService *SearchTemplateService) DeleteSearchTemplateByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]datasource.SearchTemplate{}, "id in ?", IDs).Error
	return err
}

// UpdateSearchTemplate 更新查询SQL模板记录
// Author [yourname](https://github.com/yourname)
func (searchTempService *SearchTemplateService) UpdateSearchTemplate(ctx context.Context, searchTemp datasource.SearchTemplate) (err error) {
	err = global.GVA_DB.Model(&datasource.SearchTemplate{}).Where("id = ?", searchTemp.ID).Updates(&searchTemp).Error
	return err
}

// GetSearchTemplate 根据ID获取查询SQL模板记录
// Author [yourname](https://github.com/yourname)
func (searchTempService *SearchTemplateService) GetSearchTemplate(ctx context.Context, ID string) (searchTemp datasource.SearchTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&searchTemp).Error
	return
}

// GetSearchTemplateInfoList 分页获取查询SQL模板记录
// Author [yourname](https://github.com/yourname)
func (searchTempService *SearchTemplateService) GetSearchTemplateInfoList(ctx context.Context, info datasourceReq.SearchTemplateSearch) (list []datasource.SearchTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&datasource.SearchTemplate{})
	var searchTemps []datasource.SearchTemplate
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.DatabaseID != nil {
		db = db.Where("database_id = ?", *info.DatabaseID)
	}
	if info.SearchText != nil && *info.SearchText != "" {
		db = db.Where("search_text LIKE ?", "%"+*info.SearchText+"%")
	}
	if info.Remark != nil && *info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+*info.Remark+"%")
	}
	if info.Argv1 != nil && *info.Argv1 != "" {
		db = db.Where("argv1 LIKE ?", "%"+*info.Argv1+"%")
	}
	if info.Argv2 != nil && *info.Argv2 != "" {
		db = db.Where("argv2 LIKE ?", "%"+*info.Argv2+"%")
	}
	if info.Argv3 != nil && *info.Argv3 != "" {
		db = db.Where("argv3 LIKE ?", "%"+*info.Argv3+"%")
	}
	if info.Argv4 != nil && *info.Argv4 != "" {
		db = db.Where("argv4 LIKE ?", "%"+*info.Argv4+"%")
	}
	if info.Argv5 != nil && *info.Argv5 != "" {
		db = db.Where("argv5 LIKE ?", "%"+*info.Argv5+"%")
	}
	if info.Argv6 != nil && *info.Argv6 != "" {
		db = db.Where("argv6 LIKE ?", "%"+*info.Argv6+"%")
	}
	if info.Argv7 != nil && *info.Argv7 != "" {
		db = db.Where("argv7 LIKE ?", "%"+*info.Argv7+"%")
	}
	if info.Argv8 != nil && *info.Argv8 != "" {
		db = db.Where("argv8 LIKE ?", "%"+*info.Argv8+"%")
	}
	if info.Argv9 != nil && *info.Argv9 != "" {
		db = db.Where("argv9 LIKE ?", "%"+*info.Argv9+"%")
	}
	if info.Argv10 != nil && *info.Argv10 != "" {
		db = db.Where("argv10 LIKE ?", "%"+*info.Argv10+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&searchTemps).Error
	return searchTemps, total, err
}
func (searchTempService *SearchTemplateService) GetSearchTemplateDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	databaseID := make([]map[string]any, 0)

	global.GVA_DB.Table("database_config").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&databaseID)
	res["databaseID"] = databaseID
	return
}
func (searchTempService *SearchTemplateService) GetSearchTemplatePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// 查询ima_file示例接口 SearchIma
func (searchTempService *SearchTemplateService) SearchIma(ctx context.Context) (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&datasource.SearchTemplate{})
	return db.Error
}
