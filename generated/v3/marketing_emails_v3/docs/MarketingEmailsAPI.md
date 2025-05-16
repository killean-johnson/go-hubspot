# \MarketingEmailsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3EmailsEmailId**](MarketingEmailsAPI.md#DeleteMarketingV3EmailsEmailId) | **Delete** /marketing/v3/emails/{emailId} | Delete a marketing email.
[**GetMarketingV3Emails**](MarketingEmailsAPI.md#GetMarketingV3Emails) | **Get** /marketing/v3/emails/ | Get all marketing emails for a HubSpot account.
[**GetMarketingV3EmailsEmailId**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailId) | **Get** /marketing/v3/emails/{emailId} | Get the details of a specified marketing email.
[**GetMarketingV3EmailsEmailIdAbTestGetVariation**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailIdAbTestGetVariation) | **Get** /marketing/v3/emails/{emailId}/ab-test/get-variation | Get the variation of a an A/B marketing email
[**GetMarketingV3EmailsEmailIdDraft**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailIdDraft) | **Get** /marketing/v3/emails/{emailId}/draft | Get draft version of a marketing email
[**GetMarketingV3EmailsEmailIdRevisions**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailIdRevisions) | **Get** /marketing/v3/emails/{emailId}/revisions | Get revisions of a marketing email
[**GetMarketingV3EmailsEmailIdRevisionsRevisionId**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailIdRevisionsRevisionId) | **Get** /marketing/v3/emails/{emailId}/revisions/{revisionId} | Get a revision of a marketing email.
[**PatchMarketingV3EmailsEmailId**](MarketingEmailsAPI.md#PatchMarketingV3EmailsEmailId) | **Patch** /marketing/v3/emails/{emailId} | Update a marketing email.
[**PatchMarketingV3EmailsEmailIdDraft**](MarketingEmailsAPI.md#PatchMarketingV3EmailsEmailIdDraft) | **Patch** /marketing/v3/emails/{emailId}/draft | Create or update draft version
[**PostMarketingV3Emails**](MarketingEmailsAPI.md#PostMarketingV3Emails) | **Post** /marketing/v3/emails/ | Create a new marketing email.
[**PostMarketingV3EmailsEmailIdDraftReset**](MarketingEmailsAPI.md#PostMarketingV3EmailsEmailIdDraftReset) | **Post** /marketing/v3/emails/{emailId}/draft/reset | Reset Draft
[**PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestore**](MarketingEmailsAPI.md#PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestore) | **Post** /marketing/v3/emails/{emailId}/revisions/{revisionId}/restore | Restore a revision of a marketing email
[**PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraft**](MarketingEmailsAPI.md#PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraft) | **Post** /marketing/v3/emails/{emailId}/revisions/{revisionId}/restore-to-draft | Restore a revision of a marketing email to DRAFT state



## DeleteMarketingV3EmailsEmailId

> DeleteMarketingV3EmailsEmailId(ctx, emailId).Archived(archived).Execute()

Delete a marketing email.

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
	emailId := "emailId_example" // string | The ID of the marketing email to delete.
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingEmailsAPI.DeleteMarketingV3EmailsEmailId(context.Background(), emailId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.DeleteMarketingV3EmailsEmailId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The ID of the marketing email to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3EmailsEmailIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | 

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


## GetMarketingV3Emails

> CollectionResponseWithTotalPublicEmailForwardPaging GetMarketingV3Emails(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).IncludeStats(includeStats).Type_(type_).IsPublished(isPublished).IncludedProperties(includedProperties).Archived(archived).Execute()

Get all marketing emails for a HubSpot account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createdAt := time.Now() // time.Time | Only return emails created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return emails created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return emails created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return emails last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return emails last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return emails last updated before the specified time. (optional)
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)
	includeStats := true // bool | Include statistics with emails. (optional)
	type_ := "type__example" // string | Email types to be filtered by. Multiple types can be included. All emails will be returned if not present. (optional)
	isPublished := true // bool | Filter by published/draft emails. All emails will be returned if not present. (optional)
	includedProperties := []string{"Inner_example"} // []string |  (optional)
	archived := true // bool | Specifies whether to return archived emails. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3Emails(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).IncludeStats(includeStats).Type_(type_).IsPublished(isPublished).IncludedProperties(includedProperties).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3Emails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3Emails`: CollectionResponseWithTotalPublicEmailForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3Emails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return emails created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return emails created after the specified time. | 
 **createdBefore** | **time.Time** | Only return emails created before the specified time. | 
 **updatedAt** | **time.Time** | Only return emails last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return emails last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return emails last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 
 **includeStats** | **bool** | Include statistics with emails. | 
 **type_** | **string** | Email types to be filtered by. Multiple types can be included. All emails will be returned if not present. | 
 **isPublished** | **bool** | Filter by published/draft emails. All emails will be returned if not present. | 
 **includedProperties** | **[]string** |  | 
 **archived** | **bool** | Specifies whether to return archived emails. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalPublicEmailForwardPaging**](CollectionResponseWithTotalPublicEmailForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3EmailsEmailId

> PublicEmail GetMarketingV3EmailsEmailId(ctx, emailId).IncludeStats(includeStats).IncludedProperties(includedProperties).Archived(archived).Execute()

Get the details of a specified marketing email.



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
	emailId := "emailId_example" // string | The marketing email ID.
	includeStats := true // bool | Include statistics with email (optional)
	includedProperties := []string{"Inner_example"} // []string |  (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailId(context.Background(), emailId).IncludeStats(includeStats).IncludedProperties(includedProperties).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailId`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeStats** | **bool** | Include statistics with email | 
 **includedProperties** | **[]string** |  | 
 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

[**PublicEmail**](PublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3EmailsEmailIdAbTestGetVariation

> PublicEmail GetMarketingV3EmailsEmailIdAbTestGetVariation(ctx, emailId).Execute()

Get the variation of a an A/B marketing email



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
	emailId := "emailId_example" // string | The ID of an A/B marketing email.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailIdAbTestGetVariation(context.Background(), emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdAbTestGetVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailIdAbTestGetVariation`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdAbTestGetVariation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The ID of an A/B marketing email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdAbTestGetVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicEmail**](PublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3EmailsEmailIdDraft

> PublicEmail GetMarketingV3EmailsEmailIdDraft(ctx, emailId).Execute()

Get draft version of a marketing email



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
	emailId := "emailId_example" // string | The marketing email ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailIdDraft(context.Background(), emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailIdDraft`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicEmail**](PublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3EmailsEmailIdRevisions

> CollectionResponseWithTotalVersionPublicEmail GetMarketingV3EmailsEmailIdRevisions(ctx, emailId).After(after).Before(before).Limit(limit).Execute()

Get revisions of a marketing email



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
	emailId := "emailId_example" // string | The marketing email ID.
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	before := "before_example" // string | The cursor token value to get the previous set of results. You can get this from the `paging.prev.before` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisions(context.Background(), emailId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailIdRevisions`: CollectionResponseWithTotalVersionPublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** | The cursor token value to get the previous set of results. You can get this from the &#x60;paging.prev.before&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 

### Return type

[**CollectionResponseWithTotalVersionPublicEmail**](CollectionResponseWithTotalVersionPublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3EmailsEmailIdRevisionsRevisionId

> VersionPublicEmail GetMarketingV3EmailsEmailIdRevisionsRevisionId(ctx, emailId, revisionId).Execute()

Get a revision of a marketing email.



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
	emailId := "emailId_example" // string | The marketing email ID.
	revisionId := "revisionId_example" // string | The ID of a revision.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisionsRevisionId(context.Background(), emailId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisionsRevisionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailIdRevisionsRevisionId`: VersionPublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisionsRevisionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 
**revisionId** | **string** | The ID of a revision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdRevisionsRevisionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionPublicEmail**](VersionPublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingV3EmailsEmailId

> PublicEmail PatchMarketingV3EmailsEmailId(ctx, emailId).EmailUpdateRequest(emailUpdateRequest).Archived(archived).Execute()

Update a marketing email.



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
	emailId := "emailId_example" // string | The ID of the marketing email that should get updated
	emailUpdateRequest := *openapiclient.NewEmailUpdateRequest() // EmailUpdateRequest | A marketing email object with properties that should overwrite the corresponding properties of the marketing email.
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.PatchMarketingV3EmailsEmailId(context.Background(), emailId).EmailUpdateRequest(emailUpdateRequest).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PatchMarketingV3EmailsEmailId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingV3EmailsEmailId`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PatchMarketingV3EmailsEmailId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The ID of the marketing email that should get updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3EmailsEmailIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailUpdateRequest** | [**EmailUpdateRequest**](EmailUpdateRequest.md) | A marketing email object with properties that should overwrite the corresponding properties of the marketing email. | 
 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

[**PublicEmail**](PublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingV3EmailsEmailIdDraft

> PublicEmail PatchMarketingV3EmailsEmailIdDraft(ctx, emailId).EmailUpdateRequest(emailUpdateRequest).Execute()

Create or update draft version



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
	emailId := "emailId_example" // string | The marketing email ID.
	emailUpdateRequest := *openapiclient.NewEmailUpdateRequest() // EmailUpdateRequest | A marketing email object with properties that should overwrite the corresponding properties in the email's current draft.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.PatchMarketingV3EmailsEmailIdDraft(context.Background(), emailId).EmailUpdateRequest(emailUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PatchMarketingV3EmailsEmailIdDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingV3EmailsEmailIdDraft`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PatchMarketingV3EmailsEmailIdDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3EmailsEmailIdDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailUpdateRequest** | [**EmailUpdateRequest**](EmailUpdateRequest.md) | A marketing email object with properties that should overwrite the corresponding properties in the email&#39;s current draft. | 

### Return type

[**PublicEmail**](PublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3Emails

> PublicEmail PostMarketingV3Emails(ctx).EmailCreateRequest(emailCreateRequest).Execute()

Create a new marketing email.



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
	emailCreateRequest := *openapiclient.NewEmailCreateRequest("Name_example") // EmailCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.PostMarketingV3Emails(context.Background()).EmailCreateRequest(emailCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3Emails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3Emails`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PostMarketingV3Emails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailCreateRequest** | [**EmailCreateRequest**](EmailCreateRequest.md) |  | 

### Return type

[**PublicEmail**](PublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3EmailsEmailIdDraftReset

> PostMarketingV3EmailsEmailIdDraftReset(ctx, emailId).Execute()

Reset Draft



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
	emailId := "emailId_example" // string | The marketing email ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsEmailIdDraftReset(context.Background(), emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdDraftReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsEmailIdDraftResetRequest struct via the builder pattern


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


## PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestore

> PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestore(ctx, emailId, revisionId).Execute()

Restore a revision of a marketing email



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
	emailId := "emailId_example" // string | The marketing email ID.
	revisionId := "revisionId_example" // string | The ID of a revision.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestore(context.Background(), emailId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 
**revisionId** | **string** | The ID of a revision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreRequest struct via the builder pattern


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


## PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraft

> PublicEmail PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraft(ctx, emailId, revisionId).Execute()

Restore a revision of a marketing email to DRAFT state



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
	emailId := "emailId_example" // string | The marketing email ID.
	revisionId := int64(789) // int64 | The ID of a revision.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraft(context.Background(), emailId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraft`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 
**revisionId** | **int64** | The ID of a revision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PublicEmail**](PublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

