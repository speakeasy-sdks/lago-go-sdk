package shared

type BillableMetricInputBillableMetricAggregationTypeEnum string

const (
	BillableMetricInputBillableMetricAggregationTypeEnumCountAgg          BillableMetricInputBillableMetricAggregationTypeEnum = "count_agg"
	BillableMetricInputBillableMetricAggregationTypeEnumSumAgg            BillableMetricInputBillableMetricAggregationTypeEnum = "sum_agg"
	BillableMetricInputBillableMetricAggregationTypeEnumMaxAgg            BillableMetricInputBillableMetricAggregationTypeEnum = "max_agg"
	BillableMetricInputBillableMetricAggregationTypeEnumUniqueCountAgg    BillableMetricInputBillableMetricAggregationTypeEnum = "unique_count_agg"
	BillableMetricInputBillableMetricAggregationTypeEnumRecurringCountAgg BillableMetricInputBillableMetricAggregationTypeEnum = "recurring_count_agg"
)

type BillableMetricInputBillableMetric struct {
	AggregationType *BillableMetricInputBillableMetricAggregationTypeEnum `json:"aggregation_type,omitempty"`
	Code            *string                                               `json:"code,omitempty"`
	Description     *string                                               `json:"description,omitempty"`
	FieldName       *string                                               `json:"field_name,omitempty"`
	Group           *BillableMetricGroup                                  `json:"group,omitempty"`
	Name            *string                                               `json:"name,omitempty"`
}

type BillableMetricInput struct {
	BillableMetric *BillableMetricInputBillableMetric `json:"billable_metric,omitempty"`
}
