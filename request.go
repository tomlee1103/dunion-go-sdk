package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	consts "github.com/dunion-openapi-sdk/dunion-go-sdk/const"
	"github.com/dunion-openapi-sdk/dunion-go-sdk/model"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"time"
)

const (
	AppKey    = "App-Key"
	Timestamp = "Timestamp"
	Sign      = "Sign"
)

var globalTimeoutDuration = 2 * time.Second

func SetTimeoutDuration(timeout time.Duration) {
	globalTimeoutDuration = timeout
}

func Post(ctx context.Context, appKey, accessKey, url string, body map[string]interface{}, opt ...model.Option) ([]byte, error) {
	header := map[string]string{
		AppKey:    appKey,
		Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
	}
	params := make(map[string]interface{})
	for k, v := range body {
		params[k] = v
	}
	for k, v := range header {
		params[k] = v
	}
	header[Sign] = GetSign(params, accessKey)

	bodyBytes, _ := json.Marshal(body)
	reqReader := bytes.NewReader(bodyBytes)
	req, _ := http.NewRequest("POST", url, reqReader)
	traceID := uuid4()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(consts.TraceID, traceID)
	req.Header.Set(consts.UserAgent, consts.SDKVersion)
	if unionLogger != nil {
		unionLogger.Infof("url=%s||headers=%v||data=%v", url, req.Header, body)
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	timeout := globalTimeoutDuration
	if len(opt) > 0 {
		timeout = opt[0].Timeout
	}
	client := &http.Client{Timeout: timeout}
	response, err := client.Do(req)
	if err != nil {
		if unionLogger != nil {
			unionLogger.Errorf("url=%s||headers=%v||data=%v||err=%v", url, req.Header, body, err)
		}
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

func Get(ctx context.Context, appKey, accessKey, url string, param map[string]interface{}, opt ...model.Option) ([]byte, error) {
	header := map[string]string{
		AppKey:    appKey,
		Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
	}
	params := make(map[string]interface{})
	for k, v := range param {
		params[k] = v
	}
	for k, v := range header {
		params[k] = v
	}
	header[Sign] = GetSign(params, accessKey)
	query := netUrl.Values{}
	for k, v := range params {
		query.Add(k, fmt.Sprintf("%v", v))
	}
	if query.Encode() != "" {
		url = url + "?" + query.Encode()
	}

	req, _ := http.NewRequest("GET", url, nil)
	traceID := uuid4()
	req.Header.Set(consts.TraceID, traceID)
	req.Header.Set(consts.UserAgent, consts.SDKVersion)
	if unionLogger != nil {
		unionLogger.Infof("url=%s||headers=%v||data=%v", url, req.Header, param)
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	timeout := globalTimeoutDuration
	if len(opt) > 0 {
		timeout = opt[0].Timeout
	}
	client := &http.Client{Timeout: timeout}
	response, err := client.Do(req)
	if err != nil {
		if unionLogger != nil {
			unionLogger.Errorf("url=%s||headers=%v||data=%v||err=%v", url, req.Header, param, err)
		}
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}
