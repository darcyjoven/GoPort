import service from '@/utils/request'
// @Tags SearchTemplate
// @Summary 创建查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SearchTemplate true "创建查询SQL模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /searchTemp/createSearchTemplate [post]
export const createSearchTemplate = (data) => {
  return service({
    url: '/searchTemp/createSearchTemplate',
    method: 'post',
    data
  })
}

// @Tags SearchTemplate
// @Summary 删除查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SearchTemplate true "删除查询SQL模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /searchTemp/deleteSearchTemplate [delete]
export const deleteSearchTemplate = (params) => {
  return service({
    url: '/searchTemp/deleteSearchTemplate',
    method: 'delete',
    params
  })
}

// @Tags SearchTemplate
// @Summary 批量删除查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除查询SQL模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /searchTemp/deleteSearchTemplate [delete]
export const deleteSearchTemplateByIds = (params) => {
  return service({
    url: '/searchTemp/deleteSearchTemplateByIds',
    method: 'delete',
    params
  })
}

// @Tags SearchTemplate
// @Summary 更新查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SearchTemplate true "更新查询SQL模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /searchTemp/updateSearchTemplate [put]
export const updateSearchTemplate = (data) => {
  return service({
    url: '/searchTemp/updateSearchTemplate',
    method: 'put',
    data
  })
}

// @Tags SearchTemplate
// @Summary 用id查询查询SQL模板
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SearchTemplate true "用id查询查询SQL模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /searchTemp/findSearchTemplate [get]
export const findSearchTemplate = (params) => {
  return service({
    url: '/searchTemp/findSearchTemplate',
    method: 'get',
    params
  })
}

// @Tags SearchTemplate
// @Summary 分页获取查询SQL模板列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取查询SQL模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /searchTemp/getSearchTemplateList [get]
export const getSearchTemplateList = (params) => {
  return service({
    url: '/searchTemp/getSearchTemplateList',
    method: 'get',
    params
  })
}
// @Tags SearchTemplate
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /searchTemp/findSearchTemplateDataSource [get]
export const getSearchTemplateDataSource = () => {
  return service({
    url: '/searchTemp/getSearchTemplateDataSource',
    method: 'get',
  })
}

// @Tags SearchTemplate
// @Summary 不需要鉴权的查询SQL模板接口
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.SearchTemplateSearch true "分页获取查询SQL模板列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /searchTemp/getSearchTemplatePublic [get]
export const getSearchTemplatePublic = () => {
  return service({
    url: '/searchTemp/getSearchTemplatePublic',
    method: 'get',
  })
}
