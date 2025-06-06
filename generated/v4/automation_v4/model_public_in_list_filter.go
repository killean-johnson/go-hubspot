/*
Automation V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package automation_v4

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicInListFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicInListFilter{}

// PublicInListFilter struct for PublicInListFilter
type PublicInListFilter struct {
	ListId string `json:"listId"`
	Metadata *PublicInListFilterMetadata `json:"metadata,omitempty"`
	FilterType string `json:"filterType"`
	Operator string `json:"operator"`
}

type _PublicInListFilter PublicInListFilter

// NewPublicInListFilter instantiates a new PublicInListFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicInListFilter(listId string, filterType string, operator string) *PublicInListFilter {
	this := PublicInListFilter{}
	this.ListId = listId
	this.FilterType = filterType
	this.Operator = operator
	return &this
}

// NewPublicInListFilterWithDefaults instantiates a new PublicInListFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicInListFilterWithDefaults() *PublicInListFilter {
	this := PublicInListFilter{}
	var filterType string = "IN_LIST"
	this.FilterType = filterType
	return &this
}

// GetListId returns the ListId field value
func (o *PublicInListFilter) GetListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *PublicInListFilter) GetListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *PublicInListFilter) SetListId(v string) {
	o.ListId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PublicInListFilter) GetMetadata() PublicInListFilterMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret PublicInListFilterMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicInListFilter) GetMetadataOk() (*PublicInListFilterMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PublicInListFilter) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given PublicInListFilterMetadata and assigns it to the Metadata field.
func (o *PublicInListFilter) SetMetadata(v PublicInListFilterMetadata) {
	o.Metadata = &v
}

// GetFilterType returns the FilterType field value
func (o *PublicInListFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicInListFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicInListFilter) SetFilterType(v string) {
	o.FilterType = v
}

// GetOperator returns the Operator field value
func (o *PublicInListFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PublicInListFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PublicInListFilter) SetOperator(v string) {
	o.Operator = v
}

func (o PublicInListFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicInListFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["listId"] = o.ListId
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["filterType"] = o.FilterType
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *PublicInListFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"listId",
		"filterType",
		"operator",
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

	varPublicInListFilter := _PublicInListFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicInListFilter)

	if err != nil {
		return err
	}

	*o = PublicInListFilter(varPublicInListFilter)

	return err
}

type NullablePublicInListFilter struct {
	value *PublicInListFilter
	isSet bool
}

func (v NullablePublicInListFilter) Get() *PublicInListFilter {
	return v.value
}

func (v *NullablePublicInListFilter) Set(val *PublicInListFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicInListFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicInListFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicInListFilter(val *PublicInListFilter) *NullablePublicInListFilter {
	return &NullablePublicInListFilter{value: val, isSet: true}
}

func (v NullablePublicInListFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicInListFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


