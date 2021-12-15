package main

import (
	"fmt"
	linkv "github.com/linkv-io/go-sdk"
)

func main() {
	appID := "QgVsTkmymrJNqUgJLwJJPrdjCjvUsYmM"
	appSecret := "4DC9389C0205ABEA52395210B30F905FBC4BBAB32D27746851400834A2345631C4B870230F4C04B83FF67D2220AEA72F04E1CE87757275BEFCA314D84EEAEBF7DE700423C533B6C669EB81F5A77F521EE3297ACB399745BA9294125E0AC13BF5DE1BE707A24206986AE3DA684644AD867119FBE5245C64860242D58A785C4DBB7BADEF46F85105B713A0DFE7DC3BB0E490DD351ECE66413295157AF2FF1BFD30D8DEBF44B8C782F6312D54361380B89DE43C5E5DA9758E525A3585C860B4BC9E44F6B2BDD517F2CB52E4BDDCEC5DCEDADB6F45C9DCF472DC7613BD61607A689387DDCD76D6FD2A89790C4A77138C81092089AFBEBAE9279B9CB78C8CFD5F250FC4DC7EFA106F41E48EBEA897AA29F34A40334FCA2879188FA1076C0CEABF2249A70E501E3448FE33E6414A7B80FB179AEF970D8CBCBF562FA927E477A887148D9F5618E8D62A4575AB19E7696EB05BB161114C588212F0D999397BD6ADAA64DBC2FCA44BDD533DA5CEBB1D8FE87FF1AEC9752AA0772926E897E8B183C69D0ADD39823DA83D2B7A200F04DC47C9BFE858"
	if err := linkv.Init(appID, appSecret, 30, 10); err != nil {
		fmt.Println(err)
		return
	}

	// 初始化 live对象
	im := linkv.NewLvIM()

	thirdUID := "golang"
	thirdToken, err := im.GetTokenByThirdUID(thirdUID)
	if err != nil {
		panic("im.GetTokenByThirdUID(" + err.Error() + ")")
	}
	fmt.Printf("token:%s\n", thirdToken)
	toUID := "123456"
	content := "77881122"
	var requestID string

	objectName := "000"
	if requestID, err = im.PushConverseData(thirdToken, thirdUID, toUID, content, objectName, "", ""); err != nil {
		panic("im.PushConverseData(" + err.Error() + ")")
	}
	fmt.Println("PushConverseData", requestID)

	var failUIDs []string
	if requestID, failUIDs, err = im.PushConverseDatas(thirdToken, thirdUID, []string{toUID}, content, objectName); err != nil {
		panic("im.PushConverseDatas(" + err.Error() + ")")
	}
	fmt.Println("failUIDs", failUIDs)
	fmt.Println("PushConverseDatas", requestID)

	objectName = "xxx"
	if requestID, err = im.SendEventMsg(thirdToken, thirdUID, toUID, content, objectName); err != nil {
		panic("im.SendEventMsg(" + err.Error() + ")")
	}
	fmt.Println("SendEventMsg", requestID)

	roomID := "123"
	objectName = "yyy"
	if requestID, err = im.SendRoomMessage(thirdToken, thirdUID, roomID, content, objectName); err != nil {
		panic("im.SendRoomMessage(" + err.Error() + ")")
	}
	fmt.Println("SendRoomMessage", requestID)

	if err = im.UserBlock(thirdToken, thirdUID, []string{"1234444"}, 1); err != nil {
		panic("im.UserBlock(" + err.Error() + ")")
	}
	fmt.Println("UserBlock ok")

	var bOnline bool

	if bOnline, requestID, err = im.UserStatus(thirdToken, thirdUID, toUID); err != nil {
		panic("im.UserStatus(" + err.Error() + ")")
	}
	fmt.Println("bOnline", bOnline)
	fmt.Println("UserStatus", requestID)
}
