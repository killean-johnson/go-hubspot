# \ChannelAccountStagingTokensAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchConversationsV3CustomChannelsChannelIdChannelAccountStagingTokensAccountToken**](ChannelAccountStagingTokensAPI.md#PatchConversationsV3CustomChannelsChannelIdChannelAccountStagingTokensAccountToken) | **Patch** /conversations/v3/custom-channels/{channelId}/channel-account-staging-tokens/{accountToken} | Update a channel account staging token



## PatchConversationsV3CustomChannelsChannelIdChannelAccountStagingTokensAccountToken

> PublicChannelAccountStagingToken PatchConversationsV3CustomChannelsChannelIdChannelAccountStagingTokensAccountToken(ctx, channelId, accountToken).PublicChannelAccountStagingTokenUpdateRequest(publicChannelAccountStagingTokenUpdateRequest).Execute()

Update a channel account staging token



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
	accountToken := "accountToken_example" // string | 
	publicChannelAccountStagingTokenUpdateRequest := *openapiclient.NewPublicChannelAccountStagingTokenUpdateRequest("AccountName_example", *openapiclient.NewPublicDeliveryIdentifier("Type_example", "Value_example")) // PublicChannelAccountStagingTokenUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAccountStagingTokensAPI.PatchConversationsV3CustomChannelsChannelIdChannelAccountStagingTokensAccountToken(context.Background(), channelId, accountToken).PublicChannelAccountStagingTokenUpdateRequest(publicChannelAccountStagingTokenUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAccountStagingTokensAPI.PatchConversationsV3CustomChannelsChannelIdChannelAccountStagingTokensAccountToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConversationsV3CustomChannelsChannelIdChannelAccountStagingTokensAccountToken`: PublicChannelAccountStagingToken
	fmt.Fprintf(os.Stdout, "Response from `ChannelAccountStagingTokensAPI.PatchConversationsV3CustomChannelsChannelIdChannelAccountStagingTokensAccountToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** |  | 
**accountToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConversationsV3CustomChannelsChannelIdChannelAccountStagingTokensAccountTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicChannelAccountStagingTokenUpdateRequest** | [**PublicChannelAccountStagingTokenUpdateRequest**](PublicChannelAccountStagingTokenUpdateRequest.md) |  | 

### Return type

[**PublicChannelAccountStagingToken**](PublicChannelAccountStagingToken.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

