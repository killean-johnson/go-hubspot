# \SingleSendAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV4EmailSingleSend**](SingleSendAPI.md#PostMarketingV4EmailSingleSend) | **Post** /marketing/v4/email/single-send | 



## PostMarketingV4EmailSingleSend

> EmailSendStatusView PostMarketingV4EmailSingleSend(ctx).PublicSingleSendRequestEgg(publicSingleSendRequestEgg).Execute()



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
	publicSingleSendRequestEgg := *openapiclient.NewPublicSingleSendRequestEgg(int64(123), *openapiclient.NewPublicSingleSendEmail("To_example")) // PublicSingleSendRequestEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SingleSendAPI.PostMarketingV4EmailSingleSend(context.Background()).PublicSingleSendRequestEgg(publicSingleSendRequestEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SingleSendAPI.PostMarketingV4EmailSingleSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV4EmailSingleSend`: EmailSendStatusView
	fmt.Fprintf(os.Stdout, "Response from `SingleSendAPI.PostMarketingV4EmailSingleSend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV4EmailSingleSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicSingleSendRequestEgg** | [**PublicSingleSendRequestEgg**](PublicSingleSendRequestEgg.md) |  | 

### Return type

[**EmailSendStatusView**](EmailSendStatusView.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

