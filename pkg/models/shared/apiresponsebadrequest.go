package shared

type APIResponseBadRequest struct {
	Error  *string `json:"error,omitempty"`
	Status *int32  `json:"status,omitempty"`
}
