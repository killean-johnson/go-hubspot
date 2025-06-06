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

// checks if the DatepickerField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatepickerField{}

// DatepickerField A form field used to select a date
type DatepickerField struct {
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
	// The prompt text showing when the field isn't filled in.
	Placeholder *string `json:"placeholder,omitempty"`
	// Determines how the field will be displayed and validated.
	FieldType string `json:"fieldType"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required"`
}

type _DatepickerField DatepickerField

// NewDatepickerField instantiates a new DatepickerField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatepickerField(objectTypeId string, hidden bool, name string, dependentFields []DependentField, label string, fieldType string, required bool) *DatepickerField {
	this := DatepickerField{}
	this.ObjectTypeId = objectTypeId
	this.Hidden = hidden
	this.Name = name
	this.DependentFields = dependentFields
	this.Label = label
	this.FieldType = fieldType
	this.Required = required
	return &this
}

// NewDatepickerFieldWithDefaults instantiates a new DatepickerField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatepickerFieldWithDefaults() *DatepickerField {
	this := DatepickerField{}
	var fieldType string = "datepicker"
	this.FieldType = fieldType
	return &this
}

// GetObjectTypeId returns the ObjectTypeId field value
func (o *DatepickerField) GetObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectTypeId
}

// GetObjectTypeIdOk returns a tuple with the ObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectTypeId, true
}

// SetObjectTypeId sets field value
func (o *DatepickerField) SetObjectTypeId(v string) {
	o.ObjectTypeId = v
}

// GetHidden returns the Hidden field value
func (o *DatepickerField) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *DatepickerField) SetHidden(v bool) {
	o.Hidden = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *DatepickerField) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *DatepickerField) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *DatepickerField) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetName returns the Name field value
func (o *DatepickerField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatepickerField) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatepickerField) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatepickerField) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DatepickerField) SetDescription(v string) {
	o.Description = &v
}

// GetDependentFields returns the DependentFields field value
func (o *DatepickerField) GetDependentFields() []DependentField {
	if o == nil {
		var ret []DependentField
		return ret
	}

	return o.DependentFields
}

// GetDependentFieldsOk returns a tuple with the DependentFields field value
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetDependentFieldsOk() ([]DependentField, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependentFields, true
}

// SetDependentFields sets field value
func (o *DatepickerField) SetDependentFields(v []DependentField) {
	o.DependentFields = v
}

// GetLabel returns the Label field value
func (o *DatepickerField) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *DatepickerField) SetLabel(v string) {
	o.Label = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *DatepickerField) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *DatepickerField) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *DatepickerField) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetFieldType returns the FieldType field value
func (o *DatepickerField) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *DatepickerField) SetFieldType(v string) {
	o.FieldType = v
}

// GetRequired returns the Required field value
func (o *DatepickerField) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *DatepickerField) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *DatepickerField) SetRequired(v bool) {
	o.Required = v
}

func (o DatepickerField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatepickerField) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	toSerialize["fieldType"] = o.FieldType
	toSerialize["required"] = o.Required
	return toSerialize, nil
}

func (o *DatepickerField) UnmarshalJSON(data []byte) (err error) {
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

	varDatepickerField := _DatepickerField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDatepickerField)

	if err != nil {
		return err
	}

	*o = DatepickerField(varDatepickerField)

	return err
}

type NullableDatepickerField struct {
	value *DatepickerField
	isSet bool
}

func (v NullableDatepickerField) Get() *DatepickerField {
	return v.value
}

func (v *NullableDatepickerField) Set(val *DatepickerField) {
	v.value = val
	v.isSet = true
}

func (v NullableDatepickerField) IsSet() bool {
	return v.isSet
}

func (v *NullableDatepickerField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatepickerField(val *DatepickerField) *NullableDatepickerField {
	return &NullableDatepickerField{value: val, isSet: true}
}

func (v NullableDatepickerField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatepickerField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


