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

// checks if the ConstantBoolean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstantBoolean{}

// ConstantBoolean struct for ConstantBoolean
type ConstantBoolean struct {
	PropertyName *string `json:"propertyName,omitempty"`
	Inputs []OrInputsInner `json:"inputs,omitempty"`
	Value *bool `json:"value,omitempty"`
	Operator string `json:"operator"`
}

type _ConstantBoolean ConstantBoolean

// NewConstantBoolean instantiates a new ConstantBoolean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstantBoolean(operator string) *ConstantBoolean {
	this := ConstantBoolean{}
	this.Operator = operator
	return &this
}

// NewConstantBooleanWithDefaults instantiates a new ConstantBoolean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstantBooleanWithDefaults() *ConstantBoolean {
	this := ConstantBoolean{}
	var operator string = "CONSTANT_BOOLEAN"
	this.Operator = operator
	return &this
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *ConstantBoolean) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstantBoolean) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *ConstantBoolean) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *ConstantBoolean) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *ConstantBoolean) GetInputs() []OrInputsInner {
	if o == nil || IsNil(o.Inputs) {
		var ret []OrInputsInner
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstantBoolean) GetInputsOk() ([]OrInputsInner, bool) {
	if o == nil || IsNil(o.Inputs) {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *ConstantBoolean) HasInputs() bool {
	if o != nil && !IsNil(o.Inputs) {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []OrInputsInner and assigns it to the Inputs field.
func (o *ConstantBoolean) SetInputs(v []OrInputsInner) {
	o.Inputs = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConstantBoolean) GetValue() bool {
	if o == nil || IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstantBoolean) GetValueOk() (*bool, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConstantBoolean) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *ConstantBoolean) SetValue(v bool) {
	o.Value = &v
}

// GetOperator returns the Operator field value
func (o *ConstantBoolean) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ConstantBoolean) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *ConstantBoolean) SetOperator(v string) {
	o.Operator = v
}

func (o ConstantBoolean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstantBoolean) ToMap() (map[string]interface{}, error) {
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

func (o *ConstantBoolean) UnmarshalJSON(data []byte) (err error) {
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

	varConstantBoolean := _ConstantBoolean{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConstantBoolean)

	if err != nil {
		return err
	}

	*o = ConstantBoolean(varConstantBoolean)

	return err
}

type NullableConstantBoolean struct {
	value *ConstantBoolean
	isSet bool
}

func (v NullableConstantBoolean) Get() *ConstantBoolean {
	return v.value
}

func (v *NullableConstantBoolean) Set(val *ConstantBoolean) {
	v.value = val
	v.isSet = true
}

func (v NullableConstantBoolean) IsSet() bool {
	return v.isSet
}

func (v *NullableConstantBoolean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstantBoolean(val *ConstantBoolean) *NullableConstantBoolean {
	return &NullableConstantBoolean{value: val, isSet: true}
}

func (v NullableConstantBoolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstantBoolean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


