/*
Conversations Inbox & Messages

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conversations_inbox_&amp;_messages

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the QuickReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickReply{}

// QuickReply struct for QuickReply
type QuickReply struct {
	ValueType string `json:"valueType"`
	Label *string `json:"label,omitempty"`
	Value string `json:"value"`
}

type _QuickReply QuickReply

// NewQuickReply instantiates a new QuickReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickReply(valueType string, value string) *QuickReply {
	this := QuickReply{}
	this.ValueType = valueType
	this.Value = value
	return &this
}

// NewQuickReplyWithDefaults instantiates a new QuickReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickReplyWithDefaults() *QuickReply {
	this := QuickReply{}
	return &this
}

// GetValueType returns the ValueType field value
func (o *QuickReply) GetValueType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *QuickReply) GetValueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *QuickReply) SetValueType(v string) {
	o.ValueType = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *QuickReply) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickReply) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *QuickReply) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *QuickReply) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value
func (o *QuickReply) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *QuickReply) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *QuickReply) SetValue(v string) {
	o.Value = v
}

func (o QuickReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valueType"] = o.ValueType
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *QuickReply) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"valueType",
		"value",
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

	varQuickReply := _QuickReply{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuickReply)

	if err != nil {
		return err
	}

	*o = QuickReply(varQuickReply)

	return err
}

type NullableQuickReply struct {
	value *QuickReply
	isSet bool
}

func (v NullableQuickReply) Get() *QuickReply {
	return v.value
}

func (v *NullableQuickReply) Set(val *QuickReply) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickReply) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickReply(val *QuickReply) *NullableQuickReply {
	return &NullableQuickReply{value: val, isSet: true}
}

func (v NullableQuickReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


