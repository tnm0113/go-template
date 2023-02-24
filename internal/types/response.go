package types

// Response contains HTTP response specific methods.
type Response interface {
	// Code returns HTTP response code.
	Code() int

	// Headers returns map of HTTP headers with their values.
	Headers() map[string]string

	// Empty indicates if HTTP response has content.
	Empty() bool
}

// HTTPError example
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type SucceedResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}
