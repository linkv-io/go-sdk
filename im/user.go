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

func (o *im) UserBlock(cmimToken, fromUID string, userIDs []string, min int64) error {
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

func (o *im) UserStatus(cmimToken, fromUID, userID string) (bool, string, error) {
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
			RequestID string `json:"requestId"`
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
