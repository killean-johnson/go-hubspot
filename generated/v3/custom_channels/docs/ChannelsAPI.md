# \ChannelsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConversationsV3CustomChannelsChannelId**](ChannelsAPI.md#DeleteConversationsV3CustomChannelsChannelId) | **Delete** /conversations/v3/custom-channels/{channelId} | Archive a custom channel
[**GetConversationsV3CustomChannels**](ChannelsAPI.md#GetConversationsV3CustomChannels) | **Get** /conversations/v3/custom-channels/ | Get all channels for app
[**GetConversationsV3CustomChannelsChannelId**](ChannelsAPI.md#GetConversationsV3CustomChannelsChannelId) | **Get** /conversations/v3/custom-channels/{channelId} | Get a custom channel
[**PatchConversationsV3CustomChannelsChannelId**](ChannelsAPI.md#PatchConversationsV3CustomChannelsChannelId) | **Patch** /conversations/v3/custom-channels/{channelId} | Update a custom channel
[**PostConversationsV3CustomChannels**](ChannelsAPI.md#PostConversationsV3CustomChannels) | **Post** /conversations/v3/custom-channels/ | Create a custom channel



## DeleteConversationsV3CustomChannelsChannelId

> DeleteConversationsV3CustomChannelsChannelId(ctx, channelId).Execute()

Archive a custom channel



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
	channelId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChannelsAPI.DeleteConversationsV3CustomChannelsChannelId(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.DeleteConversationsV3CustomChannelsChannelId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConversationsV3CustomChannelsChannelIdRequest struct via the builder pattern


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


## GetConversationsV3CustomChannels

> CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging GetConversationsV3CustomChannels(ctx).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Execute()

Get all channels for app

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
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	sort := []string{"Inner_example"} // []string |  (optional)
	defaultPageLength := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetConversationsV3CustomChannels(context.Background()).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetConversationsV3CustomChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3CustomChannels`: CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetConversationsV3CustomChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3CustomChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **sort** | **[]string** |  | 
 **defaultPageLength** | **int32** |  | 

### Return type

[**CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging**](CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3CustomChannelsChannelId

> PublicChannelIntegrationChannel GetConversationsV3CustomChannelsChannelId(ctx, channelId).Execute()

Get a custom channel



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
	channelId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetConversationsV3CustomChannelsChannelId(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetConversationsV3CustomChannelsChannelId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3CustomChannelsChannelId`: PublicChannelIntegrationChannel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetConversationsV3CustomChannelsChannelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3CustomChannelsChannelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicChannelIntegrationChannel**](PublicChannelIntegrationChannel.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConversationsV3CustomChannelsChannelId

> PublicChannelIntegrationChannel PatchConversationsV3CustomChannelsChannelId(ctx, channelId).PublicChannelIntegrationChannelPatch(publicChannelIntegrationChannelPatch).Execute()

Update a custom channel



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
	channelId := int32(56) // int32 | 
	publicChannelIntegrationChannelPatch := *openapiclient.NewPublicChannelIntegrationChannelPatch(map[string]interface{}(123), map[string]map[string]interface{}{"key": map[string]interface{}(123)}, map[string]interface{}(123)) // PublicChannelIntegrationChannelPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.PatchConversationsV3CustomChannelsChannelId(context.Background(), channelId).PublicChannelIntegrationChannelPatch(publicChannelIntegrationChannelPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.PatchConversationsV3CustomChannelsChannelId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConversationsV3CustomChannelsChannelId`: PublicChannelIntegrationChannel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.PatchConversationsV3CustomChannelsChannelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConversationsV3CustomChannelsChannelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicChannelIntegrationChannelPatch** | [**PublicChannelIntegrationChannelPatch**](PublicChannelIntegrationChannelPatch.md) |  | 

### Return type

[**PublicChannelIntegrationChannel**](PublicChannelIntegrationChannel.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConversationsV3CustomChannels

> PublicChannelIntegrationChannel PostConversationsV3CustomChannels(ctx).PublicChannelIntegrationChannelCreate(publicChannelIntegrationChannelCreate).Execute()

Create a custom channel



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
	publicChannelIntegrationChannelCreate := *openapiclient.NewPublicChannelIntegrationChannelCreate(map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Name_example") // PublicChannelIntegrationChannelCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.PostConversationsV3CustomChannels(context.Background()).PublicChannelIntegrationChannelCreate(publicChannelIntegrationChannelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.PostConversationsV3CustomChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConversationsV3CustomChannels`: PublicChannelIntegrationChannel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.PostConversationsV3CustomChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostConversationsV3CustomChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicChannelIntegrationChannelCreate** | [**PublicChannelIntegrationChannelCreate**](PublicChannelIntegrationChannelCreate.md) |  | 

### Return type

[**PublicChannelIntegrationChannel**](PublicChannelIntegrationChannel.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

