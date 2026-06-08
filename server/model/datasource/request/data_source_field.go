
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DataSourceFieldSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      SourceType  *string `json:"sourceType" form:"sourceType"` 
      FieldIndex  *int `json:"fieldIndex" form:"fieldIndex"` 
      FiledType  *string `json:"filedType" form:"filedType"` 
      FieldName  *string `json:"fieldName" form:"fieldName"` 
      Description  *string `json:"description" form:"description"` 
      Sortable  *string `json:"sortable" form:"sortable"` 
      Width  *float64 `json:"width" form:"width"` 
      Format  *string `json:"format" form:"format"` 
      Warp  *bool `json:"warp" form:"warp"` 
      Align  *string `json:"align" form:"align"` 
      Extra  *string `json:"extra" form:"extra"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
