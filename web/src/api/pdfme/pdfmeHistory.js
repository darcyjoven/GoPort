import service from '@/utils/request'
// @Tags PdfmeHistory
// @Summary 创建打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PdfmeHistory true "创建打印设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PMHistory/createPdfmeHistory [post]
export const createPdfmeHistory = (data) => {
  return service({
    url: '/PMHistory/createPdfmeHistory',
    method: 'post',
    data
  })
}

// @Tags PdfmeHistory
// @Summary 删除打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PdfmeHistory true "删除打印设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PMHistory/deletePdfmeHistory [delete]
export const deletePdfmeHistory = (params) => {
  return service({
    url: '/PMHistory/deletePdfmeHistory',
    method: 'delete',
    params
  })
}

// @Tags PdfmeHistory
// @Summary 批量删除打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除打印设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PMHistory/deletePdfmeHistory [delete]
export const deletePdfmeHistoryByIds = (params) => {
  return service({
    url: '/PMHistory/deletePdfmeHistoryByIds',
    method: 'delete',
    params
  })
}

// @Tags PdfmeHistory
// @Summary 更新打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PdfmeHistory true "更新打印设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PMHistory/updatePdfmeHistory [put]
export const updatePdfmeHistory = (data) => {
  return service({
    url: '/PMHistory/updatePdfmeHistory',
    method: 'put',
    data
  })
}

// @Tags PdfmeHistory
// @Summary 用id查询打印设计历史资料
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PdfmeHistory true "用id查询打印设计历史资料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PMHistory/findPdfmeHistory [get]
export const findPdfmeHistory = (params) => {
  return service({
    url: '/PMHistory/findPdfmeHistory',
    method: 'get',
    params
  })
}

// @Tags PdfmeHistory
// @Summary 分页获取打印设计历史资料列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取打印设计历史资料列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PMHistory/getPdfmeHistoryList [get]
export const getPdfmeHistoryList = (params) => {
  return service({
    url: '/PMHistory/getPdfmeHistoryList',
    method: 'get',
    params
  })
}
// @Tags PdfmeHistory
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PMHistory/findPdfmeHistoryDataSource [get]
export const getPdfmeHistoryDataSource = () => {
  return service({
    url: '/PMHistory/getPdfmeHistoryDataSource',
    method: 'get',
  })
}

// @Tags PdfmeHistory
// @Summary 不需要鉴权的打印设计历史资料接口
// @Accept application/json
// @Produce application/json
// @Param data query pdfmeReq.PdfmeHistorySearch true "分页获取打印设计历史资料列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PMHistory/getPdfmeHistoryPublic [get]
export const getPdfmeHistoryPublic = () => {
  return service({
    url: '/PMHistory/getPdfmeHistoryPublic',
    method: 'get',
  })
}
