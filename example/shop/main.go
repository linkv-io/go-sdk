package main

import (
	"fmt"
	linkv "github.com/linkv-io/go-sdk"
	"github.com/linkv-io/go-sdk/shop"
)

func main() {
	appID := ""
	appSecret := ""
	if err := linkv.Init(appID, appSecret, 30, 10); err != nil {
		fmt.Println(err)
		return
	}

	// 初始化 live对象
	live := linkv.NewLvSHOP()

	thirdUID := "test-go-tob"
	aID := "test"
	// 进行帐号绑定
	token, openID, err := live.GetTokenByThirdUID(thirdUID, aID, "test-go",
		shop.SexUnknown, "http://xxx.xxx.xxx/app/rank-list/static/img/defaultavatar.cd935fdb.png",
		"", "", "")
	if err != nil {
		panic("live.GetTokenByThirdUID(" + err.Error() + ")")
	}

	fmt.Printf("token:%v openID:%v\n", token, openID)

	fmt.Println("success")
}
