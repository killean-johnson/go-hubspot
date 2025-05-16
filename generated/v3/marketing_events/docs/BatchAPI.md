# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3MarketingEventsBatchArchiveArchiveByObjectId**](BatchAPI.md#PostMarketingV3MarketingEventsBatchArchiveArchiveByObjectId) | **Post** /marketing/v3/marketing-events/batch/archive | Delete Multiple Marketing Events by ObjectId
[**PostMarketingV3MarketingEventsBatchUpdateUpdateByObjectId**](BatchAPI.md#PostMarketingV3MarketingEventsBatchUpdateUpdateByObjectId) | **Post** /marketing/v3/marketing-events/batch/update | Update Multiple Marketing Events by ObjectId
[**PostMarketingV3MarketingEventsEventsDeleteArchive**](BatchAPI.md#PostMarketingV3MarketingEventsEventsDeleteArchive) | **Post** /marketing/v3/marketing-events/events/delete | Delete Multiple Marketing Events by External Ids
[**PostMarketingV3MarketingEventsEventsUpsertUpsert**](BatchAPI.md#PostMarketingV3MarketingEventsEventsUpsertUpsert) | **Post** /marketing/v3/marketing-events/events/upsert | Create or Update Multiple Marketing Events



## PostMarketingV3MarketingEventsBatchArchiveArchiveByObjectId

> PostMarketingV3MarketingEventsBatchArchiveArchiveByObjectId(ctx).BatchInputMarketingEventPublicObjectIdDeleteRequest(batchInputMarketingEventPublicObjectIdDeleteRequest).Execute()

Delete Multiple Marketing Events by ObjectId



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
	batchInputMarketingEventPublicObjectIdDeleteRequest := *openapiclient.NewBatchInputMarketingEventPublicObjectIdDeleteRequest([]openapiclient.MarketingEventPublicObjectIdDeleteRequest{*openapiclient.NewMarketingEventPublicObjectIdDeleteRequest("ObjectId_example")}) // BatchInputMarketingEventPublicObjectIdDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.PostMarketingV3MarketingEventsBatchArchiveArchiveByObjectId(context.Background()).BatchInputMarketingEventPublicObjectIdDeleteRequest(batchInputMarketingEventPublicObjectIdDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3MarketingEventsBatchArchiveArchiveByObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsBatchArchiveArchiveByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventPublicObjectIdDeleteRequest** | [**BatchInputMarketingEventPublicObjectIdDeleteRequest**](BatchInputMarketingEventPublicObjectIdDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsBatchUpdateUpdateByObjectId

> BatchResponseMarketingEventPublicDefaultResponseV2 PostMarketingV3MarketingEventsBatchUpdateUpdateByObjectId(ctx).BatchInputMarketingEventPublicUpdateRequestFullV2(batchInputMarketingEventPublicUpdateRequestFullV2).Execute()

Update Multiple Marketing Events by ObjectId



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
	batchInputMarketingEventPublicUpdateRequestFullV2 := *openapiclient.NewBatchInputMarketingEventPublicUpdateRequestFullV2([]openapiclient.MarketingEventPublicUpdateRequestFullV2{*openapiclient.NewMarketingEventPublicUpdateRequestFullV2([]openapiclient.PropertyValue{*openapiclient.NewPropertyValue("SourceId_example", false, "SourceLabel_example", "Source_example", "SourceMetadata_example", "DataSensitivity_example", []int64{int64(123)}, "Unit_example", "RequestId_example", false, "Name_example", "Value_example", int64(123), int64(123))}, "ObjectId_example")}) // BatchInputMarketingEventPublicUpdateRequestFullV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostMarketingV3MarketingEventsBatchUpdateUpdateByObjectId(context.Background()).BatchInputMarketingEventPublicUpdateRequestFullV2(batchInputMarketingEventPublicUpdateRequestFullV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3MarketingEventsBatchUpdateUpdateByObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsBatchUpdateUpdateByObjectId`: BatchResponseMarketingEventPublicDefaultResponseV2
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostMarketingV3MarketingEventsBatchUpdateUpdateByObjectId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsBatchUpdateUpdateByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventPublicUpdateRequestFullV2** | [**BatchInputMarketingEventPublicUpdateRequestFullV2**](BatchInputMarketingEventPublicUpdateRequestFullV2.md) |  | 

### Return type

[**BatchResponseMarketingEventPublicDefaultResponseV2**](BatchResponseMarketingEventPublicDefaultResponseV2.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsDeleteArchive

> Error PostMarketingV3MarketingEventsEventsDeleteArchive(ctx).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()

Delete Multiple Marketing Events by External Ids



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
	batchInputMarketingEventExternalUniqueIdentifier := *openapiclient.NewBatchInputMarketingEventExternalUniqueIdentifier([]openapiclient.MarketingEventExternalUniqueIdentifier{*openapiclient.NewMarketingEventExternalUniqueIdentifier("ExternalAccountId_example", "ExternalEventId_example", int32(123))}) // BatchInputMarketingEventExternalUniqueIdentifier | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostMarketingV3MarketingEventsEventsDeleteArchive(context.Background()).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3MarketingEventsEventsDeleteArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsDeleteArchive`: Error
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostMarketingV3MarketingEventsEventsDeleteArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventExternalUniqueIdentifier** | [**BatchInputMarketingEventExternalUniqueIdentifier**](BatchInputMarketingEventExternalUniqueIdentifier.md) |  | 

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


## PostMarketingV3MarketingEventsEventsUpsertUpsert

> BatchResponseMarketingEventPublicDefaultResponse PostMarketingV3MarketingEventsEventsUpsertUpsert(ctx).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()

Create or Update Multiple Marketing Events



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
	batchInputMarketingEventCreateRequestParams := *openapiclient.NewBatchInputMarketingEventCreateRequestParams([]openapiclient.MarketingEventCreateRequestParams{*openapiclient.NewMarketingEventCreateRequestParams("ExternalAccountId_example", "EventOrganizer_example", "ExternalEventId_example", "EventName_example")}) // BatchInputMarketingEventCreateRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostMarketingV3MarketingEventsEventsUpsertUpsert(context.Background()).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3MarketingEventsEventsUpsertUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsUpsertUpsert`: BatchResponseMarketingEventPublicDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostMarketingV3MarketingEventsEventsUpsertUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsUpsertUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventCreateRequestParams** | [**BatchInputMarketingEventCreateRequestParams**](BatchInputMarketingEventCreateRequestParams.md) |  | 

### Return type

[**BatchResponseMarketingEventPublicDefaultResponse**](BatchResponseMarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

