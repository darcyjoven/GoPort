import service from '@/utils/request'
// @Tags DataSourceField
// @Summary 创建数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DataSourceField true "创建数据源字段信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /DSField/createDataSourceField [post]
export const createDataSourceField = (data) => {
  return service({
    url: '/DSField/createDataSourceField',
    method: 'post',
    data
  })
}

// @Tags DataSourceField
// @Summary 删除数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DataSourceField true "删除数据源字段信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSField/deleteDataSourceField [delete]
export const deleteDataSourceField = (params) => {
  return service({
    url: '/DSField/deleteDataSourceField',
    method: 'delete',
    params
  })
}

// @Tags DataSourceField
// @Summary 批量删除数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据源字段信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSField/deleteDataSourceField [delete]
export const deleteDataSourceFieldByIds = (params) => {
  return service({
    url: '/DSField/deleteDataSourceFieldByIds',
    method: 'delete',
    params
  })
}

// @Tags DataSourceField
// @Summary 更新数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DataSourceField true "更新数据源字段信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSField/updateDataSourceField [put]
export const updateDataSourceField = (data) => {
  return service({
    url: '/DSField/updateDataSourceField',
    method: 'put',
    data
  })
}

// @Tags DataSourceField
// @Summary 用id查询数据源字段信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DataSourceField true "用id查询数据源字段信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSField/findDataSourceField [get]
export const findDataSourceField = (params) => {
  return service({
    url: '/DSField/findDataSourceField',
    method: 'get',
    params
  })
}

// @Tags DataSourceField
// @Summary 分页获取数据源字段信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据源字段信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSField/getDataSourceFieldList [get]
export const getDataSourceFieldList = (params) => {
  return service({
    url: '/DSField/getDataSourceFieldList',
    method: 'get',
    params
  })
}
// @Tags DataSourceField
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSField/findDataSourceFieldDataSource [get]
export const getDataSourceFieldDataSource = () => {
  return service({
    url: '/DSField/getDataSourceFieldDataSource',
    method: 'get',
  })
}

// @Tags DataSourceField
// @Summary 不需要鉴权的数据源字段信息接口
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.DataSourceFieldSearch true "分页获取数据源字段信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /DSField/getDataSourceFieldPublic [get]
export const getDataSourceFieldPublic = () => {
  return service({
    url: '/DSField/getDataSourceFieldPublic',
    method: 'get',
  })
}
