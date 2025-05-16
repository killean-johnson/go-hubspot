# \MessagesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConversationsV3CustomChannelsChannelIdMessagesMessageId**](MessagesAPI.md#GetConversationsV3CustomChannelsChannelIdMessagesMessageId) | **Get** /conversations/v3/custom-channels/{channelId}/messages/{messageId} | Get a message
[**PatchConversationsV3CustomChannelsChannelIdMessagesMessageId**](MessagesAPI.md#PatchConversationsV3CustomChannelsChannelIdMessagesMessageId) | **Patch** /conversations/v3/custom-channels/{channelId}/messages/{messageId} | Update message 
[**PostConversationsV3CustomChannelsChannelIdMessages**](MessagesAPI.md#PostConversationsV3CustomChannelsChannelIdMessages) | **Post** /conversations/v3/custom-channels/{channelId}/messages | Publish a message



## GetConversationsV3CustomChannelsChannelIdMessagesMessageId

> PublicConversationsMessage GetConversationsV3CustomChannelsChannelIdMessagesMessageId(ctx, channelId, messageId).Execute()

Get a message



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
	channelId := int32(56) // int32 | The channel the message was sent over
	messageId := "messageId_example" // string | The id of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.GetConversationsV3CustomChannelsChannelIdMessagesMessageId(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetConversationsV3CustomChannelsChannelIdMessagesMessageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3CustomChannelsChannelIdMessagesMessageId`: PublicConversationsMessage
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetConversationsV3CustomChannelsChannelIdMessagesMessageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** | The channel the message was sent over | 
**messageId** | **string** | The id of the message | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3CustomChannelsChannelIdMessagesMessageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PublicConversationsMessage**](PublicConversationsMessage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConversationsV3CustomChannelsChannelIdMessagesMessageId

> PublicConversationsMessage PatchConversationsV3CustomChannelsChannelIdMessagesMessageId(ctx, channelId, messageId).PublicChannelIntegrationMessageUpdateRequest(publicChannelIntegrationMessageUpdateRequest).Execute()

Update message 



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
	channelId := int32(56) // int32 | The channel the message was sent over
	messageId := "messageId_example" // string | The id of the message
	publicChannelIntegrationMessageUpdateRequest := *openapiclient.NewPublicChannelIntegrationMessageUpdateRequest("StatusType_example") // PublicChannelIntegrationMessageUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.PatchConversationsV3CustomChannelsChannelIdMessagesMessageId(context.Background(), channelId, messageId).PublicChannelIntegrationMessageUpdateRequest(publicChannelIntegrationMessageUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.PatchConversationsV3CustomChannelsChannelIdMessagesMessageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConversationsV3CustomChannelsChannelIdMessagesMessageId`: PublicConversationsMessage
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.PatchConversationsV3CustomChannelsChannelIdMessagesMessageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** | The channel the message was sent over | 
**messageId** | **string** | The id of the message | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConversationsV3CustomChannelsChannelIdMessagesMessageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicChannelIntegrationMessageUpdateRequest** | [**PublicChannelIntegrationMessageUpdateRequest**](PublicChannelIntegrationMessageUpdateRequest.md) |  | 

### Return type

[**PublicConversationsMessage**](PublicConversationsMessage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConversationsV3CustomChannelsChannelIdMessages

> PublicConversationsMessage PostConversationsV3CustomChannelsChannelIdMessages(ctx, channelId).ChannelIntegrationMessageEgg(channelIntegrationMessageEgg).Execute()

Publish a message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	channelId := int32(56) // int32 | The channel the message will be sent over
	channelIntegrationMessageEgg := *openapiclient.NewChannelIntegrationMessageEgg("MessageDirection_example", []openapiclient.ChannelIntegrationMessageEggAttachmentsInner{openapiclient.ChannelIntegrationMessageEgg_attachments_inner{ContactAttachment: openapiclient.NewContactAttachment("Type_example", *openapiclient.NewContactProfile([]openapiclient.ContactEmail{*openapiclient.NewContactEmail("Email_example")}, []openapiclient.ContactAddress{*openapiclient.NewContactAddress()}, []openapiclient.ContactUrl{*openapiclient.NewContactUrl("Url_example")}, []openapiclient.ContactPhone{*openapiclient.NewContactPhone("Phone_example")}))}}, []openapiclient.ChannelIntegrationParticipant{*openapiclient.NewChannelIntegrationParticipant(*openapiclient.NewPublicDeliveryIdentifier("Type_example", "Value_example"))}, "IntegrationThreadId_example", "Text_example", "ChannelAccountId_example", []openapiclient.ChannelIntegrationParticipant{*openapiclient.NewChannelIntegrationParticipant(*openapiclient.NewPublicDeliveryIdentifier("Type_example", "Value_example"))}, time.Now()) // ChannelIntegrationMessageEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.PostConversationsV3CustomChannelsChannelIdMessages(context.Background(), channelId).ChannelIntegrationMessageEgg(channelIntegrationMessageEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.PostConversationsV3CustomChannelsChannelIdMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConversationsV3CustomChannelsChannelIdMessages`: PublicConversationsMessage
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.PostConversationsV3CustomChannelsChannelIdMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** | The channel the message will be sent over | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConversationsV3CustomChannelsChannelIdMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelIntegrationMessageEgg** | [**ChannelIntegrationMessageEgg**](ChannelIntegrationMessageEgg.md) |  | 

### Return type

[**PublicConversationsMessage**](PublicConversationsMessage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

