
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PdfmeHistorySearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Version  *string `json:"version" form:"version"` 
      SourceID  *int `json:"sourceID" form:"sourceID"` 
      Remark  *string `json:"remark" form:"remark"` 
      Active  *bool `json:"active" form:"active"` 
    request.PageInfo
}
