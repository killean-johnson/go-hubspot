# \TokensAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameArchive**](TokensAPI.md#DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameArchive) | **Delete** /integrators/timeline/v3/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Removes a token from the event template
[**PostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreate**](TokensAPI.md#PostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreate) | **Post** /integrators/timeline/v3/{appId}/event-templates/{eventTemplateId}/tokens | Adds a token to an existing event template
[**PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdate**](TokensAPI.md#PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdate) | **Put** /integrators/timeline/v3/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Updates an existing token on an event template



## DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameArchive

> DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameArchive(ctx, eventTemplateId, tokenName, appId).Execute()

Removes a token from the event template



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
	eventTemplateId := "eventTemplateId_example" // string | The event template ID.
	tokenName := "tokenName_example" // string | The token name.
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TokensAPI.DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameArchive(context.Background(), eventTemplateId, tokenName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**tokenName** | **string** | The token name. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreate

> TimelineEventTemplateToken PostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreate(ctx, eventTemplateId, appId).TimelineEventTemplateToken(timelineEventTemplateToken).Execute()

Adds a token to an existing event template



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
	eventTemplateId := "eventTemplateId_example" // string | The event template ID.
	appId := int32(56) // int32 | The ID of the target app.
	timelineEventTemplateToken := *openapiclient.NewTimelineEventTemplateToken("petType", "Pet Type", "enumeration") // TimelineEventTemplateToken | The new token definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.PostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreate(context.Background(), eventTemplateId, appId).TimelineEventTemplateToken(timelineEventTemplateToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.PostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreate`: TimelineEventTemplateToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.PostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timelineEventTemplateToken** | [**TimelineEventTemplateToken**](TimelineEventTemplateToken.md) | The new token definition. | 

### Return type

[**TimelineEventTemplateToken**](TimelineEventTemplateToken.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdate

> TimelineEventTemplateToken PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdate(ctx, eventTemplateId, tokenName, appId).TimelineEventTemplateTokenUpdateRequest(timelineEventTemplateTokenUpdateRequest).Execute()

Updates an existing token on an event template



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
	eventTemplateId := "eventTemplateId_example" // string | The event template ID.
	tokenName := "tokenName_example" // string | The token name.
	appId := int32(56) // int32 | The ID of the target app.
	timelineEventTemplateTokenUpdateRequest := *openapiclient.NewTimelineEventTemplateTokenUpdateRequest("petType edit") // TimelineEventTemplateTokenUpdateRequest | The updated token definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdate(context.Background(), eventTemplateId, tokenName, appId).TimelineEventTemplateTokenUpdateRequest(timelineEventTemplateTokenUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdate`: TimelineEventTemplateToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**tokenName** | **string** | The token name. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **timelineEventTemplateTokenUpdateRequest** | [**TimelineEventTemplateTokenUpdateRequest**](TimelineEventTemplateTokenUpdateRequest.md) | The updated token definition. | 

### Return type

[**TimelineEventTemplateToken**](TimelineEventTemplateToken.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

