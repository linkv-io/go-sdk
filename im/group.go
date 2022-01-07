package im

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/linkv-io/go-sdk/http"
	"net/url"
	"strconv"
	"strings"
)

func (o *im) GroupCreate(userID, nickName, groupID, groupName string) (string, error) {
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
	params.Set("groupId", groupID)
	params.Set("groupName", groupName)
	params.Set("nickName", nickName)

	uri := o.GetConfig().URL + "/api/rest/group/create"

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

func (o *im) GroupJoin(userID, nickName, groupID string) (string, error) {
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
	params.Set("groupId", groupID)
	params.Set("nickName", nickName)

	uri := o.GetConfig().URL + "/api/rest/group/join"

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

func (o *im) GroupExit(userID, groupID string) (string, error) {
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
	params.Set("groupId", groupID)

	uri := o.GetConfig().URL + "/api/rest/group/exit"

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

func (o *im) GroupInfo(groupID string) (*GroupInfo, string, error) {
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

	params.Set("groupId", groupID)

	uri := o.GetConfig().URL + "/api/rest/group/groupInfo"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return nil, "", err
		}

		if resp.StatusCode != 200 {
			return nil, "", fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			ID   string `json:"requestID"`
			Code int    `json:"code"`
			Data struct {
				Basic  Group `json:"basicInfo"`
				Attach struct {
					AppID  string `json:"appid"`
					GID    string `json:"gid"`
					Topic  string `json:"topic"`
					MTime  string `json:"modifytime"`
					Author string `json:"author"`
				} `json:"attach"`
			}
			Msg string `json:"msg"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return nil, "", err
		}

		if result.Code == 200 {
			return &GroupInfo{
				Group:  result.Data.Basic,
				Topic:  result.Data.Attach.Topic,
				MTime:  result.Data.Attach.MTime,
				Author: result.Data.Attach.Author,
			}, result.ID, nil
		}

		return nil, "", fmt.Errorf("code not 200(%v) message(%v)", result.Code, result.Msg)
	}
	return nil, "", errResult
}

func (o *im) GroupModifyName(adminUserID, groupID, groupName string) (string, error) {
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

	params.Set("userId", adminUserID)
	params.Set("groupId", groupID)
	params.Set("groupName", groupName)

	uri := o.GetConfig().URL + "/api/rest/group/modifyName"

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

func (o *im) GroupModifyTopic(adminUserID, groupID, topic string) (string, error) {
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

	params.Set("userId", adminUserID)
	params.Set("groupId", groupID)
	params.Set("topic", topic)

	uri := o.GetConfig().URL + "/api/rest/group/modifyTopic"

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

func (o *im) GroupModifyAdmin(adminUserID, toUserID, groupID string) (string, error) {
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

	params.Set("toUserId", toUserID)
	params.Set("groupId", groupID)
	params.Set("adminUserId", adminUserID)

	uri := o.GetConfig().URL + "/api/rest/group/transfer"

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

func (o *im) GroupList(userID string, page, size int) ([]Group, string, error) {
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
	params.Set("page", strconv.Itoa(page))
	params.Set("size", strconv.Itoa(size))

	uri := o.GetConfig().URL + "/api/rest/group/groups"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return nil, "", err
		}

		if resp.StatusCode != 200 {
			return nil, "", fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			ID   string  `json:"requestID"`
			Code int     `json:"code"`
			Data []Group `json:"data"`
			Msg  string  `json:"msg"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return nil, "", err
		}

		if result.Code == 200 {
			return result.Data, result.ID, nil
		}

		return nil, "", fmt.Errorf("code not 200(%v) message(%v)", result.Code, result.Msg)
	}
	return nil, "", errResult
}

func (o *im) GroupUserList(groupID string, page, size int) ([]GroupUser, string, error) {
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

	params.Set("groupId", groupID)
	params.Set("page", strconv.Itoa(page))
	params.Set("size", strconv.Itoa(size))

	uri := o.GetConfig().URL + "/api/rest/group/users"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return nil, "", err
		}

		if resp.StatusCode != 200 {
			return nil, "", fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			ID   string      `json:"requestID"`
			Code int         `json:"code"`
			Data []GroupUser `json:"data"`
			Msg  string      `json:"msg"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return nil, "", err
		}

		if result.Code == 200 {
			return result.Data, result.ID, nil
		}

		return nil, "", fmt.Errorf("code not 200(%v) message(%v)", result.Code, result.Msg)
	}
	return nil, "", errResult
}

func (o *im) GroupDissolve(adminUserID, groupID string, status GroupStatus) (string, error) {
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

	params.Set("userId", adminUserID)
	params.Set("groupId", groupID)
	params.Set("block", strconv.Itoa(int(status)))

	uri := o.GetConfig().URL + "/api/rest/group/dissolve"

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

func (o *im) GroupKickUser(adminUserID, kickUserID, groupID string) (string, error) {
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

	params.Set("groupId", groupID)
	params.Set("fromUserId", adminUserID)
	params.Set("toUserId", kickUserID)

	uri := o.GetConfig().URL + "/api/rest/group/groupKickUser"

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

func (o *im) GroupGagUser(adminUserID, gagUserID, groupID string, status GroupUserStatus) (string, error) {
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

	params.Set("fromUserId", adminUserID)
	params.Set("toUserId", gagUserID)
	params.Set("groupId", groupID)
	params.Set("type", strconv.Itoa(int(status)))

	uri := o.GetConfig().URL + "/api/rest/group/gagUser"

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

func (o *im) GroupGagAllUser(adminUserID, groupID string, status GroupUserStatus) (string, error) {
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

	params.Set("fromUserId", adminUserID)
	params.Set("groupId", groupID)
	params.Set("type", strconv.Itoa(int(status)))

	uri := o.GetConfig().URL + "/api/rest/group/gagAllUser"

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

func (o *im) GroupSetNotifiable(userID, groupID string, status GroupUserStatus) (string, error) {
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

	params.Set("fromUserId", userID)
	params.Set("groupId", groupID)
	params.Set("type", strconv.Itoa(int(status)))

	uri := o.GetConfig().URL + "/api/rest/group/setNotifiable"

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
