/*
CRM Deal Splits

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deal_splits

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicDealSplitInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicDealSplitInput{}

// PublicDealSplitInput struct for PublicDealSplitInput
type PublicDealSplitInput struct {
	Percentage float32 `json:"percentage"`
	OwnerId int32 `json:"ownerId"`
}

type _PublicDealSplitInput PublicDealSplitInput

// NewPublicDealSplitInput instantiates a new PublicDealSplitInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicDealSplitInput(percentage float32, ownerId int32) *PublicDealSplitInput {
	this := PublicDealSplitInput{}
	this.Percentage = percentage
	this.OwnerId = ownerId
	return &this
}

// NewPublicDealSplitInputWithDefaults instantiates a new PublicDealSplitInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicDealSplitInputWithDefaults() *PublicDealSplitInput {
	this := PublicDealSplitInput{}
	return &this
}

// GetPercentage returns the Percentage field value
func (o *PublicDealSplitInput) GetPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *PublicDealSplitInput) GetPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value
func (o *PublicDealSplitInput) SetPercentage(v float32) {
	o.Percentage = v
}

// GetOwnerId returns the OwnerId field value
func (o *PublicDealSplitInput) GetOwnerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *PublicDealSplitInput) GetOwnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *PublicDealSplitInput) SetOwnerId(v int32) {
	o.OwnerId = v
}

func (o PublicDealSplitInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicDealSplitInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["percentage"] = o.Percentage
	toSerialize["ownerId"] = o.OwnerId
	return toSerialize, nil
}

func (o *PublicDealSplitInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"percentage",
		"ownerId",
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

	varPublicDealSplitInput := _PublicDealSplitInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicDealSplitInput)

	if err != nil {
		return err
	}

	*o = PublicDealSplitInput(varPublicDealSplitInput)

	return err
}

type NullablePublicDealSplitInput struct {
	value *PublicDealSplitInput
	isSet bool
}

func (v NullablePublicDealSplitInput) Get() *PublicDealSplitInput {
	return v.value
}

func (v *NullablePublicDealSplitInput) Set(val *PublicDealSplitInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicDealSplitInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicDealSplitInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicDealSplitInput(val *PublicDealSplitInput) *NullablePublicDealSplitInput {
	return &NullablePublicDealSplitInput{value: val, isSet: true}
}

func (v NullablePublicDealSplitInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicDealSplitInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


