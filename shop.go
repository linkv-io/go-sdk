package linkv_sdk

import (
	"github.com/linkv-io/go-sdk/shop"
)

type LvSHOP interface {
	GetTokenByThirdUID(thirdUID, aID, userName string, sex shop.SexType, portraitURI, userEmail, countryCode, birthday string) (liveToken, liveOpenID string, err error)
}

func NewLvSHOP() LvSHOP {
	return shop.New()
}
