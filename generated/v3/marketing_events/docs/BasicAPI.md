# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive**](BasicAPI.md#DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive) | **Delete** /marketing/v3/marketing-events/events/{externalEventId} | Delete Marketing Event by External Ids
[**DeleteMarketingV3MarketingEventsObjectIdArchiveByObjectId**](BasicAPI.md#DeleteMarketingV3MarketingEventsObjectIdArchiveByObjectId) | **Delete** /marketing/v3/marketing-events/{objectId} | Delete Marketing Event by objectId
[**GetMarketingV3MarketingEventsEventsExternalEventIdGetDetails**](BasicAPI.md#GetMarketingV3MarketingEventsEventsExternalEventIdGetDetails) | **Get** /marketing/v3/marketing-events/events/{externalEventId} | Get Marketing Event by External IDs
[**GetMarketingV3MarketingEventsGetAll**](BasicAPI.md#GetMarketingV3MarketingEventsGetAll) | **Get** /marketing/v3/marketing-events/ | Get all marketing event
[**GetMarketingV3MarketingEventsObjectIdGetByObjectId**](BasicAPI.md#GetMarketingV3MarketingEventsObjectIdGetByObjectId) | **Get** /marketing/v3/marketing-events/{objectId} | Get Marketing Event by objectId
[**PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate**](BasicAPI.md#PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate) | **Patch** /marketing/v3/marketing-events/events/{externalEventId} | Update Marketing Event by External IDs
[**PatchMarketingV3MarketingEventsObjectIdUpdateByObjectId**](BasicAPI.md#PatchMarketingV3MarketingEventsObjectIdUpdateByObjectId) | **Patch** /marketing/v3/marketing-events/{objectId} | Update Marketing Event by objectId
[**PostMarketingV3MarketingEventsEventsCreate**](BasicAPI.md#PostMarketingV3MarketingEventsEventsCreate) | **Post** /marketing/v3/marketing-events/events | Create a marketing event
[**PutMarketingV3MarketingEventsEventsExternalEventIdUpsert**](BasicAPI.md#PutMarketingV3MarketingEventsEventsExternalEventIdUpsert) | **Put** /marketing/v3/marketing-events/events/{externalEventId} | Create or update a marketing event



## DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive

> DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

Delete Marketing Event by External Ids



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
	r, err := apiClient.BasicAPI.DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3MarketingEventsEventsExternalEventIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMarketingV3MarketingEventsObjectIdArchiveByObjectId

> DeleteMarketingV3MarketingEventsObjectIdArchiveByObjectId(ctx, objectId).Execute()

Delete Marketing Event by objectId



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteMarketingV3MarketingEventsObjectIdArchiveByObjectId(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteMarketingV3MarketingEventsObjectIdArchiveByObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The internal ID of the marketing event in HubSpot | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3MarketingEventsObjectIdArchiveByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsEventsExternalEventIdGetDetails

> MarketingEventPublicReadResponse GetMarketingV3MarketingEventsEventsExternalEventIdGetDetails(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

Get Marketing Event by External IDs



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
	resp, r, err := apiClient.BasicAPI.GetMarketingV3MarketingEventsEventsExternalEventIdGetDetails(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetMarketingV3MarketingEventsEventsExternalEventIdGetDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsEventsExternalEventIdGetDetails`: MarketingEventPublicReadResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetMarketingV3MarketingEventsEventsExternalEventIdGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsEventsExternalEventIdGetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application | 

### Return type

[**MarketingEventPublicReadResponse**](MarketingEventPublicReadResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsGetAll

> CollectionResponseMarketingEventPublicReadResponseV2ForwardPaging GetMarketingV3MarketingEventsGetAll(ctx).After(after).Limit(limit).Execute()

Get all marketing event



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
	after := "after_example" // string | The cursor indicating the position of the last retrieved item. (optional)
	limit := int32(56) // int32 | The limit for response size. The default value is 10, the max number is 100 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetMarketingV3MarketingEventsGetAll(context.Background()).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetMarketingV3MarketingEventsGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsGetAll`: CollectionResponseMarketingEventPublicReadResponseV2ForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetMarketingV3MarketingEventsGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor indicating the position of the last retrieved item. | 
 **limit** | **int32** | The limit for response size. The default value is 10, the max number is 100 | [default to 10]

### Return type

[**CollectionResponseMarketingEventPublicReadResponseV2ForwardPaging**](CollectionResponseMarketingEventPublicReadResponseV2ForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsObjectIdGetByObjectId

> MarketingEventPublicReadResponseV2 GetMarketingV3MarketingEventsObjectIdGetByObjectId(ctx, objectId).Execute()

Get Marketing Event by objectId



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetMarketingV3MarketingEventsObjectIdGetByObjectId(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetMarketingV3MarketingEventsObjectIdGetByObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsObjectIdGetByObjectId`: MarketingEventPublicReadResponseV2
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetMarketingV3MarketingEventsObjectIdGetByObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The internal ID of the marketing event in HubSpot | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsObjectIdGetByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketingEventPublicReadResponseV2**](MarketingEventPublicReadResponseV2.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate

> MarketingEventPublicDefaultResponse PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()

Update Marketing Event by External IDs



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
	marketingEventUpdateRequestParams := *openapiclient.NewMarketingEventUpdateRequestParams() // MarketingEventUpdateRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate`: MarketingEventPublicDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3MarketingEventsEventsExternalEventIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application | 
 **marketingEventUpdateRequestParams** | [**MarketingEventUpdateRequestParams**](MarketingEventUpdateRequestParams.md) |  | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingV3MarketingEventsObjectIdUpdateByObjectId

> MarketingEventPublicDefaultResponseV2 PatchMarketingV3MarketingEventsObjectIdUpdateByObjectId(ctx, objectId).MarketingEventPublicUpdateRequestV2(marketingEventPublicUpdateRequestV2).Execute()

Update Marketing Event by objectId



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
	marketingEventPublicUpdateRequestV2 := *openapiclient.NewMarketingEventPublicUpdateRequestV2([]openapiclient.PropertyValue{*openapiclient.NewPropertyValue("SourceId_example", false, "SourceLabel_example", "Source_example", "SourceMetadata_example", "DataSensitivity_example", []int64{int64(123)}, "Unit_example", "RequestId_example", false, "Name_example", "Value_example", int64(123), int64(123))}) // MarketingEventPublicUpdateRequestV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchMarketingV3MarketingEventsObjectIdUpdateByObjectId(context.Background(), objectId).MarketingEventPublicUpdateRequestV2(marketingEventPublicUpdateRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchMarketingV3MarketingEventsObjectIdUpdateByObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingV3MarketingEventsObjectIdUpdateByObjectId`: MarketingEventPublicDefaultResponseV2
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchMarketingV3MarketingEventsObjectIdUpdateByObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The internal ID of the marketing event in HubSpot | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3MarketingEventsObjectIdUpdateByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketingEventPublicUpdateRequestV2** | [**MarketingEventPublicUpdateRequestV2**](MarketingEventPublicUpdateRequestV2.md) |  | 

### Return type

[**MarketingEventPublicDefaultResponseV2**](MarketingEventPublicDefaultResponseV2.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsCreate

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEventsCreate(ctx).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()

Create a marketing event



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
	marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("ExternalAccountId_example", "EventOrganizer_example", "ExternalEventId_example", "EventName_example") // MarketingEventCreateRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostMarketingV3MarketingEventsEventsCreate(context.Background()).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostMarketingV3MarketingEventsEventsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsCreate`: MarketingEventDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostMarketingV3MarketingEventsEventsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) |  | 

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


## PutMarketingV3MarketingEventsEventsExternalEventIdUpsert

> MarketingEventPublicDefaultResponse PutMarketingV3MarketingEventsEventsExternalEventIdUpsert(ctx, externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()

Create or update a marketing event



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
	marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("ExternalAccountId_example", "EventOrganizer_example", "ExternalEventId_example", "EventName_example") // MarketingEventCreateRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PutMarketingV3MarketingEventsEventsExternalEventIdUpsert(context.Background(), externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutMarketingV3MarketingEventsEventsExternalEventIdUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingV3MarketingEventsEventsExternalEventIdUpsert`: MarketingEventPublicDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PutMarketingV3MarketingEventsEventsExternalEventIdUpsert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event in the external event application | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingV3MarketingEventsEventsExternalEventIdUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) |  | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

