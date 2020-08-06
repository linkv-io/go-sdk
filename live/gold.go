package live

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/linkv-io/go-sdk/http"
)

func (o *live) ChangeGoldByLiveOpenID(liveOpenID, uniqueID string, typ OrderType, gold, expr int64, optionalReason string) (bool, error) {
	params := url.Values{}
	nonce := genRandomString()
	params.Add("nonce_str", nonce)
	params.Add("app_id", o.GetConfig().AppKey)

	params.Add("uid", liveOpenID)
	params.Add("request_id", uniqueID)
	params.Add("type", typ.String())
	params.Add("value", strconv.FormatInt(gold, 10))
	if expr > 0 {
		params.Add("expriation", strconv.FormatInt(time.Now().Unix()+int64(expr)*86400, 10))
	}

	if len(optionalReason) > 0 {
		params.Add("reason", optionalReason)
	}

	params.Add("sign", genSign(params, o.GetConfig().AppSecret))

	uri := o.GetConfig().URL + "/open/finanv0/changeGold"
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
	return result.Status == 200, nil
}

func (o *live) GetGoldByLiveOpenID(liveOpenID string) (int64, error) {
	params := url.Values{}
	nonce := genRandomString()
	params.Add("nonce_str", nonce)
	params.Add("app_id", o.GetConfig().AppKey)

	params.Add("uid", liveOpenID)

	params.Add("sign", genSign(params, o.GetConfig().AppSecret))

	uri := o.GetConfig().URL + "/open/finanv0/getUserTokens"
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
