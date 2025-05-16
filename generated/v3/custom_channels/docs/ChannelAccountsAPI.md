# \ChannelAccountsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConversationsV3CustomChannelsChannelIdChannelAccounts**](ChannelAccountsAPI.md#GetConversationsV3CustomChannelsChannelIdChannelAccounts) | **Get** /conversations/v3/custom-channels/{channelId}/channel-accounts | Query channel accounts
[**GetConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId**](ChannelAccountsAPI.md#GetConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId) | **Get** /conversations/v3/custom-channels/{channelId}/channel-accounts/{channelAccountId} | Get a channel account by id
[**PatchConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId**](ChannelAccountsAPI.md#PatchConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId) | **Patch** /conversations/v3/custom-channels/{channelId}/channel-accounts/{channelAccountId} | Update a channel account
[**PostConversationsV3CustomChannelsChannelIdChannelAccounts**](ChannelAccountsAPI.md#PostConversationsV3CustomChannelsChannelIdChannelAccounts) | **Post** /conversations/v3/custom-channels/{channelId}/channel-accounts | Create a channel account



## GetConversationsV3CustomChannelsChannelIdChannelAccounts

> CollectionResponseWithTotalPublicChannelAccountForwardPaging GetConversationsV3CustomChannelsChannelIdChannelAccounts(ctx, channelId).DeliveryIdentifierType(deliveryIdentifierType).DeliveryIdentifierValue(deliveryIdentifierValue).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Archived(archived).Execute()

Query channel accounts

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
	deliveryIdentifierType := []string{"Inner_example"} // []string |  (optional)
	deliveryIdentifierValue := []string{"Inner_example"} // []string |  (optional)
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	sort := []string{"Inner_example"} // []string |  (optional)
	defaultPageLength := int32(56) // int32 |  (optional)
	archived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAccountsAPI.GetConversationsV3CustomChannelsChannelIdChannelAccounts(context.Background(), channelId).DeliveryIdentifierType(deliveryIdentifierType).DeliveryIdentifierValue(deliveryIdentifierValue).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAccountsAPI.GetConversationsV3CustomChannelsChannelIdChannelAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3CustomChannelsChannelIdChannelAccounts`: CollectionResponseWithTotalPublicChannelAccountForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ChannelAccountsAPI.GetConversationsV3CustomChannelsChannelIdChannelAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3CustomChannelsChannelIdChannelAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deliveryIdentifierType** | **[]string** |  | 
 **deliveryIdentifierValue** | **[]string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 
 **defaultPageLength** | **int32** |  | 
 **archived** | **bool** |  | 

### Return type

[**CollectionResponseWithTotalPublicChannelAccountForwardPaging**](CollectionResponseWithTotalPublicChannelAccountForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId

> PublicChannelAccount GetConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId(ctx, channelId, channelAccountId).Archived(archived).Execute()

Get a channel account by id



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
	channelAccountId := int64(789) // int64 | 
	archived := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAccountsAPI.GetConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId(context.Background(), channelId, channelAccountId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAccountsAPI.GetConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId`: PublicChannelAccount
	fmt.Fprintf(os.Stdout, "Response from `ChannelAccountsAPI.GetConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** |  | 
**channelAccountId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** |  | [default to false]

### Return type

[**PublicChannelAccount**](PublicChannelAccount.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId

> PublicChannelAccount PatchConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId(ctx, channelId, channelAccountId).PublicChannelAccountUpdateRequest(publicChannelAccountUpdateRequest).Execute()

Update a channel account



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
	channelId := int32(56) // int32 | The channel to update
	channelAccountId := int64(789) // int64 | The channel account to update
	publicChannelAccountUpdateRequest := *openapiclient.NewPublicChannelAccountUpdateRequest() // PublicChannelAccountUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAccountsAPI.PatchConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId(context.Background(), channelId, channelAccountId).PublicChannelAccountUpdateRequest(publicChannelAccountUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAccountsAPI.PatchConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId`: PublicChannelAccount
	fmt.Fprintf(os.Stdout, "Response from `ChannelAccountsAPI.PatchConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** | The channel to update | 
**channelAccountId** | **int64** | The channel account to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConversationsV3CustomChannelsChannelIdChannelAccountsChannelAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicChannelAccountUpdateRequest** | [**PublicChannelAccountUpdateRequest**](PublicChannelAccountUpdateRequest.md) |  | 

### Return type

[**PublicChannelAccount**](PublicChannelAccount.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConversationsV3CustomChannelsChannelIdChannelAccounts

> PublicChannelAccount PostConversationsV3CustomChannelsChannelIdChannelAccounts(ctx, channelId).PublicChannelAccountEgg(publicChannelAccountEgg).Execute()

Create a channel account



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
	publicChannelAccountEgg := *openapiclient.NewPublicChannelAccountEgg(false, "Name_example", "InboxId_example") // PublicChannelAccountEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAccountsAPI.PostConversationsV3CustomChannelsChannelIdChannelAccounts(context.Background(), channelId).PublicChannelAccountEgg(publicChannelAccountEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAccountsAPI.PostConversationsV3CustomChannelsChannelIdChannelAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConversationsV3CustomChannelsChannelIdChannelAccounts`: PublicChannelAccount
	fmt.Fprintf(os.Stdout, "Response from `ChannelAccountsAPI.PostConversationsV3CustomChannelsChannelIdChannelAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConversationsV3CustomChannelsChannelIdChannelAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicChannelAccountEgg** | [**PublicChannelAccountEgg**](PublicChannelAccountEgg.md) |  | 

### Return type

[**PublicChannelAccount**](PublicChannelAccount.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

