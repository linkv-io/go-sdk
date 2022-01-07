package linkv_sdk

import (
	"github.com/linkv-io/go-sdk/live"
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
