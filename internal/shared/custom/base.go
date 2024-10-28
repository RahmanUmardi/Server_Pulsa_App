package custom

type BaseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ReportRes struct {
	Message string `json:"message" example:"transaction report downloaded"`
}
