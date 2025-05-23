/*
Associations

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package associations

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicDefaultAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicDefaultAssociation{}

// PublicDefaultAssociation struct for PublicDefaultAssociation
type PublicDefaultAssociation struct {
	AssociationSpec AssociationSpec `json:"associationSpec"`
	From PublicObjectId `json:"from"`
	To PublicObjectId `json:"to"`
}

type _PublicDefaultAssociation PublicDefaultAssociation

// NewPublicDefaultAssociation instantiates a new PublicDefaultAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicDefaultAssociation(associationSpec AssociationSpec, from PublicObjectId, to PublicObjectId) *PublicDefaultAssociation {
	this := PublicDefaultAssociation{}
	this.AssociationSpec = associationSpec
	this.From = from
	this.To = to
	return &this
}

// NewPublicDefaultAssociationWithDefaults instantiates a new PublicDefaultAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicDefaultAssociationWithDefaults() *PublicDefaultAssociation {
	this := PublicDefaultAssociation{}
	return &this
}

// GetAssociationSpec returns the AssociationSpec field value
func (o *PublicDefaultAssociation) GetAssociationSpec() AssociationSpec {
	if o == nil {
		var ret AssociationSpec
		return ret
	}

	return o.AssociationSpec
}

// GetAssociationSpecOk returns a tuple with the AssociationSpec field value
// and a boolean to check if the value has been set.
func (o *PublicDefaultAssociation) GetAssociationSpecOk() (*AssociationSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationSpec, true
}

// SetAssociationSpec sets field value
func (o *PublicDefaultAssociation) SetAssociationSpec(v AssociationSpec) {
	o.AssociationSpec = v
}

// GetFrom returns the From field value
func (o *PublicDefaultAssociation) GetFrom() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *PublicDefaultAssociation) GetFromOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *PublicDefaultAssociation) SetFrom(v PublicObjectId) {
	o.From = v
}

// GetTo returns the To field value
func (o *PublicDefaultAssociation) GetTo() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PublicDefaultAssociation) GetToOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *PublicDefaultAssociation) SetTo(v PublicObjectId) {
	o.To = v
}

func (o PublicDefaultAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicDefaultAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associationSpec"] = o.AssociationSpec
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	return toSerialize, nil
}

func (o *PublicDefaultAssociation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"associationSpec",
		"from",
		"to",
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

	varPublicDefaultAssociation := _PublicDefaultAssociation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicDefaultAssociation)

	if err != nil {
		return err
	}

	*o = PublicDefaultAssociation(varPublicDefaultAssociation)

	return err
}

type NullablePublicDefaultAssociation struct {
	value *PublicDefaultAssociation
	isSet bool
}

func (v NullablePublicDefaultAssociation) Get() *PublicDefaultAssociation {
	return v.value
}

func (v *NullablePublicDefaultAssociation) Set(val *PublicDefaultAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicDefaultAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicDefaultAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicDefaultAssociation(val *PublicDefaultAssociation) *NullablePublicDefaultAssociation {
	return &NullablePublicDefaultAssociation{value: val, isSet: true}
}

func (v NullablePublicDefaultAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicDefaultAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


