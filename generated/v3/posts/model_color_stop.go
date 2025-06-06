/*
Posts

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posts

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ColorStop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ColorStop{}

// ColorStop struct for ColorStop
type ColorStop struct {
	Color RGBAColor `json:"color"`
}

type _ColorStop ColorStop

// NewColorStop instantiates a new ColorStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColorStop(color RGBAColor) *ColorStop {
	this := ColorStop{}
	this.Color = color
	return &this
}

// NewColorStopWithDefaults instantiates a new ColorStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColorStopWithDefaults() *ColorStop {
	this := ColorStop{}
	return &this
}

// GetColor returns the Color field value
func (o *ColorStop) GetColor() RGBAColor {
	if o == nil {
		var ret RGBAColor
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *ColorStop) GetColorOk() (*RGBAColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *ColorStop) SetColor(v RGBAColor) {
	o.Color = v
}

func (o ColorStop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ColorStop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["color"] = o.Color
	return toSerialize, nil
}

func (o *ColorStop) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"color",
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

	varColorStop := _ColorStop{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varColorStop)

	if err != nil {
		return err
	}

	*o = ColorStop(varColorStop)

	return err
}

type NullableColorStop struct {
	value *ColorStop
	isSet bool
}

func (v NullableColorStop) Get() *ColorStop {
	return v.value
}

func (v *NullableColorStop) Set(val *ColorStop) {
	v.value = val
	v.isSet = true
}

func (v NullableColorStop) IsSet() bool {
	return v.isSet
}

func (v *NullableColorStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColorStop(val *ColorStop) *NullableColorStop {
	return &NullableColorStop{value: val, isSet: true}
}

func (v NullableColorStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColorStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


