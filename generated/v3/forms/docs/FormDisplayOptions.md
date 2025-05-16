# FormDisplayOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RenderRawHtml** | **bool** | Whether the form will render as raw HTML as opposed to inside an iFrame. | 
**CssClass** | Pointer to **string** |  | [optional] 
**Theme** | **string** | The theme used for styling the input fields. This will not apply if the form is added to a HubSpot CMS page. | 
**SubmitButtonText** | **string** | The text displayed on the form submit button. | 
**Style** | [**FormStyle**](FormStyle.md) |  | 

## Methods

### NewFormDisplayOptions

`func NewFormDisplayOptions(renderRawHtml bool, theme string, submitButtonText string, style FormStyle, ) *FormDisplayOptions`

NewFormDisplayOptions instantiates a new FormDisplayOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDisplayOptionsWithDefaults

`func NewFormDisplayOptionsWithDefaults() *FormDisplayOptions`

NewFormDisplayOptionsWithDefaults instantiates a new FormDisplayOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenderRawHtml

`func (o *FormDisplayOptions) GetRenderRawHtml() bool`

GetRenderRawHtml returns the RenderRawHtml field if non-nil, zero value otherwise.

### GetRenderRawHtmlOk

`func (o *FormDisplayOptions) GetRenderRawHtmlOk() (*bool, bool)`

GetRenderRawHtmlOk returns a tuple with the RenderRawHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderRawHtml

`func (o *FormDisplayOptions) SetRenderRawHtml(v bool)`

SetRenderRawHtml sets RenderRawHtml field to given value.


### GetCssClass

`func (o *FormDisplayOptions) GetCssClass() string`

GetCssClass returns the CssClass field if non-nil, zero value otherwise.

### GetCssClassOk

`func (o *FormDisplayOptions) GetCssClassOk() (*string, bool)`

GetCssClassOk returns a tuple with the CssClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssClass

`func (o *FormDisplayOptions) SetCssClass(v string)`

SetCssClass sets CssClass field to given value.

### HasCssClass

`func (o *FormDisplayOptions) HasCssClass() bool`

HasCssClass returns a boolean if a field has been set.

### GetTheme

`func (o *FormDisplayOptions) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *FormDisplayOptions) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *FormDisplayOptions) SetTheme(v string)`

SetTheme sets Theme field to given value.


### GetSubmitButtonText

`func (o *FormDisplayOptions) GetSubmitButtonText() string`

GetSubmitButtonText returns the SubmitButtonText field if non-nil, zero value otherwise.

### GetSubmitButtonTextOk

`func (o *FormDisplayOptions) GetSubmitButtonTextOk() (*string, bool)`

GetSubmitButtonTextOk returns a tuple with the SubmitButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitButtonText

`func (o *FormDisplayOptions) SetSubmitButtonText(v string)`

SetSubmitButtonText sets SubmitButtonText field to given value.


### GetStyle

`func (o *FormDisplayOptions) GetStyle() FormStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *FormDisplayOptions) GetStyleOk() (*FormStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *FormDisplayOptions) SetStyle(v FormStyle)`

SetStyle sets Style field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


