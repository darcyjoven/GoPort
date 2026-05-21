package api

import (
	"sync"
)

type ApiGroup struct {
	SSOLoginApi
}

var (
	ApiGroupApp = new(ApiGroup)
	once        sync.Once
)
