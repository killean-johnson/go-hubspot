/*
Object Library

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package object_library

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ObjectTypeEnablementPublicResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectTypeEnablementPublicResponse{}

// ObjectTypeEnablementPublicResponse struct for ObjectTypeEnablementPublicResponse
type ObjectTypeEnablementPublicResponse struct {
	Enablement bool `json:"enablement"`
}

type _ObjectTypeEnablementPublicResponse ObjectTypeEnablementPublicResponse

// NewObjectTypeEnablementPublicResponse instantiates a new ObjectTypeEnablementPublicResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectTypeEnablementPublicResponse(enablement bool) *ObjectTypeEnablementPublicResponse {
	this := ObjectTypeEnablementPublicResponse{}
	this.Enablement = enablement
	return &this
}

// NewObjectTypeEnablementPublicResponseWithDefaults instantiates a new ObjectTypeEnablementPublicResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectTypeEnablementPublicResponseWithDefaults() *ObjectTypeEnablementPublicResponse {
	this := ObjectTypeEnablementPublicResponse{}
	return &this
}

// GetEnablement returns the Enablement field value
func (o *ObjectTypeEnablementPublicResponse) GetEnablement() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enablement
}

// GetEnablementOk returns a tuple with the Enablement field value
// and a boolean to check if the value has been set.
func (o *ObjectTypeEnablementPublicResponse) GetEnablementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enablement, true
}

// SetEnablement sets field value
func (o *ObjectTypeEnablementPublicResponse) SetEnablement(v bool) {
	o.Enablement = v
}

func (o ObjectTypeEnablementPublicResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectTypeEnablementPublicResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enablement"] = o.Enablement
	return toSerialize, nil
}

func (o *ObjectTypeEnablementPublicResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enablement",
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

	varObjectTypeEnablementPublicResponse := _ObjectTypeEnablementPublicResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectTypeEnablementPublicResponse)

	if err != nil {
		return err
	}

	*o = ObjectTypeEnablementPublicResponse(varObjectTypeEnablementPublicResponse)

	return err
}

type NullableObjectTypeEnablementPublicResponse struct {
	value *ObjectTypeEnablementPublicResponse
	isSet bool
}

func (v NullableObjectTypeEnablementPublicResponse) Get() *ObjectTypeEnablementPublicResponse {
	return v.value
}

func (v *NullableObjectTypeEnablementPublicResponse) Set(val *ObjectTypeEnablementPublicResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectTypeEnablementPublicResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectTypeEnablementPublicResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectTypeEnablementPublicResponse(val *ObjectTypeEnablementPublicResponse) *NullableObjectTypeEnablementPublicResponse {
	return &NullableObjectTypeEnablementPublicResponse{value: val, isSet: true}
}

func (v NullableObjectTypeEnablementPublicResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectTypeEnablementPublicResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


