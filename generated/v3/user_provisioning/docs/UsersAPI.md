# \UsersAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSettingsV3UsersUserIdArchive**](UsersAPI.md#DeleteSettingsV3UsersUserIdArchive) | **Delete** /settings/v3/users/{userId} | Removes a user
[**GetSettingsV3UsersGetPage**](UsersAPI.md#GetSettingsV3UsersGetPage) | **Get** /settings/v3/users/ | Retrieves a list of users from an account
[**GetSettingsV3UsersUserIdGetById**](UsersAPI.md#GetSettingsV3UsersUserIdGetById) | **Get** /settings/v3/users/{userId} | Retrieves a user
[**PostSettingsV3UsersCreate**](UsersAPI.md#PostSettingsV3UsersCreate) | **Post** /settings/v3/users/ | Adds a user
[**PutSettingsV3UsersUserIdReplace**](UsersAPI.md#PutSettingsV3UsersUserIdReplace) | **Put** /settings/v3/users/{userId} | Modifies a user



## DeleteSettingsV3UsersUserIdArchive

> DeleteSettingsV3UsersUserIdArchive(ctx, userId).IdProperty(idProperty).Execute()

Removes a user



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
	userId := "userId_example" // string | Identifier of user to delete
	idProperty := "idProperty_example" // string | The name of a property with unique user values. Valid values are `USER_ID`(default) or `EMAIL` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DeleteSettingsV3UsersUserIdArchive(context.Background(), userId).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteSettingsV3UsersUserIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of user to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingsV3UsersUserIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idProperty** | **string** | The name of a property with unique user values. Valid values are &#x60;USER_ID&#x60;(default) or &#x60;EMAIL&#x60; | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsV3UsersGetPage

> CollectionResponsePublicUserForwardPaging GetSettingsV3UsersGetPage(ctx).Limit(limit).After(after).Execute()

Retrieves a list of users from an account



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
	limit := int32(56) // int32 | The number of users to retrieve (optional)
	after := "after_example" // string | Results will display maximum 100 users per page. Additional results will be on the next page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetSettingsV3UsersGetPage(context.Background()).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetSettingsV3UsersGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV3UsersGetPage`: CollectionResponsePublicUserForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetSettingsV3UsersGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV3UsersGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The number of users to retrieve | 
 **after** | **string** | Results will display maximum 100 users per page. Additional results will be on the next page. | 

### Return type

[**CollectionResponsePublicUserForwardPaging**](CollectionResponsePublicUserForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsV3UsersUserIdGetById

> PublicUser GetSettingsV3UsersUserIdGetById(ctx, userId).IdProperty(idProperty).Execute()

Retrieves a user



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
	userId := "userId_example" // string | Identifier of user to retrieve
	idProperty := "idProperty_example" // string | The name of a property with unique user values. Valid values are `USER_ID`(default) or `EMAIL` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetSettingsV3UsersUserIdGetById(context.Background(), userId).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetSettingsV3UsersUserIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV3UsersUserIdGetById`: PublicUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetSettingsV3UsersUserIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of user to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV3UsersUserIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idProperty** | **string** | The name of a property with unique user values. Valid values are &#x60;USER_ID&#x60;(default) or &#x60;EMAIL&#x60; | 

### Return type

[**PublicUser**](PublicUser.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsV3UsersCreate

> PublicUser PostSettingsV3UsersCreate(ctx).UserProvisionRequest(userProvisionRequest).Execute()

Adds a user



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
	userProvisionRequest := *openapiclient.NewUserProvisionRequest("newUser@email.com") // UserProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.PostSettingsV3UsersCreate(context.Background()).UserProvisionRequest(userProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostSettingsV3UsersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSettingsV3UsersCreate`: PublicUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PostSettingsV3UsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsV3UsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userProvisionRequest** | [**UserProvisionRequest**](UserProvisionRequest.md) |  | 

### Return type

[**PublicUser**](PublicUser.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSettingsV3UsersUserIdReplace

> PublicUser PutSettingsV3UsersUserIdReplace(ctx, userId).PublicUserUpdate(publicUserUpdate).IdProperty(idProperty).Execute()

Modifies a user



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
	userId := "userId_example" // string | Identifier of user to retrieve
	publicUserUpdate := *openapiclient.NewPublicUserUpdate() // PublicUserUpdate | 
	idProperty := "idProperty_example" // string | The name of a property with unique user values. Valid values are `USER_ID`(default) or `EMAIL` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.PutSettingsV3UsersUserIdReplace(context.Background(), userId).PublicUserUpdate(publicUserUpdate).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PutSettingsV3UsersUserIdReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSettingsV3UsersUserIdReplace`: PublicUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PutSettingsV3UsersUserIdReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of user to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSettingsV3UsersUserIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicUserUpdate** | [**PublicUserUpdate**](PublicUserUpdate.md) |  | 
 **idProperty** | **string** | The name of a property with unique user values. Valid values are &#x60;USER_ID&#x60;(default) or &#x60;EMAIL&#x60; | 

### Return type

[**PublicUser**](PublicUser.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

