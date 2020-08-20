package rtc

import (
	"crypto/hmac"
	"crypto/sha1"
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

// hmac sha1
func hmacSha1(key []byte, message []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

// hmac sha1 verify
func hmacSha1Verify(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
