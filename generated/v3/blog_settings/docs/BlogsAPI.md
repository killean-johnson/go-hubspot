# \BlogsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsV3BlogSettingsSettings**](BlogsAPI.md#GetCmsV3BlogSettingsSettings) | **Get** /cms/v3/blog-settings/settings | Get all Blogs
[**GetCmsV3BlogSettingsSettingsBlogId**](BlogsAPI.md#GetCmsV3BlogSettingsSettingsBlogId) | **Get** /cms/v3/blog-settings/settings/{blogId} | Retrieve a Blog
[**GetCmsV3BlogSettingsSettingsBlogIdRevisions**](BlogsAPI.md#GetCmsV3BlogSettingsSettingsBlogIdRevisions) | **Get** /cms/v3/blog-settings/settings/{blogId}/revisions | Retrieves all the previous versions of a Blog
[**GetCmsV3BlogSettingsSettingsBlogIdRevisionsRevisionId**](BlogsAPI.md#GetCmsV3BlogSettingsSettingsBlogIdRevisionsRevisionId) | **Get** /cms/v3/blog-settings/settings/{blogId}/revisions/{revisionId} | Retrieves a previous version of a Blog
[**PostCmsV3BlogSettingsSettingsMultiLanguageAttachToLangGroup**](BlogsAPI.md#PostCmsV3BlogSettingsSettingsMultiLanguageAttachToLangGroup) | **Post** /cms/v3/blog-settings/settings/multi-language/attach-to-lang-group | Attach a blog to a multi-language group
[**PostCmsV3BlogSettingsSettingsMultiLanguageCreateLanguageVariation**](BlogsAPI.md#PostCmsV3BlogSettingsSettingsMultiLanguageCreateLanguageVariation) | **Post** /cms/v3/blog-settings/settings/multi-language/create-language-variation | Create a new language variation
[**PostCmsV3BlogSettingsSettingsMultiLanguageDetachFromLangGroup**](BlogsAPI.md#PostCmsV3BlogSettingsSettingsMultiLanguageDetachFromLangGroup) | **Post** /cms/v3/blog-settings/settings/multi-language/detach-from-lang-group | Detach a blog from a multi-language group
[**PostCmsV3BlogSettingsSettingsMultiLanguageUpdateLanguages**](BlogsAPI.md#PostCmsV3BlogSettingsSettingsMultiLanguageUpdateLanguages) | **Post** /cms/v3/blog-settings/settings/multi-language/update-languages | Update languages of multi-language group
[**PutCmsV3BlogSettingsSettingsMultiLanguageSetNewLangPrimary**](BlogsAPI.md#PutCmsV3BlogSettingsSettingsMultiLanguageSetNewLangPrimary) | **Put** /cms/v3/blog-settings/settings/multi-language/set-new-lang-primary | Set a new primary language



## GetCmsV3BlogSettingsSettings

> CollectionResponseWithTotalBlogForwardPaging GetCmsV3BlogSettingsSettings(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()

Get all Blogs



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
	createdAt := time.Now() // time.Time | Only return Blogs created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return Blogs created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return Blogs created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return Blogs last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return Blogs last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return Blogs last updated before the specified time. (optional)
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name` and `id` (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)
	archived := true // bool | Specifies whether to return archived Blogs. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogsAPI.GetCmsV3BlogSettingsSettings(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogsAPI.GetCmsV3BlogSettingsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogSettingsSettings`: CollectionResponseWithTotalBlogForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BlogsAPI.GetCmsV3BlogSettingsSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogSettingsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Blogs created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Blogs created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Blogs created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Blogs last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Blogs last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Blogs last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60; and &#x60;id&#x60; | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 
 **archived** | **bool** | Specifies whether to return archived Blogs. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalBlogForwardPaging**](CollectionResponseWithTotalBlogForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogSettingsSettingsBlogId

> Blog GetCmsV3BlogSettingsSettingsBlogId(ctx, blogId).Execute()

Retrieve a Blog



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
	blogId := "blogId_example" // string | The Blog id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogsAPI.GetCmsV3BlogSettingsSettingsBlogId(context.Background(), blogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogsAPI.GetCmsV3BlogSettingsSettingsBlogId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogSettingsSettingsBlogId`: Blog
	fmt.Fprintf(os.Stdout, "Response from `BlogsAPI.GetCmsV3BlogSettingsSettingsBlogId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogId** | **string** | The Blog id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogSettingsSettingsBlogIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Blog**](Blog.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogSettingsSettingsBlogIdRevisions

> CollectionResponseWithTotalVersionBlog GetCmsV3BlogSettingsSettingsBlogIdRevisions(ctx, blogId).After(after).Before(before).Limit(limit).Execute()

Retrieves all the previous versions of a Blog



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
	blogId := "blogId_example" // string | The Blog id.
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogsAPI.GetCmsV3BlogSettingsSettingsBlogIdRevisions(context.Background(), blogId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogsAPI.GetCmsV3BlogSettingsSettingsBlogIdRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogSettingsSettingsBlogIdRevisions`: CollectionResponseWithTotalVersionBlog
	fmt.Fprintf(os.Stdout, "Response from `BlogsAPI.GetCmsV3BlogSettingsSettingsBlogIdRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogId** | **string** | The Blog id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogSettingsSettingsBlogIdRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 

### Return type

[**CollectionResponseWithTotalVersionBlog**](CollectionResponseWithTotalVersionBlog.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogSettingsSettingsBlogIdRevisionsRevisionId

> VersionBlog GetCmsV3BlogSettingsSettingsBlogIdRevisionsRevisionId(ctx, blogId, revisionId).Execute()

Retrieves a previous version of a Blog



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
	blogId := "blogId_example" // string | The Blog id.
	revisionId := "revisionId_example" // string | The Blog version id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogsAPI.GetCmsV3BlogSettingsSettingsBlogIdRevisionsRevisionId(context.Background(), blogId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogsAPI.GetCmsV3BlogSettingsSettingsBlogIdRevisionsRevisionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogSettingsSettingsBlogIdRevisionsRevisionId`: VersionBlog
	fmt.Fprintf(os.Stdout, "Response from `BlogsAPI.GetCmsV3BlogSettingsSettingsBlogIdRevisionsRevisionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogId** | **string** | The Blog id. | 
**revisionId** | **string** | The Blog version id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogSettingsSettingsBlogIdRevisionsRevisionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionBlog**](VersionBlog.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogSettingsSettingsMultiLanguageAttachToLangGroup

> PostCmsV3BlogSettingsSettingsMultiLanguageAttachToLangGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

Attach a blog to a multi-language group



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
	r, err := apiClient.BlogsAPI.PostCmsV3BlogSettingsSettingsMultiLanguageAttachToLangGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogsAPI.PostCmsV3BlogSettingsSettingsMultiLanguageAttachToLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogSettingsSettingsMultiLanguageAttachToLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogSettingsSettingsMultiLanguageCreateLanguageVariation

> Blog PostCmsV3BlogSettingsSettingsMultiLanguageCreateLanguageVariation(ctx).BlogLanguageCloneRequestVNext(blogLanguageCloneRequestVNext).Execute()

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
	blogLanguageCloneRequestVNext := *openapiclient.NewBlogLanguageCloneRequestVNext("Id_example") // BlogLanguageCloneRequestVNext | The JSON representation of the ContentLanguageCloneRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogsAPI.PostCmsV3BlogSettingsSettingsMultiLanguageCreateLanguageVariation(context.Background()).BlogLanguageCloneRequestVNext(blogLanguageCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogsAPI.PostCmsV3BlogSettingsSettingsMultiLanguageCreateLanguageVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogSettingsSettingsMultiLanguageCreateLanguageVariation`: Blog
	fmt.Fprintf(os.Stdout, "Response from `BlogsAPI.PostCmsV3BlogSettingsSettingsMultiLanguageCreateLanguageVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogSettingsSettingsMultiLanguageCreateLanguageVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogLanguageCloneRequestVNext** | [**BlogLanguageCloneRequestVNext**](BlogLanguageCloneRequestVNext.md) | The JSON representation of the ContentLanguageCloneRequest object. | 

### Return type

[**Blog**](Blog.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogSettingsSettingsMultiLanguageDetachFromLangGroup

> PostCmsV3BlogSettingsSettingsMultiLanguageDetachFromLangGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

Detach a blog from a multi-language group



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
	r, err := apiClient.BlogsAPI.PostCmsV3BlogSettingsSettingsMultiLanguageDetachFromLangGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogsAPI.PostCmsV3BlogSettingsSettingsMultiLanguageDetachFromLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogSettingsSettingsMultiLanguageDetachFromLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogSettingsSettingsMultiLanguageUpdateLanguages

> PostCmsV3BlogSettingsSettingsMultiLanguageUpdateLanguages(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

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
	r, err := apiClient.BlogsAPI.PostCmsV3BlogSettingsSettingsMultiLanguageUpdateLanguages(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogsAPI.PostCmsV3BlogSettingsSettingsMultiLanguageUpdateLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogSettingsSettingsMultiLanguageUpdateLanguagesRequest struct via the builder pattern


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


## PutCmsV3BlogSettingsSettingsMultiLanguageSetNewLangPrimary

> PutCmsV3BlogSettingsSettingsMultiLanguageSetNewLangPrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

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
	r, err := apiClient.BlogsAPI.PutCmsV3BlogSettingsSettingsMultiLanguageSetNewLangPrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogsAPI.PutCmsV3BlogSettingsSettingsMultiLanguageSetNewLangPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3BlogSettingsSettingsMultiLanguageSetNewLangPrimaryRequest struct via the builder pattern


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

