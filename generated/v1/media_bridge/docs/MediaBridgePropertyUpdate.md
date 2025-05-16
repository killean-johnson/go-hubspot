# MediaBridgePropertyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to [**[]OptionInput**](OptionInput.md) |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CalculationFormula** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**HasUniqueValue** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**FieldType** | Pointer to **string** |  | [optional] 
**FormField** | Pointer to **bool** |  | [optional] 

## Methods

### NewMediaBridgePropertyUpdate

`func NewMediaBridgePropertyUpdate() *MediaBridgePropertyUpdate`

NewMediaBridgePropertyUpdate instantiates a new MediaBridgePropertyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaBridgePropertyUpdateWithDefaults

`func NewMediaBridgePropertyUpdateWithDefaults() *MediaBridgePropertyUpdate`

NewMediaBridgePropertyUpdateWithDefaults instantiates a new MediaBridgePropertyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *MediaBridgePropertyUpdate) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *MediaBridgePropertyUpdate) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *MediaBridgePropertyUpdate) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *MediaBridgePropertyUpdate) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetHidden

`func (o *MediaBridgePropertyUpdate) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MediaBridgePropertyUpdate) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *MediaBridgePropertyUpdate) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *MediaBridgePropertyUpdate) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetOptions

`func (o *MediaBridgePropertyUpdate) GetOptions() []OptionInput`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MediaBridgePropertyUpdate) GetOptionsOk() (*[]OptionInput, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MediaBridgePropertyUpdate) SetOptions(v []OptionInput)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MediaBridgePropertyUpdate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *MediaBridgePropertyUpdate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *MediaBridgePropertyUpdate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *MediaBridgePropertyUpdate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *MediaBridgePropertyUpdate) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetDescription

`func (o *MediaBridgePropertyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MediaBridgePropertyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MediaBridgePropertyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MediaBridgePropertyUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCalculationFormula

`func (o *MediaBridgePropertyUpdate) GetCalculationFormula() string`

GetCalculationFormula returns the CalculationFormula field if non-nil, zero value otherwise.

### GetCalculationFormulaOk

`func (o *MediaBridgePropertyUpdate) GetCalculationFormulaOk() (*string, bool)`

GetCalculationFormulaOk returns a tuple with the CalculationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationFormula

`func (o *MediaBridgePropertyUpdate) SetCalculationFormula(v string)`

SetCalculationFormula sets CalculationFormula field to given value.

### HasCalculationFormula

`func (o *MediaBridgePropertyUpdate) HasCalculationFormula() bool`

HasCalculationFormula returns a boolean if a field has been set.

### GetLabel

`func (o *MediaBridgePropertyUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MediaBridgePropertyUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MediaBridgePropertyUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MediaBridgePropertyUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetHasUniqueValue

`func (o *MediaBridgePropertyUpdate) GetHasUniqueValue() bool`

GetHasUniqueValue returns the HasUniqueValue field if non-nil, zero value otherwise.

### GetHasUniqueValueOk

`func (o *MediaBridgePropertyUpdate) GetHasUniqueValueOk() (*bool, bool)`

GetHasUniqueValueOk returns a tuple with the HasUniqueValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUniqueValue

`func (o *MediaBridgePropertyUpdate) SetHasUniqueValue(v bool)`

SetHasUniqueValue sets HasUniqueValue field to given value.

### HasHasUniqueValue

`func (o *MediaBridgePropertyUpdate) HasHasUniqueValue() bool`

HasHasUniqueValue returns a boolean if a field has been set.

### GetType

`func (o *MediaBridgePropertyUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaBridgePropertyUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaBridgePropertyUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MediaBridgePropertyUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFieldType

`func (o *MediaBridgePropertyUpdate) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *MediaBridgePropertyUpdate) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *MediaBridgePropertyUpdate) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *MediaBridgePropertyUpdate) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetFormField

`func (o *MediaBridgePropertyUpdate) GetFormField() bool`

GetFormField returns the FormField field if non-nil, zero value otherwise.

### GetFormFieldOk

`func (o *MediaBridgePropertyUpdate) GetFormFieldOk() (*bool, bool)`

GetFormFieldOk returns a tuple with the FormField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormField

`func (o *MediaBridgePropertyUpdate) SetFormField(v bool)`

SetFormField sets FormField field to given value.

### HasFormField

`func (o *MediaBridgePropertyUpdate) HasFormField() bool`

HasFormField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


