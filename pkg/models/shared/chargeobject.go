package shared

type ChargeObjectChargeModelEnum string

const (
	ChargeObjectChargeModelEnumStandard   ChargeObjectChargeModelEnum = "standard"
	ChargeObjectChargeModelEnumGraduated  ChargeObjectChargeModelEnum = "graduated"
	ChargeObjectChargeModelEnumPackage    ChargeObjectChargeModelEnum = "package"
	ChargeObjectChargeModelEnumPercentage ChargeObjectChargeModelEnum = "percentage"
	ChargeObjectChargeModelEnumVolume     ChargeObjectChargeModelEnum = "volume"
)

type ChargeObject struct {
	ChargeModel          *ChargeObjectChargeModelEnum `json:"charge_model,omitempty"`
	CreatedAt            *string                      `json:"created_at,omitempty"`
	GroupProperties      []GroupPropertiesObject      `json:"group_properties,omitempty"`
	LagoBillableMetricID *string                      `json:"lago_billable_metric_id,omitempty"`
	LagoID               *string                      `json:"lago_id,omitempty"`
	Properties           map[string]interface{}       `json:"properties,omitempty"`
}
