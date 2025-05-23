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

// checks if the ApiManualEnrollmentCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiManualEnrollmentCriteria{}

// ApiManualEnrollmentCriteria struct for ApiManualEnrollmentCriteria
type ApiManualEnrollmentCriteria struct {
	// Whether or not the same object can enroll in this workflow twice.
	ShouldReEnroll bool `json:"shouldReEnroll"`
	// The type of enrollment criteria this is, this can be \"LIST_BASED\", \"EVENT_BASED\", or \"MANUAL\".
	Type string `json:"type"`
}

type _ApiManualEnrollmentCriteria ApiManualEnrollmentCriteria

// NewApiManualEnrollmentCriteria instantiates a new ApiManualEnrollmentCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiManualEnrollmentCriteria(shouldReEnroll bool, type_ string) *ApiManualEnrollmentCriteria {
	this := ApiManualEnrollmentCriteria{}
	this.ShouldReEnroll = shouldReEnroll
	this.Type = type_
	return &this
}

// NewApiManualEnrollmentCriteriaWithDefaults instantiates a new ApiManualEnrollmentCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiManualEnrollmentCriteriaWithDefaults() *ApiManualEnrollmentCriteria {
	this := ApiManualEnrollmentCriteria{}
	var type_ string = "MANUAL"
	this.Type = type_
	return &this
}

// GetShouldReEnroll returns the ShouldReEnroll field value
func (o *ApiManualEnrollmentCriteria) GetShouldReEnroll() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldReEnroll
}

// GetShouldReEnrollOk returns a tuple with the ShouldReEnroll field value
// and a boolean to check if the value has been set.
func (o *ApiManualEnrollmentCriteria) GetShouldReEnrollOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldReEnroll, true
}

// SetShouldReEnroll sets field value
func (o *ApiManualEnrollmentCriteria) SetShouldReEnroll(v bool) {
	o.ShouldReEnroll = v
}

// GetType returns the Type field value
func (o *ApiManualEnrollmentCriteria) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiManualEnrollmentCriteria) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiManualEnrollmentCriteria) SetType(v string) {
	o.Type = v
}

func (o ApiManualEnrollmentCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiManualEnrollmentCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shouldReEnroll"] = o.ShouldReEnroll
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ApiManualEnrollmentCriteria) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shouldReEnroll",
		"type",
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

	varApiManualEnrollmentCriteria := _ApiManualEnrollmentCriteria{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiManualEnrollmentCriteria)

	if err != nil {
		return err
	}

	*o = ApiManualEnrollmentCriteria(varApiManualEnrollmentCriteria)

	return err
}

type NullableApiManualEnrollmentCriteria struct {
	value *ApiManualEnrollmentCriteria
	isSet bool
}

func (v NullableApiManualEnrollmentCriteria) Get() *ApiManualEnrollmentCriteria {
	return v.value
}

func (v *NullableApiManualEnrollmentCriteria) Set(val *ApiManualEnrollmentCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableApiManualEnrollmentCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableApiManualEnrollmentCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiManualEnrollmentCriteria(val *ApiManualEnrollmentCriteria) *NullableApiManualEnrollmentCriteria {
	return &NullableApiManualEnrollmentCriteria{value: val, isSet: true}
}

func (v NullableApiManualEnrollmentCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiManualEnrollmentCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


