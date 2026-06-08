
// 自动生成模板FileSource
package datasource
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 文件源 结构体  FileSource
type FileSource struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:名称;column:name;size:255;"`  //名称
  Description  *string `json:"description" form:"description" gorm:"comment:描述;column:description;size:1000;"`  //描述
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:1000;"`  //备注
  FileType  *string `json:"fileType" form:"fileType" gorm:"comment:文件类型;column:file_type;size:文件类型;"`  //文件类型
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 文件源 FileSource自定义表名 file_source
func (FileSource) TableName() string {
    return "file_source"
}





