# \ChangePropertyAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3MarketingEventsEventsExternalEventIdCancelCancel**](ChangePropertyAPI.md#PostMarketingV3MarketingEventsEventsExternalEventIdCancelCancel) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/cancel | Mark a marketing event as cancelled
[**PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete**](ChangePropertyAPI.md#PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/complete | Mark a marketing event as completed



## PostMarketingV3MarketingEventsEventsExternalEventIdCancelCancel

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEventsExternalEventIdCancelCancel(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

Mark a marketing event as cancelled



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
	externalEventId := "externalEventId_example" // string | The id of the marketing event in the external event application
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangePropertyAPI.PostMarketingV3MarketingEventsEventsExternalEventIdCancelCancel(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangePropertyAPI.PostMarketingV3MarketingEventsEventsExternalEventIdCancelCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsExternalEventIdCancelCancel`: MarketingEventDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `ChangePropertyAPI.PostMarketingV3MarketingEventsEventsExternalEventIdCancelCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdCancelCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams).Execute()

Mark a marketing event as completed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	externalEventId := "externalEventId_example" // string | The id of the marketing event in the external event application.
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application.
	marketingEventCompleteRequestParams := *openapiclient.NewMarketingEventCompleteRequestParams(time.Now(), time.Now()) // MarketingEventCompleteRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangePropertyAPI.PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangePropertyAPI.PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete`: MarketingEventDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `ChangePropertyAPI.PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdCompleteCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application. | 
 **marketingEventCompleteRequestParams** | [**MarketingEventCompleteRequestParams**](MarketingEventCompleteRequestParams.md) |  | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

