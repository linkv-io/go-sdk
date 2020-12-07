package linkv_sdk

import (
	"github.com/linkv-io/go-sdk/config"
	"github.com/linkv-io/go-sdk/http"
	"github.com/linkv-io/go-sdk/im"
	"github.com/linkv-io/go-sdk/live"
	"github.com/linkv-io/go-sdk/rtc"
)

type SexType = live.SexType
type PlatformType = live.PlatformType

var (
	OrderTypeAdd = live.OrderTypeAdd
	OrderTypeDel = live.OrderTypeDel

	PlatformTypeH5      = live.PlatformH5
	PlatformTypeANDROID = live.PlatformANDROID
	PlatformTypeIOS     = live.PlatformIOS

	SexTypeUnknown = live.SexUnknown
	SexTypeFemale  = live.SexFemale
	SexTypeMale    = live.SexMale
)

type LvLIVE interface {
	GetTokenByThirdUID(thirdUID, aID, userName string, sex SexType, portraitURI, userEmail, countryCode, birthday string) (liveToken, liveOpenID string, err error)
	SuccessOrderByLiveOpenID(liveOpenID string, orderType, gold, money, expr int64, platformType PlatformType, orderID string) (golds int64, err error)
	ChangeGoldByLiveOpenID(liveOpenID string, orderType, gold, expr int64, optionalReason string) (ok bool, err error)
	GetGoldByLiveOpenID(liveOpenID string) (golds int64, err error)
}

func NewLvLIVE() LvLIVE {
	return live.New()
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
