# \DefinitionConfigurationsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV4AssociationsDefinitionsConfigurationsAllGetAll**](DefinitionConfigurationsAPI.md#GetCrmV4AssociationsDefinitionsConfigurationsAllGetAll) | **Get** /crm/v4/associations/definitions/configurations/all | Read All
[**GetCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeGetAllBetweenTwoObjectTypes**](DefinitionConfigurationsAPI.md#GetCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeGetAllBetweenTwoObjectTypes) | **Get** /crm/v4/associations/definitions/configurations/{fromObjectType}/{toObjectType} | Read
[**PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateBatchCreate**](DefinitionConfigurationsAPI.md#PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateBatchCreate) | **Post** /crm/v4/associations/definitions/configurations/{fromObjectType}/{toObjectType}/batch/create | Create
[**PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeBatchRemove**](DefinitionConfigurationsAPI.md#PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeBatchRemove) | **Post** /crm/v4/associations/definitions/configurations/{fromObjectType}/{toObjectType}/batch/purge | Delete
[**PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateBatchUpdate**](DefinitionConfigurationsAPI.md#PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateBatchUpdate) | **Post** /crm/v4/associations/definitions/configurations/{fromObjectType}/{toObjectType}/batch/update | Update



## GetCrmV4AssociationsDefinitionsConfigurationsAllGetAll

> CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging GetCrmV4AssociationsDefinitionsConfigurationsAllGetAll(ctx).Execute()

Read All



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionConfigurationsAPI.GetCrmV4AssociationsDefinitionsConfigurationsAllGetAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionConfigurationsAPI.GetCrmV4AssociationsDefinitionsConfigurationsAllGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV4AssociationsDefinitionsConfigurationsAllGetAll`: CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging
	fmt.Fprintf(os.Stdout, "Response from `DefinitionConfigurationsAPI.GetCrmV4AssociationsDefinitionsConfigurationsAllGetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV4AssociationsDefinitionsConfigurationsAllGetAllRequest struct via the builder pattern


### Return type

[**CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging**](CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeGetAllBetweenTwoObjectTypes

> CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging GetCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeGetAllBetweenTwoObjectTypes(ctx, fromObjectType, toObjectType).Execute()

Read



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
	toObjectType := "toObjectType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionConfigurationsAPI.GetCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeGetAllBetweenTwoObjectTypes(context.Background(), fromObjectType, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionConfigurationsAPI.GetCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeGetAllBetweenTwoObjectTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeGetAllBetweenTwoObjectTypes`: CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging
	fmt.Fprintf(os.Stdout, "Response from `DefinitionConfigurationsAPI.GetCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeGetAllBetweenTwoObjectTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeGetAllBetweenTwoObjectTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging**](CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateBatchCreate

> BatchResponsePublicAssociationDefinitionUserConfiguration PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateBatchCreate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationCreateRequest(batchInputPublicAssociationDefinitionConfigurationCreateRequest).Execute()

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
	fromObjectType := "fromObjectType_example" // string | 
	toObjectType := "toObjectType_example" // string | 
	batchInputPublicAssociationDefinitionConfigurationCreateRequest := *openapiclient.NewBatchInputPublicAssociationDefinitionConfigurationCreateRequest([]openapiclient.PublicAssociationDefinitionConfigurationCreateRequest{*openapiclient.NewPublicAssociationDefinitionConfigurationCreateRequest(int32(123), "Category_example", int32(123))}) // BatchInputPublicAssociationDefinitionConfigurationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionConfigurationsAPI.PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateBatchCreate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationCreateRequest(batchInputPublicAssociationDefinitionConfigurationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionConfigurationsAPI.PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateBatchCreate`: BatchResponsePublicAssociationDefinitionUserConfiguration
	fmt.Fprintf(os.Stdout, "Response from `DefinitionConfigurationsAPI.PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateBatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationDefinitionConfigurationCreateRequest** | [**BatchInputPublicAssociationDefinitionConfigurationCreateRequest**](BatchInputPublicAssociationDefinitionConfigurationCreateRequest.md) |  | 

### Return type

[**BatchResponsePublicAssociationDefinitionUserConfiguration**](BatchResponsePublicAssociationDefinitionUserConfiguration.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeBatchRemove

> PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeBatchRemove(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationSpec(batchInputPublicAssociationSpec).Execute()

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
	fromObjectType := "fromObjectType_example" // string | 
	toObjectType := "toObjectType_example" // string | 
	batchInputPublicAssociationSpec := *openapiclient.NewBatchInputPublicAssociationSpec([]openapiclient.PublicAssociationSpec{*openapiclient.NewPublicAssociationSpec(int32(123), "Category_example")}) // BatchInputPublicAssociationSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefinitionConfigurationsAPI.PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeBatchRemove(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationSpec(batchInputPublicAssociationSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionConfigurationsAPI.PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeBatchRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeBatchRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationSpec** | [**BatchInputPublicAssociationSpec**](BatchInputPublicAssociationSpec.md) |  | 

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


## PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateBatchUpdate

> BatchResponsePublicAssociationDefinitionConfigurationUpdateResult PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateBatchUpdate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationUpdateRequest(batchInputPublicAssociationDefinitionConfigurationUpdateRequest).Execute()

Update



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
	toObjectType := "toObjectType_example" // string | 
	batchInputPublicAssociationDefinitionConfigurationUpdateRequest := *openapiclient.NewBatchInputPublicAssociationDefinitionConfigurationUpdateRequest([]openapiclient.PublicAssociationDefinitionConfigurationUpdateRequest{*openapiclient.NewPublicAssociationDefinitionConfigurationUpdateRequest(int32(123), "Category_example", int32(123))}) // BatchInputPublicAssociationDefinitionConfigurationUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionConfigurationsAPI.PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateBatchUpdate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationUpdateRequest(batchInputPublicAssociationDefinitionConfigurationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionConfigurationsAPI.PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateBatchUpdate`: BatchResponsePublicAssociationDefinitionConfigurationUpdateResult
	fmt.Fprintf(os.Stdout, "Response from `DefinitionConfigurationsAPI.PostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateBatchUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationDefinitionConfigurationUpdateRequest** | [**BatchInputPublicAssociationDefinitionConfigurationUpdateRequest**](BatchInputPublicAssociationDefinitionConfigurationUpdateRequest.md) |  | 

### Return type

[**BatchResponsePublicAssociationDefinitionConfigurationUpdateResult**](BatchResponsePublicAssociationDefinitionConfigurationUpdateResult.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

