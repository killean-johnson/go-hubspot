# \RolesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingsV3UsersRolesGetAll**](RolesAPI.md#GetSettingsV3UsersRolesGetAll) | **Get** /settings/v3/users/roles | Retrieves the roles on an account



## GetSettingsV3UsersRolesGetAll

> CollectionResponsePublicPermissionSetNoPaging GetSettingsV3UsersRolesGetAll(ctx).Execute()

Retrieves the roles on an account



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
	resp, r, err := apiClient.RolesAPI.GetSettingsV3UsersRolesGetAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetSettingsV3UsersRolesGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV3UsersRolesGetAll`: CollectionResponsePublicPermissionSetNoPaging
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetSettingsV3UsersRolesGetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV3UsersRolesGetAllRequest struct via the builder pattern


### Return type

[**CollectionResponsePublicPermissionSetNoPaging**](CollectionResponsePublicPermissionSetNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

