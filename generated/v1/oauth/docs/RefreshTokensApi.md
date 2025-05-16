# \RefreshTokensAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOauthV1RefreshTokensTokenArchive**](RefreshTokensAPI.md#DeleteOauthV1RefreshTokensTokenArchive) | **Delete** /oauth/v1/refresh-tokens/{token} | Delete a refresh token
[**GetOauthV1RefreshTokensTokenGet**](RefreshTokensAPI.md#GetOauthV1RefreshTokensTokenGet) | **Get** /oauth/v1/refresh-tokens/{token} | Retrieve refresh token metadata



## DeleteOauthV1RefreshTokensTokenArchive

> DeleteOauthV1RefreshTokensTokenArchive(ctx, token).Execute()

Delete a refresh token



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
	token := "token_example" // string | The refresh token to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RefreshTokensAPI.DeleteOauthV1RefreshTokensTokenArchive(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefreshTokensAPI.DeleteOauthV1RefreshTokensTokenArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The refresh token to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOauthV1RefreshTokensTokenArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthV1RefreshTokensTokenGet

> RefreshTokenInfoResponse GetOauthV1RefreshTokensTokenGet(ctx, token).Execute()

Retrieve refresh token metadata



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
	token := "token_example" // string | The refresh token to retrieve information about.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefreshTokensAPI.GetOauthV1RefreshTokensTokenGet(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefreshTokensAPI.GetOauthV1RefreshTokensTokenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthV1RefreshTokensTokenGet`: RefreshTokenInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `RefreshTokensAPI.GetOauthV1RefreshTokensTokenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The refresh token to retrieve information about. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthV1RefreshTokensTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefreshTokenInfoResponse**](RefreshTokenInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

