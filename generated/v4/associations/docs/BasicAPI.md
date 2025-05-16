# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdArchive**](BasicAPI.md#DeleteCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdArchive) | **Delete** /crm/v4/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Delete
[**GetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPage**](BasicAPI.md#GetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPage) | **Get** /crm/v4/objects/{objectType}/{objectId}/associations/{toObjectType} | List
[**PutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefault**](BasicAPI.md#PutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefault) | **Put** /crm/v4/objects/{fromObjectType}/{fromObjectId}/associations/default/{toObjectType}/{toObjectId} | Create Default
[**PutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreate**](BasicAPI.md#PutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreate) | **Put** /crm/v4/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Create



## DeleteCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdArchive

> DeleteCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdArchive(ctx, objectType, objectId, toObjectType, toObjectId).Execute()

Delete



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
	objectId := "objectId_example" // string | 
	toObjectType := "toObjectType_example" // string | 
	toObjectId := "toObjectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdArchive(context.Background(), objectType, objectId, toObjectType, toObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**objectId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdArchiveRequest struct via the builder pattern


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


## GetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPage

> CollectionResponseMultiAssociatedObjectWithLabelForwardPaging GetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPage(ctx, objectType, objectId, toObjectType).After(after).Limit(limit).Execute()

List



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
	objectId := "objectId_example" // string | 
	toObjectType := "toObjectType_example" // string | 
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPage(context.Background(), objectType, objectId, toObjectType).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPage`: CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**objectId** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseMultiAssociatedObjectWithLabelForwardPaging**](CollectionResponseMultiAssociatedObjectWithLabelForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefault

> BatchResponsePublicDefaultAssociation PutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefault(ctx, fromObjectType, fromObjectId, toObjectType, toObjectId).Execute()

Create Default



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
	fromObjectType := "fromObjectType_example" // string | 
	fromObjectId := "fromObjectId_example" // string | 
	toObjectType := "toObjectType_example" // string | 
	toObjectId := "toObjectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefault(context.Background(), fromObjectType, fromObjectId, toObjectType, toObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefault`: BatchResponsePublicDefaultAssociation
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**fromObjectId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BatchResponsePublicDefaultAssociation**](BatchResponsePublicDefaultAssociation.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreate

> LabelsBetweenObjectPair PutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreate(ctx, objectType, objectId, toObjectType, toObjectId).AssociationSpec(associationSpec).Execute()

Create



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
	objectId := "objectId_example" // string | 
	toObjectType := "toObjectType_example" // string | 
	toObjectId := "toObjectId_example" // string | 
	associationSpec := []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))} // []AssociationSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreate(context.Background(), objectType, objectId, toObjectType, toObjectId).AssociationSpec(associationSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreate`: LabelsBetweenObjectPair
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**objectId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **associationSpec** | [**[]AssociationSpec**](AssociationSpec.md) |  | 

### Return type

[**LabelsBetweenObjectPair**](LabelsBetweenObjectPair.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

