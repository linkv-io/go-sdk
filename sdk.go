package linkv_sdk

import (
	"github.com/linkv-io/go-sdk/config"
	"github.com/linkv-io/go-sdk/http"
)

func Init(appID, appSecret string, httpTimeout int64, httpPoolSize int) error {
	if err := http.Init(httpTimeout, httpPoolSize); err != nil {
		return err
	}
	return config.Conf.Init(appID, appSecret)
}
