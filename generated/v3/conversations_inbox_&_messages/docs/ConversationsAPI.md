# \ConversationsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConversationsV3ConversationsThreadsThreadId**](ConversationsAPI.md#DeleteConversationsV3ConversationsThreadsThreadId) | **Delete** /conversations/v3/conversations/threads/{threadId} | Archives a thread.
[**GetConversationsV3ConversationsActorsActorId**](ConversationsAPI.md#GetConversationsV3ConversationsActorsActorId) | **Get** /conversations/v3/conversations/actors/{actorId} | Get a single actor.
[**GetConversationsV3ConversationsChannelAccounts**](ConversationsAPI.md#GetConversationsV3ConversationsChannelAccounts) | **Get** /conversations/v3/conversations/channel-accounts | Get channel accounts.
[**GetConversationsV3ConversationsChannelAccountsChannelAccountId**](ConversationsAPI.md#GetConversationsV3ConversationsChannelAccountsChannelAccountId) | **Get** /conversations/v3/conversations/channel-accounts/{channelAccountId} | Get a single channel account.
[**GetConversationsV3ConversationsChannels**](ConversationsAPI.md#GetConversationsV3ConversationsChannels) | **Get** /conversations/v3/conversations/channels | Get channels.
[**GetConversationsV3ConversationsChannelsChannelId**](ConversationsAPI.md#GetConversationsV3ConversationsChannelsChannelId) | **Get** /conversations/v3/conversations/channels/{channelId} | Get a single channel.
[**GetConversationsV3ConversationsInboxes**](ConversationsAPI.md#GetConversationsV3ConversationsInboxes) | **Get** /conversations/v3/conversations/inboxes | Get conversations inboxes.
[**GetConversationsV3ConversationsInboxesInboxId**](ConversationsAPI.md#GetConversationsV3ConversationsInboxesInboxId) | **Get** /conversations/v3/conversations/inboxes/{inboxId} | Get a single conversations inbox.
[**GetConversationsV3ConversationsThreads**](ConversationsAPI.md#GetConversationsV3ConversationsThreads) | **Get** /conversations/v3/conversations/threads | Get threads.
[**GetConversationsV3ConversationsThreadsThreadId**](ConversationsAPI.md#GetConversationsV3ConversationsThreadsThreadId) | **Get** /conversations/v3/conversations/threads/{threadId} | Get a single thread.
[**GetConversationsV3ConversationsThreadsThreadIdMessages**](ConversationsAPI.md#GetConversationsV3ConversationsThreadsThreadIdMessages) | **Get** /conversations/v3/conversations/threads/{threadId}/messages | Get message history for a thread.
[**GetConversationsV3ConversationsThreadsThreadIdMessagesMessageId**](ConversationsAPI.md#GetConversationsV3ConversationsThreadsThreadIdMessagesMessageId) | **Get** /conversations/v3/conversations/threads/{threadId}/messages/{messageId} | Get a single message.
[**GetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdOriginalContent**](ConversationsAPI.md#GetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdOriginalContent) | **Get** /conversations/v3/conversations/threads/{threadId}/messages/{messageId}/original-content | Get the original content of a single message.
[**PatchConversationsV3ConversationsThreadsThreadId**](ConversationsAPI.md#PatchConversationsV3ConversationsThreadsThreadId) | **Patch** /conversations/v3/conversations/threads/{threadId} | Update a thread.
[**PostConversationsV3ConversationsActorsBatchRead**](ConversationsAPI.md#PostConversationsV3ConversationsActorsBatchRead) | **Post** /conversations/v3/conversations/actors/batch/read | Get actors
[**PostConversationsV3ConversationsThreadsThreadIdMessages**](ConversationsAPI.md#PostConversationsV3ConversationsThreadsThreadIdMessages) | **Post** /conversations/v3/conversations/threads/{threadId}/messages | Send a message to a thread.



## DeleteConversationsV3ConversationsThreadsThreadId

> DeleteConversationsV3ConversationsThreadsThreadId(ctx, threadId).Execute()

Archives a thread.



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
	threadId := int64(789) // int64 | The unique ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConversationsAPI.DeleteConversationsV3ConversationsThreadsThreadId(context.Background(), threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.DeleteConversationsV3ConversationsThreadsThreadId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int64** | The unique ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConversationsV3ConversationsThreadsThreadIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsActorsActorId

> PublicActor GetConversationsV3ConversationsActorsActorId(ctx, actorId).Property(property).Execute()

Get a single actor.

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
	actorId := "actorId_example" // string | The unique ID of the actor.
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsActorsActorId(context.Background(), actorId).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsActorsActorId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsActorsActorId`: PublicActor
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsActorsActorId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | The unique ID of the actor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsActorsActorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **property** | **string** |  | 

### Return type

[**PublicActor**](PublicActor.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsChannelAccounts

> CollectionResponseWithTotalPublicChannelAccountForwardPaging GetConversationsV3ConversationsChannelAccounts(ctx).ChannelId(channelId).InboxId(inboxId).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Archived(archived).Execute()

Get channel accounts.

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
	channelId := []int32{int32(123)} // []int32 | Limits results to channel accounts within a particular channel. (optional)
	inboxId := []int32{int32(123)} // []int32 | Limits results to channel accounts within a particular inbox. (optional)
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	sort := []string{"Inner_example"} // []string |  (optional)
	defaultPageLength := int32(56) // int32 |  (optional)
	archived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsChannelAccounts(context.Background()).ChannelId(channelId).InboxId(inboxId).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsChannelAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsChannelAccounts`: CollectionResponseWithTotalPublicChannelAccountForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsChannelAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsChannelAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **[]int32** | Limits results to channel accounts within a particular channel. | 
 **inboxId** | **[]int32** | Limits results to channel accounts within a particular inbox. | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 
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


## GetConversationsV3ConversationsChannelAccountsChannelAccountId

> PublicChannelAccount GetConversationsV3ConversationsChannelAccountsChannelAccountId(ctx, channelAccountId).Archived(archived).Execute()

Get a single channel account.

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
	channelAccountId := int64(789) // int64 | The unique ID of the channel account.
	archived := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsChannelAccountsChannelAccountId(context.Background(), channelAccountId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsChannelAccountsChannelAccountId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsChannelAccountsChannelAccountId`: PublicChannelAccount
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsChannelAccountsChannelAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelAccountId** | **int64** | The unique ID of the channel account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsChannelAccountsChannelAccountIdRequest struct via the builder pattern


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


## GetConversationsV3ConversationsChannels

> CollectionResponseWithTotalPublicChannelForwardPaging GetConversationsV3ConversationsChannels(ctx).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Execute()

Get channels.

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
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsChannels(context.Background()).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsChannels`: CollectionResponseWithTotalPublicChannelForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **sort** | **[]string** |  | 
 **defaultPageLength** | **int32** |  | 

### Return type

[**CollectionResponseWithTotalPublicChannelForwardPaging**](CollectionResponseWithTotalPublicChannelForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsChannelsChannelId

> PublicChannel GetConversationsV3ConversationsChannelsChannelId(ctx, channelId).Execute()

Get a single channel.

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
	channelId := int32(56) // int32 | The unique ID of the channel.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsChannelsChannelId(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsChannelsChannelId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsChannelsChannelId`: PublicChannel
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsChannelsChannelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** | The unique ID of the channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsChannelsChannelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicChannel**](PublicChannel.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsInboxes

> CollectionResponseWithTotalPublicInboxForwardPaging GetConversationsV3ConversationsInboxes(ctx).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Archived(archived).Execute()

Get conversations inboxes.

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
	archived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsInboxes(context.Background()).After(after).Limit(limit).Sort(sort).DefaultPageLength(defaultPageLength).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsInboxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsInboxes`: CollectionResponseWithTotalPublicInboxForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsInboxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsInboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **sort** | **[]string** |  | 
 **defaultPageLength** | **int32** |  | 
 **archived** | **bool** |  | 

### Return type

[**CollectionResponseWithTotalPublicInboxForwardPaging**](CollectionResponseWithTotalPublicInboxForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsInboxesInboxId

> PublicInbox GetConversationsV3ConversationsInboxesInboxId(ctx, inboxId).Archived(archived).Execute()

Get a single conversations inbox.

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
	inboxId := int32(56) // int32 | The unique ID of the inbox.
	archived := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsInboxesInboxId(context.Background(), inboxId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsInboxesInboxId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsInboxesInboxId`: PublicInbox
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsInboxesInboxId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | **int32** | The unique ID of the inbox. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsInboxesInboxIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** |  | [default to false]

### Return type

[**PublicInbox**](PublicInbox.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsThreads

> CollectionResponsePublicThreadForwardPaging GetConversationsV3ConversationsThreads(ctx).After(after).Limit(limit).Sort(sort).InboxId(inboxId).AssociatedContactId(associatedContactId).ThreadStatus(threadStatus).LatestMessageTimestampAfter(latestMessageTimestampAfter).Archived(archived).Association(association).Property(property).Execute()

Get threads.

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
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	sort := []string{"Inner_example"} // []string | Set the sort order of the response. Valid options are `id` (default) and `latestMessageTimestamp` (which requires the `latestMessageTimestampAfter` field to also be set). If you’re filtering threads by `associatedContactId` , you can sort in descending order by prepending - to the sort option (e.g., `-id` or `-latestMessageTimestampAfter` ). Otherwise, results are always returned in ascending order. (optional)
	inboxId := []int32{int32(123)} // []int32 | The ID of the conversations inbox you can optionally include to retrieve the associated messages for. This parameter cannot be used in conjunction with the `associatedContactId` property. (optional)
	associatedContactId := int64(789) // int64 | The ID of a contact you can optionally include to retrieve a filtered list of conversations for. This parameter cannot be used in conjunction with the `inboxId` property. (optional)
	threadStatus := "threadStatus_example" // string | The status of the associated conversations to filter by (either `OPEN` or `CLOSED`). This property must be provided if you’re including the `associatedContactId` query parameter. (optional)
	latestMessageTimestampAfter := time.Now() // time.Time | The minimum `latestMessageTimestamp`. This is required only when sorting by `latestMessageTimestamp`. (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional)
	association := []string{"Association_example"} // []string | You can specify an association type here of `TICKET`. If this is set the response will included a thread associations object and associated ticket id if present. If there are no associations to a ticket with this conversation, then the thread associations object will not be present on the response.  (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsThreads(context.Background()).After(after).Limit(limit).Sort(sort).InboxId(inboxId).AssociatedContactId(associatedContactId).ThreadStatus(threadStatus).LatestMessageTimestampAfter(latestMessageTimestampAfter).Archived(archived).Association(association).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsThreads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsThreads`: CollectionResponsePublicThreadForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsThreads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **sort** | **[]string** | Set the sort order of the response. Valid options are &#x60;id&#x60; (default) and &#x60;latestMessageTimestamp&#x60; (which requires the &#x60;latestMessageTimestampAfter&#x60; field to also be set). If you’re filtering threads by &#x60;associatedContactId&#x60; , you can sort in descending order by prepending - to the sort option (e.g., &#x60;-id&#x60; or &#x60;-latestMessageTimestampAfter&#x60; ). Otherwise, results are always returned in ascending order. | 
 **inboxId** | **[]int32** | The ID of the conversations inbox you can optionally include to retrieve the associated messages for. This parameter cannot be used in conjunction with the &#x60;associatedContactId&#x60; property. | 
 **associatedContactId** | **int64** | The ID of a contact you can optionally include to retrieve a filtered list of conversations for. This parameter cannot be used in conjunction with the &#x60;inboxId&#x60; property. | 
 **threadStatus** | **string** | The status of the associated conversations to filter by (either &#x60;OPEN&#x60; or &#x60;CLOSED&#x60;). This property must be provided if you’re including the &#x60;associatedContactId&#x60; query parameter. | 
 **latestMessageTimestampAfter** | **time.Time** | The minimum &#x60;latestMessageTimestamp&#x60;. This is required only when sorting by &#x60;latestMessageTimestamp&#x60;. | 
 **archived** | **bool** | Whether to return only results that have been archived. | 
 **association** | **[]string** | You can specify an association type here of &#x60;TICKET&#x60;. If this is set the response will included a thread associations object and associated ticket id if present. If there are no associations to a ticket with this conversation, then the thread associations object will not be present on the response.  | 
 **property** | **string** |  | 

### Return type

[**CollectionResponsePublicThreadForwardPaging**](CollectionResponsePublicThreadForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsThreadsThreadId

> PublicThread GetConversationsV3ConversationsThreadsThreadId(ctx, threadId).Archived(archived).Association(association).Property(property).Execute()

Get a single thread.



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
	threadId := int64(789) // int64 | The unique ID of the thread.
	archived := true // bool | Whether to return only results that have been archived. Default is false. (optional)
	association := []string{"Association_example"} // []string | You can specify an association type here of `TICKET`. If this is set the response will included a thread associations object and associated ticket id if present. If there are no associations to a ticket with this conversation, then the thread associations object will not be present on the response.  (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsThreadsThreadId(context.Background(), threadId).Archived(archived).Association(association).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsThreadsThreadId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsThreadsThreadId`: PublicThread
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsThreadsThreadId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int64** | The unique ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsThreadsThreadIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. Default is false. | 
 **association** | **[]string** | You can specify an association type here of &#x60;TICKET&#x60;. If this is set the response will included a thread associations object and associated ticket id if present. If there are no associations to a ticket with this conversation, then the thread associations object will not be present on the response.  | 
 **property** | **string** |  | 

### Return type

[**PublicThread**](PublicThread.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsThreadsThreadIdMessages

> CollectionResponsePublicMessageForwardPaging GetConversationsV3ConversationsThreadsThreadIdMessages(ctx, threadId).After(after).Limit(limit).Sort(sort).Archived(archived).Property(property).Execute()

Get message history for a thread.

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
	threadId := int64(789) // int64 | The unique ID of the thread.
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	sort := []string{"Inner_example"} // []string | Sort direction. Valid options are `createdAt` (ascending), and `-createdAt` (descending, default) (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsThreadsThreadIdMessages(context.Background(), threadId).After(after).Limit(limit).Sort(sort).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsThreadsThreadIdMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsThreadsThreadIdMessages`: CollectionResponsePublicMessageForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsThreadsThreadIdMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int64** | The unique ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsThreadsThreadIdMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **sort** | **[]string** | Sort direction. Valid options are &#x60;createdAt&#x60; (ascending), and &#x60;-createdAt&#x60; (descending, default) | 
 **archived** | **bool** | Whether to return only results that have been archived. | 
 **property** | **string** |  | 

### Return type

[**CollectionResponsePublicMessageForwardPaging**](CollectionResponsePublicMessageForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsThreadsThreadIdMessagesMessageId

> PublicMessage GetConversationsV3ConversationsThreadsThreadIdMessagesMessageId(ctx, threadId, messageId).Property(property).Execute()

Get a single message.

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
	threadId := int64(789) // int64 | The unique ID of the thread.
	messageId := "messageId_example" // string | The unique ID of the message.
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsThreadsThreadIdMessagesMessageId(context.Background(), threadId, messageId).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsThreadsThreadIdMessagesMessageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsThreadsThreadIdMessagesMessageId`: PublicMessage
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsThreadsThreadIdMessagesMessageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int64** | The unique ID of the thread. | 
**messageId** | **string** | The unique ID of the message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **property** | **string** |  | 

### Return type

[**PublicMessage**](PublicMessage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdOriginalContent

> PublicMessageContent GetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdOriginalContent(ctx, threadId, messageId).Property(property).Execute()

Get the original content of a single message.



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
	threadId := int64(789) // int64 | The unique ID of the thread.
	messageId := "messageId_example" // string | The unique ID of the message.
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdOriginalContent(context.Background(), threadId, messageId).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdOriginalContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdOriginalContent`: PublicMessageContent
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdOriginalContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int64** | The unique ID of the thread. | 
**messageId** | **string** | The unique ID of the message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsV3ConversationsThreadsThreadIdMessagesMessageIdOriginalContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **property** | **string** |  | 

### Return type

[**PublicMessageContent**](PublicMessageContent.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConversationsV3ConversationsThreadsThreadId

> PublicThread PatchConversationsV3ConversationsThreadsThreadId(ctx, threadId).PublicThreadUpdateRequest(publicThreadUpdateRequest).Archived(archived).Execute()

Update a thread.



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
	threadId := int64(789) // int64 | The unique ID of the thread.
	publicThreadUpdateRequest := *openapiclient.NewPublicThreadUpdateRequest() // PublicThreadUpdateRequest | 
	archived := true // bool | Whether the thread to update is archived. Default is false. A thread's status property can not be updated if the thread is archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.PatchConversationsV3ConversationsThreadsThreadId(context.Background(), threadId).PublicThreadUpdateRequest(publicThreadUpdateRequest).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.PatchConversationsV3ConversationsThreadsThreadId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConversationsV3ConversationsThreadsThreadId`: PublicThread
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.PatchConversationsV3ConversationsThreadsThreadId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int64** | The unique ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConversationsV3ConversationsThreadsThreadIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicThreadUpdateRequest** | [**PublicThreadUpdateRequest**](PublicThreadUpdateRequest.md) |  | 
 **archived** | **bool** | Whether the thread to update is archived. Default is false. A thread&#39;s status property can not be updated if the thread is archived. | 

### Return type

[**PublicThread**](PublicThread.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConversationsV3ConversationsActorsBatchRead

> BatchResponsePublicActor PostConversationsV3ConversationsActorsBatchRead(ctx).BatchInputString(batchInputString).Property(property).Execute()

Get actors



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | 
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.PostConversationsV3ConversationsActorsBatchRead(context.Background()).BatchInputString(batchInputString).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.PostConversationsV3ConversationsActorsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConversationsV3ConversationsActorsBatchRead`: BatchResponsePublicActor
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.PostConversationsV3ConversationsActorsBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostConversationsV3ConversationsActorsBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) |  | 
 **property** | **string** |  | 

### Return type

[**BatchResponsePublicActor**](BatchResponsePublicActor.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConversationsV3ConversationsThreadsThreadIdMessages

> PublicMessage PostConversationsV3ConversationsThreadsThreadIdMessages(ctx, threadId).PublicMessageEgg(publicMessageEgg).Execute()

Send a message to a thread.



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
	threadId := int64(789) // int64 | The unique ID of the thread.
	publicMessageEgg := openapiclient.PublicMessageEgg{PublicCommentEgg: openapiclient.NewPublicCommentEgg("Type_example", "Text_example", []openapiclient.PublicCommentEggAllOfAttachments{openapiclient.PublicCommentEgg_allOf_attachments{PublicFileEgg: openapiclient.NewPublicFileEgg("Type_example", "FileId_example")}}, []openapiclient.PublicRecipientEgg{*openapiclient.NewPublicRecipientEgg([]openapiclient.PublicDeliveryIdentifier{*openapiclient.NewPublicDeliveryIdentifier("Type_example", "Value_example")})}, "SenderActorId_example", "ChannelId_example", "ChannelAccountId_example")} // PublicMessageEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.PostConversationsV3ConversationsThreadsThreadIdMessages(context.Background(), threadId).PublicMessageEgg(publicMessageEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.PostConversationsV3ConversationsThreadsThreadIdMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConversationsV3ConversationsThreadsThreadIdMessages`: PublicMessage
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.PostConversationsV3ConversationsThreadsThreadIdMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int64** | The unique ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConversationsV3ConversationsThreadsThreadIdMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicMessageEgg** | [**PublicMessageEgg**](PublicMessageEgg.md) |  | 

### Return type

[**PublicMessage**](PublicMessage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

