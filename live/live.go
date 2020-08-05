package im

import (
	. "github.com/linkv-io/go-sdk/config"
)

type LvLIVE interface {
}

func NewLvLIVE() LvLIVE {
	return &linkVLIVE{}
}

type linkVLIVE struct {
}

func (im *linkVLIVE) GetConfig() *IMConfig {
	return &Conf.IM
}
