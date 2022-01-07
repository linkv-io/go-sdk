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

func (o *im) SendEventMsg(fromUID, toUID, content, objectName string) (string, error) {
	sha1Data := sha1.Sum([]byte(o.GetConfig().AppID + "|" + o.GetConfig().AppKey + "|" + o.timestamp + "|" + o.nonce))
	sign := strings.ToUpper(hex.EncodeToString(sha1Data[:]))

	headers := make(map[string]string)
	headers["nonce"] = o.nonce
	headers["timestamp"] = o.timestamp
	headers["cmimToken"] = o.cmimToken
	headers["appUid"] = o.operatorID

	headers["sign"] = sign
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
			ID   string `json:"requestID"`
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
