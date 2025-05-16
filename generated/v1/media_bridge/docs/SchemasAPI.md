# \SchemasAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMediaBridgeV1AppIdSchemasObjectTypeAssociationsAssociationId**](SchemasAPI.md#DeleteMediaBridgeV1AppIdSchemasObjectTypeAssociationsAssociationId) | **Delete** /media-bridge/v1/{appId}/schemas/{objectType}/associations/{associationId} | 
[**GetMediaBridgeV1AppIdSchemas**](SchemasAPI.md#GetMediaBridgeV1AppIdSchemas) | **Get** /media-bridge/v1/{appId}/schemas | 
[**GetMediaBridgeV1AppIdSchemasObjectType**](SchemasAPI.md#GetMediaBridgeV1AppIdSchemasObjectType) | **Get** /media-bridge/v1/{appId}/schemas/{objectType} | 
[**PatchMediaBridgeV1AppIdSchemasObjectType**](SchemasAPI.md#PatchMediaBridgeV1AppIdSchemasObjectType) | **Patch** /media-bridge/v1/{appId}/schemas/{objectType} | 
[**PostMediaBridgeV1AppIdSchemasObjectTypeAssociations**](SchemasAPI.md#PostMediaBridgeV1AppIdSchemasObjectTypeAssociations) | **Post** /media-bridge/v1/{appId}/schemas/{objectType}/associations | 



## DeleteMediaBridgeV1AppIdSchemasObjectTypeAssociationsAssociationId

> DeleteMediaBridgeV1AppIdSchemasObjectTypeAssociationsAssociationId(ctx, objectType, associationId).Execute()



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
	associationId := "associationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SchemasAPI.DeleteMediaBridgeV1AppIdSchemasObjectTypeAssociationsAssociationId(context.Background(), objectType, associationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemasAPI.DeleteMediaBridgeV1AppIdSchemasObjectTypeAssociationsAssociationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**associationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaBridgeV1AppIdSchemasObjectTypeAssociationsAssociationIdRequest struct via the builder pattern


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


## GetMediaBridgeV1AppIdSchemas

> CollectionResponseObjectSchemaNoPaging GetMediaBridgeV1AppIdSchemas(ctx).Archived(archived).Execute()



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
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemasAPI.GetMediaBridgeV1AppIdSchemas(context.Background()).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemasAPI.GetMediaBridgeV1AppIdSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdSchemas`: CollectionResponseObjectSchemaNoPaging
	fmt.Fprintf(os.Stdout, "Response from `SchemasAPI.GetMediaBridgeV1AppIdSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponseObjectSchemaNoPaging**](CollectionResponseObjectSchemaNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaBridgeV1AppIdSchemasObjectType

> ObjectSchema GetMediaBridgeV1AppIdSchemasObjectType(ctx, objectType).Execute()



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
	resp, r, err := apiClient.SchemasAPI.GetMediaBridgeV1AppIdSchemasObjectType(context.Background(), objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemasAPI.GetMediaBridgeV1AppIdSchemasObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdSchemasObjectType`: ObjectSchema
	fmt.Fprintf(os.Stdout, "Response from `SchemasAPI.GetMediaBridgeV1AppIdSchemasObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdSchemasObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMediaBridgeV1AppIdSchemasObjectType

> ObjectTypeDefinition PatchMediaBridgeV1AppIdSchemasObjectType(ctx, objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()



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
	objectTypeDefinitionPatch := *openapiclient.NewObjectTypeDefinitionPatch(false) // ObjectTypeDefinitionPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemasAPI.PatchMediaBridgeV1AppIdSchemasObjectType(context.Background(), objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemasAPI.PatchMediaBridgeV1AppIdSchemasObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMediaBridgeV1AppIdSchemasObjectType`: ObjectTypeDefinition
	fmt.Fprintf(os.Stdout, "Response from `SchemasAPI.PatchMediaBridgeV1AppIdSchemasObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMediaBridgeV1AppIdSchemasObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectTypeDefinitionPatch** | [**ObjectTypeDefinitionPatch**](ObjectTypeDefinitionPatch.md) |  | 

### Return type

[**ObjectTypeDefinition**](ObjectTypeDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1AppIdSchemasObjectTypeAssociations

> AssociationDefinition PostMediaBridgeV1AppIdSchemasObjectTypeAssociations(ctx, objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()



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
	associationDefinitionEgg := *openapiclient.NewAssociationDefinitionEgg("FromObjectTypeId_example", "ToObjectTypeId_example") // AssociationDefinitionEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemasAPI.PostMediaBridgeV1AppIdSchemasObjectTypeAssociations(context.Background(), objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemasAPI.PostMediaBridgeV1AppIdSchemasObjectTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1AppIdSchemasObjectTypeAssociations`: AssociationDefinition
	fmt.Fprintf(os.Stdout, "Response from `SchemasAPI.PostMediaBridgeV1AppIdSchemasObjectTypeAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1AppIdSchemasObjectTypeAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **associationDefinitionEgg** | [**AssociationDefinitionEgg**](AssociationDefinitionEgg.md) |  | 

### Return type

[**AssociationDefinition**](AssociationDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

