package main

import (
	"fmt"
	linkv "github.com/linkv-io/go-sdk"
)

func main() {
	appID := "rbaiHjNHQyVprPCBSHevvVvuNynNeTvp"
	appSecret := "87EA975D424238D0A08F772321169816DD016667D5BB577EBAEB820516698416E4F94C28CB55E9FD8E010260E6C8A177C0B078FC098BCF2E9E7D4A9A71BF1EF8FBE49E05E5FC5A6A35C6550592C1DB96DF83F758EAFBC5B342D5D04C9D92B1A82A76E3756E83A4466DA22635A8A9F88901631B5BBBABC8A94577D66E8B000F4B179DA99BAA5E674E4F793D9E60EEF1C3B757006459ABB5E6315E370461EBC8E6B0A7523CA0032D33B5C0CF83264C9D83517C1C94CAB3F48B8D5062F5569D9793982455277C16F183DAE7B6C271F930A160A6CF07139712A9D3ABF85E05F8721B8BB6CAC1C23980227A1D5F31D23FA6567578AEEB6B124AF8FF76040F9598DDC9DE0DA44EF34BBB01B53E2B4713D2D701A9F913BE56F9F5B9B7D8D2006CA910D8BFA0C34C619AB0EEBDAA474E67115532511686992E88C4E32E86D82736B2FE141E9037381757ED02C7D82CA8FC9245700040D7E1E200029416295D891D388D69AC5197A65121B60D42040393FB42BC2769B1E2F649A7A17083F6AB2B1BE6E993"
	if err := linkv.Init(appID, appSecret, 30, 10); err != nil {
		fmt.Println(err)
		return
	}

	// 初始化 直播对象
	live := linkv.NewLvLIVE()

	thirdUID := "test-go-tob"
	aID := "test"
	// 进行帐号绑定
	token, openID, err := live.GetTokenByThirdUID(thirdUID, aID, "test-go",
		linkv.SexTypeUnknown, "http://meet.linkv.sg/app/rank-list/static/img/defaultavatar.cd935fdb.png",
		"", "", "")
	if err != nil {
		panic("live.GetTokenByThirdUID(" + err.Error() + ")")
	}

	fmt.Printf("token:%v\n", token)
	// 获取金币余额
	golds0, err := live.GetGoldByLiveOpenID(openID)
	if err != nil {
		panic("live.GetGoldByLiveOpenID(" + err.Error() + ")")
	}
	fmt.Printf("golds0:%v\n", golds0)

	// 完成订单
	uniqueID := "123123"
	gold := int64(10)
	golds1, err := live.SuccessOrderByLiveOpenID(openID, uniqueID, linkv.OrderTypeAdd, gold, 10, 1, linkv.PlatformTypeH5, "")
	if err != nil {
		panic("live.SuccessOrderByLiveOpenID(" + err.Error() + ")")
	}
	if (golds0 + gold) != golds1 {

		panic("(golds0+gold) != golds1")
	}
	fmt.Printf("golds1:%v\n", golds1)

	// 修改金币
	uniqueID1 := "456123"
	ok, err := live.ChangeGoldByLiveOpenID(openID, uniqueID1, linkv.OrderTypeDel, gold, 1, "测试删除")
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
