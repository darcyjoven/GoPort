package pdfme

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/pdfme"
    pdfmeReq "github.com/flipped-aurora/gin-vue-admin/server/model/pdfme/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type PdfmeHistoryApi struct {}



// CreatePdfmeHistory 创建打印设计历史资料
// @Tags PdfmeHistory
// @Summary 创建打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body pdfme.PdfmeHistory true "创建打印设计历史资料"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /PMHistory/createPdfmeHistory [post]
func (PMHistoryApi *PdfmeHistoryApi) CreatePdfmeHistory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var PMHistory pdfme.PdfmeHistory
	err := c.ShouldBindJSON(&PMHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    PMHistory.CreatedBy = utils.GetUserID(c)
	err = PMHistoryService.CreatePdfmeHistory(ctx,&PMHistory)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePdfmeHistory 删除打印设计历史资料
// @Tags PdfmeHistory
// @Summary 删除打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body pdfme.PdfmeHistory true "删除打印设计历史资料"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /PMHistory/deletePdfmeHistory [delete]
func (PMHistoryApi *PdfmeHistoryApi) DeletePdfmeHistory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := PMHistoryService.DeletePdfmeHistory(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePdfmeHistoryByIds 批量删除打印设计历史资料
// @Tags PdfmeHistory
// @Summary 批量删除打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /PMHistory/deletePdfmeHistoryByIds [delete]
func (PMHistoryApi *PdfmeHistoryApi) DeletePdfmeHistoryByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := PMHistoryService.DeletePdfmeHistoryByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePdfmeHistory 更新打印设计历史资料
// @Tags PdfmeHistory
// @Summary 更新打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body pdfme.PdfmeHistory true "更新打印设计历史资料"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /PMHistory/updatePdfmeHistory [put]
func (PMHistoryApi *PdfmeHistoryApi) UpdatePdfmeHistory(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var PMHistory pdfme.PdfmeHistory
	err := c.ShouldBindJSON(&PMHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    PMHistory.UpdatedBy = utils.GetUserID(c)
	err = PMHistoryService.UpdatePdfmeHistory(ctx,PMHistory)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPdfmeHistory 用id查询打印设计历史资料
// @Tags PdfmeHistory
// @Summary 用id查询打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询打印设计历史资料"
// @Success 200 {object} response.Response{data=pdfme.PdfmeHistory,msg=string} "查询成功"
// @Router /PMHistory/findPdfmeHistory [get]
func (PMHistoryApi *PdfmeHistoryApi) FindPdfmeHistory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rePMHistory, err := PMHistoryService.GetPdfmeHistory(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rePMHistory, c)
}
// GetPdfmeHistoryList 分页获取打印设计历史资料列表
// @Tags PdfmeHistory
// @Summary 分页获取打印设计历史资料列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query pdfmeReq.PdfmeHistorySearch true "分页获取打印设计历史资料列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /PMHistory/getPdfmeHistoryList [get]
func (PMHistoryApi *PdfmeHistoryApi) GetPdfmeHistoryList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo pdfmeReq.PdfmeHistorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := PMHistoryService.GetPdfmeHistoryInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
// GetPdfmeHistoryDataSource 获取PdfmeHistory的数据源
// @Tags PdfmeHistory
// @Summary 获取PdfmeHistory的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /PMHistory/getPdfmeHistoryDataSource [get]
func (PMHistoryApi *PdfmeHistoryApi) GetPdfmeHistoryDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
    dataSource, err := PMHistoryService.GetPdfmeHistoryDataSource(ctx)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
   		response.FailWithMessage("查询失败:" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetPdfmeHistoryPublic 不需要鉴权的打印设计历史资料接口
// @Tags PdfmeHistory
// @Summary 不需要鉴权的打印设计历史资料接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PMHistory/getPdfmeHistoryPublic [get]
func (PMHistoryApi *PdfmeHistoryApi) GetPdfmeHistoryPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    PMHistoryService.GetPdfmeHistoryPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的打印设计历史资料接口信息",
    }, "获取成功", c)
}
