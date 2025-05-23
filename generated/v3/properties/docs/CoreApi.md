# \CoreAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3PropertiesObjectTypePropertyNameArchive**](CoreAPI.md#DeleteCrmV3PropertiesObjectTypePropertyNameArchive) | **Delete** /crm/v3/properties/{objectType}/{propertyName} | Archive a property
[**GetCrmV3PropertiesObjectTypeGetAll**](CoreAPI.md#GetCrmV3PropertiesObjectTypeGetAll) | **Get** /crm/v3/properties/{objectType} | Read all properties
[**GetCrmV3PropertiesObjectTypePropertyNameGetByName**](CoreAPI.md#GetCrmV3PropertiesObjectTypePropertyNameGetByName) | **Get** /crm/v3/properties/{objectType}/{propertyName} | Read a property
[**PatchCrmV3PropertiesObjectTypePropertyNameUpdate**](CoreAPI.md#PatchCrmV3PropertiesObjectTypePropertyNameUpdate) | **Patch** /crm/v3/properties/{objectType}/{propertyName} | Update a property
[**PostCrmV3PropertiesObjectTypeCreate**](CoreAPI.md#PostCrmV3PropertiesObjectTypeCreate) | **Post** /crm/v3/properties/{objectType} | Create a property



## DeleteCrmV3PropertiesObjectTypePropertyNameArchive

> DeleteCrmV3PropertiesObjectTypePropertyNameArchive(ctx, objectType, propertyName).Execute()

Archive a property



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
	r, err := apiClient.CoreAPI.DeleteCrmV3PropertiesObjectTypePropertyNameArchive(context.Background(), objectType, propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.DeleteCrmV3PropertiesObjectTypePropertyNameArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCrmV3PropertiesObjectTypePropertyNameArchiveRequest struct via the builder pattern


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


## GetCrmV3PropertiesObjectTypeGetAll

> CollectionResponsePropertyNoPaging GetCrmV3PropertiesObjectTypeGetAll(ctx, objectType).Archived(archived).Properties(properties).Execute()

Read all properties



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
	resp, r, err := apiClient.CoreAPI.GetCrmV3PropertiesObjectTypeGetAll(context.Background(), objectType).Archived(archived).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.GetCrmV3PropertiesObjectTypeGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3PropertiesObjectTypeGetAll`: CollectionResponsePropertyNoPaging
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.GetCrmV3PropertiesObjectTypeGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PropertiesObjectTypeGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **properties** | **string** |  | 

### Return type

[**CollectionResponsePropertyNoPaging**](CollectionResponsePropertyNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3PropertiesObjectTypePropertyNameGetByName

> Property GetCrmV3PropertiesObjectTypePropertyNameGetByName(ctx, objectType, propertyName).Archived(archived).Properties(properties).Execute()

Read a property



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
	resp, r, err := apiClient.CoreAPI.GetCrmV3PropertiesObjectTypePropertyNameGetByName(context.Background(), objectType, propertyName).Archived(archived).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.GetCrmV3PropertiesObjectTypePropertyNameGetByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3PropertiesObjectTypePropertyNameGetByName`: Property
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.GetCrmV3PropertiesObjectTypePropertyNameGetByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **properties** | **string** |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3PropertiesObjectTypePropertyNameUpdate

> Property PatchCrmV3PropertiesObjectTypePropertyNameUpdate(ctx, objectType, propertyName).PropertyUpdate(propertyUpdate).Execute()

Update a property



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
	propertyUpdate := *openapiclient.NewPropertyUpdate() // PropertyUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.PatchCrmV3PropertiesObjectTypePropertyNameUpdate(context.Background(), objectType, propertyName).PropertyUpdate(propertyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.PatchCrmV3PropertiesObjectTypePropertyNameUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3PropertiesObjectTypePropertyNameUpdate`: Property
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.PatchCrmV3PropertiesObjectTypePropertyNameUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3PropertiesObjectTypePropertyNameUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertyUpdate** | [**PropertyUpdate**](PropertyUpdate.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3PropertiesObjectTypeCreate

> Property PostCrmV3PropertiesObjectTypeCreate(ctx, objectType).PropertyCreate(propertyCreate).Execute()

Create a property



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
	propertyCreate := *openapiclient.NewPropertyCreate("My Contact Property", "enumeration", "contactinformation", "Name_example", "select") // PropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.PostCrmV3PropertiesObjectTypeCreate(context.Background(), objectType).PropertyCreate(propertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.PostCrmV3PropertiesObjectTypeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3PropertiesObjectTypeCreate`: Property
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.PostCrmV3PropertiesObjectTypeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PropertiesObjectTypeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyCreate** | [**PropertyCreate**](PropertyCreate.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

