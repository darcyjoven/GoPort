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

type DataSourceFieldApi struct {}



// CreateDataSourceField 创建数据源字段信息
// @Tags DataSourceField
// @Summary 创建数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.DataSourceField true "创建数据源字段信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /DSField/createDataSourceField [post]
func (DSFieldApi *DataSourceFieldApi) CreateDataSourceField(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var DSField datasource.DataSourceField
	err := c.ShouldBindJSON(&DSField)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    DSField.CreatedBy = utils.GetUserID(c)
	err = DSFieldService.CreateDataSourceField(ctx,&DSField)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteDataSourceField 删除数据源字段信息
// @Tags DataSourceField
// @Summary 删除数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.DataSourceField true "删除数据源字段信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /DSField/deleteDataSourceField [delete]
func (DSFieldApi *DataSourceFieldApi) DeleteDataSourceField(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := DSFieldService.DeleteDataSourceField(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDataSourceFieldByIds 批量删除数据源字段信息
// @Tags DataSourceField
// @Summary 批量删除数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /DSField/deleteDataSourceFieldByIds [delete]
func (DSFieldApi *DataSourceFieldApi) DeleteDataSourceFieldByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := DSFieldService.DeleteDataSourceFieldByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDataSourceField 更新数据源字段信息
// @Tags DataSourceField
// @Summary 更新数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.DataSourceField true "更新数据源字段信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /DSField/updateDataSourceField [put]
func (DSFieldApi *DataSourceFieldApi) UpdateDataSourceField(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var DSField datasource.DataSourceField
	err := c.ShouldBindJSON(&DSField)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    DSField.UpdatedBy = utils.GetUserID(c)
	err = DSFieldService.UpdateDataSourceField(ctx,DSField)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDataSourceField 用id查询数据源字段信息
// @Tags DataSourceField
// @Summary 用id查询数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询数据源字段信息"
// @Success 200 {object} response.Response{data=datasource.DataSourceField,msg=string} "查询成功"
// @Router /DSField/findDataSourceField [get]
func (DSFieldApi *DataSourceFieldApi) FindDataSourceField(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reDSField, err := DSFieldService.GetDataSourceField(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reDSField, c)
}
// GetDataSourceFieldList 分页获取数据源字段信息列表
// @Tags DataSourceField
// @Summary 分页获取数据源字段信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.DataSourceFieldSearch true "分页获取数据源字段信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /DSField/getDataSourceFieldList [get]
func (DSFieldApi *DataSourceFieldApi) GetDataSourceFieldList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo datasourceReq.DataSourceFieldSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := DSFieldService.GetDataSourceFieldInfoList(ctx,pageInfo)
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
// GetDataSourceFieldDataSource 获取DataSourceField的数据源
// @Tags DataSourceField
// @Summary 获取DataSourceField的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /DSField/getDataSourceFieldDataSource [get]
func (DSFieldApi *DataSourceFieldApi) GetDataSourceFieldDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
    dataSource, err := DSFieldService.GetDataSourceFieldDataSource(ctx)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
   		response.FailWithMessage("查询失败:" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetDataSourceFieldPublic 不需要鉴权的数据源字段信息接口
// @Tags DataSourceField
// @Summary 不需要鉴权的数据源字段信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /DSField/getDataSourceFieldPublic [get]
func (DSFieldApi *DataSourceFieldApi) GetDataSourceFieldPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    DSFieldService.GetDataSourceFieldPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的数据源字段信息接口信息",
    }, "获取成功", c)
}
