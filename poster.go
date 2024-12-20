package model

// swagger:model
type QrcodeResponse struct {
	Response
	Data struct {
		//生成的二维码链接
		//example: https://example.com/img.jpg
		CodeLink string `json:"code_link"`
	} `json:"data"`
}

// swagger:model
type PosterResponse struct {
	Response
	Data struct {
		//生成的海报链接
		//example: https://example.com/img.jpg
		PosterLink string `json:"poster_link"`
	} `json:"data"`
}
