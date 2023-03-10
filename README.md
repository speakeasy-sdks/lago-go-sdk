# Lago Go SDK

<div align="center">
   <p>The Open Source Metering and Usage Based Billing API.</p>
   <ahref="https://github.com/speakeasy-sdks/lago-go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/lago-go-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
   <a href="https://doc.getlago.com/docs/api/intro"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/lago-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## SDK Available Operations


### AddOns

* `FindAddOn` - Find add-on by code
* `ApplyAddOn` - Apply an add-on to a customer
* `CreateAddOn` - Create a new add-on
* `DestroyAddOn` - Delete an add-on
* `FindAllAddOns` - Find add-ons
* `UpdateAddOn` - Update an existing add-on

### BillableMetrics

* `FindBillableMetric` - Find billable metric by code
* `CreateBillableMetric` - Create a new billable metric
* `DestroyBillableMetric` - Delete a billable metric
* `FindAllBillableMetricGroups` - Find Billable metric groups
* `FindAllBillableMetrics` - Find Billable metrics
* `UpdateBillableMetric` - Update an existing billable metric

### Coupons

* `FindCoupon` - Find coupon by code
* `ApplyCoupon` - Apply a coupon to a customer
* `CreateCoupon` - Create a new coupon
* `DestroyCoupon` - Delete a coupon
* `FindAllAppliedCoupons` - Find Applied Coupons
* `FindAllCoupons` - Find Coupons
* `UpdateCoupon` - Update an existing coupon

### CreditNotes

* `FindCreditNote` - Find credit note
* `CreateCreditNote` - Create a new Credit note
* `DownloadCreditNote` - Download an existing credit note
* `FindAllCreditNotes` - Find Credit notes
* `UpdateCreditNote` - Update an existing credit note
* `VoidCreditNote` - Void existing credit note

### Customers

* `DeleteCustomer` - Delete a customer
* `FindCustomer` - Find customer by external ID
* `FindCustomerCurrentUsage` - Find customer current usage
* `CreateCustomer` - Create a customer
* `FindAllCustomers` - Find customers

### Events

* `FindEvent` - Find event by transaction ID
* `CreateBatchEvents` - Create batch events
* `CreateEvent` - Create a new event

### Invoices

* `FindInvoice` - Find invoice by ID
* `DownloadInvoice` - Download an existing invoice
* `FinalizeInvoice` - Finalize a draft invoice
* `FindAllInvoices` - Find all invoices
* `RefreshInvoice` - Refresh a draft invoice
* `RetryPayment` - Retry invoice payment
* `UpdateInvoice` - Update an existing invoice status

### Organizations

* `UpdateOrganization` - Update an existing Organization

### Plans

* `FindPlan` - Fin plan by code
* `CreatePlan` - Create a new plan
* `DestroyPlan` - Delete a plan
* `FindAllPlans` - Find plans
* `UpdatePlan` - Update an existing plan

### Subscriptions

* `CreateSubscription` - Assign a plan to a customer
* `DestroySubscription` - Terminate a subscription
* `FindAllSubscriptions` - Find subscriptions
* `UpdateSubscription` - Update an existing subscription

### Wallets

* `FindWallet` - Find wallet
* `CreateWallet` - Create a new wallet
* `CreateWalletTransaction` - Create a new wallet transaction
* `DestroyWallet` - Delete a wallet
* `FindAllWalletTransactions` - Find wallet transactions
* `FindAllWallets` - Find wallets
* `UpdateWallet` - Update an existing wallet

### Webhooks

* `FetchPublicKey` - Fetch webhook public key
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
