# \SettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3MarketingEventsAppIdSettingsGetAll**](SettingsAPI.md#GetMarketingV3MarketingEventsAppIdSettingsGetAll) | **Get** /marketing/v3/marketing-events/{appId}/settings | Retrieve the application settings
[**PostMarketingV3MarketingEventsAppIdSettingsUpdate**](SettingsAPI.md#PostMarketingV3MarketingEventsAppIdSettingsUpdate) | **Post** /marketing/v3/marketing-events/{appId}/settings | Update the application settings



## GetMarketingV3MarketingEventsAppIdSettingsGetAll

> EventDetailSettings GetMarketingV3MarketingEventsAppIdSettingsGetAll(ctx, appId).Execute()

Retrieve the application settings



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
	appId := int32(56) // int32 | The id of the application to retrieve the settings for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetMarketingV3MarketingEventsAppIdSettingsGetAll(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetMarketingV3MarketingEventsAppIdSettingsGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsAppIdSettingsGetAll`: EventDetailSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetMarketingV3MarketingEventsAppIdSettingsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The id of the application to retrieve the settings for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsAppIdSettingsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsAppIdSettingsUpdate

> EventDetailSettings PostMarketingV3MarketingEventsAppIdSettingsUpdate(ctx, appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()

Update the application settings



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
	appId := int32(56) // int32 | The id of the application to update the settings for.
	eventDetailSettingsUrl := *openapiclient.NewEventDetailSettingsUrl("EventDetailsUrl_example") // EventDetailSettingsUrl | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PostMarketingV3MarketingEventsAppIdSettingsUpdate(context.Background(), appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostMarketingV3MarketingEventsAppIdSettingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsAppIdSettingsUpdate`: EventDetailSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PostMarketingV3MarketingEventsAppIdSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The id of the application to update the settings for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsAppIdSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventDetailSettingsUrl** | [**EventDetailSettingsUrl**](EventDetailSettingsUrl.md) |  | 

### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

