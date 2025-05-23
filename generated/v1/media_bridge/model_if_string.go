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

// checks if the IfString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IfString{}

// IfString struct for IfString
type IfString struct {
	PropertyName *string `json:"propertyName,omitempty"`
	Inputs []OrInputsInner `json:"inputs,omitempty"`
	IfExpression OrInputsInner `json:"ifExpression"`
	Value *string `json:"value,omitempty"`
	ElseExpression *OrInputsInner `json:"elseExpression,omitempty"`
	Operator string `json:"operator"`
}

type _IfString IfString

// NewIfString instantiates a new IfString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIfString(ifExpression OrInputsInner, operator string) *IfString {
	this := IfString{}
	this.IfExpression = ifExpression
	this.Operator = operator
	return &this
}

// NewIfStringWithDefaults instantiates a new IfString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIfStringWithDefaults() *IfString {
	this := IfString{}
	var operator string = "IF_STRING"
	this.Operator = operator
	return &this
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *IfString) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IfString) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *IfString) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *IfString) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *IfString) GetInputs() []OrInputsInner {
	if o == nil || IsNil(o.Inputs) {
		var ret []OrInputsInner
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IfString) GetInputsOk() ([]OrInputsInner, bool) {
	if o == nil || IsNil(o.Inputs) {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *IfString) HasInputs() bool {
	if o != nil && !IsNil(o.Inputs) {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []OrInputsInner and assigns it to the Inputs field.
func (o *IfString) SetInputs(v []OrInputsInner) {
	o.Inputs = v
}

// GetIfExpression returns the IfExpression field value
func (o *IfString) GetIfExpression() OrInputsInner {
	if o == nil {
		var ret OrInputsInner
		return ret
	}

	return o.IfExpression
}

// GetIfExpressionOk returns a tuple with the IfExpression field value
// and a boolean to check if the value has been set.
func (o *IfString) GetIfExpressionOk() (*OrInputsInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IfExpression, true
}

// SetIfExpression sets field value
func (o *IfString) SetIfExpression(v OrInputsInner) {
	o.IfExpression = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IfString) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IfString) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IfString) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IfString) SetValue(v string) {
	o.Value = &v
}

// GetElseExpression returns the ElseExpression field value if set, zero value otherwise.
func (o *IfString) GetElseExpression() OrInputsInner {
	if o == nil || IsNil(o.ElseExpression) {
		var ret OrInputsInner
		return ret
	}
	return *o.ElseExpression
}

// GetElseExpressionOk returns a tuple with the ElseExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IfString) GetElseExpressionOk() (*OrInputsInner, bool) {
	if o == nil || IsNil(o.ElseExpression) {
		return nil, false
	}
	return o.ElseExpression, true
}

// HasElseExpression returns a boolean if a field has been set.
func (o *IfString) HasElseExpression() bool {
	if o != nil && !IsNil(o.ElseExpression) {
		return true
	}

	return false
}

// SetElseExpression gets a reference to the given OrInputsInner and assigns it to the ElseExpression field.
func (o *IfString) SetElseExpression(v OrInputsInner) {
	o.ElseExpression = &v
}

// GetOperator returns the Operator field value
func (o *IfString) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *IfString) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *IfString) SetOperator(v string) {
	o.Operator = v
}

func (o IfString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IfString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyName) {
		toSerialize["propertyName"] = o.PropertyName
	}
	if !IsNil(o.Inputs) {
		toSerialize["inputs"] = o.Inputs
	}
	toSerialize["ifExpression"] = o.IfExpression
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ElseExpression) {
		toSerialize["elseExpression"] = o.ElseExpression
	}
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *IfString) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ifExpression",
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

	varIfString := _IfString{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIfString)

	if err != nil {
		return err
	}

	*o = IfString(varIfString)

	return err
}

type NullableIfString struct {
	value *IfString
	isSet bool
}

func (v NullableIfString) Get() *IfString {
	return v.value
}

func (v *NullableIfString) Set(val *IfString) {
	v.value = val
	v.isSet = true
}

func (v NullableIfString) IsSet() bool {
	return v.isSet
}

func (v *NullableIfString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIfString(val *IfString) *NullableIfString {
	return &NullableIfString{value: val, isSet: true}
}

func (v NullableIfString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIfString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


