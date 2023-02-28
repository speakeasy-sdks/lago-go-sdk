package lago

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
)

type webhooks struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newWebhooks(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *webhooks {
	return &webhooks{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// FetchPublicKey - Fetch webhook public key
// Webhook public key
func (s *webhooks) FetchPublicKey(ctx context.Context) (*operations.FetchPublicKeyResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/webhooks/public_key/"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.FetchPublicKeyResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `text/plain`):
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			out := string(data)
			res.FetchPublicKey200TextPlainString = &out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.APIResponseUnauthorized
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.APIResponseUnauthorized = out
		}
	}

	return res, nil
}
