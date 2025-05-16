# \SubscriptionsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive**](SubscriptionsAPI.md#DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive) | **Delete** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | Delete event subscription
[**GetWebhooksV3AppIdSubscriptionsGetAll**](SubscriptionsAPI.md#GetWebhooksV3AppIdSubscriptionsGetAll) | **Get** /webhooks/v3/{appId}/subscriptions | Read event subscriptions
[**GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById**](SubscriptionsAPI.md#GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById) | **Get** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | Read an event subscription
[**PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate**](SubscriptionsAPI.md#PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate) | **Patch** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | Update an event subscription
[**PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch**](SubscriptionsAPI.md#PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch) | **Post** /webhooks/v3/{appId}/subscriptions/batch/update | Batch create event subscriptions
[**PostWebhooksV3AppIdSubscriptionsCreate**](SubscriptionsAPI.md#PostWebhooksV3AppIdSubscriptionsCreate) | **Post** /webhooks/v3/{appId}/subscriptions | Create an event subscription



## DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive

> DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive(ctx, subscriptionId, appId).Execute()

Delete event subscription



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
	subscriptionId := int32(56) // int32 | The ID of the event subscription.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionsAPI.DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive(context.Background(), subscriptionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32** | The ID of the event subscription. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksV3AppIdSubscriptionsGetAll

> SubscriptionListResponse GetWebhooksV3AppIdSubscriptionsGetAll(ctx, appId).Execute()

Read event subscriptions



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
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.GetWebhooksV3AppIdSubscriptionsGetAll(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetWebhooksV3AppIdSubscriptionsGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhooksV3AppIdSubscriptionsGetAll`: SubscriptionListResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetWebhooksV3AppIdSubscriptionsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksV3AppIdSubscriptionsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionListResponse**](SubscriptionListResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById

> SubscriptionResponse GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById(ctx, subscriptionId, appId).Execute()

Read an event subscription



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
	subscriptionId := int32(56) // int32 | The ID of the event subscription.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById(context.Background(), subscriptionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32** | The ID of the event subscription. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate

> SubscriptionResponse PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate(ctx, subscriptionId, appId).SubscriptionPatchRequest(subscriptionPatchRequest).Execute()

Update an event subscription



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
	subscriptionId := int32(56) // int32 | The ID of the event subscription.
	appId := int32(56) // int32 | The ID of the app.
	subscriptionPatchRequest := *openapiclient.NewSubscriptionPatchRequest() // SubscriptionPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate(context.Background(), subscriptionId, appId).SubscriptionPatchRequest(subscriptionPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32** | The ID of the event subscription. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subscriptionPatchRequest** | [**SubscriptionPatchRequest**](SubscriptionPatchRequest.md) |  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch

> BatchResponseSubscriptionResponse PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch(ctx, appId).BatchInputSubscriptionBatchUpdateRequest(batchInputSubscriptionBatchUpdateRequest).Execute()

Batch create event subscriptions



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
	appId := int32(56) // int32 | The ID of the app.
	batchInputSubscriptionBatchUpdateRequest := *openapiclient.NewBatchInputSubscriptionBatchUpdateRequest([]openapiclient.SubscriptionBatchUpdateRequest{*openapiclient.NewSubscriptionBatchUpdateRequest(false, int32(123))}) // BatchInputSubscriptionBatchUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch(context.Background(), appId).BatchInputSubscriptionBatchUpdateRequest(batchInputSubscriptionBatchUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch`: BatchResponseSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputSubscriptionBatchUpdateRequest** | [**BatchInputSubscriptionBatchUpdateRequest**](BatchInputSubscriptionBatchUpdateRequest.md) |  | 

### Return type

[**BatchResponseSubscriptionResponse**](BatchResponseSubscriptionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksV3AppIdSubscriptionsCreate

> SubscriptionResponse PostWebhooksV3AppIdSubscriptionsCreate(ctx, appId).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()

Create an event subscription



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
	appId := int32(56) // int32 | The ID of the app.
	subscriptionCreateRequest := *openapiclient.NewSubscriptionCreateRequest("contact.propertyChange") // SubscriptionCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.PostWebhooksV3AppIdSubscriptionsCreate(context.Background(), appId).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.PostWebhooksV3AppIdSubscriptionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhooksV3AppIdSubscriptionsCreate`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.PostWebhooksV3AppIdSubscriptionsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksV3AppIdSubscriptionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionCreateRequest** | [**SubscriptionCreateRequest**](SubscriptionCreateRequest.md) |  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

