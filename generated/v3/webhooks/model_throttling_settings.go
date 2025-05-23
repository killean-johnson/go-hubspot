/*
Webhooks Webhooks

Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhooks

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ThrottlingSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThrottlingSettings{}

// ThrottlingSettings Configuration details for webhook throttling.
type ThrottlingSettings struct {
	// The maximum number of concurrent HTTP requests HubSpot will attempt to make to your app.
	MaxConcurrentRequests int32 `json:"maxConcurrentRequests"`
}

type _ThrottlingSettings ThrottlingSettings

// NewThrottlingSettings instantiates a new ThrottlingSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThrottlingSettings(maxConcurrentRequests int32) *ThrottlingSettings {
	this := ThrottlingSettings{}
	this.MaxConcurrentRequests = maxConcurrentRequests
	return &this
}

// NewThrottlingSettingsWithDefaults instantiates a new ThrottlingSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThrottlingSettingsWithDefaults() *ThrottlingSettings {
	this := ThrottlingSettings{}
	return &this
}

// GetMaxConcurrentRequests returns the MaxConcurrentRequests field value
func (o *ThrottlingSettings) GetMaxConcurrentRequests() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxConcurrentRequests
}

// GetMaxConcurrentRequestsOk returns a tuple with the MaxConcurrentRequests field value
// and a boolean to check if the value has been set.
func (o *ThrottlingSettings) GetMaxConcurrentRequestsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxConcurrentRequests, true
}

// SetMaxConcurrentRequests sets field value
func (o *ThrottlingSettings) SetMaxConcurrentRequests(v int32) {
	o.MaxConcurrentRequests = v
}

func (o ThrottlingSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThrottlingSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxConcurrentRequests"] = o.MaxConcurrentRequests
	return toSerialize, nil
}

func (o *ThrottlingSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"maxConcurrentRequests",
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

	varThrottlingSettings := _ThrottlingSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThrottlingSettings)

	if err != nil {
		return err
	}

	*o = ThrottlingSettings(varThrottlingSettings)

	return err
}

type NullableThrottlingSettings struct {
	value *ThrottlingSettings
	isSet bool
}

func (v NullableThrottlingSettings) Get() *ThrottlingSettings {
	return v.value
}

func (v *NullableThrottlingSettings) Set(val *ThrottlingSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottlingSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottlingSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottlingSettings(val *ThrottlingSettings) *NullableThrottlingSettings {
	return &NullableThrottlingSettings{value: val, isSet: true}
}

func (v NullableThrottlingSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThrottlingSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


