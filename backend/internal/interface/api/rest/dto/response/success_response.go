package response

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Message: "OK",
		Data:    data,
	}
}
