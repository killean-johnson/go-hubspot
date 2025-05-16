# \ListAssociationsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIds**](ListAssociationsAPI.md#DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIds) | **Delete** /marketing/v3/marketing-events/associations/{externalAccountId}/{externalEventId}/lists/{listId} | Disassociate a list from a marketing event
[**DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventId**](ListAssociationsAPI.md#DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventId) | **Delete** /marketing/v3/marketing-events/associations/{marketingEventId}/lists/{listId} | Disassociate a list from a marketing event
[**GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds**](ListAssociationsAPI.md#GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds) | **Get** /marketing/v3/marketing-events/associations/{externalAccountId}/{externalEventId}/lists | Get lists associated with a marketing event
[**GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId**](ListAssociationsAPI.md#GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId) | **Get** /marketing/v3/marketing-events/associations/{marketingEventId}/lists | Get lists associated with a marketing event
[**PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIds**](ListAssociationsAPI.md#PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIds) | **Put** /marketing/v3/marketing-events/associations/{externalAccountId}/{externalEventId}/lists/{listId} | Associate a list with a marketing event
[**PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventId**](ListAssociationsAPI.md#PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventId) | **Put** /marketing/v3/marketing-events/associations/{marketingEventId}/lists/{listId} | Associate a list with a marketing event



## DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIds

> DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIds(ctx, externalAccountId, externalEventId, listId).Execute()

Disassociate a list from a marketing event



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
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application.
	externalEventId := "externalEventId_example" // string | The id of the marketing event in the external event application.
	listId := "listId_example" // string | The ILS ID of the list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListAssociationsAPI.DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIds(context.Background(), externalAccountId, externalEventId, listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAssociationsAPI.DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application. | 
**externalEventId** | **string** | The id of the marketing event in the external event application. | 
**listId** | **string** | The ILS ID of the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIdsRequest struct via the builder pattern


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


## DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventId

> DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventId(ctx, marketingEventId, listId).Execute()

Disassociate a list from a marketing event



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
	marketingEventId := "marketingEventId_example" // string | The internal id of the marketing event in HubSpot.
	listId := "listId_example" // string | The ILS ID of the list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListAssociationsAPI.DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventId(context.Background(), marketingEventId, listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAssociationsAPI.DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingEventId** | **string** | The internal id of the marketing event in HubSpot. | 
**listId** | **string** | The ILS ID of the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventIdRequest struct via the builder pattern


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


## GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds

> CollectionResponseWithTotalPublicListNoPaging GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds(ctx, externalAccountId, externalEventId).Execute()

Get lists associated with a marketing event



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
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application
	externalEventId := "externalEventId_example" // string | The id of the marketing event in the external event application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAssociationsAPI.GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds(context.Background(), externalAccountId, externalEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAssociationsAPI.GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds`: CollectionResponseWithTotalPublicListNoPaging
	fmt.Fprintf(os.Stdout, "Response from `ListAssociationsAPI.GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application | 
**externalEventId** | **string** | The id of the marketing event in the external event application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponseWithTotalPublicListNoPaging**](CollectionResponseWithTotalPublicListNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId

> CollectionResponseWithTotalPublicListNoPaging GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId(ctx, marketingEventId).Execute()

Get lists associated with a marketing event



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
	marketingEventId := "marketingEventId_example" // string | The internal id of the marketing event in HubSpot.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAssociationsAPI.GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId(context.Background(), marketingEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAssociationsAPI.GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId`: CollectionResponseWithTotalPublicListNoPaging
	fmt.Fprintf(os.Stdout, "Response from `ListAssociationsAPI.GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingEventId** | **string** | The internal id of the marketing event in HubSpot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponseWithTotalPublicListNoPaging**](CollectionResponseWithTotalPublicListNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIds

> PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIds(ctx, externalAccountId, externalEventId, listId).Execute()

Associate a list with a marketing event



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
	externalAccountId := "externalAccountId_example" // string | The accountId that is associated with this marketing event in the external event application.
	externalEventId := "externalEventId_example" // string | The id of the marketing event in the external event application.
	listId := "listId_example" // string | The ILS ID of the list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListAssociationsAPI.PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIds(context.Background(), externalAccountId, externalEventId, listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAssociationsAPI.PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application. | 
**externalEventId** | **string** | The id of the marketing event in the external event application. | 
**listId** | **string** | The ILS ID of the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIdsRequest struct via the builder pattern


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


## PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventId

> PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventId(ctx, marketingEventId, listId).Execute()

Associate a list with a marketing event



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
	marketingEventId := "marketingEventId_example" // string | The internal id of the marketing event in HubSpot.
	listId := "listId_example" // string | The ILS ID of the list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListAssociationsAPI.PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventId(context.Background(), marketingEventId, listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAssociationsAPI.PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingEventId** | **string** | The internal id of the marketing event in HubSpot. | 
**listId** | **string** | The ILS ID of the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventIdRequest struct via the builder pattern


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

