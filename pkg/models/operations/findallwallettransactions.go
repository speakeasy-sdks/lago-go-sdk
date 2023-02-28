package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllWalletTransactionsPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FindAllWalletTransactionsQueryParams struct {
	Page            *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage         *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Status          *string `queryParam:"style=form,explode=true,name=status"`
	TransactionType *string `queryParam:"style=form,explode=true,name=transaction_type"`
}

type FindAllWalletTransactionsRequest struct {
	PathParams  FindAllWalletTransactionsPathParams
	QueryParams FindAllWalletTransactionsQueryParams
}

type FindAllWalletTransactionsResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
	WalletTransactions      *shared.WalletTransactions
}
