package responses

type BaseResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	RequestId string `json:"requestId"`
	HasMore bool `json:"hasMore"`
}


