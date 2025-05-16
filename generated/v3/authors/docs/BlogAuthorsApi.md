# \BlogAuthorsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3BlogsAuthorsObjectIdArchive**](BlogAuthorsAPI.md#DeleteCmsV3BlogsAuthorsObjectIdArchive) | **Delete** /cms/v3/blogs/authors/{objectId} | Delete a Blog Author
[**GetCmsV3BlogsAuthorsGetPage**](BlogAuthorsAPI.md#GetCmsV3BlogsAuthorsGetPage) | **Get** /cms/v3/blogs/authors | Get all Blog Authors
[**GetCmsV3BlogsAuthorsObjectIdGetById**](BlogAuthorsAPI.md#GetCmsV3BlogsAuthorsObjectIdGetById) | **Get** /cms/v3/blogs/authors/{objectId} | Retrieve a Blog Author
[**PatchCmsV3BlogsAuthorsObjectIdUpdate**](BlogAuthorsAPI.md#PatchCmsV3BlogsAuthorsObjectIdUpdate) | **Patch** /cms/v3/blogs/authors/{objectId} | Update a Blog Author
[**PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch) | **Post** /cms/v3/blogs/authors/batch/archive | Delete a batch of Blog Authors
[**PostCmsV3BlogsAuthorsBatchCreateCreateBatch**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsBatchCreateCreateBatch) | **Post** /cms/v3/blogs/authors/batch/create | Create a batch of Blog Authors
[**PostCmsV3BlogsAuthorsBatchReadReadBatch**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsBatchReadReadBatch) | **Post** /cms/v3/blogs/authors/batch/read | Retrieve a batch of Blog Authors
[**PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch) | **Post** /cms/v3/blogs/authors/batch/update | Update a batch of Blog Authors
[**PostCmsV3BlogsAuthorsCreate**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsCreate) | **Post** /cms/v3/blogs/authors | Create a new Blog Author
[**PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup) | **Post** /cms/v3/blogs/authors/multi-language/attach-to-lang-group | Attach a Blog Author to a multi-language group
[**PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation) | **Post** /cms/v3/blogs/authors/multi-language/create-language-variation | Create a new language variation
[**PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup) | **Post** /cms/v3/blogs/authors/multi-language/detach-from-lang-group | Detach a Blog Author from a multi-language group
[**PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs) | **Post** /cms/v3/blogs/authors/multi-language/update-languages | Update languages of multi-language group
[**PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary**](BlogAuthorsAPI.md#PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary) | **Put** /cms/v3/blogs/authors/multi-language/set-new-lang-primary | Set a new primary language



## DeleteCmsV3BlogsAuthorsObjectIdArchive

> DeleteCmsV3BlogsAuthorsObjectIdArchive(ctx, objectId).Archived(archived).Execute()

Delete a Blog Author



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
	objectId := "objectId_example" // string | The Blog Author id.
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogAuthorsAPI.DeleteCmsV3BlogsAuthorsObjectIdArchive(context.Background(), objectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.DeleteCmsV3BlogsAuthorsObjectIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3BlogsAuthorsObjectIdArchiveRequest struct via the builder pattern


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


## GetCmsV3BlogsAuthorsGetPage

> CollectionResponseWithTotalBlogAuthorForwardPaging GetCmsV3BlogsAuthorsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()

Get all Blog Authors



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
	createdAt := time.Now() // time.Time | Only return Blog Authors created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return Blog Authors created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return Blog Authors created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return Blog Authors last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return Blog Authors last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return Blog Authors last updated before the specified time. (optional)
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)
	archived := true // bool | Specifies whether to return deleted Blog Authors. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.GetCmsV3BlogsAuthorsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.GetCmsV3BlogsAuthorsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsAuthorsGetPage`: CollectionResponseWithTotalBlogAuthorForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.GetCmsV3BlogsAuthorsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsAuthorsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Blog Authors created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Blog Authors created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Blog Authors created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Blog Authors last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Blog Authors last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Blog Authors last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**CollectionResponseWithTotalBlogAuthorForwardPaging**](CollectionResponseWithTotalBlogAuthorForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsAuthorsObjectIdGetById

> BlogAuthor GetCmsV3BlogsAuthorsObjectIdGetById(ctx, objectId).Archived(archived).Property(property).Execute()

Retrieve a Blog Author



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
	objectId := "objectId_example" // string | The Blog Author id.
	archived := true // bool | Specifies whether to return deleted Blog Authors. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.GetCmsV3BlogsAuthorsObjectIdGetById(context.Background(), objectId).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.GetCmsV3BlogsAuthorsObjectIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsAuthorsObjectIdGetById`: BlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.GetCmsV3BlogsAuthorsObjectIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsAuthorsObjectIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsAuthorsObjectIdUpdate

> BlogAuthor PatchCmsV3BlogsAuthorsObjectIdUpdate(ctx, objectId).BlogAuthor(blogAuthor).Archived(archived).Execute()

Update a Blog Author



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
	objectId := "objectId_example" // string | The Blog Author id.
	blogAuthor := *openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example") // BlogAuthor | The JSON representation of the updated Blog Author.
	archived := true // bool | Specifies whether to update deleted Blog Authors. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PatchCmsV3BlogsAuthorsObjectIdUpdate(context.Background(), objectId).BlogAuthor(blogAuthor).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PatchCmsV3BlogsAuthorsObjectIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3BlogsAuthorsObjectIdUpdate`: BlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PatchCmsV3BlogsAuthorsObjectIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsAuthorsObjectIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogAuthor** | [**BlogAuthor**](BlogAuthor.md) | The JSON representation of the updated Blog Author. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch

> PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch(ctx).BatchInputString(batchInputString).Execute()

Delete a batch of Blog Authors



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Author ids.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch(context.Background()).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchArchiveArchiveBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Author ids. | 

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


## PostCmsV3BlogsAuthorsBatchCreateCreateBatch

> BatchResponseBlogAuthor PostCmsV3BlogsAuthorsBatchCreateCreateBatch(ctx).BatchInputBlogAuthor(batchInputBlogAuthor).Execute()

Create a batch of Blog Authors



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
	batchInputBlogAuthor := *openapiclient.NewBatchInputBlogAuthor([]openapiclient.BlogAuthor{*openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example")}) // BatchInputBlogAuthor | The JSON array of new Blog Authors to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchCreateCreateBatch(context.Background()).BatchInputBlogAuthor(batchInputBlogAuthor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchCreateCreateBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthorsBatchCreateCreateBatch`: BatchResponseBlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchCreateCreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchCreateCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputBlogAuthor** | [**BatchInputBlogAuthor**](BatchInputBlogAuthor.md) | The JSON array of new Blog Authors to create. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsBatchReadReadBatch

> BatchResponseBlogAuthor PostCmsV3BlogsAuthorsBatchReadReadBatch(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Blog Authors



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Author ids.
	archived := true // bool | Specifies whether to return deleted Blog Authors. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchReadReadBatch(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchReadReadBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthorsBatchReadReadBatch`: BatchResponseBlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchReadReadBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchReadReadBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Author ids. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch

> BatchResponseBlogAuthor PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Blog Authors



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
	batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | A JSON array of the JSON representations of the updated Blog Authors.
	archived := true // bool | Specifies whether to update deleted Blog Authors. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch`: BatchResponseBlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchUpdateUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Authors. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsCreate

> BlogAuthor PostCmsV3BlogsAuthorsCreate(ctx).BlogAuthor(blogAuthor).Execute()

Create a new Blog Author



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
	blogAuthor := *openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example") // BlogAuthor | The JSON representation of a new Blog Author.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsCreate(context.Background()).BlogAuthor(blogAuthor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthorsCreate`: BlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthorsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogAuthor** | [**BlogAuthor**](BlogAuthor.md) | The JSON representation of a new Blog Author. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup

> PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

Attach a Blog Author to a multi-language group



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
	r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation

> BlogAuthor PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation(ctx).BlogAuthorCloneRequestVNext(blogAuthorCloneRequestVNext).Execute()

Create a new language variation



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
	blogAuthorCloneRequestVNext := *openapiclient.NewBlogAuthorCloneRequestVNext("Id_example", *openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example")) // BlogAuthorCloneRequestVNext | The JSON representation of the ContentLanguageCloneRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation(context.Background()).BlogAuthorCloneRequestVNext(blogAuthorCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation`: BlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogAuthorCloneRequestVNext** | [**BlogAuthorCloneRequestVNext**](BlogAuthorCloneRequestVNext.md) | The JSON representation of the ContentLanguageCloneRequest object. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup

> PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

Detach a Blog Author from a multi-language group



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
	r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs

> PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

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
	r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangsRequest struct via the builder pattern


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


## PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary

> PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

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
	r, err := apiClient.BlogAuthorsAPI.PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimaryRequest struct via the builder pattern


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

