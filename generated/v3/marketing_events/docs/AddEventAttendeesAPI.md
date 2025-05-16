# \AddEventAttendeesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateRecordByContactIds**](AddEventAttendeesAPI.md#PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateRecordByContactIds) | **Post** /marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/create | Record Participants by ContactId with Marketing Event External Ids
[**PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateRecordByContactEmails**](AddEventAttendeesAPI.md#PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateRecordByContactEmails) | **Post** /marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/email-create | Record Participants by Email with Marketing Event External Ids
[**PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateCreateRecordByContactId**](AddEventAttendeesAPI.md#PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateCreateRecordByContactId) | **Post** /marketing/v3/marketing-events/{objectId}/attendance/{subscriberState}/create | Record Participants by ContactId with Marketing Event Object Id
[**PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateEmailCreateRecordByEmail**](AddEventAttendeesAPI.md#PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateEmailCreateRecordByEmail) | **Post** /marketing/v3/marketing-events/{objectId}/attendance/{subscriberState}/email-create | Record Participants by Email with Marketing Event Object Id



## PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateRecordByContactIds

> BatchResponseSubscriberVidResponse PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateRecordByContactIds(ctx, externalEventId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).ExternalAccountId(externalAccountId).Execute()

Record Participants by ContactId with Marketing Event External Ids



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
	batchInputMarketingEventSubscriber := *openapiclient.NewBatchInputMarketingEventSubscriber([]openapiclient.MarketingEventSubscriber{*openapiclient.NewMarketingEventSubscriber(int64(123))}) // BatchInputMarketingEventSubscriber | 
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddEventAttendeesAPI.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateRecordByContactIds(context.Background(), externalEventId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).ExternalAccountId(externalAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddEventAttendeesAPI.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateRecordByContactIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateRecordByContactIds`: BatchResponseSubscriberVidResponse
	fmt.Fprintf(os.Stdout, "Response from `AddEventAttendeesAPI.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateRecordByContactIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateRecordByContactIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventSubscriber** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md) |  | 
 **externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application | 

### Return type

[**BatchResponseSubscriberVidResponse**](BatchResponseSubscriberVidResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateRecordByContactEmails

> BatchResponseSubscriberEmailResponse PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateRecordByContactEmails(ctx, externalEventId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).ExternalAccountId(externalAccountId).Execute()

Record Participants by Email with Marketing Event External Ids



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
	batchInputMarketingEventEmailSubscriber := *openapiclient.NewBatchInputMarketingEventEmailSubscriber([]openapiclient.MarketingEventEmailSubscriber{*openapiclient.NewMarketingEventEmailSubscriber("Email_example", int64(123))}) // BatchInputMarketingEventEmailSubscriber | 
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddEventAttendeesAPI.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateRecordByContactEmails(context.Background(), externalEventId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).ExternalAccountId(externalAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddEventAttendeesAPI.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateRecordByContactEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateRecordByContactEmails`: BatchResponseSubscriberEmailResponse
	fmt.Fprintf(os.Stdout, "Response from `AddEventAttendeesAPI.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateRecordByContactEmails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateRecordByContactEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventEmailSubscriber** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md) |  | 
 **externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application | 

### Return type

[**BatchResponseSubscriberEmailResponse**](BatchResponseSubscriberEmailResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateCreateRecordByContactId

> BatchResponseSubscriberVidResponse PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateCreateRecordByContactId(ctx, objectId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()

Record Participants by ContactId with Marketing Event Object Id



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
	objectId := "objectId_example" // string | The internal id of the marketing event in HubSpot
	subscriberState := "subscriberState_example" // string | The attendance state value. It may be 'register', 'attend' or 'cancel'
	batchInputMarketingEventSubscriber := *openapiclient.NewBatchInputMarketingEventSubscriber([]openapiclient.MarketingEventSubscriber{*openapiclient.NewMarketingEventSubscriber(int64(123))}) // BatchInputMarketingEventSubscriber | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddEventAttendeesAPI.PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateCreateRecordByContactId(context.Background(), objectId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddEventAttendeesAPI.PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateCreateRecordByContactId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateCreateRecordByContactId`: BatchResponseSubscriberVidResponse
	fmt.Fprintf(os.Stdout, "Response from `AddEventAttendeesAPI.PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateCreateRecordByContactId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The internal id of the marketing event in HubSpot | 
**subscriberState** | **string** | The attendance state value. It may be &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateCreateRecordByContactIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventSubscriber** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md) |  | 

### Return type

[**BatchResponseSubscriberVidResponse**](BatchResponseSubscriberVidResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateEmailCreateRecordByEmail

> BatchResponseSubscriberEmailResponse PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateEmailCreateRecordByEmail(ctx, objectId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()

Record Participants by Email with Marketing Event Object Id



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
	objectId := "objectId_example" // string | The internal ID of the marketing event in HubSpot
	subscriberState := "subscriberState_example" // string | The attendance state value. It may be 'register', 'attend' or 'cancel'
	batchInputMarketingEventEmailSubscriber := *openapiclient.NewBatchInputMarketingEventEmailSubscriber([]openapiclient.MarketingEventEmailSubscriber{*openapiclient.NewMarketingEventEmailSubscriber("Email_example", int64(123))}) // BatchInputMarketingEventEmailSubscriber | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddEventAttendeesAPI.PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateEmailCreateRecordByEmail(context.Background(), objectId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddEventAttendeesAPI.PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateEmailCreateRecordByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateEmailCreateRecordByEmail`: BatchResponseSubscriberEmailResponse
	fmt.Fprintf(os.Stdout, "Response from `AddEventAttendeesAPI.PostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateEmailCreateRecordByEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The internal ID of the marketing event in HubSpot | 
**subscriberState** | **string** | The attendance state value. It may be &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsObjectIdAttendanceSubscriberStateEmailCreateRecordByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventEmailSubscriber** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md) |  | 

### Return type

[**BatchResponseSubscriberEmailResponse**](BatchResponseSubscriberEmailResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

