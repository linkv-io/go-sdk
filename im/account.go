package im

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/linkv-io/go-sdk/http"
	"net/url"
	"sort"
	"strings"
)

func (o *im) GetTokenByThirdUID(thirdUID string) (string, error) {
	nonce := genGUID()
	timestamp := getTimestampS()

	arr := []string{nonce, timestamp, o.GetConfig().AppSecret}
	sort.Strings(arr)
	md5Data := md5.Sum([]byte(strings.Join(arr, "")))
	cmimToken := strings.ToLower(hex.EncodeToString(md5Data[:]))

	sha1Data := sha1.Sum([]byte(o.GetConfig().AppID + "|" + o.GetConfig().AppKey + "|" + timestamp + "|" + nonce))
	sign := strings.ToUpper(hex.EncodeToString(sha1Data[:]))

	headers := make(map[string]string)
	headers["nonce"] = nonce
	headers["timestamp"] = timestamp
	headers["cmimToken"] = cmimToken
	headers["signature"] = cmimToken
	headers["sign"] = sign
	headers["appkey"] = o.GetConfig().AppKey
	headers["appUid"] = thirdUID
	headers["appId"] = o.GetConfig().AppID

	params := url.Values{}
	params.Set("userId", thirdUID)

	uri := o.GetConfig().URL + "/api/rest/getToken"

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
			Code  int    `json:"code,string"`
			Token string `json:"token"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return "", err
		}

		if result.Code == 200 {
			return result.Token, nil
		}

		return "", fmt.Errorf("code not 200(%v)", result.Code)
	}
	return "", errResult

}
