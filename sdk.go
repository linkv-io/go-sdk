package linkv_sdk

import (
	"github.com/linkv-io/go-sdk/config"
	"github.com/linkv-io/go-sdk/im"
	"github.com/linkv-io/go-sdk/rtc"
)

var (
	IM  = &im.IM{}
	RTC = &rtc.RTC{}
)

func Init(appID, appSecret string) error {
	return config.Conf.Init(appID, appSecret)
}
