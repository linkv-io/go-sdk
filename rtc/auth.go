package rtc

import (
	"encoding/hex"
	"strconv"
	"time"
)

func (o *rtc) GenAuth() (string, string, int64) {
	now := time.Now().Unix()
	expireTS := strconv.FormatInt(now, 10)
	return o.GetConfig().AppID, hex.EncodeToString(hmacSha1([]byte(o.GetConfig().AppKey), []byte(o.GetConfig().AppID+expireTS))), now
}
