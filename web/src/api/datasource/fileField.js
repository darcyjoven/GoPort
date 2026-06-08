import service from '@/utils/request'
// @Tags FileField
// @Summary 创建文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FileField true "创建文件字段定义"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /FField/createFileField [post]
export const createFileField = (data) => {
  return service({
    url: '/FField/createFileField',
    method: 'post',
    data
  })
}

// @Tags FileField
// @Summary 删除文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FileField true "删除文件字段定义"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FField/deleteFileField [delete]
export const deleteFileField = (params) => {
  return service({
    url: '/FField/deleteFileField',
    method: 'delete',
    params
  })
}

// @Tags FileField
// @Summary 批量删除文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文件字段定义"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FField/deleteFileField [delete]
export const deleteFileFieldByIds = (params) => {
  return service({
    url: '/FField/deleteFileFieldByIds',
    method: 'delete',
    params
  })
}

// @Tags FileField
// @Summary 更新文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FileField true "更新文件字段定义"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /FField/updateFileField [put]
export const updateFileField = (data) => {
  return service({
    url: '/FField/updateFileField',
    method: 'put',
    data
  })
}

// @Tags FileField
// @Summary 用id查询文件字段定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.FileField true "用id查询文件字段定义"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FField/findFileField [get]
export const findFileField = (params) => {
  return service({
    url: '/FField/findFileField',
    method: 'get',
    params
  })
}

// @Tags FileField
// @Summary 分页获取文件字段定义列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取文件字段定义列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FField/getFileFieldList [get]
export const getFileFieldList = (params) => {
  return service({
    url: '/FField/getFileFieldList',
    method: 'get',
    params
  })
}
// @Tags FileField
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FField/findFileFieldDataSource [get]
export const getFileFieldDataSource = () => {
  return service({
    url: '/FField/getFileFieldDataSource',
    method: 'get',
  })
}

// @Tags FileField
// @Summary 不需要鉴权的文件字段定义接口
// @Accept application/json
// @Produce application/json
// @Param data query datasourceReq.FileFieldSearch true "分页获取文件字段定义列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /FField/getFileFieldPublic [get]
export const getFileFieldPublic = () => {
  return service({
    url: '/FField/getFileFieldPublic',
    method: 'get',
  })
}
