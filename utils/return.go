package utils

type ResultUtil struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type ResultDataUtil struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}