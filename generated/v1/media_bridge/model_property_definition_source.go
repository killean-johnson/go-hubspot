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

// checks if the PropertyDefinitionSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyDefinitionSource{}

// PropertyDefinitionSource struct for PropertyDefinitionSource
type PropertyDefinitionSource struct {
	Name *string `json:"name,omitempty"`
	Type string `json:"type"`
}

type _PropertyDefinitionSource PropertyDefinitionSource

// NewPropertyDefinitionSource instantiates a new PropertyDefinitionSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyDefinitionSource(type_ string) *PropertyDefinitionSource {
	this := PropertyDefinitionSource{}
	this.Type = type_
	return &this
}

// NewPropertyDefinitionSourceWithDefaults instantiates a new PropertyDefinitionSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyDefinitionSourceWithDefaults() *PropertyDefinitionSource {
	this := PropertyDefinitionSource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PropertyDefinitionSource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinitionSource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PropertyDefinitionSource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PropertyDefinitionSource) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *PropertyDefinitionSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PropertyDefinitionSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PropertyDefinitionSource) SetType(v string) {
	o.Type = v
}

func (o PropertyDefinitionSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyDefinitionSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *PropertyDefinitionSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varPropertyDefinitionSource := _PropertyDefinitionSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPropertyDefinitionSource)

	if err != nil {
		return err
	}

	*o = PropertyDefinitionSource(varPropertyDefinitionSource)

	return err
}

type NullablePropertyDefinitionSource struct {
	value *PropertyDefinitionSource
	isSet bool
}

func (v NullablePropertyDefinitionSource) Get() *PropertyDefinitionSource {
	return v.value
}

func (v *NullablePropertyDefinitionSource) Set(val *PropertyDefinitionSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyDefinitionSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyDefinitionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyDefinitionSource(val *PropertyDefinitionSource) *NullablePropertyDefinitionSource {
	return &NullablePropertyDefinitionSource{value: val, isSet: true}
}

func (v NullablePropertyDefinitionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyDefinitionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


