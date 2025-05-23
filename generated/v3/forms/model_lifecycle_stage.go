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

// checks if the LifecycleStage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleStage{}

// LifecycleStage struct for LifecycleStage
type LifecycleStage struct {
	// The objectTypeId for both contact and company
	ObjectTypeId string `json:"objectTypeId"`
	// The internal name of the contact's lifecycle stage set when submitting a form
	Value string `json:"value"`
}

type _LifecycleStage LifecycleStage

// NewLifecycleStage instantiates a new LifecycleStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleStage(objectTypeId string, value string) *LifecycleStage {
	this := LifecycleStage{}
	this.ObjectTypeId = objectTypeId
	this.Value = value
	return &this
}

// NewLifecycleStageWithDefaults instantiates a new LifecycleStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleStageWithDefaults() *LifecycleStage {
	this := LifecycleStage{}
	return &this
}

// GetObjectTypeId returns the ObjectTypeId field value
func (o *LifecycleStage) GetObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectTypeId
}

// GetObjectTypeIdOk returns a tuple with the ObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *LifecycleStage) GetObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectTypeId, true
}

// SetObjectTypeId sets field value
func (o *LifecycleStage) SetObjectTypeId(v string) {
	o.ObjectTypeId = v
}

// GetValue returns the Value field value
func (o *LifecycleStage) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LifecycleStage) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *LifecycleStage) SetValue(v string) {
	o.Value = v
}

func (o LifecycleStage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleStage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objectTypeId"] = o.ObjectTypeId
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *LifecycleStage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectTypeId",
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

	varLifecycleStage := _LifecycleStage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLifecycleStage)

	if err != nil {
		return err
	}

	*o = LifecycleStage(varLifecycleStage)

	return err
}

type NullableLifecycleStage struct {
	value *LifecycleStage
	isSet bool
}

func (v NullableLifecycleStage) Get() *LifecycleStage {
	return v.value
}

func (v *NullableLifecycleStage) Set(val *LifecycleStage) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleStage) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleStage(val *LifecycleStage) *NullableLifecycleStage {
	return &NullableLifecycleStage{value: val, isSet: true}
}

func (v NullableLifecycleStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


