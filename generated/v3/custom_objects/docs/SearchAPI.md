# \SearchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsObjectTypeSearchDoSearch**](SearchAPI.md#PostCrmV3ObjectsObjectTypeSearchDoSearch) | **Post** /crm/v3/objects/{objectType}/search | 



## PostCrmV3ObjectsObjectTypeSearchDoSearch

> CollectionResponseWithTotalSimplePublicObjectForwardPaging PostCrmV3ObjectsObjectTypeSearchDoSearch(ctx, objectType).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()



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
	objectType := "objectType_example" // string | 
	publicObjectSearchRequest := *openapiclient.NewPublicObjectSearchRequest() // PublicObjectSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.PostCrmV3ObjectsObjectTypeSearchDoSearch(context.Background(), objectType).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.PostCrmV3ObjectsObjectTypeSearchDoSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ObjectsObjectTypeSearchDoSearch`: CollectionResponseWithTotalSimplePublicObjectForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.PostCrmV3ObjectsObjectTypeSearchDoSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsObjectTypeSearchDoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicObjectSearchRequest** | [**PublicObjectSearchRequest**](PublicObjectSearchRequest.md) |  | 

### Return type

[**CollectionResponseWithTotalSimplePublicObjectForwardPaging**](CollectionResponseWithTotalSimplePublicObjectForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

