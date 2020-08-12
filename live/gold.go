package live

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/linkv-io/go-sdk/http"
)

func (o *live) ChangeGoldByLiveOpenID(liveOpenID string, orderType, gold, expr int64, optionalReason string) (bool, error) {
	params := url.Values{}
	nonce := genRandomString()
	params.Add("nonce_str", nonce)
	params.Add("app_id", o.GetConfig().AppKey)

	params.Add("uid", liveOpenID)
	params.Add("request_id", genUniqueIDString())
	params.Add("type", strconv.FormatInt(orderType, 10))
	params.Add("value", strconv.FormatInt(gold, 10))
	if expr > 0 {
		params.Add("expriation", strconv.FormatInt(time.Now().Unix()+int64(expr)*86400, 10))
	}

	if len(optionalReason) > 0 {
		params.Add("reason", optionalReason)
	}

	params.Add("sign", genSign(params, o.GetConfig().AppSecret))

	uri := o.GetConfig().URL + "/open/finanv0/changeGold"

	var errResult error
	for i := 0; i < 3; i++ {
		jsonData, resp, err := http.PostDataWithHeader(uri, params, nil)
		if err != nil {
			return false, err
		}
		if resp.StatusCode != 200 {
			return false, fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}
		var result struct {
			Status int    `json:"status,string"`
			Msg    string `json:"msg"`
		}
		if err := json.Unmarshal(jsonData, &result); err != nil {
			return false, err
		}

		if result.Status == 200 {
			return true, nil
		}

		if result.Status == 500 {
			errResult = fmt.Errorf("message(%v)", result.Msg)
			time.Sleep(waitTime)
			continue
		}
		return false, nil
	}
	return false, errResult
}

func (o *live) GetGoldByLiveOpenID(liveOpenID string) (int64, error) {
	params := url.Values{}
	nonce := genRandomString()
	params.Add("nonce_str", nonce)
	params.Add("app_id", o.GetConfig().AppKey)

	params.Add("uid", liveOpenID)

	params.Add("sign", genSign(params, o.GetConfig().AppSecret))

	uri := o.GetConfig().URL + "/open/finanv0/getUserTokens"

	var errResult error
	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.GetDataWithHeader(uri, params, nil)
		if err != nil {
			return 0, err
		}
		if resp.StatusCode != 200 {
			return 0, fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}
		var result struct {
			Status int    `json:"status,string"`
			Msg    string `json:"msg"`
		}
		if err := json.Unmarshal(jsonData, &result); err != nil {
			return 0, err
		}

		if result.Status == 200 {
			var resultData struct {
				Data struct {
					Token     int64  `json:"livemeTokens,string"`
					IsMigrate string `json:"isMigrate"`
				} `json:"data"`
			}
			if err := json.Unmarshal(jsonData, &resultData); err != nil {
				return 0, err
			}
			return resultData.Data.Token, nil
		}
		if result.Status == 500 {
			errResult = fmt.Errorf("message(%v)", result.Msg)
			time.Sleep(waitTime)
			continue
		}
		return 0, fmt.Errorf("message(%v)", result.Msg)
	}
	return 0, errResult
}
