package rtc

import (
	. "github.com/linkv-io/go-sdk/config"
)

type LvRTC interface {
}

func NewLvRTC() LvRTC {
	return &linkVRTC{}
}

type linkVRTC struct {
}

func (rtc *linkVRTC) GetConfig() *RTCConfig {
	return &Conf.RTC
}
