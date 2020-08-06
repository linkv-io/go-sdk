package rtc

import (
	. "github.com/linkv-io/go-sdk/config"
)

func New() *rtc {
	return &rtc{}
}

type rtc struct {
}

func (o *rtc) GetConfig() *RTCConfig {
	return &Conf.RTC
}
