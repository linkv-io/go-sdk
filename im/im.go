package im

import (
	. "github.com/linkv-io/go-sdk/config"
)

type LvIM interface {
}

func NewLvIM() LvIM {
	return &linkVIM{}
}

type linkVIM struct {
}

func (im *linkVIM) GetConfig() *IMConfig {
	return &Conf.IM
}
