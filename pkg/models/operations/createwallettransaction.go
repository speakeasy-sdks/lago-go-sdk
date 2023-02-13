package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type CreateWalletTransactionRequest struct {
	Request shared.WalletTransactionInput `request:"mediaType=application/json"`
}

type CreateWalletTransactionResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	StatusCode                     int64
	WalletTransactions             *shared.WalletTransactions
}
