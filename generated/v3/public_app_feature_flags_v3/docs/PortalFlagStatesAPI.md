# \PortalFlagStatesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId**](PortalFlagStatesAPI.md#DeleteFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId) | **Delete** /feature-flags/v3/{appId}/flags/{flagName}/portals/{portalId} | Delete an account flag state
[**GetFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId**](PortalFlagStatesAPI.md#GetFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId) | **Get** /feature-flags/v3/{appId}/flags/{flagName}/portals/{portalId} | Retrieve account flag state
[**PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchDelete**](PortalFlagStatesAPI.md#PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchDelete) | **Post** /feature-flags/v3/{appId}/flags/{flagName}/portals/batch/delete | Batch delete account flag state
[**PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchUpsert**](PortalFlagStatesAPI.md#PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchUpsert) | **Post** /feature-flags/v3/{appId}/flags/{flagName}/portals/batch/upsert | Batch set account flag state
[**PutFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId**](PortalFlagStatesAPI.md#PutFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId) | **Put** /feature-flags/v3/{appId}/flags/{flagName}/portals/{portalId} | Set an account flag state



## DeleteFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId

> PortalFlagStateResponse DeleteFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId(ctx, flagName, portalId, appId).Execute()

Delete an account flag state



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
	portalId := int32(56) // int32 | The ID of the account that installed the app.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalFlagStatesAPI.DeleteFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId(context.Background(), flagName, portalId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalFlagStatesAPI.DeleteFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId`: PortalFlagStateResponse
	fmt.Fprintf(os.Stdout, "Response from `PortalFlagStatesAPI.DeleteFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagName** | **string** | The name of the flag, either &#x60;hs-release-app-cards&#x60; or &#x60;hs-hide-crm-cards&#x60;. | 
**portalId** | **int32** | The ID of the account that installed the app. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PortalFlagStateResponse**](PortalFlagStateResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId

> PortalFlagStateResponse GetFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId(ctx, flagName, portalId, appId).Execute()

Retrieve account flag state



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
	portalId := int32(56) // int32 | The ID of the account that installed the app.
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalFlagStatesAPI.GetFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId(context.Background(), flagName, portalId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalFlagStatesAPI.GetFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId`: PortalFlagStateResponse
	fmt.Fprintf(os.Stdout, "Response from `PortalFlagStatesAPI.GetFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagName** | **string** | The name of the flag, either &#x60;hs-release-app-cards&#x60; or &#x60;hs-hide-crm-cards&#x60;. | 
**portalId** | **int32** | The ID of the account that installed the app. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PortalFlagStateResponse**](PortalFlagStateResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchDelete

> PortalFlagStateBatchResponse PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchDelete(ctx, flagName, appId).PortalFlagStateBatchDeleteRequest(portalFlagStateBatchDeleteRequest).Execute()

Batch delete account flag state



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
	portalFlagStateBatchDeleteRequest := *openapiclient.NewPortalFlagStateBatchDeleteRequest([]int32{int32(123)}) // PortalFlagStateBatchDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalFlagStatesAPI.PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchDelete(context.Background(), flagName, appId).PortalFlagStateBatchDeleteRequest(portalFlagStateBatchDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalFlagStatesAPI.PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchDelete`: PortalFlagStateBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `PortalFlagStatesAPI.PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagName** | **string** | The name of the flag, either &#x60;hs-release-app-cards&#x60; or &#x60;hs-hide-crm-cards&#x60;. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portalFlagStateBatchDeleteRequest** | [**PortalFlagStateBatchDeleteRequest**](PortalFlagStateBatchDeleteRequest.md) |  | 

### Return type

[**PortalFlagStateBatchResponse**](PortalFlagStateBatchResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchUpsert

> PortalFlagStateBatchResponse PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchUpsert(ctx, flagName, appId).PortalFlagStateBatchPutRequest(portalFlagStateBatchPutRequest).Execute()

Batch set account flag state



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
	portalFlagStateBatchPutRequest := *openapiclient.NewPortalFlagStateBatchPutRequest([]openapiclient.BatchPortalEntry{*openapiclient.NewBatchPortalEntry(int32(123), "FlagState_example")}) // PortalFlagStateBatchPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalFlagStatesAPI.PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchUpsert(context.Background(), flagName, appId).PortalFlagStateBatchPutRequest(portalFlagStateBatchPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalFlagStatesAPI.PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchUpsert`: PortalFlagStateBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `PortalFlagStatesAPI.PostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchUpsert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagName** | **string** | The name of the flag, either &#x60;hs-release-app-cards&#x60; or &#x60;hs-hide-crm-cards&#x60;. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFeatureFlagsV3AppIdFlagsFlagNamePortalsBatchUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portalFlagStateBatchPutRequest** | [**PortalFlagStateBatchPutRequest**](PortalFlagStateBatchPutRequest.md) |  | 

### Return type

[**PortalFlagStateBatchResponse**](PortalFlagStateBatchResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId

> PortalFlagStateResponse PutFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId(ctx, flagName, portalId, appId).PortalFlagStatePutRequest(portalFlagStatePutRequest).Execute()

Set an account flag state



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
	portalId := int32(56) // int32 | The ID of the account that installed the app.
	appId := int32(56) // int32 | The ID of the app.
	portalFlagStatePutRequest := *openapiclient.NewPortalFlagStatePutRequest("FlagState_example") // PortalFlagStatePutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalFlagStatesAPI.PutFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId(context.Background(), flagName, portalId, appId).PortalFlagStatePutRequest(portalFlagStatePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalFlagStatesAPI.PutFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId`: PortalFlagStateResponse
	fmt.Fprintf(os.Stdout, "Response from `PortalFlagStatesAPI.PutFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagName** | **string** | The name of the flag, either &#x60;hs-release-app-cards&#x60; or &#x60;hs-hide-crm-cards&#x60;. | 
**portalId** | **int32** | The ID of the account that installed the app. | 
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFeatureFlagsV3AppIdFlagsFlagNamePortalsPortalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **portalFlagStatePutRequest** | [**PortalFlagStatePutRequest**](PortalFlagStatePutRequest.md) |  | 

### Return type

[**PortalFlagStateResponse**](PortalFlagStateResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

