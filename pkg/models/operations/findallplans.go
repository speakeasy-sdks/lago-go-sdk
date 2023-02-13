package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllPlansQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllPlansRequest struct {
	QueryParams FindAllPlansQueryParams
}

type FindAllPlansResponse struct {
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Plans                   *shared.Plans
	StatusCode              int64
}
