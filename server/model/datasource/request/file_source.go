
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type FileSourceSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      Description  *string `json:"description" form:"description"` 
      Remark  *string `json:"remark" form:"remark"` 
      FileType  *string `json:"fileType" form:"fileType"` 
    request.PageInfo
}
