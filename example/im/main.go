package main

import (
	"fmt"
	linkv "github.com/linkv-io/go-sdk"
)

func main() {
	appID := "QgVsTkmymrJNqUgJLwJJPrdjCjvUsYmM"
	appSecret := "4DC9389C0205ABEA52395210B30F905FBC4BBAB32D27746851400834A2345631C4B870230F4C04B83FF67D2220AEA72F04E1CE87757275BEFCA314D84EEAEBF7DE700423C533B6C669EB81F5A77F521EE3297ACB399745BA9294125E0AC13BF5DE1BE707A24206986AE3DA684644AD867119FBE5245C64860242D58A785C4DBB7BADEF46F85105B713A0DFE7DC3BB0E490DD351ECE66413295157AF2FF1BFD30D8DEBF44B8C782F6312D54361380B89DE43C5E5DA9758E525A3585C860B4BC9E44F6B2BDD517F2CB52E4BDDCEC5DCEDADB6F45C9DCF472DC7613BD61607A689387DDCD76D6FD2A89790C4A77138C81092089AFBEBAE9279B9CB78C8CFD5F250FC4DC7EFA106F41E48EBEA897AA29F34A40334FCA2879188FA1076C0CEABF2249A70E501E3448FE33E6414A7B80FB179AEF970D8CBCBF562FA927E477A887148D9F5618E8D62A4575AB19E7696EB05BB161114C588212F0D999397BD6ADAA64DB41270B073136A57BFB8B8137F2024154FEA42CBB30597E50BF93BD0A78B98602278F60BE134FD9E118B8EBBA14BA126B"
	if err := linkv.Init(appID, appSecret, 30, 10); err != nil {
		fmt.Println(err)
		return
	}

	// 初始化 live对象
	im := linkv.NewLvIM("123")

	thirdUID := "golang"
	// thirdToken, err := im.GetTokenByThirdUID(thirdUID)
	// if err != nil {
	// 	panic("im.GetTokenByThirdUID(" + err.Error() + ")")
	// }
	// fmt.Printf("token:%s\n", thirdToken)
	var err error
	toUID := "123456"
	content := "77881122"
	var requestID string

	objectName := "000"
	if requestID, err = im.PushConverseData(thirdUID, toUID, content, objectName, "", ""); err != nil {
		panic("im.PushConverseData(" + err.Error() + ")")
	}
	fmt.Println("PushConverseData", requestID)

	var failUIDs []string
	if requestID, failUIDs, err = im.PushConverseDatas(thirdUID, []string{toUID}, content, objectName); err != nil {
		panic("im.PushConverseDatas(" + err.Error() + ")")
	}
	fmt.Println("failUIDs", failUIDs)
	fmt.Println("PushConverseDatas", requestID)

	objectName = "xxx"
	if requestID, err = im.SendEventMsg(thirdUID, toUID, content, objectName); err != nil {
		panic("im.SendEventMsg(" + err.Error() + ")")
	}
	fmt.Println("SendEventMsg", requestID)

	roomID := "123"
	objectName = "yyy"
	if requestID, err = im.SendRoomMessage(thirdUID, roomID, content, objectName); err != nil {
		panic("im.SendRoomMessage(" + err.Error() + ")")
	}
	fmt.Println("SendRoomMessage", requestID)

	if err = im.UserBlock([]string{"1234444"}, 1); err != nil {
		panic("im.UserBlock(" + err.Error() + ")")
	}
	fmt.Println("UserBlock ok")

	var bOnline bool

	if bOnline, requestID, err = im.UserStatus(toUID); err != nil {
		panic("im.UserStatus(" + err.Error() + ")")
	}
	fmt.Println("bOnline", bOnline)
	fmt.Println("UserStatus", requestID)

	if requestID, failUIDs, err = im.AddUserBlack(thirdUID, []string{"123123"}); err != nil {
		panic("im.AddUserBlack(" + err.Error() + ")")
	}
	fmt.Println("AddUserBlack requestID:", requestID, "failUIDs:", failUIDs)

	if requestID, failUIDs, err = im.RemoveUserBlack(thirdUID, []string{"123123"}); err != nil {
		panic("im.RemoveUserBlack" + err.Error() + ")")
	}
	fmt.Println("RemoveUserBlack requestID:", requestID, "failUIDs:", failUIDs)

	groupID := "test9"
	groupName := "test"
	adminUserID := "1"
	adminNickName := "admin"
	if requestID, err = im.GroupCreate(adminUserID, adminNickName, groupID, groupName); err != nil {
		panic("im.GroupCreate" + err.Error() + ")")
	}
	fmt.Println("GroupCreate requestID:", requestID)

	userID := "10"
	nickName := "10a"
	if requestID, err = im.GroupJoin(userID, nickName, groupID); err != nil {
		panic("im.GroupJoin" + err.Error() + ")")
	}
	fmt.Println("GroupJoin requestID:", requestID)

	if requestID, err = im.GroupExit(userID, groupID); err != nil {
		panic("im.GroupExit" + err.Error() + ")")
	}
	fmt.Println("GroupExit requestID:", requestID)

	if requestID, err = im.GroupJoin(userID, nickName, groupID); err != nil {
		panic("im.GroupJoin" + err.Error() + ")")
	}
	fmt.Println("GroupJoin requestID:", requestID)

	var ginfo *linkv.IMGroupInfo
	if ginfo, requestID, err = im.GroupInfo(groupID); err != nil {
		panic("im.GroupInfo" + err.Error() + ")")
	}
	fmt.Println("GroupInfo requestID:", requestID, "info:", *ginfo)

	if requestID, err = im.GroupModifyName(adminUserID, groupID, "bb"); err != nil {
		panic("im.GroupModifyName" + err.Error() + ")")
	}
	fmt.Println("GroupModifyName requestID:", requestID)

	if requestID, err = im.GroupModifyTopic(adminUserID, groupID, "bb"); err != nil {
		panic("im.GroupModifyTopic" + err.Error() + ")")
	}
	fmt.Println("GroupModifyTopic requestID:", requestID)

	fmt.Println("GroupModifyAdmin requestID:", requestID)
	if requestID, err = im.GroupDissolve(adminUserID, groupID, linkv.IMStatusGroupBlock); err != nil {
		panic("im.GroupDissolve" + err.Error() + ")")
	}
	fmt.Println("GroupDissolve requestID:", requestID)

	if requestID, err = im.GroupDissolve(adminUserID, groupID, linkv.IMStatusGroupUnBlock); err != nil {
		panic("im.GroupDissolve" + err.Error() + ")")
	}
	fmt.Println("GroupDissolve requestID:", requestID)

	if requestID, err = im.GroupKickUser(adminUserID, userID, groupID); err != nil {
		panic("im.GroupKickUser" + err.Error() + ")")
	}
	fmt.Println("GroupKickUser requestID:", requestID)

	if requestID, err = im.GroupJoin(userID, nickName, groupID); err != nil {
		panic("im.GroupJoin" + err.Error() + ")")
	}
	fmt.Println("GroupJoin requestID:", requestID)

	if requestID, err = im.GroupGagUser(adminUserID, userID, groupID, linkv.IMStatusGroupUserDisable); err != nil {
		panic("im.GroupGagUser" + err.Error() + ")")
	}
	fmt.Println("GroupGagUser requestID:", requestID)

	if requestID, err = im.GroupGagAllUser(adminUserID, groupID, linkv.IMStatusGroupUserDisable); err != nil {
		panic("im.GroupGagAllUser" + err.Error() + ")")
	}
	fmt.Println("GroupGagAllUser requestID:", requestID)

	if requestID, err = im.GroupModifyAdmin(adminUserID, userID, groupID); err != nil {
		panic("im.GroupModifyAdmin" + err.Error() + ")")
	}
	fmt.Println("GroupModifyAdmin requestID:", requestID)

	if requestID, err = im.GroupModifyAdmin(userID, adminUserID, groupID); err != nil {
		panic("im.GroupModifyAdmin" + err.Error() + ")")
	}

	var gList []linkv.IMGroup
	if gList, requestID, err = im.GroupList(userID, 1, 10); err != nil {
		panic("im.GroupList" + err.Error() + ")")
	}
	fmt.Println("im.GroupList requestID:", requestID, "groups:", gList)

	if requestID, err = im.GroupSetNotifiable(userID, groupID, linkv.IMStatusGroupUserDisable); err != nil {
		panic("im.GroupSetNotifiable" + err.Error() + ")")
	}
	fmt.Println("GroupSetNotifiable requestID:", requestID)

	var guList []linkv.IMGroupUser
	if guList, requestID, err = im.GroupUserList(userID, 1, 10); err != nil {
		panic("GroupUserList" + err.Error() + ")")
	}
	fmt.Println("GroupUserList requestID:", requestID, "group users:", guList)
}
