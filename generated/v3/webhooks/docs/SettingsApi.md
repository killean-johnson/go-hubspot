# \SettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhooksV3AppIdSettingsClear**](SettingsAPI.md#DeleteWebhooksV3AppIdSettingsClear) | **Delete** /webhooks/v3/{appId}/settings | Delete webhook settings
[**GetWebhooksV3AppIdSettingsGetAll**](SettingsAPI.md#GetWebhooksV3AppIdSettingsGetAll) | **Get** /webhooks/v3/{appId}/settings | Read webhook settings
[**PutWebhooksV3AppIdSettingsConfigure**](SettingsAPI.md#PutWebhooksV3AppIdSettingsConfigure) | **Put** /webhooks/v3/{appId}/settings | Update webhook settings



## DeleteWebhooksV3AppIdSettingsClear

> DeleteWebhooksV3AppIdSettingsClear(ctx, appId).Execute()

Delete webhook settings



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
	r, err := apiClient.SettingsAPI.DeleteWebhooksV3AppIdSettingsClear(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteWebhooksV3AppIdSettingsClear``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWebhooksV3AppIdSettingsClearRequest struct via the builder pattern


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


## GetWebhooksV3AppIdSettingsGetAll

> SettingsResponse GetWebhooksV3AppIdSettingsGetAll(ctx, appId).Execute()

Read webhook settings



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
	resp, r, err := apiClient.SettingsAPI.GetWebhooksV3AppIdSettingsGetAll(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetWebhooksV3AppIdSettingsGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhooksV3AppIdSettingsGetAll`: SettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetWebhooksV3AppIdSettingsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksV3AppIdSettingsGetAllRequest struct via the builder pattern


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


## PutWebhooksV3AppIdSettingsConfigure

> SettingsResponse PutWebhooksV3AppIdSettingsConfigure(ctx, appId).SettingsChangeRequest(settingsChangeRequest).Execute()

Update webhook settings



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
	settingsChangeRequest := *openapiclient.NewSettingsChangeRequest(*openapiclient.NewThrottlingSettings(int32(123)), "https://www.example.com/hubspot/target") // SettingsChangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PutWebhooksV3AppIdSettingsConfigure(context.Background(), appId).SettingsChangeRequest(settingsChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PutWebhooksV3AppIdSettingsConfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutWebhooksV3AppIdSettingsConfigure`: SettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PutWebhooksV3AppIdSettingsConfigure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhooksV3AppIdSettingsConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsChangeRequest** | [**SettingsChangeRequest**](SettingsChangeRequest.md) |  | 

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

