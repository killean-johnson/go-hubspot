# \PublicPropertyValidationsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3PropertyValidationsObjectTypeId**](PublicPropertyValidationsAPI.md#GetCrmV3PropertyValidationsObjectTypeId) | **Get** /crm/v3/property-validations/{objectTypeId} | Read all property validation rules for an object
[**GetCrmV3PropertyValidationsObjectTypeIdPropertyName**](PublicPropertyValidationsAPI.md#GetCrmV3PropertyValidationsObjectTypeIdPropertyName) | **Get** /crm/v3/property-validations/{objectTypeId}/{propertyName} | Read validation rules for a property



## GetCrmV3PropertyValidationsObjectTypeId

> CollectionResponsePublicPropertyValidationRuleMapNoPaging GetCrmV3PropertyValidationsObjectTypeId(ctx, objectTypeId).Execute()

Read all property validation rules for an object



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
	objectTypeId := "objectTypeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicPropertyValidationsAPI.GetCrmV3PropertyValidationsObjectTypeId(context.Background(), objectTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicPropertyValidationsAPI.GetCrmV3PropertyValidationsObjectTypeId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3PropertyValidationsObjectTypeId`: CollectionResponsePublicPropertyValidationRuleMapNoPaging
	fmt.Fprintf(os.Stdout, "Response from `PublicPropertyValidationsAPI.GetCrmV3PropertyValidationsObjectTypeId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PropertyValidationsObjectTypeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponsePublicPropertyValidationRuleMapNoPaging**](CollectionResponsePublicPropertyValidationRuleMapNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3PropertyValidationsObjectTypeIdPropertyName

> CollectionResponsePublicPropertyValidationRuleNoPaging GetCrmV3PropertyValidationsObjectTypeIdPropertyName(ctx, objectTypeId, propertyName).Execute()

Read validation rules for a property



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
	objectTypeId := "objectTypeId_example" // string | 
	propertyName := "propertyName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicPropertyValidationsAPI.GetCrmV3PropertyValidationsObjectTypeIdPropertyName(context.Background(), objectTypeId, propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicPropertyValidationsAPI.GetCrmV3PropertyValidationsObjectTypeIdPropertyName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3PropertyValidationsObjectTypeIdPropertyName`: CollectionResponsePublicPropertyValidationRuleNoPaging
	fmt.Fprintf(os.Stdout, "Response from `PublicPropertyValidationsAPI.GetCrmV3PropertyValidationsObjectTypeIdPropertyName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectTypeId** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PropertyValidationsObjectTypeIdPropertyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicPropertyValidationRuleNoPaging**](CollectionResponsePublicPropertyValidationRuleNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

