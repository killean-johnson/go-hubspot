# \SitePagesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3PagesSitePagesObjectIdArchive**](SitePagesAPI.md#DeleteCmsV3PagesSitePagesObjectIdArchive) | **Delete** /cms/v3/pages/site-pages/{objectId} | Delete a Site Page
[**GetCmsV3PagesSitePagesGetPage**](SitePagesAPI.md#GetCmsV3PagesSitePagesGetPage) | **Get** /cms/v3/pages/site-pages | Get all Site Pages
[**GetCmsV3PagesSitePagesObjectIdDraftGetDraftById**](SitePagesAPI.md#GetCmsV3PagesSitePagesObjectIdDraftGetDraftById) | **Get** /cms/v3/pages/site-pages/{objectId}/draft | Retrieve the full draft version of the Site Page
[**GetCmsV3PagesSitePagesObjectIdGetById**](SitePagesAPI.md#GetCmsV3PagesSitePagesObjectIdGetById) | **Get** /cms/v3/pages/site-pages/{objectId} | Retrieve a Site Page
[**GetCmsV3PagesSitePagesObjectIdRevisionsGetPreviousVersions**](SitePagesAPI.md#GetCmsV3PagesSitePagesObjectIdRevisionsGetPreviousVersions) | **Get** /cms/v3/pages/site-pages/{objectId}/revisions | Retrieves all the previous versions of a Site Page
[**GetCmsV3PagesSitePagesObjectIdRevisionsRevisionIdGetPreviousVersion**](SitePagesAPI.md#GetCmsV3PagesSitePagesObjectIdRevisionsRevisionIdGetPreviousVersion) | **Get** /cms/v3/pages/site-pages/{objectId}/revisions/{revisionId} | Retrieves a previous version of a Site Page
[**PatchCmsV3PagesSitePagesObjectIdDraftUpdateDraft**](SitePagesAPI.md#PatchCmsV3PagesSitePagesObjectIdDraftUpdateDraft) | **Patch** /cms/v3/pages/site-pages/{objectId}/draft | Update a Site Page draft
[**PatchCmsV3PagesSitePagesObjectIdUpdate**](SitePagesAPI.md#PatchCmsV3PagesSitePagesObjectIdUpdate) | **Patch** /cms/v3/pages/site-pages/{objectId} | Update a Site Page
[**PostCmsV3PagesSitePagesAbTestCreateVariationCreateABTestVariation**](SitePagesAPI.md#PostCmsV3PagesSitePagesAbTestCreateVariationCreateABTestVariation) | **Post** /cms/v3/pages/site-pages/ab-test/create-variation | Create a new A/B test variation
[**PostCmsV3PagesSitePagesAbTestEndEndActiveABTest**](SitePagesAPI.md#PostCmsV3PagesSitePagesAbTestEndEndActiveABTest) | **Post** /cms/v3/pages/site-pages/ab-test/end | End an active A/B test
[**PostCmsV3PagesSitePagesAbTestRerunRerunPreviousABTest**](SitePagesAPI.md#PostCmsV3PagesSitePagesAbTestRerunRerunPreviousABTest) | **Post** /cms/v3/pages/site-pages/ab-test/rerun | Rerun a previous A/B test
[**PostCmsV3PagesSitePagesBatchArchiveArchiveBatch**](SitePagesAPI.md#PostCmsV3PagesSitePagesBatchArchiveArchiveBatch) | **Post** /cms/v3/pages/site-pages/batch/archive | Delete a batch of Site Pages
[**PostCmsV3PagesSitePagesBatchCreateCreateBatch**](SitePagesAPI.md#PostCmsV3PagesSitePagesBatchCreateCreateBatch) | **Post** /cms/v3/pages/site-pages/batch/create | Create a batch of Site Pages
[**PostCmsV3PagesSitePagesBatchReadReadBatch**](SitePagesAPI.md#PostCmsV3PagesSitePagesBatchReadReadBatch) | **Post** /cms/v3/pages/site-pages/batch/read | Retrieve a batch of Site Pages
[**PostCmsV3PagesSitePagesBatchUpdateUpdateBatch**](SitePagesAPI.md#PostCmsV3PagesSitePagesBatchUpdateUpdateBatch) | **Post** /cms/v3/pages/site-pages/batch/update | Update a batch of Site Pages
[**PostCmsV3PagesSitePagesCloneClone**](SitePagesAPI.md#PostCmsV3PagesSitePagesCloneClone) | **Post** /cms/v3/pages/site-pages/clone | Clone a Site Page
[**PostCmsV3PagesSitePagesCreate**](SitePagesAPI.md#PostCmsV3PagesSitePagesCreate) | **Post** /cms/v3/pages/site-pages | Create a new Site Page
[**PostCmsV3PagesSitePagesMultiLanguageAttachToLangGroupAttachToLangGroup**](SitePagesAPI.md#PostCmsV3PagesSitePagesMultiLanguageAttachToLangGroupAttachToLangGroup) | **Post** /cms/v3/pages/site-pages/multi-language/attach-to-lang-group | Attach a site page to a multi-language group
[**PostCmsV3PagesSitePagesMultiLanguageCreateLanguageVariationCreateLangVariation**](SitePagesAPI.md#PostCmsV3PagesSitePagesMultiLanguageCreateLanguageVariationCreateLangVariation) | **Post** /cms/v3/pages/site-pages/multi-language/create-language-variation | Create a new language variation
[**PostCmsV3PagesSitePagesMultiLanguageDetachFromLangGroupDetachFromLangGroup**](SitePagesAPI.md#PostCmsV3PagesSitePagesMultiLanguageDetachFromLangGroupDetachFromLangGroup) | **Post** /cms/v3/pages/site-pages/multi-language/detach-from-lang-group | Detach a site page from a multi-language group
[**PostCmsV3PagesSitePagesMultiLanguageUpdateLanguagesUpdateLangs**](SitePagesAPI.md#PostCmsV3PagesSitePagesMultiLanguageUpdateLanguagesUpdateLangs) | **Post** /cms/v3/pages/site-pages/multi-language/update-languages | Update languages of multi-language group
[**PostCmsV3PagesSitePagesObjectIdDraftPushLivePushLive**](SitePagesAPI.md#PostCmsV3PagesSitePagesObjectIdDraftPushLivePushLive) | **Post** /cms/v3/pages/site-pages/{objectId}/draft/push-live | Push Site Page draft edits live
[**PostCmsV3PagesSitePagesObjectIdDraftResetResetDraft**](SitePagesAPI.md#PostCmsV3PagesSitePagesObjectIdDraftResetResetDraft) | **Post** /cms/v3/pages/site-pages/{objectId}/draft/reset | Reset the Site Page draft to the live version
[**PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion**](SitePagesAPI.md#PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion) | **Post** /cms/v3/pages/site-pages/{objectId}/revisions/{revisionId}/restore | Restore a previous version of a Site Page
[**PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft**](SitePagesAPI.md#PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft) | **Post** /cms/v3/pages/site-pages/{objectId}/revisions/{revisionId}/restore-to-draft | Restore a previous version of a Site Page, to the draft version of the Site Page
[**PostCmsV3PagesSitePagesScheduleSchedule**](SitePagesAPI.md#PostCmsV3PagesSitePagesScheduleSchedule) | **Post** /cms/v3/pages/site-pages/schedule | Schedule a Site Page to be Published
[**PutCmsV3PagesSitePagesMultiLanguageSetNewLangPrimarySetLangPrimary**](SitePagesAPI.md#PutCmsV3PagesSitePagesMultiLanguageSetNewLangPrimarySetLangPrimary) | **Put** /cms/v3/pages/site-pages/multi-language/set-new-lang-primary | Set a new primary language



## DeleteCmsV3PagesSitePagesObjectIdArchive

> DeleteCmsV3PagesSitePagesObjectIdArchive(ctx, objectId).Archived(archived).Execute()

Delete a Site Page



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
	objectId := "objectId_example" // string | The Site Page id.
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.DeleteCmsV3PagesSitePagesObjectIdArchive(context.Background(), objectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.DeleteCmsV3PagesSitePagesObjectIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Site Page id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3PagesSitePagesObjectIdArchiveRequest struct via the builder pattern


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


## GetCmsV3PagesSitePagesGetPage

> CollectionResponseWithTotalPageForwardPaging GetCmsV3PagesSitePagesGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()

Get all Site Pages



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
	createdAt := time.Now() // time.Time | Only return Site Pages created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return Site Pages created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return Site Pages created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return Site Pages last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return Site Pages last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return Site Pages last updated before the specified time. (optional)
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)
	archived := true // bool | Specifies whether to return deleted Site Pages. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.GetCmsV3PagesSitePagesGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.GetCmsV3PagesSitePagesGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3PagesSitePagesGetPage`: CollectionResponseWithTotalPageForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.GetCmsV3PagesSitePagesGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3PagesSitePagesGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Site Pages created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Site Pages created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Site Pages created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Site Pages last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Site Pages last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Site Pages last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 
 **archived** | **bool** | Specifies whether to return deleted Site Pages. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**CollectionResponseWithTotalPageForwardPaging**](CollectionResponseWithTotalPageForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3PagesSitePagesObjectIdDraftGetDraftById

> Page GetCmsV3PagesSitePagesObjectIdDraftGetDraftById(ctx, objectId).Execute()

Retrieve the full draft version of the Site Page



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
	objectId := "objectId_example" // string | The Site Page id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.GetCmsV3PagesSitePagesObjectIdDraftGetDraftById(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.GetCmsV3PagesSitePagesObjectIdDraftGetDraftById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3PagesSitePagesObjectIdDraftGetDraftById`: Page
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.GetCmsV3PagesSitePagesObjectIdDraftGetDraftById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Site Page id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3PagesSitePagesObjectIdDraftGetDraftByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Page**](Page.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3PagesSitePagesObjectIdGetById

> Page GetCmsV3PagesSitePagesObjectIdGetById(ctx, objectId).Archived(archived).Property(property).Execute()

Retrieve a Site Page



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
	objectId := "objectId_example" // string | The Site Page id.
	archived := true // bool | Specifies whether to return deleted Site Pages. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.GetCmsV3PagesSitePagesObjectIdGetById(context.Background(), objectId).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.GetCmsV3PagesSitePagesObjectIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3PagesSitePagesObjectIdGetById`: Page
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.GetCmsV3PagesSitePagesObjectIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Site Page id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3PagesSitePagesObjectIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Site Pages. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**Page**](Page.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3PagesSitePagesObjectIdRevisionsGetPreviousVersions

> CollectionResponseWithTotalVersionPage GetCmsV3PagesSitePagesObjectIdRevisionsGetPreviousVersions(ctx, objectId).After(after).Before(before).Limit(limit).Execute()

Retrieves all the previous versions of a Site Page



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
	objectId := "objectId_example" // string | The Site Page id.
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.GetCmsV3PagesSitePagesObjectIdRevisionsGetPreviousVersions(context.Background(), objectId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.GetCmsV3PagesSitePagesObjectIdRevisionsGetPreviousVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3PagesSitePagesObjectIdRevisionsGetPreviousVersions`: CollectionResponseWithTotalVersionPage
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.GetCmsV3PagesSitePagesObjectIdRevisionsGetPreviousVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Site Page id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3PagesSitePagesObjectIdRevisionsGetPreviousVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 

### Return type

[**CollectionResponseWithTotalVersionPage**](CollectionResponseWithTotalVersionPage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3PagesSitePagesObjectIdRevisionsRevisionIdGetPreviousVersion

> VersionPage GetCmsV3PagesSitePagesObjectIdRevisionsRevisionIdGetPreviousVersion(ctx, objectId, revisionId).Execute()

Retrieves a previous version of a Site Page



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
	objectId := "objectId_example" // string | The Site Page id.
	revisionId := "revisionId_example" // string | The Site Page version id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.GetCmsV3PagesSitePagesObjectIdRevisionsRevisionIdGetPreviousVersion(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.GetCmsV3PagesSitePagesObjectIdRevisionsRevisionIdGetPreviousVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3PagesSitePagesObjectIdRevisionsRevisionIdGetPreviousVersion`: VersionPage
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.GetCmsV3PagesSitePagesObjectIdRevisionsRevisionIdGetPreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Site Page id. | 
**revisionId** | **string** | The Site Page version id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3PagesSitePagesObjectIdRevisionsRevisionIdGetPreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionPage**](VersionPage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3PagesSitePagesObjectIdDraftUpdateDraft

> Page PatchCmsV3PagesSitePagesObjectIdDraftUpdateDraft(ctx, objectId).Page(page).Execute()

Update a Site Page draft



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
	objectId := "objectId_example" // string | The Site Page id.
	page := *openapiclient.NewPage(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", false, "HtmlTitle_example", map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, "Subcategory_example", map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "ContentGroupId_example", false, "TemplatePath_example", "Url_example", []map[string]interface{}{map[string]interface{}(123)}, time.Now(), map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // Page | The JSON representation of the updated Site Page to be applied to the draft.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PatchCmsV3PagesSitePagesObjectIdDraftUpdateDraft(context.Background(), objectId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PatchCmsV3PagesSitePagesObjectIdDraftUpdateDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3PagesSitePagesObjectIdDraftUpdateDraft`: Page
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PatchCmsV3PagesSitePagesObjectIdDraftUpdateDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Site Page id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3PagesSitePagesObjectIdDraftUpdateDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | [**Page**](Page.md) | The JSON representation of the updated Site Page to be applied to the draft. | 

### Return type

[**Page**](Page.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3PagesSitePagesObjectIdUpdate

> Page PatchCmsV3PagesSitePagesObjectIdUpdate(ctx, objectId).Page(page).Archived(archived).Execute()

Update a Site Page



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
	objectId := "objectId_example" // string | The Site Page id.
	page := *openapiclient.NewPage(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", false, "HtmlTitle_example", map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, "Subcategory_example", map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "ContentGroupId_example", false, "TemplatePath_example", "Url_example", []map[string]interface{}{map[string]interface{}(123)}, time.Now(), map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // Page | The JSON representation of the updated Site Page.
	archived := true // bool | Specifies whether to update deleted Site Pages. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PatchCmsV3PagesSitePagesObjectIdUpdate(context.Background(), objectId).Page(page).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PatchCmsV3PagesSitePagesObjectIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3PagesSitePagesObjectIdUpdate`: Page
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PatchCmsV3PagesSitePagesObjectIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Site Page id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3PagesSitePagesObjectIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | [**Page**](Page.md) | The JSON representation of the updated Site Page. | 
 **archived** | **bool** | Specifies whether to update deleted Site Pages. Defaults to &#x60;false&#x60;. | 

### Return type

[**Page**](Page.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesAbTestCreateVariationCreateABTestVariation

> Page PostCmsV3PagesSitePagesAbTestCreateVariationCreateABTestVariation(ctx).AbTestCreateRequestVNext(abTestCreateRequestVNext).Execute()

Create a new A/B test variation



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
	abTestCreateRequestVNext := *openapiclient.NewAbTestCreateRequestVNext("VariationName_example", "ContentId_example") // AbTestCreateRequestVNext | The JSON representation of the AbTestCreateRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesAbTestCreateVariationCreateABTestVariation(context.Background()).AbTestCreateRequestVNext(abTestCreateRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesAbTestCreateVariationCreateABTestVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3PagesSitePagesAbTestCreateVariationCreateABTestVariation`: Page
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PostCmsV3PagesSitePagesAbTestCreateVariationCreateABTestVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesAbTestCreateVariationCreateABTestVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **abTestCreateRequestVNext** | [**AbTestCreateRequestVNext**](AbTestCreateRequestVNext.md) | The JSON representation of the AbTestCreateRequest object. | 

### Return type

[**Page**](Page.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesAbTestEndEndActiveABTest

> PostCmsV3PagesSitePagesAbTestEndEndActiveABTest(ctx).AbTestEndRequestVNext(abTestEndRequestVNext).Execute()

End an active A/B test



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
	abTestEndRequestVNext := *openapiclient.NewAbTestEndRequestVNext("WinnerId_example", "AbTestId_example") // AbTestEndRequestVNext | The JSON representation of the AbTestEndRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesAbTestEndEndActiveABTest(context.Background()).AbTestEndRequestVNext(abTestEndRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesAbTestEndEndActiveABTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesAbTestEndEndActiveABTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **abTestEndRequestVNext** | [**AbTestEndRequestVNext**](AbTestEndRequestVNext.md) | The JSON representation of the AbTestEndRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesAbTestRerunRerunPreviousABTest

> PostCmsV3PagesSitePagesAbTestRerunRerunPreviousABTest(ctx).AbTestRerunRequestVNext(abTestRerunRequestVNext).Execute()

Rerun a previous A/B test



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
	abTestRerunRequestVNext := *openapiclient.NewAbTestRerunRequestVNext("VariationId_example", "AbTestId_example") // AbTestRerunRequestVNext | The JSON representation of the AbTestRerunRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesAbTestRerunRerunPreviousABTest(context.Background()).AbTestRerunRequestVNext(abTestRerunRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesAbTestRerunRerunPreviousABTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesAbTestRerunRerunPreviousABTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **abTestRerunRequestVNext** | [**AbTestRerunRequestVNext**](AbTestRerunRequestVNext.md) | The JSON representation of the AbTestRerunRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesBatchArchiveArchiveBatch

> PostCmsV3PagesSitePagesBatchArchiveArchiveBatch(ctx).BatchInputString(batchInputString).Execute()

Delete a batch of Site Pages



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Site Page ids.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesBatchArchiveArchiveBatch(context.Background()).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesBatchArchiveArchiveBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesBatchArchiveArchiveBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Site Page ids. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesBatchCreateCreateBatch

> BatchResponsePage PostCmsV3PagesSitePagesBatchCreateCreateBatch(ctx).BatchInputPage(batchInputPage).Execute()

Create a batch of Site Pages



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
	batchInputPage := *openapiclient.NewBatchInputPage([]openapiclient.Page{*openapiclient.NewPage(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", false, "HtmlTitle_example", map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, "Subcategory_example", map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "ContentGroupId_example", false, "TemplatePath_example", "Url_example", []map[string]interface{}{map[string]interface{}(123)}, time.Now(), map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example")}) // BatchInputPage | The JSON array of new Site Pages to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesBatchCreateCreateBatch(context.Background()).BatchInputPage(batchInputPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesBatchCreateCreateBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3PagesSitePagesBatchCreateCreateBatch`: BatchResponsePage
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PostCmsV3PagesSitePagesBatchCreateCreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesBatchCreateCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputPage** | [**BatchInputPage**](BatchInputPage.md) | The JSON array of new Site Pages to create. | 

### Return type

[**BatchResponsePage**](BatchResponsePage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesBatchReadReadBatch

> BatchResponsePage PostCmsV3PagesSitePagesBatchReadReadBatch(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Site Pages



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Site Page ids.
	archived := true // bool | Specifies whether to return deleted Site Pages. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesBatchReadReadBatch(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesBatchReadReadBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3PagesSitePagesBatchReadReadBatch`: BatchResponsePage
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PostCmsV3PagesSitePagesBatchReadReadBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesBatchReadReadBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Site Page ids. | 
 **archived** | **bool** | Specifies whether to return deleted Site Pages. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponsePage**](BatchResponsePage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesBatchUpdateUpdateBatch

> BatchResponsePage PostCmsV3PagesSitePagesBatchUpdateUpdateBatch(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Site Pages



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
	batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | The JSON representation of the updated Site Pages.
	archived := true // bool | Specifies whether to update deleted Site Pages. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesBatchUpdateUpdateBatch(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesBatchUpdateUpdateBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3PagesSitePagesBatchUpdateUpdateBatch`: BatchResponsePage
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PostCmsV3PagesSitePagesBatchUpdateUpdateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesBatchUpdateUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | The JSON representation of the updated Site Pages. | 
 **archived** | **bool** | Specifies whether to update deleted Site Pages. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponsePage**](BatchResponsePage.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesCloneClone

> Page PostCmsV3PagesSitePagesCloneClone(ctx).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()

Clone a Site Page



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
	contentCloneRequestVNext := *openapiclient.NewContentCloneRequestVNext("Id_example") // ContentCloneRequestVNext | The JSON representation of the ContentCloneRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesCloneClone(context.Background()).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesCloneClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3PagesSitePagesCloneClone`: Page
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PostCmsV3PagesSitePagesCloneClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesCloneCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentCloneRequestVNext** | [**ContentCloneRequestVNext**](ContentCloneRequestVNext.md) | The JSON representation of the ContentCloneRequest object. | 

### Return type

[**Page**](Page.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesCreate

> PostCmsV3PagesSitePagesCreate(ctx).Page(page).Execute()

Create a new Site Page



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
	page := *openapiclient.NewPage(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", false, "HtmlTitle_example", map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, "Subcategory_example", map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "ContentGroupId_example", false, "TemplatePath_example", "Url_example", []map[string]interface{}{map[string]interface{}(123)}, time.Now(), map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // Page | The JSON representation of a new Site Page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesCreate(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**Page**](Page.md) | The JSON representation of a new Site Page. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesMultiLanguageAttachToLangGroupAttachToLangGroup

> PostCmsV3PagesSitePagesMultiLanguageAttachToLangGroupAttachToLangGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

Attach a site page to a multi-language group



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
	attachToLangPrimaryRequestVNext := *openapiclient.NewAttachToLangPrimaryRequestVNext("Language_example", "Id_example", "PrimaryId_example") // AttachToLangPrimaryRequestVNext | The JSON representation of the AttachToLangPrimaryRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesMultiLanguageAttachToLangGroupAttachToLangGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesMultiLanguageAttachToLangGroupAttachToLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesMultiLanguageAttachToLangGroupAttachToLangGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachToLangPrimaryRequestVNext** | [**AttachToLangPrimaryRequestVNext**](AttachToLangPrimaryRequestVNext.md) | The JSON representation of the AttachToLangPrimaryRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesMultiLanguageCreateLanguageVariationCreateLangVariation

> Page PostCmsV3PagesSitePagesMultiLanguageCreateLanguageVariationCreateLangVariation(ctx).ContentLanguageCloneRequestVNext(contentLanguageCloneRequestVNext).Execute()

Create a new language variation



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
	contentLanguageCloneRequestVNext := *openapiclient.NewContentLanguageCloneRequestVNext("Id_example") // ContentLanguageCloneRequestVNext | The JSON representation of the ContentLanguageCloneRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesMultiLanguageCreateLanguageVariationCreateLangVariation(context.Background()).ContentLanguageCloneRequestVNext(contentLanguageCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesMultiLanguageCreateLanguageVariationCreateLangVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3PagesSitePagesMultiLanguageCreateLanguageVariationCreateLangVariation`: Page
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PostCmsV3PagesSitePagesMultiLanguageCreateLanguageVariationCreateLangVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesMultiLanguageCreateLanguageVariationCreateLangVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguageCloneRequestVNext** | [**ContentLanguageCloneRequestVNext**](ContentLanguageCloneRequestVNext.md) | The JSON representation of the ContentLanguageCloneRequest object. | 

### Return type

[**Page**](Page.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesMultiLanguageDetachFromLangGroupDetachFromLangGroup

> PostCmsV3PagesSitePagesMultiLanguageDetachFromLangGroupDetachFromLangGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

Detach a site page from a multi-language group



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
	detachFromLangGroupRequestVNext := *openapiclient.NewDetachFromLangGroupRequestVNext("Id_example") // DetachFromLangGroupRequestVNext | The JSON representation of the DetachFromLangGroupRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesMultiLanguageDetachFromLangGroupDetachFromLangGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesMultiLanguageDetachFromLangGroupDetachFromLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesMultiLanguageDetachFromLangGroupDetachFromLangGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detachFromLangGroupRequestVNext** | [**DetachFromLangGroupRequestVNext**](DetachFromLangGroupRequestVNext.md) | The JSON representation of the DetachFromLangGroupRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesMultiLanguageUpdateLanguagesUpdateLangs

> PostCmsV3PagesSitePagesMultiLanguageUpdateLanguagesUpdateLangs(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

Update languages of multi-language group



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
	updateLanguagesRequestVNext := *openapiclient.NewUpdateLanguagesRequestVNext(map[string]string{"key": "Inner_example"}, "PrimaryId_example") // UpdateLanguagesRequestVNext | The JSON representation of the UpdateLanguagesRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesMultiLanguageUpdateLanguagesUpdateLangs(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesMultiLanguageUpdateLanguagesUpdateLangs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesMultiLanguageUpdateLanguagesUpdateLangsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLanguagesRequestVNext** | [**UpdateLanguagesRequestVNext**](UpdateLanguagesRequestVNext.md) | The JSON representation of the UpdateLanguagesRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesObjectIdDraftPushLivePushLive

> PostCmsV3PagesSitePagesObjectIdDraftPushLivePushLive(ctx, objectId).Execute()

Push Site Page draft edits live



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
	objectId := "objectId_example" // string | The id of the Site Page for which it's draft will be pushed live.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesObjectIdDraftPushLivePushLive(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesObjectIdDraftPushLivePushLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The id of the Site Page for which it&#39;s draft will be pushed live. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesObjectIdDraftPushLivePushLiveRequest struct via the builder pattern


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


## PostCmsV3PagesSitePagesObjectIdDraftResetResetDraft

> PostCmsV3PagesSitePagesObjectIdDraftResetResetDraft(ctx, objectId).Execute()

Reset the Site Page draft to the live version



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
	objectId := "objectId_example" // string | The id of the Site Page for which it's draft will be reset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesObjectIdDraftResetResetDraft(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesObjectIdDraftResetResetDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The id of the Site Page for which it&#39;s draft will be reset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesObjectIdDraftResetResetDraftRequest struct via the builder pattern


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


## PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion

> Page PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion(ctx, objectId, revisionId).Execute()

Restore a previous version of a Site Page



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
	objectId := "objectId_example" // string | The Site Page id.
	revisionId := "revisionId_example" // string | The Site Page version id to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion`: Page
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Site Page id. | 
**revisionId** | **string** | The Site Page version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreRestorePreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Page**](Page.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft

> Page PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft(ctx, objectId, revisionId).Execute()

Restore a previous version of a Site Page, to the draft version of the Site Page



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
	objectId := "objectId_example" // string | The Site Page id.
	revisionId := int64(789) // int64 | The Site Page version id to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft`: Page
	fmt.Fprintf(os.Stdout, "Response from `SitePagesAPI.PostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Site Page id. | 
**revisionId** | **int64** | The Site Page version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Page**](Page.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3PagesSitePagesScheduleSchedule

> PostCmsV3PagesSitePagesScheduleSchedule(ctx).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()

Schedule a Site Page to be Published



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
	contentScheduleRequestVNext := *openapiclient.NewContentScheduleRequestVNext(time.Now(), "Id_example") // ContentScheduleRequestVNext | The JSON representation of the ContentScheduleRequestVNext object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PostCmsV3PagesSitePagesScheduleSchedule(context.Background()).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PostCmsV3PagesSitePagesScheduleSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3PagesSitePagesScheduleScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentScheduleRequestVNext** | [**ContentScheduleRequestVNext**](ContentScheduleRequestVNext.md) | The JSON representation of the ContentScheduleRequestVNext object. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCmsV3PagesSitePagesMultiLanguageSetNewLangPrimarySetLangPrimary

> PutCmsV3PagesSitePagesMultiLanguageSetNewLangPrimarySetLangPrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

Set a new primary language



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
	setNewLanguagePrimaryRequestVNext := *openapiclient.NewSetNewLanguagePrimaryRequestVNext("Id_example") // SetNewLanguagePrimaryRequestVNext | The JSON representation of the SetNewLanguagePrimaryRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitePagesAPI.PutCmsV3PagesSitePagesMultiLanguageSetNewLangPrimarySetLangPrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitePagesAPI.PutCmsV3PagesSitePagesMultiLanguageSetNewLangPrimarySetLangPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3PagesSitePagesMultiLanguageSetNewLangPrimarySetLangPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setNewLanguagePrimaryRequestVNext** | [**SetNewLanguagePrimaryRequestVNext**](SetNewLanguagePrimaryRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

