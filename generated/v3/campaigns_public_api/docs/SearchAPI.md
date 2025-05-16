# \SearchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3Campaigns**](SearchAPI.md#GetMarketingV3Campaigns) | **Get** /marketing/v3/campaigns/ | Campaign search



## GetMarketingV3Campaigns

> CollectionResponseWithTotalPublicCampaignForwardPaging GetMarketingV3Campaigns(ctx).Sort(sort).After(after).Limit(limit).Name(name).Properties(properties).Execute()

Campaign search



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
	sort := "sort_example" // string | The field by which to sort the results. Allowed values are hs_name, createdAt, updatedAt. An optional '-' before the property name can denote descending order Default: hs_name (optional)
	after := "after_example" // string | A cursor for pagination. If provided, the results will start after the given cursor. Example: NTI1Cg%3D%3D (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Allowed values range from 1 to 100 Default: 50 (optional)
	name := "name_example" // string | A filter to return campaigns whose names contain the specified substring. This allows partial matching of campaign names, returning all campaigns that include the given substring in their name. If this parameter is not provided, the search will return all campaigns (optional)
	properties := []string{"Inner_example"} // []string | A comma-separated list of the properties to be returned in the response. If any of the specified properties has empty value on the requested object(s), they will be ignored and not returned in response. If this parameter is empty, the response will include an empty properties map (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.GetMarketingV3Campaigns(context.Background()).Sort(sort).After(after).Limit(limit).Name(name).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetMarketingV3Campaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3Campaigns`: CollectionResponseWithTotalPublicCampaignForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetMarketingV3Campaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3CampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | The field by which to sort the results. Allowed values are hs_name, createdAt, updatedAt. An optional &#39;-&#39; before the property name can denote descending order Default: hs_name | 
 **after** | **string** | A cursor for pagination. If provided, the results will start after the given cursor. Example: NTI1Cg%3D%3D | 
 **limit** | **int32** | The maximum number of results to return. Allowed values range from 1 to 100 Default: 50 | 
 **name** | **string** | A filter to return campaigns whose names contain the specified substring. This allows partial matching of campaign names, returning all campaigns that include the given substring in their name. If this parameter is not provided, the search will return all campaigns | 
 **properties** | **[]string** | A comma-separated list of the properties to be returned in the response. If any of the specified properties has empty value on the requested object(s), they will be ignored and not returned in response. If this parameter is empty, the response will include an empty properties map | 

### Return type

[**CollectionResponseWithTotalPublicCampaignForwardPaging**](CollectionResponseWithTotalPublicCampaignForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

