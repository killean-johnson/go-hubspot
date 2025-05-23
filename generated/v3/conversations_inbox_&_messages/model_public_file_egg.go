/*
Conversations Inbox & Messages

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conversations_inbox_&amp;_messages

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicFileEgg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicFileEgg{}

// PublicFileEgg struct for PublicFileEgg
type PublicFileEgg struct {
	Type string `json:"type"`
	FileId string `json:"fileId"`
}

type _PublicFileEgg PublicFileEgg

// NewPublicFileEgg instantiates a new PublicFileEgg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicFileEgg(type_ string, fileId string) *PublicFileEgg {
	this := PublicFileEgg{}
	this.Type = type_
	this.FileId = fileId
	return &this
}

// NewPublicFileEggWithDefaults instantiates a new PublicFileEgg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicFileEggWithDefaults() *PublicFileEgg {
	this := PublicFileEgg{}
	var type_ string = "FILE"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PublicFileEgg) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PublicFileEgg) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PublicFileEgg) SetType(v string) {
	o.Type = v
}

// GetFileId returns the FileId field value
func (o *PublicFileEgg) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *PublicFileEgg) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *PublicFileEgg) SetFileId(v string) {
	o.FileId = v
}

func (o PublicFileEgg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicFileEgg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["fileId"] = o.FileId
	return toSerialize, nil
}

func (o *PublicFileEgg) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"fileId",
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

	varPublicFileEgg := _PublicFileEgg{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicFileEgg)

	if err != nil {
		return err
	}

	*o = PublicFileEgg(varPublicFileEgg)

	return err
}

type NullablePublicFileEgg struct {
	value *PublicFileEgg
	isSet bool
}

func (v NullablePublicFileEgg) Get() *PublicFileEgg {
	return v.value
}

func (v *NullablePublicFileEgg) Set(val *PublicFileEgg) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicFileEgg) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicFileEgg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicFileEgg(val *PublicFileEgg) *NullablePublicFileEgg {
	return &NullablePublicFileEgg{value: val, isSet: true}
}

func (v NullablePublicFileEgg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicFileEgg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


