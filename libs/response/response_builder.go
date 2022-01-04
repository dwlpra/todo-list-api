package response

type RespModel struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Result struct {
	Data interface{}
	Err  error
}

type EmptyResp struct{}
