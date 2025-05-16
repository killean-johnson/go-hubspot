# \MeetingsLinksAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSchedulerV3MeetingsMeetingLinks**](MeetingsLinksAPI.md#GetSchedulerV3MeetingsMeetingLinks) | **Get** /scheduler/v3/meetings/meeting-links | 
[**GetSchedulerV3MeetingsMeetingLinksBookAvailabilityPageSlug**](MeetingsLinksAPI.md#GetSchedulerV3MeetingsMeetingLinksBookAvailabilityPageSlug) | **Get** /scheduler/v3/meetings/meeting-links/book/availability-page/{slug} | 
[**GetSchedulerV3MeetingsMeetingLinksBookSlug**](MeetingsLinksAPI.md#GetSchedulerV3MeetingsMeetingLinksBookSlug) | **Get** /scheduler/v3/meetings/meeting-links/book/{slug} | 
[**PostSchedulerV3MeetingsMeetingLinksBook**](MeetingsLinksAPI.md#PostSchedulerV3MeetingsMeetingLinksBook) | **Post** /scheduler/v3/meetings/meeting-links/book | 



## GetSchedulerV3MeetingsMeetingLinks

> CollectionResponseWithTotalExternalLinkMetadataForwardPaging GetSchedulerV3MeetingsMeetingLinks(ctx).After(after).Limit(limit).Name(name).OrganizerUserId(organizerUserId).Type_(type_).Execute()



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
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	name := "name_example" // string |  (optional)
	organizerUserId := "organizerUserId_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingsLinksAPI.GetSchedulerV3MeetingsMeetingLinks(context.Background()).After(after).Limit(limit).Name(name).OrganizerUserId(organizerUserId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsLinksAPI.GetSchedulerV3MeetingsMeetingLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedulerV3MeetingsMeetingLinks`: CollectionResponseWithTotalExternalLinkMetadataForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `MeetingsLinksAPI.GetSchedulerV3MeetingsMeetingLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulerV3MeetingsMeetingLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** |  | 
 **limit** | **int32** |  | 
 **name** | **string** |  | 
 **organizerUserId** | **string** |  | 
 **type_** | **string** |  | 

### Return type

[**CollectionResponseWithTotalExternalLinkMetadataForwardPaging**](CollectionResponseWithTotalExternalLinkMetadataForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedulerV3MeetingsMeetingLinksBookAvailabilityPageSlug

> ExternalLinkAvailabilityAndBusyTimes GetSchedulerV3MeetingsMeetingLinksBookAvailabilityPageSlug(ctx, slug).Timezone(timezone).MonthOffset(monthOffset).Execute()



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
	slug := "slug_example" // string | 
	timezone := "timezone_example" // string |  (optional)
	monthOffset := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingsLinksAPI.GetSchedulerV3MeetingsMeetingLinksBookAvailabilityPageSlug(context.Background(), slug).Timezone(timezone).MonthOffset(monthOffset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsLinksAPI.GetSchedulerV3MeetingsMeetingLinksBookAvailabilityPageSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedulerV3MeetingsMeetingLinksBookAvailabilityPageSlug`: ExternalLinkAvailabilityAndBusyTimes
	fmt.Fprintf(os.Stdout, "Response from `MeetingsLinksAPI.GetSchedulerV3MeetingsMeetingLinksBookAvailabilityPageSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulerV3MeetingsMeetingLinksBookAvailabilityPageSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timezone** | **string** |  | 
 **monthOffset** | **int32** |  | 

### Return type

[**ExternalLinkAvailabilityAndBusyTimes**](ExternalLinkAvailabilityAndBusyTimes.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedulerV3MeetingsMeetingLinksBookSlug

> ExternalBookingInfo GetSchedulerV3MeetingsMeetingLinksBookSlug(ctx, slug).Timezone(timezone).Execute()



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
	slug := "slug_example" // string | 
	timezone := "timezone_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingsLinksAPI.GetSchedulerV3MeetingsMeetingLinksBookSlug(context.Background(), slug).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsLinksAPI.GetSchedulerV3MeetingsMeetingLinksBookSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedulerV3MeetingsMeetingLinksBookSlug`: ExternalBookingInfo
	fmt.Fprintf(os.Stdout, "Response from `MeetingsLinksAPI.GetSchedulerV3MeetingsMeetingLinksBookSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulerV3MeetingsMeetingLinksBookSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timezone** | **string** |  | 

### Return type

[**ExternalBookingInfo**](ExternalBookingInfo.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSchedulerV3MeetingsMeetingLinksBook

> ExternalMeetingBookingResponse PostSchedulerV3MeetingsMeetingLinksBook(ctx).ExternalMeetingBooking(externalMeetingBooking).Execute()



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
	externalMeetingBooking := *openapiclient.NewExternalMeetingBooking(int64(123), "FirstName_example", "LastName_example", []string{"LikelyAvailableUserIds_example"}, []openapiclient.ExternalLegalConsentResponse{*openapiclient.NewExternalLegalConsentResponse("CommunicationTypeId_example", false)}, time.Now(), []openapiclient.ExternalBookingFormField{*openapiclient.NewExternalBookingFormField("Name_example", "Value_example")}, "Slug_example", "Email_example") // ExternalMeetingBooking | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingsLinksAPI.PostSchedulerV3MeetingsMeetingLinksBook(context.Background()).ExternalMeetingBooking(externalMeetingBooking).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsLinksAPI.PostSchedulerV3MeetingsMeetingLinksBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSchedulerV3MeetingsMeetingLinksBook`: ExternalMeetingBookingResponse
	fmt.Fprintf(os.Stdout, "Response from `MeetingsLinksAPI.PostSchedulerV3MeetingsMeetingLinksBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSchedulerV3MeetingsMeetingLinksBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalMeetingBooking** | [**ExternalMeetingBooking**](ExternalMeetingBooking.md) |  | 

### Return type

[**ExternalMeetingBookingResponse**](ExternalMeetingBookingResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

