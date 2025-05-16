# \CallbacksAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete**](CallbacksAPI.md#PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete) | **Post** /automation/v4/actions/callbacks/{callbackId}/complete | Completes a callback
[**PostAutomationV4ActionsCallbacksCompleteCompleteBatch**](CallbacksAPI.md#PostAutomationV4ActionsCallbacksCompleteCompleteBatch) | **Post** /automation/v4/actions/callbacks/complete | Complete a batch of callbacks



## PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete

> PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete(ctx, callbackId).CallbackCompletionRequest(callbackCompletionRequest).Execute()

Completes a callback



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
	callbackId := "callbackId_example" // string | The ID of the action execution.
	callbackCompletionRequest := *openapiclient.NewCallbackCompletionRequest(map[string]string{"key": "Inner_example"}) // CallbackCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CallbacksAPI.PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete(context.Background(), callbackId).CallbackCompletionRequest(callbackCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallbacksAPI.PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callbackId** | **string** | The ID of the action execution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **callbackCompletionRequest** | [**CallbackCompletionRequest**](CallbackCompletionRequest.md) |  | 

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


## PostAutomationV4ActionsCallbacksCompleteCompleteBatch

> PostAutomationV4ActionsCallbacksCompleteCompleteBatch(ctx).BatchInputCallbackCompletionBatchRequest(batchInputCallbackCompletionBatchRequest).Execute()

Complete a batch of callbacks



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
	batchInputCallbackCompletionBatchRequest := *openapiclient.NewBatchInputCallbackCompletionBatchRequest([]openapiclient.CallbackCompletionBatchRequest{*openapiclient.NewCallbackCompletionBatchRequest(map[string]string{"key": "Inner_example"}, "CallbackId_example")}) // BatchInputCallbackCompletionBatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CallbacksAPI.PostAutomationV4ActionsCallbacksCompleteCompleteBatch(context.Background()).BatchInputCallbackCompletionBatchRequest(batchInputCallbackCompletionBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallbacksAPI.PostAutomationV4ActionsCallbacksCompleteCompleteBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationV4ActionsCallbacksCompleteCompleteBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputCallbackCompletionBatchRequest** | [**BatchInputCallbackCompletionBatchRequest**](BatchInputCallbackCompletionBatchRequest.md) |  | 

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

