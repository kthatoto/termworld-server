package commands

type Response struct {
	RequestId string `json:"requestId"`
	Success   bool   `json:"success"`
	Message   string `json:"message"`
}
