package shared

type APIResponseUnprocessableEntity struct {
	Code         *string                `json:"code,omitempty"`
	Error        *string                `json:"error,omitempty"`
	ErrorDetails map[string]interface{} `json:"error_details,omitempty"`
	Status       *int32                 `json:"status,omitempty"`
}
