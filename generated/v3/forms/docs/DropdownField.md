# DropdownField

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
**Placeholder** | Pointer to **string** | The prompt text showing when the field isn&#39;t filled in. | [optional] 
**FieldType** | **string** | Determines how the field will be displayed and validated. | [default to "dropdown"]
**Required** | **bool** | Whether a value for this field is required when submitting the form. | 

## Methods

### NewDropdownField

`func NewDropdownField(objectTypeId string, hidden bool, name string, options []EnumeratedFieldOption, defaultValues []string, dependentFields []DependentField, label string, fieldType string, required bool, ) *DropdownField`

NewDropdownField instantiates a new DropdownField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropdownFieldWithDefaults

`func NewDropdownFieldWithDefaults() *DropdownField`

NewDropdownFieldWithDefaults instantiates a new DropdownField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *DropdownField) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *DropdownField) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *DropdownField) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetHidden

`func (o *DropdownField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *DropdownField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *DropdownField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetName

`func (o *DropdownField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DropdownField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DropdownField) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *DropdownField) GetOptions() []EnumeratedFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DropdownField) GetOptionsOk() (*[]EnumeratedFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DropdownField) SetOptions(v []EnumeratedFieldOption)`

SetOptions sets Options field to given value.


### GetDescription

`func (o *DropdownField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DropdownField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DropdownField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DropdownField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValues

`func (o *DropdownField) GetDefaultValues() []string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *DropdownField) GetDefaultValuesOk() (*[]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *DropdownField) SetDefaultValues(v []string)`

SetDefaultValues sets DefaultValues field to given value.


### GetDependentFields

`func (o *DropdownField) GetDependentFields() []DependentField`

GetDependentFields returns the DependentFields field if non-nil, zero value otherwise.

### GetDependentFieldsOk

`func (o *DropdownField) GetDependentFieldsOk() (*[]DependentField, bool)`

GetDependentFieldsOk returns a tuple with the DependentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFields

`func (o *DropdownField) SetDependentFields(v []DependentField)`

SetDependentFields sets DependentFields field to given value.


### GetLabel

`func (o *DropdownField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DropdownField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DropdownField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPlaceholder

`func (o *DropdownField) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *DropdownField) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *DropdownField) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *DropdownField) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetFieldType

`func (o *DropdownField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *DropdownField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *DropdownField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetRequired

`func (o *DropdownField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DropdownField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DropdownField) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


