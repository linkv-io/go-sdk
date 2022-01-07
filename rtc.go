package linkv_sdk

import (
	"github.com/linkv-io/go-sdk/rtc"
)

type LvRTC interface {
	GenAuth() (appID string, authKey string, expireTS string)
}

func NewLvRTC() LvRTC {
	return rtc.New()
}
