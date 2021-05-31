package main

import (
	"fmt"
	linkv "github.com/linkv-io/go-sdk"
)

func main() {
	appID := ""
	appSecret := ""
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
	objectName := "RC:textMsg"
	content := "77881122"
	ok, err := im.PushConverseData(thirdUID, toUID, objectName, content, "", "", "", "", "", "")
	if err != nil {
		panic("im.PushConverseData(" + err.Error() + ")")
	}
	fmt.Println(ok)
}
