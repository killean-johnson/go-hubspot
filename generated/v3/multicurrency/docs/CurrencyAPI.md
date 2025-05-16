# \CurrencyAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingsV3CurrenciesCodes**](CurrencyAPI.md#GetSettingsV3CurrenciesCodes) | **Get** /settings/v3/currencies/codes | 
[**GetSettingsV3CurrenciesCompanyCurrency**](CurrencyAPI.md#GetSettingsV3CurrenciesCompanyCurrency) | **Get** /settings/v3/currencies/company-currency | 
[**GetSettingsV3CurrenciesExchangeRates**](CurrencyAPI.md#GetSettingsV3CurrenciesExchangeRates) | **Get** /settings/v3/currencies/exchange-rates | 
[**GetSettingsV3CurrenciesExchangeRatesCurrent**](CurrencyAPI.md#GetSettingsV3CurrenciesExchangeRatesCurrent) | **Get** /settings/v3/currencies/exchange-rates/current | 
[**GetSettingsV3CurrenciesExchangeRatesExchangeRateId**](CurrencyAPI.md#GetSettingsV3CurrenciesExchangeRatesExchangeRateId) | **Get** /settings/v3/currencies/exchange-rates/{exchangeRateId} | 
[**PatchSettingsV3CurrenciesExchangeRatesExchangeRateId**](CurrencyAPI.md#PatchSettingsV3CurrenciesExchangeRatesExchangeRateId) | **Patch** /settings/v3/currencies/exchange-rates/{exchangeRateId} | 
[**PostSettingsV3CurrenciesExchangeRates**](CurrencyAPI.md#PostSettingsV3CurrenciesExchangeRates) | **Post** /settings/v3/currencies/exchange-rates | 
[**PostSettingsV3CurrenciesExchangeRatesBatchCreate**](CurrencyAPI.md#PostSettingsV3CurrenciesExchangeRatesBatchCreate) | **Post** /settings/v3/currencies/exchange-rates/batch/create | 
[**PostSettingsV3CurrenciesExchangeRatesBatchRead**](CurrencyAPI.md#PostSettingsV3CurrenciesExchangeRatesBatchRead) | **Post** /settings/v3/currencies/exchange-rates/batch/read | 
[**PostSettingsV3CurrenciesExchangeRatesBatchUpdate**](CurrencyAPI.md#PostSettingsV3CurrenciesExchangeRatesBatchUpdate) | **Post** /settings/v3/currencies/exchange-rates/batch/update | 
[**PostSettingsV3CurrenciesExchangeRatesUpdateVisibility**](CurrencyAPI.md#PostSettingsV3CurrenciesExchangeRatesUpdateVisibility) | **Post** /settings/v3/currencies/exchange-rates/update-visibility | 
[**PutSettingsV3CurrenciesCompanyCurrency**](CurrencyAPI.md#PutSettingsV3CurrenciesCompanyCurrency) | **Put** /settings/v3/currencies/company-currency | 



## GetSettingsV3CurrenciesCodes

> CollectionResponseCurrencyCodeInfoNoPaging GetSettingsV3CurrenciesCodes(ctx).Execute()



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
	resp, r, err := apiClient.CurrencyAPI.GetSettingsV3CurrenciesCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.GetSettingsV3CurrenciesCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV3CurrenciesCodes`: CollectionResponseCurrencyCodeInfoNoPaging
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.GetSettingsV3CurrenciesCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV3CurrenciesCodesRequest struct via the builder pattern


### Return type

[**CollectionResponseCurrencyCodeInfoNoPaging**](CollectionResponseCurrencyCodeInfoNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsV3CurrenciesCompanyCurrency

> CompanyCurrency GetSettingsV3CurrenciesCompanyCurrency(ctx).Execute()



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
	resp, r, err := apiClient.CurrencyAPI.GetSettingsV3CurrenciesCompanyCurrency(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.GetSettingsV3CurrenciesCompanyCurrency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV3CurrenciesCompanyCurrency`: CompanyCurrency
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.GetSettingsV3CurrenciesCompanyCurrency`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV3CurrenciesCompanyCurrencyRequest struct via the builder pattern


### Return type

[**CompanyCurrency**](CompanyCurrency.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsV3CurrenciesExchangeRates

> CollectionResponseExchangeRateForwardPaging GetSettingsV3CurrenciesExchangeRates(ctx).Limit(limit).After(after).FromCurrencyCode(fromCurrencyCode).ToCurrencyCode(toCurrencyCode).Execute()



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
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 100)
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	fromCurrencyCode := "fromCurrencyCode_example" // string |  (optional)
	toCurrencyCode := "toCurrencyCode_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.GetSettingsV3CurrenciesExchangeRates(context.Background()).Limit(limit).After(after).FromCurrencyCode(fromCurrencyCode).ToCurrencyCode(toCurrencyCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.GetSettingsV3CurrenciesExchangeRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV3CurrenciesExchangeRates`: CollectionResponseExchangeRateForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.GetSettingsV3CurrenciesExchangeRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV3CurrenciesExchangeRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to display per page. | [default to 100]
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **fromCurrencyCode** | **string** |  | 
 **toCurrencyCode** | **string** |  | 

### Return type

[**CollectionResponseExchangeRateForwardPaging**](CollectionResponseExchangeRateForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsV3CurrenciesExchangeRatesCurrent

> CollectionResponseExchangeRateNoPaging GetSettingsV3CurrenciesExchangeRatesCurrent(ctx).Execute()



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
	resp, r, err := apiClient.CurrencyAPI.GetSettingsV3CurrenciesExchangeRatesCurrent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.GetSettingsV3CurrenciesExchangeRatesCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV3CurrenciesExchangeRatesCurrent`: CollectionResponseExchangeRateNoPaging
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.GetSettingsV3CurrenciesExchangeRatesCurrent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV3CurrenciesExchangeRatesCurrentRequest struct via the builder pattern


### Return type

[**CollectionResponseExchangeRateNoPaging**](CollectionResponseExchangeRateNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsV3CurrenciesExchangeRatesExchangeRateId

> ExchangeRate GetSettingsV3CurrenciesExchangeRatesExchangeRateId(ctx, exchangeRateId).Execute()



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
	exchangeRateId := "exchangeRateId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.GetSettingsV3CurrenciesExchangeRatesExchangeRateId(context.Background(), exchangeRateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.GetSettingsV3CurrenciesExchangeRatesExchangeRateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV3CurrenciesExchangeRatesExchangeRateId`: ExchangeRate
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.GetSettingsV3CurrenciesExchangeRatesExchangeRateId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeRateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV3CurrenciesExchangeRatesExchangeRateIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExchangeRate**](ExchangeRate.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSettingsV3CurrenciesExchangeRatesExchangeRateId

> ExchangeRate PatchSettingsV3CurrenciesExchangeRatesExchangeRateId(ctx, exchangeRateId).ExchangeRateMultiplier(exchangeRateMultiplier).Execute()



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
	exchangeRateId := "exchangeRateId_example" // string | 
	exchangeRateMultiplier := *openapiclient.NewExchangeRateMultiplier(float32(123)) // ExchangeRateMultiplier | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.PatchSettingsV3CurrenciesExchangeRatesExchangeRateId(context.Background(), exchangeRateId).ExchangeRateMultiplier(exchangeRateMultiplier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.PatchSettingsV3CurrenciesExchangeRatesExchangeRateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSettingsV3CurrenciesExchangeRatesExchangeRateId`: ExchangeRate
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.PatchSettingsV3CurrenciesExchangeRatesExchangeRateId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeRateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSettingsV3CurrenciesExchangeRatesExchangeRateIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exchangeRateMultiplier** | [**ExchangeRateMultiplier**](ExchangeRateMultiplier.md) |  | 

### Return type

[**ExchangeRate**](ExchangeRate.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsV3CurrenciesExchangeRates

> ExchangeRate PostSettingsV3CurrenciesExchangeRates(ctx).ExchangeRateCreateRequest(exchangeRateCreateRequest).Execute()



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
	exchangeRateCreateRequest := *openapiclient.NewExchangeRateCreateRequest(float32(123), "FromCurrencyCode_example") // ExchangeRateCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.PostSettingsV3CurrenciesExchangeRates(context.Background()).ExchangeRateCreateRequest(exchangeRateCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.PostSettingsV3CurrenciesExchangeRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSettingsV3CurrenciesExchangeRates`: ExchangeRate
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.PostSettingsV3CurrenciesExchangeRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsV3CurrenciesExchangeRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchangeRateCreateRequest** | [**ExchangeRateCreateRequest**](ExchangeRateCreateRequest.md) |  | 

### Return type

[**ExchangeRate**](ExchangeRate.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsV3CurrenciesExchangeRatesBatchCreate

> BatchResponseExchangeRate PostSettingsV3CurrenciesExchangeRatesBatchCreate(ctx).BatchInputExchangeRateCreateRequest(batchInputExchangeRateCreateRequest).Execute()



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
	batchInputExchangeRateCreateRequest := *openapiclient.NewBatchInputExchangeRateCreateRequest([]openapiclient.ExchangeRateCreateRequest{*openapiclient.NewExchangeRateCreateRequest(float32(123), "FromCurrencyCode_example")}) // BatchInputExchangeRateCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesBatchCreate(context.Background()).BatchInputExchangeRateCreateRequest(batchInputExchangeRateCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSettingsV3CurrenciesExchangeRatesBatchCreate`: BatchResponseExchangeRate
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesBatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsV3CurrenciesExchangeRatesBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputExchangeRateCreateRequest** | [**BatchInputExchangeRateCreateRequest**](BatchInputExchangeRateCreateRequest.md) |  | 

### Return type

[**BatchResponseExchangeRate**](BatchResponseExchangeRate.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsV3CurrenciesExchangeRatesBatchRead

> BatchResponseExchangeRate PostSettingsV3CurrenciesExchangeRatesBatchRead(ctx).BatchInputPublicObjectId(batchInputPublicObjectId).Execute()



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
	batchInputPublicObjectId := *openapiclient.NewBatchInputPublicObjectId([]openapiclient.PublicObjectId{*openapiclient.NewPublicObjectId("Id_example")}) // BatchInputPublicObjectId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesBatchRead(context.Background()).BatchInputPublicObjectId(batchInputPublicObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSettingsV3CurrenciesExchangeRatesBatchRead`: BatchResponseExchangeRate
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsV3CurrenciesExchangeRatesBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputPublicObjectId** | [**BatchInputPublicObjectId**](BatchInputPublicObjectId.md) |  | 

### Return type

[**BatchResponseExchangeRate**](BatchResponseExchangeRate.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsV3CurrenciesExchangeRatesBatchUpdate

> BatchResponseExchangeRate PostSettingsV3CurrenciesExchangeRatesBatchUpdate(ctx).BatchInputExchangeRateUpdateRequest(batchInputExchangeRateUpdateRequest).Execute()



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
	batchInputExchangeRateUpdateRequest := *openapiclient.NewBatchInputExchangeRateUpdateRequest([]openapiclient.ExchangeRateUpdateRequest{*openapiclient.NewExchangeRateUpdateRequest("Id_example", float32(123))}) // BatchInputExchangeRateUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesBatchUpdate(context.Background()).BatchInputExchangeRateUpdateRequest(batchInputExchangeRateUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSettingsV3CurrenciesExchangeRatesBatchUpdate`: BatchResponseExchangeRate
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesBatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsV3CurrenciesExchangeRatesBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputExchangeRateUpdateRequest** | [**BatchInputExchangeRateUpdateRequest**](BatchInputExchangeRateUpdateRequest.md) |  | 

### Return type

[**BatchResponseExchangeRate**](BatchResponseExchangeRate.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsV3CurrenciesExchangeRatesUpdateVisibility

> PostSettingsV3CurrenciesExchangeRatesUpdateVisibility(ctx).CurrencyPairUpdate(currencyPairUpdate).Execute()



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
	currencyPairUpdate := *openapiclient.NewCurrencyPairUpdate("ToCurrencyCode_example", false, "FromCurrencyCode_example") // CurrencyPairUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesUpdateVisibility(context.Background()).CurrencyPairUpdate(currencyPairUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.PostSettingsV3CurrenciesExchangeRatesUpdateVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsV3CurrenciesExchangeRatesUpdateVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyPairUpdate** | [**CurrencyPairUpdate**](CurrencyPairUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSettingsV3CurrenciesCompanyCurrency

> CompanyCurrency PutSettingsV3CurrenciesCompanyCurrency(ctx).CompanyCurrencyUpdateRequest(companyCurrencyUpdateRequest).Execute()



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
	companyCurrencyUpdateRequest := *openapiclient.NewCompanyCurrencyUpdateRequest("CurrencyCode_example") // CompanyCurrencyUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.PutSettingsV3CurrenciesCompanyCurrency(context.Background()).CompanyCurrencyUpdateRequest(companyCurrencyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.PutSettingsV3CurrenciesCompanyCurrency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSettingsV3CurrenciesCompanyCurrency`: CompanyCurrency
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.PutSettingsV3CurrenciesCompanyCurrency`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutSettingsV3CurrenciesCompanyCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyCurrencyUpdateRequest** | [**CompanyCurrencyUpdateRequest**](CompanyCurrencyUpdateRequest.md) |  | 

### Return type

[**CompanyCurrency**](CompanyCurrency.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

