# \PublicTaxRateAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaxRatesV1TaxRates**](PublicTaxRateAPI.md#GetTaxRatesV1TaxRates) | **Get** /tax-rates/v1/tax-rates | 
[**GetTaxRatesV1TaxRatesTaxRateGroupId**](PublicTaxRateAPI.md#GetTaxRatesV1TaxRatesTaxRateGroupId) | **Get** /tax-rates/v1/tax-rates/{taxRateGroupId} | 



## GetTaxRatesV1TaxRates

> CollectionResponsePublicTaxRateGroupForwardPaging GetTaxRatesV1TaxRates(ctx).After(after).Limit(limit).Active(active).Execute()



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
	active := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicTaxRateAPI.GetTaxRatesV1TaxRates(context.Background()).After(after).Limit(limit).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicTaxRateAPI.GetTaxRatesV1TaxRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxRatesV1TaxRates`: CollectionResponsePublicTaxRateGroupForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `PublicTaxRateAPI.GetTaxRatesV1TaxRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxRatesV1TaxRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** |  | 
 **limit** | **int32** |  | 
 **active** | **bool** |  | 

### Return type

[**CollectionResponsePublicTaxRateGroupForwardPaging**](CollectionResponsePublicTaxRateGroupForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxRatesV1TaxRatesTaxRateGroupId

> PublicTaxRateGroup GetTaxRatesV1TaxRatesTaxRateGroupId(ctx, taxRateGroupId).Execute()



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
	taxRateGroupId := "taxRateGroupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicTaxRateAPI.GetTaxRatesV1TaxRatesTaxRateGroupId(context.Background(), taxRateGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicTaxRateAPI.GetTaxRatesV1TaxRatesTaxRateGroupId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxRatesV1TaxRatesTaxRateGroupId`: PublicTaxRateGroup
	fmt.Fprintf(os.Stdout, "Response from `PublicTaxRateAPI.GetTaxRatesV1TaxRatesTaxRateGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxRateGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxRatesV1TaxRatesTaxRateGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicTaxRateGroup**](PublicTaxRateGroup.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

