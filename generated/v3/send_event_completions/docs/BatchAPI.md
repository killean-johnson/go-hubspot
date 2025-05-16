# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostEventsV3SendBatchSend**](BatchAPI.md#PostEventsV3SendBatchSend) | **Post** /events/v3/send/batch | Send a batch of event completions



## PostEventsV3SendBatchSend

> PostEventsV3SendBatchSend(ctx).BatchedBehavioralEventHttpCompletionRequest(batchedBehavioralEventHttpCompletionRequest).Execute()

Send a batch of event completions



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
	batchedBehavioralEventHttpCompletionRequest := *openapiclient.NewBatchedBehavioralEventHttpCompletionRequest([]openapiclient.BehavioralEventHttpCompletionRequest{*openapiclient.NewBehavioralEventHttpCompletionRequest("pe123456_account_login")}) // BatchedBehavioralEventHttpCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.PostEventsV3SendBatchSend(context.Background()).BatchedBehavioralEventHttpCompletionRequest(batchedBehavioralEventHttpCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostEventsV3SendBatchSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEventsV3SendBatchSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchedBehavioralEventHttpCompletionRequest** | [**BatchedBehavioralEventHttpCompletionRequest**](BatchedBehavioralEventHttpCompletionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

