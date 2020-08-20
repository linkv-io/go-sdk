package rtc

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	. "github.com/linkv-io/go-sdk/config"
	"strconv"
	"time"
)

func New() *rtc {
	return &rtc{}
}

type rtc struct {
}

func (o *rtc) GetConfig() *RTCConfig {
	return &Conf.RTC
}

func (o *rtc) GenAuth() (string, int64) {
	now := time.Now().Unix()
	expireTS := strconv.FormatInt(now, 10)
	return hex.EncodeToString(hmacSha1([]byte(o.GetConfig().AppKey), []byte(o.GetConfig().AppID+expireTS))), now
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
