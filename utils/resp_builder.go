package utils

type Resp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Result struct {
	Data interface{}
	Err  error
}

type EmptyResp struct{}
