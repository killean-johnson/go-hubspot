# \PublicVNextEmailsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3EmailsAbTestCreateVariation**](PublicVNextEmailsAPI.md#PostMarketingV3EmailsAbTestCreateVariation) | **Post** /marketing/v3/emails/ab-test/create-variation | 
[**PostMarketingV3EmailsClone**](PublicVNextEmailsAPI.md#PostMarketingV3EmailsClone) | **Post** /marketing/v3/emails/clone | 



## PostMarketingV3EmailsAbTestCreateVariation

> PublicEmail PostMarketingV3EmailsAbTestCreateVariation(ctx).AbTestCreateRequestVNext(abTestCreateRequestVNext).Execute()



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
	abTestCreateRequestVNext := *openapiclient.NewAbTestCreateRequestVNext("VariationName_example", "ContentId_example") // AbTestCreateRequestVNext | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicVNextEmailsAPI.PostMarketingV3EmailsAbTestCreateVariation(context.Background()).AbTestCreateRequestVNext(abTestCreateRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicVNextEmailsAPI.PostMarketingV3EmailsAbTestCreateVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3EmailsAbTestCreateVariation`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `PublicVNextEmailsAPI.PostMarketingV3EmailsAbTestCreateVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsAbTestCreateVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **abTestCreateRequestVNext** | [**AbTestCreateRequestVNext**](AbTestCreateRequestVNext.md) |  | 

### Return type

[**PublicEmail**](PublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3EmailsClone

> PublicEmail PostMarketingV3EmailsClone(ctx).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()



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
	contentCloneRequestVNext := *openapiclient.NewContentCloneRequestVNext("Id_example") // ContentCloneRequestVNext | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicVNextEmailsAPI.PostMarketingV3EmailsClone(context.Background()).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicVNextEmailsAPI.PostMarketingV3EmailsClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3EmailsClone`: PublicEmail
	fmt.Fprintf(os.Stdout, "Response from `PublicVNextEmailsAPI.PostMarketingV3EmailsClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3EmailsCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentCloneRequestVNext** | [**ContentCloneRequestVNext**](ContentCloneRequestVNext.md) |  | 

### Return type

[**PublicEmail**](PublicEmail.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

