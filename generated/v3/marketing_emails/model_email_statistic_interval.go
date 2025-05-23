/*
Marketing Emails

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_emails

import (
	"encoding/json"
)

// checks if the EmailStatisticInterval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailStatisticInterval{}

// EmailStatisticInterval struct for EmailStatisticInterval
type EmailStatisticInterval struct {
	Interval *Interval `json:"interval,omitempty"`
	Aggregations *EmailStatisticsData `json:"aggregations,omitempty"`
}

// NewEmailStatisticInterval instantiates a new EmailStatisticInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailStatisticInterval() *EmailStatisticInterval {
	this := EmailStatisticInterval{}
	return &this
}

// NewEmailStatisticIntervalWithDefaults instantiates a new EmailStatisticInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailStatisticIntervalWithDefaults() *EmailStatisticInterval {
	this := EmailStatisticInterval{}
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *EmailStatisticInterval) GetInterval() Interval {
	if o == nil || IsNil(o.Interval) {
		var ret Interval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatisticInterval) GetIntervalOk() (*Interval, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *EmailStatisticInterval) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given Interval and assigns it to the Interval field.
func (o *EmailStatisticInterval) SetInterval(v Interval) {
	o.Interval = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *EmailStatisticInterval) GetAggregations() EmailStatisticsData {
	if o == nil || IsNil(o.Aggregations) {
		var ret EmailStatisticsData
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatisticInterval) GetAggregationsOk() (*EmailStatisticsData, bool) {
	if o == nil || IsNil(o.Aggregations) {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *EmailStatisticInterval) HasAggregations() bool {
	if o != nil && !IsNil(o.Aggregations) {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given EmailStatisticsData and assigns it to the Aggregations field.
func (o *EmailStatisticInterval) SetAggregations(v EmailStatisticsData) {
	o.Aggregations = &v
}

func (o EmailStatisticInterval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailStatisticInterval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Aggregations) {
		toSerialize["aggregations"] = o.Aggregations
	}
	return toSerialize, nil
}

type NullableEmailStatisticInterval struct {
	value *EmailStatisticInterval
	isSet bool
}

func (v NullableEmailStatisticInterval) Get() *EmailStatisticInterval {
	return v.value
}

func (v *NullableEmailStatisticInterval) Set(val *EmailStatisticInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailStatisticInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailStatisticInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailStatisticInterval(val *EmailStatisticInterval) *NullableEmailStatisticInterval {
	return &NullableEmailStatisticInterval{value: val, isSet: true}
}

func (v NullableEmailStatisticInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailStatisticInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


