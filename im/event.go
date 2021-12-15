package im

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/linkv-io/go-sdk/http"
	"net/url"
	"strings"
)

func (o *im) SendEventMsg(cmimToken, fromUID, toUID, content, objectName string) (string, error) {
	nonce := genGUID()
	timestamp := getTimestampS()

	sha1Data := sha1.Sum([]byte(o.GetConfig().AppID + "|" + o.GetConfig().AppKey + "|" + timestamp + "|" + nonce))
	sign := strings.ToUpper(hex.EncodeToString(sha1Data[:]))

	headers := make(map[string]string)
	headers["nonce"] = nonce
	headers["timestamp"] = timestamp
	headers["cmimToken"] = cmimToken
	headers["sign"] = sign
	headers["appUid"] = fromUID
	headers["appkey"] = o.GetConfig().AppKey
	headers["appId"] = o.GetConfig().AppID

	params := url.Values{}
	params.Set("fromUserId", fromUID)
	params.Set("toUserId", toUID)
	params.Set("content", content)
	params.Set("appId", o.GetConfig().AppID)
	params.Set("objectName", objectName)

	uri := o.GetConfig().URL + "/api/rest/v1/sendEventMsg"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return "", err
		}

		if resp.StatusCode != 200 {
			return "", fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			ID   string `json:"requestId"`
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return "", err
		}

		if result.Code == 200 {
			return result.ID, nil
		}

		return "", fmt.Errorf("code not 200(%v) message(%v)", result.Code, result.Msg)
	}
	return "", errResult
}
