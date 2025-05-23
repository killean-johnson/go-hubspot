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

// checks if the ApiDailyEnrollmentSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDailyEnrollmentSchedule{}

// ApiDailyEnrollmentSchedule struct for ApiDailyEnrollmentSchedule
type ApiDailyEnrollmentSchedule struct {
	// The type of enrollment schedule this is, can be: \"DAILY\", \"WEEKLY\", \"MONTHLY_SPECIFIC_DAYS\", \"MONTHLY_RELATIVE_DAYS\", \"YEARLY\"
	Type string `json:"type"`
	TimeOfDay ApiTimeOfDay `json:"timeOfDay"`
}

type _ApiDailyEnrollmentSchedule ApiDailyEnrollmentSchedule

// NewApiDailyEnrollmentSchedule instantiates a new ApiDailyEnrollmentSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDailyEnrollmentSchedule(type_ string, timeOfDay ApiTimeOfDay) *ApiDailyEnrollmentSchedule {
	this := ApiDailyEnrollmentSchedule{}
	this.Type = type_
	this.TimeOfDay = timeOfDay
	return &this
}

// NewApiDailyEnrollmentScheduleWithDefaults instantiates a new ApiDailyEnrollmentSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDailyEnrollmentScheduleWithDefaults() *ApiDailyEnrollmentSchedule {
	this := ApiDailyEnrollmentSchedule{}
	var type_ string = "DAILY"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ApiDailyEnrollmentSchedule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiDailyEnrollmentSchedule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiDailyEnrollmentSchedule) SetType(v string) {
	o.Type = v
}

// GetTimeOfDay returns the TimeOfDay field value
func (o *ApiDailyEnrollmentSchedule) GetTimeOfDay() ApiTimeOfDay {
	if o == nil {
		var ret ApiTimeOfDay
		return ret
	}

	return o.TimeOfDay
}

// GetTimeOfDayOk returns a tuple with the TimeOfDay field value
// and a boolean to check if the value has been set.
func (o *ApiDailyEnrollmentSchedule) GetTimeOfDayOk() (*ApiTimeOfDay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeOfDay, true
}

// SetTimeOfDay sets field value
func (o *ApiDailyEnrollmentSchedule) SetTimeOfDay(v ApiTimeOfDay) {
	o.TimeOfDay = v
}

func (o ApiDailyEnrollmentSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDailyEnrollmentSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timeOfDay"] = o.TimeOfDay
	return toSerialize, nil
}

func (o *ApiDailyEnrollmentSchedule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"timeOfDay",
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

	varApiDailyEnrollmentSchedule := _ApiDailyEnrollmentSchedule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiDailyEnrollmentSchedule)

	if err != nil {
		return err
	}

	*o = ApiDailyEnrollmentSchedule(varApiDailyEnrollmentSchedule)

	return err
}

type NullableApiDailyEnrollmentSchedule struct {
	value *ApiDailyEnrollmentSchedule
	isSet bool
}

func (v NullableApiDailyEnrollmentSchedule) Get() *ApiDailyEnrollmentSchedule {
	return v.value
}

func (v *NullableApiDailyEnrollmentSchedule) Set(val *ApiDailyEnrollmentSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDailyEnrollmentSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDailyEnrollmentSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDailyEnrollmentSchedule(val *ApiDailyEnrollmentSchedule) *NullableApiDailyEnrollmentSchedule {
	return &NullableApiDailyEnrollmentSchedule{value: val, isSet: true}
}

func (v NullableApiDailyEnrollmentSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDailyEnrollmentSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


