package main

import (
	"fmt"
	linkv "github.com/linkv-io/go-sdk"
	li "github.com/linkv-io/go-sdk/live"
)

func main() {
	appID := "eWpPVvgCktvjGiAxHeUbZwiSCcHCUGRw"
	appSecret := "7ABDFC88F75EFEF10C6F8D52436A5F2608ADDD0423E41EE5111BFAA24464A3343EA56ED3A62A7D09978456D5AF9957849234B6BEFDF0090E495FE8CC29ED9225D618DABD7E0E7750FAACD09DAEF7219D8B055516E6B1E11C2CC8CCBF387E338AF986877A1CA7EEB14C23EC5E8975B8261F84F222E711BC64FE23304BE5AC60B25C84D7592A353DF92CA8DEECD33F65720CF64EDD2B7C7E95E0FFDA8B19F27E423460D25E46D42E823D33275C15848E9E548FF3EE250CB69F2371882CA7F253349E48241574D590D87323F42E348448C7859D05D3ABB0EA88A9B93B472FA98075A21D9EEAAF916065E92D40F839233C34CE5DADBA024056BCCAD26D369C4643572E00D1B2F1D44F134C0039786FBF5E0C43107CA96D95CBF31C4ACCF45E2BD2F63483F092CA8A1AAF8B67BA5C3D2B7846C6B13E6E78C68A2B2CBD2CE0203C432BF91F78074B585E503A34BFFC51BCFF6B73FA24E0F9810A71A41D12BDBE2E6E257C43971204F443A762520B7855777A4E747B7915214316144FFA00AE9E96F294"
	if err := linkv.Init(appID, appSecret, 30, 10); err != nil {
		fmt.Println(err)
		return
	}

	// 初始化 live对象
	live := linkv.NewLvLIVE()

	thirdUID := "test-go-tob"
	aID := "test"
	// 进行帐号绑定
	token, openID, err := live.GetTokenByThirdUID(thirdUID, aID, "test-go",
		li.SexUnknown, "http://meet.linkv.sg/app/rank-list/static/img/defaultavatar.cd935fdb.png",
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
