# \WorkflowIDMappingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAutomationV4WorkflowIdMappingsBatchRead**](WorkflowIDMappingsAPI.md#PostAutomationV4WorkflowIdMappingsBatchRead) | **Post** /automation/v4/workflow-id-mappings/batch/read | Retrieve migrated workflow mappings



## PostAutomationV4WorkflowIdMappingsBatchRead

> BatchResponseFlowIdWorkflowIdMappingResponse PostAutomationV4WorkflowIdMappingsBatchRead(ctx).ApiFlowBatchMigrationInput(apiFlowBatchMigrationInput).Execute()

Retrieve migrated workflow mappings



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
	apiFlowBatchMigrationInput := *openapiclient.NewApiFlowBatchMigrationInput([]openapiclient.ApiFlowBatchMigrationInputInputsInner{openapiclient.ApiFlowBatchMigrationInput_inputs_inner{ApiFlowBatchFetchMigrationFlowIdCoordinate: openapiclient.NewApiFlowBatchFetchMigrationFlowIdCoordinate("FlowMigrationStatuses_example", "Type_example")}}) // ApiFlowBatchMigrationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowIDMappingsAPI.PostAutomationV4WorkflowIdMappingsBatchRead(context.Background()).ApiFlowBatchMigrationInput(apiFlowBatchMigrationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowIDMappingsAPI.PostAutomationV4WorkflowIdMappingsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationV4WorkflowIdMappingsBatchRead`: BatchResponseFlowIdWorkflowIdMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowIDMappingsAPI.PostAutomationV4WorkflowIdMappingsBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationV4WorkflowIdMappingsBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiFlowBatchMigrationInput** | [**ApiFlowBatchMigrationInput**](ApiFlowBatchMigrationInput.md) |  | 

### Return type

[**BatchResponseFlowIdWorkflowIdMappingResponse**](BatchResponseFlowIdWorkflowIdMappingResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

