# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3BlogsPostsObjectIdArchive**](BasicAPI.md#DeleteCmsV3BlogsPostsObjectIdArchive) | **Delete** /cms/v3/blogs/posts/{objectId} | Delete a blog post
[**GetCmsV3BlogsPostsGetPage**](BasicAPI.md#GetCmsV3BlogsPostsGetPage) | **Get** /cms/v3/blogs/posts | Get all posts
[**GetCmsV3BlogsPostsObjectIdDraftGetDraftById**](BasicAPI.md#GetCmsV3BlogsPostsObjectIdDraftGetDraftById) | **Get** /cms/v3/blogs/posts/{objectId}/draft | Retrieve the full draft version of the Blog Post
[**GetCmsV3BlogsPostsObjectIdGetById**](BasicAPI.md#GetCmsV3BlogsPostsObjectIdGetById) | **Get** /cms/v3/blogs/posts/{objectId} | Retrieve a blog post
[**GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions**](BasicAPI.md#GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions) | **Get** /cms/v3/blogs/posts/{objectId}/revisions | Retrieves all previous versions of a post
[**GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion**](BasicAPI.md#GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion) | **Get** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId} | Retrieve a previous version of a blog post
[**PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft**](BasicAPI.md#PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft) | **Patch** /cms/v3/blogs/posts/{objectId}/draft | Update the draft of a post
[**PatchCmsV3BlogsPostsObjectIdUpdate**](BasicAPI.md#PatchCmsV3BlogsPostsObjectIdUpdate) | **Patch** /cms/v3/blogs/posts/{objectId} | Update a post
[**PostCmsV3BlogsPostsCloneClone**](BasicAPI.md#PostCmsV3BlogsPostsCloneClone) | **Post** /cms/v3/blogs/posts/clone | Clone a blog post
[**PostCmsV3BlogsPostsCreate**](BasicAPI.md#PostCmsV3BlogsPostsCreate) | **Post** /cms/v3/blogs/posts | Create a new post
[**PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive**](BasicAPI.md#PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive) | **Post** /cms/v3/blogs/posts/{objectId}/draft/push-live | Publish blog post draft
[**PostCmsV3BlogsPostsObjectIdDraftResetResetDraft**](BasicAPI.md#PostCmsV3BlogsPostsObjectIdDraftResetResetDraft) | **Post** /cms/v3/blogs/posts/{objectId}/draft/reset | Reset post draft to the live version
[**PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion**](BasicAPI.md#PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore | Restore a previous version
[**PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft**](BasicAPI.md#PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore-to-draft | Restore a draft to a previous version
[**PostCmsV3BlogsPostsScheduleSchedule**](BasicAPI.md#PostCmsV3BlogsPostsScheduleSchedule) | **Post** /cms/v3/blogs/posts/schedule | Schedule a post to be published



## DeleteCmsV3BlogsPostsObjectIdArchive

> DeleteCmsV3BlogsPostsObjectIdArchive(ctx, objectId).Archived(archived).Execute()

Delete a blog post



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
	objectId := "objectId_example" // string | The ID of the blog post to delete.
	archived := true // bool | Whether to return only results that have been deleted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteCmsV3BlogsPostsObjectIdArchive(context.Background(), objectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteCmsV3BlogsPostsObjectIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3BlogsPostsObjectIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been deleted. | 

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


## GetCmsV3BlogsPostsGetPage

> CollectionResponseWithTotalBlogPostForwardPaging GetCmsV3BlogsPostsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()

Get all posts



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
	createdAt := time.Now() // time.Time | Only return blog posts created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return blog posts created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return blog posts created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return blog posts last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return blog posts last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return blog posts last updated before the specified time. (optional)
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `createdAt` (default), `name`, `updatedAt`, `createdBy`, `updatedBy`. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 20. (optional)
	archived := true // bool | Specifies whether to return deleted blog posts. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCmsV3BlogsPostsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCmsV3BlogsPostsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPostsGetPage`: CollectionResponseWithTotalBlogPostForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCmsV3BlogsPostsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return blog posts created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return blog posts created after the specified time. | 
 **createdBefore** | **time.Time** | Only return blog posts created before the specified time. | 
 **updatedAt** | **time.Time** | Only return blog posts last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return blog posts last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return blog posts last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;createdAt&#x60; (default), &#x60;name&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 20. | 
 **archived** | **bool** | Specifies whether to return deleted blog posts. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**CollectionResponseWithTotalBlogPostForwardPaging**](CollectionResponseWithTotalBlogPostForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdDraftGetDraftById

> BlogPost GetCmsV3BlogsPostsObjectIdDraftGetDraftById(ctx, objectId).Execute()

Retrieve the full draft version of the Blog Post



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
	objectId := "objectId_example" // string | The ID of the blog post to retrieve the draft of.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCmsV3BlogsPostsObjectIdDraftGetDraftById(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCmsV3BlogsPostsObjectIdDraftGetDraftById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPostsObjectIdDraftGetDraftById`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCmsV3BlogsPostsObjectIdDraftGetDraftById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post to retrieve the draft of. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdDraftGetDraftByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdGetById

> BlogPost GetCmsV3BlogsPostsObjectIdGetById(ctx, objectId).Archived(archived).Property(property).Execute()

Retrieve a blog post



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
	objectId := "objectId_example" // string | The ID of the blog post to retrieve.
	archived := true // bool | Specifies whether to return deleted blog posts. Defaults to `false`. (optional)
	property := "property_example" // string | Specific properties to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCmsV3BlogsPostsObjectIdGetById(context.Background(), objectId).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCmsV3BlogsPostsObjectIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPostsObjectIdGetById`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCmsV3BlogsPostsObjectIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted blog posts. Defaults to &#x60;false&#x60;. | 
 **property** | **string** | Specific properties to return. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions

> CollectionResponseWithTotalVersionBlogPost GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions(ctx, objectId).After(after).Before(before).Limit(limit).Execute()

Retrieves all previous versions of a post



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
	objectId := "objectId_example" // string | The ID of the blog post to retrieve previous versions of.
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions(context.Background(), objectId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions`: CollectionResponseWithTotalVersionBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post to retrieve previous versions of. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 

### Return type

[**CollectionResponseWithTotalVersionBlogPost**](CollectionResponseWithTotalVersionBlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion

> VersionBlogPost GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion(ctx, objectId, revisionId).Execute()

Retrieve a previous version of a blog post



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
	objectId := "objectId_example" // string | The ID of the blog post.
	revisionId := "revisionId_example" // string | The ID of the version to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion`: VersionBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post. | 
**revisionId** | **string** | The ID of the version to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionBlogPost**](VersionBlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft

> BlogPost PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft(ctx, objectId).BlogPost(blogPost).Execute()

Update the draft of a post



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
	objectId := "objectId_example" // string | The ID of the blog post to update the draft of.
	blogPost := *openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", "CampaignName_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", false, []map[string]interface{}{map[string]interface{}(123)}, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // BlogPost | The JSON representation of the updated Blog Post to be applied to the draft.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft(context.Background(), objectId).BlogPost(blogPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post to update the draft of. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsPostsObjectIdDraftUpdateDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post to be applied to the draft. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsPostsObjectIdUpdate

> BlogPost PatchCmsV3BlogsPostsObjectIdUpdate(ctx, objectId).BlogPost(blogPost).Archived(archived).Execute()

Update a post



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
	objectId := "objectId_example" // string | The ID of the blog post to update.
	blogPost := *openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", "CampaignName_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", false, []map[string]interface{}{map[string]interface{}(123)}, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // BlogPost | The JSON representation of the updated Blog Post.
	archived := true // bool | Specifies whether to update deleted blog posts. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchCmsV3BlogsPostsObjectIdUpdate(context.Background(), objectId).BlogPost(blogPost).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchCmsV3BlogsPostsObjectIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3BlogsPostsObjectIdUpdate`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchCmsV3BlogsPostsObjectIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsPostsObjectIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post. | 
 **archived** | **bool** | Specifies whether to update deleted blog posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsCloneClone

> BlogPost PostCmsV3BlogsPostsCloneClone(ctx).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()

Clone a blog post



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
	resp, r, err := apiClient.BasicAPI.PostCmsV3BlogsPostsCloneClone(context.Background()).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCmsV3BlogsPostsCloneClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsCloneClone`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostCmsV3BlogsPostsCloneClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsCloneCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentCloneRequestVNext** | [**ContentCloneRequestVNext**](ContentCloneRequestVNext.md) | The JSON representation of the ContentCloneRequest object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsCreate

> BlogPost PostCmsV3BlogsPostsCreate(ctx).BlogPost(blogPost).Execute()

Create a new post



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
	blogPost := *openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", "CampaignName_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", false, []map[string]interface{}{map[string]interface{}(123)}, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // BlogPost | The JSON representation of a new Blog Post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostCmsV3BlogsPostsCreate(context.Background()).BlogPost(blogPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCmsV3BlogsPostsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsCreate`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostCmsV3BlogsPostsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of a new Blog Post. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive

> PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive(ctx, objectId).Execute()

Publish blog post draft



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
	objectId := "objectId_example" // string | The ID of the post to publish.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the post to publish. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdDraftPushLivePushLiveRequest struct via the builder pattern


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


## PostCmsV3BlogsPostsObjectIdDraftResetResetDraft

> PostCmsV3BlogsPostsObjectIdDraftResetResetDraft(ctx, objectId).Execute()

Reset post draft to the live version



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
	objectId := "objectId_example" // string | The ID of the blog post to reset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.PostCmsV3BlogsPostsObjectIdDraftResetResetDraft(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCmsV3BlogsPostsObjectIdDraftResetResetDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post to reset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdDraftResetResetDraftRequest struct via the builder pattern


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


## PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion

> BlogPost PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion(ctx, objectId, revisionId).Execute()

Restore a previous version



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
	objectId := "objectId_example" // string | The ID of the blog post.
	revisionId := "revisionId_example" // string | The ID of the version to restore the blog post to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post. | 
**revisionId** | **string** | The ID of the version to restore the blog post to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft

> BlogPost PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft(ctx, objectId, revisionId).Execute()

Restore a draft to a previous version



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
	objectId := "objectId_example" // string | The ID of the blog post.
	revisionId := int64(789) // int64 | The ID of the version to restore the blog post to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the blog post. | 
**revisionId** | **int64** | The ID of the version to restore the blog post to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsScheduleSchedule

> PostCmsV3BlogsPostsScheduleSchedule(ctx).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()

Schedule a post to be published



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
	r, err := apiClient.BasicAPI.PostCmsV3BlogsPostsScheduleSchedule(context.Background()).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCmsV3BlogsPostsScheduleSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsScheduleScheduleRequest struct via the builder pattern


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

