package util

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strings"
)

func GetSign(params map[string]interface{}, accessKey string) string {
	// key排序
	arr := sort.StringSlice{}
	for k := range params {
		if k != "sign" {
			arr = append(arr, k)
		}
	}
	arr.Sort()
	// 参数拼接
	var build strings.Builder
	for idx, k := range arr {
		if idx != 0 {
			build.WriteString("&")
		}
		build.WriteString(fmt.Sprintf("%s=%v", k, params[k]))
	}
	build.WriteString(accessKey)
	// URL encode
	sourceStr := url.QueryEscape(build.String())
	// sha1加密
	h := sha1.New()
	_, _ = io.WriteString(h, sourceStr)
	shaStr := hex.EncodeToString(h.Sum([]byte("")))
	// 返回base64字符串
	b64Str := base64.StdEncoding.EncodeToString([]byte(shaStr))
	// base64字符串含有=和/，再一次URL encode
	return url.QueryEscape(b64Str)
}
