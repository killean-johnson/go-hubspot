# \FormsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3FormsFormIdArchive**](FormsAPI.md#DeleteMarketingV3FormsFormIdArchive) | **Delete** /marketing/v3/forms/{formId} | Archive a form definition
[**GetMarketingV3FormsFormIdGetById**](FormsAPI.md#GetMarketingV3FormsFormIdGetById) | **Get** /marketing/v3/forms/{formId} | Get a form definition
[**GetMarketingV3FormsGetPage**](FormsAPI.md#GetMarketingV3FormsGetPage) | **Get** /marketing/v3/forms/ | Get a list of forms
[**PatchMarketingV3FormsFormIdUpdate**](FormsAPI.md#PatchMarketingV3FormsFormIdUpdate) | **Patch** /marketing/v3/forms/{formId} | Partially update a form definition
[**PostMarketingV3FormsCreate**](FormsAPI.md#PostMarketingV3FormsCreate) | **Post** /marketing/v3/forms/ | Create a form
[**PutMarketingV3FormsFormIdReplace**](FormsAPI.md#PutMarketingV3FormsFormIdReplace) | **Put** /marketing/v3/forms/{formId} | Update a form definition



## DeleteMarketingV3FormsFormIdArchive

> DeleteMarketingV3FormsFormIdArchive(ctx, formId).Execute()

Archive a form definition



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
	formId := "formId_example" // string | The ID of the form to archive.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FormsAPI.DeleteMarketingV3FormsFormIdArchive(context.Background(), formId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.DeleteMarketingV3FormsFormIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **string** | The ID of the form to archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3FormsFormIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3FormsFormIdGetById

> FormDefinitionBase GetMarketingV3FormsFormIdGetById(ctx, formId).Archived(archived).Execute()

Get a form definition



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
	formId := "formId_example" // string | The unique identifier of the form
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.GetMarketingV3FormsFormIdGetById(context.Background(), formId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.GetMarketingV3FormsFormIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3FormsFormIdGetById`: FormDefinitionBase
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.GetMarketingV3FormsFormIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **string** | The unique identifier of the form | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3FormsFormIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

[**FormDefinitionBase**](FormDefinitionBase.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3FormsGetPage

> CollectionResponseFormDefinitionBaseForwardPaging GetMarketingV3FormsGetPage(ctx).After(after).Limit(limit).Archived(archived).FormTypes(formTypes).Execute()

Get a list of forms



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
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional)
	formTypes := []string{"FormTypes_example"} // []string | The form types to be included in the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.GetMarketingV3FormsGetPage(context.Background()).After(after).Limit(limit).Archived(archived).FormTypes(formTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.GetMarketingV3FormsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3FormsGetPage`: CollectionResponseFormDefinitionBaseForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.GetMarketingV3FormsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3FormsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **archived** | **bool** | Whether to return only results that have been archived. | 
 **formTypes** | **[]string** | The form types to be included in the results. | 

### Return type

[**CollectionResponseFormDefinitionBaseForwardPaging**](CollectionResponseFormDefinitionBaseForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingV3FormsFormIdUpdate

> FormDefinitionBase PatchMarketingV3FormsFormIdUpdate(ctx, formId).HubSpotFormDefinitionPatchRequest(hubSpotFormDefinitionPatchRequest).Execute()

Partially update a form definition



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
	formId := "formId_example" // string | The ID of the form to update.
	hubSpotFormDefinitionPatchRequest := *openapiclient.NewHubSpotFormDefinitionPatchRequest() // HubSpotFormDefinitionPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.PatchMarketingV3FormsFormIdUpdate(context.Background(), formId).HubSpotFormDefinitionPatchRequest(hubSpotFormDefinitionPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.PatchMarketingV3FormsFormIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingV3FormsFormIdUpdate`: FormDefinitionBase
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.PatchMarketingV3FormsFormIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **string** | The ID of the form to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3FormsFormIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubSpotFormDefinitionPatchRequest** | [**HubSpotFormDefinitionPatchRequest**](HubSpotFormDefinitionPatchRequest.md) |  | 

### Return type

[**FormDefinitionBase**](FormDefinitionBase.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3FormsCreate

> FormDefinitionBase PostMarketingV3FormsCreate(ctx).FormDefinitionCreateRequestBase(formDefinitionCreateRequestBase).Execute()

Create a form



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	formDefinitionCreateRequestBase := openapiclient.FormDefinitionCreateRequestBase{HubSpotFormDefinitionCreateRequest: openapiclient.NewHubSpotFormDefinitionCreateRequest("FormType_example", "Name_example", time.Now(), time.Now(), false, []openapiclient.FieldGroup{*openapiclient.NewFieldGroup("GroupType_example", "RichTextType_example", []openapiclient.DependentFieldDependentField{openapiclient.DependentField_dependentField{DatepickerField: openapiclient.NewDatepickerField("ObjectTypeId_example", false, "Name_example", []openapiclient.DependentField{*openapiclient.NewDependentField(*openapiclient.NewDependentFieldFilter("RangeStart_example", []string{"Values_example"}, "Value_example", "Operator_example", "RangeEnd_example"), openapiclient.DependentField_dependentField{DatepickerField: openapiclient.NewDatepickerField("ObjectTypeId_example", false, "Name_example", []openapiclient.DependentField{*openapiclient.NewDependentField(*openapiclient.NewDependentFieldFilter("RangeStart_example", []string{"Values_example"}, "Value_example", "Operator_example", "RangeEnd_example"), openapiclient.DependentField_dependentField{DatepickerField: })}, "Label_example", "FieldType_example", false)})}, "Label_example", "FieldType_example", false)}})}, *openapiclient.NewHubSpotFormConfiguration(false, false, false, *openapiclient.NewFormPostSubmitAction("Type_example", "Value_example"), "Language_example", false, false, false, false, false, []string{"NotifyRecipients_example"}), *openapiclient.NewFormDisplayOptions(false, "Theme_example", "SubmitButtonText_example", *openapiclient.NewFormStyle("LabelTextSize_example", "LegalConsentTextColor_example", "FontFamily_example", "LegalConsentTextSize_example", "BackgroundWidth_example", "HelpTextSize_example", "SubmitFontColor_example", "LabelTextColor_example", "SubmitAlignment_example", "SubmitSize_example", "HelpTextColor_example", "SubmitColor_example")), openapiclient.HubSpotFormDefinitionCreateRequest_allOf_legalConsentOptions{LegalConsentOptionsExplicitConsentToProcess: openapiclient.NewLegalConsentOptionsExplicitConsentToProcess([]openapiclient.LegalConsentCheckbox{*openapiclient.NewLegalConsentCheckbox(int32(123), "Label_example", false)}, "Type_example", "PrivacyText_example")})} // FormDefinitionCreateRequestBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.PostMarketingV3FormsCreate(context.Background()).FormDefinitionCreateRequestBase(formDefinitionCreateRequestBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.PostMarketingV3FormsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3FormsCreate`: FormDefinitionBase
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.PostMarketingV3FormsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3FormsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formDefinitionCreateRequestBase** | [**FormDefinitionCreateRequestBase**](FormDefinitionCreateRequestBase.md) |  | 

### Return type

[**FormDefinitionBase**](FormDefinitionBase.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingV3FormsFormIdReplace

> FormDefinitionBase PutMarketingV3FormsFormIdReplace(ctx, formId).HubSpotFormDefinition(hubSpotFormDefinition).Execute()

Update a form definition



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	formId := "formId_example" // string | 
	hubSpotFormDefinition := *openapiclient.NewHubSpotFormDefinition("FormType_example", "Id_example", "Name_example", time.Now(), time.Now(), false, []openapiclient.FieldGroup{*openapiclient.NewFieldGroup("GroupType_example", "RichTextType_example", []openapiclient.DependentFieldDependentField{openapiclient.DependentField_dependentField{DatepickerField: openapiclient.NewDatepickerField("ObjectTypeId_example", false, "Name_example", []openapiclient.DependentField{*openapiclient.NewDependentField(*openapiclient.NewDependentFieldFilter("RangeStart_example", []string{"Values_example"}, "Value_example", "Operator_example", "RangeEnd_example"), openapiclient.DependentField_dependentField{DatepickerField: openapiclient.NewDatepickerField("ObjectTypeId_example", false, "Name_example", []openapiclient.DependentField{*openapiclient.NewDependentField(*openapiclient.NewDependentFieldFilter("RangeStart_example", []string{"Values_example"}, "Value_example", "Operator_example", "RangeEnd_example"), openapiclient.DependentField_dependentField{DatepickerField: })}, "Label_example", "FieldType_example", false)})}, "Label_example", "FieldType_example", false)}})}, *openapiclient.NewHubSpotFormConfiguration(false, false, false, *openapiclient.NewFormPostSubmitAction("Type_example", "Value_example"), "Language_example", false, false, false, false, false, []string{"NotifyRecipients_example"}), *openapiclient.NewFormDisplayOptions(false, "Theme_example", "SubmitButtonText_example", *openapiclient.NewFormStyle("LabelTextSize_example", "LegalConsentTextColor_example", "FontFamily_example", "LegalConsentTextSize_example", "BackgroundWidth_example", "HelpTextSize_example", "SubmitFontColor_example", "LabelTextColor_example", "SubmitAlignment_example", "SubmitSize_example", "HelpTextColor_example", "SubmitColor_example")), openapiclient.HubSpotFormDefinition_allOf_legalConsentOptions{LegalConsentOptionsExplicitConsentToProcess: openapiclient.NewLegalConsentOptionsExplicitConsentToProcess([]openapiclient.LegalConsentCheckbox{*openapiclient.NewLegalConsentCheckbox(int32(123), "Label_example", false)}, "Type_example", "PrivacyText_example")}) // HubSpotFormDefinition | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.PutMarketingV3FormsFormIdReplace(context.Background(), formId).HubSpotFormDefinition(hubSpotFormDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.PutMarketingV3FormsFormIdReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingV3FormsFormIdReplace`: FormDefinitionBase
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.PutMarketingV3FormsFormIdReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingV3FormsFormIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubSpotFormDefinition** | [**HubSpotFormDefinition**](HubSpotFormDefinition.md) |  | 

### Return type

[**FormDefinitionBase**](FormDefinitionBase.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

