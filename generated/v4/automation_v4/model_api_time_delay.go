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

// checks if the ApiTimeDelay type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTimeDelay{}

// ApiTimeDelay struct for ApiTimeDelay
type ApiTimeDelay struct {
	TimeZoneStrategy *ApiTimeDelayTimeZoneStrategy `json:"timeZoneStrategy,omitempty"`
	Delta int32 `json:"delta"`
	DaysOfWeek []string `json:"daysOfWeek"`
	TimeUnit string `json:"timeUnit"`
	TimeOfDay *ApiTimeOfDay `json:"timeOfDay,omitempty"`
}

type _ApiTimeDelay ApiTimeDelay

// NewApiTimeDelay instantiates a new ApiTimeDelay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTimeDelay(delta int32, daysOfWeek []string, timeUnit string) *ApiTimeDelay {
	this := ApiTimeDelay{}
	this.Delta = delta
	this.DaysOfWeek = daysOfWeek
	this.TimeUnit = timeUnit
	return &this
}

// NewApiTimeDelayWithDefaults instantiates a new ApiTimeDelay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTimeDelayWithDefaults() *ApiTimeDelay {
	this := ApiTimeDelay{}
	return &this
}

// GetTimeZoneStrategy returns the TimeZoneStrategy field value if set, zero value otherwise.
func (o *ApiTimeDelay) GetTimeZoneStrategy() ApiTimeDelayTimeZoneStrategy {
	if o == nil || IsNil(o.TimeZoneStrategy) {
		var ret ApiTimeDelayTimeZoneStrategy
		return ret
	}
	return *o.TimeZoneStrategy
}

// GetTimeZoneStrategyOk returns a tuple with the TimeZoneStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTimeDelay) GetTimeZoneStrategyOk() (*ApiTimeDelayTimeZoneStrategy, bool) {
	if o == nil || IsNil(o.TimeZoneStrategy) {
		return nil, false
	}
	return o.TimeZoneStrategy, true
}

// HasTimeZoneStrategy returns a boolean if a field has been set.
func (o *ApiTimeDelay) HasTimeZoneStrategy() bool {
	if o != nil && !IsNil(o.TimeZoneStrategy) {
		return true
	}

	return false
}

// SetTimeZoneStrategy gets a reference to the given ApiTimeDelayTimeZoneStrategy and assigns it to the TimeZoneStrategy field.
func (o *ApiTimeDelay) SetTimeZoneStrategy(v ApiTimeDelayTimeZoneStrategy) {
	o.TimeZoneStrategy = &v
}

// GetDelta returns the Delta field value
func (o *ApiTimeDelay) GetDelta() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value
// and a boolean to check if the value has been set.
func (o *ApiTimeDelay) GetDeltaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delta, true
}

// SetDelta sets field value
func (o *ApiTimeDelay) SetDelta(v int32) {
	o.Delta = v
}

// GetDaysOfWeek returns the DaysOfWeek field value
func (o *ApiTimeDelay) GetDaysOfWeek() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DaysOfWeek
}

// GetDaysOfWeekOk returns a tuple with the DaysOfWeek field value
// and a boolean to check if the value has been set.
func (o *ApiTimeDelay) GetDaysOfWeekOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DaysOfWeek, true
}

// SetDaysOfWeek sets field value
func (o *ApiTimeDelay) SetDaysOfWeek(v []string) {
	o.DaysOfWeek = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *ApiTimeDelay) GetTimeUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *ApiTimeDelay) GetTimeUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *ApiTimeDelay) SetTimeUnit(v string) {
	o.TimeUnit = v
}

// GetTimeOfDay returns the TimeOfDay field value if set, zero value otherwise.
func (o *ApiTimeDelay) GetTimeOfDay() ApiTimeOfDay {
	if o == nil || IsNil(o.TimeOfDay) {
		var ret ApiTimeOfDay
		return ret
	}
	return *o.TimeOfDay
}

// GetTimeOfDayOk returns a tuple with the TimeOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTimeDelay) GetTimeOfDayOk() (*ApiTimeOfDay, bool) {
	if o == nil || IsNil(o.TimeOfDay) {
		return nil, false
	}
	return o.TimeOfDay, true
}

// HasTimeOfDay returns a boolean if a field has been set.
func (o *ApiTimeDelay) HasTimeOfDay() bool {
	if o != nil && !IsNil(o.TimeOfDay) {
		return true
	}

	return false
}

// SetTimeOfDay gets a reference to the given ApiTimeOfDay and assigns it to the TimeOfDay field.
func (o *ApiTimeDelay) SetTimeOfDay(v ApiTimeOfDay) {
	o.TimeOfDay = &v
}

func (o ApiTimeDelay) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTimeDelay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeZoneStrategy) {
		toSerialize["timeZoneStrategy"] = o.TimeZoneStrategy
	}
	toSerialize["delta"] = o.Delta
	toSerialize["daysOfWeek"] = o.DaysOfWeek
	toSerialize["timeUnit"] = o.TimeUnit
	if !IsNil(o.TimeOfDay) {
		toSerialize["timeOfDay"] = o.TimeOfDay
	}
	return toSerialize, nil
}

func (o *ApiTimeDelay) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"delta",
		"daysOfWeek",
		"timeUnit",
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

	varApiTimeDelay := _ApiTimeDelay{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTimeDelay)

	if err != nil {
		return err
	}

	*o = ApiTimeDelay(varApiTimeDelay)

	return err
}

type NullableApiTimeDelay struct {
	value *ApiTimeDelay
	isSet bool
}

func (v NullableApiTimeDelay) Get() *ApiTimeDelay {
	return v.value
}

func (v *NullableApiTimeDelay) Set(val *ApiTimeDelay) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTimeDelay) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTimeDelay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTimeDelay(val *ApiTimeDelay) *NullableApiTimeDelay {
	return &NullableApiTimeDelay{value: val, isSet: true}
}

func (v NullableApiTimeDelay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTimeDelay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


