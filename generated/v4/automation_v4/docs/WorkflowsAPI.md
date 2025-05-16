# \WorkflowsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutomationV4FlowsFlowId**](WorkflowsAPI.md#DeleteAutomationV4FlowsFlowId) | **Delete** /automation/v4/flows/{flowId} | Delete a workflow
[**GetAutomationV4Flows**](WorkflowsAPI.md#GetAutomationV4Flows) | **Get** /automation/v4/flows | Retrieve workflows
[**GetAutomationV4FlowsFlowId**](WorkflowsAPI.md#GetAutomationV4FlowsFlowId) | **Get** /automation/v4/flows/{flowId} | Retrieve workflow details
[**PostAutomationV4Flows**](WorkflowsAPI.md#PostAutomationV4Flows) | **Post** /automation/v4/flows | Create a new workflow
[**PostAutomationV4FlowsBatchRead**](WorkflowsAPI.md#PostAutomationV4FlowsBatchRead) | **Post** /automation/v4/flows/batch/read | Retrieve a batch of workflows
[**PutAutomationV4FlowsFlowId**](WorkflowsAPI.md#PutAutomationV4FlowsFlowId) | **Put** /automation/v4/flows/{flowId} | Update a workflow



## DeleteAutomationV4FlowsFlowId

> DeleteAutomationV4FlowsFlowId(ctx, flowId).Execute()

Delete a workflow



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
	flowId := int64(789) // int64 | The ID of the workflow to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.DeleteAutomationV4FlowsFlowId(context.Background(), flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.DeleteAutomationV4FlowsFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **int64** | The ID of the workflow to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationV4FlowsFlowIdRequest struct via the builder pattern


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


## GetAutomationV4Flows

> CollectionResponseApiFlowListingForwardPaging GetAutomationV4Flows(ctx).After(after).Limit(limit).Execute()

Retrieve workflows



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
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetAutomationV4Flows(context.Background()).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetAutomationV4Flows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4Flows`: CollectionResponseApiFlowListingForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetAutomationV4Flows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4FlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 100]

### Return type

[**CollectionResponseApiFlowListingForwardPaging**](CollectionResponseApiFlowListingForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4FlowsFlowId

> ApiFlow GetAutomationV4FlowsFlowId(ctx, flowId).Execute()

Retrieve workflow details



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
	flowId := "flowId_example" // string | The ID of the workflow to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetAutomationV4FlowsFlowId(context.Background(), flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetAutomationV4FlowsFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4FlowsFlowId`: ApiFlow
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetAutomationV4FlowsFlowId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | The ID of the workflow to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4FlowsFlowIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiFlow**](ApiFlow.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutomationV4Flows

> ApiFlow PostAutomationV4Flows(ctx).ApiFlowCreateRequest(apiFlowCreateRequest).Execute()

Create a new workflow



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
	apiFlowCreateRequest := openapiclient.ApiFlowCreateRequest{ApiContactFlowCreateRequest: openapiclient.NewApiContactFlowCreateRequest("Type_example", "ObjectTypeId_example", false, false, "FlowType_example", []openapiclient.ApiContactFlowCreateRequestAllOfActions{openapiclient.ApiContactFlowCreateRequest_allOf_actions{ApiABTestBranchAction: openapiclient.NewApiABTestBranchAction([]openapiclient.ApiConnection{*openapiclient.NewApiConnection("NextActionId_example", "EdgeType_example")}, "1", "AB_TEST_BRANCH")}}, []openapiclient.ApiTimeWindow{*openapiclient.NewApiTimeWindow(*openapiclient.NewApiTimeOfDay(int32(123), int32(123)), *openapiclient.NewApiTimeOfDay(int32(123), int32(123)), "Day_example")}, []openapiclient.ApiBlockedDate{*openapiclient.NewApiBlockedDate("Month_example", int32(123))}, map[string]string{"key": "Inner_example"}, []openapiclient.ApiContactFlowCreateRequestAllOfDataSources{openapiclient.ApiContactFlowCreateRequest_allOf_dataSources{ApiAssociationDataSource: openapiclient.NewApiAssociationDataSource("ObjectTypeId_example", "Name_example", int32(123), "AssociationCategory_example", "Type_example")}}, []int32{int32(123)})} // ApiFlowCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PostAutomationV4Flows(context.Background()).ApiFlowCreateRequest(apiFlowCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PostAutomationV4Flows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationV4Flows`: ApiFlow
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PostAutomationV4Flows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationV4FlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiFlowCreateRequest** | [**ApiFlowCreateRequest**](ApiFlowCreateRequest.md) |  | 

### Return type

[**ApiFlow**](ApiFlow.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutomationV4FlowsBatchRead

> BatchResponseApiFlow PostAutomationV4FlowsBatchRead(ctx).ApiFlowBatchInput(apiFlowBatchInput).Execute()

Retrieve a batch of workflows



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
	apiFlowBatchInput := *openapiclient.NewApiFlowBatchInput([]openapiclient.ApiFlowBatchInputInputsInner{openapiclient.ApiFlowBatchInput_inputs_inner{ApiFlowBatchFetchFlowIdCoordinate: openapiclient.NewApiFlowBatchFetchFlowIdCoordinate("Type_example", "FlowId_example")}}) // ApiFlowBatchInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PostAutomationV4FlowsBatchRead(context.Background()).ApiFlowBatchInput(apiFlowBatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PostAutomationV4FlowsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationV4FlowsBatchRead`: BatchResponseApiFlow
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PostAutomationV4FlowsBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationV4FlowsBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiFlowBatchInput** | [**ApiFlowBatchInput**](ApiFlowBatchInput.md) |  | 

### Return type

[**BatchResponseApiFlow**](BatchResponseApiFlow.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAutomationV4FlowsFlowId

> ApiFlow PutAutomationV4FlowsFlowId(ctx, flowId).ApiFlowPutRequest(apiFlowPutRequest).Execute()

Update a workflow



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
	flowId := "flowId_example" // string | The ID of the workflow.
	apiFlowPutRequest := openapiclient.ApiFlowPutRequest{ApiContactFlowPutRequest: openapiclient.NewApiContactFlowPutRequest("Type_example", "RevisionId_example", false, []openapiclient.ApiContactFlowPutRequestAllOfActions{openapiclient.ApiContactFlowPutRequest_allOf_actions{ApiABTestBranchAction: openapiclient.NewApiABTestBranchAction([]openapiclient.ApiConnection{*openapiclient.NewApiConnection("NextActionId_example", "EdgeType_example")}, "1", "AB_TEST_BRANCH")}}, []openapiclient.ApiTimeWindow{*openapiclient.NewApiTimeWindow(*openapiclient.NewApiTimeOfDay(int32(123), int32(123)), *openapiclient.NewApiTimeOfDay(int32(123), int32(123)), "Day_example")}, []openapiclient.ApiBlockedDate{*openapiclient.NewApiBlockedDate("Month_example", int32(123))}, map[string]string{"key": "Inner_example"}, []int32{int32(123)}, false)} // ApiFlowPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PutAutomationV4FlowsFlowId(context.Background(), flowId).ApiFlowPutRequest(apiFlowPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PutAutomationV4FlowsFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAutomationV4FlowsFlowId`: ApiFlow
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PutAutomationV4FlowsFlowId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | The ID of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAutomationV4FlowsFlowIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiFlowPutRequest** | [**ApiFlowPutRequest**](ApiFlowPutRequest.md) |  | 

### Return type

[**ApiFlow**](ApiFlow.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

