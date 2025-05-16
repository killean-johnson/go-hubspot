# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3PropertiesObjectTypeBatchArchiveArchive**](BatchAPI.md#PostCrmV3PropertiesObjectTypeBatchArchiveArchive) | **Post** /crm/v3/properties/{objectType}/batch/archive | Archive a batch of properties
[**PostCrmV3PropertiesObjectTypeBatchCreateCreate**](BatchAPI.md#PostCrmV3PropertiesObjectTypeBatchCreateCreate) | **Post** /crm/v3/properties/{objectType}/batch/create | Create a batch of properties
[**PostCrmV3PropertiesObjectTypeBatchReadRead**](BatchAPI.md#PostCrmV3PropertiesObjectTypeBatchReadRead) | **Post** /crm/v3/properties/{objectType}/batch/read | Read a batch of properties



## PostCrmV3PropertiesObjectTypeBatchArchiveArchive

> PostCrmV3PropertiesObjectTypeBatchArchiveArchive(ctx, objectType).BatchInputPropertyName(batchInputPropertyName).Execute()

Archive a batch of properties



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
	batchInputPropertyName := *openapiclient.NewBatchInputPropertyName([]openapiclient.PropertyName{*openapiclient.NewPropertyName("my_custom_property")}) // BatchInputPropertyName | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.PostCrmV3PropertiesObjectTypeBatchArchiveArchive(context.Background(), objectType).BatchInputPropertyName(batchInputPropertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmV3PropertiesObjectTypeBatchArchiveArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PropertiesObjectTypeBatchArchiveArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyName** | [**BatchInputPropertyName**](BatchInputPropertyName.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3PropertiesObjectTypeBatchCreateCreate

> BatchResponseProperty PostCrmV3PropertiesObjectTypeBatchCreateCreate(ctx, objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()

Create a batch of properties



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
	batchInputPropertyCreate := *openapiclient.NewBatchInputPropertyCreate([]openapiclient.PropertyCreate{*openapiclient.NewPropertyCreate("My Contact Property", "enumeration", "contactinformation", "Name_example", "select")}) // BatchInputPropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostCrmV3PropertiesObjectTypeBatchCreateCreate(context.Background(), objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmV3PropertiesObjectTypeBatchCreateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3PropertiesObjectTypeBatchCreateCreate`: BatchResponseProperty
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmV3PropertiesObjectTypeBatchCreateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PropertiesObjectTypeBatchCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyCreate** | [**BatchInputPropertyCreate**](BatchInputPropertyCreate.md) |  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3PropertiesObjectTypeBatchReadRead

> BatchResponseProperty PostCrmV3PropertiesObjectTypeBatchReadRead(ctx, objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Execute()

Read a batch of properties



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
	batchReadInputPropertyName := *openapiclient.NewBatchReadInputPropertyName(false, []openapiclient.PropertyName{*openapiclient.NewPropertyName("my_custom_property")}) // BatchReadInputPropertyName | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostCrmV3PropertiesObjectTypeBatchReadRead(context.Background(), objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmV3PropertiesObjectTypeBatchReadRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3PropertiesObjectTypeBatchReadRead`: BatchResponseProperty
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmV3PropertiesObjectTypeBatchReadRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PropertiesObjectTypeBatchReadReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchReadInputPropertyName** | [**BatchReadInputPropertyName**](BatchReadInputPropertyName.md) |  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

