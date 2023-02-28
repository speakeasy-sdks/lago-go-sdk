package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindWalletPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FindWalletRequest struct {
	PathParams FindWalletPathParams
}

type FindWalletResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
	Wallet                  *shared.Wallet
}
