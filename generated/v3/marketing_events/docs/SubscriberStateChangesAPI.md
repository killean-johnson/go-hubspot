# \SubscriberStateChangesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertUpsertByContactEmail**](SubscriberStateChangesAPI.md#PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertUpsertByContactEmail) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/email-upsert | Record a subscriber state by contact email
[**PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertUpsertByContactId**](SubscriberStateChangesAPI.md#PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertUpsertByContactId) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/upsert | Record a subscriber state by contact ID



## PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertUpsertByContactEmail

> Error PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertUpsertByContactEmail(ctx, externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()

Record a subscriber state by contact email



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
	subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: 'register', 'attend' or 'cancel'.
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application
	batchInputMarketingEventEmailSubscriber := *openapiclient.NewBatchInputMarketingEventEmailSubscriber([]openapiclient.MarketingEventEmailSubscriber{*openapiclient.NewMarketingEventEmailSubscriber("Email_example", int64(123))}) // BatchInputMarketingEventEmailSubscriber | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertUpsertByContactEmail(context.Background(), externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertUpsertByContactEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertUpsertByContactEmail`: Error
	fmt.Fprintf(os.Stdout, "Response from `SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertUpsertByContactEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertUpsertByContactEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application | 
 **batchInputMarketingEventEmailSubscriber** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertUpsertByContactId

> Error PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertUpsertByContactId(ctx, externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()

Record a subscriber state by contact ID



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
	subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: 'register', 'attend' or 'cancel'.
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application
	batchInputMarketingEventSubscriber := *openapiclient.NewBatchInputMarketingEventSubscriber([]openapiclient.MarketingEventSubscriber{*openapiclient.NewMarketingEventSubscriber(int64(123))}) // BatchInputMarketingEventSubscriber | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertUpsertByContactId(context.Background(), externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertUpsertByContactId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertUpsertByContactId`: Error
	fmt.Fprintf(os.Stdout, "Response from `SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertUpsertByContactId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertUpsertByContactIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application | 
 **batchInputMarketingEventSubscriber** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

