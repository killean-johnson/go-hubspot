# \PipelineStagesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive**](PipelineStagesAPI.md#DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Delete a pipeline stage
[**GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll**](PipelineStagesAPI.md#GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Return all stages of a pipeline
[**GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById**](PipelineStagesAPI.md#GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Return a pipeline stage by ID
[**PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate**](PipelineStagesAPI.md#PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Update a pipeline stage
[**PostCrmV3PipelinesObjectTypePipelineIdStagesCreate**](PipelineStagesAPI.md#PostCrmV3PipelinesObjectTypePipelineIdStagesCreate) | **Post** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Create a pipeline stage
[**PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace**](PipelineStagesAPI.md#PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Replace a pipeline stage



## DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive

> DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive(ctx, objectType, pipelineId, stageId).Execute()

Delete a pipeline stage



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
	objectType := "objectType_example" // string | 
	pipelineId := "pipelineId_example" // string | 
	stageId := "stageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PipelineStagesAPI.DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive(context.Background(), objectType, pipelineId, stageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesAPI.DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll

> CollectionResponsePipelineStageNoPaging GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll(ctx, objectType, pipelineId).Execute()

Return all stages of a pipeline



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
	objectType := "objectType_example" // string | 
	pipelineId := "pipelineId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelineStagesAPI.GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll(context.Background(), objectType, pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesAPI.GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll`: CollectionResponsePipelineStageNoPaging
	fmt.Fprintf(os.Stdout, "Response from `PipelineStagesAPI.GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PipelinesObjectTypePipelineIdStagesGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePipelineStageNoPaging**](CollectionResponsePipelineStageNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById

> PipelineStage GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById(ctx, objectType, pipelineId, stageId).Execute()

Return a pipeline stage by ID



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
	objectType := "objectType_example" // string | 
	pipelineId := "pipelineId_example" // string | 
	stageId := "stageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelineStagesAPI.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById(context.Background(), objectType, pipelineId, stageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesAPI.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById`: PipelineStage
	fmt.Fprintf(os.Stdout, "Response from `PipelineStagesAPI.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate

> PipelineStage PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate(ctx, objectType, pipelineId, stageId).PipelineStagePatchInput(pipelineStagePatchInput).Execute()

Update a pipeline stage



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
	objectType := "objectType_example" // string | 
	pipelineId := "pipelineId_example" // string | 
	stageId := "stageId_example" // string | 
	pipelineStagePatchInput := *openapiclient.NewPipelineStagePatchInput() // PipelineStagePatchInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelineStagesAPI.PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate(context.Background(), objectType, pipelineId, stageId).PipelineStagePatchInput(pipelineStagePatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesAPI.PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate`: PipelineStage
	fmt.Fprintf(os.Stdout, "Response from `PipelineStagesAPI.PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pipelineStagePatchInput** | [**PipelineStagePatchInput**](PipelineStagePatchInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3PipelinesObjectTypePipelineIdStagesCreate

> PipelineStage PostCrmV3PipelinesObjectTypePipelineIdStagesCreate(ctx, objectType, pipelineId).PipelineStageInput(pipelineStageInput).Execute()

Create a pipeline stage



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
	objectType := "objectType_example" // string | 
	pipelineId := "pipelineId_example" // string | 
	pipelineStageInput := *openapiclient.NewPipelineStageInput(int32(1), "Done") // PipelineStageInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelineStagesAPI.PostCrmV3PipelinesObjectTypePipelineIdStagesCreate(context.Background(), objectType, pipelineId).PipelineStageInput(pipelineStageInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesAPI.PostCrmV3PipelinesObjectTypePipelineIdStagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3PipelinesObjectTypePipelineIdStagesCreate`: PipelineStage
	fmt.Fprintf(os.Stdout, "Response from `PipelineStagesAPI.PostCrmV3PipelinesObjectTypePipelineIdStagesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PipelinesObjectTypePipelineIdStagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineStageInput** | [**PipelineStageInput**](PipelineStageInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace

> PipelineStage PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace(ctx, objectType, pipelineId, stageId).PipelineStageInput(pipelineStageInput).Execute()

Replace a pipeline stage



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
	objectType := "objectType_example" // string | 
	pipelineId := "pipelineId_example" // string | 
	stageId := "stageId_example" // string | 
	pipelineStageInput := *openapiclient.NewPipelineStageInput(int32(1), "Done") // PipelineStageInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelineStagesAPI.PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace(context.Background(), objectType, pipelineId, stageId).PipelineStageInput(pipelineStageInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesAPI.PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace`: PipelineStage
	fmt.Fprintf(os.Stdout, "Response from `PipelineStagesAPI.PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pipelineStageInput** | [**PipelineStageInput**](PipelineStageInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

