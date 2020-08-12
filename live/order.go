package live

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/linkv-io/go-sdk/http"
)

func (o *live) SuccessOrderByLiveOpenID(liveOpenID string, orderType, gold, money, expr int64, platformType, optionalOrderID string) (int64, error) {
	params := url.Values{}
	nonce := genRandomString()
	params.Add("nonce_str", nonce)
	params.Add("app_id", o.GetConfig().AppKey)

	params.Add("uid", liveOpenID)
	params.Add("request_id", genUniqueIDString())
	params.Add("type", strconv.FormatInt(orderType, 10))
	params.Add("value", strconv.FormatInt(gold, 10))
	params.Add("money", strconv.FormatInt(money, 10))
	params.Add("expriation", strconv.FormatInt(time.Now().Unix()+int64(expr)*86400, 10))
	params.Add("channel", o.GetConfig().Alias)
	params.Add("platform", platformType)

	if len(optionalOrderID) > 0 {
		params.Add("order_id", optionalOrderID)
	}

	params.Add("sign", genSign(params, o.GetConfig().AppSecret))

	uri := o.GetConfig().URL + "/open/finanv0/orderSuccess"
	var errResult error
	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, nil)
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
		if result.Status != 200 {
			if result.Status == 500 {
				errResult = fmt.Errorf("message(%v)", result.Msg)
				time.Sleep(waitTime)
				continue
			}
			return 0, fmt.Errorf("message(%v)", result.Msg)
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
	return 0, errResult
}
