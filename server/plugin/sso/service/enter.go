package service

import (
	"sync"
)

type ServiceGroup struct {
	SSOLoginService
}

var (
	ServiceGroupApp = new(ServiceGroup)
	once             sync.Once
)
