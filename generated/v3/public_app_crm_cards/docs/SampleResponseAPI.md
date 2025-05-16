# \SampleResponseAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ExtensionsCardsDevSampleResponseGetCardsSampleResponse**](SampleResponseAPI.md#GetCrmV3ExtensionsCardsDevSampleResponseGetCardsSampleResponse) | **Get** /crm/v3/extensions/cards-dev/sample-response | Get sample card detail response



## GetCrmV3ExtensionsCardsDevSampleResponseGetCardsSampleResponse

> IntegratorCardPayloadResponse GetCrmV3ExtensionsCardsDevSampleResponseGetCardsSampleResponse(ctx).Execute()

Get sample card detail response



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SampleResponseAPI.GetCrmV3ExtensionsCardsDevSampleResponseGetCardsSampleResponse(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SampleResponseAPI.GetCrmV3ExtensionsCardsDevSampleResponseGetCardsSampleResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExtensionsCardsDevSampleResponseGetCardsSampleResponse`: IntegratorCardPayloadResponse
	fmt.Fprintf(os.Stdout, "Response from `SampleResponseAPI.GetCrmV3ExtensionsCardsDevSampleResponseGetCardsSampleResponse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCardsDevSampleResponseGetCardsSampleResponseRequest struct via the builder pattern


### Return type

[**IntegratorCardPayloadResponse**](IntegratorCardPayloadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

