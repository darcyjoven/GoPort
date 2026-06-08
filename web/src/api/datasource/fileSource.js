import service from '@/utils/request'
// @Tags FileSource
// @Summary 创建文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FileSource true "创建文件源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /FSource/createFileSource [post]
export const createFileSource = (data) => {
  return service({
    url: '/FSource/createFileSource',
    method: 'post',
    data
  })
}

// @Tags FileSource
// @Summary 删除文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FileSource true "删除文件源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FSource/deleteFileSource [delete]
export const deleteFileSource = (params) => {
  return service({
    url: '/FSource/deleteFileSource',
    method: 'delete',
    params
  })
}

// @Tags FileSource
// @Summary 批量删除文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文件源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FSource/deleteFileSource [delete]
export const deleteFileSourceByIds = (params) => {
  return service({
    url: '/FSource/deleteFileSourceByIds',
    method: 'delete',
    params
  })
}

// @Tags FileSource
// @Summary 更新文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FileSource true "更新文件源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /FSource/updateFileSource [put]
export const updateFileSource = (data) => {
  return service({
    url: '/FSource/updateFileSource',
    method: 'put',
    data
  })
}

// @Tags FileSource
// @Summary 用id查询文件源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.FileSource true "用id查询文件源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FSource/findFileSource [get]
export const findFileSource = (params) => {
  return service({
    url: '/FSource/findFileSource',
    method: 'get',
    params
  })
}

// @Tags FileSource
// @Summary 分页获取文件源列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取文件源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FSource/getFileSourceList [get]
export const getFileSourceList = (params) => {
  return service({
    url: '/FSource/getFileSourceList',
    method: 'get',
    params
  })
}

// @Tags FileSource
// @Summary 不需要鉴权的文件源接口
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.FileSourceSearch true "分页获取文件源列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /FSource/getFileSourcePublic [get]
export const getFileSourcePublic = () => {
  return service({
    url: '/FSource/getFileSourcePublic',
    method: 'get',
  })
}
