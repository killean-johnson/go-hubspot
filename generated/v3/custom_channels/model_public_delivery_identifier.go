/*
Conversations Custom Channels

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package custom_channels

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicDeliveryIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicDeliveryIdentifier{}

// PublicDeliveryIdentifier struct for PublicDeliveryIdentifier
type PublicDeliveryIdentifier struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

type _PublicDeliveryIdentifier PublicDeliveryIdentifier

// NewPublicDeliveryIdentifier instantiates a new PublicDeliveryIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicDeliveryIdentifier(type_ string, value string) *PublicDeliveryIdentifier {
	this := PublicDeliveryIdentifier{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewPublicDeliveryIdentifierWithDefaults instantiates a new PublicDeliveryIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicDeliveryIdentifierWithDefaults() *PublicDeliveryIdentifier {
	this := PublicDeliveryIdentifier{}
	return &this
}

// GetType returns the Type field value
func (o *PublicDeliveryIdentifier) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PublicDeliveryIdentifier) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PublicDeliveryIdentifier) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *PublicDeliveryIdentifier) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PublicDeliveryIdentifier) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PublicDeliveryIdentifier) SetValue(v string) {
	o.Value = v
}

func (o PublicDeliveryIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicDeliveryIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *PublicDeliveryIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varPublicDeliveryIdentifier := _PublicDeliveryIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicDeliveryIdentifier)

	if err != nil {
		return err
	}

	*o = PublicDeliveryIdentifier(varPublicDeliveryIdentifier)

	return err
}

type NullablePublicDeliveryIdentifier struct {
	value *PublicDeliveryIdentifier
	isSet bool
}

func (v NullablePublicDeliveryIdentifier) Get() *PublicDeliveryIdentifier {
	return v.value
}

func (v *NullablePublicDeliveryIdentifier) Set(val *PublicDeliveryIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicDeliveryIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicDeliveryIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicDeliveryIdentifier(val *PublicDeliveryIdentifier) *NullablePublicDeliveryIdentifier {
	return &NullablePublicDeliveryIdentifier{value: val, isSet: true}
}

func (v NullablePublicDeliveryIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicDeliveryIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


