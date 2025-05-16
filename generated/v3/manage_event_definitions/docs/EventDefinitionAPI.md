# \EventDefinitionAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEventsV3EventDefinitionsEventName**](EventDefinitionAPI.md#DeleteEventsV3EventDefinitionsEventName) | **Delete** /events/v3/event-definitions/{eventName} | Delete an event definition
[**DeleteEventsV3EventDefinitionsEventNamePropertyPropertyName**](EventDefinitionAPI.md#DeleteEventsV3EventDefinitionsEventNamePropertyPropertyName) | **Delete** /events/v3/event-definitions/{eventName}/property/{propertyName} | Delete a property from a custom event definition
[**GetEventsV3EventDefinitions**](EventDefinitionAPI.md#GetEventsV3EventDefinitions) | **Get** /events/v3/event-definitions | Retrieve event definitions
[**GetEventsV3EventDefinitionsEventName**](EventDefinitionAPI.md#GetEventsV3EventDefinitionsEventName) | **Get** /events/v3/event-definitions/{eventName} | Retrieve a custom event definition
[**PatchEventsV3EventDefinitionsEventName**](EventDefinitionAPI.md#PatchEventsV3EventDefinitionsEventName) | **Patch** /events/v3/event-definitions/{eventName} | Update a custom event definition
[**PatchEventsV3EventDefinitionsEventNamePropertyPropertyName**](EventDefinitionAPI.md#PatchEventsV3EventDefinitionsEventNamePropertyPropertyName) | **Patch** /events/v3/event-definitions/{eventName}/property/{propertyName} | Update an existing custom event property
[**PostEventsV3EventDefinitions**](EventDefinitionAPI.md#PostEventsV3EventDefinitions) | **Post** /events/v3/event-definitions | Create custom event definition
[**PostEventsV3EventDefinitionsEventNameProperty**](EventDefinitionAPI.md#PostEventsV3EventDefinitionsEventNameProperty) | **Post** /events/v3/event-definitions/{eventName}/property | Create a property for an event definition



## DeleteEventsV3EventDefinitionsEventName

> DeleteEventsV3EventDefinitionsEventName(ctx, eventName).Execute()

Delete an event definition



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
	eventName := "eventName_example" // string | The name of the event definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventDefinitionAPI.DeleteEventsV3EventDefinitionsEventName(context.Background(), eventName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventDefinitionAPI.DeleteEventsV3EventDefinitionsEventName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** | The name of the event definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventsV3EventDefinitionsEventNameRequest struct via the builder pattern


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


## DeleteEventsV3EventDefinitionsEventNamePropertyPropertyName

> DeleteEventsV3EventDefinitionsEventNamePropertyPropertyName(ctx, eventName, propertyName).Execute()

Delete a property from a custom event definition



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
	eventName := "eventName_example" // string | The internal name of the custom event.
	propertyName := "propertyName_example" // string | The internal name of the property to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventDefinitionAPI.DeleteEventsV3EventDefinitionsEventNamePropertyPropertyName(context.Background(), eventName, propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventDefinitionAPI.DeleteEventsV3EventDefinitionsEventNamePropertyPropertyName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** | The internal name of the custom event. | 
**propertyName** | **string** | The internal name of the property to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventsV3EventDefinitionsEventNamePropertyPropertyNameRequest struct via the builder pattern


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


## GetEventsV3EventDefinitions

> CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging GetEventsV3EventDefinitions(ctx).SearchString(searchString).After(after).Limit(limit).IncludeProperties(includeProperties).SortOrder(sortOrder).Execute()

Retrieve event definitions



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
	searchString := "searchString_example" // string | Characters in the event name that the user is searching for. This search is a naive “contains” search, no fuzzy matching is done. (optional)
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	includeProperties := true // bool |  (optional)
	sortOrder := "sortOrder_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventDefinitionAPI.GetEventsV3EventDefinitions(context.Background()).SearchString(searchString).After(after).Limit(limit).IncludeProperties(includeProperties).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventDefinitionAPI.GetEventsV3EventDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsV3EventDefinitions`: CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `EventDefinitionAPI.GetEventsV3EventDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsV3EventDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchString** | **string** | Characters in the event name that the user is searching for. This search is a naive “contains” search, no fuzzy matching is done. | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **includeProperties** | **bool** |  | 
 **sortOrder** | **string** |  | 

### Return type

[**CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging**](CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsV3EventDefinitionsEventName

> ExternalBehavioralEventTypeDefinition GetEventsV3EventDefinitionsEventName(ctx, eventName).Execute()

Retrieve a custom event definition



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
	eventName := "eventName_example" // string | The internal name of the custom event.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventDefinitionAPI.GetEventsV3EventDefinitionsEventName(context.Background(), eventName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventDefinitionAPI.GetEventsV3EventDefinitionsEventName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsV3EventDefinitionsEventName`: ExternalBehavioralEventTypeDefinition
	fmt.Fprintf(os.Stdout, "Response from `EventDefinitionAPI.GetEventsV3EventDefinitionsEventName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** | The internal name of the custom event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsV3EventDefinitionsEventNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalBehavioralEventTypeDefinition**](ExternalBehavioralEventTypeDefinition.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEventsV3EventDefinitionsEventName

> ExternalBehavioralEventTypeDefinition PatchEventsV3EventDefinitionsEventName(ctx, eventName).ExternalBehavioralEventTypeDefinitionPatch(externalBehavioralEventTypeDefinitionPatch).Execute()

Update a custom event definition



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
	eventName := "eventName_example" // string | The internal name of the custom event.
	externalBehavioralEventTypeDefinitionPatch := *openapiclient.NewExternalBehavioralEventTypeDefinitionPatch() // ExternalBehavioralEventTypeDefinitionPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventDefinitionAPI.PatchEventsV3EventDefinitionsEventName(context.Background(), eventName).ExternalBehavioralEventTypeDefinitionPatch(externalBehavioralEventTypeDefinitionPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventDefinitionAPI.PatchEventsV3EventDefinitionsEventName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEventsV3EventDefinitionsEventName`: ExternalBehavioralEventTypeDefinition
	fmt.Fprintf(os.Stdout, "Response from `EventDefinitionAPI.PatchEventsV3EventDefinitionsEventName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** | The internal name of the custom event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEventsV3EventDefinitionsEventNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalBehavioralEventTypeDefinitionPatch** | [**ExternalBehavioralEventTypeDefinitionPatch**](ExternalBehavioralEventTypeDefinitionPatch.md) |  | 

### Return type

[**ExternalBehavioralEventTypeDefinition**](ExternalBehavioralEventTypeDefinition.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEventsV3EventDefinitionsEventNamePropertyPropertyName

> Property PatchEventsV3EventDefinitionsEventNamePropertyPropertyName(ctx, eventName, propertyName).ExternalBehavioralEventPropertyDefinitionPatch(externalBehavioralEventPropertyDefinitionPatch).Execute()

Update an existing custom event property



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
	eventName := "eventName_example" // string | The internal name of the custom event.
	propertyName := "propertyName_example" // string | The internal name of the property to update.
	externalBehavioralEventPropertyDefinitionPatch := *openapiclient.NewExternalBehavioralEventPropertyDefinitionPatch() // ExternalBehavioralEventPropertyDefinitionPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventDefinitionAPI.PatchEventsV3EventDefinitionsEventNamePropertyPropertyName(context.Background(), eventName, propertyName).ExternalBehavioralEventPropertyDefinitionPatch(externalBehavioralEventPropertyDefinitionPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventDefinitionAPI.PatchEventsV3EventDefinitionsEventNamePropertyPropertyName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEventsV3EventDefinitionsEventNamePropertyPropertyName`: Property
	fmt.Fprintf(os.Stdout, "Response from `EventDefinitionAPI.PatchEventsV3EventDefinitionsEventNamePropertyPropertyName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** | The internal name of the custom event. | 
**propertyName** | **string** | The internal name of the property to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEventsV3EventDefinitionsEventNamePropertyPropertyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalBehavioralEventPropertyDefinitionPatch** | [**ExternalBehavioralEventPropertyDefinitionPatch**](ExternalBehavioralEventPropertyDefinitionPatch.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEventsV3EventDefinitions

> ExternalBehavioralEventTypeDefinition PostEventsV3EventDefinitions(ctx).ExternalBehavioralEventTypeDefinitionEgg(externalBehavioralEventTypeDefinitionEgg).Execute()

Create custom event definition



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
	externalBehavioralEventTypeDefinitionEgg := *openapiclient.NewExternalBehavioralEventTypeDefinitionEgg([]openapiclient.ExternalBehavioralEventPropertyCreate{*openapiclient.NewExternalBehavioralEventPropertyCreate("Label_example", "Type_example")}, "Label_example") // ExternalBehavioralEventTypeDefinitionEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventDefinitionAPI.PostEventsV3EventDefinitions(context.Background()).ExternalBehavioralEventTypeDefinitionEgg(externalBehavioralEventTypeDefinitionEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventDefinitionAPI.PostEventsV3EventDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEventsV3EventDefinitions`: ExternalBehavioralEventTypeDefinition
	fmt.Fprintf(os.Stdout, "Response from `EventDefinitionAPI.PostEventsV3EventDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEventsV3EventDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalBehavioralEventTypeDefinitionEgg** | [**ExternalBehavioralEventTypeDefinitionEgg**](ExternalBehavioralEventTypeDefinitionEgg.md) |  | 

### Return type

[**ExternalBehavioralEventTypeDefinition**](ExternalBehavioralEventTypeDefinition.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEventsV3EventDefinitionsEventNameProperty

> Property PostEventsV3EventDefinitionsEventNameProperty(ctx, eventName).ExternalBehavioralEventPropertyCreate(externalBehavioralEventPropertyCreate).Execute()

Create a property for an event definition



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
	eventName := "eventName_example" // string | The internal name of the custom event.
	externalBehavioralEventPropertyCreate := *openapiclient.NewExternalBehavioralEventPropertyCreate("Label_example", "Type_example") // ExternalBehavioralEventPropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventDefinitionAPI.PostEventsV3EventDefinitionsEventNameProperty(context.Background(), eventName).ExternalBehavioralEventPropertyCreate(externalBehavioralEventPropertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventDefinitionAPI.PostEventsV3EventDefinitionsEventNameProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEventsV3EventDefinitionsEventNameProperty`: Property
	fmt.Fprintf(os.Stdout, "Response from `EventDefinitionAPI.PostEventsV3EventDefinitionsEventNameProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** | The internal name of the custom event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEventsV3EventDefinitionsEventNamePropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalBehavioralEventPropertyCreate** | [**ExternalBehavioralEventPropertyCreate**](ExternalBehavioralEventPropertyCreate.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

