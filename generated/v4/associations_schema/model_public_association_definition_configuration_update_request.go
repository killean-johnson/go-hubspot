/*
Associations Schema

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package associations_schema

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicAssociationDefinitionConfigurationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicAssociationDefinitionConfigurationUpdateRequest{}

// PublicAssociationDefinitionConfigurationUpdateRequest struct for PublicAssociationDefinitionConfigurationUpdateRequest
type PublicAssociationDefinitionConfigurationUpdateRequest struct {
	TypeId int32 `json:"typeId"`
	Category string `json:"category"`
	MaxToObjectIds int32 `json:"maxToObjectIds"`
}

type _PublicAssociationDefinitionConfigurationUpdateRequest PublicAssociationDefinitionConfigurationUpdateRequest

// NewPublicAssociationDefinitionConfigurationUpdateRequest instantiates a new PublicAssociationDefinitionConfigurationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssociationDefinitionConfigurationUpdateRequest(typeId int32, category string, maxToObjectIds int32) *PublicAssociationDefinitionConfigurationUpdateRequest {
	this := PublicAssociationDefinitionConfigurationUpdateRequest{}
	this.TypeId = typeId
	this.Category = category
	this.MaxToObjectIds = maxToObjectIds
	return &this
}

// NewPublicAssociationDefinitionConfigurationUpdateRequestWithDefaults instantiates a new PublicAssociationDefinitionConfigurationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssociationDefinitionConfigurationUpdateRequestWithDefaults() *PublicAssociationDefinitionConfigurationUpdateRequest {
	this := PublicAssociationDefinitionConfigurationUpdateRequest{}
	return &this
}

// GetTypeId returns the TypeId field value
func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetTypeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetTypeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *PublicAssociationDefinitionConfigurationUpdateRequest) SetTypeId(v int32) {
	o.TypeId = v
}

// GetCategory returns the Category field value
func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *PublicAssociationDefinitionConfigurationUpdateRequest) SetCategory(v string) {
	o.Category = v
}

// GetMaxToObjectIds returns the MaxToObjectIds field value
func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetMaxToObjectIds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxToObjectIds
}

// GetMaxToObjectIdsOk returns a tuple with the MaxToObjectIds field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetMaxToObjectIdsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxToObjectIds, true
}

// SetMaxToObjectIds sets field value
func (o *PublicAssociationDefinitionConfigurationUpdateRequest) SetMaxToObjectIds(v int32) {
	o.MaxToObjectIds = v
}

func (o PublicAssociationDefinitionConfigurationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicAssociationDefinitionConfigurationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["typeId"] = o.TypeId
	toSerialize["category"] = o.Category
	toSerialize["maxToObjectIds"] = o.MaxToObjectIds
	return toSerialize, nil
}

func (o *PublicAssociationDefinitionConfigurationUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"typeId",
		"category",
		"maxToObjectIds",
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

	varPublicAssociationDefinitionConfigurationUpdateRequest := _PublicAssociationDefinitionConfigurationUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicAssociationDefinitionConfigurationUpdateRequest)

	if err != nil {
		return err
	}

	*o = PublicAssociationDefinitionConfigurationUpdateRequest(varPublicAssociationDefinitionConfigurationUpdateRequest)

	return err
}

type NullablePublicAssociationDefinitionConfigurationUpdateRequest struct {
	value *PublicAssociationDefinitionConfigurationUpdateRequest
	isSet bool
}

func (v NullablePublicAssociationDefinitionConfigurationUpdateRequest) Get() *PublicAssociationDefinitionConfigurationUpdateRequest {
	return v.value
}

func (v *NullablePublicAssociationDefinitionConfigurationUpdateRequest) Set(val *PublicAssociationDefinitionConfigurationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssociationDefinitionConfigurationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssociationDefinitionConfigurationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssociationDefinitionConfigurationUpdateRequest(val *PublicAssociationDefinitionConfigurationUpdateRequest) *NullablePublicAssociationDefinitionConfigurationUpdateRequest {
	return &NullablePublicAssociationDefinitionConfigurationUpdateRequest{value: val, isSet: true}
}

func (v NullablePublicAssociationDefinitionConfigurationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssociationDefinitionConfigurationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


