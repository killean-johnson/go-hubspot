# \ReportAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV4AssociationsUsageHighUsageReportUserIdRequest**](ReportAPI.md#PostCrmV4AssociationsUsageHighUsageReportUserIdRequest) | **Post** /crm/v4/associations/usage/high-usage-report/{userId} | Report



## PostCrmV4AssociationsUsageHighUsageReportUserIdRequest

> ReportCreationResponse PostCrmV4AssociationsUsageHighUsageReportUserIdRequest(ctx, userId).Execute()

Report



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
	userId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportAPI.PostCrmV4AssociationsUsageHighUsageReportUserIdRequest(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.PostCrmV4AssociationsUsageHighUsageReportUserIdRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV4AssociationsUsageHighUsageReportUserIdRequest`: ReportCreationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportAPI.PostCrmV4AssociationsUsageHighUsageReportUserIdRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsUsageHighUsageReportUserIdRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportCreationResponse**](ReportCreationResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

