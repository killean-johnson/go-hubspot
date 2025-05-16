# \ValidationAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCmsV3SourceCodeEnvironmentValidatePathDoValidate**](ValidationAPI.md#PostCmsV3SourceCodeEnvironmentValidatePathDoValidate) | **Post** /cms/v3/source-code/{environment}/validate/{path} | Validate the contents of a file



## PostCmsV3SourceCodeEnvironmentValidatePathDoValidate

> Error PostCmsV3SourceCodeEnvironmentValidatePathDoValidate(ctx, path).File(file).Execute()

Validate the contents of a file



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
	path := "path_example" // string | The file system location of the file.
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationAPI.PostCmsV3SourceCodeEnvironmentValidatePathDoValidate(context.Background(), path).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationAPI.PostCmsV3SourceCodeEnvironmentValidatePathDoValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3SourceCodeEnvironmentValidatePathDoValidate`: Error
	fmt.Fprintf(os.Stdout, "Response from `ValidationAPI.PostCmsV3SourceCodeEnvironmentValidatePathDoValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3SourceCodeEnvironmentValidatePathDoValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

