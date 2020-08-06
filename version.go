package linkv_sdk

import "github.com/linkv-io/go-sdk/http"

var version = "0.4.2"

func init() {
	http.SetVersion(version)
}
