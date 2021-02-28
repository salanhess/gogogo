package model

type GetBaiduResponse struct {
	Args    BaiduResp   `json:"args"`
	Headers interface{} `json:"headers"`
	Url     string      `json:"url"`
}

type BaiduResp struct {
	Hello string `json:"hello"`
}
