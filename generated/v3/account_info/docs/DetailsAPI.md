# \DetailsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountInfoV3Details**](DetailsAPI.md#GetAccountInfoV3Details) | **Get** /account-info/v3/details | Retrieve account details



## GetAccountInfoV3Details

> PortalInformationResponse GetAccountInfoV3Details(ctx).Execute()

Retrieve account details



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DetailsAPI.GetAccountInfoV3Details(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetailsAPI.GetAccountInfoV3Details``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInfoV3Details`: PortalInformationResponse
	fmt.Fprintf(os.Stdout, "Response from `DetailsAPI.GetAccountInfoV3Details`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoV3DetailsRequest struct via the builder pattern


### Return type

[**PortalInformationResponse**](PortalInformationResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

