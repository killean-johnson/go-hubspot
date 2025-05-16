# \RevisionsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPage**](RevisionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPage) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions | Retrieve revisions for a given definition
[**GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetById**](RevisionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetById) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions/{revisionId} | Retrieve a specific revision of a definition



## GetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPage

> CollectionResponsePublicActionRevisionForwardPaging GetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPage(ctx, definitionId, appId).Limit(limit).After(after).Execute()

Retrieve revisions for a given definition



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
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPage(context.Background(), definitionId, appId).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPage`: CollectionResponsePublicActionRevisionForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the definition. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The maximum number of results to display per page. | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 

### Return type

[**CollectionResponsePublicActionRevisionForwardPaging**](CollectionResponsePublicActionRevisionForwardPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetById

> PublicActionRevision GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetById(ctx, definitionId, revisionId, appId).Execute()

Retrieve a specific revision of a definition



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
	revisionId := "revisionId_example" // string | The ID of the revision.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetById(context.Background(), definitionId, revisionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetById`: PublicActionRevision
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the definition. | 
**revisionId** | **string** | The ID of the revision. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PublicActionRevision**](PublicActionRevision.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

