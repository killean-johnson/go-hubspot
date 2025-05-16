# \MarketingEmailsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3EmailsEmailIdArchive**](MarketingEmailsAPI.md#DeleteMarketingV3EmailsEmailIdArchive) | **Delete** /marketing/v3/emails/{emailId} | Delete a marketing email.
[**GetMarketingV3EmailsEmailIdAbTestGetVariationGetAbTestVariation**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailIdAbTestGetVariationGetAbTestVariation) | **Get** /marketing/v3/emails/{emailId}/ab-test/get-variation | Get the variation of a an A/B marketing email
[**GetMarketingV3EmailsEmailIdDraftGetDraft**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailIdDraftGetDraft) | **Get** /marketing/v3/emails/{emailId}/draft | Get draft version of a marketing email
[**GetMarketingV3EmailsEmailIdGetById**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailIdGetById) | **Get** /marketing/v3/emails/{emailId} | Get the details of a specified marketing email.
[**GetMarketingV3EmailsEmailIdRevisionsGetRevisions**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailIdRevisionsGetRevisions) | **Get** /marketing/v3/emails/{emailId}/revisions | Get revisions of a marketing email
[**GetMarketingV3EmailsEmailIdRevisionsRevisionIdGetRevisionById**](MarketingEmailsAPI.md#GetMarketingV3EmailsEmailIdRevisionsRevisionIdGetRevisionById) | **Get** /marketing/v3/emails/{emailId}/revisions/{revisionId} | Get a revision of a marketing email.
[**GetMarketingV3EmailsGetPage**](MarketingEmailsAPI.md#GetMarketingV3EmailsGetPage) | **Get** /marketing/v3/emails/ | Get all marketing emails for a HubSpot account.
[**PatchMarketingV3EmailsEmailIdDraftUpsertDraft**](MarketingEmailsAPI.md#PatchMarketingV3EmailsEmailIdDraftUpsertDraft) | **Patch** /marketing/v3/emails/{emailId}/draft | Create or update draft version
[**PatchMarketingV3EmailsEmailIdUpdate**](MarketingEmailsAPI.md#PatchMarketingV3EmailsEmailIdUpdate) | **Patch** /marketing/v3/emails/{emailId} | Update a marketing email.
[**PostMarketingV3EmailsAbTestCreateVariationCreateAbTestVariation**](MarketingEmailsAPI.md#PostMarketingV3EmailsAbTestCreateVariationCreateAbTestVariation) | **Post** /marketing/v3/emails/ab-test/create-variation | Create an A/B test variation of a marketing email.
[**PostMarketingV3EmailsCloneClone**](MarketingEmailsAPI.md#PostMarketingV3EmailsCloneClone) | **Post** /marketing/v3/emails/clone | Clone a marketing email.
[**PostMarketingV3EmailsCreate**](MarketingEmailsAPI.md#PostMarketingV3EmailsCreate) | **Post** /marketing/v3/emails/ | Create a new marketing email.
[**PostMarketingV3EmailsEmailIdDraftResetResetDraft**](MarketingEmailsAPI.md#PostMarketingV3EmailsEmailIdDraftResetResetDraft) | **Post** /marketing/v3/emails/{emailId}/draft/reset | Reset Draft
[**PostMarketingV3EmailsEmailIdPublishPublishOrSend**](MarketingEmailsAPI.md#PostMarketingV3EmailsEmailIdPublishPublishOrSend) | **Post** /marketing/v3/emails/{emailId}/publish | Publish or send a marketing email.
[**PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreRestoreRevision**](MarketingEmailsAPI.md#PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreRestoreRevision) | **Post** /marketing/v3/emails/{emailId}/revisions/{revisionId}/restore | Restore a revision of a marketing email
[**PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRestoreDraftRevision**](MarketingEmailsAPI.md#PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRestoreDraftRevision) | **Post** /marketing/v3/emails/{emailId}/revisions/{revisionId}/restore-to-draft | Restore a revision of a marketing email to DRAFT state
[**PostMarketingV3EmailsEmailIdUnpublishUnpublishOrCancel**](MarketingEmailsAPI.md#PostMarketingV3EmailsEmailIdUnpublishUnpublishOrCancel) | **Post** /marketing/v3/emails/{emailId}/unpublish | Unpublish or cancel a marketing email.



## DeleteMarketingV3EmailsEmailIdArchive

> DeleteMarketingV3EmailsEmailIdArchive(ctx, emailId).Archived(archived).Execute()

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
	r, err := apiClient.MarketingEmailsAPI.DeleteMarketingV3EmailsEmailIdArchive(context.Background(), emailId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.DeleteMarketingV3EmailsEmailIdArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMarketingV3EmailsEmailIdArchiveRequest struct via the builder pattern


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


## GetMarketingV3EmailsEmailIdAbTestGetVariationGetAbTestVariation

> PublicEmail GetMarketingV3EmailsEmailIdAbTestGetVariationGetAbTestVariation(ctx, emailId).Execute()

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
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailIdAbTestGetVariationGetAbTestVariation(context.Background(), emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdAbTestGetVariationGetAbTestVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailIdAbTestGetVariationGetAbTestVariation`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdAbTestGetVariationGetAbTestVariation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The ID of an A/B marketing email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdAbTestGetVariationGetAbTestVariationRequest struct via the builder pattern


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


## GetMarketingV3EmailsEmailIdDraftGetDraft

> PublicEmail GetMarketingV3EmailsEmailIdDraftGetDraft(ctx, emailId).Execute()

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
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailIdDraftGetDraft(context.Background(), emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdDraftGetDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailIdDraftGetDraft`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdDraftGetDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdDraftGetDraftRequest struct via the builder pattern


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


## GetMarketingV3EmailsEmailIdGetById

> PublicEmail GetMarketingV3EmailsEmailIdGetById(ctx, emailId).IncludeStats(includeStats).MarketingCampaignNames(marketingCampaignNames).WorkflowNames(workflowNames).IncludedProperties(includedProperties).Archived(archived).Execute()

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
	includeStats := true // bool | Include statistics with email. (optional)
	marketingCampaignNames := true // bool | If set to true, loads `campaignName` and `campaignUtm`. (optional)
	workflowNames := true // bool | If set to true, loads workflows in which the email is used within a \"send email\" action.  (optional)
	includedProperties := []string{"Inner_example"} // []string |  (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailIdGetById(context.Background(), emailId).IncludeStats(includeStats).MarketingCampaignNames(marketingCampaignNames).WorkflowNames(workflowNames).IncludedProperties(includedProperties).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailIdGetById`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeStats** | **bool** | Include statistics with email. | 
 **marketingCampaignNames** | **bool** | If set to true, loads &#x60;campaignName&#x60; and &#x60;campaignUtm&#x60;. | 
 **workflowNames** | **bool** | If set to true, loads workflows in which the email is used within a \&quot;send email\&quot; action.  | 
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


## GetMarketingV3EmailsEmailIdRevisionsGetRevisions

> CollectionResponseWithTotalVersionPublicEmail GetMarketingV3EmailsEmailIdRevisionsGetRevisions(ctx, emailId).After(after).Before(before).Limit(limit).Execute()

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
	limit := int32(56) // int32 | The maximum number of results to return. Default is 10. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisionsGetRevisions(context.Background(), emailId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisionsGetRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailIdRevisionsGetRevisions`: CollectionResponseWithTotalVersionPublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisionsGetRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdRevisionsGetRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** | The cursor token value to get the previous set of results. You can get this from the &#x60;paging.prev.before&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 10. | 

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


## GetMarketingV3EmailsEmailIdRevisionsRevisionIdGetRevisionById

> VersionPublicEmail GetMarketingV3EmailsEmailIdRevisionsRevisionIdGetRevisionById(ctx, emailId, revisionId).Execute()

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
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisionsRevisionIdGetRevisionById(context.Background(), emailId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisionsRevisionIdGetRevisionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsEmailIdRevisionsRevisionIdGetRevisionById`: VersionPublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsEmailIdRevisionsRevisionIdGetRevisionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 
**revisionId** | **string** | The ID of a revision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsEmailIdRevisionsRevisionIdGetRevisionByIdRequest struct via the builder pattern


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


## GetMarketingV3EmailsGetPage

> CollectionResponseWithTotalPublicEmailForwardPaging GetMarketingV3EmailsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).IncludeStats(includeStats).MarketingCampaignNames(marketingCampaignNames).WorkflowNames(workflowNames).Type_(type_).IsPublished(isPublished).IncludedProperties(includedProperties).Campaign(campaign).Archived(archived).Execute()

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
	limit := int32(56) // int32 | The maximum number of results to return. Default is 10. (optional)
	includeStats := true // bool | Include statistics with emails. (optional)
	marketingCampaignNames := true // bool |  (optional)
	workflowNames := true // bool |  (optional)
	type_ := "type__example" // string | Email types to be filtered by. Multiple types can be included. All emails will be returned if not present. (optional)
	isPublished := true // bool | Filter by published/draft emails. All emails will be returned if not present. (optional)
	includedProperties := []string{"Inner_example"} // []string |  (optional)
	campaign := "campaign_example" // string | Filter by campaign GUID. All emails will be returned if not present. (optional)
	archived := true // bool | Specifies whether to return archived emails. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.GetMarketingV3EmailsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).IncludeStats(includeStats).MarketingCampaignNames(marketingCampaignNames).WorkflowNames(workflowNames).Type_(type_).IsPublished(isPublished).IncludedProperties(includedProperties).Campaign(campaign).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.GetMarketingV3EmailsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsGetPage`: CollectionResponseWithTotalPublicEmailForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.GetMarketingV3EmailsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsGetPageRequest struct via the builder pattern


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
 **limit** | **int32** | The maximum number of results to return. Default is 10. | 
 **includeStats** | **bool** | Include statistics with emails. | 
 **marketingCampaignNames** | **bool** |  | 
 **workflowNames** | **bool** |  | 
 **type_** | **string** | Email types to be filtered by. Multiple types can be included. All emails will be returned if not present. | 
 **isPublished** | **bool** | Filter by published/draft emails. All emails will be returned if not present. | 
 **includedProperties** | **[]string** |  | 
 **campaign** | **string** | Filter by campaign GUID. All emails will be returned if not present. | 
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


## PatchMarketingV3EmailsEmailIdDraftUpsertDraft

> PublicEmail PatchMarketingV3EmailsEmailIdDraftUpsertDraft(ctx, emailId).EmailUpdateRequest(emailUpdateRequest).Execute()

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
	resp, r, err := apiClient.MarketingEmailsAPI.PatchMarketingV3EmailsEmailIdDraftUpsertDraft(context.Background(), emailId).EmailUpdateRequest(emailUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PatchMarketingV3EmailsEmailIdDraftUpsertDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingV3EmailsEmailIdDraftUpsertDraft`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PatchMarketingV3EmailsEmailIdDraftUpsertDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3EmailsEmailIdDraftUpsertDraftRequest struct via the builder pattern


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


## PatchMarketingV3EmailsEmailIdUpdate

> PublicEmail PatchMarketingV3EmailsEmailIdUpdate(ctx, emailId).EmailUpdateRequest(emailUpdateRequest).Archived(archived).Execute()

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
	resp, r, err := apiClient.MarketingEmailsAPI.PatchMarketingV3EmailsEmailIdUpdate(context.Background(), emailId).EmailUpdateRequest(emailUpdateRequest).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PatchMarketingV3EmailsEmailIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingV3EmailsEmailIdUpdate`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PatchMarketingV3EmailsEmailIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The ID of the marketing email that should get updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3EmailsEmailIdUpdateRequest struct via the builder pattern


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


## PostMarketingV3EmailsAbTestCreateVariationCreateAbTestVariation

> PublicEmail PostMarketingV3EmailsAbTestCreateVariationCreateAbTestVariation(ctx).AbTestCreateRequestVNext(abTestCreateRequestVNext).Execute()

Create an A/B test variation of a marketing email.



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
	abTestCreateRequestVNext := *openapiclient.NewAbTestCreateRequestVNext("VariationName_example", "ContentId_example") // AbTestCreateRequestVNext | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsAbTestCreateVariationCreateAbTestVariation(context.Background()).AbTestCreateRequestVNext(abTestCreateRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsAbTestCreateVariationCreateAbTestVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3EmailsAbTestCreateVariationCreateAbTestVariation`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PostMarketingV3EmailsAbTestCreateVariationCreateAbTestVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsAbTestCreateVariationCreateAbTestVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **abTestCreateRequestVNext** | [**AbTestCreateRequestVNext**](AbTestCreateRequestVNext.md) |  | 

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


## PostMarketingV3EmailsCloneClone

> PublicEmail PostMarketingV3EmailsCloneClone(ctx).EmailCloneRequestVNext(emailCloneRequestVNext).Execute()

Clone a marketing email.



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
	emailCloneRequestVNext := *openapiclient.NewEmailCloneRequestVNext("Id_example") // EmailCloneRequestVNext | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsCloneClone(context.Background()).EmailCloneRequestVNext(emailCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsCloneClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3EmailsCloneClone`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PostMarketingV3EmailsCloneClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsCloneCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailCloneRequestVNext** | [**EmailCloneRequestVNext**](EmailCloneRequestVNext.md) |  | 

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


## PostMarketingV3EmailsCreate

> PublicEmail PostMarketingV3EmailsCreate(ctx).EmailCreateRequest(emailCreateRequest).Execute()

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
	resp, r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsCreate(context.Background()).EmailCreateRequest(emailCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3EmailsCreate`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PostMarketingV3EmailsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsCreateRequest struct via the builder pattern


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


## PostMarketingV3EmailsEmailIdDraftResetResetDraft

> PostMarketingV3EmailsEmailIdDraftResetResetDraft(ctx, emailId).Execute()

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
	r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsEmailIdDraftResetResetDraft(context.Background(), emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdDraftResetResetDraft``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsEmailIdDraftResetResetDraftRequest struct via the builder pattern


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


## PostMarketingV3EmailsEmailIdPublishPublishOrSend

> PostMarketingV3EmailsEmailIdPublishPublishOrSend(ctx, emailId).Execute()

Publish or send a marketing email.



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
	r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsEmailIdPublishPublishOrSend(context.Background(), emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdPublishPublishOrSend``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsEmailIdPublishPublishOrSendRequest struct via the builder pattern


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


## PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreRestoreRevision

> PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreRestoreRevision(ctx, emailId, revisionId).Execute()

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
	r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreRestoreRevision(context.Background(), emailId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreRestoreRevision``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreRestoreRevisionRequest struct via the builder pattern


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


## PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRestoreDraftRevision

> PublicEmail PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRestoreDraftRevision(ctx, emailId, revisionId).Execute()

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
	resp, r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRestoreDraftRevision(context.Background(), emailId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRestoreDraftRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRestoreDraftRevision`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRestoreDraftRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** | The marketing email ID. | 
**revisionId** | **int64** | The ID of a revision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsEmailIdRevisionsRevisionIdRestoreToDraftRestoreDraftRevisionRequest struct via the builder pattern


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


## PostMarketingV3EmailsEmailIdUnpublishUnpublishOrCancel

> PostMarketingV3EmailsEmailIdUnpublishUnpublishOrCancel(ctx, emailId).Execute()

Unpublish or cancel a marketing email.



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
	emailId := "emailId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingEmailsAPI.PostMarketingV3EmailsEmailIdUnpublishUnpublishOrCancel(context.Background(), emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEmailsAPI.PostMarketingV3EmailsEmailIdUnpublishUnpublishOrCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsEmailIdUnpublishUnpublishOrCancelRequest struct via the builder pattern


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

