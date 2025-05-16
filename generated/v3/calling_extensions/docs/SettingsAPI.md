# \SettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ExtensionsCallingAppIdSettingsArchive**](SettingsAPI.md#DeleteCrmV3ExtensionsCallingAppIdSettingsArchive) | **Delete** /crm/v3/extensions/calling/{appId}/settings | Delete calling settings
[**GetCrmV3ExtensionsCallingAppIdSettingsGetById**](SettingsAPI.md#GetCrmV3ExtensionsCallingAppIdSettingsGetById) | **Get** /crm/v3/extensions/calling/{appId}/settings | Retrieve settings
[**PatchCrmV3ExtensionsCallingAppIdSettingsUpdate**](SettingsAPI.md#PatchCrmV3ExtensionsCallingAppIdSettingsUpdate) | **Patch** /crm/v3/extensions/calling/{appId}/settings | Update settings
[**PostCrmV3ExtensionsCallingAppIdSettingsCreate**](SettingsAPI.md#PostCrmV3ExtensionsCallingAppIdSettingsCreate) | **Post** /crm/v3/extensions/calling/{appId}/settings | Configure a calling extension



## DeleteCrmV3ExtensionsCallingAppIdSettingsArchive

> DeleteCrmV3ExtensionsCallingAppIdSettingsArchive(ctx, appId).Execute()

Delete calling settings



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
	r, err := apiClient.SettingsAPI.DeleteCrmV3ExtensionsCallingAppIdSettingsArchive(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteCrmV3ExtensionsCallingAppIdSettingsArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCrmV3ExtensionsCallingAppIdSettingsArchiveRequest struct via the builder pattern


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


## GetCrmV3ExtensionsCallingAppIdSettingsGetById

> SettingsResponse GetCrmV3ExtensionsCallingAppIdSettingsGetById(ctx, appId).Execute()

Retrieve settings



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
	resp, r, err := apiClient.SettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsGetById(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExtensionsCallingAppIdSettingsGetById`: SettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCallingAppIdSettingsGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsCallingAppIdSettingsUpdate

> SettingsResponse PatchCrmV3ExtensionsCallingAppIdSettingsUpdate(ctx, appId).SettingsPatchRequest(settingsPatchRequest).Execute()

Update settings



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
	settingsPatchRequest := *openapiclient.NewSettingsPatchRequest() // SettingsPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsUpdate(context.Background(), appId).SettingsPatchRequest(settingsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3ExtensionsCallingAppIdSettingsUpdate`: SettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsPatchRequest** | [**SettingsPatchRequest**](SettingsPatchRequest.md) |  | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCallingAppIdSettingsCreate

> SettingsResponse PostCrmV3ExtensionsCallingAppIdSettingsCreate(ctx, appId).SettingsRequest(settingsRequest).Execute()

Configure a calling extension



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
	settingsRequest := *openapiclient.NewSettingsRequest("HubPhone", "https://www.example.com/hubspot/iframe") // SettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsCreate(context.Background(), appId).SettingsRequest(settingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ExtensionsCallingAppIdSettingsCreate`: SettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsRequest** | [**SettingsRequest**](SettingsRequest.md) |  | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

