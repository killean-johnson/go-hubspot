# \PipelinesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3PipelinesObjectTypePipelineIdArchive**](PipelinesAPI.md#DeleteCrmV3PipelinesObjectTypePipelineIdArchive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId} | Delete a pipeline
[**GetCrmV3PipelinesObjectTypeGetAll**](PipelinesAPI.md#GetCrmV3PipelinesObjectTypeGetAll) | **Get** /crm/v3/pipelines/{objectType} | Retrieve all pipelines
[**GetCrmV3PipelinesObjectTypePipelineIdGetById**](PipelinesAPI.md#GetCrmV3PipelinesObjectTypePipelineIdGetById) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId} | Return a pipeline by ID
[**PatchCrmV3PipelinesObjectTypePipelineIdUpdate**](PipelinesAPI.md#PatchCrmV3PipelinesObjectTypePipelineIdUpdate) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId} | Update a pipeline
[**PostCrmV3PipelinesObjectTypeCreate**](PipelinesAPI.md#PostCrmV3PipelinesObjectTypeCreate) | **Post** /crm/v3/pipelines/{objectType} | Create a pipeline
[**PutCrmV3PipelinesObjectTypePipelineIdReplace**](PipelinesAPI.md#PutCrmV3PipelinesObjectTypePipelineIdReplace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId} | Replace a pipeline



## DeleteCrmV3PipelinesObjectTypePipelineIdArchive

> DeleteCrmV3PipelinesObjectTypePipelineIdArchive(ctx, objectType, pipelineId).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()

Delete a pipeline



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
	validateReferencesBeforeDelete := true // bool |  (optional) (default to false)
	validateDealStageUsagesBeforeDelete := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PipelinesAPI.DeleteCrmV3PipelinesObjectTypePipelineIdArchive(context.Background(), objectType, pipelineId).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.DeleteCrmV3PipelinesObjectTypePipelineIdArchive``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3PipelinesObjectTypePipelineIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateReferencesBeforeDelete** | **bool** |  | [default to false]
 **validateDealStageUsagesBeforeDelete** | **bool** |  | [default to false]

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


## GetCrmV3PipelinesObjectTypeGetAll

> CollectionResponsePipelineNoPaging GetCrmV3PipelinesObjectTypeGetAll(ctx, objectType).Execute()

Retrieve all pipelines



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.GetCrmV3PipelinesObjectTypeGetAll(context.Background(), objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.GetCrmV3PipelinesObjectTypeGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3PipelinesObjectTypeGetAll`: CollectionResponsePipelineNoPaging
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.GetCrmV3PipelinesObjectTypeGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PipelinesObjectTypeGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponsePipelineNoPaging**](CollectionResponsePipelineNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3PipelinesObjectTypePipelineIdGetById

> Pipeline GetCrmV3PipelinesObjectTypePipelineIdGetById(ctx, objectType, pipelineId).Execute()

Return a pipeline by ID



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
	resp, r, err := apiClient.PipelinesAPI.GetCrmV3PipelinesObjectTypePipelineIdGetById(context.Background(), objectType, pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.GetCrmV3PipelinesObjectTypePipelineIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3PipelinesObjectTypePipelineIdGetById`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.GetCrmV3PipelinesObjectTypePipelineIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PipelinesObjectTypePipelineIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3PipelinesObjectTypePipelineIdUpdate

> Pipeline PatchCrmV3PipelinesObjectTypePipelineIdUpdate(ctx, objectType, pipelineId).PipelinePatchInput(pipelinePatchInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()

Update a pipeline



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
	pipelinePatchInput := *openapiclient.NewPipelinePatchInput() // PipelinePatchInput | 
	validateReferencesBeforeDelete := true // bool |  (optional) (default to false)
	validateDealStageUsagesBeforeDelete := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.PatchCrmV3PipelinesObjectTypePipelineIdUpdate(context.Background(), objectType, pipelineId).PipelinePatchInput(pipelinePatchInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PatchCrmV3PipelinesObjectTypePipelineIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3PipelinesObjectTypePipelineIdUpdate`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PatchCrmV3PipelinesObjectTypePipelineIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelinePatchInput** | [**PipelinePatchInput**](PipelinePatchInput.md) |  | 
 **validateReferencesBeforeDelete** | **bool** |  | [default to false]
 **validateDealStageUsagesBeforeDelete** | **bool** |  | [default to false]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3PipelinesObjectTypeCreate

> Pipeline PostCrmV3PipelinesObjectTypeCreate(ctx, objectType).PipelineInput(pipelineInput).Execute()

Create a pipeline



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
	pipelineInput := *openapiclient.NewPipelineInput(int32(0), []openapiclient.PipelineStageInput{*openapiclient.NewPipelineStageInput(int32(1), "Done")}, "My replaced pipeline") // PipelineInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.PostCrmV3PipelinesObjectTypeCreate(context.Background(), objectType).PipelineInput(pipelineInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PostCrmV3PipelinesObjectTypeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3PipelinesObjectTypeCreate`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PostCrmV3PipelinesObjectTypeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PipelinesObjectTypeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pipelineInput** | [**PipelineInput**](PipelineInput.md) |  | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3PipelinesObjectTypePipelineIdReplace

> Pipeline PutCrmV3PipelinesObjectTypePipelineIdReplace(ctx, objectType, pipelineId).PipelineInput(pipelineInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()

Replace a pipeline



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
	pipelineInput := *openapiclient.NewPipelineInput(int32(0), []openapiclient.PipelineStageInput{*openapiclient.NewPipelineStageInput(int32(1), "Done")}, "My replaced pipeline") // PipelineInput | 
	validateReferencesBeforeDelete := true // bool |  (optional) (default to false)
	validateDealStageUsagesBeforeDelete := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.PutCrmV3PipelinesObjectTypePipelineIdReplace(context.Background(), objectType, pipelineId).PipelineInput(pipelineInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PutCrmV3PipelinesObjectTypePipelineIdReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3PipelinesObjectTypePipelineIdReplace`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PutCrmV3PipelinesObjectTypePipelineIdReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineInput** | [**PipelineInput**](PipelineInput.md) |  | 
 **validateReferencesBeforeDelete** | **bool** |  | [default to false]
 **validateDealStageUsagesBeforeDelete** | **bool** |  | [default to false]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

