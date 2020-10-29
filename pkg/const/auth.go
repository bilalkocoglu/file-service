package _const

import "time"

const (
	AuthorizationHeader = "Authorization"
	CurrentUser         = "CurrentUser"
	TokenExpireTime     = time.Minute * 180 * 10000 //3 hours
	SecretKey           = "ahyalandunya"
)
