# \SearchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsDealsSearchDoSearch**](SearchAPI.md#PostCrmV3ObjectsDealsSearchDoSearch) | **Post** /crm/v3/objects/deals/search | 



## PostCrmV3ObjectsDealsSearchDoSearch

> CollectionResponseWithTotalSimplePublicObjectForwardPaging PostCrmV3ObjectsDealsSearchDoSearch(ctx).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()



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
	publicObjectSearchRequest := *openapiclient.NewPublicObjectSearchRequest() // PublicObjectSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.PostCrmV3ObjectsDealsSearchDoSearch(context.Background()).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.PostCrmV3ObjectsDealsSearchDoSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ObjectsDealsSearchDoSearch`: CollectionResponseWithTotalSimplePublicObjectForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.PostCrmV3ObjectsDealsSearchDoSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsDealsSearchDoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicObjectSearchRequest** | [**PublicObjectSearchRequest**](PublicObjectSearchRequest.md) |  | 

### Return type

[**CollectionResponseWithTotalSimplePublicObjectForwardPaging**](CollectionResponseWithTotalSimplePublicObjectForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

