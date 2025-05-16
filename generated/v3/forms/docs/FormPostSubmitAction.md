# FormPostSubmitAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The action to take after submit. The default action is displaying a thank you message. | 
**Value** | **string** | The thank you text or the page to redirect to. | 

## Methods

### NewFormPostSubmitAction

`func NewFormPostSubmitAction(type_ string, value string, ) *FormPostSubmitAction`

NewFormPostSubmitAction instantiates a new FormPostSubmitAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormPostSubmitActionWithDefaults

`func NewFormPostSubmitActionWithDefaults() *FormPostSubmitAction`

NewFormPostSubmitActionWithDefaults instantiates a new FormPostSubmitAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormPostSubmitAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormPostSubmitAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormPostSubmitAction) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *FormPostSubmitAction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FormPostSubmitAction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FormPostSubmitAction) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


