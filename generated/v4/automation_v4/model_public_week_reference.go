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

// checks if the PublicWeekReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicWeekReference{}

// PublicWeekReference struct for PublicWeekReference
type PublicWeekReference struct {
	DayOfWeek string `json:"dayOfWeek"`
	Hour *int32 `json:"hour,omitempty"`
	Millisecond *int32 `json:"millisecond,omitempty"`
	ReferenceType string `json:"referenceType"`
	Minute *int32 `json:"minute,omitempty"`
	Second *int32 `json:"second,omitempty"`
}

type _PublicWeekReference PublicWeekReference

// NewPublicWeekReference instantiates a new PublicWeekReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicWeekReference(dayOfWeek string, referenceType string) *PublicWeekReference {
	this := PublicWeekReference{}
	this.DayOfWeek = dayOfWeek
	this.ReferenceType = referenceType
	return &this
}

// NewPublicWeekReferenceWithDefaults instantiates a new PublicWeekReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicWeekReferenceWithDefaults() *PublicWeekReference {
	this := PublicWeekReference{}
	var referenceType string = "WEEK"
	this.ReferenceType = referenceType
	return &this
}

// GetDayOfWeek returns the DayOfWeek field value
func (o *PublicWeekReference) GetDayOfWeek() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value
// and a boolean to check if the value has been set.
func (o *PublicWeekReference) GetDayOfWeekOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfWeek, true
}

// SetDayOfWeek sets field value
func (o *PublicWeekReference) SetDayOfWeek(v string) {
	o.DayOfWeek = v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *PublicWeekReference) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicWeekReference) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *PublicWeekReference) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *PublicWeekReference) SetHour(v int32) {
	o.Hour = &v
}

// GetMillisecond returns the Millisecond field value if set, zero value otherwise.
func (o *PublicWeekReference) GetMillisecond() int32 {
	if o == nil || IsNil(o.Millisecond) {
		var ret int32
		return ret
	}
	return *o.Millisecond
}

// GetMillisecondOk returns a tuple with the Millisecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicWeekReference) GetMillisecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Millisecond) {
		return nil, false
	}
	return o.Millisecond, true
}

// HasMillisecond returns a boolean if a field has been set.
func (o *PublicWeekReference) HasMillisecond() bool {
	if o != nil && !IsNil(o.Millisecond) {
		return true
	}

	return false
}

// SetMillisecond gets a reference to the given int32 and assigns it to the Millisecond field.
func (o *PublicWeekReference) SetMillisecond(v int32) {
	o.Millisecond = &v
}

// GetReferenceType returns the ReferenceType field value
func (o *PublicWeekReference) GetReferenceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value
// and a boolean to check if the value has been set.
func (o *PublicWeekReference) GetReferenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceType, true
}

// SetReferenceType sets field value
func (o *PublicWeekReference) SetReferenceType(v string) {
	o.ReferenceType = v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *PublicWeekReference) GetMinute() int32 {
	if o == nil || IsNil(o.Minute) {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicWeekReference) GetMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *PublicWeekReference) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *PublicWeekReference) SetMinute(v int32) {
	o.Minute = &v
}

// GetSecond returns the Second field value if set, zero value otherwise.
func (o *PublicWeekReference) GetSecond() int32 {
	if o == nil || IsNil(o.Second) {
		var ret int32
		return ret
	}
	return *o.Second
}

// GetSecondOk returns a tuple with the Second field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicWeekReference) GetSecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Second) {
		return nil, false
	}
	return o.Second, true
}

// HasSecond returns a boolean if a field has been set.
func (o *PublicWeekReference) HasSecond() bool {
	if o != nil && !IsNil(o.Second) {
		return true
	}

	return false
}

// SetSecond gets a reference to the given int32 and assigns it to the Second field.
func (o *PublicWeekReference) SetSecond(v int32) {
	o.Second = &v
}

func (o PublicWeekReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicWeekReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dayOfWeek"] = o.DayOfWeek
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if !IsNil(o.Millisecond) {
		toSerialize["millisecond"] = o.Millisecond
	}
	toSerialize["referenceType"] = o.ReferenceType
	if !IsNil(o.Minute) {
		toSerialize["minute"] = o.Minute
	}
	if !IsNil(o.Second) {
		toSerialize["second"] = o.Second
	}
	return toSerialize, nil
}

func (o *PublicWeekReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dayOfWeek",
		"referenceType",
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

	varPublicWeekReference := _PublicWeekReference{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicWeekReference)

	if err != nil {
		return err
	}

	*o = PublicWeekReference(varPublicWeekReference)

	return err
}

type NullablePublicWeekReference struct {
	value *PublicWeekReference
	isSet bool
}

func (v NullablePublicWeekReference) Get() *PublicWeekReference {
	return v.value
}

func (v *NullablePublicWeekReference) Set(val *PublicWeekReference) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicWeekReference) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicWeekReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicWeekReference(val *PublicWeekReference) *NullablePublicWeekReference {
	return &NullablePublicWeekReference{value: val, isSet: true}
}

func (v NullablePublicWeekReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicWeekReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


