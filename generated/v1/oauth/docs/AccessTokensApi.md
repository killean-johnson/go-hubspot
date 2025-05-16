# \AccessTokensAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauthV1AccessTokensTokenGet**](AccessTokensAPI.md#GetOauthV1AccessTokensTokenGet) | **Get** /oauth/v1/access-tokens/{token} | Retrieve OAuth token metadata



## GetOauthV1AccessTokensTokenGet

> AccessTokenInfoResponse GetOauthV1AccessTokensTokenGet(ctx, token).Execute()

Retrieve OAuth token metadata



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
	token := "token_example" // string | The access token that you want to retrieve information about.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.GetOauthV1AccessTokensTokenGet(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.GetOauthV1AccessTokensTokenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthV1AccessTokensTokenGet`: AccessTokenInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.GetOauthV1AccessTokensTokenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The access token that you want to retrieve information about. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthV1AccessTokensTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessTokenInfoResponse**](AccessTokenInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

