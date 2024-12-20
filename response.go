package model

import "fmt"

type Response struct {
	Errno   int64  `json:"errno"`
	ErrMsg  string `json:"errmsg"`
	TraceID string `json:"traceid"`
}

func (r Response) ErrorMsg() string {
	return fmt.Sprintf("错误码：%d, 错误信息: %s, traceID: %s", r.Errno, r.ErrMsg, r.TraceID)
}
