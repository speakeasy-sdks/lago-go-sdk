<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/lago-go-sdk"
    "github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/lago-go-sdk/pkg/models/operations"
)

func main() {
    opts := []lago.SDKOption{
        lago.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            }
        ),
    }

    s := lago.New(opts...)
    
    req := operations.FindAddOnRequest{
        PathParams: operations.FindAddOnPathParams{
            Code: "unde",
        },
    }
    
    res, err := s.AddOns.FindAddOn(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddOn != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->