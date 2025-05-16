# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostEventsV3SendSend**](BasicAPI.md#PostEventsV3SendSend) | **Post** /events/v3/send | Send a custom event completion



## PostEventsV3SendSend

> PostEventsV3SendSend(ctx).BehavioralEventHttpCompletionRequest(behavioralEventHttpCompletionRequest).Execute()

Send a custom event completion



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
	behavioralEventHttpCompletionRequest := *openapiclient.NewBehavioralEventHttpCompletionRequest("pe123456_account_login") // BehavioralEventHttpCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.PostEventsV3SendSend(context.Background()).BehavioralEventHttpCompletionRequest(behavioralEventHttpCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostEventsV3SendSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEventsV3SendSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **behavioralEventHttpCompletionRequest** | [**BehavioralEventHttpCompletionRequest**](BehavioralEventHttpCompletionRequest.md) |  | 

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

