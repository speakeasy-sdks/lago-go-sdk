package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type UpdateWalletPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateWalletRequest struct {
	PathParams UpdateWalletPathParams
	Request    shared.WalletUpdateInput `request:"mediaType=application/json"`
}

type UpdateWalletResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	StatusCode                     int
	Wallet                         *shared.Wallet
}
