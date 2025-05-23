/*
Marketing Emails V3

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_emails_v3

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EmailStatisticInterval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailStatisticInterval{}

// EmailStatisticInterval struct for EmailStatisticInterval
type EmailStatisticInterval struct {
	Interval Interval `json:"interval"`
	Aggregations EmailStatisticsData `json:"aggregations"`
}

type _EmailStatisticInterval EmailStatisticInterval

// NewEmailStatisticInterval instantiates a new EmailStatisticInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailStatisticInterval(interval Interval, aggregations EmailStatisticsData) *EmailStatisticInterval {
	this := EmailStatisticInterval{}
	this.Interval = interval
	this.Aggregations = aggregations
	return &this
}

// NewEmailStatisticIntervalWithDefaults instantiates a new EmailStatisticInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailStatisticIntervalWithDefaults() *EmailStatisticInterval {
	this := EmailStatisticInterval{}
	return &this
}

// GetInterval returns the Interval field value
func (o *EmailStatisticInterval) GetInterval() Interval {
	if o == nil {
		var ret Interval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *EmailStatisticInterval) GetIntervalOk() (*Interval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *EmailStatisticInterval) SetInterval(v Interval) {
	o.Interval = v
}

// GetAggregations returns the Aggregations field value
func (o *EmailStatisticInterval) GetAggregations() EmailStatisticsData {
	if o == nil {
		var ret EmailStatisticsData
		return ret
	}

	return o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value
// and a boolean to check if the value has been set.
func (o *EmailStatisticInterval) GetAggregationsOk() (*EmailStatisticsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregations, true
}

// SetAggregations sets field value
func (o *EmailStatisticInterval) SetAggregations(v EmailStatisticsData) {
	o.Aggregations = v
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
	toSerialize["interval"] = o.Interval
	toSerialize["aggregations"] = o.Aggregations
	return toSerialize, nil
}

func (o *EmailStatisticInterval) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interval",
		"aggregations",
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

	varEmailStatisticInterval := _EmailStatisticInterval{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailStatisticInterval)

	if err != nil {
		return err
	}

	*o = EmailStatisticInterval(varEmailStatisticInterval)

	return err
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


