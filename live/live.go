package live

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	. "github.com/linkv-io/go-sdk/config"
)

var (
	OrderTypeAdd int64 = 1
	OrderTypeDel int64 = 2

	SexTypeUnknown int = -1
	SexTypeFemale  int = 0
	SexTypeMale    int = 1

	PlatformTypeH5      string = "h5"
	PlatformTypeANDROID string = "android"
	PlatformTypeIOS     string = "ios"

	waitTime = time.Millisecond * 300
)

func New() *live {
	return &live{}
}

type live struct {
}

func (o *live) GetConfig() *LiveConfig {
	return &Conf.Live
}

func genUniqueIDString(appKey string) string {
	nLen := 9
	container := string([]byte(appKey)[2:]) + "-"
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

func genRandomString() string {
	nLen := 16
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < nLen; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
		if i == 7 {
			container += strconv.FormatInt(time.Now().Unix(), 10)
		}
	}
	return container
}

func genSign(params url.Values, md5Secret string) string {
	data := encode(params) + "&key=" + md5Secret
	md5Data := md5.Sum([]byte(data))
	return strings.ToLower(hex.EncodeToString(md5Data[:]))
}

func encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(vs[0])
	}
	return buf.String()
}
