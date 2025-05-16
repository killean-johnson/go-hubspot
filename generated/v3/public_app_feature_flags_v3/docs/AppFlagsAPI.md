# \AppFlagsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeatureFlagsV3AppIdFlagsFlagName**](AppFlagsAPI.md#DeleteFeatureFlagsV3AppIdFlagsFlagName) | **Delete** /feature-flags/v3/{appId}/flags/{flagName} | Delete an app&#39;s feature flag
[**GetFeatureFlagsV3AppIdFlagsFlagName**](AppFlagsAPI.md#GetFeatureFlagsV3AppIdFlagsFlagName) | **Get** /feature-flags/v3/{appId}/flags/{flagName} | Retrieve an app&#39;s feature flags
[**GetFeatureFlagsV3AppIdFlagsFlagNamePortals**](AppFlagsAPI.md#GetFeatureFlagsV3AppIdFlagsFlagNamePortals) | **Get** /feature-flags/v3/{appId}/flags/{flagName}/portals | Retrieve accounts with a set flag state
[**PutFeatureFlagsV3AppIdFlagsFlagName**](AppFlagsAPI.md#PutFeatureFlagsV3AppIdFlagsFlagName) | **Put** /feature-flags/v3/{appId}/flags/{flagName} | Set an app&#39;s feature flag



## DeleteFeatureFlagsV3AppIdFlagsFlagName

> FlagResponse DeleteFeatureFlagsV3AppIdFlagsFlagName(ctx, flagName, appId).Execute()

Delete an app's feature flag



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
	flagName := "flagName_example" // string | The name of the flag, either `hs-release-app-cards` or `hs-hide-crm-cards`.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppFlagsAPI.DeleteFeatureFlagsV3AppIdFlagsFlagName(context.Background(), flagName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppFlagsAPI.DeleteFeatureFlagsV3AppIdFlagsFlagName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFeatureFlagsV3AppIdFlagsFlagName`: FlagResponse
	fmt.Fprintf(os.Stdout, "Response from `AppFlagsAPI.DeleteFeatureFlagsV3AppIdFlagsFlagName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagName** | **string** | The name of the flag, either &#x60;hs-release-app-cards&#x60; or &#x60;hs-hide-crm-cards&#x60;. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureFlagsV3AppIdFlagsFlagNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlagResponse**](FlagResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagsV3AppIdFlagsFlagName

> FlagResponse GetFeatureFlagsV3AppIdFlagsFlagName(ctx, flagName, appId).Execute()

Retrieve an app's feature flags



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
	flagName := "flagName_example" // string | The name of the flag, either `hs-release-app-cards` or `hs-hide-crm-cards`.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppFlagsAPI.GetFeatureFlagsV3AppIdFlagsFlagName(context.Background(), flagName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppFlagsAPI.GetFeatureFlagsV3AppIdFlagsFlagName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlagsV3AppIdFlagsFlagName`: FlagResponse
	fmt.Fprintf(os.Stdout, "Response from `AppFlagsAPI.GetFeatureFlagsV3AppIdFlagsFlagName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagName** | **string** | The name of the flag, either &#x60;hs-release-app-cards&#x60; or &#x60;hs-hide-crm-cards&#x60;. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagsV3AppIdFlagsFlagNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlagResponse**](FlagResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagsV3AppIdFlagsFlagNamePortals

> PortalFlagStateBatchResponse GetFeatureFlagsV3AppIdFlagsFlagNamePortals(ctx, flagName, appId).StartPortalId(startPortalId).Limit(limit).Execute()

Retrieve accounts with a set flag state



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
	flagName := "flagName_example" // string | The name of the flag, either `hs-release-app-cards` or `hs-hide-crm-cards`.
	appId := int32(56) // int32 | The ID of the app.
	startPortalId := int32(56) // int32 | The initial account ID for listing, enabling pagination. (optional)
	limit := int32(56) // int32 | The maximum number of results to return in a single request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppFlagsAPI.GetFeatureFlagsV3AppIdFlagsFlagNamePortals(context.Background(), flagName, appId).StartPortalId(startPortalId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppFlagsAPI.GetFeatureFlagsV3AppIdFlagsFlagNamePortals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlagsV3AppIdFlagsFlagNamePortals`: PortalFlagStateBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `AppFlagsAPI.GetFeatureFlagsV3AppIdFlagsFlagNamePortals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagName** | **string** | The name of the flag, either &#x60;hs-release-app-cards&#x60; or &#x60;hs-hide-crm-cards&#x60;. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagsV3AppIdFlagsFlagNamePortalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startPortalId** | **int32** | The initial account ID for listing, enabling pagination. | 
 **limit** | **int32** | The maximum number of results to return in a single request. | 

### Return type

[**PortalFlagStateBatchResponse**](PortalFlagStateBatchResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFeatureFlagsV3AppIdFlagsFlagName

> FlagResponse PutFeatureFlagsV3AppIdFlagsFlagName(ctx, flagName, appId).FlagPutRequest(flagPutRequest).Execute()

Set an app's feature flag



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
	flagName := "flagName_example" // string | The name of the flag, either `hs-release-app-cards` or `hs-hide-crm-cards`.
	appId := int32(56) // int32 | The ID of the app.
	flagPutRequest := *openapiclient.NewFlagPutRequest("DefaultState_example") // FlagPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppFlagsAPI.PutFeatureFlagsV3AppIdFlagsFlagName(context.Background(), flagName, appId).FlagPutRequest(flagPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppFlagsAPI.PutFeatureFlagsV3AppIdFlagsFlagName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFeatureFlagsV3AppIdFlagsFlagName`: FlagResponse
	fmt.Fprintf(os.Stdout, "Response from `AppFlagsAPI.PutFeatureFlagsV3AppIdFlagsFlagName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagName** | **string** | The name of the flag, either &#x60;hs-release-app-cards&#x60; or &#x60;hs-hide-crm-cards&#x60;. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFeatureFlagsV3AppIdFlagsFlagNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flagPutRequest** | [**FlagPutRequest**](FlagPutRequest.md) |  | 

### Return type

[**FlagResponse**](FlagResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

