package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FetchPublicKeyResponse struct {
	APIResponseUnauthorized          *shared.APIResponseUnauthorized
	ContentType                      string
	StatusCode                       int
	FetchPublicKey200TextPlainString *string
}
