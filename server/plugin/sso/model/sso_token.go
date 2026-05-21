package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type SSOToken struct {
	global.GVA_MODEL
	ProviderName  string    `json:"providerName" gorm:"index;comment:SSO提供商名称"`
	Username      string    `json:"username" gorm:"index;comment:用户名"`
	Token         string    `json:"token" gorm:"uniqueIndex;comment:SSO验证token"`
	ExpiresAt     time.Time `json:"expiresAt" gorm:"index;comment:过期时间"`
	RemoteIP      string    `json:"remoteIP" gorm:"comment:请求IP"`
}

func (SSOToken) TableName() string {
	return "sso_tokens"
}
