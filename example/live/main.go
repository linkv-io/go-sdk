package main

import (
	"fmt"
	linkv "github.com/linkv-io/go-sdk"
	li "github.com/linkv-io/go-sdk/live"
)

func main() {
	appID := ""
	appSecret := ""
	if err := linkv.Init(appID, appSecret, 30, 10); err != nil {
		fmt.Println(err)
		return
	}

	// 初始化 live对象
	live := linkv.NewLvLIVE()

	thirdUID := "503543"
	aID := "test"
	// 进行帐号绑定
	token, openID, err := live.GetTokenByThirdUID(thirdUID, aID, "test-go",
		li.SexUnknown, "http://xxx.xx.xxx/app/rank-list/static/img/defaultavatar.cd935fdb.png",
		"", "", "")
	if err != nil {
		panic("live.GetTokenByThirdUID(" + err.Error() + ")")
	}

	fmt.Printf("token:%v\n", token)
	return
	// 获取金币余额
	golds0, err := live.GetGoldByLiveOpenID(openID)
	if err != nil {
		panic("live.GetGoldByLiveOpenID(" + err.Error() + ")")
	}
	fmt.Printf("golds0:%v\n", golds0)

	// 完成订单
	orderID := ""
	gold := int64(10)
	golds1, err := live.SuccessOrderByLiveOpenID(openID, li.OrderTypeAdd, gold, 10, 1, li.PlatformH5, orderID)
	if err != nil {
		panic("live.SuccessOrderByLiveOpenID(" + err.Error() + ")")
	}
	if (golds0 + gold) != golds1 {

		panic("(golds0+gold) != golds1")
	}
	fmt.Printf("golds1:%v\n", golds1)

	// 修改金币
	ok, err := live.ChangeGoldByLiveOpenID(openID, li.OrderTypeDel, gold, 1, "测试删除")
	if err != nil {
		panic("live.ChangeGoldByLiveOpenID(" + err.Error() + ")")
	}
	if !ok {
		panic("!ok")
	}

	// 获取金币余额
	golds2, err := live.GetGoldByLiveOpenID(openID)
	if err != nil {
		panic("live.GetGoldByLiveOpenID(" + err.Error() + ")")
	}
	fmt.Printf("golds2:%v\n", golds2)
	if golds0 != golds2 {
		panic("golds0 != golds2")
	}
	fmt.Println("success")
}
