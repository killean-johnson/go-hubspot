# \StatisticsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3EmailsStatisticsHistogramGetHistogram**](StatisticsAPI.md#GetMarketingV3EmailsStatisticsHistogramGetHistogram) | **Get** /marketing/v3/emails/statistics/histogram | Get aggregated statistic intervals.
[**GetMarketingV3EmailsStatisticsListGetEmailsList**](StatisticsAPI.md#GetMarketingV3EmailsStatisticsListGetEmailsList) | **Get** /marketing/v3/emails/statistics/list | Get aggregated statistics.



## GetMarketingV3EmailsStatisticsHistogramGetHistogram

> CollectionResponseWithTotalEmailStatisticIntervalNoPaging GetMarketingV3EmailsStatisticsHistogramGetHistogram(ctx).Interval(interval).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).EmailIds(emailIds).Execute()

Get aggregated statistic intervals.



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
	interval := "interval_example" // string | The interval to aggregate statistics for. (optional)
	startTimestamp := "startTimestamp_example" // string | The start timestamp of the time span, in ISO8601 representation. (optional)
	endTimestamp := "endTimestamp_example" // string | The end timestamp of the time span, in ISO8601 representation. (optional)
	emailIds := []int64{int64(123)} // []int64 | Filter by email IDs. Only include statistics of emails with these IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.GetMarketingV3EmailsStatisticsHistogramGetHistogram(context.Background()).Interval(interval).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).EmailIds(emailIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.GetMarketingV3EmailsStatisticsHistogramGetHistogram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsStatisticsHistogramGetHistogram`: CollectionResponseWithTotalEmailStatisticIntervalNoPaging
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.GetMarketingV3EmailsStatisticsHistogramGetHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | **string** | The interval to aggregate statistics for. | 
 **startTimestamp** | **string** | The start timestamp of the time span, in ISO8601 representation. | 
 **endTimestamp** | **string** | The end timestamp of the time span, in ISO8601 representation. | 
 **emailIds** | **[]int64** | Filter by email IDs. Only include statistics of emails with these IDs. | 

### Return type

[**CollectionResponseWithTotalEmailStatisticIntervalNoPaging**](CollectionResponseWithTotalEmailStatisticIntervalNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3EmailsStatisticsListGetEmailsList

> AggregateEmailStatistics GetMarketingV3EmailsStatisticsListGetEmailsList(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).EmailIds(emailIds).Property(property).Execute()

Get aggregated statistics.



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
	startTimestamp := "startTimestamp_example" // string | The start timestamp of the time span, in ISO8601 representation. (optional)
	endTimestamp := "endTimestamp_example" // string | The end timestamp of the time span, in ISO8601 representation. (optional)
	emailIds := []int64{int64(123)} // []int64 | Filter by email IDs. Only include statistics of emails with these IDs. (optional)
	property := "property_example" // string | Specifies which email properties should be returned. All properties will be returned by default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.GetMarketingV3EmailsStatisticsListGetEmailsList(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).EmailIds(emailIds).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.GetMarketingV3EmailsStatisticsListGetEmailsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3EmailsStatisticsListGetEmailsList`: AggregateEmailStatistics
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.GetMarketingV3EmailsStatisticsListGetEmailsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3EmailsStatisticsListGetEmailsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **string** | The start timestamp of the time span, in ISO8601 representation. | 
 **endTimestamp** | **string** | The end timestamp of the time span, in ISO8601 representation. | 
 **emailIds** | **[]int64** | Filter by email IDs. Only include statistics of emails with these IDs. | 
 **property** | **string** | Specifies which email properties should be returned. All properties will be returned by default. | 

### Return type

[**AggregateEmailStatistics**](AggregateEmailStatistics.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

