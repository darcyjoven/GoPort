import service from '@/utils/request'
// @Tags PdfmeDesign
// @Summary 创建打印设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PdfmeDesign true "创建打印设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PMDesign/createPdfmeDesign [post]
export const createPdfmeDesign = (data) => {
  return service({
    url: '/PMDesign/createPdfmeDesign',
    method: 'post',
    data
  })
}

// @Tags PdfmeDesign
// @Summary 删除打印设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PdfmeDesign true "删除打印设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PMDesign/deletePdfmeDesign [delete]
export const deletePdfmeDesign = (params) => {
  return service({
    url: '/PMDesign/deletePdfmeDesign',
    method: 'delete',
    params
  })
}

// @Tags PdfmeDesign
// @Summary 批量删除打印设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除打印设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PMDesign/deletePdfmeDesign [delete]
export const deletePdfmeDesignByIds = (params) => {
  return service({
    url: '/PMDesign/deletePdfmeDesignByIds',
    method: 'delete',
    params
  })
}

// @Tags PdfmeDesign
// @Summary 更新打印设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PdfmeDesign true "更新打印设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PMDesign/updatePdfmeDesign [put]
export const updatePdfmeDesign = (data) => {
  return service({
    url: '/PMDesign/updatePdfmeDesign',
    method: 'put',
    data
  })
}

// @Tags PdfmeDesign
// @Summary 用id查询打印设计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PdfmeDesign true "用id查询打印设计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PMDesign/findPdfmeDesign [get]
export const findPdfmeDesign = (params) => {
  return service({
    url: '/PMDesign/findPdfmeDesign',
    method: 'get',
    params
  })
}

// @Tags PdfmeDesign
// @Summary 分页获取打印设计列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取打印设计列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PMDesign/getPdfmeDesignList [get]
export const getPdfmeDesignList = (params) => {
  return service({
    url: '/PMDesign/getPdfmeDesignList',
    method: 'get',
    params
  })
}

// @Tags PdfmeDesign
// @Summary 不需要鉴权的打印设计接口
// @Accept application/json
// @Produce application/json
// @Param data query pdfmeReq.PdfmeDesignSearch true "分页获取打印设计列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PMDesign/getPdfmeDesignPublic [get]
export const getPdfmeDesignPublic = () => {
  return service({
    url: '/PMDesign/getPdfmeDesignPublic',
    method: 'get',
  })
}
