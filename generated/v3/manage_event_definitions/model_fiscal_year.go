/*
Events Manage Event Definitions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package manage_event_definitions

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FiscalYear type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FiscalYear{}

// FiscalYear struct for FiscalYear
type FiscalYear struct {
	Hour *int32 `json:"hour,omitempty"`
	Month int32 `json:"month"`
	Millisecond *int32 `json:"millisecond,omitempty"`
	ReferenceType string `json:"referenceType"`
	Day int32 `json:"day"`
	Minute *int32 `json:"minute,omitempty"`
	Second *int32 `json:"second,omitempty"`
}

type _FiscalYear FiscalYear

// NewFiscalYear instantiates a new FiscalYear object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiscalYear(month int32, referenceType string, day int32) *FiscalYear {
	this := FiscalYear{}
	this.Month = month
	this.ReferenceType = referenceType
	this.Day = day
	return &this
}

// NewFiscalYearWithDefaults instantiates a new FiscalYear object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiscalYearWithDefaults() *FiscalYear {
	this := FiscalYear{}
	var referenceType string = "FISCAL_YEAR"
	this.ReferenceType = referenceType
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *FiscalYear) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalYear) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *FiscalYear) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *FiscalYear) SetHour(v int32) {
	o.Hour = &v
}

// GetMonth returns the Month field value
func (o *FiscalYear) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *FiscalYear) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *FiscalYear) SetMonth(v int32) {
	o.Month = v
}

// GetMillisecond returns the Millisecond field value if set, zero value otherwise.
func (o *FiscalYear) GetMillisecond() int32 {
	if o == nil || IsNil(o.Millisecond) {
		var ret int32
		return ret
	}
	return *o.Millisecond
}

// GetMillisecondOk returns a tuple with the Millisecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalYear) GetMillisecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Millisecond) {
		return nil, false
	}
	return o.Millisecond, true
}

// HasMillisecond returns a boolean if a field has been set.
func (o *FiscalYear) HasMillisecond() bool {
	if o != nil && !IsNil(o.Millisecond) {
		return true
	}

	return false
}

// SetMillisecond gets a reference to the given int32 and assigns it to the Millisecond field.
func (o *FiscalYear) SetMillisecond(v int32) {
	o.Millisecond = &v
}

// GetReferenceType returns the ReferenceType field value
func (o *FiscalYear) GetReferenceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value
// and a boolean to check if the value has been set.
func (o *FiscalYear) GetReferenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceType, true
}

// SetReferenceType sets field value
func (o *FiscalYear) SetReferenceType(v string) {
	o.ReferenceType = v
}

// GetDay returns the Day field value
func (o *FiscalYear) GetDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *FiscalYear) GetDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *FiscalYear) SetDay(v int32) {
	o.Day = v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *FiscalYear) GetMinute() int32 {
	if o == nil || IsNil(o.Minute) {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalYear) GetMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *FiscalYear) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *FiscalYear) SetMinute(v int32) {
	o.Minute = &v
}

// GetSecond returns the Second field value if set, zero value otherwise.
func (o *FiscalYear) GetSecond() int32 {
	if o == nil || IsNil(o.Second) {
		var ret int32
		return ret
	}
	return *o.Second
}

// GetSecondOk returns a tuple with the Second field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalYear) GetSecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Second) {
		return nil, false
	}
	return o.Second, true
}

// HasSecond returns a boolean if a field has been set.
func (o *FiscalYear) HasSecond() bool {
	if o != nil && !IsNil(o.Second) {
		return true
	}

	return false
}

// SetSecond gets a reference to the given int32 and assigns it to the Second field.
func (o *FiscalYear) SetSecond(v int32) {
	o.Second = &v
}

func (o FiscalYear) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiscalYear) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	toSerialize["month"] = o.Month
	if !IsNil(o.Millisecond) {
		toSerialize["millisecond"] = o.Millisecond
	}
	toSerialize["referenceType"] = o.ReferenceType
	toSerialize["day"] = o.Day
	if !IsNil(o.Minute) {
		toSerialize["minute"] = o.Minute
	}
	if !IsNil(o.Second) {
		toSerialize["second"] = o.Second
	}
	return toSerialize, nil
}

func (o *FiscalYear) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"month",
		"referenceType",
		"day",
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

	varFiscalYear := _FiscalYear{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFiscalYear)

	if err != nil {
		return err
	}

	*o = FiscalYear(varFiscalYear)

	return err
}

type NullableFiscalYear struct {
	value *FiscalYear
	isSet bool
}

func (v NullableFiscalYear) Get() *FiscalYear {
	return v.value
}

func (v *NullableFiscalYear) Set(val *FiscalYear) {
	v.value = val
	v.isSet = true
}

func (v NullableFiscalYear) IsSet() bool {
	return v.isSet
}

func (v *NullableFiscalYear) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiscalYear(val *FiscalYear) *NullableFiscalYear {
	return &NullableFiscalYear{value: val, isSet: true}
}

func (v NullableFiscalYear) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiscalYear) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


