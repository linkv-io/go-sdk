package main

import (
	"encoding/json"
	"fmt"
	linkv "github.com/linkv-io/go-sdk"
	"net/http"
)

type AuthResult struct {
	AppID    string `json:"app_id"`
	Auth     string `json:"auth"`
	ExpireTS string `json:"expire_ts"`
}

func main() {
	appID := "rbaiHjNHQyVprPCBSHevvVvuNynNeTvp"
	appSecret := "87EA975D424238D0A08F772321169816DD016667D5BB577EBAEB820516698416E4F94C28CB55E9FD8E010260E6C8A177C0B078FC098BCF2E9E7D4A9A71BF1EF8FBE49E05E5FC5A6A35C6550592C1DB96DF83F758EAFBC5B342D5D04C9D92B1A82A76E3756E83A4466DA22635A8A9F88901631B5BBBABC8A94577D66E8B000F4B179DA99BAA5E674E4F793D9E60EEF1C3B757006459ABB5E6315E370461EBC8E6B0A7523CA0032D33B5C0CF83264C9D83517C1C94CAB3F48B8D5062F5569D9793982455277C16F183DAE7B6C271F930A160A6CF07139712A9D3ABF85E05F8721B8BB6CAC1C23980227A1D5F31D23FA6567578AEEB6B124AF8FF76040F9598DDC9DE0DA44EF34BBB01B53E2B4713D2D701A9F913BE56F9F5B9B7D8D2006CA910D8BFA0C34C619AB0EEBDAA474E67115532511686992E88C4E32E86D82736B2FE141E9037381757ED02C7D82CA8FC9245700040D7E1E200029416295D891D388D69AC5197A65121B60D42040393FB42BC2769B1E2F649A7A17083F6AB2B1BE6E993"
	if err := linkv.Init(appID, appSecret, 30, 10); err != nil {
		fmt.Println(err)
		return
	}

	// 初始化 live对象
	rtc := linkv.NewLvRTC()

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
			w.Write([]byte{})
			return
		}
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		appID, auth, expireTS := rtc.GenAuth()

		data := &AuthResult{appID, auth, expireTS}
		jsonData, _ := json.Marshal(data)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		fmt.Fprintf(w, "%s", jsonData)
	})

	http.ListenAndServe(":8080", nil)
}
