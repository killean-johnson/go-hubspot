# \RetrieveParticipantStateAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3MarketingEventsParticipationsContactsContactIdentifierBreakdownGetParticipationsBreakdownByContactId**](RetrieveParticipantStateAPI.md#GetMarketingV3MarketingEventsParticipationsContactsContactIdentifierBreakdownGetParticipationsBreakdownByContactId) | **Get** /marketing/v3/marketing-events/participations/contacts/{contactIdentifier}/breakdown | Read participations breakdown by Contact identifier
[**GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdBreakdownGetParticipationsBreakdownByExternalEventId**](RetrieveParticipantStateAPI.md#GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdBreakdownGetParticipationsBreakdownByExternalEventId) | **Get** /marketing/v3/marketing-events/participations/{externalAccountId}/{externalEventId}/breakdown | Read participations breakdown by Marketing Event external identifier
[**GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdGetParticipationsCountersByEventExternalId**](RetrieveParticipantStateAPI.md#GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdGetParticipationsCountersByEventExternalId) | **Get** /marketing/v3/marketing-events/participations/{externalAccountId}/{externalEventId} | Read participations counters by Marketing Event external identifier
[**GetMarketingV3MarketingEventsParticipationsMarketingEventIdBreakdownGetParticipationsBreakdownByMarketingEventId**](RetrieveParticipantStateAPI.md#GetMarketingV3MarketingEventsParticipationsMarketingEventIdBreakdownGetParticipationsBreakdownByMarketingEventId) | **Get** /marketing/v3/marketing-events/participations/{marketingEventId}/breakdown | Read participations breakdown by Marketing Event internal identifier
[**GetMarketingV3MarketingEventsParticipationsMarketingEventIdGetParticipationsCountersByMarketingEventId**](RetrieveParticipantStateAPI.md#GetMarketingV3MarketingEventsParticipationsMarketingEventIdGetParticipationsCountersByMarketingEventId) | **Get** /marketing/v3/marketing-events/participations/{marketingEventId} | Read participations counters by Marketing Event internal identifier



## GetMarketingV3MarketingEventsParticipationsContactsContactIdentifierBreakdownGetParticipationsBreakdownByContactId

> CollectionResponseWithTotalParticipationBreakdownForwardPaging GetMarketingV3MarketingEventsParticipationsContactsContactIdentifierBreakdownGetParticipationsBreakdownByContactId(ctx, contactIdentifier).State(state).Limit(limit).After(after).Execute()

Read participations breakdown by Contact identifier



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
	contactIdentifier := "contactIdentifier_example" // string | The identifier of the Contact. It may be email or internal id.
	state := "state_example" // string | The participation state value. It may be REGISTERED, CANCELLED, ATTENDED, NO_SHOW (optional)
	limit := int32(56) // int32 | The limit for response size. The default value is 10, the max number is 100 (optional) (default to 10)
	after := "after_example" // string | The cursor indicating the position of the last retrieved item. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsContactsContactIdentifierBreakdownGetParticipationsBreakdownByContactId(context.Background(), contactIdentifier).State(state).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsContactsContactIdentifierBreakdownGetParticipationsBreakdownByContactId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsParticipationsContactsContactIdentifierBreakdownGetParticipationsBreakdownByContactId`: CollectionResponseWithTotalParticipationBreakdownForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsContactsContactIdentifierBreakdownGetParticipationsBreakdownByContactId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactIdentifier** | **string** | The identifier of the Contact. It may be email or internal id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsParticipationsContactsContactIdentifierBreakdownGetParticipationsBreakdownByContactIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | The participation state value. It may be REGISTERED, CANCELLED, ATTENDED, NO_SHOW | 
 **limit** | **int32** | The limit for response size. The default value is 10, the max number is 100 | [default to 10]
 **after** | **string** | The cursor indicating the position of the last retrieved item. | 

### Return type

[**CollectionResponseWithTotalParticipationBreakdownForwardPaging**](CollectionResponseWithTotalParticipationBreakdownForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdBreakdownGetParticipationsBreakdownByExternalEventId

> CollectionResponseWithTotalParticipationBreakdownForwardPaging GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdBreakdownGetParticipationsBreakdownByExternalEventId(ctx, externalAccountId, externalEventId).ContactIdentifier(contactIdentifier).State(state).Limit(limit).After(after).Execute()

Read participations breakdown by Marketing Event external identifier



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
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application.
	externalEventId := "externalEventId_example" // string | The id of the marketing event in the external event application.
	contactIdentifier := "contactIdentifier_example" // string | The identifier of the Contact. It may be email or internal id. (optional)
	state := "state_example" // string | The participation state value. It may be REGISTERED, CANCELLED, ATTENDED, NO_SHOW (optional)
	limit := int32(56) // int32 | The limit for response size. The default value is 10, the max number is 100 (optional) (default to 10)
	after := "after_example" // string | The cursor indicating the position of the last retrieved item. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdBreakdownGetParticipationsBreakdownByExternalEventId(context.Background(), externalAccountId, externalEventId).ContactIdentifier(contactIdentifier).State(state).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdBreakdownGetParticipationsBreakdownByExternalEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdBreakdownGetParticipationsBreakdownByExternalEventId`: CollectionResponseWithTotalParticipationBreakdownForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdBreakdownGetParticipationsBreakdownByExternalEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application. | 
**externalEventId** | **string** | The id of the marketing event in the external event application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdBreakdownGetParticipationsBreakdownByExternalEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contactIdentifier** | **string** | The identifier of the Contact. It may be email or internal id. | 
 **state** | **string** | The participation state value. It may be REGISTERED, CANCELLED, ATTENDED, NO_SHOW | 
 **limit** | **int32** | The limit for response size. The default value is 10, the max number is 100 | [default to 10]
 **after** | **string** | The cursor indicating the position of the last retrieved item. | 

### Return type

[**CollectionResponseWithTotalParticipationBreakdownForwardPaging**](CollectionResponseWithTotalParticipationBreakdownForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdGetParticipationsCountersByEventExternalId

> AttendanceCounters GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdGetParticipationsCountersByEventExternalId(ctx, externalAccountId, externalEventId).Execute()

Read participations counters by Marketing Event external identifier



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
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application.
	externalEventId := "externalEventId_example" // string | The id of the marketing event in the external event application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdGetParticipationsCountersByEventExternalId(context.Background(), externalAccountId, externalEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdGetParticipationsCountersByEventExternalId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdGetParticipationsCountersByEventExternalId`: AttendanceCounters
	fmt.Fprintf(os.Stdout, "Response from `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdGetParticipationsCountersByEventExternalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application. | 
**externalEventId** | **string** | The id of the marketing event in the external event application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsParticipationsExternalAccountIdExternalEventIdGetParticipationsCountersByEventExternalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttendanceCounters**](AttendanceCounters.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsParticipationsMarketingEventIdBreakdownGetParticipationsBreakdownByMarketingEventId

> CollectionResponseWithTotalParticipationBreakdownForwardPaging GetMarketingV3MarketingEventsParticipationsMarketingEventIdBreakdownGetParticipationsBreakdownByMarketingEventId(ctx, marketingEventId).ContactIdentifier(contactIdentifier).State(state).Limit(limit).After(after).Execute()

Read participations breakdown by Marketing Event internal identifier



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
	marketingEventId := int64(789) // int64 | The internal id of the marketing event in HubSpot.
	contactIdentifier := "contactIdentifier_example" // string | The identifier of the Contact. It may be email or internal id. (optional)
	state := "state_example" // string | The participation state value. It may be REGISTERED, CANCELLED, ATTENDED, NO_SHOW (optional)
	limit := int32(56) // int32 | The limit for response size. The default value is 10, the max number is 100 (optional) (default to 10)
	after := "after_example" // string | The cursor indicating the position of the last retrieved item. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsMarketingEventIdBreakdownGetParticipationsBreakdownByMarketingEventId(context.Background(), marketingEventId).ContactIdentifier(contactIdentifier).State(state).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsMarketingEventIdBreakdownGetParticipationsBreakdownByMarketingEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsParticipationsMarketingEventIdBreakdownGetParticipationsBreakdownByMarketingEventId`: CollectionResponseWithTotalParticipationBreakdownForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsMarketingEventIdBreakdownGetParticipationsBreakdownByMarketingEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingEventId** | **int64** | The internal id of the marketing event in HubSpot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsParticipationsMarketingEventIdBreakdownGetParticipationsBreakdownByMarketingEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactIdentifier** | **string** | The identifier of the Contact. It may be email or internal id. | 
 **state** | **string** | The participation state value. It may be REGISTERED, CANCELLED, ATTENDED, NO_SHOW | 
 **limit** | **int32** | The limit for response size. The default value is 10, the max number is 100 | [default to 10]
 **after** | **string** | The cursor indicating the position of the last retrieved item. | 

### Return type

[**CollectionResponseWithTotalParticipationBreakdownForwardPaging**](CollectionResponseWithTotalParticipationBreakdownForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsParticipationsMarketingEventIdGetParticipationsCountersByMarketingEventId

> AttendanceCounters GetMarketingV3MarketingEventsParticipationsMarketingEventIdGetParticipationsCountersByMarketingEventId(ctx, marketingEventId).Execute()

Read participations counters by Marketing Event internal identifier



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
	marketingEventId := int64(789) // int64 | The internal id of the marketing event in HubSpot.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsMarketingEventIdGetParticipationsCountersByMarketingEventId(context.Background(), marketingEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsMarketingEventIdGetParticipationsCountersByMarketingEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsParticipationsMarketingEventIdGetParticipationsCountersByMarketingEventId`: AttendanceCounters
	fmt.Fprintf(os.Stdout, "Response from `RetrieveParticipantStateAPI.GetMarketingV3MarketingEventsParticipationsMarketingEventIdGetParticipationsCountersByMarketingEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingEventId** | **int64** | The internal id of the marketing event in HubSpot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsParticipationsMarketingEventIdGetParticipationsCountersByMarketingEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttendanceCounters**](AttendanceCounters.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

