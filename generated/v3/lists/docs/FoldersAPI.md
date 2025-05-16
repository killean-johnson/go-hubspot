# \FoldersAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ListsFoldersFolderIdRemove**](FoldersAPI.md#DeleteCrmV3ListsFoldersFolderIdRemove) | **Delete** /crm/v3/lists/folders/{folderId} | Deletes a folder
[**GetCrmV3ListsFoldersGetAll**](FoldersAPI.md#GetCrmV3ListsFoldersGetAll) | **Get** /crm/v3/lists/folders | Retrieves a folder.
[**PostCrmV3ListsFoldersCreate**](FoldersAPI.md#PostCrmV3ListsFoldersCreate) | **Post** /crm/v3/lists/folders | Creates a folder
[**PutCrmV3ListsFoldersFolderIdMoveNewParentFolderIdMove**](FoldersAPI.md#PutCrmV3ListsFoldersFolderIdMoveNewParentFolderIdMove) | **Put** /crm/v3/lists/folders/{folderId}/move/{newParentFolderId} | Moves a folder
[**PutCrmV3ListsFoldersFolderIdRenameRename**](FoldersAPI.md#PutCrmV3ListsFoldersFolderIdRenameRename) | **Put** /crm/v3/lists/folders/{folderId}/rename | Rename a folder
[**PutCrmV3ListsFoldersMoveListMoveList**](FoldersAPI.md#PutCrmV3ListsFoldersMoveListMoveList) | **Put** /crm/v3/lists/folders/move-list | Moves a list to a given folder



## DeleteCrmV3ListsFoldersFolderIdRemove

> DeleteCrmV3ListsFoldersFolderIdRemove(ctx, folderId).Execute()

Deletes a folder



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
	folderId := "folderId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FoldersAPI.DeleteCrmV3ListsFoldersFolderIdRemove(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.DeleteCrmV3ListsFoldersFolderIdRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ListsFoldersFolderIdRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ListsFoldersGetAll

> ListFolderFetchResponse GetCrmV3ListsFoldersGetAll(ctx).FolderId(folderId).Execute()

Retrieves a folder.



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
	folderId := "folderId_example" // string | The Id of the folder to retrieve. (optional) (default to "0")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetCrmV3ListsFoldersGetAll(context.Background()).FolderId(folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetCrmV3ListsFoldersGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ListsFoldersGetAll`: ListFolderFetchResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetCrmV3ListsFoldersGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ListsFoldersGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string** | The Id of the folder to retrieve. | [default to &quot;0&quot;]

### Return type

[**ListFolderFetchResponse**](ListFolderFetchResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ListsFoldersCreate

> ListFolderCreateResponse PostCrmV3ListsFoldersCreate(ctx).ListFolderCreateRequest(listFolderCreateRequest).Execute()

Creates a folder



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
	listFolderCreateRequest := *openapiclient.NewListFolderCreateRequest("Name_example") // ListFolderCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.PostCrmV3ListsFoldersCreate(context.Background()).ListFolderCreateRequest(listFolderCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.PostCrmV3ListsFoldersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ListsFoldersCreate`: ListFolderCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.PostCrmV3ListsFoldersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ListsFoldersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listFolderCreateRequest** | [**ListFolderCreateRequest**](ListFolderCreateRequest.md) |  | 

### Return type

[**ListFolderCreateResponse**](ListFolderCreateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ListsFoldersFolderIdMoveNewParentFolderIdMove

> ListFolderFetchResponse PutCrmV3ListsFoldersFolderIdMoveNewParentFolderIdMove(ctx, folderId, newParentFolderId).Execute()

Moves a folder



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
	folderId := "folderId_example" // string | 
	newParentFolderId := "newParentFolderId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.PutCrmV3ListsFoldersFolderIdMoveNewParentFolderIdMove(context.Background(), folderId, newParentFolderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.PutCrmV3ListsFoldersFolderIdMoveNewParentFolderIdMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3ListsFoldersFolderIdMoveNewParentFolderIdMove`: ListFolderFetchResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.PutCrmV3ListsFoldersFolderIdMoveNewParentFolderIdMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 
**newParentFolderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ListsFoldersFolderIdMoveNewParentFolderIdMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListFolderFetchResponse**](ListFolderFetchResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ListsFoldersFolderIdRenameRename

> ListFolderFetchResponse PutCrmV3ListsFoldersFolderIdRenameRename(ctx, folderId).NewFolderName(newFolderName).Execute()

Rename a folder



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
	folderId := "folderId_example" // string | 
	newFolderName := "newFolderName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.PutCrmV3ListsFoldersFolderIdRenameRename(context.Background(), folderId).NewFolderName(newFolderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.PutCrmV3ListsFoldersFolderIdRenameRename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3ListsFoldersFolderIdRenameRename`: ListFolderFetchResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.PutCrmV3ListsFoldersFolderIdRenameRename`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ListsFoldersFolderIdRenameRenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newFolderName** | **string** |  | 

### Return type

[**ListFolderFetchResponse**](ListFolderFetchResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ListsFoldersMoveListMoveList

> PutCrmV3ListsFoldersMoveListMoveList(ctx).ListMoveRequest(listMoveRequest).Execute()

Moves a list to a given folder



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
	listMoveRequest := *openapiclient.NewListMoveRequest("ListId_example", "NewFolderId_example") // ListMoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FoldersAPI.PutCrmV3ListsFoldersMoveListMoveList(context.Background()).ListMoveRequest(listMoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.PutCrmV3ListsFoldersMoveListMoveList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ListsFoldersMoveListMoveListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listMoveRequest** | [**ListMoveRequest**](ListMoveRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

