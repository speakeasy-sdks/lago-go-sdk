package shared

type APIResponseNotAllowed struct {
	Code   *string `json:"code,omitempty"`
	Error  *string `json:"error,omitempty"`
	Status *int    `json:"status,omitempty"`
}
