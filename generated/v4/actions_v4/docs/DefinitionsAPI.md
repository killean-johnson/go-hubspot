# \DefinitionsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutomationV4ActionsAppIdDefinitionIdArchive**](DefinitionsAPI.md#DeleteAutomationV4ActionsAppIdDefinitionIdArchive) | **Delete** /automation/v4/actions/{appId}/{definitionId} | Delete an action definition
[**GetAutomationV4ActionsAppIdDefinitionIdGetById**](DefinitionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdGetById) | **Get** /automation/v4/actions/{appId}/{definitionId} | Retrieve a custom action definition
[**GetAutomationV4ActionsAppIdGetPage**](DefinitionsAPI.md#GetAutomationV4ActionsAppIdGetPage) | **Get** /automation/v4/actions/{appId} | Retrieve custom action definitions
[**PatchAutomationV4ActionsAppIdDefinitionIdUpdate**](DefinitionsAPI.md#PatchAutomationV4ActionsAppIdDefinitionIdUpdate) | **Patch** /automation/v4/actions/{appId}/{definitionId} | Update an existing action definition
[**PostAutomationV4ActionsAppIdCreate**](DefinitionsAPI.md#PostAutomationV4ActionsAppIdCreate) | **Post** /automation/v4/actions/{appId} | Create a new custom action definition



## DeleteAutomationV4ActionsAppIdDefinitionIdArchive

> DeleteAutomationV4ActionsAppIdDefinitionIdArchive(ctx, definitionId, appId).Execute()

Delete an action definition



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
	definitionId := "definitionId_example" // string | The ID of the custom action definition.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefinitionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdArchive(context.Background(), definitionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom action definition. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationV4ActionsAppIdDefinitionIdArchiveRequest struct via the builder pattern


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


## GetAutomationV4ActionsAppIdDefinitionIdGetById

> PublicActionDefinition GetAutomationV4ActionsAppIdDefinitionIdGetById(ctx, definitionId, appId).Archived(archived).Execute()

Retrieve a custom action definition



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
	definitionId := "definitionId_example" // string | The ID of the custom action.
	appId := int32(56) // int32 | The ID of the app.
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionsAPI.GetAutomationV4ActionsAppIdDefinitionIdGetById(context.Background(), definitionId, appId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsAPI.GetAutomationV4ActionsAppIdDefinitionIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdGetById`: PublicActionDefinition
	fmt.Fprintf(os.Stdout, "Response from `DefinitionsAPI.GetAutomationV4ActionsAppIdDefinitionIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom action. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**PublicActionDefinition**](PublicActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppIdGetPage

> CollectionResponsePublicActionDefinitionForwardPaging GetAutomationV4ActionsAppIdGetPage(ctx, appId).Limit(limit).After(after).Archived(archived).Execute()

Retrieve custom action definitions



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
	appId := int32(56) // int32 | The ID of the app.
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionsAPI.GetAutomationV4ActionsAppIdGetPage(context.Background(), appId).Limit(limit).After(after).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsAPI.GetAutomationV4ActionsAppIdGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdGetPage`: CollectionResponsePublicActionDefinitionForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `DefinitionsAPI.GetAutomationV4ActionsAppIdGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of results to display per page. | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponsePublicActionDefinitionForwardPaging**](CollectionResponsePublicActionDefinitionForwardPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAutomationV4ActionsAppIdDefinitionIdUpdate

> PublicActionDefinition PatchAutomationV4ActionsAppIdDefinitionIdUpdate(ctx, definitionId, appId).PublicActionDefinitionPatch(publicActionDefinitionPatch).Execute()

Update an existing action definition



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
	definitionId := "definitionId_example" // string | The ID of the custom action definition.
	appId := int32(56) // int32 | The ID of the app.
	publicActionDefinitionPatch := *openapiclient.NewPublicActionDefinitionPatch() // PublicActionDefinitionPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionsAPI.PatchAutomationV4ActionsAppIdDefinitionIdUpdate(context.Background(), definitionId, appId).PublicActionDefinitionPatch(publicActionDefinitionPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsAPI.PatchAutomationV4ActionsAppIdDefinitionIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAutomationV4ActionsAppIdDefinitionIdUpdate`: PublicActionDefinition
	fmt.Fprintf(os.Stdout, "Response from `DefinitionsAPI.PatchAutomationV4ActionsAppIdDefinitionIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom action definition. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAutomationV4ActionsAppIdDefinitionIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicActionDefinitionPatch** | [**PublicActionDefinitionPatch**](PublicActionDefinitionPatch.md) |  | 

### Return type

[**PublicActionDefinition**](PublicActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutomationV4ActionsAppIdCreate

> PublicActionDefinition PostAutomationV4ActionsAppIdCreate(ctx, appId).PublicActionDefinitionEgg(publicActionDefinitionEgg).Execute()

Create a new custom action definition



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
	appId := int32(56) // int32 | The ID of the app.
	publicActionDefinitionEgg := *openapiclient.NewPublicActionDefinitionEgg([]openapiclient.InputFieldDefinition{*openapiclient.NewInputFieldDefinition(false, *openapiclient.NewFieldTypeDefinition("Name_example", []openapiclient.Option{*openapiclient.NewOption(false, int32(123), float32(123), "Description_example", false, "Label_example", "Value_example")}, "Type_example", false))}, []openapiclient.PublicActionFunction{*openapiclient.NewPublicActionFunction("FunctionSource_example", "FunctionType_example")}, "ActionUrl_example", false, []string{"ObjectTypes_example"}, map[string]PublicActionLabels{"key": *openapiclient.NewPublicActionLabels("ActionName_example")}) // PublicActionDefinitionEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionsAPI.PostAutomationV4ActionsAppIdCreate(context.Background(), appId).PublicActionDefinitionEgg(publicActionDefinitionEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsAPI.PostAutomationV4ActionsAppIdCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationV4ActionsAppIdCreate`: PublicActionDefinition
	fmt.Fprintf(os.Stdout, "Response from `DefinitionsAPI.PostAutomationV4ActionsAppIdCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationV4ActionsAppIdCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicActionDefinitionEgg** | [**PublicActionDefinitionEgg**](PublicActionDefinitionEgg.md) |  | 

### Return type

[**PublicActionDefinition**](PublicActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

