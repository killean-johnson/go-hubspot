# \PublicExportsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ExportsExportAsyncTasksTaskIdStatusGetStatus**](PublicExportsAPI.md#GetCrmV3ExportsExportAsyncTasksTaskIdStatusGetStatus) | **Get** /crm/v3/exports/export/async/tasks/{taskId}/status | Get the status of the export including the URL to download the file
[**PostCrmV3ExportsExportAsyncStart**](PublicExportsAPI.md#PostCrmV3ExportsExportAsyncStart) | **Post** /crm/v3/exports/export/async | Start an export



## GetCrmV3ExportsExportAsyncTasksTaskIdStatusGetStatus

> ActionResponseWithSingleResultURI GetCrmV3ExportsExportAsyncTasksTaskIdStatusGetStatus(ctx, taskId).Execute()

Get the status of the export including the URL to download the file



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
	taskId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicExportsAPI.GetCrmV3ExportsExportAsyncTasksTaskIdStatusGetStatus(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicExportsAPI.GetCrmV3ExportsExportAsyncTasksTaskIdStatusGetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExportsExportAsyncTasksTaskIdStatusGetStatus`: ActionResponseWithSingleResultURI
	fmt.Fprintf(os.Stdout, "Response from `PublicExportsAPI.GetCrmV3ExportsExportAsyncTasksTaskIdStatusGetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExportsExportAsyncTasksTaskIdStatusGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResponseWithSingleResultURI**](ActionResponseWithSingleResultURI.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExportsExportAsyncStart

> TaskLocator PostCrmV3ExportsExportAsyncStart(ctx).PublicExportRequest(publicExportRequest).Execute()

Start an export



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
	publicExportRequest := openapiclient.PublicExportRequest{PublicExportListRequest: openapiclient.NewPublicExportListRequest("ExportType_example", "Format_example", "ExportName_example", []string{"ObjectProperties_example"}, "ObjectType_example", "Language_example", []string{"ExportInternalValuesOptions_example"}, false, "ListId_example")} // PublicExportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicExportsAPI.PostCrmV3ExportsExportAsyncStart(context.Background()).PublicExportRequest(publicExportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicExportsAPI.PostCrmV3ExportsExportAsyncStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ExportsExportAsyncStart`: TaskLocator
	fmt.Fprintf(os.Stdout, "Response from `PublicExportsAPI.PostCrmV3ExportsExportAsyncStart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExportsExportAsyncStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicExportRequest** | [**PublicExportRequest**](PublicExportRequest.md) |  | 

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

