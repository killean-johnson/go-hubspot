# \FunctionsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType**](FunctionsAPI.md#DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Delete a function for a definition
[**DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive**](FunctionsAPI.md#DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Archive a function for a definition
[**GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById**](FunctionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Retrieve a function from a given definition
[**GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType**](FunctionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Retrieve functions by a type for a given definition
[**GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage**](FunctionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions | Retrieve functions for a given definition
[**PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType**](FunctionsAPI.md#PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Insert a function for a definition
[**PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace**](FunctionsAPI.md#PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Update a function for a definition



## DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType

> DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType(ctx, definitionId, functionType, appId).Execute()

Delete a function for a definition



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
	definitionId := "definitionId_example" // string | The ID of the definition.
	functionType := "functionType_example" // string | The type of function. Can be `PRE_ACTION_EXECUTION`, `PRE_FETCH_OPTIONS`, `POST_FETCH_OPTIONS`, `POST_ACTION_EXECUTION`.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FunctionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType(context.Background(), definitionId, functionType, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the definition. | 
**functionType** | **string** | The type of function. Can be &#x60;PRE_ACTION_EXECUTION&#x60;, &#x60;PRE_FETCH_OPTIONS&#x60;, &#x60;POST_FETCH_OPTIONS&#x60;, &#x60;POST_ACTION_EXECUTION&#x60;. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionTypeRequest struct via the builder pattern


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


## DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive

> DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive(ctx, definitionId, functionType, functionId, appId).Execute()

Archive a function for a definition

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
	definitionId := "definitionId_example" // string | 
	functionType := "functionType_example" // string | 
	functionId := "functionId_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FunctionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive(context.Background(), definitionId, functionType, functionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**functionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchiveRequest struct via the builder pattern


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


## GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById

> PublicActionFunction GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById(ctx, definitionId, functionType, functionId, appId).Execute()

Retrieve a function from a given definition



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
	definitionId := "definitionId_example" // string | The ID of the definition.
	functionType := "functionType_example" // string | The type of function. Can be `PRE_ACTION_EXECUTION`, `PRE_FETCH_OPTIONS`, `POST_FETCH_OPTIONS`, `POST_ACTION_EXECUTION`.
	functionId := "functionId_example" // string | The ID of the function.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById(context.Background(), definitionId, functionType, functionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById`: PublicActionFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the definition. | 
**functionType** | **string** | The type of function. Can be &#x60;PRE_ACTION_EXECUTION&#x60;, &#x60;PRE_FETCH_OPTIONS&#x60;, &#x60;POST_FETCH_OPTIONS&#x60;, &#x60;POST_ACTION_EXECUTION&#x60;. | 
**functionId** | **string** | The ID of the function. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PublicActionFunction**](PublicActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType

> PublicActionFunction GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType(ctx, definitionId, functionType, appId).Execute()

Retrieve functions by a type for a given definition

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
	definitionId := "definitionId_example" // string | The ID of the definition.
	functionType := "functionType_example" // string | The type of function. Can be `PRE_ACTION_EXECUTION`, `PRE_FETCH_OPTIONS`, `POST_FETCH_OPTIONS`, `POST_ACTION_EXECUTION`.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType(context.Background(), definitionId, functionType, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType`: PublicActionFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the definition. | 
**functionType** | **string** | The type of function. Can be &#x60;PRE_ACTION_EXECUTION&#x60;, &#x60;PRE_FETCH_OPTIONS&#x60;, &#x60;POST_FETCH_OPTIONS&#x60;, &#x60;POST_ACTION_EXECUTION&#x60;. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PublicActionFunction**](PublicActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage

> CollectionResponsePublicActionFunctionIdentifierNoPaging GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage(ctx, definitionId, appId).Execute()

Retrieve functions for a given definition



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
	definitionId := "definitionId_example" // string | The ID of the definition.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage(context.Background(), definitionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage`: CollectionResponsePublicActionFunctionIdentifierNoPaging
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the definition. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicActionFunctionIdentifierNoPaging**](CollectionResponsePublicActionFunctionIdentifierNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType

> PublicActionFunctionIdentifier PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType(ctx, definitionId, functionType, appId).Body(body).Execute()

Insert a function for a definition



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
	definitionId := "definitionId_example" // string | The ID of the definition.
	functionType := "functionType_example" // string | The type of function. Can be `PRE_ACTION_EXECUTION`, `PRE_FETCH_OPTIONS`, `POST_FETCH_OPTIONS`, `POST_ACTION_EXECUTION`.
	appId := int32(56) // int32 | The ID of the app.
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType(context.Background(), definitionId, functionType, appId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType`: PublicActionFunctionIdentifier
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the definition. | 
**functionType** | **string** | The type of function. Can be &#x60;PRE_ACTION_EXECUTION&#x60;, &#x60;PRE_FETCH_OPTIONS&#x60;, &#x60;POST_FETCH_OPTIONS&#x60;, &#x60;POST_ACTION_EXECUTION&#x60;. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

### Return type

[**PublicActionFunctionIdentifier**](PublicActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace

> PublicActionFunctionIdentifier PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace(ctx, definitionId, functionType, functionId, appId).Body(body).Execute()

Update a function for a definition



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
	definitionId := "definitionId_example" // string | The ID of the definition.
	functionType := "functionType_example" // string | The type of function. Can be `PRE_ACTION_EXECUTION`, `PRE_FETCH_OPTIONS`, `POST_FETCH_OPTIONS`, `POST_ACTION_EXECUTION`.
	functionId := "functionId_example" // string | The ID of the function.
	appId := int32(56) // int32 | The ID of the app.
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace(context.Background(), definitionId, functionType, functionId, appId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace`: PublicActionFunctionIdentifier
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the definition. | 
**functionType** | **string** | The type of function. Can be &#x60;PRE_ACTION_EXECUTION&#x60;, &#x60;PRE_FETCH_OPTIONS&#x60;, &#x60;POST_FETCH_OPTIONS&#x60;, &#x60;POST_ACTION_EXECUTION&#x60;. | 
**functionId** | **string** | The ID of the function. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**PublicActionFunctionIdentifier**](PublicActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

