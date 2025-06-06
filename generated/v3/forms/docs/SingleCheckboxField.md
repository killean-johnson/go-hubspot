# SingleCheckboxField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypeId** | **string** | A unique ID for this field&#39;s CRM object type. For example a CONTACT field will have the object type ID 0-1. | 
**Hidden** | **bool** | Whether a field should be hidden or not. Hidden fields won&#39;t appear on the form, but can be used to pass a value to a property without requiring the customer to fill it in. | 
**DefaultValue** | Pointer to **string** | The value filled in by default. This value will be submitted unless the customer modifies it. | [optional] 
**Name** | **string** | The identifier of the field. In combination with the object type ID, it must be unique. | 
**Description** | Pointer to **string** | Additional text helping the customer to complete the field. | [optional] 
**DependentFields** | [**[]DependentField**](DependentField.md) | A list of other fields to make visible based on the value filled in for this field. | 
**Label** | **string** | The main label for the form field. | 
**FieldType** | **string** | Determines how the field will be displayed and validated. | [default to "single_checkbox"]
**Required** | **bool** | Whether a value for this field is required when submitting the form. | 

## Methods

### NewSingleCheckboxField

`func NewSingleCheckboxField(objectTypeId string, hidden bool, name string, dependentFields []DependentField, label string, fieldType string, required bool, ) *SingleCheckboxField`

NewSingleCheckboxField instantiates a new SingleCheckboxField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleCheckboxFieldWithDefaults

`func NewSingleCheckboxFieldWithDefaults() *SingleCheckboxField`

NewSingleCheckboxFieldWithDefaults instantiates a new SingleCheckboxField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *SingleCheckboxField) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *SingleCheckboxField) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *SingleCheckboxField) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetHidden

`func (o *SingleCheckboxField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *SingleCheckboxField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *SingleCheckboxField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetDefaultValue

`func (o *SingleCheckboxField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SingleCheckboxField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SingleCheckboxField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SingleCheckboxField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetName

`func (o *SingleCheckboxField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SingleCheckboxField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SingleCheckboxField) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SingleCheckboxField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SingleCheckboxField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SingleCheckboxField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SingleCheckboxField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDependentFields

`func (o *SingleCheckboxField) GetDependentFields() []DependentField`

GetDependentFields returns the DependentFields field if non-nil, zero value otherwise.

### GetDependentFieldsOk

`func (o *SingleCheckboxField) GetDependentFieldsOk() (*[]DependentField, bool)`

GetDependentFieldsOk returns a tuple with the DependentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFields

`func (o *SingleCheckboxField) SetDependentFields(v []DependentField)`

SetDependentFields sets DependentFields field to given value.


### GetLabel

`func (o *SingleCheckboxField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SingleCheckboxField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SingleCheckboxField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFieldType

`func (o *SingleCheckboxField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *SingleCheckboxField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *SingleCheckboxField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetRequired

`func (o *SingleCheckboxField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SingleCheckboxField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SingleCheckboxField) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


