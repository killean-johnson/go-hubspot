# \EventsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsV3EventsEventTypesGetTypes**](EventsAPI.md#GetEventsV3EventsEventTypesGetTypes) | **Get** /events/v3/events/event-types | Event Types
[**GetEventsV3EventsGetPage**](EventsAPI.md#GetEventsV3EventsGetPage) | **Get** /events/v3/events/ | Retrieve event data



## GetEventsV3EventsEventTypesGetTypes

> VisibleExternalEventTypeNames GetEventsV3EventsEventTypesGetTypes(ctx).Execute()

Event Types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEventsV3EventsEventTypesGetTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventsV3EventsEventTypesGetTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsV3EventsEventTypesGetTypes`: VisibleExternalEventTypeNames
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventsV3EventsEventTypesGetTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsV3EventsEventTypesGetTypesRequest struct via the builder pattern


### Return type

[**VisibleExternalEventTypeNames**](VisibleExternalEventTypeNames.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsV3EventsGetPage

> CollectionResponseExternalUnifiedEvent GetEventsV3EventsGetPage(ctx).ObjectType(objectType).EventType(eventType).After(after).Before(before).Limit(limit).Sort(sort).OccurredAfter(occurredAfter).OccurredBefore(occurredBefore).ObjectId(objectId).ObjectPropertyPropname(objectPropertyPropname).PropertyPropname(propertyPropname).Id(id).Execute()

Retrieve event data



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
	objectType := "objectType_example" // string | The type of CRM object to filter event instances on (e.g., `contact`). To retrieve event data for a specific CRM record, include the additional `objectId` query parameter (below).  (optional)
	eventType := "eventType_example" // string | The event type name. You can retrieve available event types using the [event types endpoint](#get-%2Fevents%2Fv3%2Fevents%2Fevent-types). (optional)
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	sort := []string{"Inner_example"} // []string | Sort direction based on the timestamp of the event instance, `ASCENDING` or `DESCENDING`. (optional)
	occurredAfter := time.Now() // time.Time | Filter for event data that occurred after a specific datetime. (optional)
	occurredBefore := time.Now() // time.Time | Filter for event data that occurred before a specific datetime. (optional)
	objectId := int64(789) // int64 | The ID of the CRM Object to filter event instances on. When including this parameter, you must also include the `objectType` parameter. (optional)
	objectPropertyPropname := map[string]interface{}{ ... } // map[string]interface{} | Instead of retrieving event data for a specific object by its ID, you can specify a unique identifier property. For contacts, you can use the `email` property. (e.g., `objectProperty.email=name@domain.com`). (optional)
	propertyPropname := map[string]interface{}{ ... } // map[string]interface{} | Filter for event completions that contain a specific value for an event property (e.g., `property.hs_city=portland`). For properties values with spaces, replaces spaces with `%20` or `+` (e.g., `property.hs_city=new+york`). (optional)
	id := []string{"Inner_example"} // []string | ID of an event instance. IDs are 1:1 with event instances. If you provide this filter and additional filters, the other filters must match the values on the event instance to yield results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEventsV3EventsGetPage(context.Background()).ObjectType(objectType).EventType(eventType).After(after).Before(before).Limit(limit).Sort(sort).OccurredAfter(occurredAfter).OccurredBefore(occurredBefore).ObjectId(objectId).ObjectPropertyPropname(objectPropertyPropname).PropertyPropname(propertyPropname).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventsV3EventsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsV3EventsGetPage`: CollectionResponseExternalUnifiedEvent
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventsV3EventsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsV3EventsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectType** | **string** | The type of CRM object to filter event instances on (e.g., &#x60;contact&#x60;). To retrieve event data for a specific CRM record, include the additional &#x60;objectId&#x60; query parameter (below).  | 
 **eventType** | **string** | The event type name. You can retrieve available event types using the [event types endpoint](#get-%2Fevents%2Fv3%2Fevents%2Fevent-types). | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **sort** | **[]string** | Sort direction based on the timestamp of the event instance, &#x60;ASCENDING&#x60; or &#x60;DESCENDING&#x60;. | 
 **occurredAfter** | **time.Time** | Filter for event data that occurred after a specific datetime. | 
 **occurredBefore** | **time.Time** | Filter for event data that occurred before a specific datetime. | 
 **objectId** | **int64** | The ID of the CRM Object to filter event instances on. When including this parameter, you must also include the &#x60;objectType&#x60; parameter. | 
 **objectPropertyPropname** | [**map[string]interface{}**](map[string]interface{}.md) | Instead of retrieving event data for a specific object by its ID, you can specify a unique identifier property. For contacts, you can use the &#x60;email&#x60; property. (e.g., &#x60;objectProperty.email&#x3D;name@domain.com&#x60;). | 
 **propertyPropname** | [**map[string]interface{}**](map[string]interface{}.md) | Filter for event completions that contain a specific value for an event property (e.g., &#x60;property.hs_city&#x3D;portland&#x60;). For properties values with spaces, replaces spaces with &#x60;%20&#x60; or &#x60;+&#x60; (e.g., &#x60;property.hs_city&#x3D;new+york&#x60;). | 
 **id** | **[]string** | ID of an event instance. IDs are 1:1 with event instances. If you provide this filter and additional filters, the other filters must match the values on the event instance to yield results. | 

### Return type

[**CollectionResponseExternalUnifiedEvent**](CollectionResponseExternalUnifiedEvent.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

