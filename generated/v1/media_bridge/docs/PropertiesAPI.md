# \PropertiesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName**](PropertiesAPI.md#DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName) | **Delete** /media-bridge/v1/{appId}/properties/{objectType}/{propertyName} | 
[**GetMediaBridgeV1AppIdPropertiesObjectType**](PropertiesAPI.md#GetMediaBridgeV1AppIdPropertiesObjectType) | **Get** /media-bridge/v1/{appId}/properties/{objectType} | 
[**GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName**](PropertiesAPI.md#GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName) | **Get** /media-bridge/v1/{appId}/properties/{objectType}/{propertyName} | 
[**PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName**](PropertiesAPI.md#PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName) | **Patch** /media-bridge/v1/{appId}/properties/{objectType}/{propertyName} | 
[**PostMediaBridgeV1AppIdPropertiesObjectType**](PropertiesAPI.md#PostMediaBridgeV1AppIdPropertiesObjectType) | **Post** /media-bridge/v1/{appId}/properties/{objectType} | 
[**PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive**](PropertiesAPI.md#PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive) | **Post** /media-bridge/v1/{appId}/properties/{objectType}/batch/archive | 
[**PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate**](PropertiesAPI.md#PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate) | **Post** /media-bridge/v1/{appId}/properties/{objectType}/batch/create | 
[**PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead**](PropertiesAPI.md#PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead) | **Post** /media-bridge/v1/{appId}/properties/{objectType}/batch/read | 



## DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName

> DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName(ctx, objectType, propertyName).Execute()



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
	propertyName := "propertyName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PropertiesAPI.DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName(context.Background(), objectType, propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertiesAPI.DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest struct via the builder pattern


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


## GetMediaBridgeV1AppIdPropertiesObjectType

> CollectionResponsePropertyNoPaging GetMediaBridgeV1AppIdPropertiesObjectType(ctx, objectType).Archived(archived).Properties(properties).Execute()



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
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	properties := "properties_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertiesAPI.GetMediaBridgeV1AppIdPropertiesObjectType(context.Background(), objectType).Archived(archived).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertiesAPI.GetMediaBridgeV1AppIdPropertiesObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdPropertiesObjectType`: CollectionResponsePropertyNoPaging
	fmt.Fprintf(os.Stdout, "Response from `PropertiesAPI.GetMediaBridgeV1AppIdPropertiesObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **properties** | **string** |  | 

### Return type

[**CollectionResponsePropertyNoPaging**](CollectionResponsePropertyNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName

> Property GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName(ctx, objectType, propertyName).Archived(archived).Properties(properties).Execute()



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
	propertyName := "propertyName_example" // string | 
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	properties := "properties_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertiesAPI.GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName(context.Background(), objectType, propertyName).Archived(archived).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertiesAPI.GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName`: Property
	fmt.Fprintf(os.Stdout, "Response from `PropertiesAPI.GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **properties** | **string** |  | 

### Return type

[**Property**](Property.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName

> Property PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName(ctx, objectType, propertyName).MediaBridgePropertyUpdate(mediaBridgePropertyUpdate).Execute()



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
	propertyName := "propertyName_example" // string | 
	mediaBridgePropertyUpdate := *openapiclient.NewMediaBridgePropertyUpdate() // MediaBridgePropertyUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertiesAPI.PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName(context.Background(), objectType, propertyName).MediaBridgePropertyUpdate(mediaBridgePropertyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertiesAPI.PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName`: Property
	fmt.Fprintf(os.Stdout, "Response from `PropertiesAPI.PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaBridgePropertyUpdate** | [**MediaBridgePropertyUpdate**](MediaBridgePropertyUpdate.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1AppIdPropertiesObjectType

> Property PostMediaBridgeV1AppIdPropertiesObjectType(ctx, objectType).PropertyCreate(propertyCreate).Execute()



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
	propertyCreate := *openapiclient.NewPropertyCreate("Label_example", "Type_example", "GroupName_example", "Name_example", "FieldType_example") // PropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectType(context.Background(), objectType).PropertyCreate(propertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1AppIdPropertiesObjectType`: Property
	fmt.Fprintf(os.Stdout, "Response from `PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1AppIdPropertiesObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyCreate** | [**PropertyCreate**](PropertyCreate.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive

> PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive(ctx, objectType).BatchInputPropertyName(batchInputPropertyName).Execute()



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
	batchInputPropertyName := *openapiclient.NewBatchInputPropertyName([]openapiclient.PropertyName{*openapiclient.NewPropertyName("Name_example")}) // BatchInputPropertyName | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive(context.Background(), objectType).BatchInputPropertyName(batchInputPropertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyName** | [**BatchInputPropertyName**](BatchInputPropertyName.md) |  | 

### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate

> BatchResponseProperty PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate(ctx, objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()



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
	batchInputPropertyCreate := *openapiclient.NewBatchInputPropertyCreate([]openapiclient.PropertyCreate{*openapiclient.NewPropertyCreate("Label_example", "Type_example", "GroupName_example", "Name_example", "FieldType_example")}) // BatchInputPropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate(context.Background(), objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate`: BatchResponseProperty
	fmt.Fprintf(os.Stdout, "Response from `PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyCreate** | [**BatchInputPropertyCreate**](BatchInputPropertyCreate.md) |  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead

> BatchResponseProperty PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead(ctx, objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Execute()



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
	batchReadInputPropertyName := *openapiclient.NewBatchReadInputPropertyName(false, []openapiclient.PropertyName{*openapiclient.NewPropertyName("Name_example")}) // BatchReadInputPropertyName | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead(context.Background(), objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead`: BatchResponseProperty
	fmt.Fprintf(os.Stdout, "Response from `PropertiesAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchReadInputPropertyName** | [**BatchReadInputPropertyName**](BatchReadInputPropertyName.md) |  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

