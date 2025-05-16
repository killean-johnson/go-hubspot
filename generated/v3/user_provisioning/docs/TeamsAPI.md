# \TeamsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingsV3UsersTeamsGetAll**](TeamsAPI.md#GetSettingsV3UsersTeamsGetAll) | **Get** /settings/v3/users/teams | See details about this account&#39;s teams



## GetSettingsV3UsersTeamsGetAll

> CollectionResponsePublicTeamNoPaging GetSettingsV3UsersTeamsGetAll(ctx).Execute()

See details about this account's teams



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
	resp, r, err := apiClient.TeamsAPI.GetSettingsV3UsersTeamsGetAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetSettingsV3UsersTeamsGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV3UsersTeamsGetAll`: CollectionResponsePublicTeamNoPaging
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetSettingsV3UsersTeamsGetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV3UsersTeamsGetAllRequest struct via the builder pattern


### Return type

[**CollectionResponsePublicTeamNoPaging**](CollectionResponsePublicTeamNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

