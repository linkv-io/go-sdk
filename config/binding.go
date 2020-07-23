package config

import (
	"fmt"
	"plugin"
)

const VERSION = "0.0.4"
const FILE = "decrypt"

var _binding = binding{}

type binding struct {
	core     *plugin.Plugin
	deCrypto func(string, string) (string, error)
}

func (b *binding) init() error {
	p, err := dlopenPlatformSpecific(FILE, "")
	if err != nil {
		return err
	}
	b.core = p

	fun, err := p.Lookup("Decrypt")
	if err != nil {
		panic(err)
	}
	deCrypto, ok := fun.(func(string, string) (string, error))
	if !ok {
		return fmt.Errorf("映射动态链接库函数失败")
	}
	b.deCrypto = deCrypto

	return nil
}
