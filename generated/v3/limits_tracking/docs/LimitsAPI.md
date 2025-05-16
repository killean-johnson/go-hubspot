# \LimitsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3LimitsAssociationsLabels**](LimitsAPI.md#GetCrmV3LimitsAssociationsLabels) | **Get** /crm/v3/limits/associations/labels | Read association label limits
[**GetCrmV3LimitsAssociationsRecordsFrom**](LimitsAPI.md#GetCrmV3LimitsAssociationsRecordsFrom) | **Get** /crm/v3/limits/associations/records/from | Read record association limits
[**GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdTo**](LimitsAPI.md#GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdTo) | **Get** /crm/v3/limits/associations/records/{fromObjectTypeId}/to | Read record association limits from an object
[**GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToObjectTypeId**](LimitsAPI.md#GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToObjectTypeId) | **Get** /crm/v3/limits/associations/records/{fromObjectTypeId}/{toObjectTypeId} | Read record association limits between two objects
[**GetCrmV3LimitsCalculatedProperties**](LimitsAPI.md#GetCrmV3LimitsCalculatedProperties) | **Get** /crm/v3/limits/calculated-properties | Read calculation property limits
[**GetCrmV3LimitsCustomObjectTypes**](LimitsAPI.md#GetCrmV3LimitsCustomObjectTypes) | **Get** /crm/v3/limits/custom-object-types | Read custom object limits
[**GetCrmV3LimitsCustomProperties**](LimitsAPI.md#GetCrmV3LimitsCustomProperties) | **Get** /crm/v3/limits/custom-properties | Read custom property limits
[**GetCrmV3LimitsPipelines**](LimitsAPI.md#GetCrmV3LimitsPipelines) | **Get** /crm/v3/limits/pipelines | Read pipeline limits
[**GetCrmV3LimitsRecords**](LimitsAPI.md#GetCrmV3LimitsRecords) | **Get** /crm/v3/limits/records | Read record limits



## GetCrmV3LimitsAssociationsLabels

> CollectionResponseAssociationLabelLimitResponseNoPaging GetCrmV3LimitsAssociationsLabels(ctx).FromObjectTypeId(fromObjectTypeId).ToObjectTypeId(toObjectTypeId).Execute()

Read association label limits



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
	fromObjectTypeId := "fromObjectTypeId_example" // string |  (optional)
	toObjectTypeId := "toObjectTypeId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitsAPI.GetCrmV3LimitsAssociationsLabels(context.Background()).FromObjectTypeId(fromObjectTypeId).ToObjectTypeId(toObjectTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.GetCrmV3LimitsAssociationsLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3LimitsAssociationsLabels`: CollectionResponseAssociationLabelLimitResponseNoPaging
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.GetCrmV3LimitsAssociationsLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3LimitsAssociationsLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromObjectTypeId** | **string** |  | 
 **toObjectTypeId** | **string** |  | 

### Return type

[**CollectionResponseAssociationLabelLimitResponseNoPaging**](CollectionResponseAssociationLabelLimitResponseNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3LimitsAssociationsRecordsFrom

> CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging GetCrmV3LimitsAssociationsRecordsFrom(ctx).Execute()

Read record association limits



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
	resp, r, err := apiClient.LimitsAPI.GetCrmV3LimitsAssociationsRecordsFrom(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.GetCrmV3LimitsAssociationsRecordsFrom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3LimitsAssociationsRecordsFrom`: CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.GetCrmV3LimitsAssociationsRecordsFrom`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3LimitsAssociationsRecordsFromRequest struct via the builder pattern


### Return type

[**CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging**](CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdTo

> CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdTo(ctx, fromObjectTypeId).Execute()

Read record association limits from an object



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
	fromObjectTypeId := "fromObjectTypeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitsAPI.GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdTo(context.Background(), fromObjectTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdTo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdTo`: CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging**](CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToObjectTypeId

> AssociationRecordLimitResponse GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToObjectTypeId(ctx, fromObjectTypeId, toObjectTypeId).Execute()

Read record association limits between two objects



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
	fromObjectTypeId := "fromObjectTypeId_example" // string | 
	toObjectTypeId := "toObjectTypeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitsAPI.GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToObjectTypeId(context.Background(), fromObjectTypeId, toObjectTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToObjectTypeId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToObjectTypeId`: AssociationRecordLimitResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.GetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToObjectTypeId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectTypeId** | **string** |  | 
**toObjectTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3LimitsAssociationsRecordsFromObjectTypeIdToObjectTypeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AssociationRecordLimitResponse**](AssociationRecordLimitResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3LimitsCalculatedProperties

> CalculatedPropertyLimitResponse GetCrmV3LimitsCalculatedProperties(ctx).Execute()

Read calculation property limits



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
	resp, r, err := apiClient.LimitsAPI.GetCrmV3LimitsCalculatedProperties(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.GetCrmV3LimitsCalculatedProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3LimitsCalculatedProperties`: CalculatedPropertyLimitResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.GetCrmV3LimitsCalculatedProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3LimitsCalculatedPropertiesRequest struct via the builder pattern


### Return type

[**CalculatedPropertyLimitResponse**](CalculatedPropertyLimitResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3LimitsCustomObjectTypes

> CustomObjectLimitResponse GetCrmV3LimitsCustomObjectTypes(ctx).Execute()

Read custom object limits



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
	resp, r, err := apiClient.LimitsAPI.GetCrmV3LimitsCustomObjectTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.GetCrmV3LimitsCustomObjectTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3LimitsCustomObjectTypes`: CustomObjectLimitResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.GetCrmV3LimitsCustomObjectTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3LimitsCustomObjectTypesRequest struct via the builder pattern


### Return type

[**CustomObjectLimitResponse**](CustomObjectLimitResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3LimitsCustomProperties

> CustomPropertyLimitResponse GetCrmV3LimitsCustomProperties(ctx).Execute()

Read custom property limits



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
	resp, r, err := apiClient.LimitsAPI.GetCrmV3LimitsCustomProperties(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.GetCrmV3LimitsCustomProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3LimitsCustomProperties`: CustomPropertyLimitResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.GetCrmV3LimitsCustomProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3LimitsCustomPropertiesRequest struct via the builder pattern


### Return type

[**CustomPropertyLimitResponse**](CustomPropertyLimitResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3LimitsPipelines

> PipelineLimitResponse GetCrmV3LimitsPipelines(ctx).Execute()

Read pipeline limits



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
	resp, r, err := apiClient.LimitsAPI.GetCrmV3LimitsPipelines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.GetCrmV3LimitsPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3LimitsPipelines`: PipelineLimitResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.GetCrmV3LimitsPipelines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3LimitsPipelinesRequest struct via the builder pattern


### Return type

[**PipelineLimitResponse**](PipelineLimitResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3LimitsRecords

> RecordLimitResponse GetCrmV3LimitsRecords(ctx).Execute()

Read record limits



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
	resp, r, err := apiClient.LimitsAPI.GetCrmV3LimitsRecords(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.GetCrmV3LimitsRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3LimitsRecords`: RecordLimitResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.GetCrmV3LimitsRecords`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3LimitsRecordsRequest struct via the builder pattern


### Return type

[**RecordLimitResponse**](RecordLimitResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

