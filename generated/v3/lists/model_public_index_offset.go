/*
Lists

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
)

// checks if the PublicIndexOffset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicIndexOffset{}

// PublicIndexOffset struct for PublicIndexOffset
type PublicIndexOffset struct {
	Milliseconds *int32 `json:"milliseconds,omitempty"`
	Hours *int32 `json:"hours,omitempty"`
	Seconds *int32 `json:"seconds,omitempty"`
	Months *int32 `json:"months,omitempty"`
	Weeks *int32 `json:"weeks,omitempty"`
	Minutes *int32 `json:"minutes,omitempty"`
	Quarters *int32 `json:"quarters,omitempty"`
	Days *int32 `json:"days,omitempty"`
	Years *int32 `json:"years,omitempty"`
}

// NewPublicIndexOffset instantiates a new PublicIndexOffset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIndexOffset() *PublicIndexOffset {
	this := PublicIndexOffset{}
	return &this
}

// NewPublicIndexOffsetWithDefaults instantiates a new PublicIndexOffset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIndexOffsetWithDefaults() *PublicIndexOffset {
	this := PublicIndexOffset{}
	return &this
}

// GetMilliseconds returns the Milliseconds field value if set, zero value otherwise.
func (o *PublicIndexOffset) GetMilliseconds() int32 {
	if o == nil || IsNil(o.Milliseconds) {
		var ret int32
		return ret
	}
	return *o.Milliseconds
}

// GetMillisecondsOk returns a tuple with the Milliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexOffset) GetMillisecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Milliseconds) {
		return nil, false
	}
	return o.Milliseconds, true
}

// HasMilliseconds returns a boolean if a field has been set.
func (o *PublicIndexOffset) HasMilliseconds() bool {
	if o != nil && !IsNil(o.Milliseconds) {
		return true
	}

	return false
}

// SetMilliseconds gets a reference to the given int32 and assigns it to the Milliseconds field.
func (o *PublicIndexOffset) SetMilliseconds(v int32) {
	o.Milliseconds = &v
}

// GetHours returns the Hours field value if set, zero value otherwise.
func (o *PublicIndexOffset) GetHours() int32 {
	if o == nil || IsNil(o.Hours) {
		var ret int32
		return ret
	}
	return *o.Hours
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexOffset) GetHoursOk() (*int32, bool) {
	if o == nil || IsNil(o.Hours) {
		return nil, false
	}
	return o.Hours, true
}

// HasHours returns a boolean if a field has been set.
func (o *PublicIndexOffset) HasHours() bool {
	if o != nil && !IsNil(o.Hours) {
		return true
	}

	return false
}

// SetHours gets a reference to the given int32 and assigns it to the Hours field.
func (o *PublicIndexOffset) SetHours(v int32) {
	o.Hours = &v
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *PublicIndexOffset) GetSeconds() int32 {
	if o == nil || IsNil(o.Seconds) {
		var ret int32
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexOffset) GetSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *PublicIndexOffset) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int32 and assigns it to the Seconds field.
func (o *PublicIndexOffset) SetSeconds(v int32) {
	o.Seconds = &v
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *PublicIndexOffset) GetMonths() int32 {
	if o == nil || IsNil(o.Months) {
		var ret int32
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexOffset) GetMonthsOk() (*int32, bool) {
	if o == nil || IsNil(o.Months) {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *PublicIndexOffset) HasMonths() bool {
	if o != nil && !IsNil(o.Months) {
		return true
	}

	return false
}

// SetMonths gets a reference to the given int32 and assigns it to the Months field.
func (o *PublicIndexOffset) SetMonths(v int32) {
	o.Months = &v
}

// GetWeeks returns the Weeks field value if set, zero value otherwise.
func (o *PublicIndexOffset) GetWeeks() int32 {
	if o == nil || IsNil(o.Weeks) {
		var ret int32
		return ret
	}
	return *o.Weeks
}

// GetWeeksOk returns a tuple with the Weeks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexOffset) GetWeeksOk() (*int32, bool) {
	if o == nil || IsNil(o.Weeks) {
		return nil, false
	}
	return o.Weeks, true
}

// HasWeeks returns a boolean if a field has been set.
func (o *PublicIndexOffset) HasWeeks() bool {
	if o != nil && !IsNil(o.Weeks) {
		return true
	}

	return false
}

// SetWeeks gets a reference to the given int32 and assigns it to the Weeks field.
func (o *PublicIndexOffset) SetWeeks(v int32) {
	o.Weeks = &v
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *PublicIndexOffset) GetMinutes() int32 {
	if o == nil || IsNil(o.Minutes) {
		var ret int32
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexOffset) GetMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *PublicIndexOffset) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given int32 and assigns it to the Minutes field.
func (o *PublicIndexOffset) SetMinutes(v int32) {
	o.Minutes = &v
}

// GetQuarters returns the Quarters field value if set, zero value otherwise.
func (o *PublicIndexOffset) GetQuarters() int32 {
	if o == nil || IsNil(o.Quarters) {
		var ret int32
		return ret
	}
	return *o.Quarters
}

// GetQuartersOk returns a tuple with the Quarters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexOffset) GetQuartersOk() (*int32, bool) {
	if o == nil || IsNil(o.Quarters) {
		return nil, false
	}
	return o.Quarters, true
}

// HasQuarters returns a boolean if a field has been set.
func (o *PublicIndexOffset) HasQuarters() bool {
	if o != nil && !IsNil(o.Quarters) {
		return true
	}

	return false
}

// SetQuarters gets a reference to the given int32 and assigns it to the Quarters field.
func (o *PublicIndexOffset) SetQuarters(v int32) {
	o.Quarters = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *PublicIndexOffset) GetDays() int32 {
	if o == nil || IsNil(o.Days) {
		var ret int32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexOffset) GetDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *PublicIndexOffset) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given int32 and assigns it to the Days field.
func (o *PublicIndexOffset) SetDays(v int32) {
	o.Days = &v
}

// GetYears returns the Years field value if set, zero value otherwise.
func (o *PublicIndexOffset) GetYears() int32 {
	if o == nil || IsNil(o.Years) {
		var ret int32
		return ret
	}
	return *o.Years
}

// GetYearsOk returns a tuple with the Years field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexOffset) GetYearsOk() (*int32, bool) {
	if o == nil || IsNil(o.Years) {
		return nil, false
	}
	return o.Years, true
}

// HasYears returns a boolean if a field has been set.
func (o *PublicIndexOffset) HasYears() bool {
	if o != nil && !IsNil(o.Years) {
		return true
	}

	return false
}

// SetYears gets a reference to the given int32 and assigns it to the Years field.
func (o *PublicIndexOffset) SetYears(v int32) {
	o.Years = &v
}

func (o PublicIndexOffset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicIndexOffset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Milliseconds) {
		toSerialize["milliseconds"] = o.Milliseconds
	}
	if !IsNil(o.Hours) {
		toSerialize["hours"] = o.Hours
	}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	if !IsNil(o.Months) {
		toSerialize["months"] = o.Months
	}
	if !IsNil(o.Weeks) {
		toSerialize["weeks"] = o.Weeks
	}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Quarters) {
		toSerialize["quarters"] = o.Quarters
	}
	if !IsNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	if !IsNil(o.Years) {
		toSerialize["years"] = o.Years
	}
	return toSerialize, nil
}

type NullablePublicIndexOffset struct {
	value *PublicIndexOffset
	isSet bool
}

func (v NullablePublicIndexOffset) Get() *PublicIndexOffset {
	return v.value
}

func (v *NullablePublicIndexOffset) Set(val *PublicIndexOffset) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIndexOffset) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIndexOffset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIndexOffset(val *PublicIndexOffset) *NullablePublicIndexOffset {
	return &NullablePublicIndexOffset{value: val, isSet: true}
}

func (v NullablePublicIndexOffset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIndexOffset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


