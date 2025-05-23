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

// checks if the Or type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Or{}

// Or struct for Or
type Or struct {
	PropertyName *string `json:"propertyName,omitempty"`
	Inputs []OrInputsInner `json:"inputs,omitempty"`
	EnclosedInParentheses bool `json:"enclosedInParentheses"`
	Value *bool `json:"value,omitempty"`
	Operator string `json:"operator"`
}

type _Or Or

// NewOr instantiates a new Or object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOr(enclosedInParentheses bool, operator string) *Or {
	this := Or{}
	this.EnclosedInParentheses = enclosedInParentheses
	this.Operator = operator
	return &this
}

// NewOrWithDefaults instantiates a new Or object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrWithDefaults() *Or {
	this := Or{}
	var operator string = "OR"
	this.Operator = operator
	return &this
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *Or) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Or) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *Or) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *Or) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *Or) GetInputs() []OrInputsInner {
	if o == nil || IsNil(o.Inputs) {
		var ret []OrInputsInner
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Or) GetInputsOk() ([]OrInputsInner, bool) {
	if o == nil || IsNil(o.Inputs) {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *Or) HasInputs() bool {
	if o != nil && !IsNil(o.Inputs) {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []OrInputsInner and assigns it to the Inputs field.
func (o *Or) SetInputs(v []OrInputsInner) {
	o.Inputs = v
}

// GetEnclosedInParentheses returns the EnclosedInParentheses field value
func (o *Or) GetEnclosedInParentheses() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnclosedInParentheses
}

// GetEnclosedInParenthesesOk returns a tuple with the EnclosedInParentheses field value
// and a boolean to check if the value has been set.
func (o *Or) GetEnclosedInParenthesesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnclosedInParentheses, true
}

// SetEnclosedInParentheses sets field value
func (o *Or) SetEnclosedInParentheses(v bool) {
	o.EnclosedInParentheses = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Or) GetValue() bool {
	if o == nil || IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Or) GetValueOk() (*bool, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Or) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *Or) SetValue(v bool) {
	o.Value = &v
}

// GetOperator returns the Operator field value
func (o *Or) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *Or) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *Or) SetOperator(v string) {
	o.Operator = v
}

func (o Or) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Or) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyName) {
		toSerialize["propertyName"] = o.PropertyName
	}
	if !IsNil(o.Inputs) {
		toSerialize["inputs"] = o.Inputs
	}
	toSerialize["enclosedInParentheses"] = o.EnclosedInParentheses
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *Or) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enclosedInParentheses",
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

	varOr := _Or{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOr)

	if err != nil {
		return err
	}

	*o = Or(varOr)

	return err
}

type NullableOr struct {
	value *Or
	isSet bool
}

func (v NullableOr) Get() *Or {
	return v.value
}

func (v *NullableOr) Set(val *Or) {
	v.value = val
	v.isSet = true
}

func (v NullableOr) IsSet() bool {
	return v.isSet
}

func (v *NullableOr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOr(val *Or) *NullableOr {
	return &NullableOr{value: val, isSet: true}
}

func (v NullableOr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


