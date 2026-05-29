
// 自动生成模板DatabaseConfig
package datasource
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	"gorm.io/datatypes"
)

// 数据库连接配置 结构体  DatabaseConfig
type DatabaseConfig struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:数据库名称;column:name;size:255;" binding:"required"`  //数据库名称
  DbType  *string `json:"dbType" form:"dbType" gorm:"comment:数据库类型;column:db_type;" binding:"required"`  //数据库类型
  Host  *string `json:"host" form:"host" gorm:"comment:主机地址/IP;column:host;size:255;" binding:"required"`  //主机地址/IP
  Port  *int32 `json:"port" form:"port" gorm:"comment:端口号;column:port;"`  //端口号
  Server  *string `json:"server" form:"server" gorm:"comment:数据库名;column:server;size:255;"`  //数据库名
  Username  *string `json:"username" form:"username" gorm:"comment:用户名;column:username;size:255;"`  //用户名
  Password  *string `json:"password" form:"password" gorm:"comment:密码;column:password;size:255;"`  //密码
  ExtraParams  datatypes.JSON `json:"extraParams" form:"extraParams" gorm:"comment:其它参数;column:extra_params;size:1000;" swaggertype:"object"`  //其它参数
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:1000;"`  //备注
  Enable  *bool `json:"enable" form:"enable" gorm:"comment:是否启用;column:enable;" binding:"required"`  //是否启用
  LastTestTime  *time.Time `json:"lastTestTime" form:"lastTestTime" gorm:"comment:上次测试时间;column:last_test_time;"`  //上次测试时间
}


// TableName 数据库连接配置 DatabaseConfig自定义表名 database_config
func (DatabaseConfig) TableName() string {
    return "database_config"
}





