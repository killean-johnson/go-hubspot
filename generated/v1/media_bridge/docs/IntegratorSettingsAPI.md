# \IntegratorSettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMediaBridgeV1AppIdSettingsOembedDomains**](IntegratorSettingsAPI.md#DeleteMediaBridgeV1AppIdSettingsOembedDomains) | **Delete** /media-bridge/v1/{appId}/settings/oembed-domains | 
[**GetMediaBridgeV1AppIdSettingsEventVisibility**](IntegratorSettingsAPI.md#GetMediaBridgeV1AppIdSettingsEventVisibility) | **Get** /media-bridge/v1/{appId}/settings/event-visibility | 
[**GetMediaBridgeV1AppIdSettingsObjectDefinitionsMediaType**](IntegratorSettingsAPI.md#GetMediaBridgeV1AppIdSettingsObjectDefinitionsMediaType) | **Get** /media-bridge/v1/{appId}/settings/object-definitions/{mediaType} | 
[**GetMediaBridgeV1AppIdSettingsOembedDomains**](IntegratorSettingsAPI.md#GetMediaBridgeV1AppIdSettingsOembedDomains) | **Get** /media-bridge/v1/{appId}/settings/oembed-domains | 
[**GetMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId**](IntegratorSettingsAPI.md#GetMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId) | **Get** /media-bridge/v1/{appId}/settings/oembed-domains/{oEmbedDomainId} | 
[**PatchMediaBridgeV1AppIdSettingsEventVisibility**](IntegratorSettingsAPI.md#PatchMediaBridgeV1AppIdSettingsEventVisibility) | **Patch** /media-bridge/v1/{appId}/settings/event-visibility | 
[**PatchMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId**](IntegratorSettingsAPI.md#PatchMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId) | **Patch** /media-bridge/v1/{appId}/settings/oembed-domains/{oEmbedDomainId} | 
[**PostMediaBridgeV1AppIdSettingsObjectDefinitions**](IntegratorSettingsAPI.md#PostMediaBridgeV1AppIdSettingsObjectDefinitions) | **Post** /media-bridge/v1/{appId}/settings/object-definitions | 
[**PostMediaBridgeV1AppIdSettingsOembedDomains**](IntegratorSettingsAPI.md#PostMediaBridgeV1AppIdSettingsOembedDomains) | **Post** /media-bridge/v1/{appId}/settings/oembed-domains | 
[**PostMediaBridgeV1AppIdSettingsRegister**](IntegratorSettingsAPI.md#PostMediaBridgeV1AppIdSettingsRegister) | **Post** /media-bridge/v1/{appId}/settings/register | 
[**PutMediaBridgeV1AppIdSettings**](IntegratorSettingsAPI.md#PutMediaBridgeV1AppIdSettings) | **Put** /media-bridge/v1/{appId}/settings | 



## DeleteMediaBridgeV1AppIdSettingsOembedDomains

> DeleteMediaBridgeV1AppIdSettingsOembedDomains(ctx).Id(id).DomainPortalId(domainPortalId).Execute()



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
	id := int64(789) // int64 |  (optional)
	domainPortalId := int32(56) // int32 |  (optional) (default to -1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegratorSettingsAPI.DeleteMediaBridgeV1AppIdSettingsOembedDomains(context.Background()).Id(id).DomainPortalId(domainPortalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.DeleteMediaBridgeV1AppIdSettingsOembedDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaBridgeV1AppIdSettingsOembedDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** |  | 
 **domainPortalId** | **int32** |  | [default to -1]

### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaBridgeV1AppIdSettingsEventVisibility

> EventVisibilityResponse GetMediaBridgeV1AppIdSettingsEventVisibility(ctx).Execute()



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
	resp, r, err := apiClient.IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsEventVisibility(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsEventVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdSettingsEventVisibility`: EventVisibilityResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsEventVisibility`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdSettingsEventVisibilityRequest struct via the builder pattern


### Return type

[**EventVisibilityResponse**](EventVisibilityResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaBridgeV1AppIdSettingsObjectDefinitionsMediaType

> ObjectDefinitionResponse GetMediaBridgeV1AppIdSettingsObjectDefinitionsMediaType(ctx, mediaType).IncludeFullDefinition(includeFullDefinition).Execute()



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
	mediaType := "mediaType_example" // string | 
	includeFullDefinition := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsObjectDefinitionsMediaType(context.Background(), mediaType).IncludeFullDefinition(includeFullDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsObjectDefinitionsMediaType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdSettingsObjectDefinitionsMediaType`: ObjectDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsObjectDefinitionsMediaType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdSettingsObjectDefinitionsMediaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeFullDefinition** | **bool** |  | 

### Return type

[**ObjectDefinitionResponse**](ObjectDefinitionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaBridgeV1AppIdSettingsOembedDomains

> OEmbedDomainsCollectionResponse GetMediaBridgeV1AppIdSettingsOembedDomains(ctx).DomainPortalId(domainPortalId).Execute()



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
	domainPortalId := int32(56) // int32 |  (optional) (default to -1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsOembedDomains(context.Background()).DomainPortalId(domainPortalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsOembedDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdSettingsOembedDomains`: OEmbedDomainsCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsOembedDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdSettingsOembedDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainPortalId** | **int32** |  | [default to -1]

### Return type

[**OEmbedDomainsCollectionResponse**](OEmbedDomainsCollectionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId

> IntegratorOEmbedDomainModel GetMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId(ctx, oEmbedDomainId).Execute()



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
	oEmbedDomainId := "oEmbedDomainId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId(context.Background(), oEmbedDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId`: IntegratorOEmbedDomainModel
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.GetMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oEmbedDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegratorOEmbedDomainModel**](IntegratorOEmbedDomainModel.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMediaBridgeV1AppIdSettingsEventVisibility

> EventVisibilityChange PatchMediaBridgeV1AppIdSettingsEventVisibility(ctx).EventVisibilityChange(eventVisibilityChange).Execute()



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
	eventVisibilityChange := *openapiclient.NewEventVisibilityChange("EventType_example", int64(123)) // EventVisibilityChange | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorSettingsAPI.PatchMediaBridgeV1AppIdSettingsEventVisibility(context.Background()).EventVisibilityChange(eventVisibilityChange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.PatchMediaBridgeV1AppIdSettingsEventVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMediaBridgeV1AppIdSettingsEventVisibility`: EventVisibilityChange
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.PatchMediaBridgeV1AppIdSettingsEventVisibility`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchMediaBridgeV1AppIdSettingsEventVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventVisibilityChange** | [**EventVisibilityChange**](EventVisibilityChange.md) |  | 

### Return type

[**EventVisibilityChange**](EventVisibilityChange.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId

> IntegratorOEmbedDomainModel PatchMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId(ctx, oEmbedDomainId).IntegratorOEmbedDomainRequest(integratorOEmbedDomainRequest).Execute()



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
	oEmbedDomainId := "oEmbedDomainId_example" // string | 
	integratorOEmbedDomainRequest := *openapiclient.NewIntegratorOEmbedDomainRequest(*openapiclient.NewEndpoints(false, []string{"Schemes_example"}, "Url_example")) // IntegratorOEmbedDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorSettingsAPI.PatchMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId(context.Background(), oEmbedDomainId).IntegratorOEmbedDomainRequest(integratorOEmbedDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.PatchMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId`: IntegratorOEmbedDomainModel
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.PatchMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oEmbedDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMediaBridgeV1AppIdSettingsOembedDomainsOEmbedDomainIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integratorOEmbedDomainRequest** | [**IntegratorOEmbedDomainRequest**](IntegratorOEmbedDomainRequest.md) |  | 

### Return type

[**IntegratorOEmbedDomainModel**](IntegratorOEmbedDomainModel.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1AppIdSettingsObjectDefinitions

> map[string]IntegratorObjectCreationResponse PostMediaBridgeV1AppIdSettingsObjectDefinitions(ctx).IntegratorObjectCreationRequest(integratorObjectCreationRequest).Execute()



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
	integratorObjectCreationRequest := *openapiclient.NewIntegratorObjectCreationRequest([]string{"MediaTypes_example"}) // IntegratorObjectCreationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorSettingsAPI.PostMediaBridgeV1AppIdSettingsObjectDefinitions(context.Background()).IntegratorObjectCreationRequest(integratorObjectCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.PostMediaBridgeV1AppIdSettingsObjectDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1AppIdSettingsObjectDefinitions`: map[string]IntegratorObjectCreationResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.PostMediaBridgeV1AppIdSettingsObjectDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1AppIdSettingsObjectDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integratorObjectCreationRequest** | [**IntegratorObjectCreationRequest**](IntegratorObjectCreationRequest.md) |  | 

### Return type

[**map[string]IntegratorObjectCreationResponse**](IntegratorObjectCreationResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1AppIdSettingsOembedDomains

> IntegratorOEmbedDomainModel PostMediaBridgeV1AppIdSettingsOembedDomains(ctx).IntegratorOEmbedDomainRequest(integratorOEmbedDomainRequest).Execute()



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
	integratorOEmbedDomainRequest := *openapiclient.NewIntegratorOEmbedDomainRequest(*openapiclient.NewEndpoints(false, []string{"Schemes_example"}, "Url_example")) // IntegratorOEmbedDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorSettingsAPI.PostMediaBridgeV1AppIdSettingsOembedDomains(context.Background()).IntegratorOEmbedDomainRequest(integratorOEmbedDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.PostMediaBridgeV1AppIdSettingsOembedDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1AppIdSettingsOembedDomains`: IntegratorOEmbedDomainModel
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.PostMediaBridgeV1AppIdSettingsOembedDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1AppIdSettingsOembedDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integratorOEmbedDomainRequest** | [**IntegratorOEmbedDomainRequest**](IntegratorOEmbedDomainRequest.md) |  | 

### Return type

[**IntegratorOEmbedDomainModel**](IntegratorOEmbedDomainModel.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMediaBridgeV1AppIdSettingsRegister

> MediaBridgeProviderRegistrationResponse PostMediaBridgeV1AppIdSettingsRegister(ctx).MediaBridgeProviderPartial(mediaBridgeProviderPartial).Execute()



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
	mediaBridgeProviderPartial := *openapiclient.NewMediaBridgeProviderPartial(int64(123)) // MediaBridgeProviderPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorSettingsAPI.PostMediaBridgeV1AppIdSettingsRegister(context.Background()).MediaBridgeProviderPartial(mediaBridgeProviderPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.PostMediaBridgeV1AppIdSettingsRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMediaBridgeV1AppIdSettingsRegister`: MediaBridgeProviderRegistrationResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.PostMediaBridgeV1AppIdSettingsRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMediaBridgeV1AppIdSettingsRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaBridgeProviderPartial** | [**MediaBridgeProviderPartial**](MediaBridgeProviderPartial.md) |  | 

### Return type

[**MediaBridgeProviderRegistrationResponse**](MediaBridgeProviderRegistrationResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMediaBridgeV1AppIdSettings

> MediaBridgeProviderRegistrationResponse PutMediaBridgeV1AppIdSettings(ctx).MediaBridgeProviderPartial(mediaBridgeProviderPartial).Execute()



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
	mediaBridgeProviderPartial := *openapiclient.NewMediaBridgeProviderPartial(int64(123)) // MediaBridgeProviderPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorSettingsAPI.PutMediaBridgeV1AppIdSettings(context.Background()).MediaBridgeProviderPartial(mediaBridgeProviderPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorSettingsAPI.PutMediaBridgeV1AppIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMediaBridgeV1AppIdSettings`: MediaBridgeProviderRegistrationResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegratorSettingsAPI.PutMediaBridgeV1AppIdSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutMediaBridgeV1AppIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaBridgeProviderPartial** | [**MediaBridgeProviderPartial**](MediaBridgeProviderPartial.md) |  | 

### Return type

[**MediaBridgeProviderRegistrationResponse**](MediaBridgeProviderRegistrationResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

