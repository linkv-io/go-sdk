package live

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/linkv-io/go-sdk/http"
)

func (o *live) GetTokenByThirdUID(thirdUID, aID, userName string, sex int, portraitURI, userEmail, countryCode, birthday string) (string, string, error) {
	params := url.Values{}
	nonce := genRandomString()
	params.Add("nonce_str", nonce)
	params.Add("app_id", o.GetConfig().AppKey)

	params.Add("userId", thirdUID)
	params.Add("aid", aID)

	if len(userName) > 0 {
		params.Add("name", userName)
	}
	if len(portraitURI) > 0 {
		params.Add("portraitUri", portraitURI)
	}
	if len(userEmail) > 0 {
		params.Add("email", userEmail)
	}
	if len(countryCode) > 0 {
		params.Add("countryCode", countryCode)
	}
	if len(birthday) > 0 {
		params.Add("birthday", birthday)
	}

	if sex != SexTypeUnknown {
		params.Add("sex", strconv.Itoa(sex))
	}

	params.Add("sign", genSign(params, o.GetConfig().AppSecret))

	uri := o.GetConfig().URL + "/open/v0/thGetToken"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, nil)
		if err != nil {
			return "", "", err
		}

		if resp.StatusCode != 200 {
			return "", "", fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			Status int    `json:"status,string"`
			Msg    string `json:"msg"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return "", "", err
		}
		if result.Status != 200 {
			if result.Status == 500 {
				errResult = fmt.Errorf("message(%v)", result.Msg)
				time.Sleep(waitTime)
				continue
			}
			return "", "", fmt.Errorf("message(%v)", result.Msg)
		}
		var resultData struct {
			Data struct {
				Token  string `json:"token"`
				OpenID string `json:"openId"`
			} `json:"data"`
		}
		if err := json.Unmarshal(jsonData, &resultData); err != nil {
			return "", "", err
		}
		return resultData.Data.Token, resultData.Data.OpenID, nil
	}
	return "", "", errResult
}
