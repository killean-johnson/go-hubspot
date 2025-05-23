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

// checks if the PropertyName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyName{}

// PropertyName struct for PropertyName
type PropertyName struct {
	Name string `json:"name"`
}

type _PropertyName PropertyName

// NewPropertyName instantiates a new PropertyName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyName(name string) *PropertyName {
	this := PropertyName{}
	this.Name = name
	return &this
}

// NewPropertyNameWithDefaults instantiates a new PropertyName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyNameWithDefaults() *PropertyName {
	this := PropertyName{}
	return &this
}

// GetName returns the Name field value
func (o *PropertyName) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PropertyName) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PropertyName) SetName(v string) {
	o.Name = v
}

func (o PropertyName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *PropertyName) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varPropertyName := _PropertyName{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPropertyName)

	if err != nil {
		return err
	}

	*o = PropertyName(varPropertyName)

	return err
}

type NullablePropertyName struct {
	value *PropertyName
	isSet bool
}

func (v NullablePropertyName) Get() *PropertyName {
	return v.value
}

func (v *NullablePropertyName) Set(val *PropertyName) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyName) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyName(val *PropertyName) *NullablePropertyName {
	return &NullablePropertyName{value: val, isSet: true}
}

func (v NullablePropertyName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


