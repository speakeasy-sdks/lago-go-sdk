package shared

type APIResponseNotFound struct {
	Code   *string `json:"code,omitempty"`
	Error  *string `json:"error,omitempty"`
	Status *int32  `json:"status,omitempty"`
}
