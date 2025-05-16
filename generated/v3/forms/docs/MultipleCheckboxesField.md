# MultipleCheckboxesField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypeId** | **string** | A unique ID for this field&#39;s CRM object type. For example a CONTACT field will have the object type ID 0-1. | 
**Hidden** | **bool** | Whether a field should be hidden or not. Hidden fields won&#39;t appear on the form, but can be used to pass a value to a property without requiring the customer to fill it in. | 
**Name** | **string** | The identifier of the field. In combination with the object type ID, it must be unique. | 
**Options** | [**[]EnumeratedFieldOption**](EnumeratedFieldOption.md) | The list of available choices for this field. | 
**Description** | Pointer to **string** | Additional text helping the customer to complete the field. | [optional] 
**DefaultValues** | **[]string** | The values selected by default. Those values will be submitted unless the customer modifies them. | 
**DependentFields** | [**[]DependentField**](DependentField.md) | A list of other fields to make visible based on the value filled in for this field. | 
**Label** | **string** | The main label for the form field. | 
**FieldType** | **string** | Determines how the field will be displayed and validated. | [default to "multiple_checkboxes"]
**Required** | **bool** | Whether a value for this field is required when submitting the form. | 

## Methods

### NewMultipleCheckboxesField

`func NewMultipleCheckboxesField(objectTypeId string, hidden bool, name string, options []EnumeratedFieldOption, defaultValues []string, dependentFields []DependentField, label string, fieldType string, required bool, ) *MultipleCheckboxesField`

NewMultipleCheckboxesField instantiates a new MultipleCheckboxesField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleCheckboxesFieldWithDefaults

`func NewMultipleCheckboxesFieldWithDefaults() *MultipleCheckboxesField`

NewMultipleCheckboxesFieldWithDefaults instantiates a new MultipleCheckboxesField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *MultipleCheckboxesField) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *MultipleCheckboxesField) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *MultipleCheckboxesField) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetHidden

`func (o *MultipleCheckboxesField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MultipleCheckboxesField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *MultipleCheckboxesField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetName

`func (o *MultipleCheckboxesField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultipleCheckboxesField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MultipleCheckboxesField) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *MultipleCheckboxesField) GetOptions() []EnumeratedFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MultipleCheckboxesField) GetOptionsOk() (*[]EnumeratedFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MultipleCheckboxesField) SetOptions(v []EnumeratedFieldOption)`

SetOptions sets Options field to given value.


### GetDescription

`func (o *MultipleCheckboxesField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MultipleCheckboxesField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MultipleCheckboxesField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MultipleCheckboxesField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValues

`func (o *MultipleCheckboxesField) GetDefaultValues() []string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *MultipleCheckboxesField) GetDefaultValuesOk() (*[]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *MultipleCheckboxesField) SetDefaultValues(v []string)`

SetDefaultValues sets DefaultValues field to given value.


### GetDependentFields

`func (o *MultipleCheckboxesField) GetDependentFields() []DependentField`

GetDependentFields returns the DependentFields field if non-nil, zero value otherwise.

### GetDependentFieldsOk

`func (o *MultipleCheckboxesField) GetDependentFieldsOk() (*[]DependentField, bool)`

GetDependentFieldsOk returns a tuple with the DependentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFields

`func (o *MultipleCheckboxesField) SetDependentFields(v []DependentField)`

SetDependentFields sets DependentFields field to given value.


### GetLabel

`func (o *MultipleCheckboxesField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MultipleCheckboxesField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MultipleCheckboxesField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFieldType

`func (o *MultipleCheckboxesField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *MultipleCheckboxesField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *MultipleCheckboxesField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetRequired

`func (o *MultipleCheckboxesField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *MultipleCheckboxesField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *MultipleCheckboxesField) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


