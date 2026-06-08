import service from '@/utils/request'
// @Tags APISource
// @Summary 创建api配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.APISource true "创建api配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ASource/createAPISource [post]
export const createAPISource = (data) => {
  return service({
    url: '/ASource/createAPISource',
    method: 'post',
    data
  })
}

// @Tags APISource
// @Summary 删除api配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.APISource true "删除api配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ASource/deleteAPISource [delete]
export const deleteAPISource = (params) => {
  return service({
    url: '/ASource/deleteAPISource',
    method: 'delete',
    params
  })
}

// @Tags APISource
// @Summary 批量删除api配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除api配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ASource/deleteAPISource [delete]
export const deleteAPISourceByIds = (params) => {
  return service({
    url: '/ASource/deleteAPISourceByIds',
    method: 'delete',
    params
  })
}

// @Tags APISource
// @Summary 更新api配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.APISource true "更新api配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ASource/updateAPISource [put]
export const updateAPISource = (data) => {
  return service({
    url: '/ASource/updateAPISource',
    method: 'put',
    data
  })
}

// @Tags APISource
// @Summary 用id查询api配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.APISource true "用id查询api配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ASource/findAPISource [get]
export const findAPISource = (params) => {
  return service({
    url: '/ASource/findAPISource',
    method: 'get',
    params
  })
}

// @Tags APISource
// @Summary 分页获取api配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取api配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ASource/getAPISourceList [get]
export const getAPISourceList = (params) => {
  return service({
    url: '/ASource/getAPISourceList',
    method: 'get',
    params
  })
}

// @Tags APISource
// @Summary 不需要鉴权的api配置接口
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.APISourceSearch true "分页获取api配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ASource/getAPISourcePublic [get]
export const getAPISourcePublic = () => {
  return service({
    url: '/ASource/getAPISourcePublic',
    method: 'get',
  })
}
