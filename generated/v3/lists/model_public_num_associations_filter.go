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

// checks if the PublicNumAssociationsFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNumAssociationsFilter{}

// PublicNumAssociationsFilter struct for PublicNumAssociationsFilter
type PublicNumAssociationsFilter struct {
	CoalescingRefineBy PublicFormSubmissionFilterCoalescingRefineBy `json:"coalescingRefineBy"`
	AssociationTypeId int32 `json:"associationTypeId"`
	AssociationCategory string `json:"associationCategory"`
	FilterType string `json:"filterType"`
}

type _PublicNumAssociationsFilter PublicNumAssociationsFilter

// NewPublicNumAssociationsFilter instantiates a new PublicNumAssociationsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNumAssociationsFilter(coalescingRefineBy PublicFormSubmissionFilterCoalescingRefineBy, associationTypeId int32, associationCategory string, filterType string) *PublicNumAssociationsFilter {
	this := PublicNumAssociationsFilter{}
	this.CoalescingRefineBy = coalescingRefineBy
	this.AssociationTypeId = associationTypeId
	this.AssociationCategory = associationCategory
	this.FilterType = filterType
	return &this
}

// NewPublicNumAssociationsFilterWithDefaults instantiates a new PublicNumAssociationsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNumAssociationsFilterWithDefaults() *PublicNumAssociationsFilter {
	this := PublicNumAssociationsFilter{}
	var filterType string = "NUM_ASSOCIATIONS"
	this.FilterType = filterType
	return &this
}

// GetCoalescingRefineBy returns the CoalescingRefineBy field value
func (o *PublicNumAssociationsFilter) GetCoalescingRefineBy() PublicFormSubmissionFilterCoalescingRefineBy {
	if o == nil {
		var ret PublicFormSubmissionFilterCoalescingRefineBy
		return ret
	}

	return o.CoalescingRefineBy
}

// GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field value
// and a boolean to check if the value has been set.
func (o *PublicNumAssociationsFilter) GetCoalescingRefineByOk() (*PublicFormSubmissionFilterCoalescingRefineBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoalescingRefineBy, true
}

// SetCoalescingRefineBy sets field value
func (o *PublicNumAssociationsFilter) SetCoalescingRefineBy(v PublicFormSubmissionFilterCoalescingRefineBy) {
	o.CoalescingRefineBy = v
}

// GetAssociationTypeId returns the AssociationTypeId field value
func (o *PublicNumAssociationsFilter) GetAssociationTypeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssociationTypeId
}

// GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field value
// and a boolean to check if the value has been set.
func (o *PublicNumAssociationsFilter) GetAssociationTypeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationTypeId, true
}

// SetAssociationTypeId sets field value
func (o *PublicNumAssociationsFilter) SetAssociationTypeId(v int32) {
	o.AssociationTypeId = v
}

// GetAssociationCategory returns the AssociationCategory field value
func (o *PublicNumAssociationsFilter) GetAssociationCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssociationCategory
}

// GetAssociationCategoryOk returns a tuple with the AssociationCategory field value
// and a boolean to check if the value has been set.
func (o *PublicNumAssociationsFilter) GetAssociationCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationCategory, true
}

// SetAssociationCategory sets field value
func (o *PublicNumAssociationsFilter) SetAssociationCategory(v string) {
	o.AssociationCategory = v
}

// GetFilterType returns the FilterType field value
func (o *PublicNumAssociationsFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicNumAssociationsFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicNumAssociationsFilter) SetFilterType(v string) {
	o.FilterType = v
}

func (o PublicNumAssociationsFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicNumAssociationsFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coalescingRefineBy"] = o.CoalescingRefineBy
	toSerialize["associationTypeId"] = o.AssociationTypeId
	toSerialize["associationCategory"] = o.AssociationCategory
	toSerialize["filterType"] = o.FilterType
	return toSerialize, nil
}

func (o *PublicNumAssociationsFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"coalescingRefineBy",
		"associationTypeId",
		"associationCategory",
		"filterType",
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

	varPublicNumAssociationsFilter := _PublicNumAssociationsFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicNumAssociationsFilter)

	if err != nil {
		return err
	}

	*o = PublicNumAssociationsFilter(varPublicNumAssociationsFilter)

	return err
}

type NullablePublicNumAssociationsFilter struct {
	value *PublicNumAssociationsFilter
	isSet bool
}

func (v NullablePublicNumAssociationsFilter) Get() *PublicNumAssociationsFilter {
	return v.value
}

func (v *NullablePublicNumAssociationsFilter) Set(val *PublicNumAssociationsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNumAssociationsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNumAssociationsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNumAssociationsFilter(val *PublicNumAssociationsFilter) *NullablePublicNumAssociationsFilter {
	return &NullablePublicNumAssociationsFilter{value: val, isSet: true}
}

func (v NullablePublicNumAssociationsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNumAssociationsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


