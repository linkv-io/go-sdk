package rtc

import (
	"encoding/hex"
	"strconv"
	"time"
)

func (o *rtc) GenAuth() (string, string, string) {
	expireTS := strconv.FormatInt(time.Now().Unix(), 10)
	data := o.GetConfig().AppID + expireTS
	return o.GetConfig().AppID, hex.EncodeToString(hmacSha1([]byte(o.GetConfig().AppKey), []byte(data))), expireTS
}
