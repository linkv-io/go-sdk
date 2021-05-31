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
	appID := ""
	appSecret := ""
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
