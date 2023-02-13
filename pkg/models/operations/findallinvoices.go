package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllInvoicesQueryParams struct {
	ExternalCustomerID *string `queryParam:"style=form,explode=true,name=external_customer_id"`
	IssuingDateFrom    *string `queryParam:"style=form,explode=true,name=issuing_date_from"`
	IssuingDateTo      *string `queryParam:"style=form,explode=true,name=issuing_date_to"`
	Page               *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage            *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Status             *string `queryParam:"style=form,explode=true,name=status"`
}

type FindAllInvoicesRequest struct {
	QueryParams FindAllInvoicesQueryParams
}

type FindAllInvoicesResponse struct {
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Invoices                *shared.Invoices
	StatusCode              int64
}
