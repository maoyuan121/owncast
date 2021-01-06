package models

// BaseAPIResponse 是一个简单的对 API 请求的响应体
type BaseAPIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
