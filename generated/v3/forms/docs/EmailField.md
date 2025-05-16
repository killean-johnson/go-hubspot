# EmailField

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
**Placeholder** | Pointer to **string** | The prompt text showing when the field isn&#39;t filled in. | [optional] 
**FieldType** | **string** | Determines how the field will be displayed and validated. | [default to "email"]
**Required** | **bool** | Whether a value for this field is required when submitting the form. | 
**Validation** | [**EmailFieldValidation**](EmailFieldValidation.md) |  | 

## Methods

### NewEmailField

`func NewEmailField(objectTypeId string, hidden bool, name string, dependentFields []DependentField, label string, fieldType string, required bool, validation EmailFieldValidation, ) *EmailField`

NewEmailField instantiates a new EmailField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailFieldWithDefaults

`func NewEmailFieldWithDefaults() *EmailField`

NewEmailFieldWithDefaults instantiates a new EmailField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *EmailField) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *EmailField) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *EmailField) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetHidden

`func (o *EmailField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *EmailField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *EmailField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetDefaultValue

`func (o *EmailField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *EmailField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *EmailField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *EmailField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetName

`func (o *EmailField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailField) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EmailField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDependentFields

`func (o *EmailField) GetDependentFields() []DependentField`

GetDependentFields returns the DependentFields field if non-nil, zero value otherwise.

### GetDependentFieldsOk

`func (o *EmailField) GetDependentFieldsOk() (*[]DependentField, bool)`

GetDependentFieldsOk returns a tuple with the DependentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFields

`func (o *EmailField) SetDependentFields(v []DependentField)`

SetDependentFields sets DependentFields field to given value.


### GetLabel

`func (o *EmailField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EmailField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EmailField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPlaceholder

`func (o *EmailField) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *EmailField) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *EmailField) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *EmailField) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetFieldType

`func (o *EmailField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *EmailField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *EmailField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetRequired

`func (o *EmailField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *EmailField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *EmailField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetValidation

`func (o *EmailField) GetValidation() EmailFieldValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *EmailField) GetValidationOk() (*EmailFieldValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *EmailField) SetValidation(v EmailFieldValidation)`

SetValidation sets Validation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


