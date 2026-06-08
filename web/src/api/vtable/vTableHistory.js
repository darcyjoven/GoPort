import service from '@/utils/request'
// @Tags VTableHistory
// @Summary 创建表格设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VTableHistory true "创建表格设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /VTHistory/createVTableHistory [post]
export const createVTableHistory = (data) => {
  return service({
    url: '/VTHistory/createVTableHistory',
    method: 'post',
    data
  })
}

// @Tags VTableHistory
// @Summary 删除表格设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VTableHistory true "删除表格设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VTHistory/deleteVTableHistory [delete]
export const deleteVTableHistory = (params) => {
  return service({
    url: '/VTHistory/deleteVTableHistory',
    method: 'delete',
    params
  })
}

// @Tags VTableHistory
// @Summary 批量删除表格设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除表格设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VTHistory/deleteVTableHistory [delete]
export const deleteVTableHistoryByIds = (params) => {
  return service({
    url: '/VTHistory/deleteVTableHistoryByIds',
    method: 'delete',
    params
  })
}

// @Tags VTableHistory
// @Summary 更新表格设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VTableHistory true "更新表格设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /VTHistory/updateVTableHistory [put]
export const updateVTableHistory = (data) => {
  return service({
    url: '/VTHistory/updateVTableHistory',
    method: 'put',
    data
  })
}

// @Tags VTableHistory
// @Summary 用id查询表格设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.VTableHistory true "用id查询表格设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /VTHistory/findVTableHistory [get]
export const findVTableHistory = (params) => {
  return service({
    url: '/VTHistory/findVTableHistory',
    method: 'get',
    params
  })
}

// @Tags VTableHistory
// @Summary 分页获取表格设计历史资料列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取表格设计历史资料列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /VTHistory/getVTableHistoryList [get]
export const getVTableHistoryList = (params) => {
  return service({
    url: '/VTHistory/getVTableHistoryList',
    method: 'get',
    params
  })
}
// @Tags VTableHistory
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /VTHistory/findVTableHistoryDataSource [get]
export const getVTableHistoryDataSource = () => {
  return service({
    url: '/VTHistory/getVTableHistoryDataSource',
    method: 'get',
  })
}

// @Tags VTableHistory
// @Summary 不需要鉴权的表格设计历史资料接口
// @Accept application/json
// @Produce application/json
// @Param data query vtableReq.VTableHistorySearch true "分页获取表格设计历史资料列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /VTHistory/getVTableHistoryPublic [get]
export const getVTableHistoryPublic = () => {
  return service({
    url: '/VTHistory/getVTableHistoryPublic',
    method: 'get',
  })
}
