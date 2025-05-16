# \PublicSequencesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutomationV4Sequences**](PublicSequencesAPI.md#GetAutomationV4Sequences) | **Get** /automation/v4/sequences/ | 
[**GetAutomationV4SequencesSequenceId**](PublicSequencesAPI.md#GetAutomationV4SequencesSequenceId) | **Get** /automation/v4/sequences/{sequenceId} | 



## GetAutomationV4Sequences

> CollectionResponseWithTotalPublicSequenceLiteResponseForwardPaging GetAutomationV4Sequences(ctx).UserId(userId).After(after).Limit(limit).Name(name).Execute()



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
	userId := "userId_example" // string | 
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicSequencesAPI.GetAutomationV4Sequences(context.Background()).UserId(userId).After(after).Limit(limit).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicSequencesAPI.GetAutomationV4Sequences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4Sequences`: CollectionResponseWithTotalPublicSequenceLiteResponseForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `PublicSequencesAPI.GetAutomationV4Sequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4SequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | 
 **name** | **string** |  | 

### Return type

[**CollectionResponseWithTotalPublicSequenceLiteResponseForwardPaging**](CollectionResponseWithTotalPublicSequenceLiteResponseForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4SequencesSequenceId

> PublicSequenceResponse GetAutomationV4SequencesSequenceId(ctx, sequenceId).UserId(userId).Execute()



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
	sequenceId := "sequenceId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicSequencesAPI.GetAutomationV4SequencesSequenceId(context.Background(), sequenceId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicSequencesAPI.GetAutomationV4SequencesSequenceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4SequencesSequenceId`: PublicSequenceResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicSequencesAPI.GetAutomationV4SequencesSequenceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4SequencesSequenceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** |  | 

### Return type

[**PublicSequenceResponse**](PublicSequenceResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

