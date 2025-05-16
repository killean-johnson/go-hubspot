# \SettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive**](SettingsAPI.md#DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive) | **Delete** /crm/v3/extensions/videoconferencing/settings/{appId} | Delete settings
[**GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById**](SettingsAPI.md#GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById) | **Get** /crm/v3/extensions/videoconferencing/settings/{appId} | Get settings
[**PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace**](SettingsAPI.md#PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace) | **Put** /crm/v3/extensions/videoconferencing/settings/{appId} | Update settings



## DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive

> DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive(ctx, appId).Execute()

Delete settings



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
	appId := int32(56) // int32 | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveRequest struct via the builder pattern


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


## GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById

> ExternalSettings GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById(ctx, appId).Execute()

Get settings



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
	appId := int32(56) // int32 | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById`: ExternalSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalSettings**](ExternalSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace

> ExternalSettings PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace(ctx, appId).ExternalSettings(externalSettings).Execute()

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
	appId := int32(56) // int32 | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal.
	externalSettings := *openapiclient.NewExternalSettings("https://example.com/create-meeting") // ExternalSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace(context.Background(), appId).ExternalSettings(externalSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace`: ExternalSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalSettings** | [**ExternalSettings**](ExternalSettings.md) |  | 

### Return type

[**ExternalSettings**](ExternalSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

