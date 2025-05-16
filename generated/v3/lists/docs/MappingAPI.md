# \MappingAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ListsIdmappingTranslateLegacyListIdToListId**](MappingAPI.md#GetCrmV3ListsIdmappingTranslateLegacyListIdToListId) | **Get** /crm/v3/lists/idmapping | Translate Legacy List Id to Modern List Id
[**PostCrmV3ListsIdmappingTranslateLegacyListIdToListIdBatch**](MappingAPI.md#PostCrmV3ListsIdmappingTranslateLegacyListIdToListIdBatch) | **Post** /crm/v3/lists/idmapping | Translate Legacy List Id to Modern List Id in Batch



## GetCrmV3ListsIdmappingTranslateLegacyListIdToListId

> PublicMigrationMapping GetCrmV3ListsIdmappingTranslateLegacyListIdToListId(ctx).LegacyListId(legacyListId).Execute()

Translate Legacy List Id to Modern List Id



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
	legacyListId := "legacyListId_example" // string | The legacy list id from lists v1 API. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingAPI.GetCrmV3ListsIdmappingTranslateLegacyListIdToListId(context.Background()).LegacyListId(legacyListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingAPI.GetCrmV3ListsIdmappingTranslateLegacyListIdToListId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ListsIdmappingTranslateLegacyListIdToListId`: PublicMigrationMapping
	fmt.Fprintf(os.Stdout, "Response from `MappingAPI.GetCrmV3ListsIdmappingTranslateLegacyListIdToListId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ListsIdmappingTranslateLegacyListIdToListIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legacyListId** | **string** | The legacy list id from lists v1 API. | 

### Return type

[**PublicMigrationMapping**](PublicMigrationMapping.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ListsIdmappingTranslateLegacyListIdToListIdBatch

> PublicBatchMigrationMapping PostCrmV3ListsIdmappingTranslateLegacyListIdToListIdBatch(ctx).RequestBody(requestBody).Execute()

Translate Legacy List Id to Modern List Id in Batch



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
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingAPI.PostCrmV3ListsIdmappingTranslateLegacyListIdToListIdBatch(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingAPI.PostCrmV3ListsIdmappingTranslateLegacyListIdToListIdBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ListsIdmappingTranslateLegacyListIdToListIdBatch`: PublicBatchMigrationMapping
	fmt.Fprintf(os.Stdout, "Response from `MappingAPI.PostCrmV3ListsIdmappingTranslateLegacyListIdToListIdBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ListsIdmappingTranslateLegacyListIdToListIdBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**PublicBatchMigrationMapping**](PublicBatchMigrationMapping.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

