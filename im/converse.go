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

func (o *im) PushConverseData(fromUID, toUID, content, objectName, toAppID, toUserExtSysUserID string) (string, error) {
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

	if len(toAppID) > 0 {
		params.Set("toUserAppid", toAppID)
	}

	if len(toUserExtSysUserID) > 0 {
		params.Set("toUserExtSysUserId", toUserExtSysUserID)
	}

	uri := o.GetConfig().URL + "/api/rest/message/v1/converse/pushConverseData"

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

func (o *im) PushConverseDatas(fromUID string, toUIDs []string, content, objectName string) (string, []string, error) {

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
	params.Set("toUserIds", strings.Join(toUIDs, ","))
	params.Set("content", content)
	params.Set("appId", o.GetConfig().AppID)
	params.Set("objectName", objectName)

	uri := o.GetConfig().URL + "/api/rest/message/v1/converse/pushConverseDatas"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return "", nil, err
		}

		if resp.StatusCode != 200 {
			return "", nil, fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			ID   string         `json:"requestID"`
			Code int            `json:"code"`
			Msg  string         `json:"msg"`
			Data map[string]int `json:"messageSendResult"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return "", nil, err
		}

		if result.Code != 200 {
			return "", nil, fmt.Errorf("code not 200(%v) message(%v)", result.Code, result.Msg)
		}
		var list []string
		for k, v := range result.Data {
			if v != 200 {
				list = append(list, k)
			}
		}

		return result.ID, list, nil

	}
	return "", nil, errResult
}
