# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsDealsSplitsBatchReadRead**](BatchAPI.md#PostCrmV3ObjectsDealsSplitsBatchReadRead) | **Post** /crm/v3/objects/deals/splits/batch/read | Read a batch of deal split objects by their associated deal object internal ID
[**PostCrmV3ObjectsDealsSplitsBatchUpsertUpsert**](BatchAPI.md#PostCrmV3ObjectsDealsSplitsBatchUpsertUpsert) | **Post** /crm/v3/objects/deals/splits/batch/upsert | Create or replace deal splits for deals with the provided IDs. Deal split percentages for each deal must sum up to 1.0 (100%) and may have up to 8 decimal places



## PostCrmV3ObjectsDealsSplitsBatchReadRead

> BatchResponseDealToDealSplits PostCrmV3ObjectsDealsSplitsBatchReadRead(ctx).BatchInputPublicObjectId(batchInputPublicObjectId).Execute()

Read a batch of deal split objects by their associated deal object internal ID

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
	batchInputPublicObjectId := *openapiclient.NewBatchInputPublicObjectId([]openapiclient.PublicObjectId{*openapiclient.NewPublicObjectId("Id_example")}) // BatchInputPublicObjectId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostCrmV3ObjectsDealsSplitsBatchReadRead(context.Background()).BatchInputPublicObjectId(batchInputPublicObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmV3ObjectsDealsSplitsBatchReadRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ObjectsDealsSplitsBatchReadRead`: BatchResponseDealToDealSplits
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmV3ObjectsDealsSplitsBatchReadRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsDealsSplitsBatchReadReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputPublicObjectId** | [**BatchInputPublicObjectId**](BatchInputPublicObjectId.md) |  | 

### Return type

[**BatchResponseDealToDealSplits**](BatchResponseDealToDealSplits.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsDealsSplitsBatchUpsertUpsert

> BatchResponseDealToDealSplits PostCrmV3ObjectsDealsSplitsBatchUpsertUpsert(ctx).PublicDealSplitsBatchCreateRequest(publicDealSplitsBatchCreateRequest).Execute()

Create or replace deal splits for deals with the provided IDs. Deal split percentages for each deal must sum up to 1.0 (100%) and may have up to 8 decimal places

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
	publicDealSplitsBatchCreateRequest := *openapiclient.NewPublicDealSplitsBatchCreateRequest([]openapiclient.PublicDealSplitsCreateRequest{*openapiclient.NewPublicDealSplitsCreateRequest([]openapiclient.PublicDealSplitInput{*openapiclient.NewPublicDealSplitInput(float32(123), int32(123))}, int32(123))}) // PublicDealSplitsBatchCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostCrmV3ObjectsDealsSplitsBatchUpsertUpsert(context.Background()).PublicDealSplitsBatchCreateRequest(publicDealSplitsBatchCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmV3ObjectsDealsSplitsBatchUpsertUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ObjectsDealsSplitsBatchUpsertUpsert`: BatchResponseDealToDealSplits
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmV3ObjectsDealsSplitsBatchUpsertUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsDealsSplitsBatchUpsertUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicDealSplitsBatchCreateRequest** | [**PublicDealSplitsBatchCreateRequest**](PublicDealSplitsBatchCreateRequest.md) |  | 

### Return type

[**BatchResponseDealToDealSplits**](BatchResponseDealToDealSplits.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

