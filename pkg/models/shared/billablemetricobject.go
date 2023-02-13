package shared

type BillableMetricObjectAggregationTypeEnum string

const (
	BillableMetricObjectAggregationTypeEnumCountAgg          BillableMetricObjectAggregationTypeEnum = "count_agg"
	BillableMetricObjectAggregationTypeEnumSumAgg            BillableMetricObjectAggregationTypeEnum = "sum_agg"
	BillableMetricObjectAggregationTypeEnumMaxAgg            BillableMetricObjectAggregationTypeEnum = "max_agg"
	BillableMetricObjectAggregationTypeEnumUniqueCountAgg    BillableMetricObjectAggregationTypeEnum = "unique_count_agg"
	BillableMetricObjectAggregationTypeEnumRecurringCountAgg BillableMetricObjectAggregationTypeEnum = "recurring_count_agg"
)

type BillableMetricObject struct {
	ActiveSubscriptionsCount *int64                                   `json:"active_subscriptions_count,omitempty"`
	AggregationType          *BillableMetricObjectAggregationTypeEnum `json:"aggregation_type,omitempty"`
	Code                     *string                                  `json:"code,omitempty"`
	CreatedAt                *string                                  `json:"created_at,omitempty"`
	Description              *string                                  `json:"description,omitempty"`
	DraftInvoicesCount       *int64                                   `json:"draft_invoices_count,omitempty"`
	FieldName                *string                                  `json:"field_name,omitempty"`
	Group                    *BillableMetricGroup                     `json:"group,omitempty"`
	LagoID                   *string                                  `json:"lago_id,omitempty"`
	Name                     *string                                  `json:"name,omitempty"`
}
