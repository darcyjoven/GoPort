
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type FileFieldSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      SourceID  *int `json:"sourceID" form:"sourceID"` 
      Index  *int `json:"index" form:"index"` 
      Key  *string `json:"key" form:"key"` 
      Name  *string `json:"name" form:"name"` 
      Description  *string `json:"description" form:"description"` 
    request.PageInfo
}
