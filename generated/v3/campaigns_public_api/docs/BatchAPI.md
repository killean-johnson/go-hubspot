# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3CampaignsBatchArchive**](BatchAPI.md#PostMarketingV3CampaignsBatchArchive) | **Post** /marketing/v3/campaigns/batch/archive | Delete a batch of campaigns
[**PostMarketingV3CampaignsBatchCreate**](BatchAPI.md#PostMarketingV3CampaignsBatchCreate) | **Post** /marketing/v3/campaigns/batch/create | Create a batch of campaigns
[**PostMarketingV3CampaignsBatchRead**](BatchAPI.md#PostMarketingV3CampaignsBatchRead) | **Post** /marketing/v3/campaigns/batch/read | Read a batch of campaigns
[**PostMarketingV3CampaignsBatchUpdate**](BatchAPI.md#PostMarketingV3CampaignsBatchUpdate) | **Post** /marketing/v3/campaigns/batch/update | Update a batch of campaigns



## PostMarketingV3CampaignsBatchArchive

> PostMarketingV3CampaignsBatchArchive(ctx).BatchInputPublicCampaignDeleteInput(batchInputPublicCampaignDeleteInput).Execute()

Delete a batch of campaigns



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
	batchInputPublicCampaignDeleteInput := *openapiclient.NewBatchInputPublicCampaignDeleteInput([]openapiclient.PublicCampaignDeleteInput{*openapiclient.NewPublicCampaignDeleteInput("Id_example")}) // BatchInputPublicCampaignDeleteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.PostMarketingV3CampaignsBatchArchive(context.Background()).BatchInputPublicCampaignDeleteInput(batchInputPublicCampaignDeleteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3CampaignsBatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3CampaignsBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputPublicCampaignDeleteInput** | [**BatchInputPublicCampaignDeleteInput**](BatchInputPublicCampaignDeleteInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3CampaignsBatchCreate

> BatchResponsePublicCampaign PostMarketingV3CampaignsBatchCreate(ctx).BatchInputPublicCampaignInput(batchInputPublicCampaignInput).Execute()

Create a batch of campaigns



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
	batchInputPublicCampaignInput := *openapiclient.NewBatchInputPublicCampaignInput([]openapiclient.PublicCampaignInput{*openapiclient.NewPublicCampaignInput(map[string]string{"key": "Inner_example"})}) // BatchInputPublicCampaignInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostMarketingV3CampaignsBatchCreate(context.Background()).BatchInputPublicCampaignInput(batchInputPublicCampaignInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3CampaignsBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3CampaignsBatchCreate`: BatchResponsePublicCampaign
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostMarketingV3CampaignsBatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3CampaignsBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputPublicCampaignInput** | [**BatchInputPublicCampaignInput**](BatchInputPublicCampaignInput.md) |  | 

### Return type

[**BatchResponsePublicCampaign**](BatchResponsePublicCampaign.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3CampaignsBatchRead

> BatchResponsePublicCampaignWithAssets PostMarketingV3CampaignsBatchRead(ctx).BatchInputPublicCampaignReadInput(batchInputPublicCampaignReadInput).StartDate(startDate).EndDate(endDate).Properties(properties).Execute()

Read a batch of campaigns



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
	batchInputPublicCampaignReadInput := *openapiclient.NewBatchInputPublicCampaignReadInput([]openapiclient.PublicCampaignReadInput{*openapiclient.NewPublicCampaignReadInput("Id_example")}) // BatchInputPublicCampaignReadInput | 
	startDate := "startDate_example" // string | Start date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched.  (optional)
	endDate := "endDate_example" // string | End date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched. (optional)
	properties := []string{"Inner_example"} // []string | A comma-separated list of the properties to be returned in the response. If any of the specified properties has empty value on the requested object(s), they will be ignored and not returned in response. If this parameter is empty, the response will include an empty properties map. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostMarketingV3CampaignsBatchRead(context.Background()).BatchInputPublicCampaignReadInput(batchInputPublicCampaignReadInput).StartDate(startDate).EndDate(endDate).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3CampaignsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3CampaignsBatchRead`: BatchResponsePublicCampaignWithAssets
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostMarketingV3CampaignsBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3CampaignsBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputPublicCampaignReadInput** | [**BatchInputPublicCampaignReadInput**](BatchInputPublicCampaignReadInput.md) |  | 
 **startDate** | **string** | Start date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched.  | 
 **endDate** | **string** | End date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched. | 
 **properties** | **[]string** | A comma-separated list of the properties to be returned in the response. If any of the specified properties has empty value on the requested object(s), they will be ignored and not returned in response. If this parameter is empty, the response will include an empty properties map. | 

### Return type

[**BatchResponsePublicCampaignWithAssets**](BatchResponsePublicCampaignWithAssets.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3CampaignsBatchUpdate

> BatchResponsePublicCampaign PostMarketingV3CampaignsBatchUpdate(ctx).BatchInputPublicCampaignBatchUpdateItem(batchInputPublicCampaignBatchUpdateItem).Execute()

Update a batch of campaigns



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
	batchInputPublicCampaignBatchUpdateItem := *openapiclient.NewBatchInputPublicCampaignBatchUpdateItem([]openapiclient.PublicCampaignBatchUpdateItem{*openapiclient.NewPublicCampaignBatchUpdateItem("Id_example", map[string]string{"key": "Inner_example"})}) // BatchInputPublicCampaignBatchUpdateItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostMarketingV3CampaignsBatchUpdate(context.Background()).BatchInputPublicCampaignBatchUpdateItem(batchInputPublicCampaignBatchUpdateItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3CampaignsBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3CampaignsBatchUpdate`: BatchResponsePublicCampaign
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostMarketingV3CampaignsBatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3CampaignsBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputPublicCampaignBatchUpdateItem** | [**BatchInputPublicCampaignBatchUpdateItem**](BatchInputPublicCampaignBatchUpdateItem.md) |  | 

### Return type

[**BatchResponsePublicCampaign**](BatchResponsePublicCampaign.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

