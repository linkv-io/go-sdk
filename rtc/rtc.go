package rtc

import (
	. "github.com/linkv-io/go-sdk/config"
)

type RTC struct {
}

func (rtc *RTC) GetConfig() *RTCConfig {
	return &Conf.RTC
}
