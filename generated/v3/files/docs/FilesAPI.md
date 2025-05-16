# \FilesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFilesV3FilesFileIdArchive**](FilesAPI.md#DeleteFilesV3FilesFileIdArchive) | **Delete** /files/v3/files/{fileId} | Delete file by ID
[**DeleteFilesV3FilesFileIdGdprDeleteDelete**](FilesAPI.md#DeleteFilesV3FilesFileIdGdprDeleteDelete) | **Delete** /files/v3/files/{fileId}/gdpr-delete | GDPR-delete file
[**GetFilesV3FilesFileIdGetById**](FilesAPI.md#GetFilesV3FilesFileIdGetById) | **Get** /files/v3/files/{fileId} | Retrieve file by ID
[**GetFilesV3FilesFileIdSignedUrlGetSignedUrl**](FilesAPI.md#GetFilesV3FilesFileIdSignedUrlGetSignedUrl) | **Get** /files/v3/files/{fileId}/signed-url | Get signed URL to access private file
[**GetFilesV3FilesImportFromUrlAsyncTasksTaskIdStatusCheckImport**](FilesAPI.md#GetFilesV3FilesImportFromUrlAsyncTasksTaskIdStatusCheckImport) | **Get** /files/v3/files/import-from-url/async/tasks/{taskId}/status | Check import status
[**GetFilesV3FilesSearchDoSearch**](FilesAPI.md#GetFilesV3FilesSearchDoSearch) | **Get** /files/v3/files/search | Search files
[**GetFilesV3FilesStatPathGetMetadata**](FilesAPI.md#GetFilesV3FilesStatPathGetMetadata) | **Get** /files/v3/files/stat/{path} | Retrieve file by path
[**PatchFilesV3FilesFileIdUpdateProperties**](FilesAPI.md#PatchFilesV3FilesFileIdUpdateProperties) | **Patch** /files/v3/files/{fileId} | Update file properties
[**PostFilesV3FilesImportFromUrlAsyncImportFromUrl**](FilesAPI.md#PostFilesV3FilesImportFromUrlAsyncImportFromUrl) | **Post** /files/v3/files/import-from-url/async | Import file from URL
[**PostFilesV3FilesUpload**](FilesAPI.md#PostFilesV3FilesUpload) | **Post** /files/v3/files | Upload file
[**PutFilesV3FilesFileIdReplace**](FilesAPI.md#PutFilesV3FilesFileIdReplace) | **Put** /files/v3/files/{fileId} | Replace file



## DeleteFilesV3FilesFileIdArchive

> DeleteFilesV3FilesFileIdArchive(ctx, fileId).Execute()

Delete file by ID



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
	fileId := "fileId_example" // string | FileId to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesAPI.DeleteFilesV3FilesFileIdArchive(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteFilesV3FilesFileIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | FileId to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilesV3FilesFileIdArchiveRequest struct via the builder pattern


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


## DeleteFilesV3FilesFileIdGdprDeleteDelete

> DeleteFilesV3FilesFileIdGdprDeleteDelete(ctx, fileId).Execute()

GDPR-delete file



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
	fileId := "fileId_example" // string | ID of file to GDPR delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesAPI.DeleteFilesV3FilesFileIdGdprDeleteDelete(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteFilesV3FilesFileIdGdprDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ID of file to GDPR delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilesV3FilesFileIdGdprDeleteDeleteRequest struct via the builder pattern


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


## GetFilesV3FilesFileIdGetById

> File GetFilesV3FilesFileIdGetById(ctx, fileId).Properties(properties).Execute()

Retrieve file by ID



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
	fileId := "fileId_example" // string | ID of the desired file.
	properties := []string{"Inner_example"} // []string | null (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFilesV3FilesFileIdGetById(context.Background(), fileId).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFilesV3FilesFileIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesV3FilesFileIdGetById`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFilesV3FilesFileIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ID of the desired file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesV3FilesFileIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | **[]string** | null | 

### Return type

[**File**](File.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesV3FilesFileIdSignedUrlGetSignedUrl

> SignedUrl GetFilesV3FilesFileIdSignedUrlGetSignedUrl(ctx, fileId).Size(size).ExpirationSeconds(expirationSeconds).Upscale(upscale).Execute()

Get signed URL to access private file



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
	fileId := "fileId_example" // string | ID of file.
	size := "size_example" // string | For image files. This will resize the image to the desired size before sharing. Does not affect the original file, just the file served by this signed URL. (optional)
	expirationSeconds := int64(789) // int64 | How long in seconds the link will provide access to the file. (optional)
	upscale := true // bool | If size is provided, this will upscale the image to fit the size dimensions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFilesV3FilesFileIdSignedUrlGetSignedUrl(context.Background(), fileId).Size(size).ExpirationSeconds(expirationSeconds).Upscale(upscale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFilesV3FilesFileIdSignedUrlGetSignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesV3FilesFileIdSignedUrlGetSignedUrl`: SignedUrl
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFilesV3FilesFileIdSignedUrlGetSignedUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ID of file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesV3FilesFileIdSignedUrlGetSignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **string** | For image files. This will resize the image to the desired size before sharing. Does not affect the original file, just the file served by this signed URL. | 
 **expirationSeconds** | **int64** | How long in seconds the link will provide access to the file. | 
 **upscale** | **bool** | If size is provided, this will upscale the image to fit the size dimensions. | 

### Return type

[**SignedUrl**](SignedUrl.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesV3FilesImportFromUrlAsyncTasksTaskIdStatusCheckImport

> FileActionResponse GetFilesV3FilesImportFromUrlAsyncTasksTaskIdStatusCheckImport(ctx, taskId).Execute()

Check import status



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
	taskId := "taskId_example" // string | Import by URL task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFilesV3FilesImportFromUrlAsyncTasksTaskIdStatusCheckImport(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFilesV3FilesImportFromUrlAsyncTasksTaskIdStatusCheckImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesV3FilesImportFromUrlAsyncTasksTaskIdStatusCheckImport`: FileActionResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFilesV3FilesImportFromUrlAsyncTasksTaskIdStatusCheckImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Import by URL task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesV3FilesImportFromUrlAsyncTasksTaskIdStatusCheckImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileActionResponse**](FileActionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesV3FilesSearchDoSearch

> CollectionResponseFile GetFilesV3FilesSearchDoSearch(ctx).Properties(properties).After(after).Before(before).Limit(limit).Sort(sort).Ids(ids).IdLte(idLte).IdGte(idGte).CreatedAt(createdAt).CreatedAtLte(createdAtLte).CreatedAtGte(createdAtGte).UpdatedAt(updatedAt).UpdatedAtLte(updatedAtLte).UpdatedAtGte(updatedAtGte).Name(name).Path(path).ParentFolderIds(parentFolderIds).Size(size).SizeLte(sizeLte).SizeGte(sizeGte).Height(height).HeightLte(heightLte).HeightGte(heightGte).Width(width).WidthLte(widthLte).WidthGte(widthGte).Encoding(encoding).Type_(type_).Extension(extension).Url(url).IsUsableInContent(isUsableInContent).AllowsAnonymousAccess(allowsAnonymousAccess).FileMd5(fileMd5).ExpiresAt(expiresAt).ExpiresAtLte(expiresAtLte).ExpiresAtGte(expiresAtGte).Execute()

Search files



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
	properties := []string{"Inner_example"} // []string | A list of file properties to return. (optional)
	after := "after_example" // string | Offset search results by this value. The default offset is 0 and the maximum offset of items for a given search is 10,000. Narrow your search down if you are reaching this limit. (optional)
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | Number of items to return. Default limit is 10, maximum limit is 100. (optional)
	sort := []string{"Inner_example"} // []string | Sort files by a given field. (optional)
	ids := []int64{int64(123)} // []int64 | Search by a list of file IDs. (optional)
	idLte := int64(789) // int64 |  (optional)
	idGte := int64(789) // int64 |  (optional)
	createdAt := time.Now() // time.Time | Search files by time of creation. (optional)
	createdAtLte := time.Now() // time.Time | Search files by less than or equal to time of creation. Can be used with `createdAtGte` to create a range. (optional)
	createdAtGte := time.Now() // time.Time | Search files by greater than or equal to time of creation. Can be used with `createdAtLte` to create a range. (optional)
	updatedAt := time.Now() // time.Time | Search files by time of latest updated. (optional)
	updatedAtLte := time.Now() // time.Time | Search files by less than or equal to time of latest update. Can be used with `updatedAtGte` to create a range. (optional)
	updatedAtGte := time.Now() // time.Time | Search files by greater than or equal to time of latest update. Can be used with `updatedAtLte` to create a range. (optional)
	name := "name_example" // string | Search for files containing the given name. (optional)
	path := "path_example" // string | Search files by path. (optional)
	parentFolderIds := []int64{int64(123)} // []int64 | Search files within given `folderId`. (optional)
	size := int64(789) // int64 | Search files by exact file size in bytes. (optional)
	sizeLte := int64(789) // int64 | Search files by less than or equal to file size. Can be used with `sizeGte` to create a range. (optional)
	sizeGte := int64(789) // int64 | Search files by greater than or equal to file size. Can be used with `sizeLte` to create a range. (optional)
	height := int32(56) // int32 | Search files by height of image or video. (optional)
	heightLte := int32(56) // int32 | Search files by less than or equal to height of image or video. Can be used with `heightGte` to create a range. (optional)
	heightGte := int32(56) // int32 | Search files by greater than or equal to height of image or video. Can be used with `heightLte` to create a range. (optional)
	width := int32(56) // int32 | Search files by width of image or video. (optional)
	widthLte := int32(56) // int32 | Search files by less than or equal to width of image or video. Can be used with `widthGte` to create a range. (optional)
	widthGte := int32(56) // int32 | Search files by greater than or equal to width of image or video. Can be used with `widthLte` to create a range. (optional)
	encoding := "encoding_example" // string | Search files by specified encoding. (optional)
	type_ := "type__example" // string | Filter by provided file type. (optional)
	extension := "extension_example" // string | Search files by given extension. (optional)
	url := "url_example" // string | Search by file URL. (optional)
	isUsableInContent := true // bool | If `true`, shows files that have been marked to be used in new content. If `false`, shows files that should not be used in new content. (optional)
	allowsAnonymousAccess := true // bool | Search files by access. If `true`, will show only public files. If `false`, will show only private files. (optional)
	fileMd5 := "fileMd5_example" // string | Search files by a specific md5 hash. (optional)
	expiresAt := time.Now() // time.Time | Search files by exact expires time. Time must be epoch time in milliseconds. (optional)
	expiresAtLte := time.Now() // time.Time | Search files by less than or equal to expires time. Can be used with `expiresAtGte` to create a range. (optional)
	expiresAtGte := time.Now() // time.Time | Search files by greater than or equal to expires time. Can be used with `expiresAtLte` to create a range. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFilesV3FilesSearchDoSearch(context.Background()).Properties(properties).After(after).Before(before).Limit(limit).Sort(sort).Ids(ids).IdLte(idLte).IdGte(idGte).CreatedAt(createdAt).CreatedAtLte(createdAtLte).CreatedAtGte(createdAtGte).UpdatedAt(updatedAt).UpdatedAtLte(updatedAtLte).UpdatedAtGte(updatedAtGte).Name(name).Path(path).ParentFolderIds(parentFolderIds).Size(size).SizeLte(sizeLte).SizeGte(sizeGte).Height(height).HeightLte(heightLte).HeightGte(heightGte).Width(width).WidthLte(widthLte).WidthGte(widthGte).Encoding(encoding).Type_(type_).Extension(extension).Url(url).IsUsableInContent(isUsableInContent).AllowsAnonymousAccess(allowsAnonymousAccess).FileMd5(fileMd5).ExpiresAt(expiresAt).ExpiresAtLte(expiresAtLte).ExpiresAtGte(expiresAtGte).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFilesV3FilesSearchDoSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesV3FilesSearchDoSearch`: CollectionResponseFile
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFilesV3FilesSearchDoSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesV3FilesSearchDoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **properties** | **[]string** | A list of file properties to return. | 
 **after** | **string** | Offset search results by this value. The default offset is 0 and the maximum offset of items for a given search is 10,000. Narrow your search down if you are reaching this limit. | 
 **before** | **string** |  | 
 **limit** | **int32** | Number of items to return. Default limit is 10, maximum limit is 100. | 
 **sort** | **[]string** | Sort files by a given field. | 
 **ids** | **[]int64** | Search by a list of file IDs. | 
 **idLte** | **int64** |  | 
 **idGte** | **int64** |  | 
 **createdAt** | **time.Time** | Search files by time of creation. | 
 **createdAtLte** | **time.Time** | Search files by less than or equal to time of creation. Can be used with &#x60;createdAtGte&#x60; to create a range. | 
 **createdAtGte** | **time.Time** | Search files by greater than or equal to time of creation. Can be used with &#x60;createdAtLte&#x60; to create a range. | 
 **updatedAt** | **time.Time** | Search files by time of latest updated. | 
 **updatedAtLte** | **time.Time** | Search files by less than or equal to time of latest update. Can be used with &#x60;updatedAtGte&#x60; to create a range. | 
 **updatedAtGte** | **time.Time** | Search files by greater than or equal to time of latest update. Can be used with &#x60;updatedAtLte&#x60; to create a range. | 
 **name** | **string** | Search for files containing the given name. | 
 **path** | **string** | Search files by path. | 
 **parentFolderIds** | **[]int64** | Search files within given &#x60;folderId&#x60;. | 
 **size** | **int64** | Search files by exact file size in bytes. | 
 **sizeLte** | **int64** | Search files by less than or equal to file size. Can be used with &#x60;sizeGte&#x60; to create a range. | 
 **sizeGte** | **int64** | Search files by greater than or equal to file size. Can be used with &#x60;sizeLte&#x60; to create a range. | 
 **height** | **int32** | Search files by height of image or video. | 
 **heightLte** | **int32** | Search files by less than or equal to height of image or video. Can be used with &#x60;heightGte&#x60; to create a range. | 
 **heightGte** | **int32** | Search files by greater than or equal to height of image or video. Can be used with &#x60;heightLte&#x60; to create a range. | 
 **width** | **int32** | Search files by width of image or video. | 
 **widthLte** | **int32** | Search files by less than or equal to width of image or video. Can be used with &#x60;widthGte&#x60; to create a range. | 
 **widthGte** | **int32** | Search files by greater than or equal to width of image or video. Can be used with &#x60;widthLte&#x60; to create a range. | 
 **encoding** | **string** | Search files by specified encoding. | 
 **type_** | **string** | Filter by provided file type. | 
 **extension** | **string** | Search files by given extension. | 
 **url** | **string** | Search by file URL. | 
 **isUsableInContent** | **bool** | If &#x60;true&#x60;, shows files that have been marked to be used in new content. If &#x60;false&#x60;, shows files that should not be used in new content. | 
 **allowsAnonymousAccess** | **bool** | Search files by access. If &#x60;true&#x60;, will show only public files. If &#x60;false&#x60;, will show only private files. | 
 **fileMd5** | **string** | Search files by a specific md5 hash. | 
 **expiresAt** | **time.Time** | Search files by exact expires time. Time must be epoch time in milliseconds. | 
 **expiresAtLte** | **time.Time** | Search files by less than or equal to expires time. Can be used with &#x60;expiresAtGte&#x60; to create a range. | 
 **expiresAtGte** | **time.Time** | Search files by greater than or equal to expires time. Can be used with &#x60;expiresAtLte&#x60; to create a range. | 

### Return type

[**CollectionResponseFile**](CollectionResponseFile.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesV3FilesStatPathGetMetadata

> FileStat GetFilesV3FilesStatPathGetMetadata(ctx, path).Properties(properties).Execute()

Retrieve file by path



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
	path := "path_example" // string | The path of the file. 
	properties := []string{"Inner_example"} // []string | Properties to return in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFilesV3FilesStatPathGetMetadata(context.Background(), path).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFilesV3FilesStatPathGetMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesV3FilesStatPathGetMetadata`: FileStat
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFilesV3FilesStatPathGetMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path of the file.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesV3FilesStatPathGetMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | **[]string** | Properties to return in the response. | 

### Return type

[**FileStat**](FileStat.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFilesV3FilesFileIdUpdateProperties

> File PatchFilesV3FilesFileIdUpdateProperties(ctx, fileId).FileUpdateInput(fileUpdateInput).Execute()

Update file properties



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
	fileId := "fileId_example" // string | ID of file to update
	fileUpdateInput := *openapiclient.NewFileUpdateInput() // FileUpdateInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.PatchFilesV3FilesFileIdUpdateProperties(context.Background(), fileId).FileUpdateInput(fileUpdateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.PatchFilesV3FilesFileIdUpdateProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFilesV3FilesFileIdUpdateProperties`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.PatchFilesV3FilesFileIdUpdateProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ID of file to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFilesV3FilesFileIdUpdatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileUpdateInput** | [**FileUpdateInput**](FileUpdateInput.md) |  | 

### Return type

[**File**](File.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFilesV3FilesImportFromUrlAsyncImportFromUrl

> ImportFromUrlTaskLocator PostFilesV3FilesImportFromUrlAsyncImportFromUrl(ctx).ImportFromUrlInput(importFromUrlInput).Execute()

Import file from URL



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
	importFromUrlInput := *openapiclient.NewImportFromUrlInput("Access_example", "Url_example") // ImportFromUrlInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.PostFilesV3FilesImportFromUrlAsyncImportFromUrl(context.Background()).ImportFromUrlInput(importFromUrlInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.PostFilesV3FilesImportFromUrlAsyncImportFromUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFilesV3FilesImportFromUrlAsyncImportFromUrl`: ImportFromUrlTaskLocator
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.PostFilesV3FilesImportFromUrlAsyncImportFromUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFilesV3FilesImportFromUrlAsyncImportFromUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importFromUrlInput** | [**ImportFromUrlInput**](ImportFromUrlInput.md) |  | 

### Return type

[**ImportFromUrlTaskLocator**](ImportFromUrlTaskLocator.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFilesV3FilesUpload

> File PostFilesV3FilesUpload(ctx).File(file).FolderId(folderId).FolderPath(folderPath).FileName(fileName).CharsetHunch(charsetHunch).Options(options).Execute()

Upload file



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
	file := os.NewFile(1234, "some_file") // *os.File | File to be uploaded. (optional)
	folderId := "folderId_example" // string | Either 'folderId' or 'folderPath' is required. folderId is the ID of the folder the file will be uploaded to. (optional)
	folderPath := "folderPath_example" // string | Either 'folderPath' or 'folderId' is required. This field represents the destination folder path for the uploaded file. If a path doesn't exist, the system will try to create one. (optional)
	fileName := "fileName_example" // string | Desired name for the uploaded file. (optional)
	charsetHunch := "charsetHunch_example" // string | Character set of the uploaded file. (optional)
	options := "options_example" // string | JSON string representing FileUploadOptions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.PostFilesV3FilesUpload(context.Background()).File(file).FolderId(folderId).FolderPath(folderPath).FileName(fileName).CharsetHunch(charsetHunch).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.PostFilesV3FilesUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFilesV3FilesUpload`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.PostFilesV3FilesUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFilesV3FilesUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | File to be uploaded. | 
 **folderId** | **string** | Either &#39;folderId&#39; or &#39;folderPath&#39; is required. folderId is the ID of the folder the file will be uploaded to. | 
 **folderPath** | **string** | Either &#39;folderPath&#39; or &#39;folderId&#39; is required. This field represents the destination folder path for the uploaded file. If a path doesn&#39;t exist, the system will try to create one. | 
 **fileName** | **string** | Desired name for the uploaded file. | 
 **charsetHunch** | **string** | Character set of the uploaded file. | 
 **options** | **string** | JSON string representing FileUploadOptions. | 

### Return type

[**File**](File.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFilesV3FilesFileIdReplace

> File PutFilesV3FilesFileIdReplace(ctx, fileId).File(file).CharsetHunch(charsetHunch).Options(options).Execute()

Replace file



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
	fileId := "fileId_example" // string | ID of the desired file.
	file := os.NewFile(1234, "some_file") // *os.File | File data that will replace existing file in the file manager. (optional)
	charsetHunch := "charsetHunch_example" // string | Character set of given file data. (optional)
	options := "options_example" // string | JSON string representing FileReplaceOptions. Includes options to set the access and expiresAt properties, which will automatically update when the file is replaced. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.PutFilesV3FilesFileIdReplace(context.Background(), fileId).File(file).CharsetHunch(charsetHunch).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.PutFilesV3FilesFileIdReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFilesV3FilesFileIdReplace`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.PutFilesV3FilesFileIdReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ID of the desired file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFilesV3FilesFileIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | File data that will replace existing file in the file manager. | 
 **charsetHunch** | **string** | Character set of given file data. | 
 **options** | **string** | JSON string representing FileReplaceOptions. Includes options to set the access and expiresAt properties, which will automatically update when the file is replaced. | 

### Return type

[**File**](File.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

