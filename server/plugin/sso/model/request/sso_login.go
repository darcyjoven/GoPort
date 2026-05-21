package request

type SSOLoginReq struct {
	Provider string `form:"provider" json:"provider"`
	Username string `form:"username" json:"username"`
	Department string `form:"department" json:"department"`
	RealName string `form:"realName" json:"realName"`
	TargetPage string `form:"targetPage" json:"targetPage"`
	Timestamp int64 `form:"timestamp" json:"timestamp"`
	Sign string `form:"sign" json:"sign"`
	Token string `form:"token" json:"token"`
}
