
// 自动生成模板PdfmeHistory
package pdfme
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 打印设计历史资料 结构体  PdfmeHistory
type PdfmeHistory struct {
    global.GVA_MODEL
  Version  *string `json:"version" form:"version" gorm:"comment:版本;column:version;size:20;" binding:"required"`  //版本
  Config  datatypes.JSON `json:"config" form:"config" gorm:"comment:样式配置;column:config;" swaggertype:"object"`  //样式配置
  DefaultData  datatypes.JSON `json:"defaultData" form:"defaultData" gorm:"comment:默认数据;column:default_data;" swaggertype:"object"`  //默认数据
  SourceID  *int64 `json:"sourceID" form:"sourceID" gorm:"comment:数据源ID;column:source_i_d;"`  //数据源ID
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:1000;"`  //备注
  Active  *bool `json:"active" form:"active" gorm:"comment:生效否;column:active;"`  //生效否
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 打印设计历史资料 PdfmeHistory自定义表名 pdfme_history
func (PdfmeHistory) TableName() string {
    return "pdfme_history"
}





