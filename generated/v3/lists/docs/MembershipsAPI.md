# \MembershipsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ListsListIdMembershipsRemoveAll**](MembershipsAPI.md#DeleteCrmV3ListsListIdMembershipsRemoveAll) | **Delete** /crm/v3/lists/{listId}/memberships | Delete All Records from a List
[**GetCrmV3ListsListIdMembershipsGetPage**](MembershipsAPI.md#GetCrmV3ListsListIdMembershipsGetPage) | **Get** /crm/v3/lists/{listId}/memberships | Fetch List Memberships Ordered by ID
[**GetCrmV3ListsListIdMembershipsJoinOrderGetPageOrderedByAddedToListDate**](MembershipsAPI.md#GetCrmV3ListsListIdMembershipsJoinOrderGetPageOrderedByAddedToListDate) | **Get** /crm/v3/lists/{listId}/memberships/join-order | Fetch List Memberships Ordered by Added to List Date
[**GetCrmV3ListsRecordsObjectTypeIdRecordIdMembershipsGetLists**](MembershipsAPI.md#GetCrmV3ListsRecordsObjectTypeIdRecordIdMembershipsGetLists) | **Get** /crm/v3/lists/records/{objectTypeId}/{recordId}/memberships | Get lists record is member of
[**PutCrmV3ListsListIdMembershipsAddAdd**](MembershipsAPI.md#PutCrmV3ListsListIdMembershipsAddAdd) | **Put** /crm/v3/lists/{listId}/memberships/add | Add Records to a List
[**PutCrmV3ListsListIdMembershipsAddAndRemoveAddAndRemove**](MembershipsAPI.md#PutCrmV3ListsListIdMembershipsAddAndRemoveAddAndRemove) | **Put** /crm/v3/lists/{listId}/memberships/add-and-remove | Add and/or Remove Records from a List
[**PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList**](MembershipsAPI.md#PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList) | **Put** /crm/v3/lists/{listId}/memberships/add-from/{sourceListId} | Add All Records from a Source List to a Destination List
[**PutCrmV3ListsListIdMembershipsRemoveRemove**](MembershipsAPI.md#PutCrmV3ListsListIdMembershipsRemoveRemove) | **Put** /crm/v3/lists/{listId}/memberships/remove | Remove Records from a List



## DeleteCrmV3ListsListIdMembershipsRemoveAll

> DeleteCrmV3ListsListIdMembershipsRemoveAll(ctx, listId).Execute()

Delete All Records from a List



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
	listId := "listId_example" // string | The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembershipsAPI.DeleteCrmV3ListsListIdMembershipsRemoveAll(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.DeleteCrmV3ListsListIdMembershipsRemoveAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ListsListIdMembershipsRemoveAllRequest struct via the builder pattern


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


## GetCrmV3ListsListIdMembershipsGetPage

> ApiCollectionResponseJoinTimeAndRecordId GetCrmV3ListsListIdMembershipsGetPage(ctx, listId).After(after).Before(before).Limit(limit).Execute()

Fetch List Memberships Ordered by ID



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
	listId := "listId_example" // string | The **ILS ID** of the list.
	after := "after_example" // string | The paging offset token for the page that comes `after` the previously requested records.  If provided, then the records in the response will be the records following the offset, sorted in *ascending* order. Takes precedence over the `before` offset. (optional)
	before := "before_example" // string | The paging offset token for the page that comes `before` the previously requested records.  If provided, then the records in the response will be the records preceding the offset, sorted in *descending* order. (optional)
	limit := int32(56) // int32 | The number of records to return in the response. The maximum `limit` is 250. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.GetCrmV3ListsListIdMembershipsGetPage(context.Background(), listId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetCrmV3ListsListIdMembershipsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ListsListIdMembershipsGetPage`: ApiCollectionResponseJoinTimeAndRecordId
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetCrmV3ListsListIdMembershipsGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The **ILS ID** of the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ListsListIdMembershipsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The paging offset token for the page that comes &#x60;after&#x60; the previously requested records.  If provided, then the records in the response will be the records following the offset, sorted in *ascending* order. Takes precedence over the &#x60;before&#x60; offset. | 
 **before** | **string** | The paging offset token for the page that comes &#x60;before&#x60; the previously requested records.  If provided, then the records in the response will be the records preceding the offset, sorted in *descending* order. | 
 **limit** | **int32** | The number of records to return in the response. The maximum &#x60;limit&#x60; is 250. | [default to 100]

### Return type

[**ApiCollectionResponseJoinTimeAndRecordId**](ApiCollectionResponseJoinTimeAndRecordId.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ListsListIdMembershipsJoinOrderGetPageOrderedByAddedToListDate

> ApiCollectionResponseJoinTimeAndRecordId GetCrmV3ListsListIdMembershipsJoinOrderGetPageOrderedByAddedToListDate(ctx, listId).After(after).Before(before).Limit(limit).Execute()

Fetch List Memberships Ordered by Added to List Date



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
	listId := "listId_example" // string | The **ILS ID** of the list.
	after := "after_example" // string | The paging offset token for the page that comes `after` the previously requested records.  If provided, then the records in the response will be the records following the offset, sorted in *ascending* order. Takes precedence over the `before` offset. (optional)
	before := "before_example" // string | The paging offset token for the page that comes `before` the previously requested records.  If provided, then the records in the response will be the records preceding the offset, sorted in *descending* order. (optional)
	limit := int32(56) // int32 | The number of records to return in the response. The maximum `limit` is 250. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.GetCrmV3ListsListIdMembershipsJoinOrderGetPageOrderedByAddedToListDate(context.Background(), listId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetCrmV3ListsListIdMembershipsJoinOrderGetPageOrderedByAddedToListDate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ListsListIdMembershipsJoinOrderGetPageOrderedByAddedToListDate`: ApiCollectionResponseJoinTimeAndRecordId
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetCrmV3ListsListIdMembershipsJoinOrderGetPageOrderedByAddedToListDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The **ILS ID** of the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ListsListIdMembershipsJoinOrderGetPageOrderedByAddedToListDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The paging offset token for the page that comes &#x60;after&#x60; the previously requested records.  If provided, then the records in the response will be the records following the offset, sorted in *ascending* order. Takes precedence over the &#x60;before&#x60; offset. | 
 **before** | **string** | The paging offset token for the page that comes &#x60;before&#x60; the previously requested records.  If provided, then the records in the response will be the records preceding the offset, sorted in *descending* order. | 
 **limit** | **int32** | The number of records to return in the response. The maximum &#x60;limit&#x60; is 250. | [default to 100]

### Return type

[**ApiCollectionResponseJoinTimeAndRecordId**](ApiCollectionResponseJoinTimeAndRecordId.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ListsRecordsObjectTypeIdRecordIdMembershipsGetLists

> ApiCollectionResponseRecordListMembershipNoPaging GetCrmV3ListsRecordsObjectTypeIdRecordIdMembershipsGetLists(ctx, objectTypeId, recordId).Execute()

Get lists record is member of



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
	objectTypeId := "objectTypeId_example" // string | Object type id of the record
	recordId := "recordId_example" // string | Id of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.GetCrmV3ListsRecordsObjectTypeIdRecordIdMembershipsGetLists(context.Background(), objectTypeId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetCrmV3ListsRecordsObjectTypeIdRecordIdMembershipsGetLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ListsRecordsObjectTypeIdRecordIdMembershipsGetLists`: ApiCollectionResponseRecordListMembershipNoPaging
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetCrmV3ListsRecordsObjectTypeIdRecordIdMembershipsGetLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectTypeId** | **string** | Object type id of the record | 
**recordId** | **string** | Id of the record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ListsRecordsObjectTypeIdRecordIdMembershipsGetListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiCollectionResponseRecordListMembershipNoPaging**](ApiCollectionResponseRecordListMembershipNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ListsListIdMembershipsAddAdd

> MembershipsUpdateResponse PutCrmV3ListsListIdMembershipsAddAdd(ctx, listId).RequestBody(requestBody).Execute()

Add Records to a List



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
	listId := "listId_example" // string | The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.PutCrmV3ListsListIdMembershipsAddAdd(context.Background(), listId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.PutCrmV3ListsListIdMembershipsAddAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3ListsListIdMembershipsAddAdd`: MembershipsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.PutCrmV3ListsListIdMembershipsAddAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ListsListIdMembershipsAddAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

[**MembershipsUpdateResponse**](MembershipsUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ListsListIdMembershipsAddAndRemoveAddAndRemove

> MembershipsUpdateResponse PutCrmV3ListsListIdMembershipsAddAndRemoveAddAndRemove(ctx, listId).MembershipChangeRequest(membershipChangeRequest).Execute()

Add and/or Remove Records from a List



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
	listId := "listId_example" // string | The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	membershipChangeRequest := *openapiclient.NewMembershipChangeRequest([]string{"RecordIdsToRemove_example"}, []string{"RecordIdsToAdd_example"}) // MembershipChangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.PutCrmV3ListsListIdMembershipsAddAndRemoveAddAndRemove(context.Background(), listId).MembershipChangeRequest(membershipChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.PutCrmV3ListsListIdMembershipsAddAndRemoveAddAndRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3ListsListIdMembershipsAddAndRemoveAddAndRemove`: MembershipsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.PutCrmV3ListsListIdMembershipsAddAndRemoveAddAndRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ListsListIdMembershipsAddAndRemoveAddAndRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membershipChangeRequest** | [**MembershipChangeRequest**](MembershipChangeRequest.md) |  | 

### Return type

[**MembershipsUpdateResponse**](MembershipsUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList

> PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList(ctx, listId, sourceListId).Execute()

Add All Records from a Source List to a Destination List



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
	listId := "listId_example" // string | The **ILS ID** of the `MANUAL` or `SNAPSHOT` *destination list*, which the *source list* records are added to.
	sourceListId := "sourceListId_example" // string | The **ILS ID** of the *source list* to grab the records from, which are then added to the *destination list*.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembershipsAPI.PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList(context.Background(), listId, sourceListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; *destination list*, which the *source list* records are added to. | 
**sourceListId** | **string** | The **ILS ID** of the *source list* to grab the records from, which are then added to the *destination list*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListRequest struct via the builder pattern


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


## PutCrmV3ListsListIdMembershipsRemoveRemove

> MembershipsUpdateResponse PutCrmV3ListsListIdMembershipsRemoveRemove(ctx, listId).RequestBody(requestBody).Execute()

Remove Records from a List



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
	listId := "listId_example" // string | The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.PutCrmV3ListsListIdMembershipsRemoveRemove(context.Background(), listId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.PutCrmV3ListsListIdMembershipsRemoveRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3ListsListIdMembershipsRemoveRemove`: MembershipsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.PutCrmV3ListsListIdMembershipsRemoveRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ListsListIdMembershipsRemoveRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

[**MembershipsUpdateResponse**](MembershipsUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

