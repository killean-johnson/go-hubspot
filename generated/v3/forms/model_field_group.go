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

// checks if the FieldGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldGroup{}

// FieldGroup A collection of up to three form fields usually displayed in a row.
type FieldGroup struct {
	// 
	GroupType string `json:"groupType"`
	// The type of rich text included. The default value is text.
	RichTextType string `json:"richTextType"`
	// A block of rich text or an image. Those can be used to add extra information for the customers filling in the form. If the field group includes fields, the rich text will be displayed before the fields.
	RichText *string `json:"richText,omitempty"`
	// The form fields included in the group
	Fields []DependentFieldDependentField `json:"fields"`
}

type _FieldGroup FieldGroup

// NewFieldGroup instantiates a new FieldGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldGroup(groupType string, richTextType string, fields []DependentFieldDependentField) *FieldGroup {
	this := FieldGroup{}
	this.GroupType = groupType
	this.RichTextType = richTextType
	this.Fields = fields
	return &this
}

// NewFieldGroupWithDefaults instantiates a new FieldGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldGroupWithDefaults() *FieldGroup {
	this := FieldGroup{}
	return &this
}

// GetGroupType returns the GroupType field value
func (o *FieldGroup) GetGroupType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value
// and a boolean to check if the value has been set.
func (o *FieldGroup) GetGroupTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupType, true
}

// SetGroupType sets field value
func (o *FieldGroup) SetGroupType(v string) {
	o.GroupType = v
}

// GetRichTextType returns the RichTextType field value
func (o *FieldGroup) GetRichTextType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RichTextType
}

// GetRichTextTypeOk returns a tuple with the RichTextType field value
// and a boolean to check if the value has been set.
func (o *FieldGroup) GetRichTextTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RichTextType, true
}

// SetRichTextType sets field value
func (o *FieldGroup) SetRichTextType(v string) {
	o.RichTextType = v
}

// GetRichText returns the RichText field value if set, zero value otherwise.
func (o *FieldGroup) GetRichText() string {
	if o == nil || IsNil(o.RichText) {
		var ret string
		return ret
	}
	return *o.RichText
}

// GetRichTextOk returns a tuple with the RichText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldGroup) GetRichTextOk() (*string, bool) {
	if o == nil || IsNil(o.RichText) {
		return nil, false
	}
	return o.RichText, true
}

// HasRichText returns a boolean if a field has been set.
func (o *FieldGroup) HasRichText() bool {
	if o != nil && !IsNil(o.RichText) {
		return true
	}

	return false
}

// SetRichText gets a reference to the given string and assigns it to the RichText field.
func (o *FieldGroup) SetRichText(v string) {
	o.RichText = &v
}

// GetFields returns the Fields field value
func (o *FieldGroup) GetFields() []DependentFieldDependentField {
	if o == nil {
		var ret []DependentFieldDependentField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *FieldGroup) GetFieldsOk() ([]DependentFieldDependentField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *FieldGroup) SetFields(v []DependentFieldDependentField) {
	o.Fields = v
}

func (o FieldGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupType"] = o.GroupType
	toSerialize["richTextType"] = o.RichTextType
	if !IsNil(o.RichText) {
		toSerialize["richText"] = o.RichText
	}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *FieldGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupType",
		"richTextType",
		"fields",
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

	varFieldGroup := _FieldGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFieldGroup)

	if err != nil {
		return err
	}

	*o = FieldGroup(varFieldGroup)

	return err
}

type NullableFieldGroup struct {
	value *FieldGroup
	isSet bool
}

func (v NullableFieldGroup) Get() *FieldGroup {
	return v.value
}

func (v *NullableFieldGroup) Set(val *FieldGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldGroup(val *FieldGroup) *NullableFieldGroup {
	return &NullableFieldGroup{value: val, isSet: true}
}

func (v NullableFieldGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


