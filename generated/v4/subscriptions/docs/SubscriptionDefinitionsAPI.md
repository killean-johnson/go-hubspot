# \SubscriptionDefinitionsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommunicationPreferencesV4Definitions**](SubscriptionDefinitionsAPI.md#GetCommunicationPreferencesV4Definitions) | **Get** /communication-preferences/v4/definitions | Retrieve all subscription status definitions



## GetCommunicationPreferencesV4Definitions

> ActionResponseWithResultsSubscriptionDefinition GetCommunicationPreferencesV4Definitions(ctx).BusinessUnitId(businessUnitId).IncludeTranslations(includeTranslations).Execute()

Retrieve all subscription status definitions



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
	businessUnitId := int64(789) // int64 | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use `0`. (optional)
	includeTranslations := true // bool | Set to `true` to return subscription translations associated with each definition. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionDefinitionsAPI.GetCommunicationPreferencesV4Definitions(context.Background()).BusinessUnitId(businessUnitId).IncludeTranslations(includeTranslations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionDefinitionsAPI.GetCommunicationPreferencesV4Definitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommunicationPreferencesV4Definitions`: ActionResponseWithResultsSubscriptionDefinition
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionDefinitionsAPI.GetCommunicationPreferencesV4Definitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationPreferencesV4DefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessUnitId** | **int64** | If you have the [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api), include this parameter to filter results by business unit ID. The default Account business unit will always use &#x60;0&#x60;. | 
 **includeTranslations** | **bool** | Set to &#x60;true&#x60; to return subscription translations associated with each definition. | 

### Return type

[**ActionResponseWithResultsSubscriptionDefinition**](ActionResponseWithResultsSubscriptionDefinition.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

