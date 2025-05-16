# \SubscriptionStatusAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommunicationPreferencesV4StatusesSubscriberIdString**](SubscriptionStatusAPI.md#GetCommunicationPreferencesV4StatusesSubscriberIdString) | **Get** /communication-preferences/v4/statuses/{subscriberIdString} | Get subscription preferences for a specific contact
[**GetCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll**](SubscriptionStatusAPI.md#GetCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll) | **Get** /communication-preferences/v4/statuses/{subscriberIdString}/unsubscribe-all | Retrieve a contact&#39;s unsubscribed status
[**PostCommunicationPreferencesV4StatusesBatchRead**](SubscriptionStatusAPI.md#PostCommunicationPreferencesV4StatusesBatchRead) | **Post** /communication-preferences/v4/statuses/batch/read | Batch retrieve subscription statuses
[**PostCommunicationPreferencesV4StatusesBatchUnsubscribeAll**](SubscriptionStatusAPI.md#PostCommunicationPreferencesV4StatusesBatchUnsubscribeAll) | **Post** /communication-preferences/v4/statuses/batch/unsubscribe-all | Batch unsubscribe contacts from all subscriptions
[**PostCommunicationPreferencesV4StatusesBatchUnsubscribeAllRead**](SubscriptionStatusAPI.md#PostCommunicationPreferencesV4StatusesBatchUnsubscribeAllRead) | **Post** /communication-preferences/v4/statuses/batch/unsubscribe-all/read | Batch retrieve contacts who have opted out of all communications
[**PostCommunicationPreferencesV4StatusesBatchWrite**](SubscriptionStatusAPI.md#PostCommunicationPreferencesV4StatusesBatchWrite) | **Post** /communication-preferences/v4/statuses/batch/write | Batch update subscription status
[**PostCommunicationPreferencesV4StatusesSubscriberIdString**](SubscriptionStatusAPI.md#PostCommunicationPreferencesV4StatusesSubscriberIdString) | **Post** /communication-preferences/v4/statuses/{subscriberIdString} | Update a contact&#39;s subscription status
[**PostCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll**](SubscriptionStatusAPI.md#PostCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll) | **Post** /communication-preferences/v4/statuses/{subscriberIdString}/unsubscribe-all | Unsubscribe a contact from all subscriptions



## GetCommunicationPreferencesV4StatusesSubscriberIdString

> ActionResponseWithResultsPublicStatus GetCommunicationPreferencesV4StatusesSubscriberIdString(ctx, subscriberIdString).Channel(channel).BusinessUnitId(businessUnitId).Execute()

Get subscription preferences for a specific contact



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
	subscriberIdString := "subscriberIdString_example" // string | The contact's email address.
	channel := "channel_example" // string | The channel type for the subscription type. Currently, the only supported channel type is `EMAIL`.
	businessUnitId := int64(789) // int64 | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use `0`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionStatusAPI.GetCommunicationPreferencesV4StatusesSubscriberIdString(context.Background(), subscriberIdString).Channel(channel).BusinessUnitId(businessUnitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionStatusAPI.GetCommunicationPreferencesV4StatusesSubscriberIdString``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommunicationPreferencesV4StatusesSubscriberIdString`: ActionResponseWithResultsPublicStatus
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionStatusAPI.GetCommunicationPreferencesV4StatusesSubscriberIdString`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberIdString** | **string** | The contact&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationPreferencesV4StatusesSubscriberIdStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **string** | The channel type for the subscription type. Currently, the only supported channel type is &#x60;EMAIL&#x60;. | 
 **businessUnitId** | **int64** | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use &#x60;0&#x60;. | 

### Return type

[**ActionResponseWithResultsPublicStatus**](ActionResponseWithResultsPublicStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll

> ActionResponseWithResultsPublicWideStatus GetCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll(ctx, subscriberIdString).Channel(channel).BusinessUnitId(businessUnitId).Verbose(verbose).Execute()

Retrieve a contact's unsubscribed status



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
	subscriberIdString := "subscriberIdString_example" // string | The contact's email address.
	channel := "channel_example" // string | The channel type for the subscription type. Currently, the only supported channel type is `EMAIL`.
	businessUnitId := int64(789) // int64 | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use `0`. (optional)
	verbose := true // bool | Set to `true` to include the details of the updated subscription statuses in the response. Not including this parameter will result in an empty response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionStatusAPI.GetCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll(context.Background(), subscriberIdString).Channel(channel).BusinessUnitId(businessUnitId).Verbose(verbose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionStatusAPI.GetCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll`: ActionResponseWithResultsPublicWideStatus
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionStatusAPI.GetCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberIdString** | **string** | The contact&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **string** | The channel type for the subscription type. Currently, the only supported channel type is &#x60;EMAIL&#x60;. | 
 **businessUnitId** | **int64** | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use &#x60;0&#x60;. | 
 **verbose** | **bool** | Set to &#x60;true&#x60; to include the details of the updated subscription statuses in the response. Not including this parameter will result in an empty response. | [default to false]

### Return type

[**ActionResponseWithResultsPublicWideStatus**](ActionResponseWithResultsPublicWideStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV4StatusesBatchRead

> BatchResponsePublicStatusBulkResponse PostCommunicationPreferencesV4StatusesBatchRead(ctx).Channel(channel).BatchInputString(batchInputString).BusinessUnitId(businessUnitId).Execute()

Batch retrieve subscription statuses



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
	channel := "channel_example" // string | The channel type for the subscription type. Currently, the only supported channel type is `EMAIL`.
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | 
	businessUnitId := int64(789) // int64 | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use `0`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchRead(context.Background()).Channel(channel).BatchInputString(batchInputString).BusinessUnitId(businessUnitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommunicationPreferencesV4StatusesBatchRead`: BatchResponsePublicStatusBulkResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV4StatusesBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | The channel type for the subscription type. Currently, the only supported channel type is &#x60;EMAIL&#x60;. | 
 **batchInputString** | [**BatchInputString**](BatchInputString.md) |  | 
 **businessUnitId** | **int64** | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use &#x60;0&#x60;. | 

### Return type

[**BatchResponsePublicStatusBulkResponse**](BatchResponsePublicStatusBulkResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV4StatusesBatchUnsubscribeAll

> BatchResponsePublicBulkOptOutFromAllResponse PostCommunicationPreferencesV4StatusesBatchUnsubscribeAll(ctx).Channel(channel).BatchInputString(batchInputString).BusinessUnitId(businessUnitId).Verbose(verbose).Execute()

Batch unsubscribe contacts from all subscriptions



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
	channel := "channel_example" // string | The channel type for the subscription type. Currently, the only supported channel type is `EMAIL`.
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | 
	businessUnitId := int64(789) // int64 | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use `0`. (optional)
	verbose := true // bool | Set to `true` to include the details of the updated subscription statuses in the response. Not including this parameter will result in an empty response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchUnsubscribeAll(context.Background()).Channel(channel).BatchInputString(batchInputString).BusinessUnitId(businessUnitId).Verbose(verbose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchUnsubscribeAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommunicationPreferencesV4StatusesBatchUnsubscribeAll`: BatchResponsePublicBulkOptOutFromAllResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchUnsubscribeAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV4StatusesBatchUnsubscribeAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | The channel type for the subscription type. Currently, the only supported channel type is &#x60;EMAIL&#x60;. | 
 **batchInputString** | [**BatchInputString**](BatchInputString.md) |  | 
 **businessUnitId** | **int64** | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use &#x60;0&#x60;. | 
 **verbose** | **bool** | Set to &#x60;true&#x60; to include the details of the updated subscription statuses in the response. Not including this parameter will result in an empty response. | [default to false]

### Return type

[**BatchResponsePublicBulkOptOutFromAllResponse**](BatchResponsePublicBulkOptOutFromAllResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV4StatusesBatchUnsubscribeAllRead

> BatchResponsePublicWideStatusBulkResponse PostCommunicationPreferencesV4StatusesBatchUnsubscribeAllRead(ctx).Channel(channel).BatchInputString(batchInputString).BusinessUnitId(businessUnitId).Execute()

Batch retrieve contacts who have opted out of all communications



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
	channel := "channel_example" // string | The channel type for the subscription type. Currently, the only supported channel type is `EMAIL`.
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | 
	businessUnitId := int64(789) // int64 | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use `0`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchUnsubscribeAllRead(context.Background()).Channel(channel).BatchInputString(batchInputString).BusinessUnitId(businessUnitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchUnsubscribeAllRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommunicationPreferencesV4StatusesBatchUnsubscribeAllRead`: BatchResponsePublicWideStatusBulkResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchUnsubscribeAllRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV4StatusesBatchUnsubscribeAllReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | The channel type for the subscription type. Currently, the only supported channel type is &#x60;EMAIL&#x60;. | 
 **batchInputString** | [**BatchInputString**](BatchInputString.md) |  | 
 **businessUnitId** | **int64** | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use &#x60;0&#x60;. | 

### Return type

[**BatchResponsePublicWideStatusBulkResponse**](BatchResponsePublicWideStatusBulkResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV4StatusesBatchWrite

> BatchResponsePublicStatus PostCommunicationPreferencesV4StatusesBatchWrite(ctx).BatchInputPublicStatusRequest(batchInputPublicStatusRequest).Execute()

Batch update subscription status



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
	batchInputPublicStatusRequest := *openapiclient.NewBatchInputPublicStatusRequest([]openapiclient.PublicStatusRequest{*openapiclient.NewPublicStatusRequest("StatusState_example", "Channel_example", "SubscriberIdString_example", int32(123))}) // BatchInputPublicStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchWrite(context.Background()).BatchInputPublicStatusRequest(batchInputPublicStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchWrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommunicationPreferencesV4StatusesBatchWrite`: BatchResponsePublicStatus
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesBatchWrite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV4StatusesBatchWriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputPublicStatusRequest** | [**BatchInputPublicStatusRequest**](BatchInputPublicStatusRequest.md) |  | 

### Return type

[**BatchResponsePublicStatus**](BatchResponsePublicStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV4StatusesSubscriberIdString

> ActionResponseWithResultsPublicStatus PostCommunicationPreferencesV4StatusesSubscriberIdString(ctx, subscriberIdString).PartialPublicStatusRequest(partialPublicStatusRequest).Execute()

Update a contact's subscription status



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
	subscriberIdString := "subscriberIdString_example" // string | The contact's email address.
	partialPublicStatusRequest := *openapiclient.NewPartialPublicStatusRequest("StatusState_example", "Channel_example", int64(123)) // PartialPublicStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesSubscriberIdString(context.Background(), subscriberIdString).PartialPublicStatusRequest(partialPublicStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesSubscriberIdString``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommunicationPreferencesV4StatusesSubscriberIdString`: ActionResponseWithResultsPublicStatus
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesSubscriberIdString`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberIdString** | **string** | The contact&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV4StatusesSubscriberIdStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partialPublicStatusRequest** | [**PartialPublicStatusRequest**](PartialPublicStatusRequest.md) |  | 

### Return type

[**ActionResponseWithResultsPublicStatus**](ActionResponseWithResultsPublicStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll

> ActionResponseWithResultsPublicStatus PostCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll(ctx, subscriberIdString).Channel(channel).BusinessUnitId(businessUnitId).Verbose(verbose).Execute()

Unsubscribe a contact from all subscriptions



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
	subscriberIdString := "subscriberIdString_example" // string | The contact's email address.
	channel := "channel_example" // string | The channel type for the subscription type. Currently, the only supported channel type is `EMAIL`.
	businessUnitId := int64(789) // int64 | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use `0`. (optional)
	verbose := true // bool | Set to `true` to include the details of the updated subscription statuses in the response. Not including this parameter will result in an empty response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll(context.Background(), subscriberIdString).Channel(channel).BusinessUnitId(businessUnitId).Verbose(verbose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll`: ActionResponseWithResultsPublicStatus
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionStatusAPI.PostCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberIdString** | **string** | The contact&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV4StatusesSubscriberIdStringUnsubscribeAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **string** | The channel type for the subscription type. Currently, the only supported channel type is &#x60;EMAIL&#x60;. | 
 **businessUnitId** | **int64** | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use &#x60;0&#x60;. | 
 **verbose** | **bool** | Set to &#x60;true&#x60; to include the details of the updated subscription statuses in the response. Not including this parameter will result in an empty response. | [default to false]

### Return type

[**ActionResponseWithResultsPublicStatus**](ActionResponseWithResultsPublicStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

