# \EventsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMediaBridgeV1EventsAttentionSpan**](EventsAPI.md#PostMediaBridgeV1EventsAttentionSpan) | **Post** /media-bridge/v1/events/attention-span | 
[**PostMediaBridgeV1EventsMediaPlayed**](EventsAPI.md#PostMediaBridgeV1EventsMediaPlayed) | **Post** /media-bridge/v1/events/media-played | 
[**PostMediaBridgeV1EventsMediaPlayedPercent**](EventsAPI.md#PostMediaBridgeV1EventsMediaPlayedPercent) | **Post** /media-bridge/v1/events/media-played-percent | 



## PostMediaBridgeV1EventsAttentionSpan

> AttentionSpanEvent PostMediaBridgeV1EventsAttentionSpan(ctx).Body(body).Execute()



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.PostMediaBridgeV1EventsAttentionSpan(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostMediaBridgeV1EventsAttentionSpan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1EventsAttentionSpan`: AttentionSpanEvent
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.PostMediaBridgeV1EventsAttentionSpan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1EventsAttentionSpanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**AttentionSpanEvent**](AttentionSpanEvent.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1EventsMediaPlayed

> MediaPlayedEvent PostMediaBridgeV1EventsMediaPlayed(ctx).Body(body).Execute()



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.PostMediaBridgeV1EventsMediaPlayed(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostMediaBridgeV1EventsMediaPlayed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1EventsMediaPlayed`: MediaPlayedEvent
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.PostMediaBridgeV1EventsMediaPlayed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1EventsMediaPlayedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**MediaPlayedEvent**](MediaPlayedEvent.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1EventsMediaPlayedPercent

> MediaPlayedPercentageEvent PostMediaBridgeV1EventsMediaPlayedPercent(ctx).Body(body).Execute()



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.PostMediaBridgeV1EventsMediaPlayedPercent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostMediaBridgeV1EventsMediaPlayedPercent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1EventsMediaPlayedPercent`: MediaPlayedPercentageEvent
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.PostMediaBridgeV1EventsMediaPlayedPercent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1EventsMediaPlayedPercentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**MediaPlayedPercentageEvent**](MediaPlayedPercentageEvent.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

