# \RowsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftPurgeDraftTableRow**](RowsAPI.md#DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftPurgeDraftTableRow) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Permanently deletes a row
[**GetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRows**](RowsAPI.md#GetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRows) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft | Get rows from draft table
[**GetCmsV3HubdbTablesTableIdOrNameRowsGetTableRows**](RowsAPI.md#GetCmsV3HubdbTablesTableIdOrNameRowsGetTableRows) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Get rows for a table
[**GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowById**](RowsAPI.md#GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowById) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Get a row from the draft table
[**GetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRow**](RowsAPI.md#GetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRow) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId} | Get a table row
[**PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRow**](RowsAPI.md#PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRow) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Updates an existing row
[**PostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRow**](RowsAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRow) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Add a new row to a table
[**PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRow**](RowsAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRow) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft/clone | Clone a row
[**PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRow**](RowsAPI.md#PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRow) | **Put** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Replaces an existing row



## DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftPurgeDraftTableRow

> DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftPurgeDraftTableRow(ctx, tableIdOrName, rowId).Execute()

Permanently deletes a row



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	rowId := "rowId_example" // string | The ID of the row

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RowsAPI.DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftPurgeDraftTableRow(context.Background(), tableIdOrName, rowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftPurgeDraftTableRow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftPurgeDraftTableRowRequest struct via the builder pattern


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


## GetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRows

> UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3 GetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRows(ctx, tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Offset(offset).Archived(archived).Execute()

Get rows from draft table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to query.
	sort := []string{"Inner_example"} // []string | Specifies the column names to sort the results by. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is `1000`. (optional)
	properties := []string{"Inner_example"} // []string | Specify the column names to get results containing only the required columns instead of all column details. If you want to include multiple columns in the result, use this query param as many times.  (optional)
	offset := int32(56) // int32 |  (optional)
	archived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRows(context.Background(), tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Offset(offset).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRows`: UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **[]string** | Specifies the column names to sort the results by. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is &#x60;1000&#x60;. | 
 **properties** | **[]string** | Specify the column names to get results containing only the required columns instead of all column details. If you want to include multiple columns in the result, use this query param as many times.  | 
 **offset** | **int32** |  | 
 **archived** | **bool** |  | 

### Return type

[**UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3**](UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameRowsGetTableRows

> UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3 GetCmsV3HubdbTablesTableIdOrNameRowsGetTableRows(ctx, tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Offset(offset).Archived(archived).Execute()

Get rows for a table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to query.
	sort := []string{"Inner_example"} // []string | Specifies the column names to sort the results by. See the above description for more details. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is `1000`. (optional)
	properties := []string{"Inner_example"} // []string | Specify the column names to get results containing only the required columns instead of all column details. (optional)
	offset := int32(56) // int32 |  (optional)
	archived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsGetTableRows(context.Background(), tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Offset(offset).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsGetTableRows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameRowsGetTableRows`: UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsGetTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameRowsGetTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **[]string** | Specifies the column names to sort the results by. See the above description for more details. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is &#x60;1000&#x60;. | 
 **properties** | **[]string** | Specify the column names to get results containing only the required columns instead of all column details. | 
 **offset** | **int32** |  | 
 **archived** | **bool** |  | 

### Return type

[**UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3**](UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowById

> HubDbTableRowV3 GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowById(ctx, tableIdOrName, rowId).Archived(archived).Execute()

Get a row from the draft table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	rowId := "rowId_example" // string | The ID of the row
	archived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowById(context.Background(), tableIdOrName, rowId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowById`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** |  | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRow

> HubDbTableRowV3 GetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRow(ctx, tableIdOrName, rowId).Archived(archived).Execute()

Get a table row



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	rowId := "rowId_example" // string | The ID of the row
	archived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRow(context.Background(), tableIdOrName, rowId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRow`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** |  | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRow

> HubDbTableRowV3 PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRow(ctx, tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()

Updates an existing row



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	rowId := "rowId_example" // string | The ID of the row
	hubDbTableRowV3Request := *openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // HubDbTableRowV3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRow(context.Background(), tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRow`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hubDbTableRowV3Request** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md) |  | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRow

> HubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRow(ctx, tableIdOrName).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()

Add a new row to a table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the target table.
	hubDbTableRowV3Request := *openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // HubDbTableRowV3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRow(context.Background(), tableIdOrName).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRow`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the target table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubDbTableRowV3Request** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md) |  | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRow

> HubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRow(ctx, tableIdOrName, rowId).Name(name).Execute()

Clone a row



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	rowId := "rowId_example" // string | The ID of the row
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRow(context.Background(), tableIdOrName, rowId).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRow`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** |  | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRow

> HubDbTableRowV3 PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRow(ctx, tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()

Replaces an existing row



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	rowId := "rowId_example" // string | The ID of the row
	hubDbTableRowV3Request := *openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // HubDbTableRowV3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRow(context.Background(), tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRow`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hubDbTableRowV3Request** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md) |  | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

