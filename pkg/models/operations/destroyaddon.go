package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DestroyAddOnPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type DestroyAddOnRequest struct {
	PathParams DestroyAddOnPathParams
}

type DestroyAddOnResponse struct {
	AddOn                   *shared.AddOn
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
}
