# \BudgetAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3CampaignsCampaignGuidBudgetTotals**](BudgetAPI.md#GetMarketingV3CampaignsCampaignGuidBudgetTotals) | **Get** /marketing/v3/campaigns/{campaignGuid}/budget/totals | Read budget



## GetMarketingV3CampaignsCampaignGuidBudgetTotals

> PublicBudgetTotals GetMarketingV3CampaignsCampaignGuidBudgetTotals(ctx, campaignGuid).Execute()

Read budget



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignGuid := "campaignGuid_example" // string | Unique identifier for the campaign, formatted as a UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetAPI.GetMarketingV3CampaignsCampaignGuidBudgetTotals(context.Background(), campaignGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetAPI.GetMarketingV3CampaignsCampaignGuidBudgetTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3CampaignsCampaignGuidBudgetTotals`: PublicBudgetTotals
	fmt.Fprintf(os.Stdout, "Response from `BudgetAPI.GetMarketingV3CampaignsCampaignGuidBudgetTotals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3CampaignsCampaignGuidBudgetTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicBudgetTotals**](PublicBudgetTotals.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

