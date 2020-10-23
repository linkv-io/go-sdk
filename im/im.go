package im

import (
	"bytes"
	"crypto/rand"
	. "github.com/linkv-io/go-sdk/config"
	"math/big"
	"strconv"
	"time"
)

func New() *im {
	return &im{}
}

type im struct {
}

func (o *im) GetConfig() *IMConfig {
	return &Conf.IM
}

func genRandomString(nLen int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < nLen; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func genGUID() string {
	return genRandomString(9) + "-" + genRandomString(4) + "-" + genRandomString(4) + "-" + genRandomString(12)
}

func getTimestampS() string {
	t := time.Now()
	return strconv.FormatInt(t.Unix(), 10)
}

func getTimestampMS() string {
	t := time.Now()
	return strconv.FormatInt(t.Unix()*1000+t.UnixNano(), 10)
}
