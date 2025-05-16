# PaymentLinkRadioField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypeId** | **string** |  | 
**Hidden** | **bool** |  | 
**Name** | **string** |  | 
**Options** | [**[]EnumeratedFieldOption**](EnumeratedFieldOption.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**DefaultValues** | **[]string** |  | 
**DependentFields** | [**[]DependentField**](DependentField.md) |  | 
**Label** | **string** |  | 
**FieldType** | **string** |  | [default to "payment_link_radio"]
**Required** | **bool** |  | 

## Methods

### NewPaymentLinkRadioField

`func NewPaymentLinkRadioField(objectTypeId string, hidden bool, name string, options []EnumeratedFieldOption, defaultValues []string, dependentFields []DependentField, label string, fieldType string, required bool, ) *PaymentLinkRadioField`

NewPaymentLinkRadioField instantiates a new PaymentLinkRadioField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinkRadioFieldWithDefaults

`func NewPaymentLinkRadioFieldWithDefaults() *PaymentLinkRadioField`

NewPaymentLinkRadioFieldWithDefaults instantiates a new PaymentLinkRadioField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *PaymentLinkRadioField) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *PaymentLinkRadioField) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *PaymentLinkRadioField) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetHidden

`func (o *PaymentLinkRadioField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PaymentLinkRadioField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PaymentLinkRadioField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetName

`func (o *PaymentLinkRadioField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentLinkRadioField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentLinkRadioField) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *PaymentLinkRadioField) GetOptions() []EnumeratedFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PaymentLinkRadioField) GetOptionsOk() (*[]EnumeratedFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PaymentLinkRadioField) SetOptions(v []EnumeratedFieldOption)`

SetOptions sets Options field to given value.


### GetDescription

`func (o *PaymentLinkRadioField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentLinkRadioField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentLinkRadioField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentLinkRadioField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValues

`func (o *PaymentLinkRadioField) GetDefaultValues() []string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *PaymentLinkRadioField) GetDefaultValuesOk() (*[]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *PaymentLinkRadioField) SetDefaultValues(v []string)`

SetDefaultValues sets DefaultValues field to given value.


### GetDependentFields

`func (o *PaymentLinkRadioField) GetDependentFields() []DependentField`

GetDependentFields returns the DependentFields field if non-nil, zero value otherwise.

### GetDependentFieldsOk

`func (o *PaymentLinkRadioField) GetDependentFieldsOk() (*[]DependentField, bool)`

GetDependentFieldsOk returns a tuple with the DependentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFields

`func (o *PaymentLinkRadioField) SetDependentFields(v []DependentField)`

SetDependentFields sets DependentFields field to given value.


### GetLabel

`func (o *PaymentLinkRadioField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaymentLinkRadioField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaymentLinkRadioField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFieldType

`func (o *PaymentLinkRadioField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *PaymentLinkRadioField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *PaymentLinkRadioField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetRequired

`func (o *PaymentLinkRadioField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PaymentLinkRadioField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PaymentLinkRadioField) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


