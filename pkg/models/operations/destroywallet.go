package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DestroyWalletPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DestroyWalletRequest struct {
	PathParams DestroyWalletPathParams
}

type DestroyWalletResponse struct {
	APIResponseNotAllowed   *shared.APIResponseNotAllowed
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int64
	Wallet                  *shared.Wallet
}
