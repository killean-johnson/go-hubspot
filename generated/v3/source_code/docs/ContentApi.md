# \ContentAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3SourceCodeEnvironmentContentPathArchive**](ContentAPI.md#DeleteCmsV3SourceCodeEnvironmentContentPathArchive) | **Delete** /cms/v3/source-code/{environment}/content/{path} | Delete a file
[**GetCmsV3SourceCodeEnvironmentContentPathDownload**](ContentAPI.md#GetCmsV3SourceCodeEnvironmentContentPathDownload) | **Get** /cms/v3/source-code/{environment}/content/{path} | Download a file
[**PostCmsV3SourceCodeEnvironmentContentPathCreate**](ContentAPI.md#PostCmsV3SourceCodeEnvironmentContentPathCreate) | **Post** /cms/v3/source-code/{environment}/content/{path} | Create a file
[**PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate**](ContentAPI.md#PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate) | **Put** /cms/v3/source-code/{environment}/content/{path} | Create or update a file



## DeleteCmsV3SourceCodeEnvironmentContentPathArchive

> DeleteCmsV3SourceCodeEnvironmentContentPathArchive(ctx, environment, path).Execute()

Delete a file



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
	environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
	path := "path_example" // string | The file system location of the file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContentAPI.DeleteCmsV3SourceCodeEnvironmentContentPathArchive(context.Background(), environment, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.DeleteCmsV3SourceCodeEnvironmentContentPathArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3SourceCodeEnvironmentContentPathArchiveRequest struct via the builder pattern


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


## GetCmsV3SourceCodeEnvironmentContentPathDownload

> Error GetCmsV3SourceCodeEnvironmentContentPathDownload(ctx, environment, path).Execute()

Download a file



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
	environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
	path := "path_example" // string | The file system location of the file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetCmsV3SourceCodeEnvironmentContentPathDownload(context.Background(), environment, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetCmsV3SourceCodeEnvironmentContentPathDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3SourceCodeEnvironmentContentPathDownload`: Error
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetCmsV3SourceCodeEnvironmentContentPathDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3SourceCodeEnvironmentContentPathDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Error**](Error.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3SourceCodeEnvironmentContentPathCreate

> AssetFileMetadata PostCmsV3SourceCodeEnvironmentContentPathCreate(ctx, environment, path).File(file).Execute()

Create a file



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
	environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
	path := "path_example" // string | The file system location of the file.
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.PostCmsV3SourceCodeEnvironmentContentPathCreate(context.Background(), environment, path).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.PostCmsV3SourceCodeEnvironmentContentPathCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3SourceCodeEnvironmentContentPathCreate`: AssetFileMetadata
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.PostCmsV3SourceCodeEnvironmentContentPathCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3SourceCodeEnvironmentContentPathCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate

> AssetFileMetadata PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate(ctx, environment, path).File(file).Execute()

Create or update a file



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
	environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
	path := "path_example" // string | The file system location of the file.
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate(context.Background(), environment, path).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate`: AssetFileMetadata
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

