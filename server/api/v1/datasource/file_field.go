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

type FileFieldApi struct {}



// CreateFileField 创建文件字段定义
// @Tags FileField
// @Summary 创建文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.FileField true "创建文件字段定义"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /FField/createFileField [post]
func (FFieldApi *FileFieldApi) CreateFileField(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var FField datasource.FileField
	err := c.ShouldBindJSON(&FField)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    FField.CreatedBy = utils.GetUserID(c)
	err = FFieldService.CreateFileField(ctx,&FField)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteFileField 删除文件字段定义
// @Tags FileField
// @Summary 删除文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.FileField true "删除文件字段定义"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /FField/deleteFileField [delete]
func (FFieldApi *FileFieldApi) DeleteFileField(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := FFieldService.DeleteFileField(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteFileFieldByIds 批量删除文件字段定义
// @Tags FileField
// @Summary 批量删除文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /FField/deleteFileFieldByIds [delete]
func (FFieldApi *FileFieldApi) DeleteFileFieldByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := FFieldService.DeleteFileFieldByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateFileField 更新文件字段定义
// @Tags FileField
// @Summary 更新文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.FileField true "更新文件字段定义"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /FField/updateFileField [put]
func (FFieldApi *FileFieldApi) UpdateFileField(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var FField datasource.FileField
	err := c.ShouldBindJSON(&FField)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    FField.UpdatedBy = utils.GetUserID(c)
	err = FFieldService.UpdateFileField(ctx,FField)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindFileField 用id查询文件字段定义
// @Tags FileField
// @Summary 用id查询文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询文件字段定义"
// @Success 200 {object} response.Response{data=datasource.FileField,msg=string} "查询成功"
// @Router /FField/findFileField [get]
func (FFieldApi *FileFieldApi) FindFileField(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reFField, err := FFieldService.GetFileField(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reFField, c)
}
// GetFileFieldList 分页获取文件字段定义列表
// @Tags FileField
// @Summary 分页获取文件字段定义列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.FileFieldSearch true "分页获取文件字段定义列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /FField/getFileFieldList [get]
func (FFieldApi *FileFieldApi) GetFileFieldList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo datasourceReq.FileFieldSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := FFieldService.GetFileFieldInfoList(ctx,pageInfo)
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
// GetFileFieldDataSource 获取FileField的数据源
// @Tags FileField
// @Summary 获取FileField的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /FField/getFileFieldDataSource [get]
func (FFieldApi *FileFieldApi) GetFileFieldDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
    dataSource, err := FFieldService.GetFileFieldDataSource(ctx)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
   		response.FailWithMessage("查询失败:" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetFileFieldPublic 不需要鉴权的文件字段定义接口
// @Tags FileField
// @Summary 不需要鉴权的文件字段定义接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /FField/getFileFieldPublic [get]
func (FFieldApi *FileFieldApi) GetFileFieldPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    FFieldService.GetFileFieldPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的文件字段定义接口信息",
    }, "获取成功", c)
}
