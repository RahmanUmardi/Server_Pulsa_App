package custom

type BaseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
