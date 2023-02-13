package shared

type BillableMetricGroup struct {
	Key    *string       `json:"key,omitempty"`
	Values []interface{} `json:"values,omitempty"`
}
