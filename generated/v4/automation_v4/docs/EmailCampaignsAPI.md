# \EmailCampaignsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutomationV4FlowsEmailCampaigns**](EmailCampaignsAPI.md#GetAutomationV4FlowsEmailCampaigns) | **Get** /automation/v4/flows/email-campaigns | Retrieve workflow emails



## GetAutomationV4FlowsEmailCampaigns

> CollectionResponseApiFlowEmailCampaign GetAutomationV4FlowsEmailCampaigns(ctx).After(after).Before(before).Limit(limit).FlowId(flowId).Execute()

Retrieve workflow emails



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
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	flowId := []string{"Inner_example"} // []string | The ID of the workflow. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailCampaignsAPI.GetAutomationV4FlowsEmailCampaigns(context.Background()).After(after).Before(before).Limit(limit).FlowId(flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailCampaignsAPI.GetAutomationV4FlowsEmailCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4FlowsEmailCampaigns`: CollectionResponseApiFlowEmailCampaign
	fmt.Fprintf(os.Stdout, "Response from `EmailCampaignsAPI.GetAutomationV4FlowsEmailCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4FlowsEmailCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **flowId** | **[]string** | The ID of the workflow. | 

### Return type

[**CollectionResponseApiFlowEmailCampaign**](CollectionResponseApiFlowEmailCampaign.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

