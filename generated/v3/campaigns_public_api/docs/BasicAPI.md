# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3CampaignsCampaignGuid**](BasicAPI.md#DeleteMarketingV3CampaignsCampaignGuid) | **Delete** /marketing/v3/campaigns/{campaignGuid} | Delete campaign 
[**GetMarketingV3CampaignsCampaignGuid**](BasicAPI.md#GetMarketingV3CampaignsCampaignGuid) | **Get** /marketing/v3/campaigns/{campaignGuid} | Read a campaign
[**PatchMarketingV3CampaignsCampaignGuid**](BasicAPI.md#PatchMarketingV3CampaignsCampaignGuid) | **Patch** /marketing/v3/campaigns/{campaignGuid} | Update campaign
[**PostMarketingV3Campaigns**](BasicAPI.md#PostMarketingV3Campaigns) | **Post** /marketing/v3/campaigns/ | Create a campaign



## DeleteMarketingV3CampaignsCampaignGuid

> DeleteMarketingV3CampaignsCampaignGuid(ctx, campaignGuid).Execute()

Delete campaign 



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
	campaignGuid := "campaignGuid_example" // string | Unique identifier for the campaign, formatted as a UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteMarketingV3CampaignsCampaignGuid(context.Background(), campaignGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteMarketingV3CampaignsCampaignGuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3CampaignsCampaignGuidRequest struct via the builder pattern


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


## GetMarketingV3CampaignsCampaignGuid

> PublicCampaignWithAssets GetMarketingV3CampaignsCampaignGuid(ctx, campaignGuid).StartDate(startDate).EndDate(endDate).Properties(properties).Execute()

Read a campaign



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
	campaignGuid := "campaignGuid_example" // string | Unique identifier for the campaign, formatted as a UUID.
	startDate := "startDate_example" // string | Start date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched. (optional)
	endDate := "endDate_example" // string |  End date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched. (optional)
	properties := []string{"Inner_example"} // []string | A comma-separated list of the properties to be returned in the response. If any of the specified properties has empty value on the requested object, they will be ignored and not returned in response. If this parameter is empty, the response will include an empty properties map. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetMarketingV3CampaignsCampaignGuid(context.Background(), campaignGuid).StartDate(startDate).EndDate(endDate).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetMarketingV3CampaignsCampaignGuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3CampaignsCampaignGuid`: PublicCampaignWithAssets
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetMarketingV3CampaignsCampaignGuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3CampaignsCampaignGuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Start date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched. | 
 **endDate** | **string** |  End date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched. | 
 **properties** | **[]string** | A comma-separated list of the properties to be returned in the response. If any of the specified properties has empty value on the requested object, they will be ignored and not returned in response. If this parameter is empty, the response will include an empty properties map. | 

### Return type

[**PublicCampaignWithAssets**](PublicCampaignWithAssets.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingV3CampaignsCampaignGuid

> PublicCampaign PatchMarketingV3CampaignsCampaignGuid(ctx, campaignGuid).PublicCampaignInput(publicCampaignInput).Execute()

Update campaign



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
	campaignGuid := "campaignGuid_example" // string | Unique identifier for the campaign, formatted as a UUID. 
	publicCampaignInput := *openapiclient.NewPublicCampaignInput(map[string]string{"key": "Inner_example"}) // PublicCampaignInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchMarketingV3CampaignsCampaignGuid(context.Background(), campaignGuid).PublicCampaignInput(publicCampaignInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchMarketingV3CampaignsCampaignGuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingV3CampaignsCampaignGuid`: PublicCampaign
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchMarketingV3CampaignsCampaignGuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3CampaignsCampaignGuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicCampaignInput** | [**PublicCampaignInput**](PublicCampaignInput.md) |  | 

### Return type

[**PublicCampaign**](PublicCampaign.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3Campaigns

> PublicCampaign PostMarketingV3Campaigns(ctx).PublicCampaignInput(publicCampaignInput).Execute()

Create a campaign



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
	publicCampaignInput := *openapiclient.NewPublicCampaignInput(map[string]string{"key": "Inner_example"}) // PublicCampaignInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostMarketingV3Campaigns(context.Background()).PublicCampaignInput(publicCampaignInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostMarketingV3Campaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3Campaigns`: PublicCampaign
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostMarketingV3Campaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3CampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicCampaignInput** | [**PublicCampaignInput**](PublicCampaignInput.md) |  | 

### Return type

[**PublicCampaign**](PublicCampaign.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

