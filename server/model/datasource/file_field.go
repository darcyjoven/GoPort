
// 自动生成模板FileField
package datasource
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 文件字段定义 结构体  FileField
type FileField struct {
    global.GVA_MODEL
  SourceID  *int64 `json:"sourceID" form:"sourceID" gorm:"comment:文件源ID;column:source_id;"`  //文件源ID
  Index  *int32 `json:"index" form:"index" gorm:"comment:序列;column:index;"`  //序列
  Key  *string `json:"key" form:"key" gorm:"comment:原始字段值;column:key;size:255;"`  //原始字段值
  Name  *string `json:"name" form:"name" gorm:"comment:字段名称;column:name;size:255;"`  //字段名称
  Description  *string `json:"description" form:"description" gorm:"comment:说明;column:description;size:1000;"`  //说明
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 文件字段定义 FileField自定义表名 file_field
func (FileField) TableName() string {
    return "file_field"
}





