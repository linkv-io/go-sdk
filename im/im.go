package im

import (
	. "github.com/linkv-io/go-sdk/config"
)

func New() *im {
	return &im{}
}

type im struct {
}

func (o *im) GetConfig() *IMConfig {
	return &Conf.IM
}
