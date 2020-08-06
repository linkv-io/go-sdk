package linkv_sdk

import (
	"github.com/linkv-io/go-sdk/config"
	"github.com/linkv-io/go-sdk/http"
	"github.com/linkv-io/go-sdk/im"
	"github.com/linkv-io/go-sdk/live"
	"github.com/linkv-io/go-sdk/rtc"
)

type (
	SexType      = live.SexType
	OrderType    = live.OrderType
	PlatformType = live.PlatformType
)

var (
	OrderTypeAdd = live.OrderTypeAdd
	OrderTypeDel = live.OrderTypeDel

	PlatformTypeH5      = live.PlatformTypeH5
	PlatformTypeANDROID = live.PlatformTypeANDROID
	PlatformTypeIOS     = live.PlatformTypeIOS

	SexTypeUnknown = live.SexTypeUnknown
	SexTypeFemale  = live.SexTypeFemale
	SexTypeMale    = live.SexTypeMale
)

type LvLIVE interface {
	GetTokenByThirdUID(thirdUID, aID, userName string, sex SexType, portraitURI, userEmail, countryCode, birthday string) (liveToken, liveOpenID string, err error)
	SuccessOrderByLiveOpenID(liveOpenID, uniqueID string, typ OrderType, gold, money, expr int64, plat PlatformType, optionalOrderID string) (golds int64, err error)
	ChangeGoldByLiveOpenID(liveOpenID, uniqueID string, typ OrderType, gold, expr int64, optionalReason string) (ok bool, err error)
	GetGoldByLiveOpenID(liveOpenID string) (golds int64, err error)
}

func NewLvLIVE() LvLIVE {
	return live.New()
}

type LvIM interface {
}

func NewLvIM() LvIM {
	return im.New()
}

type LvRTC interface {
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
