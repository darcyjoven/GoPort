package datasource

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/datasource"
    datasourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/datasource/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type DatabaseConfigApi struct {}



// CreateDatabaseConfig 创建数据库连接配置
// @Tags DatabaseConfig
// @Summary 创建数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.DatabaseConfig true "创建数据库连接配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dbConfig/createDatabaseConfig [post]
func (dbConfigApi *DatabaseConfigApi) CreateDatabaseConfig(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var dbConfig datasource.DatabaseConfig
	err := c.ShouldBindJSON(&dbConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dbConfigService.CreateDatabaseConfig(ctx,&dbConfig)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteDatabaseConfig 删除数据库连接配置
// @Tags DatabaseConfig
// @Summary 删除数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.DatabaseConfig true "删除数据库连接配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dbConfig/deleteDatabaseConfig [delete]
func (dbConfigApi *DatabaseConfigApi) DeleteDatabaseConfig(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := dbConfigService.DeleteDatabaseConfig(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDatabaseConfigByIds 批量删除数据库连接配置
// @Tags DatabaseConfig
// @Summary 批量删除数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dbConfig/deleteDatabaseConfigByIds [delete]
func (dbConfigApi *DatabaseConfigApi) DeleteDatabaseConfigByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := dbConfigService.DeleteDatabaseConfigByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDatabaseConfig 更新数据库连接配置
// @Tags DatabaseConfig
// @Summary 更新数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datasource.DatabaseConfig true "更新数据库连接配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dbConfig/updateDatabaseConfig [put]
func (dbConfigApi *DatabaseConfigApi) UpdateDatabaseConfig(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var dbConfig datasource.DatabaseConfig
	err := c.ShouldBindJSON(&dbConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dbConfigService.UpdateDatabaseConfig(ctx,dbConfig)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDatabaseConfig 用id查询数据库连接配置
// @Tags DatabaseConfig
// @Summary 用id查询数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询数据库连接配置"
// @Success 200 {object} response.Response{data=datasource.DatabaseConfig,msg=string} "查询成功"
// @Router /dbConfig/findDatabaseConfig [get]
func (dbConfigApi *DatabaseConfigApi) FindDatabaseConfig(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	redbConfig, err := dbConfigService.GetDatabaseConfig(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(redbConfig, c)
}
// GetDatabaseConfigList 分页获取数据库连接配置列表
// @Tags DatabaseConfig
// @Summary 分页获取数据库连接配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.DatabaseConfigSearch true "分页获取数据库连接配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dbConfig/getDatabaseConfigList [get]
func (dbConfigApi *DatabaseConfigApi) GetDatabaseConfigList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo datasourceReq.DatabaseConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dbConfigService.GetDatabaseConfigInfoList(ctx,pageInfo)
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

// GetDatabaseConfigPublic 不需要鉴权的数据库连接配置接口
// @Tags DatabaseConfig
// @Summary 不需要鉴权的数据库连接配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dbConfig/getDatabaseConfigPublic [get]
func (dbConfigApi *DatabaseConfigApi) GetDatabaseConfigPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    dbConfigService.GetDatabaseConfigPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的数据库连接配置接口信息",
    }, "获取成功", c)
}
