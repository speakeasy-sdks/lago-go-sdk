package shared

type GroupObject struct {
	Key    *string `json:"key,omitempty"`
	LagoID *string `json:"lago_id,omitempty"`
	Value  *string `json:"value,omitempty"`
}
