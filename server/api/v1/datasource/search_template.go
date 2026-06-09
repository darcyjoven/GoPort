package datasource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
	datasourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/datasource/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SearchTemplateApi struct{}

// CreateSearchTemplate 创建查询SQL模板
// @Tags SearchTemplate
// @Summary 创建查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.SearchTemplate true "创建查询SQL模板"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /searchTemp/createSearchTemplate [post]
func (searchTempApi *SearchTemplateApi) CreateSearchTemplate(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var searchTemp datasource.SearchTemplate
	err := c.ShouldBindJSON(&searchTemp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = searchTempService.CreateSearchTemplate(ctx, &searchTemp)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSearchTemplate 删除查询SQL模板
// @Tags SearchTemplate
// @Summary 删除查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.SearchTemplate true "删除查询SQL模板"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /searchTemp/deleteSearchTemplate [delete]
func (searchTempApi *SearchTemplateApi) DeleteSearchTemplate(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := searchTempService.DeleteSearchTemplate(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSearchTemplateByIds 批量删除查询SQL模板
// @Tags SearchTemplate
// @Summary 批量删除查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /searchTemp/deleteSearchTemplateByIds [delete]
func (searchTempApi *SearchTemplateApi) DeleteSearchTemplateByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := searchTempService.DeleteSearchTemplateByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSearchTemplate 更新查询SQL模板
// @Tags SearchTemplate
// @Summary 更新查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.SearchTemplate true "更新查询SQL模板"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /searchTemp/updateSearchTemplate [put]
func (searchTempApi *SearchTemplateApi) UpdateSearchTemplate(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var searchTemp datasource.SearchTemplate
	err := c.ShouldBindJSON(&searchTemp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = searchTempService.UpdateSearchTemplate(ctx, searchTemp)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSearchTemplate 用id查询查询SQL模板
// @Tags SearchTemplate
// @Summary 用id查询查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询查询SQL模板"
// @Success 200 {object} response.Response{data=datasource.SearchTemplate,msg=string} "查询成功"
// @Router /searchTemp/findSearchTemplate [get]
func (searchTempApi *SearchTemplateApi) FindSearchTemplate(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	researchTemp, err := searchTempService.GetSearchTemplate(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(researchTemp, c)
}

// GetSearchTemplateList 分页获取查询SQL模板列表
// @Tags SearchTemplate
// @Summary 分页获取查询SQL模板列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.SearchTemplateSearch true "分页获取查询SQL模板列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /searchTemp/getSearchTemplateList [get]
func (searchTempApi *SearchTemplateApi) GetSearchTemplateList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo datasourceReq.SearchTemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := searchTempService.GetSearchTemplateInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetSearchTemplateDataSource 获取SearchTemplate的数据源
// @Tags SearchTemplate
// @Summary 获取SearchTemplate的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /searchTemp/getSearchTemplateDataSource [get]
func (searchTempApi *SearchTemplateApi) GetSearchTemplateDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := searchTempService.GetSearchTemplateDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetSearchTemplatePublic 不需要鉴权的查询SQL模板接口
// @Tags SearchTemplate
// @Summary 不需要鉴权的查询SQL模板接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /searchTemp/getSearchTemplatePublic [get]
func (searchTempApi *SearchTemplateApi) GetSearchTemplatePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	searchTempService.GetSearchTemplatePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的查询SQL模板接口信息",
	}, "获取成功", c)
}

// 查询ima_file示例接口 SearchIma
// @Tags SearchTemplate
// @Summary SearchIma
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /searchTemp/searchIma [GET]
func (searchTempApi *SearchTemplateApi) SearchIma(c *gin.Context) {
	ctx := c.Request.Context()

	// 获取Oracle数据库连接（使用别名"erp"）
	oracleDB := global.GetGlobalDBByDBName("erp")
	if oracleDB == nil {
		global.GVA_LOG.Error("未找到Oracle数据库连接!", zap.String("alias", "erp"))
		response.FailWithMessage("未找到Oracle数据库连接，请检查配置", c)
		return
	}

	// SQL查询语句
	sql := `select ima01, ima02, ima021 from forewin.ima_file where ima01 like 'M.IN%'`

	result, err := utils.QueryArray(ctx, oracleDB, sql)
	if err != nil {
		global.GVA_LOG.Error("Oracle查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithDetailed(result, "查询成功", c)
}
