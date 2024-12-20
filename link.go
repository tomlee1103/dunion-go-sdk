package model

//swagger:model
type LinkResponse struct {
	Response
	Data struct {
		//生成的链接
		//example: https://v.didi.cn/p/abcd
		Link string `json:"link"`
		//实例ID，可通过此ID去生成海报或者二维码
		DSI string `json:"dsi"`
		//小程序appid
		AppId string `json:"app_id,omitempty"`
		//小程序原始ID
		AppSource string `json:"app_source,omitempty"`
	} `json:"data"`
}
