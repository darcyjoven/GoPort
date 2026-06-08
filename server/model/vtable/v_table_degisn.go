
// 自动生成模板VTableDesign
package vtable
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 表格设计 结构体  VTableDesign
type VTableDesign struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:名称;column:name;size:255;" binding:"required"`  //名称
  Desciption  *string `json:"desciption" form:"desciption" gorm:"comment:说明;column:desciption;size:1000;"`  //说明
  Module  *string `json:"module" form:"module" gorm:"comment:模块;column:module;size:20;" binding:"required"`  //模块
  CurrentVersion  *string `json:"currentVersion" form:"currentVersion" gorm:"comment:当前版本;column:current_version;size:20;" binding:"required"`  //当前版本
  Checkout  *bool `json:"checkout" form:"checkout" gorm:"comment:签出否;column:checkout;"`  //签出否
  CheckoutVersion  *string `json:"checkoutVersion" form:"checkoutVersion" gorm:"comment:签出版本;column:checkout_version;size:20;"`  //签出版本
  LastCheckOutTime  *time.Time `json:"lastCheckOutTime" form:"lastCheckOutTime" gorm:"comment:上次签出时间;column:last_check_out_time;"`  //上次签出时间
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:1000;"`  //备注
  Active  *bool `json:"active" form:"active" gorm:"comment:生效否;column:active;"`  //生效否
}


// TableName 表格设计 VTableDesign自定义表名 vtable_degisn
func (VTableDesign) TableName() string {
    return "vtable_degisn"
}





