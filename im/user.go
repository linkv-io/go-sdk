package im

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/linkv-io/go-sdk/http"
)

func (o *im) UserBlock(userIDs []string, min int64) error {
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
	params.Set("appId", o.GetConfig().AppID)
	params.Set("userIds", strings.Join(userIDs, ","))
	if min > 0 {
		params.Set("minute", strconv.FormatInt(min, 10))
	}
	uri := o.GetConfig().URL + "/api/rest/user/block"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return err
		}

		if result.Code == 200 {
			return nil
		}

		return fmt.Errorf("code not 200(%v) message(%v)", result.Code, result.Msg)
	}
	return errResult
}

func (o *im) UserStatus(userID string) (bool, string, error) {
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
	params.Set("appId", o.GetConfig().AppID)
	params.Set("userId", userID)

	uri := o.GetConfig().URL + "/api/rest/userOnlineStatus"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return false, "", err
		}

		if resp.StatusCode != 200 {
			return false, "", fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			Code      int    `json:"code"`
			Online    int    `json:"data"`
			Msg       string `json:"msg"`
			RequestID string `json:"requestID"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return false, "", err
		}

		if result.Code == 200 {
			return result.Online == 1, result.RequestID, nil
		}

		return false, "", fmt.Errorf("code not 200(%v) message(%v)", result.Code, result.Msg)
	}
	return false, "", errResult
}

func (o *im) AddUserBlack(fromUID string, userIDs []string) (requestID string, failUIDs []string, err error) {
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
	params.Set("userId", o.GetConfig().AppID)
	params.Set("blackUserIds", strings.Join(userIDs, ","))
	uri := o.GetConfig().URL + "/api/rest/user/addUserBlack"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return "", userIDs, err
		}

		if resp.StatusCode != 200 {
			return "", userIDs, fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			ID   string         `json:"requestID"`
			Code int            `json:"code"`
			Msg  string         `json:"msg"`
			Data map[string]int `json:"messageSendResult"`
		}

		if err = json.Unmarshal(jsonData, &result); err != nil {
			return "", userIDs, err
		}

		if result.Code != 200 {
			return "", userIDs, fmt.Errorf("code not 200(%v) message(%v)", result.Code, result.Msg)
		}

		var list []string
		for k, v := range result.Data {
			if v != 200 {
				list = append(list, k)
			}
		}
		return result.ID, list, nil

	}
	return "", userIDs, errResult
}

func (o *im) RemoveUserBlack(fromUID string, userIDs []string) (requestID string, failUIDs []string, err error) {
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
	params.Set("userId", o.GetConfig().AppID)
	params.Set("blackUserIds", strings.Join(userIDs, ","))
	uri := o.GetConfig().URL + "/api/rest/user/removeUserBlack"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return "", userIDs, err
		}

		if resp.StatusCode != 200 {
			return "", userIDs, fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			ID   string         `json:"requestID"`
			Code int            `json:"code"`
			Msg  string         `json:"msg"`
			Data map[string]int `json:"messageSendResult"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return "", userIDs, err
		}
		if result.Code != 200 {
			return "", userIDs, fmt.Errorf("code not 200(%v) message(%v)", result.Code, result.Msg)
		}

		var list []string
		for k, v := range result.Data {
			if v != 200 {
				list = append(list, k)
			}
		}
		return result.ID, list, nil
	}
	return "", userIDs, errResult
}
