# \EnablementAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ObjectLibraryEnablement**](EnablementAPI.md#GetCrmV3ObjectLibraryEnablement) | **Get** /crm/v3/object-library/enablement | Read All
[**GetCrmV3ObjectLibraryEnablementObjectTypeId**](EnablementAPI.md#GetCrmV3ObjectLibraryEnablementObjectTypeId) | **Get** /crm/v3/object-library/enablement/{objectTypeId} | Read



## GetCrmV3ObjectLibraryEnablement

> PortalObjectTypeEnablementPublicResponse GetCrmV3ObjectLibraryEnablement(ctx).Execute()

Read All



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
	resp, r, err := apiClient.EnablementAPI.GetCrmV3ObjectLibraryEnablement(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnablementAPI.GetCrmV3ObjectLibraryEnablement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ObjectLibraryEnablement`: PortalObjectTypeEnablementPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `EnablementAPI.GetCrmV3ObjectLibraryEnablement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ObjectLibraryEnablementRequest struct via the builder pattern


### Return type

[**PortalObjectTypeEnablementPublicResponse**](PortalObjectTypeEnablementPublicResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ObjectLibraryEnablementObjectTypeId

> ObjectTypeEnablementPublicResponse GetCrmV3ObjectLibraryEnablementObjectTypeId(ctx, objectTypeId).Execute()

Read



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
	resp, r, err := apiClient.EnablementAPI.GetCrmV3ObjectLibraryEnablementObjectTypeId(context.Background(), objectTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnablementAPI.GetCrmV3ObjectLibraryEnablementObjectTypeId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ObjectLibraryEnablementObjectTypeId`: ObjectTypeEnablementPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `EnablementAPI.GetCrmV3ObjectLibraryEnablementObjectTypeId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ObjectLibraryEnablementObjectTypeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectTypeEnablementPublicResponse**](ObjectTypeEnablementPublicResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

