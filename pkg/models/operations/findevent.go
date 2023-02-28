package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindEventPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FindEventRequest struct {
	PathParams FindEventPathParams
}

type FindEventResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Event                   *shared.Event
	StatusCode              int
}
