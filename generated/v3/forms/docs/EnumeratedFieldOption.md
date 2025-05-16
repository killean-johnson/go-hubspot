# EnumeratedFieldOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayOrder** | **int32** | The order the choices will be displayed in. | 
**Description** | Pointer to **string** |  | [optional] 
**Label** | **string** | The visible label for this choice. | 
**Value** | **string** | The value which will be submitted if this choice is selected. | 

## Methods

### NewEnumeratedFieldOption

`func NewEnumeratedFieldOption(displayOrder int32, label string, value string, ) *EnumeratedFieldOption`

NewEnumeratedFieldOption instantiates a new EnumeratedFieldOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumeratedFieldOptionWithDefaults

`func NewEnumeratedFieldOptionWithDefaults() *EnumeratedFieldOption`

NewEnumeratedFieldOptionWithDefaults instantiates a new EnumeratedFieldOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayOrder

`func (o *EnumeratedFieldOption) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *EnumeratedFieldOption) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *EnumeratedFieldOption) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetDescription

`func (o *EnumeratedFieldOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnumeratedFieldOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnumeratedFieldOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnumeratedFieldOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *EnumeratedFieldOption) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EnumeratedFieldOption) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EnumeratedFieldOption) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *EnumeratedFieldOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnumeratedFieldOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnumeratedFieldOption) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


