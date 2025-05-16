# \ChannelConnectionSettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ExtensionsCallingAppIdSettingsChannelConnectionArchive**](ChannelConnectionSettingsAPI.md#DeleteCrmV3ExtensionsCallingAppIdSettingsChannelConnectionArchive) | **Delete** /crm/v3/extensions/calling/{appId}/settings/channel-connection | Delete channel connection settings
[**GetCrmV3ExtensionsCallingAppIdSettingsChannelConnectionGetById**](ChannelConnectionSettingsAPI.md#GetCrmV3ExtensionsCallingAppIdSettingsChannelConnectionGetById) | **Get** /crm/v3/extensions/calling/{appId}/settings/channel-connection | Retrieve channel connection settings
[**PatchCrmV3ExtensionsCallingAppIdSettingsChannelConnectionUpdate**](ChannelConnectionSettingsAPI.md#PatchCrmV3ExtensionsCallingAppIdSettingsChannelConnectionUpdate) | **Patch** /crm/v3/extensions/calling/{appId}/settings/channel-connection | Update channel connection settings
[**PostCrmV3ExtensionsCallingAppIdSettingsChannelConnectionCreate**](ChannelConnectionSettingsAPI.md#PostCrmV3ExtensionsCallingAppIdSettingsChannelConnectionCreate) | **Post** /crm/v3/extensions/calling/{appId}/settings/channel-connection | Configure channel connection settings



## DeleteCrmV3ExtensionsCallingAppIdSettingsChannelConnectionArchive

> DeleteCrmV3ExtensionsCallingAppIdSettingsChannelConnectionArchive(ctx, appId).Execute()

Delete channel connection settings



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
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChannelConnectionSettingsAPI.DeleteCrmV3ExtensionsCallingAppIdSettingsChannelConnectionArchive(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelConnectionSettingsAPI.DeleteCrmV3ExtensionsCallingAppIdSettingsChannelConnectionArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ExtensionsCallingAppIdSettingsChannelConnectionArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ExtensionsCallingAppIdSettingsChannelConnectionGetById

> ChannelConnectionSettingsResponse GetCrmV3ExtensionsCallingAppIdSettingsChannelConnectionGetById(ctx, appId).Execute()

Retrieve channel connection settings



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
	appId := int32(56) // int32 | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelConnectionSettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsChannelConnectionGetById(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelConnectionSettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsChannelConnectionGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExtensionsCallingAppIdSettingsChannelConnectionGetById`: ChannelConnectionSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelConnectionSettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsChannelConnectionGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCallingAppIdSettingsChannelConnectionGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChannelConnectionSettingsResponse**](ChannelConnectionSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsCallingAppIdSettingsChannelConnectionUpdate

> ChannelConnectionSettingsResponse PatchCrmV3ExtensionsCallingAppIdSettingsChannelConnectionUpdate(ctx, appId).ChannelConnectionSettingsPatchRequest(channelConnectionSettingsPatchRequest).Execute()

Update channel connection settings



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
	appId := int32(56) // int32 | The ID of the app.
	channelConnectionSettingsPatchRequest := *openapiclient.NewChannelConnectionSettingsPatchRequest() // ChannelConnectionSettingsPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelConnectionSettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsChannelConnectionUpdate(context.Background(), appId).ChannelConnectionSettingsPatchRequest(channelConnectionSettingsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelConnectionSettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsChannelConnectionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3ExtensionsCallingAppIdSettingsChannelConnectionUpdate`: ChannelConnectionSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelConnectionSettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsChannelConnectionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsCallingAppIdSettingsChannelConnectionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelConnectionSettingsPatchRequest** | [**ChannelConnectionSettingsPatchRequest**](ChannelConnectionSettingsPatchRequest.md) |  | 

### Return type

[**ChannelConnectionSettingsResponse**](ChannelConnectionSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCallingAppIdSettingsChannelConnectionCreate

> ChannelConnectionSettingsResponse PostCrmV3ExtensionsCallingAppIdSettingsChannelConnectionCreate(ctx, appId).ChannelConnectionSettingsRequest(channelConnectionSettingsRequest).Execute()

Configure channel connection settings



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
	appId := int32(56) // int32 | The ID of the app.
	channelConnectionSettingsRequest := *openapiclient.NewChannelConnectionSettingsRequest(false, "Url_example") // ChannelConnectionSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelConnectionSettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsChannelConnectionCreate(context.Background(), appId).ChannelConnectionSettingsRequest(channelConnectionSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelConnectionSettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsChannelConnectionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ExtensionsCallingAppIdSettingsChannelConnectionCreate`: ChannelConnectionSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelConnectionSettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsChannelConnectionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCallingAppIdSettingsChannelConnectionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelConnectionSettingsRequest** | [**ChannelConnectionSettingsRequest**](ChannelConnectionSettingsRequest.md) |  | 

### Return type

[**ChannelConnectionSettingsResponse**](ChannelConnectionSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

