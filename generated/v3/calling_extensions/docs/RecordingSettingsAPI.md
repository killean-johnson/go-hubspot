# \RecordingSettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat**](RecordingSettingsAPI.md#GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat) | **Get** /crm/v3/extensions/calling/{appId}/settings/recording | Retrieve recording settings
[**PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat**](RecordingSettingsAPI.md#PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat) | **Patch** /crm/v3/extensions/calling/{appId}/settings/recording | Update recording settings
[**PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat**](RecordingSettingsAPI.md#PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat) | **Post** /crm/v3/extensions/calling/{appId}/settings/recording | Enable the app for call recording
[**PostCrmV3ExtensionsCallingRecordingsReadyMarkAsReady**](RecordingSettingsAPI.md#PostCrmV3ExtensionsCallingRecordingsReadyMarkAsReady) | **Post** /crm/v3/extensions/calling/recordings/ready | Mark recording as ready for transcription



## GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat

> RecordingSettingsResponse GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat(ctx, appId).Execute()

Retrieve recording settings



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
	resp, r, err := apiClient.RecordingSettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat`: RecordingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordingSettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecordingSettingsResponse**](RecordingSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat

> RecordingSettingsResponse PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat(ctx, appId).RecordingSettingsPatchRequest(recordingSettingsPatchRequest).Execute()

Update recording settings



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
	recordingSettingsPatchRequest := *openapiclient.NewRecordingSettingsPatchRequest() // RecordingSettingsPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordingSettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat(context.Background(), appId).RecordingSettingsPatchRequest(recordingSettingsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat`: RecordingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordingSettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordingSettingsPatchRequest** | [**RecordingSettingsPatchRequest**](RecordingSettingsPatchRequest.md) |  | 

### Return type

[**RecordingSettingsResponse**](RecordingSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat

> RecordingSettingsResponse PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat(ctx, appId).RecordingSettingsRequest(recordingSettingsRequest).Execute()

Enable the app for call recording



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
	recordingSettingsRequest := *openapiclient.NewRecordingSettingsRequest("UrlToRetrieveAuthedRecording_example") // RecordingSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordingSettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat(context.Background(), appId).RecordingSettingsRequest(recordingSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat`: RecordingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordingSettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordingSettingsRequest** | [**RecordingSettingsRequest**](RecordingSettingsRequest.md) |  | 

### Return type

[**RecordingSettingsResponse**](RecordingSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCallingRecordingsReadyMarkAsReady

> PostCrmV3ExtensionsCallingRecordingsReadyMarkAsReady(ctx).MarkRecordingAsReadyRequest(markRecordingAsReadyRequest).Execute()

Mark recording as ready for transcription



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
	markRecordingAsReadyRequest := *openapiclient.NewMarkRecordingAsReadyRequest(int64(123)) // MarkRecordingAsReadyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecordingSettingsAPI.PostCrmV3ExtensionsCallingRecordingsReadyMarkAsReady(context.Background()).MarkRecordingAsReadyRequest(markRecordingAsReadyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsAPI.PostCrmV3ExtensionsCallingRecordingsReadyMarkAsReady``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCallingRecordingsReadyMarkAsReadyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markRecordingAsReadyRequest** | [**MarkRecordingAsReadyRequest**](MarkRecordingAsReadyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

