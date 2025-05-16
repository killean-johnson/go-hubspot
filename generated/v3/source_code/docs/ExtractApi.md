# \ExtractAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus**](ExtractAPI.md#GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus) | **Get** /cms/v3/source-code/extract/async/tasks/{taskId}/status | Get extraction status
[**PostCmsV3SourceCodeExtractAsyncDoAsync**](ExtractAPI.md#PostCmsV3SourceCodeExtractAsyncDoAsync) | **Post** /cms/v3/source-code/extract/async | Extract a zip file



## GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus

> ActionResponse GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus(ctx, taskId).Execute()

Get extraction status



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
	taskId := int32(56) // int32 | The extraction task ID returned by the initial `extract/async` request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtractAPI.GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtractAPI.GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus`: ActionResponse
	fmt.Fprintf(os.Stdout, "Response from `ExtractAPI.GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** | The extraction task ID returned by the initial &#x60;extract/async&#x60; request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3SourceCodeExtractAsyncDoAsync

> TaskLocator PostCmsV3SourceCodeExtractAsyncDoAsync(ctx).FileExtractRequest(fileExtractRequest).Execute()

Extract a zip file



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
	fileExtractRequest := *openapiclient.NewFileExtractRequest("Path_example") // FileExtractRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtractAPI.PostCmsV3SourceCodeExtractAsyncDoAsync(context.Background()).FileExtractRequest(fileExtractRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtractAPI.PostCmsV3SourceCodeExtractAsyncDoAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3SourceCodeExtractAsyncDoAsync`: TaskLocator
	fmt.Fprintf(os.Stdout, "Response from `ExtractAPI.PostCmsV3SourceCodeExtractAsyncDoAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3SourceCodeExtractAsyncDoAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileExtractRequest** | [**FileExtractRequest**](FileExtractRequest.md) |  | 

### Return type

[**TaskLocator**](TaskLocator.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

