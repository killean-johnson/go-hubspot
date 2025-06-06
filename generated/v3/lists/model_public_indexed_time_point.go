/*
Lists

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicIndexedTimePoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicIndexedTimePoint{}

// PublicIndexedTimePoint struct for PublicIndexedTimePoint
type PublicIndexedTimePoint struct {
	Offset *PublicIndexOffset `json:"offset,omitempty"`
	TimezoneSource *string `json:"timezoneSource,omitempty"`
	IndexReference PublicIndexedTimePointIndexReference `json:"indexReference"`
	TimeType string `json:"timeType"`
	ZoneId string `json:"zoneId"`
}

type _PublicIndexedTimePoint PublicIndexedTimePoint

// NewPublicIndexedTimePoint instantiates a new PublicIndexedTimePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIndexedTimePoint(indexReference PublicIndexedTimePointIndexReference, timeType string, zoneId string) *PublicIndexedTimePoint {
	this := PublicIndexedTimePoint{}
	this.IndexReference = indexReference
	this.TimeType = timeType
	this.ZoneId = zoneId
	return &this
}

// NewPublicIndexedTimePointWithDefaults instantiates a new PublicIndexedTimePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIndexedTimePointWithDefaults() *PublicIndexedTimePoint {
	this := PublicIndexedTimePoint{}
	var timeType string = "INDEXED"
	this.TimeType = timeType
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PublicIndexedTimePoint) GetOffset() PublicIndexOffset {
	if o == nil || IsNil(o.Offset) {
		var ret PublicIndexOffset
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexedTimePoint) GetOffsetOk() (*PublicIndexOffset, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PublicIndexedTimePoint) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given PublicIndexOffset and assigns it to the Offset field.
func (o *PublicIndexedTimePoint) SetOffset(v PublicIndexOffset) {
	o.Offset = &v
}

// GetTimezoneSource returns the TimezoneSource field value if set, zero value otherwise.
func (o *PublicIndexedTimePoint) GetTimezoneSource() string {
	if o == nil || IsNil(o.TimezoneSource) {
		var ret string
		return ret
	}
	return *o.TimezoneSource
}

// GetTimezoneSourceOk returns a tuple with the TimezoneSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIndexedTimePoint) GetTimezoneSourceOk() (*string, bool) {
	if o == nil || IsNil(o.TimezoneSource) {
		return nil, false
	}
	return o.TimezoneSource, true
}

// HasTimezoneSource returns a boolean if a field has been set.
func (o *PublicIndexedTimePoint) HasTimezoneSource() bool {
	if o != nil && !IsNil(o.TimezoneSource) {
		return true
	}

	return false
}

// SetTimezoneSource gets a reference to the given string and assigns it to the TimezoneSource field.
func (o *PublicIndexedTimePoint) SetTimezoneSource(v string) {
	o.TimezoneSource = &v
}

// GetIndexReference returns the IndexReference field value
func (o *PublicIndexedTimePoint) GetIndexReference() PublicIndexedTimePointIndexReference {
	if o == nil {
		var ret PublicIndexedTimePointIndexReference
		return ret
	}

	return o.IndexReference
}

// GetIndexReferenceOk returns a tuple with the IndexReference field value
// and a boolean to check if the value has been set.
func (o *PublicIndexedTimePoint) GetIndexReferenceOk() (*PublicIndexedTimePointIndexReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexReference, true
}

// SetIndexReference sets field value
func (o *PublicIndexedTimePoint) SetIndexReference(v PublicIndexedTimePointIndexReference) {
	o.IndexReference = v
}

// GetTimeType returns the TimeType field value
func (o *PublicIndexedTimePoint) GetTimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeType
}

// GetTimeTypeOk returns a tuple with the TimeType field value
// and a boolean to check if the value has been set.
func (o *PublicIndexedTimePoint) GetTimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeType, true
}

// SetTimeType sets field value
func (o *PublicIndexedTimePoint) SetTimeType(v string) {
	o.TimeType = v
}

// GetZoneId returns the ZoneId field value
func (o *PublicIndexedTimePoint) GetZoneId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value
// and a boolean to check if the value has been set.
func (o *PublicIndexedTimePoint) GetZoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneId, true
}

// SetZoneId sets field value
func (o *PublicIndexedTimePoint) SetZoneId(v string) {
	o.ZoneId = v
}

func (o PublicIndexedTimePoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicIndexedTimePoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.TimezoneSource) {
		toSerialize["timezoneSource"] = o.TimezoneSource
	}
	toSerialize["indexReference"] = o.IndexReference
	toSerialize["timeType"] = o.TimeType
	toSerialize["zoneId"] = o.ZoneId
	return toSerialize, nil
}

func (o *PublicIndexedTimePoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"indexReference",
		"timeType",
		"zoneId",
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

	varPublicIndexedTimePoint := _PublicIndexedTimePoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicIndexedTimePoint)

	if err != nil {
		return err
	}

	*o = PublicIndexedTimePoint(varPublicIndexedTimePoint)

	return err
}

type NullablePublicIndexedTimePoint struct {
	value *PublicIndexedTimePoint
	isSet bool
}

func (v NullablePublicIndexedTimePoint) Get() *PublicIndexedTimePoint {
	return v.value
}

func (v *NullablePublicIndexedTimePoint) Set(val *PublicIndexedTimePoint) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIndexedTimePoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIndexedTimePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIndexedTimePoint(val *PublicIndexedTimePoint) *NullablePublicIndexedTimePoint {
	return &NullablePublicIndexedTimePoint{value: val, isSet: true}
}

func (v NullablePublicIndexedTimePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIndexedTimePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


