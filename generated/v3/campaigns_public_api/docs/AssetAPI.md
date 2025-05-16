# \AssetAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId**](AssetAPI.md#DeleteMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId) | **Delete** /marketing/v3/campaigns/{campaignGuid}/assets/{assetType}/{assetId} | Remove asset association
[**GetMarketingV3CampaignsCampaignGuidAssetsAssetType**](AssetAPI.md#GetMarketingV3CampaignsCampaignGuidAssetsAssetType) | **Get** /marketing/v3/campaigns/{campaignGuid}/assets/{assetType} | List assets
[**PutMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId**](AssetAPI.md#PutMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId) | **Put** /marketing/v3/campaigns/{campaignGuid}/assets/{assetType}/{assetId} | Add asset association



## DeleteMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId

> DeleteMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId(ctx, campaignGuid, assetType, assetId).Execute()

Remove asset association



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
	assetType := "assetType_example" // string | The type of asset Important: Currently, only the following asset types are available for disassociation via the API: FORM, OBJECT_LIST, EXTERNAL_WEB_URL
	assetId := "assetId_example" // string | Id of the asset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetAPI.DeleteMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId(context.Background(), campaignGuid, assetType, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.DeleteMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID. | 
**assetType** | **string** | The type of asset Important: Currently, only the following asset types are available for disassociation via the API: FORM, OBJECT_LIST, EXTERNAL_WEB_URL | 
**assetId** | **string** | Id of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetIdRequest struct via the builder pattern


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


## GetMarketingV3CampaignsCampaignGuidAssetsAssetType

> CollectionResponsePublicCampaignAssetForwardPaging GetMarketingV3CampaignsCampaignGuidAssetsAssetType(ctx, campaignGuid, assetType).After(after).Limit(limit).StartDate(startDate).EndDate(endDate).Execute()

List assets



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
	assetType := "assetType_example" // string | The type of asset to fetch.
	after := "after_example" // string | A cursor for pagination. If provided, the results will start after the given cursor. Example: NTI1Cg%3D%3D (optional)
	limit := "limit_example" // string | The maximum number of results to return. Default: 10 (optional)
	startDate := "startDate_example" // string | Start date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched.  (optional)
	endDate := "endDate_example" // string | End date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetAPI.GetMarketingV3CampaignsCampaignGuidAssetsAssetType(context.Background(), campaignGuid, assetType).After(after).Limit(limit).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.GetMarketingV3CampaignsCampaignGuidAssetsAssetType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3CampaignsCampaignGuidAssetsAssetType`: CollectionResponsePublicCampaignAssetForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.GetMarketingV3CampaignsCampaignGuidAssetsAssetType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID. | 
**assetType** | **string** | The type of asset to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3CampaignsCampaignGuidAssetsAssetTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | A cursor for pagination. If provided, the results will start after the given cursor. Example: NTI1Cg%3D%3D | 
 **limit** | **string** | The maximum number of results to return. Default: 10 | 
 **startDate** | **string** | Start date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched.  | 
 **endDate** | **string** | End date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched. | 

### Return type

[**CollectionResponsePublicCampaignAssetForwardPaging**](CollectionResponsePublicCampaignAssetForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId

> PutMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId(ctx, campaignGuid, assetType, assetId).Execute()

Add asset association



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
	campaignGuid := "campaignGuid_example" // string | Unique identifier for the campaign, formatted as a UUID
	assetType := "assetType_example" // string | The type of asset Important: Currently, only the following asset types are available for association via the API: FORM, OBJECT_LIST, EXTERNAL_WEB_URL
	assetId := "assetId_example" // string | Id of the asset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetAPI.PutMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId(context.Background(), campaignGuid, assetType, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.PutMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | Unique identifier for the campaign, formatted as a UUID | 
**assetType** | **string** | The type of asset Important: Currently, only the following asset types are available for association via the API: FORM, OBJECT_LIST, EXTERNAL_WEB_URL | 
**assetId** | **string** | Id of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingV3CampaignsCampaignGuidAssetsAssetTypeAssetIdRequest struct via the builder pattern


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

