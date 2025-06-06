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

// checks if the SingleCheckboxField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleCheckboxField{}

// SingleCheckboxField A form field consisting of a single checkbox.
type SingleCheckboxField struct {
	// A unique ID for this field's CRM object type. For example a CONTACT field will have the object type ID 0-1.
	ObjectTypeId string `json:"objectTypeId"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form, but can be used to pass a value to a property without requiring the customer to fill it in.
	Hidden bool `json:"hidden"`
	// The value filled in by default. This value will be submitted unless the customer modifies it.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// The identifier of the field. In combination with the object type ID, it must be unique.
	Name string `json:"name"`
	// Additional text helping the customer to complete the field.
	Description *string `json:"description,omitempty"`
	// A list of other fields to make visible based on the value filled in for this field.
	DependentFields []DependentField `json:"dependentFields"`
	// The main label for the form field.
	Label string `json:"label"`
	// Determines how the field will be displayed and validated.
	FieldType string `json:"fieldType"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required"`
}

type _SingleCheckboxField SingleCheckboxField

// NewSingleCheckboxField instantiates a new SingleCheckboxField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleCheckboxField(objectTypeId string, hidden bool, name string, dependentFields []DependentField, label string, fieldType string, required bool) *SingleCheckboxField {
	this := SingleCheckboxField{}
	this.ObjectTypeId = objectTypeId
	this.Hidden = hidden
	this.Name = name
	this.DependentFields = dependentFields
	this.Label = label
	this.FieldType = fieldType
	this.Required = required
	return &this
}

// NewSingleCheckboxFieldWithDefaults instantiates a new SingleCheckboxField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleCheckboxFieldWithDefaults() *SingleCheckboxField {
	this := SingleCheckboxField{}
	var fieldType string = "single_checkbox"
	this.FieldType = fieldType
	return &this
}

// GetObjectTypeId returns the ObjectTypeId field value
func (o *SingleCheckboxField) GetObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectTypeId
}

// GetObjectTypeIdOk returns a tuple with the ObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *SingleCheckboxField) GetObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectTypeId, true
}

// SetObjectTypeId sets field value
func (o *SingleCheckboxField) SetObjectTypeId(v string) {
	o.ObjectTypeId = v
}

// GetHidden returns the Hidden field value
func (o *SingleCheckboxField) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *SingleCheckboxField) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *SingleCheckboxField) SetHidden(v bool) {
	o.Hidden = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *SingleCheckboxField) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleCheckboxField) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *SingleCheckboxField) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *SingleCheckboxField) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetName returns the Name field value
func (o *SingleCheckboxField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SingleCheckboxField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SingleCheckboxField) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SingleCheckboxField) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleCheckboxField) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SingleCheckboxField) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SingleCheckboxField) SetDescription(v string) {
	o.Description = &v
}

// GetDependentFields returns the DependentFields field value
func (o *SingleCheckboxField) GetDependentFields() []DependentField {
	if o == nil {
		var ret []DependentField
		return ret
	}

	return o.DependentFields
}

// GetDependentFieldsOk returns a tuple with the DependentFields field value
// and a boolean to check if the value has been set.
func (o *SingleCheckboxField) GetDependentFieldsOk() ([]DependentField, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependentFields, true
}

// SetDependentFields sets field value
func (o *SingleCheckboxField) SetDependentFields(v []DependentField) {
	o.DependentFields = v
}

// GetLabel returns the Label field value
func (o *SingleCheckboxField) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SingleCheckboxField) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SingleCheckboxField) SetLabel(v string) {
	o.Label = v
}

// GetFieldType returns the FieldType field value
func (o *SingleCheckboxField) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *SingleCheckboxField) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *SingleCheckboxField) SetFieldType(v string) {
	o.FieldType = v
}

// GetRequired returns the Required field value
func (o *SingleCheckboxField) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *SingleCheckboxField) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *SingleCheckboxField) SetRequired(v bool) {
	o.Required = v
}

func (o SingleCheckboxField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleCheckboxField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objectTypeId"] = o.ObjectTypeId
	toSerialize["hidden"] = o.Hidden
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["dependentFields"] = o.DependentFields
	toSerialize["label"] = o.Label
	toSerialize["fieldType"] = o.FieldType
	toSerialize["required"] = o.Required
	return toSerialize, nil
}

func (o *SingleCheckboxField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectTypeId",
		"hidden",
		"name",
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

	varSingleCheckboxField := _SingleCheckboxField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSingleCheckboxField)

	if err != nil {
		return err
	}

	*o = SingleCheckboxField(varSingleCheckboxField)

	return err
}

type NullableSingleCheckboxField struct {
	value *SingleCheckboxField
	isSet bool
}

func (v NullableSingleCheckboxField) Get() *SingleCheckboxField {
	return v.value
}

func (v *NullableSingleCheckboxField) Set(val *SingleCheckboxField) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleCheckboxField) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleCheckboxField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleCheckboxField(val *SingleCheckboxField) *NullableSingleCheckboxField {
	return &NullableSingleCheckboxField{value: val, isSet: true}
}

func (v NullableSingleCheckboxField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleCheckboxField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


