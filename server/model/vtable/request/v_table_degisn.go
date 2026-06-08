
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type VTableDesignSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      Desciption  *string `json:"desciption" form:"desciption"` 
      Module  *string `json:"module" form:"module"` 
      CurrentVersion  *string `json:"currentVersion" form:"currentVersion"` 
      Checkout  *bool `json:"checkout" form:"checkout"` 
      CheckoutVersion  *string `json:"checkoutVersion" form:"checkoutVersion"` 
      LastCheckOutTimeRange  []time.Time  `json:"lastCheckOutTimeRange" form:"lastCheckOutTimeRange[]"`
      Remark  *string `json:"remark" form:"remark"` 
      Active  *bool `json:"active" form:"active"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
