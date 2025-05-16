# DependentFieldDependentField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypeId** | **string** |  | 
**Hidden** | **bool** |  | 
**DefaultValue** | Pointer to **string** | The value filled in by default. This value will be submitted unless the customer modifies it. | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DependentFields** | [**[]DependentField**](DependentField.md) |  | 
**Label** | **string** |  | 
**Placeholder** | Pointer to **string** | The prompt text showing when the field isn&#39;t filled in. | [optional] 
**FieldType** | **string** |  | [default to "payment_link_radio"]
**Required** | **bool** |  | 
**Validation** | [**NumberFieldValidation**](NumberFieldValidation.md) |  | 
**UseCountryCodeSelect** | **bool** | Whether to display a country code drop down next to the phone field. | 
**Options** | [**[]EnumeratedFieldOption**](EnumeratedFieldOption.md) |  | 
**DefaultValues** | **[]string** |  | 
**AllowMultipleFiles** | **bool** | Whether to allow the upload of multiple files. | 

## Methods

### NewDependentFieldDependentField

`func NewDependentFieldDependentField(objectTypeId string, hidden bool, name string, dependentFields []DependentField, label string, fieldType string, required bool, validation NumberFieldValidation, useCountryCodeSelect bool, options []EnumeratedFieldOption, defaultValues []string, allowMultipleFiles bool, ) *DependentFieldDependentField`

NewDependentFieldDependentField instantiates a new DependentFieldDependentField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentFieldDependentFieldWithDefaults

`func NewDependentFieldDependentFieldWithDefaults() *DependentFieldDependentField`

NewDependentFieldDependentFieldWithDefaults instantiates a new DependentFieldDependentField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *DependentFieldDependentField) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *DependentFieldDependentField) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *DependentFieldDependentField) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetHidden

`func (o *DependentFieldDependentField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *DependentFieldDependentField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *DependentFieldDependentField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetDefaultValue

`func (o *DependentFieldDependentField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DependentFieldDependentField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DependentFieldDependentField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DependentFieldDependentField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetName

`func (o *DependentFieldDependentField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependentFieldDependentField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependentFieldDependentField) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DependentFieldDependentField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DependentFieldDependentField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DependentFieldDependentField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DependentFieldDependentField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDependentFields

`func (o *DependentFieldDependentField) GetDependentFields() []DependentField`

GetDependentFields returns the DependentFields field if non-nil, zero value otherwise.

### GetDependentFieldsOk

`func (o *DependentFieldDependentField) GetDependentFieldsOk() (*[]DependentField, bool)`

GetDependentFieldsOk returns a tuple with the DependentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFields

`func (o *DependentFieldDependentField) SetDependentFields(v []DependentField)`

SetDependentFields sets DependentFields field to given value.


### GetLabel

`func (o *DependentFieldDependentField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DependentFieldDependentField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DependentFieldDependentField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPlaceholder

`func (o *DependentFieldDependentField) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *DependentFieldDependentField) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *DependentFieldDependentField) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *DependentFieldDependentField) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetFieldType

`func (o *DependentFieldDependentField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *DependentFieldDependentField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *DependentFieldDependentField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetRequired

`func (o *DependentFieldDependentField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DependentFieldDependentField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DependentFieldDependentField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetValidation

`func (o *DependentFieldDependentField) GetValidation() NumberFieldValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *DependentFieldDependentField) GetValidationOk() (*NumberFieldValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *DependentFieldDependentField) SetValidation(v NumberFieldValidation)`

SetValidation sets Validation field to given value.


### GetUseCountryCodeSelect

`func (o *DependentFieldDependentField) GetUseCountryCodeSelect() bool`

GetUseCountryCodeSelect returns the UseCountryCodeSelect field if non-nil, zero value otherwise.

### GetUseCountryCodeSelectOk

`func (o *DependentFieldDependentField) GetUseCountryCodeSelectOk() (*bool, bool)`

GetUseCountryCodeSelectOk returns a tuple with the UseCountryCodeSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCountryCodeSelect

`func (o *DependentFieldDependentField) SetUseCountryCodeSelect(v bool)`

SetUseCountryCodeSelect sets UseCountryCodeSelect field to given value.


### GetOptions

`func (o *DependentFieldDependentField) GetOptions() []EnumeratedFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DependentFieldDependentField) GetOptionsOk() (*[]EnumeratedFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DependentFieldDependentField) SetOptions(v []EnumeratedFieldOption)`

SetOptions sets Options field to given value.


### GetDefaultValues

`func (o *DependentFieldDependentField) GetDefaultValues() []string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *DependentFieldDependentField) GetDefaultValuesOk() (*[]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *DependentFieldDependentField) SetDefaultValues(v []string)`

SetDefaultValues sets DefaultValues field to given value.


### GetAllowMultipleFiles

`func (o *DependentFieldDependentField) GetAllowMultipleFiles() bool`

GetAllowMultipleFiles returns the AllowMultipleFiles field if non-nil, zero value otherwise.

### GetAllowMultipleFilesOk

`func (o *DependentFieldDependentField) GetAllowMultipleFilesOk() (*bool, bool)`

GetAllowMultipleFilesOk returns a tuple with the AllowMultipleFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleFiles

`func (o *DependentFieldDependentField) SetAllowMultipleFiles(v bool)`

SetAllowMultipleFiles sets AllowMultipleFiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


