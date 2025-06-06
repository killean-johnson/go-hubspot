# \TokensAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostOauthV1TokenCreate**](TokensAPI.md#PostOauthV1TokenCreate) | **Post** /oauth/v1/token | Refresh an access token



## PostOauthV1TokenCreate

> TokenResponseIF PostOauthV1TokenCreate(ctx).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).ClientSecret(clientSecret).RefreshToken(refreshToken).Execute()

Refresh an access token



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
	grantType := "grantType_example" // string |  (optional)
	code := "code_example" // string |  (optional)
	redirectUri := "redirectUri_example" // string |  (optional)
	clientId := "clientId_example" // string |  (optional)
	clientSecret := "clientSecret_example" // string |  (optional)
	refreshToken := "refreshToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.PostOauthV1TokenCreate(context.Background()).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).ClientSecret(clientSecret).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.PostOauthV1TokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOauthV1TokenCreate`: TokenResponseIF
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.PostOauthV1TokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOauthV1TokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **code** | **string** |  | 
 **redirectUri** | **string** |  | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 
 **refreshToken** | **string** |  | 

### Return type

[**TokenResponseIF**](TokenResponseIF.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

