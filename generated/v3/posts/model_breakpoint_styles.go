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

// checks if the BreakpointStyles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BreakpointStyles{}

// BreakpointStyles struct for BreakpointStyles
type BreakpointStyles struct {
	Padding map[string]interface{} `json:"padding"`
	Margin map[string]interface{} `json:"margin"`
	Hidden bool `json:"hidden"`
}

type _BreakpointStyles BreakpointStyles

// NewBreakpointStyles instantiates a new BreakpointStyles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBreakpointStyles(padding map[string]interface{}, margin map[string]interface{}, hidden bool) *BreakpointStyles {
	this := BreakpointStyles{}
	this.Padding = padding
	this.Margin = margin
	this.Hidden = hidden
	return &this
}

// NewBreakpointStylesWithDefaults instantiates a new BreakpointStyles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBreakpointStylesWithDefaults() *BreakpointStyles {
	this := BreakpointStyles{}
	return &this
}

// GetPadding returns the Padding field value
func (o *BreakpointStyles) GetPadding() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Padding
}

// GetPaddingOk returns a tuple with the Padding field value
// and a boolean to check if the value has been set.
func (o *BreakpointStyles) GetPaddingOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Padding, true
}

// SetPadding sets field value
func (o *BreakpointStyles) SetPadding(v map[string]interface{}) {
	o.Padding = v
}

// GetMargin returns the Margin field value
func (o *BreakpointStyles) GetMargin() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Margin
}

// GetMarginOk returns a tuple with the Margin field value
// and a boolean to check if the value has been set.
func (o *BreakpointStyles) GetMarginOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Margin, true
}

// SetMargin sets field value
func (o *BreakpointStyles) SetMargin(v map[string]interface{}) {
	o.Margin = v
}

// GetHidden returns the Hidden field value
func (o *BreakpointStyles) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *BreakpointStyles) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *BreakpointStyles) SetHidden(v bool) {
	o.Hidden = v
}

func (o BreakpointStyles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BreakpointStyles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["padding"] = o.Padding
	toSerialize["margin"] = o.Margin
	toSerialize["hidden"] = o.Hidden
	return toSerialize, nil
}

func (o *BreakpointStyles) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"padding",
		"margin",
		"hidden",
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

	varBreakpointStyles := _BreakpointStyles{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBreakpointStyles)

	if err != nil {
		return err
	}

	*o = BreakpointStyles(varBreakpointStyles)

	return err
}

type NullableBreakpointStyles struct {
	value *BreakpointStyles
	isSet bool
}

func (v NullableBreakpointStyles) Get() *BreakpointStyles {
	return v.value
}

func (v *NullableBreakpointStyles) Set(val *BreakpointStyles) {
	v.value = val
	v.isSet = true
}

func (v NullableBreakpointStyles) IsSet() bool {
	return v.isSet
}

func (v *NullableBreakpointStyles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBreakpointStyles(val *BreakpointStyles) *NullableBreakpointStyles {
	return &NullableBreakpointStyles{value: val, isSet: true}
}

func (v NullableBreakpointStyles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBreakpointStyles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


