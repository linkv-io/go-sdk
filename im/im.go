package im

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	. "github.com/linkv-io/go-sdk/config"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"
)

func New(operatorID string) *im {
	obj := &im{operatorID: operatorID}
	obj.nonce, obj.timestamp, obj.cmimToken = genCore(obj.GetConfig().AppSecret)
	go func() {
		waitTime := time.Hour*24*7 - time.Second*5
		t := time.NewTimer(waitTime)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				break
			}
			obj.nonce, obj.timestamp, obj.cmimToken = genCore(obj.GetConfig().AppSecret)
			t.Reset(waitTime)
		}
	}()
	return obj
}

type im struct {
	operatorID string
	nonce      string
	timestamp  string
	cmimToken  string
}

func (o *im) GetConfig() *IMConfig {
	im := &Conf.IM
	return im
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

func genCore(secret string) (string, string, string) {
	nonce := genGUID()
	timestamp := getTimestampS()
	arr := []string{nonce, timestamp, secret}
	sort.Strings(arr)
	md5Data := md5.Sum([]byte(strings.Join(arr, "")))
	cmimToken := strings.ToLower(hex.EncodeToString(md5Data[:]))
	return nonce, timestamp, cmimToken
}
