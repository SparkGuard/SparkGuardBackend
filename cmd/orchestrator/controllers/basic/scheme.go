package basic

type DefaultSuccessResponse struct {
	Message string `json:"message"`
}

type DefaultErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
