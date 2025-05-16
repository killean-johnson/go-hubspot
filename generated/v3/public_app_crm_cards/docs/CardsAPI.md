# \CardsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ExtensionsCardsDevAppIdCardIdArchive**](CardsAPI.md#DeleteCrmV3ExtensionsCardsDevAppIdCardIdArchive) | **Delete** /crm/v3/extensions/cards-dev/{appId}/{cardId} | Delete a card
[**GetCrmV3ExtensionsCardsDevAppIdCardIdGetById**](CardsAPI.md#GetCrmV3ExtensionsCardsDevAppIdCardIdGetById) | **Get** /crm/v3/extensions/cards-dev/{appId}/{cardId} | Get a card.
[**GetCrmV3ExtensionsCardsDevAppIdGetAll**](CardsAPI.md#GetCrmV3ExtensionsCardsDevAppIdGetAll) | **Get** /crm/v3/extensions/cards-dev/{appId} | Get all cards
[**PatchCrmV3ExtensionsCardsDevAppIdCardIdUpdate**](CardsAPI.md#PatchCrmV3ExtensionsCardsDevAppIdCardIdUpdate) | **Patch** /crm/v3/extensions/cards-dev/{appId}/{cardId} | Update a card
[**PostCrmV3ExtensionsCardsDevAppIdCreate**](CardsAPI.md#PostCrmV3ExtensionsCardsDevAppIdCreate) | **Post** /crm/v3/extensions/cards-dev/{appId} | Create a new card



## DeleteCrmV3ExtensionsCardsDevAppIdCardIdArchive

> DeleteCrmV3ExtensionsCardsDevAppIdCardIdArchive(ctx, cardId, appId).Execute()

Delete a card



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
	cardId := "cardId_example" // string | The ID of the card to delete.
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CardsAPI.DeleteCrmV3ExtensionsCardsDevAppIdCardIdArchive(context.Background(), cardId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.DeleteCrmV3ExtensionsCardsDevAppIdCardIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** | The ID of the card to delete. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ExtensionsCardsDevAppIdCardIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetCrmV3ExtensionsCardsDevAppIdCardIdGetById

> PublicCardResponse GetCrmV3ExtensionsCardsDevAppIdCardIdGetById(ctx, cardId, appId).Execute()

Get a card.



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
	cardId := "cardId_example" // string | The ID of the target card.
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.GetCrmV3ExtensionsCardsDevAppIdCardIdGetById(context.Background(), cardId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCrmV3ExtensionsCardsDevAppIdCardIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExtensionsCardsDevAppIdCardIdGetById`: PublicCardResponse
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCrmV3ExtensionsCardsDevAppIdCardIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** | The ID of the target card. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCardsDevAppIdCardIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PublicCardResponse**](PublicCardResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ExtensionsCardsDevAppIdGetAll

> PublicCardListResponse GetCrmV3ExtensionsCardsDevAppIdGetAll(ctx, appId).Execute()

Get all cards



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
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.GetCrmV3ExtensionsCardsDevAppIdGetAll(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCrmV3ExtensionsCardsDevAppIdGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExtensionsCardsDevAppIdGetAll`: PublicCardListResponse
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCrmV3ExtensionsCardsDevAppIdGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCardsDevAppIdGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicCardListResponse**](PublicCardListResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsCardsDevAppIdCardIdUpdate

> PublicCardResponse PatchCrmV3ExtensionsCardsDevAppIdCardIdUpdate(ctx, cardId, appId).CardPatchRequest(cardPatchRequest).Execute()

Update a card



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
	cardId := "cardId_example" // string | The ID of the card to update.
	appId := int32(56) // int32 | The ID of the target app.
	cardPatchRequest := *openapiclient.NewCardPatchRequest() // CardPatchRequest | Card definition fields to be updated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.PatchCrmV3ExtensionsCardsDevAppIdCardIdUpdate(context.Background(), cardId, appId).CardPatchRequest(cardPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.PatchCrmV3ExtensionsCardsDevAppIdCardIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3ExtensionsCardsDevAppIdCardIdUpdate`: PublicCardResponse
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.PatchCrmV3ExtensionsCardsDevAppIdCardIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** | The ID of the card to update. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsCardsDevAppIdCardIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cardPatchRequest** | [**CardPatchRequest**](CardPatchRequest.md) | Card definition fields to be updated. | 

### Return type

[**PublicCardResponse**](PublicCardResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCardsDevAppIdCreate

> PublicCardResponse PostCrmV3ExtensionsCardsDevAppIdCreate(ctx, appId).CardCreateRequest(cardCreateRequest).Execute()

Create a new card



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
	appId := int32(56) // int32 | The ID of the target app.
	cardCreateRequest := *openapiclient.NewCardCreateRequest(*openapiclient.NewCardFetchBody([]openapiclient.CardObjectTypeBody{*openapiclient.NewCardObjectTypeBody("Name_example", []string{"PropertiesToSend_example"})}, "TargetUrl_example"), *openapiclient.NewCardDisplayBody([]openapiclient.CardDisplayProperty{*openapiclient.NewCardDisplayProperty("DataType_example", "Name_example", []openapiclient.DisplayOption{*openapiclient.NewDisplayOption("Name_example", "Label_example", "Type_example")}, "Label_example")}), "PetSpot", *openapiclient.NewCardActions([]string{"BaseUrls_example"})) // CardCreateRequest | The new card definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.PostCrmV3ExtensionsCardsDevAppIdCreate(context.Background(), appId).CardCreateRequest(cardCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.PostCrmV3ExtensionsCardsDevAppIdCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ExtensionsCardsDevAppIdCreate`: PublicCardResponse
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.PostCrmV3ExtensionsCardsDevAppIdCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCardsDevAppIdCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cardCreateRequest** | [**CardCreateRequest**](CardCreateRequest.md) | The new card definition. | 

### Return type

[**PublicCardResponse**](PublicCardResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

