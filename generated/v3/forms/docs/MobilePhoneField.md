# MobilePhoneField

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
**FieldType** | **string** | Determines how the field will be displayed and validated. | [default to "mobile_phone"]
**Required** | **bool** | Whether a value for this field is required when submitting the form. | 
**Validation** | [**PhoneFieldValidation**](PhoneFieldValidation.md) |  | 

## Methods

### NewMobilePhoneField

`func NewMobilePhoneField(objectTypeId string, hidden bool, name string, dependentFields []DependentField, label string, fieldType string, required bool, validation PhoneFieldValidation, ) *MobilePhoneField`

NewMobilePhoneField instantiates a new MobilePhoneField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobilePhoneFieldWithDefaults

`func NewMobilePhoneFieldWithDefaults() *MobilePhoneField`

NewMobilePhoneFieldWithDefaults instantiates a new MobilePhoneField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *MobilePhoneField) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *MobilePhoneField) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *MobilePhoneField) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetHidden

`func (o *MobilePhoneField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MobilePhoneField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *MobilePhoneField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetDefaultValue

`func (o *MobilePhoneField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MobilePhoneField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MobilePhoneField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MobilePhoneField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetName

`func (o *MobilePhoneField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobilePhoneField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobilePhoneField) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MobilePhoneField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MobilePhoneField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MobilePhoneField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MobilePhoneField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDependentFields

`func (o *MobilePhoneField) GetDependentFields() []DependentField`

GetDependentFields returns the DependentFields field if non-nil, zero value otherwise.

### GetDependentFieldsOk

`func (o *MobilePhoneField) GetDependentFieldsOk() (*[]DependentField, bool)`

GetDependentFieldsOk returns a tuple with the DependentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFields

`func (o *MobilePhoneField) SetDependentFields(v []DependentField)`

SetDependentFields sets DependentFields field to given value.


### GetLabel

`func (o *MobilePhoneField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MobilePhoneField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MobilePhoneField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPlaceholder

`func (o *MobilePhoneField) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *MobilePhoneField) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *MobilePhoneField) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *MobilePhoneField) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetFieldType

`func (o *MobilePhoneField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *MobilePhoneField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *MobilePhoneField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetRequired

`func (o *MobilePhoneField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *MobilePhoneField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *MobilePhoneField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetValidation

`func (o *MobilePhoneField) GetValidation() PhoneFieldValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *MobilePhoneField) GetValidationOk() (*PhoneFieldValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *MobilePhoneField) SetValidation(v PhoneFieldValidation)`

SetValidation sets Validation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


