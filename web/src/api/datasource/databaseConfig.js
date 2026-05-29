import service from '@/utils/request'
// @Tags DatabaseConfig
// @Summary 创建数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DatabaseConfig true "创建数据库连接配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dbConfig/createDatabaseConfig [post]
export const createDatabaseConfig = (data) => {
  return service({
    url: '/dbConfig/createDatabaseConfig',
    method: 'post',
    data
  })
}

// @Tags DatabaseConfig
// @Summary 删除数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DatabaseConfig true "删除数据库连接配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dbConfig/deleteDatabaseConfig [delete]
export const deleteDatabaseConfig = (params) => {
  return service({
    url: '/dbConfig/deleteDatabaseConfig',
    method: 'delete',
    params
  })
}

// @Tags DatabaseConfig
// @Summary 批量删除数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据库连接配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dbConfig/deleteDatabaseConfig [delete]
export const deleteDatabaseConfigByIds = (params) => {
  return service({
    url: '/dbConfig/deleteDatabaseConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags DatabaseConfig
// @Summary 更新数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DatabaseConfig true "更新数据库连接配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dbConfig/updateDatabaseConfig [put]
export const updateDatabaseConfig = (data) => {
  return service({
    url: '/dbConfig/updateDatabaseConfig',
    method: 'put',
    data
  })
}

// @Tags DatabaseConfig
// @Summary 用id查询数据库连接配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DatabaseConfig true "用id查询数据库连接配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dbConfig/findDatabaseConfig [get]
export const findDatabaseConfig = (params) => {
  return service({
    url: '/dbConfig/findDatabaseConfig',
    method: 'get',
    params
  })
}

// @Tags DatabaseConfig
// @Summary 分页获取数据库连接配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据库连接配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dbConfig/getDatabaseConfigList [get]
export const getDatabaseConfigList = (params) => {
  return service({
    url: '/dbConfig/getDatabaseConfigList',
    method: 'get',
    params
  })
}

// @Tags DatabaseConfig
// @Summary 不需要鉴权的数据库连接配置接口
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.DatabaseConfigSearch true "分页获取数据库连接配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dbConfig/getDatabaseConfigPublic [get]
export const getDatabaseConfigPublic = () => {
  return service({
    url: '/dbConfig/getDatabaseConfigPublic',
    method: 'get',
  })
}
