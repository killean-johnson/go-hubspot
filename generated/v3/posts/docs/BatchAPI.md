# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCmsV3BlogsPostsBatchArchiveArchive**](BatchAPI.md#PostCmsV3BlogsPostsBatchArchiveArchive) | **Post** /cms/v3/blogs/posts/batch/archive | Delete a batch of blog posts
[**PostCmsV3BlogsPostsBatchCreateCreate**](BatchAPI.md#PostCmsV3BlogsPostsBatchCreateCreate) | **Post** /cms/v3/blogs/posts/batch/create | Create a batch of blog posts
[**PostCmsV3BlogsPostsBatchReadRead**](BatchAPI.md#PostCmsV3BlogsPostsBatchReadRead) | **Post** /cms/v3/blogs/posts/batch/read | Retrieve a batch of Blog Posts
[**PostCmsV3BlogsPostsBatchUpdateUpdate**](BatchAPI.md#PostCmsV3BlogsPostsBatchUpdateUpdate) | **Post** /cms/v3/blogs/posts/batch/update | Update a batch of Blog Posts



## PostCmsV3BlogsPostsBatchArchiveArchive

> PostCmsV3BlogsPostsBatchArchiveArchive(ctx).BatchInputString(batchInputString).Execute()

Delete a batch of blog posts



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Post ids.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.PostCmsV3BlogsPostsBatchArchiveArchive(context.Background()).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCmsV3BlogsPostsBatchArchiveArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchArchiveArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 

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


## PostCmsV3BlogsPostsBatchCreateCreate

> BatchResponseBlogPost PostCmsV3BlogsPostsBatchCreateCreate(ctx).BatchInputBlogPost(batchInputBlogPost).Execute()

Create a batch of blog posts



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
	batchInputBlogPost := *openapiclient.NewBatchInputBlogPost([]openapiclient.BlogPost{*openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", "CampaignName_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", false, []map[string]interface{}{map[string]interface{}(123)}, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example")}) // BatchInputBlogPost | The JSON array of new Blog Posts to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostCmsV3BlogsPostsBatchCreateCreate(context.Background()).BatchInputBlogPost(batchInputBlogPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCmsV3BlogsPostsBatchCreateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsBatchCreateCreate`: BatchResponseBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCmsV3BlogsPostsBatchCreateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputBlogPost** | [**BatchInputBlogPost**](BatchInputBlogPost.md) | The JSON array of new Blog Posts to create. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsBatchReadRead

> BatchResponseBlogPost PostCmsV3BlogsPostsBatchReadRead(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Blog Posts



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Post ids.
	archived := true // bool | Specifies whether to return deleted blog posts Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostCmsV3BlogsPostsBatchReadRead(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCmsV3BlogsPostsBatchReadRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsBatchReadRead`: BatchResponseBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCmsV3BlogsPostsBatchReadRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchReadReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 
 **archived** | **bool** | Specifies whether to return deleted blog posts Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsBatchUpdateUpdate

> BatchResponseBlogPost PostCmsV3BlogsPostsBatchUpdateUpdate(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Blog Posts



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
	batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | A JSON array of the JSON representations of the updated Blog Posts.
	archived := true // bool | Specifies whether to update deleted Blog Posts. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostCmsV3BlogsPostsBatchUpdateUpdate(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCmsV3BlogsPostsBatchUpdateUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsBatchUpdateUpdate`: BatchResponseBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCmsV3BlogsPostsBatchUpdateUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchUpdateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Posts. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

