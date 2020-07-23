package im

import (
	. "github.com/linkv-io/go-sdk/config"
)

type IM struct {
}

func (im *IM) GetConfig() *IMConfig {
	return &Conf.IM
}
