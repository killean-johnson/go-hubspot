/*
CMS Media Bridge

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package media_bridge

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the StringEquals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringEquals{}

// StringEquals struct for StringEquals
type StringEquals struct {
	PropertyName *string `json:"propertyName,omitempty"`
	Inputs []OrInputsInner `json:"inputs,omitempty"`
	Value *bool `json:"value,omitempty"`
	Operator string `json:"operator"`
}

type _StringEquals StringEquals

// NewStringEquals instantiates a new StringEquals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringEquals(operator string) *StringEquals {
	this := StringEquals{}
	this.Operator = operator
	return &this
}

// NewStringEqualsWithDefaults instantiates a new StringEquals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringEqualsWithDefaults() *StringEquals {
	this := StringEquals{}
	var operator string = "STRING_EQUALS"
	this.Operator = operator
	return &this
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *StringEquals) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringEquals) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *StringEquals) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *StringEquals) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *StringEquals) GetInputs() []OrInputsInner {
	if o == nil || IsNil(o.Inputs) {
		var ret []OrInputsInner
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringEquals) GetInputsOk() ([]OrInputsInner, bool) {
	if o == nil || IsNil(o.Inputs) {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *StringEquals) HasInputs() bool {
	if o != nil && !IsNil(o.Inputs) {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []OrInputsInner and assigns it to the Inputs field.
func (o *StringEquals) SetInputs(v []OrInputsInner) {
	o.Inputs = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StringEquals) GetValue() bool {
	if o == nil || IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringEquals) GetValueOk() (*bool, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StringEquals) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *StringEquals) SetValue(v bool) {
	o.Value = &v
}

// GetOperator returns the Operator field value
func (o *StringEquals) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *StringEquals) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *StringEquals) SetOperator(v string) {
	o.Operator = v
}

func (o StringEquals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringEquals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyName) {
		toSerialize["propertyName"] = o.PropertyName
	}
	if !IsNil(o.Inputs) {
		toSerialize["inputs"] = o.Inputs
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *StringEquals) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operator",
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

	varStringEquals := _StringEquals{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStringEquals)

	if err != nil {
		return err
	}

	*o = StringEquals(varStringEquals)

	return err
}

type NullableStringEquals struct {
	value *StringEquals
	isSet bool
}

func (v NullableStringEquals) Get() *StringEquals {
	return v.value
}

func (v *NullableStringEquals) Set(val *StringEquals) {
	v.value = val
	v.isSet = true
}

func (v NullableStringEquals) IsSet() bool {
	return v.isSet
}

func (v *NullableStringEquals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringEquals(val *StringEquals) *NullableStringEquals {
	return &NullableStringEquals{value: val, isSet: true}
}

func (v NullableStringEquals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringEquals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


