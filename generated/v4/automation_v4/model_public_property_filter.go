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

// checks if the PublicPropertyFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicPropertyFilter{}

// PublicPropertyFilter struct for PublicPropertyFilter
type PublicPropertyFilter struct {
	Property string `json:"property"`
	FilterType string `json:"filterType"`
	Operation PublicSurveyMonkeyValueFilterValueComparison `json:"operation"`
}

type _PublicPropertyFilter PublicPropertyFilter

// NewPublicPropertyFilter instantiates a new PublicPropertyFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicPropertyFilter(property string, filterType string, operation PublicSurveyMonkeyValueFilterValueComparison) *PublicPropertyFilter {
	this := PublicPropertyFilter{}
	this.Property = property
	this.FilterType = filterType
	this.Operation = operation
	return &this
}

// NewPublicPropertyFilterWithDefaults instantiates a new PublicPropertyFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicPropertyFilterWithDefaults() *PublicPropertyFilter {
	this := PublicPropertyFilter{}
	var filterType string = "PROPERTY"
	this.FilterType = filterType
	return &this
}

// GetProperty returns the Property field value
func (o *PublicPropertyFilter) GetProperty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *PublicPropertyFilter) GetPropertyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *PublicPropertyFilter) SetProperty(v string) {
	o.Property = v
}

// GetFilterType returns the FilterType field value
func (o *PublicPropertyFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicPropertyFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicPropertyFilter) SetFilterType(v string) {
	o.FilterType = v
}

// GetOperation returns the Operation field value
func (o *PublicPropertyFilter) GetOperation() PublicSurveyMonkeyValueFilterValueComparison {
	if o == nil {
		var ret PublicSurveyMonkeyValueFilterValueComparison
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *PublicPropertyFilter) GetOperationOk() (*PublicSurveyMonkeyValueFilterValueComparison, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *PublicPropertyFilter) SetOperation(v PublicSurveyMonkeyValueFilterValueComparison) {
	o.Operation = v
}

func (o PublicPropertyFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicPropertyFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["property"] = o.Property
	toSerialize["filterType"] = o.FilterType
	toSerialize["operation"] = o.Operation
	return toSerialize, nil
}

func (o *PublicPropertyFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"property",
		"filterType",
		"operation",
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

	varPublicPropertyFilter := _PublicPropertyFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicPropertyFilter)

	if err != nil {
		return err
	}

	*o = PublicPropertyFilter(varPublicPropertyFilter)

	return err
}

type NullablePublicPropertyFilter struct {
	value *PublicPropertyFilter
	isSet bool
}

func (v NullablePublicPropertyFilter) Get() *PublicPropertyFilter {
	return v.value
}

func (v *NullablePublicPropertyFilter) Set(val *PublicPropertyFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicPropertyFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicPropertyFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicPropertyFilter(val *PublicPropertyFilter) *NullablePublicPropertyFilter {
	return &NullablePublicPropertyFilter{value: val, isSet: true}
}

func (v NullablePublicPropertyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicPropertyFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


