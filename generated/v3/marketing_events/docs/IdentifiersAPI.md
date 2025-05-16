# \IdentifiersAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3MarketingEventsEventsSearchDoSearch**](IdentifiersAPI.md#GetMarketingV3MarketingEventsEventsSearchDoSearch) | **Get** /marketing/v3/marketing-events/events/search | Find App-Specific Marketing Events by External Event Id
[**GetMarketingV3MarketingEventsExternalEventIdIdentifiersSearchPortalEvents**](IdentifiersAPI.md#GetMarketingV3MarketingEventsExternalEventIdIdentifiersSearchPortalEvents) | **Get** /marketing/v3/marketing-events/{externalEventId}/identifiers | Find Marketing Events by External Event Id



## GetMarketingV3MarketingEventsEventsSearchDoSearch

> CollectionResponseSearchPublicResponseWrapperNoPaging GetMarketingV3MarketingEventsEventsSearchDoSearch(ctx).Q(q).Execute()

Find App-Specific Marketing Events by External Event Id



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
	q := "q_example" // string | The id of the marketing event in the external event application (externalEventId)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentifiersAPI.GetMarketingV3MarketingEventsEventsSearchDoSearch(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentifiersAPI.GetMarketingV3MarketingEventsEventsSearchDoSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsEventsSearchDoSearch`: CollectionResponseSearchPublicResponseWrapperNoPaging
	fmt.Fprintf(os.Stdout, "Response from `IdentifiersAPI.GetMarketingV3MarketingEventsEventsSearchDoSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsEventsSearchDoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The id of the marketing event in the external event application (externalEventId) | 

### Return type

[**CollectionResponseSearchPublicResponseWrapperNoPaging**](CollectionResponseSearchPublicResponseWrapperNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsExternalEventIdIdentifiersSearchPortalEvents

> CollectionResponseWithTotalMarketingEventIdentifiersResponseNoPaging GetMarketingV3MarketingEventsExternalEventIdIdentifiersSearchPortalEvents(ctx, externalEventId).Execute()

Find Marketing Events by External Event Id



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
	externalEventId := "externalEventId_example" // string | The id of the marketing event in the external event application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentifiersAPI.GetMarketingV3MarketingEventsExternalEventIdIdentifiersSearchPortalEvents(context.Background(), externalEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentifiersAPI.GetMarketingV3MarketingEventsExternalEventIdIdentifiersSearchPortalEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsExternalEventIdIdentifiersSearchPortalEvents`: CollectionResponseWithTotalMarketingEventIdentifiersResponseNoPaging
	fmt.Fprintf(os.Stdout, "Response from `IdentifiersAPI.GetMarketingV3MarketingEventsExternalEventIdIdentifiersSearchPortalEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsExternalEventIdIdentifiersSearchPortalEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponseWithTotalMarketingEventIdentifiersResponseNoPaging**](CollectionResponseWithTotalMarketingEventIdentifiersResponseNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

