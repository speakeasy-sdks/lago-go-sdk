package shared

type APIResponseUnauthorized struct {
	Error  *string `json:"error,omitempty"`
	Status *int32  `json:"status,omitempty"`
}
