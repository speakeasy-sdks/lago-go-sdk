package shared

type ChargeUsageObjectBillableMetricAggregationTypeEnum string

const (
	ChargeUsageObjectBillableMetricAggregationTypeEnumCountAgg          ChargeUsageObjectBillableMetricAggregationTypeEnum = "count_agg"
	ChargeUsageObjectBillableMetricAggregationTypeEnumSumAgg            ChargeUsageObjectBillableMetricAggregationTypeEnum = "sum_agg"
	ChargeUsageObjectBillableMetricAggregationTypeEnumMaxAgg            ChargeUsageObjectBillableMetricAggregationTypeEnum = "max_agg"
	ChargeUsageObjectBillableMetricAggregationTypeEnumUniqueCountAgg    ChargeUsageObjectBillableMetricAggregationTypeEnum = "unique_count_agg"
	ChargeUsageObjectBillableMetricAggregationTypeEnumRecurringCountAgg ChargeUsageObjectBillableMetricAggregationTypeEnum = "recurring_count_agg"
)

type ChargeUsageObjectBillableMetric struct {
	AggregationType *ChargeUsageObjectBillableMetricAggregationTypeEnum `json:"aggregation_type,omitempty"`
	Code            *string                                             `json:"code,omitempty"`
	LagoID          *string                                             `json:"lago_id,omitempty"`
	Name            *string                                             `json:"name,omitempty"`
}

type ChargeUsageObjectChargeChargeModelEnum string

const (
	ChargeUsageObjectChargeChargeModelEnumStandard   ChargeUsageObjectChargeChargeModelEnum = "standard"
	ChargeUsageObjectChargeChargeModelEnumGraduated  ChargeUsageObjectChargeChargeModelEnum = "graduated"
	ChargeUsageObjectChargeChargeModelEnumPackage    ChargeUsageObjectChargeChargeModelEnum = "package"
	ChargeUsageObjectChargeChargeModelEnumPercentage ChargeUsageObjectChargeChargeModelEnum = "percentage"
	ChargeUsageObjectChargeChargeModelEnumVolume     ChargeUsageObjectChargeChargeModelEnum = "volume"
)

type ChargeUsageObjectCharge struct {
	ChargeModel *ChargeUsageObjectChargeChargeModelEnum `json:"charge_model,omitempty"`
	LagoID      *string                                 `json:"lago_id,omitempty"`
}

type ChargeUsageObjectGroups struct {
	AmountCents *int64   `json:"amount_cents,omitempty"`
	Key         *string  `json:"key,omitempty"`
	LagoID      *string  `json:"lago_id,omitempty"`
	Units       *float64 `json:"units,omitempty"`
	Value       *string  `json:"value,omitempty"`
}

type ChargeUsageObject struct {
	AmountCents    *int64                           `json:"amount_cents,omitempty"`
	AmountCurrency *string                          `json:"amount_currency,omitempty"`
	BillableMetric *ChargeUsageObjectBillableMetric `json:"billable_metric,omitempty"`
	Charge         *ChargeUsageObjectCharge         `json:"charge,omitempty"`
	Groups         []ChargeUsageObjectGroups        `json:"groups,omitempty"`
	Units          *float64                         `json:"units,omitempty"`
}
