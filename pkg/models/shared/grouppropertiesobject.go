package shared

type GroupPropertiesObject struct {
	GroupID *string                `json:"group_id,omitempty"`
	Values  map[string]interface{} `json:"values,omitempty"`
}
