# \MetadataAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsV3SourceCodeEnvironmentMetadataPathGet**](MetadataAPI.md#GetCmsV3SourceCodeEnvironmentMetadataPathGet) | **Get** /cms/v3/source-code/{environment}/metadata/{path} | Get the metadata for a file



## GetCmsV3SourceCodeEnvironmentMetadataPathGet

> AssetFileMetadata GetCmsV3SourceCodeEnvironmentMetadataPathGet(ctx, environment, path).Properties(properties).Execute()

Get the metadata for a file



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
	properties := "properties_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.GetCmsV3SourceCodeEnvironmentMetadataPathGet(context.Background(), environment, path).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetCmsV3SourceCodeEnvironmentMetadataPathGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3SourceCodeEnvironmentMetadataPathGet`: AssetFileMetadata
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetCmsV3SourceCodeEnvironmentMetadataPathGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3SourceCodeEnvironmentMetadataPathGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **properties** | **string** |  | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

