import service from '@/utils/request'
// @Tags VTableDesign
// @Summary 创建表格设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VTableDesign true "创建表格设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /VTDesign/createVTableDesign [post]
export const createVTableDesign = (data) => {
  return service({
    url: '/VTDesign/createVTableDesign',
    method: 'post',
    data
  })
}

// @Tags VTableDesign
// @Summary 删除表格设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VTableDesign true "删除表格设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VTDesign/deleteVTableDesign [delete]
export const deleteVTableDesign = (params) => {
  return service({
    url: '/VTDesign/deleteVTableDesign',
    method: 'delete',
    params
  })
}

// @Tags VTableDesign
// @Summary 批量删除表格设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除表格设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VTDesign/deleteVTableDesign [delete]
export const deleteVTableDesignByIds = (params) => {
  return service({
    url: '/VTDesign/deleteVTableDesignByIds',
    method: 'delete',
    params
  })
}

// @Tags VTableDesign
// @Summary 更新表格设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VTableDesign true "更新表格设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /VTDesign/updateVTableDesign [put]
export const updateVTableDesign = (data) => {
  return service({
    url: '/VTDesign/updateVTableDesign',
    method: 'put',
    data
  })
}

// @Tags VTableDesign
// @Summary 用id查询表格设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.VTableDesign true "用id查询表格设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /VTDesign/findVTableDesign [get]
export const findVTableDesign = (params) => {
  return service({
    url: '/VTDesign/findVTableDesign',
    method: 'get',
    params
  })
}

// @Tags VTableDesign
// @Summary 分页获取表格设计列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取表格设计列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /VTDesign/getVTableDesignList [get]
export const getVTableDesignList = (params) => {
  return service({
    url: '/VTDesign/getVTableDesignList',
    method: 'get',
    params
  })
}

// @Tags VTableDesign
// @Summary 不需要鉴权的表格设计接口
// @Accept application/json
// @Produce application/json
// @Param data query vtableReq.VTableDesignSearch true "分页获取表格设计列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /VTDesign/getVTableDesignPublic [get]
export const getVTableDesignPublic = () => {
  return service({
    url: '/VTDesign/getVTableDesignPublic',
    method: 'get',
  })
}
