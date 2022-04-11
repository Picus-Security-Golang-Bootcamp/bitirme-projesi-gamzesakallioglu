package api

type ApiResponse struct {
	Code int64 `json:"code,omitempty"`

	Details interface{} `json:"details,omitempty"`

	Message string `json:"message,omitempty"`
}
