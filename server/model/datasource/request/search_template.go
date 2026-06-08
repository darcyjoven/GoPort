package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SearchTemplateSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           *string     `json:"name" form:"name"`
	DatabaseID     *int        `json:"databaseID" form:"databaseID"`
	SearchText     *string     `json:"searchText" form:"searchText"`
	Remark         *string     `json:"remark" form:"remark"`
	Argv1          *string     `json:"argv1" form:"argv1"`
	Argv2          *string     `json:"argv2" form:"argv2"`
	Argv3          *string     `json:"argv3" form:"argv3"`
	Argv4          *string     `json:"argv4" form:"argv4"`
	Argv5          *string     `json:"argv5" form:"argv5"`
	Argv6          *string     `json:"argv6" form:"argv6"`
	Argv7          *string     `json:"argv7" form:"argv7"`
	Argv8          *string     `json:"argv8" form:"argv8"`
	Argv9          *string     `json:"argv9" form:"argv9"`
	Argv10         *string     `json:"argv10" form:"argv10"`
	request.PageInfo
}
