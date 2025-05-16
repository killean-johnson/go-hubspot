# \ReportsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3CampaignsCampaignGuidReportsContactsContactType**](ReportsAPI.md#GetMarketingV3CampaignsCampaignGuidReportsContactsContactType) | **Get** /marketing/v3/campaigns/{campaignGuid}/reports/contacts/{contactType} | Fetch contact IDs
[**GetMarketingV3CampaignsCampaignGuidReportsMetrics**](ReportsAPI.md#GetMarketingV3CampaignsCampaignGuidReportsMetrics) | **Get** /marketing/v3/campaigns/{campaignGuid}/reports/metrics | Get Campaign Metrics 
[**GetMarketingV3CampaignsCampaignGuidReportsRevenue**](ReportsAPI.md#GetMarketingV3CampaignsCampaignGuidReportsRevenue) | **Get** /marketing/v3/campaigns/{campaignGuid}/reports/revenue | Fetch revenue



## GetMarketingV3CampaignsCampaignGuidReportsContactsContactType

> CollectionResponseContactReferenceForwardPaging GetMarketingV3CampaignsCampaignGuidReportsContactsContactType(ctx, campaignGuid, contactType).StartDate(startDate).EndDate(endDate).Limit(limit).After(after).Execute()

Fetch contact IDs



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
	contactType := "contactType_example" // string | The type of metric to filter the influenced contacts. Allowed values: contactFirstTouch, contactLastTouch, influencedContacts
	startDate := "startDate_example" // string | The start date for the report data, formatted as YYYY-MM-DD. Default value: 2006-01-01 (optional)
	endDate := "endDate_example" // string | End date for the report data, formatted as YYYY-MM-DD. Default value: Current date (optional)
	limit := int32(56) // int32 | Limit for the number of contacts to fetch Default: 100 (optional)
	after := "after_example" // string | A cursor for pagination. If provided, the results will start after the given cursor. Example: NTI1Cg%3D%3D (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetMarketingV3CampaignsCampaignGuidReportsContactsContactType(context.Background(), campaignGuid, contactType).StartDate(startDate).EndDate(endDate).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetMarketingV3CampaignsCampaignGuidReportsContactsContactType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3CampaignsCampaignGuidReportsContactsContactType`: CollectionResponseContactReferenceForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetMarketingV3CampaignsCampaignGuidReportsContactsContactType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID. | 
**contactType** | **string** | The type of metric to filter the influenced contacts. Allowed values: contactFirstTouch, contactLastTouch, influencedContacts | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3CampaignsCampaignGuidReportsContactsContactTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **string** | The start date for the report data, formatted as YYYY-MM-DD. Default value: 2006-01-01 | 
 **endDate** | **string** | End date for the report data, formatted as YYYY-MM-DD. Default value: Current date | 
 **limit** | **int32** | Limit for the number of contacts to fetch Default: 100 | 
 **after** | **string** | A cursor for pagination. If provided, the results will start after the given cursor. Example: NTI1Cg%3D%3D | 

### Return type

[**CollectionResponseContactReferenceForwardPaging**](CollectionResponseContactReferenceForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3CampaignsCampaignGuidReportsMetrics

> MetricsCounters GetMarketingV3CampaignsCampaignGuidReportsMetrics(ctx, campaignGuid).StartDate(startDate).EndDate(endDate).Execute()

Get Campaign Metrics 



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
	startDate := "startDate_example" // string | The start date for the report data, formatted as YYYY-MM-DD. Default value: 2006-01-01 (optional)
	endDate := "endDate_example" // string | End date for the report data, formatted as YYYY-MM-DD. Default value: Current date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetMarketingV3CampaignsCampaignGuidReportsMetrics(context.Background(), campaignGuid).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetMarketingV3CampaignsCampaignGuidReportsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3CampaignsCampaignGuidReportsMetrics`: MetricsCounters
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetMarketingV3CampaignsCampaignGuidReportsMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3CampaignsCampaignGuidReportsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | The start date for the report data, formatted as YYYY-MM-DD. Default value: 2006-01-01 | 
 **endDate** | **string** | End date for the report data, formatted as YYYY-MM-DD. Default value: Current date | 

### Return type

[**MetricsCounters**](MetricsCounters.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3CampaignsCampaignGuidReportsRevenue

> RevenueAttributionAggregate GetMarketingV3CampaignsCampaignGuidReportsRevenue(ctx, campaignGuid).AttributionModel(attributionModel).StartDate(startDate).EndDate(endDate).Execute()

Fetch revenue



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
	attributionModel := "attributionModel_example" // string | Allowed values: LINEAR, FIRST_INTERACTION, LAST_INTERACTION, FULL_PATH, U_SHAPED, W_SHAPED, TIME_DECAY, J_SHAPED, INVERSE_J_SHAPED Default value: LINEAR (optional)
	startDate := "startDate_example" // string | The start date for the report data, formatted as YYYY-MM-DD. Default value: 2006-01-01 (optional)
	endDate := "endDate_example" // string | End date for the report data, formatted as YYYY-MM-DD. Default value: Current date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetMarketingV3CampaignsCampaignGuidReportsRevenue(context.Background(), campaignGuid).AttributionModel(attributionModel).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetMarketingV3CampaignsCampaignGuidReportsRevenue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3CampaignsCampaignGuidReportsRevenue`: RevenueAttributionAggregate
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetMarketingV3CampaignsCampaignGuidReportsRevenue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3CampaignsCampaignGuidReportsRevenueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributionModel** | **string** | Allowed values: LINEAR, FIRST_INTERACTION, LAST_INTERACTION, FULL_PATH, U_SHAPED, W_SHAPED, TIME_DECAY, J_SHAPED, INVERSE_J_SHAPED Default value: LINEAR | 
 **startDate** | **string** | The start date for the report data, formatted as YYYY-MM-DD. Default value: 2006-01-01 | 
 **endDate** | **string** | End date for the report data, formatted as YYYY-MM-DD. Default value: Current date | 

### Return type

[**RevenueAttributionAggregate**](RevenueAttributionAggregate.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

