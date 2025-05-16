# \GroupsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName**](GroupsAPI.md#DeleteMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName) | **Delete** /media-bridge/v1/{appId}/properties/{objectType}/groups/{groupName} | 
[**GetMediaBridgeV1AppIdPropertiesObjectTypeGroups**](GroupsAPI.md#GetMediaBridgeV1AppIdPropertiesObjectTypeGroups) | **Get** /media-bridge/v1/{appId}/properties/{objectType}/groups | 
[**GetMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName**](GroupsAPI.md#GetMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName) | **Get** /media-bridge/v1/{appId}/properties/{objectType}/groups/{groupName} | 
[**PatchMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName**](GroupsAPI.md#PatchMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName) | **Patch** /media-bridge/v1/{appId}/properties/{objectType}/groups/{groupName} | 
[**PostMediaBridgeV1AppIdPropertiesObjectTypeGroups**](GroupsAPI.md#PostMediaBridgeV1AppIdPropertiesObjectTypeGroups) | **Post** /media-bridge/v1/{appId}/properties/{objectType}/groups | 



## DeleteMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName

> DeleteMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName(ctx, objectType, groupName).Execute()



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
	groupName := "groupName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.DeleteMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName(context.Background(), objectType, groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupNameRequest struct via the builder pattern


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


## GetMediaBridgeV1AppIdPropertiesObjectTypeGroups

> CollectionResponsePropertyGroupNoPaging GetMediaBridgeV1AppIdPropertiesObjectTypeGroups(ctx, objectType).Execute()



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
	resp, r, err := apiClient.GroupsAPI.GetMediaBridgeV1AppIdPropertiesObjectTypeGroups(context.Background(), objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetMediaBridgeV1AppIdPropertiesObjectTypeGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdPropertiesObjectTypeGroups`: CollectionResponsePropertyGroupNoPaging
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetMediaBridgeV1AppIdPropertiesObjectTypeGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdPropertiesObjectTypeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponsePropertyGroupNoPaging**](CollectionResponsePropertyGroupNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName

> PropertyGroup GetMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName(ctx, objectType, groupName).Execute()



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
	groupName := "groupName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName(context.Background(), objectType, groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName`: PropertyGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName

> PropertyGroup PatchMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName(ctx, objectType, groupName).PropertyGroupUpdate(propertyGroupUpdate).Execute()



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
	groupName := "groupName_example" // string | 
	propertyGroupUpdate := *openapiclient.NewPropertyGroupUpdate() // PropertyGroupUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.PatchMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName(context.Background(), objectType, groupName).PropertyGroupUpdate(propertyGroupUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.PatchMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName`: PropertyGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.PatchMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMediaBridgeV1AppIdPropertiesObjectTypeGroupsGroupNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertyGroupUpdate** | [**PropertyGroupUpdate**](PropertyGroupUpdate.md) |  | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1AppIdPropertiesObjectTypeGroups

> PropertyGroup PostMediaBridgeV1AppIdPropertiesObjectTypeGroups(ctx, objectType).PropertyGroupCreate(propertyGroupCreate).Execute()



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
	propertyGroupCreate := *openapiclient.NewPropertyGroupCreate("Name_example", "Label_example") // PropertyGroupCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeGroups(context.Background(), objectType).PropertyGroupCreate(propertyGroupCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1AppIdPropertiesObjectTypeGroups`: PropertyGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.PostMediaBridgeV1AppIdPropertiesObjectTypeGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1AppIdPropertiesObjectTypeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyGroupCreate** | [**PropertyGroupCreate**](PropertyGroupCreate.md) |  | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

