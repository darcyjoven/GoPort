
// 自动生成模板DataSourceField
package datasource
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 数据源字段信息 结构体  DataSourceField
type DataSourceField struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"index;comment:名称;column:name;size:1000;" binding:"required"`  //名称
  SourceType  *string `json:"sourceType" form:"sourceType" gorm:"comment:数据源类型;column:source_type;size:5;" binding:"required"`  //数据源类型
  SourceID  *int64 `json:"sourceID" form:"sourceID" gorm:"comment:数据源ID;column:source_i_d;" binding:"required"`  //数据源ID
  FieldIndex  *int16 `json:"fieldIndex" form:"fieldIndex" gorm:"comment:字段顺序;column:field_index;" binding:"required"`  //字段顺序
  FieldKey  *string `json:"fieldKey" form:"fieldKey" gorm:"comment:原始字段内容;column:field_key;size:255;"`  //原始字段内容
  FiledType  *string `json:"filedType" form:"filedType" gorm:"comment:字段类型;column:filed_type;size:20;" binding:"required"`  //字段类型
  FieldName  *string `json:"fieldName" form:"fieldName" gorm:"comment:字段别名;column:field_name;size:255;" binding:"required"`  //字段别名
  Description  *string `json:"description" form:"description" gorm:"comment:字段描述;column:description;size:1000;"`  //字段描述
  Sortable  *string `json:"sortable" form:"sortable" gorm:"comment:排序;column:sortable;size:5;" binding:"required"`  //排序
  Width  *float64 `json:"width" form:"width" gorm:"comment:宽度;column:width;size:20,6;"`  //宽度
  Format  *string `json:"format" form:"format" gorm:"comment:格式化;column:format;size:1000;"`  //格式化
  Warp  *bool `json:"warp" form:"warp" gorm:"default:false;comment:换行;column:warp;" binding:"required"`  //换行
  Align  *string `json:"align" form:"align" gorm:"default:5;comment:对齐方式;column:align;size:10;"`  //对齐方式
  Extra  *string `json:"extra" form:"extra" gorm:"comment:其它配置;column:extra;size:1000;"`  //其它配置
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 数据源字段信息 DataSourceField自定义表名 datasource_field
func (DataSourceField) TableName() string {
    return "datasource_field"
}





