# \ActivityAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountInfoV3ActivityAuditLogs**](ActivityAPI.md#GetAccountInfoV3ActivityAuditLogs) | **Get** /account-info/v3/activity/audit-logs | Retrieve audit logs
[**GetAccountInfoV3ActivityLogin**](ActivityAPI.md#GetAccountInfoV3ActivityLogin) | **Get** /account-info/v3/activity/login | Retrieve login activity
[**GetAccountInfoV3ActivitySecurity**](ActivityAPI.md#GetAccountInfoV3ActivitySecurity) | **Get** /account-info/v3/activity/security | Retrieve security history



## GetAccountInfoV3ActivityAuditLogs

> CollectionResponsePublicApiUserActionEventForwardPaging GetAccountInfoV3ActivityAuditLogs(ctx).After(after).Limit(limit).ActingUserId(actingUserId).OccurredAfter(occurredAfter).OccurredBefore(occurredBefore).Sort(sort).Execute()

Retrieve audit logs



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
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	actingUserId := []int32{int32(123)} // []int32 | The ID of a user, for retrieving user-specific logs. (optional)
	occurredAfter := time.Now() // time.Time | A timestamp, as a starting point for retrieving activity logs.  (optional)
	occurredBefore := time.Now() // time.Time | A timestamp, as an end point for retrieving activity logs.  (optional)
	sort := []string{"Inner_example"} // []string | Set to `occurredAt` to order results by the time of the event. By default, events are ordered from oldest to newest. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetAccountInfoV3ActivityAuditLogs(context.Background()).After(after).Limit(limit).ActingUserId(actingUserId).OccurredAfter(occurredAfter).OccurredBefore(occurredBefore).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetAccountInfoV3ActivityAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInfoV3ActivityAuditLogs`: CollectionResponsePublicApiUserActionEventForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetAccountInfoV3ActivityAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoV3ActivityAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **actingUserId** | **[]int32** | The ID of a user, for retrieving user-specific logs. | 
 **occurredAfter** | **time.Time** | A timestamp, as a starting point for retrieving activity logs.  | 
 **occurredBefore** | **time.Time** | A timestamp, as an end point for retrieving activity logs.  | 
 **sort** | **[]string** | Set to &#x60;occurredAt&#x60; to order results by the time of the event. By default, events are ordered from oldest to newest. | 

### Return type

[**CollectionResponsePublicApiUserActionEventForwardPaging**](CollectionResponsePublicApiUserActionEventForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInfoV3ActivityLogin

> CollectionResponsePublicLoginAuditForwardPaging GetAccountInfoV3ActivityLogin(ctx).After(after).Limit(limit).UserId(userId).Execute()

Retrieve login activity



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
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. Max value of limit is 200. (optional)
	userId := int32(56) // int32 | The ID of a user, for retrieving user-specific logs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetAccountInfoV3ActivityLogin(context.Background()).After(after).Limit(limit).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetAccountInfoV3ActivityLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInfoV3ActivityLogin`: CollectionResponsePublicLoginAuditForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetAccountInfoV3ActivityLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoV3ActivityLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. Max value of limit is 200. | 
 **userId** | **int32** | The ID of a user, for retrieving user-specific logs. | 

### Return type

[**CollectionResponsePublicLoginAuditForwardPaging**](CollectionResponsePublicLoginAuditForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInfoV3ActivitySecurity

> CollectionResponseHydratedCriticalActionForwardPaging GetAccountInfoV3ActivitySecurity(ctx).After(after).Limit(limit).UserId(userId).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Execute()

Retrieve security history



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
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. Max value of limit is 200. (optional)
	userId := int32(56) // int32 | The ID of a user, for retrieving user-specific logs. (optional)
	fromTimestamp := int64(789) // int64 | The start time, for retrieving logs within a specific timeframe. (optional)
	toTimestamp := int64(789) // int64 | The end time, for retrieving logs within a specific timeframe. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetAccountInfoV3ActivitySecurity(context.Background()).After(after).Limit(limit).UserId(userId).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetAccountInfoV3ActivitySecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInfoV3ActivitySecurity`: CollectionResponseHydratedCriticalActionForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetAccountInfoV3ActivitySecurity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoV3ActivitySecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. Max value of limit is 200. | 
 **userId** | **int32** | The ID of a user, for retrieving user-specific logs. | 
 **fromTimestamp** | **int64** | The start time, for retrieving logs within a specific timeframe. | 
 **toTimestamp** | **int64** | The end time, for retrieving logs within a specific timeframe. | 

### Return type

[**CollectionResponseHydratedCriticalActionForwardPaging**](CollectionResponseHydratedCriticalActionForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

