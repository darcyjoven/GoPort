// 自动生成模板SearchTemplate
package datasource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 查询SQL模板 结构体  SearchTemplate
type SearchTemplate struct {
	global.GVA_MODEL
	Name       *string `json:"name" form:"name" gorm:"index;comment:模板名称;column:name;size:255;" binding:"required"`               //模板名称
	DatabaseID *int32  `json:"databaseID" form:"databaseID" gorm:"comment:数据库配置;column:database_id;" binding:"required"`          //数据库配置
	SearchText *string `json:"searchText" form:"searchText" gorm:"comment:查询内容;column:search_text;size:1000;" binding:"required"` //查询内容
	Remark     *string `json:"remark" form:"remark" gorm:"column:remark;size:1000;"`                                              //备注
	Argv1      *string `json:"argv1" form:"argv1" gorm:"comment:参数1默认值;column:argv1;size:255;"`                                   //参数1默认值
	Argv2      *string `json:"argv2" form:"argv2" gorm:"comment:参数2默认值;column:argv2;size:255;"`                                   //参数2默认值
	Argv3      *string `json:"argv3" form:"argv3" gorm:"comment:参数3默认值;column:argv3;size:255;"`                                   //参数3默认值
	Argv4      *string `json:"argv4" form:"argv4" gorm:"comment:参数4默认值;column:argv4;size:255;"`                                   //参数4默认值
	Argv5      *string `json:"argv5" form:"argv5" gorm:"comment:参数5默认值;column:argv5;size:255;"`                                   //参数5默认值
	Argv6      *string `json:"argv6" form:"argv6" gorm:"comment:参数6;column:argv6;size:255;"`                                      //参数6
	Argv7      *string `json:"argv7" form:"argv7" gorm:"comment:参数7;column:argv7;size:255;"`                                      //参数7
	Argv8      *string `json:"argv8" form:"argv8" gorm:"comment:参数8;column:argv8;size:255;"`                                      //参数8
	Argv9      *string `json:"argv9" form:"argv9" gorm:"comment:参数9;column:argv9;size:255;"`                                      //参数9
	Argv10     *string `json:"argv10" form:"argv10" gorm:"comment:参数10;column:argv10;size:255;"`                                  //参数10
}

// TableName 查询SQL模板 SearchTemplate自定义表名 search_template
func (SearchTemplate) TableName() string {
	return "search_template"
}
