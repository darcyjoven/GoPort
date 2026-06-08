package datasource

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
    datasourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/datasource/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type FileSourceApi struct {}



// CreateFileSource 创建文件源
// @Tags FileSource
// @Summary 创建文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.FileSource true "创建文件源"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /FSource/createFileSource [post]
func (FSourceApi *FileSourceApi) CreateFileSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var FSource datasource.FileSource
	err := c.ShouldBindJSON(&FSource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    FSource.CreatedBy = utils.GetUserID(c)
	err = FSourceService.CreateFileSource(ctx,&FSource)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteFileSource 删除文件源
// @Tags FileSource
// @Summary 删除文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.FileSource true "删除文件源"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /FSource/deleteFileSource [delete]
func (FSourceApi *FileSourceApi) DeleteFileSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := FSourceService.DeleteFileSource(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteFileSourceByIds 批量删除文件源
// @Tags FileSource
// @Summary 批量删除文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /FSource/deleteFileSourceByIds [delete]
func (FSourceApi *FileSourceApi) DeleteFileSourceByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := FSourceService.DeleteFileSourceByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateFileSource 更新文件源
// @Tags FileSource
// @Summary 更新文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.FileSource true "更新文件源"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /FSource/updateFileSource [put]
func (FSourceApi *FileSourceApi) UpdateFileSource(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var FSource datasource.FileSource
	err := c.ShouldBindJSON(&FSource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    FSource.UpdatedBy = utils.GetUserID(c)
	err = FSourceService.UpdateFileSource(ctx,FSource)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindFileSource 用id查询文件源
// @Tags FileSource
// @Summary 用id查询文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询文件源"
// @Success 200 {object} response.Response{data=datasource.FileSource,msg=string} "查询成功"
// @Router /FSource/findFileSource [get]
func (FSourceApi *FileSourceApi) FindFileSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reFSource, err := FSourceService.GetFileSource(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reFSource, c)
}
// GetFileSourceList 分页获取文件源列表
// @Tags FileSource
// @Summary 分页获取文件源列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.FileSourceSearch true "分页获取文件源列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /FSource/getFileSourceList [get]
func (FSourceApi *FileSourceApi) GetFileSourceList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo datasourceReq.FileSourceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := FSourceService.GetFileSourceInfoList(ctx,pageInfo)
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

// GetFileSourcePublic 不需要鉴权的文件源接口
// @Tags FileSource
// @Summary 不需要鉴权的文件源接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /FSource/getFileSourcePublic [get]
func (FSourceApi *FileSourceApi) GetFileSourcePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    FSourceService.GetFileSourcePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的文件源接口信息",
    }, "获取成功", c)
}
