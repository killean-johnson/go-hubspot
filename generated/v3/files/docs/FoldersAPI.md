# \FoldersAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFilesV3FoldersFolderIdArchive**](FoldersAPI.md#DeleteFilesV3FoldersFolderIdArchive) | **Delete** /files/v3/folders/{folderId} | Delete folder by ID
[**DeleteFilesV3FoldersFolderPathArchiveByPath**](FoldersAPI.md#DeleteFilesV3FoldersFolderPathArchiveByPath) | **Delete** /files/v3/folders/{folderPath} | Delete folder by path
[**GetFilesV3FoldersFolderIdGetById**](FoldersAPI.md#GetFilesV3FoldersFolderIdGetById) | **Get** /files/v3/folders/{folderId} | Retrieve folder by ID
[**GetFilesV3FoldersFolderPathGetByPath**](FoldersAPI.md#GetFilesV3FoldersFolderPathGetByPath) | **Get** /files/v3/folders/{folderPath} | Retrieve folder by path
[**GetFilesV3FoldersSearchDoSearch**](FoldersAPI.md#GetFilesV3FoldersSearchDoSearch) | **Get** /files/v3/folders/search | Search folders
[**GetFilesV3FoldersUpdateAsyncTasksTaskIdStatusCheckUpdateStatus**](FoldersAPI.md#GetFilesV3FoldersUpdateAsyncTasksTaskIdStatusCheckUpdateStatus) | **Get** /files/v3/folders/update/async/tasks/{taskId}/status | Check folder update status
[**PatchFilesV3FoldersFolderIdUpdateProperties**](FoldersAPI.md#PatchFilesV3FoldersFolderIdUpdateProperties) | **Patch** /files/v3/folders/{folderId} | Update folder properties by folder ID
[**PostFilesV3FoldersCreate**](FoldersAPI.md#PostFilesV3FoldersCreate) | **Post** /files/v3/folders | Create folder
[**PostFilesV3FoldersUpdateAsyncUpdatePropertiesRecursively**](FoldersAPI.md#PostFilesV3FoldersUpdateAsyncUpdatePropertiesRecursively) | **Post** /files/v3/folders/update/async | Update folder properties



## DeleteFilesV3FoldersFolderIdArchive

> DeleteFilesV3FoldersFolderIdArchive(ctx, folderId).Execute()

Delete folder by ID



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
	folderId := "folderId_example" // string | ID of folder to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FoldersAPI.DeleteFilesV3FoldersFolderIdArchive(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.DeleteFilesV3FoldersFolderIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | ID of folder to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilesV3FoldersFolderIdArchiveRequest struct via the builder pattern


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


## DeleteFilesV3FoldersFolderPathArchiveByPath

> DeleteFilesV3FoldersFolderPathArchiveByPath(ctx, folderPath).Execute()

Delete folder by path



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
	folderPath := "folderPath_example" // string | Path of folder to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FoldersAPI.DeleteFilesV3FoldersFolderPathArchiveByPath(context.Background(), folderPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.DeleteFilesV3FoldersFolderPathArchiveByPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderPath** | **string** | Path of folder to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilesV3FoldersFolderPathArchiveByPathRequest struct via the builder pattern


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


## GetFilesV3FoldersFolderIdGetById

> Folder GetFilesV3FoldersFolderIdGetById(ctx, folderId).Properties(properties).Execute()

Retrieve folder by ID



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
	folderId := "folderId_example" // string | ID of desired folder
	properties := []string{"Inner_example"} // []string | Properties to set on returned folder. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFilesV3FoldersFolderIdGetById(context.Background(), folderId).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFilesV3FoldersFolderIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesV3FoldersFolderIdGetById`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFilesV3FoldersFolderIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | ID of desired folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesV3FoldersFolderIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | **[]string** | Properties to set on returned folder. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesV3FoldersFolderPathGetByPath

> Folder GetFilesV3FoldersFolderPathGetByPath(ctx, folderPath).Properties(properties).Execute()

Retrieve folder by path



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
	folderPath := "folderPath_example" // string | Path of desired folder.
	properties := []string{"Inner_example"} // []string | Properties to set on returned folder. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFilesV3FoldersFolderPathGetByPath(context.Background(), folderPath).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFilesV3FoldersFolderPathGetByPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesV3FoldersFolderPathGetByPath`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFilesV3FoldersFolderPathGetByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderPath** | **string** | Path of desired folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesV3FoldersFolderPathGetByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | **[]string** | Properties to set on returned folder. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesV3FoldersSearchDoSearch

> CollectionResponseFolder GetFilesV3FoldersSearchDoSearch(ctx).Properties(properties).After(after).Before(before).Limit(limit).Sort(sort).Ids(ids).IdLte(idLte).IdGte(idGte).CreatedAt(createdAt).CreatedAtLte(createdAtLte).CreatedAtGte(createdAtGte).UpdatedAt(updatedAt).UpdatedAtLte(updatedAtLte).UpdatedAtGte(updatedAtGte).Name(name).Path(path).ParentFolderIds(parentFolderIds).Execute()

Search folders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	properties := []string{"Inner_example"} // []string | Properties that should be included in the returned folders. (optional)
	after := "after_example" // string | Offset search results by this value. The default offset is 0 and the maximum offset of items for a given search is 10,000. Narrow your search down if you are reaching this limit. (optional)
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | Number of items to return. Default limit is 10, maximum limit is 100. (optional)
	sort := []string{"Inner_example"} // []string | Sort results by given property. For example -name sorts by name field descending, name sorts by name field ascending. (optional)
	ids := []int64{int64(123)} // []int64 |  (optional)
	idLte := int64(789) // int64 |  (optional)
	idGte := int64(789) // int64 |  (optional)
	createdAt := time.Now() // time.Time | Search folders by exact time of creation. Time must be epoch time in milliseconds. (optional)
	createdAtLte := time.Now() // time.Time | Search folders by less than or equal to time of creation. Can be used with createdAtGte to create a range. (optional)
	createdAtGte := time.Now() // time.Time | Search folders by greater than or equal to time of creation. Can be used with createdAtLte to create a range. (optional)
	updatedAt := time.Now() // time.Time | Search folders by exact time of latest updated. Time must be epoch time in milliseconds. (optional)
	updatedAtLte := time.Now() // time.Time | Search folders by less than or equal to time of latest update. Can be used with updatedAtGte to create a range. (optional)
	updatedAtGte := time.Now() // time.Time | Search folders by greater than or equal to time of latest update. Can be used with updatedAtLte to create a range. (optional)
	name := "name_example" // string | Search for folders containing the specified name. (optional)
	path := "path_example" // string | Search folders by path. (optional)
	parentFolderIds := []int64{int64(123)} // []int64 | Search folders with the given parent folderId. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFilesV3FoldersSearchDoSearch(context.Background()).Properties(properties).After(after).Before(before).Limit(limit).Sort(sort).Ids(ids).IdLte(idLte).IdGte(idGte).CreatedAt(createdAt).CreatedAtLte(createdAtLte).CreatedAtGte(createdAtGte).UpdatedAt(updatedAt).UpdatedAtLte(updatedAtLte).UpdatedAtGte(updatedAtGte).Name(name).Path(path).ParentFolderIds(parentFolderIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFilesV3FoldersSearchDoSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesV3FoldersSearchDoSearch`: CollectionResponseFolder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFilesV3FoldersSearchDoSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesV3FoldersSearchDoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **properties** | **[]string** | Properties that should be included in the returned folders. | 
 **after** | **string** | Offset search results by this value. The default offset is 0 and the maximum offset of items for a given search is 10,000. Narrow your search down if you are reaching this limit. | 
 **before** | **string** |  | 
 **limit** | **int32** | Number of items to return. Default limit is 10, maximum limit is 100. | 
 **sort** | **[]string** | Sort results by given property. For example -name sorts by name field descending, name sorts by name field ascending. | 
 **ids** | **[]int64** |  | 
 **idLte** | **int64** |  | 
 **idGte** | **int64** |  | 
 **createdAt** | **time.Time** | Search folders by exact time of creation. Time must be epoch time in milliseconds. | 
 **createdAtLte** | **time.Time** | Search folders by less than or equal to time of creation. Can be used with createdAtGte to create a range. | 
 **createdAtGte** | **time.Time** | Search folders by greater than or equal to time of creation. Can be used with createdAtLte to create a range. | 
 **updatedAt** | **time.Time** | Search folders by exact time of latest updated. Time must be epoch time in milliseconds. | 
 **updatedAtLte** | **time.Time** | Search folders by less than or equal to time of latest update. Can be used with updatedAtGte to create a range. | 
 **updatedAtGte** | **time.Time** | Search folders by greater than or equal to time of latest update. Can be used with updatedAtLte to create a range. | 
 **name** | **string** | Search for folders containing the specified name. | 
 **path** | **string** | Search folders by path. | 
 **parentFolderIds** | **[]int64** | Search folders with the given parent folderId. | 

### Return type

[**CollectionResponseFolder**](CollectionResponseFolder.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesV3FoldersUpdateAsyncTasksTaskIdStatusCheckUpdateStatus

> FolderActionResponse GetFilesV3FoldersUpdateAsyncTasksTaskIdStatusCheckUpdateStatus(ctx, taskId).Execute()

Check folder update status



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
	taskId := "taskId_example" // string | The ID of the folder update task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFilesV3FoldersUpdateAsyncTasksTaskIdStatusCheckUpdateStatus(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFilesV3FoldersUpdateAsyncTasksTaskIdStatusCheckUpdateStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesV3FoldersUpdateAsyncTasksTaskIdStatusCheckUpdateStatus`: FolderActionResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFilesV3FoldersUpdateAsyncTasksTaskIdStatusCheckUpdateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the folder update task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesV3FoldersUpdateAsyncTasksTaskIdStatusCheckUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderActionResponse**](FolderActionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFilesV3FoldersFolderIdUpdateProperties

> Folder PatchFilesV3FoldersFolderIdUpdateProperties(ctx, folderId).FolderUpdateInput(folderUpdateInput).Execute()

Update folder properties by folder ID



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
	folderId := "folderId_example" // string | ID of folder to update
	folderUpdateInput := *openapiclient.NewFolderUpdateInput() // FolderUpdateInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.PatchFilesV3FoldersFolderIdUpdateProperties(context.Background(), folderId).FolderUpdateInput(folderUpdateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.PatchFilesV3FoldersFolderIdUpdateProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFilesV3FoldersFolderIdUpdateProperties`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.PatchFilesV3FoldersFolderIdUpdateProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | ID of folder to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFilesV3FoldersFolderIdUpdatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderUpdateInput** | [**FolderUpdateInput**](FolderUpdateInput.md) |  | 

### Return type

[**Folder**](Folder.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFilesV3FoldersCreate

> Folder PostFilesV3FoldersCreate(ctx).FolderInput(folderInput).Execute()

Create folder



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
	folderInput := *openapiclient.NewFolderInput("Name_example") // FolderInput | Folder creation options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.PostFilesV3FoldersCreate(context.Background()).FolderInput(folderInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.PostFilesV3FoldersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFilesV3FoldersCreate`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.PostFilesV3FoldersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFilesV3FoldersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderInput** | [**FolderInput**](FolderInput.md) | Folder creation options | 

### Return type

[**Folder**](Folder.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFilesV3FoldersUpdateAsyncUpdatePropertiesRecursively

> FolderUpdateTaskLocator PostFilesV3FoldersUpdateAsyncUpdatePropertiesRecursively(ctx).FolderUpdateInputWithId(folderUpdateInputWithId).Execute()

Update folder properties



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
	folderUpdateInputWithId := *openapiclient.NewFolderUpdateInputWithId("Id_example") // FolderUpdateInputWithId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.PostFilesV3FoldersUpdateAsyncUpdatePropertiesRecursively(context.Background()).FolderUpdateInputWithId(folderUpdateInputWithId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.PostFilesV3FoldersUpdateAsyncUpdatePropertiesRecursively``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFilesV3FoldersUpdateAsyncUpdatePropertiesRecursively`: FolderUpdateTaskLocator
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.PostFilesV3FoldersUpdateAsyncUpdatePropertiesRecursively`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFilesV3FoldersUpdateAsyncUpdatePropertiesRecursivelyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderUpdateInputWithId** | [**FolderUpdateInputWithId**](FolderUpdateInputWithId.md) |  | 

### Return type

[**FolderUpdateTaskLocator**](FolderUpdateTaskLocator.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

