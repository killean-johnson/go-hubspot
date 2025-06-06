/*
Lists

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicMigrationMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicMigrationMapping{}

// PublicMigrationMapping struct for PublicMigrationMapping
type PublicMigrationMapping struct {
	// The V3 list id for the list
	ListId string `json:"listId"`
	// The legacy list id for the list
	LegacyListId string `json:"legacyListId"`
}

type _PublicMigrationMapping PublicMigrationMapping

// NewPublicMigrationMapping instantiates a new PublicMigrationMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicMigrationMapping(listId string, legacyListId string) *PublicMigrationMapping {
	this := PublicMigrationMapping{}
	this.ListId = listId
	this.LegacyListId = legacyListId
	return &this
}

// NewPublicMigrationMappingWithDefaults instantiates a new PublicMigrationMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicMigrationMappingWithDefaults() *PublicMigrationMapping {
	this := PublicMigrationMapping{}
	return &this
}

// GetListId returns the ListId field value
func (o *PublicMigrationMapping) GetListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *PublicMigrationMapping) GetListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *PublicMigrationMapping) SetListId(v string) {
	o.ListId = v
}

// GetLegacyListId returns the LegacyListId field value
func (o *PublicMigrationMapping) GetLegacyListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegacyListId
}

// GetLegacyListIdOk returns a tuple with the LegacyListId field value
// and a boolean to check if the value has been set.
func (o *PublicMigrationMapping) GetLegacyListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegacyListId, true
}

// SetLegacyListId sets field value
func (o *PublicMigrationMapping) SetLegacyListId(v string) {
	o.LegacyListId = v
}

func (o PublicMigrationMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicMigrationMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["listId"] = o.ListId
	toSerialize["legacyListId"] = o.LegacyListId
	return toSerialize, nil
}

func (o *PublicMigrationMapping) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"listId",
		"legacyListId",
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

	varPublicMigrationMapping := _PublicMigrationMapping{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicMigrationMapping)

	if err != nil {
		return err
	}

	*o = PublicMigrationMapping(varPublicMigrationMapping)

	return err
}

type NullablePublicMigrationMapping struct {
	value *PublicMigrationMapping
	isSet bool
}

func (v NullablePublicMigrationMapping) Get() *PublicMigrationMapping {
	return v.value
}

func (v *NullablePublicMigrationMapping) Set(val *PublicMigrationMapping) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicMigrationMapping) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicMigrationMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicMigrationMapping(val *PublicMigrationMapping) *NullablePublicMigrationMapping {
	return &NullablePublicMigrationMapping{value: val, isSet: true}
}

func (v NullablePublicMigrationMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicMigrationMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


