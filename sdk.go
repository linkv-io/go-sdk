package linkv_sdk

import (
	"github.com/linkv-io/go-sdk/config"
	"github.com/linkv-io/go-sdk/http"
	"github.com/linkv-io/go-sdk/im"
	"github.com/linkv-io/go-sdk/live"
	"github.com/linkv-io/go-sdk/rtc"
	"github.com/linkv-io/go-sdk/shop"
)

type LvLIVE interface {
	GetTokenByThirdUID(thirdUID, aID, userName string, sex live.SexType, portraitURI, userEmail, countryCode, birthday string) (liveToken, liveOpenID string, err error)
	SuccessOrderByLiveOpenID(liveOpenID string, orderType, gold, money, expr int64, platformType live.PlatformType, orderID string) (golds int64, err error)
	ChangeGoldByLiveOpenID(liveOpenID string, orderType, gold, expr int64, optionalReason string) (ok bool, err error)
	GetGoldByLiveOpenID(liveOpenID string) (golds int64, err error)
}

func NewLvLIVE() LvLIVE {
	return live.New()
}

type LvSHOP interface {
	GetTokenByThirdUID(thirdUID, aID, userName string, sex shop.SexType, portraitURI, userEmail, countryCode, birthday string) (liveToken, liveOpenID string, err error)
}

func NewLvSHOP() LvSHOP {
	return shop.New()
}

type LvIM interface {
	GetTokenByThirdUID(thirdUID string) (imToken string, err error)
	PushConverseData(fromUID, toUID, objectName, content, pushContent, pushData, deviceID, toAppID, toUserExtSysUserID, isCheckSensitiveWords string) (success bool, err error)
}

func NewLvIM() LvIM {
	return im.New()
}

type LvRTC interface {
	GenAuth() (appID string, authKey string, expireTS string)
}

func NewLvRTC() LvRTC {
	return rtc.New()
}

func Init(appID, appSecret string, httpTimeout int64, httpPoolSize int) error {
	if err := http.Init(httpTimeout, httpPoolSize); err != nil {
		return err
	}
	return config.Conf.Init(appID, appSecret)
}
