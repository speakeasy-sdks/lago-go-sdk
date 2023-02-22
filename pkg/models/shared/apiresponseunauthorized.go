package shared

type APIResponseUnauthorized struct {
	Error  *string `json:"error,omitempty"`
	Status *int    `json:"status,omitempty"`
}
