# \RedirectsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3UrlRedirectsUrlRedirectIdArchive**](RedirectsAPI.md#DeleteCmsV3UrlRedirectsUrlRedirectIdArchive) | **Delete** /cms/v3/url-redirects/{urlRedirectId} | Delete a redirect
[**GetCmsV3UrlRedirectsGetPage**](RedirectsAPI.md#GetCmsV3UrlRedirectsGetPage) | **Get** /cms/v3/url-redirects/ | Get current redirects
[**GetCmsV3UrlRedirectsUrlRedirectIdGetById**](RedirectsAPI.md#GetCmsV3UrlRedirectsUrlRedirectIdGetById) | **Get** /cms/v3/url-redirects/{urlRedirectId} | Get details for a redirect
[**PatchCmsV3UrlRedirectsUrlRedirectIdUpdate**](RedirectsAPI.md#PatchCmsV3UrlRedirectsUrlRedirectIdUpdate) | **Patch** /cms/v3/url-redirects/{urlRedirectId} | Update a redirect
[**PostCmsV3UrlRedirectsCreate**](RedirectsAPI.md#PostCmsV3UrlRedirectsCreate) | **Post** /cms/v3/url-redirects/ | Create a redirect



## DeleteCmsV3UrlRedirectsUrlRedirectIdArchive

> DeleteCmsV3UrlRedirectsUrlRedirectIdArchive(ctx, urlRedirectId).Execute()

Delete a redirect



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
	urlRedirectId := "urlRedirectId_example" // string | The ID of the target redirect.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RedirectsAPI.DeleteCmsV3UrlRedirectsUrlRedirectIdArchive(context.Background(), urlRedirectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectsAPI.DeleteCmsV3UrlRedirectsUrlRedirectIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlRedirectId** | **string** | The ID of the target redirect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3UrlRedirectsUrlRedirectIdArchiveRequest struct via the builder pattern


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


## GetCmsV3UrlRedirectsGetPage

> CollectionResponseWithTotalUrlMappingForwardPaging GetCmsV3UrlRedirectsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()

Get current redirects



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
	createdAt := time.Now() // time.Time | Only return redirects created on exactly this date. (optional)
	createdAfter := time.Now() // time.Time | Only return redirects created after this date. (optional)
	createdBefore := time.Now() // time.Time | Only return redirects created before this date. (optional)
	updatedAt := time.Now() // time.Time | Only return redirects last updated on exactly this date. (optional)
	updatedAfter := time.Now() // time.Time | Only return redirects last updated after this date. (optional)
	updatedBefore := time.Now() // time.Time | Only return redirects last updated before this date. (optional)
	sort := []string{"Inner_example"} // []string |  (optional)
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | Maximum number of result per page (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectsAPI.GetCmsV3UrlRedirectsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectsAPI.GetCmsV3UrlRedirectsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3UrlRedirectsGetPage`: CollectionResponseWithTotalUrlMappingForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `RedirectsAPI.GetCmsV3UrlRedirectsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3UrlRedirectsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return redirects created on exactly this date. | 
 **createdAfter** | **time.Time** | Only return redirects created after this date. | 
 **createdBefore** | **time.Time** | Only return redirects created before this date. | 
 **updatedAt** | **time.Time** | Only return redirects last updated on exactly this date. | 
 **updatedAfter** | **time.Time** | Only return redirects last updated after this date. | 
 **updatedBefore** | **time.Time** | Only return redirects last updated before this date. | 
 **sort** | **[]string** |  | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | Maximum number of result per page | 
 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

[**CollectionResponseWithTotalUrlMappingForwardPaging**](CollectionResponseWithTotalUrlMappingForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3UrlRedirectsUrlRedirectIdGetById

> UrlMapping GetCmsV3UrlRedirectsUrlRedirectIdGetById(ctx, urlRedirectId).Execute()

Get details for a redirect



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
	urlRedirectId := "urlRedirectId_example" // string | The ID of the target redirect.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectsAPI.GetCmsV3UrlRedirectsUrlRedirectIdGetById(context.Background(), urlRedirectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectsAPI.GetCmsV3UrlRedirectsUrlRedirectIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3UrlRedirectsUrlRedirectIdGetById`: UrlMapping
	fmt.Fprintf(os.Stdout, "Response from `RedirectsAPI.GetCmsV3UrlRedirectsUrlRedirectIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlRedirectId** | **string** | The ID of the target redirect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3UrlRedirectsUrlRedirectIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3UrlRedirectsUrlRedirectIdUpdate

> UrlMapping PatchCmsV3UrlRedirectsUrlRedirectIdUpdate(ctx, urlRedirectId).UrlMapping(urlMapping).Execute()

Update a redirect



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
	urlRedirectId := "urlRedirectId_example" // string | 
	urlMapping := *openapiclient.NewUrlMapping(false, int32(123), false, false, "Destination_example", false, false, int32(123), "RoutePrefix_example", false, "Id_example") // UrlMapping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectsAPI.PatchCmsV3UrlRedirectsUrlRedirectIdUpdate(context.Background(), urlRedirectId).UrlMapping(urlMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectsAPI.PatchCmsV3UrlRedirectsUrlRedirectIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3UrlRedirectsUrlRedirectIdUpdate`: UrlMapping
	fmt.Fprintf(os.Stdout, "Response from `RedirectsAPI.PatchCmsV3UrlRedirectsUrlRedirectIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlRedirectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3UrlRedirectsUrlRedirectIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlMapping** | [**UrlMapping**](UrlMapping.md) |  | 

### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3UrlRedirectsCreate

> UrlMapping PostCmsV3UrlRedirectsCreate(ctx).UrlMappingCreateRequestBody(urlMappingCreateRequestBody).Execute()

Create a redirect



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
	urlMappingCreateRequestBody := *openapiclient.NewUrlMappingCreateRequestBody(int32(123), "RoutePrefix_example", "Destination_example") // UrlMappingCreateRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectsAPI.PostCmsV3UrlRedirectsCreate(context.Background()).UrlMappingCreateRequestBody(urlMappingCreateRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectsAPI.PostCmsV3UrlRedirectsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3UrlRedirectsCreate`: UrlMapping
	fmt.Fprintf(os.Stdout, "Response from `RedirectsAPI.PostCmsV3UrlRedirectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3UrlRedirectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlMappingCreateRequestBody** | [**UrlMappingCreateRequestBody**](UrlMappingCreateRequestBody.md) |  | 

### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

