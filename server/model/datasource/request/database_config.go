
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DatabaseConfigSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      DbType  *string `json:"dbType" form:"dbType"` 
      Host  *string `json:"host" form:"host"` 
      Port  *int `json:"port" form:"port"` 
      Server  *string `json:"server" form:"server"` 
      Username  *string `json:"username" form:"username"` 
      Password  *string `json:"password" form:"password"` 
      Remark  *string `json:"remark" form:"remark"` 
      Enable  *bool `json:"enable" form:"enable"` 
      LastTestTime  *time.Time `json:"lastTestTime" form:"lastTestTime"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
