# OptionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hidden** | **bool** |  | 
**DisplayOrder** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Label** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewOptionInput

`func NewOptionInput(hidden bool, displayOrder int32, label string, value string, ) *OptionInput`

NewOptionInput instantiates a new OptionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionInputWithDefaults

`func NewOptionInputWithDefaults() *OptionInput`

NewOptionInputWithDefaults instantiates a new OptionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHidden

`func (o *OptionInput) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *OptionInput) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *OptionInput) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetDisplayOrder

`func (o *OptionInput) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *OptionInput) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *OptionInput) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetDescription

`func (o *OptionInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *OptionInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OptionInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OptionInput) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *OptionInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OptionInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OptionInput) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


