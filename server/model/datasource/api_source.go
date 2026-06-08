
// 自动生成模板APISource
package datasource
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// api配置 结构体  APISource
type APISource struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"index;comment:模板名称;column:name;size:255;" binding:"required"`  //模板名称
  Path  *string `json:"path" form:"path" gorm:"comment:api路径;column:path;size:1000;" binding:"required"`  //api路径
  Remark  *string `json:"remark" form:"remark" gorm:"column:remark;size:1000;"`  //备注
  Argv1  *string `json:"argv1" form:"argv1" gorm:"comment:参数1默认值;column:argv1;size:255;"`  //参数1默认值
  Argv2  *string `json:"argv2" form:"argv2" gorm:"comment:参数2默认值;column:argv2;size:255;"`  //参数2默认值
  Argv3  *string `json:"argv3" form:"argv3" gorm:"comment:参数3默认值;column:argv3;size:255;"`  //参数3默认值
  Argv4  *string `json:"argv4" form:"argv4" gorm:"comment:参数4默认值;column:argv4;size:255;"`  //参数4默认值
  Argv5  *string `json:"argv5" form:"argv5" gorm:"comment:参数5默认值;column:argv5;size:255;"`  //参数5默认值
  Argv6  *string `json:"argv6" form:"argv6" gorm:"comment:参数6默认值;column:argv6;size:255;"`  //参数6默认值
  Argv7  *string `json:"argv7" form:"argv7" gorm:"comment:参数7默认值;column:argv7;size:255;"`  //参数7默认值
  Argv8  *string `json:"argv8" form:"argv8" gorm:"comment:参数8默认值;column:argv8;size:255;"`  //参数8默认值
  Argv9  *string `json:"argv9" form:"argv9" gorm:"comment:参数9默认值;column:argv9;size:255;"`  //参数9默认值
  Argv10  *string `json:"argv10" form:"argv10" gorm:"comment:参数10默认值;column:argv10;size:255;"`  //参数10默认值
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName api配置 APISource自定义表名 api_source
func (APISource) TableName() string {
    return "api_source"
}





