# \PublicEnrollmentsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutomationV4SequencesEnrollmentsContactContactId**](PublicEnrollmentsAPI.md#GetAutomationV4SequencesEnrollmentsContactContactId) | **Get** /automation/v4/sequences/enrollments/contact/{contactId} | 
[**PostAutomationV4SequencesEnrollments**](PublicEnrollmentsAPI.md#PostAutomationV4SequencesEnrollments) | **Post** /automation/v4/sequences/enrollments | 



## GetAutomationV4SequencesEnrollmentsContactContactId

> PublicSequenceEnrollmentResponse GetAutomationV4SequencesEnrollmentsContactContactId(ctx, contactId).Execute()



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
	contactId := "contactId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicEnrollmentsAPI.GetAutomationV4SequencesEnrollmentsContactContactId(context.Background(), contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicEnrollmentsAPI.GetAutomationV4SequencesEnrollmentsContactContactId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4SequencesEnrollmentsContactContactId`: PublicSequenceEnrollmentResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicEnrollmentsAPI.GetAutomationV4SequencesEnrollmentsContactContactId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4SequencesEnrollmentsContactContactIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicSequenceEnrollmentResponse**](PublicSequenceEnrollmentResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutomationV4SequencesEnrollments

> PublicSequenceEnrollmentLiteResponse PostAutomationV4SequencesEnrollments(ctx).UserId(userId).PublicSequenceEnrollmentRequest(publicSequenceEnrollmentRequest).Execute()



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
	publicSequenceEnrollmentRequest := *openapiclient.NewPublicSequenceEnrollmentRequest("ContactId_example", "SenderEmail_example", "SequenceId_example") // PublicSequenceEnrollmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicEnrollmentsAPI.PostAutomationV4SequencesEnrollments(context.Background()).UserId(userId).PublicSequenceEnrollmentRequest(publicSequenceEnrollmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicEnrollmentsAPI.PostAutomationV4SequencesEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationV4SequencesEnrollments`: PublicSequenceEnrollmentLiteResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicEnrollmentsAPI.PostAutomationV4SequencesEnrollments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationV4SequencesEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **publicSequenceEnrollmentRequest** | [**PublicSequenceEnrollmentRequest**](PublicSequenceEnrollmentRequest.md) |  | 

### Return type

[**PublicSequenceEnrollmentLiteResponse**](PublicSequenceEnrollmentLiteResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

