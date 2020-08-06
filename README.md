[![API Reference](https://img.shields.io/badge/api-reference-blue.svg)]()
[![Build Status](https://img.shields.io/static/v1?label=build&message=passing&color=32CD32)]()
[![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/linkv-io/dart-sdk/blob/master/LICENSE)

# go-sdk

LINKV SDK for the go programming language.

## Requirement

Go 1.8+ (plugin)


## Installing

```sh
go get "github.com/linkv-io/go-sdk"
```

## Usage

```go
package main

import (
	"fmt"
	linkv "github.com/linkv-io/go-sdk"
)

func main() {
	appID := "jQMfhmxPLWKaojdbRKuRWmEvMsuvDnxh"
	appSecret := "CD9570C72C6BE30468E1D207F741F5916A767DD2FAB696755B6B904A19F21CF1C0BF1838EDA612782B77C8BEB931DEB6108E0C83A4227925C899AAFF81A2B2E21AA64E619E9AB4EEA6928630EFD0DEF0D35318572BC298961F9434ACFD73D4765B47707532A81B60229EFBEDED4354947624894FE69A6CCD3E6BBAB276DBEA89E39AF295DA58E13C723399D9F6B9B0085097300BFD20951FABD7FCFDAF106108269EFDA851C89526003C16BB35ADA7419C69E13803D444B109FF02929C877A6F03EB39A0FB7B8E3B868419850F87C8BD9680415A0E55CCB18477C4FA4DBC32BC8F2938BB87DE1DFFA3C5DE19EF89F7A5E643EBC6B89318446825109984675617847D963C6100151DAE894C98DD9AC1CE2EED4B99E2BA12F0A6B379FD6852745881123A0182335F6685AB775F57A99BA7"
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
	uniqueID := "123"
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
	uniqueID1 := "456"
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
```

## License

This SDK is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt and NOTICE.txt for more information.