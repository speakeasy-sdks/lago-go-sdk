package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAddOnPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type FindAddOnRequest struct {
	PathParams FindAddOnPathParams
}

type FindAddOnResponse struct {
	AddOn                   *shared.AddOn
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int64
}
