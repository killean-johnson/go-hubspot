/*
Forms

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package forms

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PaymentLinkRadioField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentLinkRadioField{}

// PaymentLinkRadioField struct for PaymentLinkRadioField
type PaymentLinkRadioField struct {
	ObjectTypeId string `json:"objectTypeId"`
	Hidden bool `json:"hidden"`
	Name string `json:"name"`
	Options []EnumeratedFieldOption `json:"options"`
	Description *string `json:"description,omitempty"`
	DefaultValues []string `json:"defaultValues"`
	DependentFields []DependentField `json:"dependentFields"`
	Label string `json:"label"`
	FieldType string `json:"fieldType"`
	Required bool `json:"required"`
}

type _PaymentLinkRadioField PaymentLinkRadioField

// NewPaymentLinkRadioField instantiates a new PaymentLinkRadioField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinkRadioField(objectTypeId string, hidden bool, name string, options []EnumeratedFieldOption, defaultValues []string, dependentFields []DependentField, label string, fieldType string, required bool) *PaymentLinkRadioField {
	this := PaymentLinkRadioField{}
	this.ObjectTypeId = objectTypeId
	this.Hidden = hidden
	this.Name = name
	this.Options = options
	this.DefaultValues = defaultValues
	this.DependentFields = dependentFields
	this.Label = label
	this.FieldType = fieldType
	this.Required = required
	return &this
}

// NewPaymentLinkRadioFieldWithDefaults instantiates a new PaymentLinkRadioField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinkRadioFieldWithDefaults() *PaymentLinkRadioField {
	this := PaymentLinkRadioField{}
	var fieldType string = "payment_link_radio"
	this.FieldType = fieldType
	return &this
}

// GetObjectTypeId returns the ObjectTypeId field value
func (o *PaymentLinkRadioField) GetObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectTypeId
}

// GetObjectTypeIdOk returns a tuple with the ObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectTypeId, true
}

// SetObjectTypeId sets field value
func (o *PaymentLinkRadioField) SetObjectTypeId(v string) {
	o.ObjectTypeId = v
}

// GetHidden returns the Hidden field value
func (o *PaymentLinkRadioField) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *PaymentLinkRadioField) SetHidden(v bool) {
	o.Hidden = v
}

// GetName returns the Name field value
func (o *PaymentLinkRadioField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentLinkRadioField) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *PaymentLinkRadioField) GetOptions() []EnumeratedFieldOption {
	if o == nil {
		var ret []EnumeratedFieldOption
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetOptionsOk() ([]EnumeratedFieldOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *PaymentLinkRadioField) SetOptions(v []EnumeratedFieldOption) {
	o.Options = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentLinkRadioField) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentLinkRadioField) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentLinkRadioField) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultValues returns the DefaultValues field value
func (o *PaymentLinkRadioField) GetDefaultValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetDefaultValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValues, true
}

// SetDefaultValues sets field value
func (o *PaymentLinkRadioField) SetDefaultValues(v []string) {
	o.DefaultValues = v
}

// GetDependentFields returns the DependentFields field value
func (o *PaymentLinkRadioField) GetDependentFields() []DependentField {
	if o == nil {
		var ret []DependentField
		return ret
	}

	return o.DependentFields
}

// GetDependentFieldsOk returns a tuple with the DependentFields field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetDependentFieldsOk() ([]DependentField, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependentFields, true
}

// SetDependentFields sets field value
func (o *PaymentLinkRadioField) SetDependentFields(v []DependentField) {
	o.DependentFields = v
}

// GetLabel returns the Label field value
func (o *PaymentLinkRadioField) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PaymentLinkRadioField) SetLabel(v string) {
	o.Label = v
}

// GetFieldType returns the FieldType field value
func (o *PaymentLinkRadioField) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *PaymentLinkRadioField) SetFieldType(v string) {
	o.FieldType = v
}

// GetRequired returns the Required field value
func (o *PaymentLinkRadioField) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRadioField) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *PaymentLinkRadioField) SetRequired(v bool) {
	o.Required = v
}

func (o PaymentLinkRadioField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentLinkRadioField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objectTypeId"] = o.ObjectTypeId
	toSerialize["hidden"] = o.Hidden
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["defaultValues"] = o.DefaultValues
	toSerialize["dependentFields"] = o.DependentFields
	toSerialize["label"] = o.Label
	toSerialize["fieldType"] = o.FieldType
	toSerialize["required"] = o.Required
	return toSerialize, nil
}

func (o *PaymentLinkRadioField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectTypeId",
		"hidden",
		"name",
		"options",
		"defaultValues",
		"dependentFields",
		"label",
		"fieldType",
		"required",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaymentLinkRadioField := _PaymentLinkRadioField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentLinkRadioField)

	if err != nil {
		return err
	}

	*o = PaymentLinkRadioField(varPaymentLinkRadioField)

	return err
}

type NullablePaymentLinkRadioField struct {
	value *PaymentLinkRadioField
	isSet bool
}

func (v NullablePaymentLinkRadioField) Get() *PaymentLinkRadioField {
	return v.value
}

func (v *NullablePaymentLinkRadioField) Set(val *PaymentLinkRadioField) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinkRadioField) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinkRadioField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinkRadioField(val *PaymentLinkRadioField) *NullablePaymentLinkRadioField {
	return &NullablePaymentLinkRadioField{value: val, isSet: true}
}

func (v NullablePaymentLinkRadioField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinkRadioField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


