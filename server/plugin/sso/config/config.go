package config

type SSO struct {
	Providers []Provider `mapstructure:"providers" json:"providers" yaml:"providers"`
}

type Provider struct {
	Name             string `mapstructure:"name" json:"name" yaml:"name"`
	Enabled          bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	VerifyURL        string `mapstructure:"verifyURL" json:"verifyURL" yaml:"verifyURL"`
	VerifyText       string `mapstructure:"verifyText" json:"verifyText" yaml:"verifyText"`
	SuccessText      string `mapstructure:"successText" json:"successText" yaml:"successText"`
	SecretKey        string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
	TokenExpireTime  string `mapstructure:"tokenExpireTime" json:"tokenExpireTime" yaml:"tokenExpireTime"`
	AllowAutoRegister bool  `mapstructure:"allowAutoRegister" json:"allowAutoRegister" yaml:"allowAutoRegister"`
	DefaultAuthorityId uint `mapstructure:"defaultAuthorityId" json:"defaultAuthorityId" yaml:"defaultAuthorityId"`
}
