/*
Marketing Events

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EventDetailSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventDetailSettings{}

// EventDetailSettings struct for EventDetailSettings
type EventDetailSettings struct {
	// The id of the application the settings are for
	AppId int32 `json:"appId"`
	// The url that will be used to fetch marketing event details by id
	EventDetailsUrl string `json:"eventDetailsUrl"`
}

type _EventDetailSettings EventDetailSettings

// NewEventDetailSettings instantiates a new EventDetailSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventDetailSettings(appId int32, eventDetailsUrl string) *EventDetailSettings {
	this := EventDetailSettings{}
	this.AppId = appId
	this.EventDetailsUrl = eventDetailsUrl
	return &this
}

// NewEventDetailSettingsWithDefaults instantiates a new EventDetailSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDetailSettingsWithDefaults() *EventDetailSettings {
	this := EventDetailSettings{}
	return &this
}

// GetAppId returns the AppId field value
func (o *EventDetailSettings) GetAppId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *EventDetailSettings) GetAppIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *EventDetailSettings) SetAppId(v int32) {
	o.AppId = v
}

// GetEventDetailsUrl returns the EventDetailsUrl field value
func (o *EventDetailSettings) GetEventDetailsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventDetailsUrl
}

// GetEventDetailsUrlOk returns a tuple with the EventDetailsUrl field value
// and a boolean to check if the value has been set.
func (o *EventDetailSettings) GetEventDetailsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDetailsUrl, true
}

// SetEventDetailsUrl sets field value
func (o *EventDetailSettings) SetEventDetailsUrl(v string) {
	o.EventDetailsUrl = v
}

func (o EventDetailSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventDetailSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	toSerialize["eventDetailsUrl"] = o.EventDetailsUrl
	return toSerialize, nil
}

func (o *EventDetailSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appId",
		"eventDetailsUrl",
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

	varEventDetailSettings := _EventDetailSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventDetailSettings)

	if err != nil {
		return err
	}

	*o = EventDetailSettings(varEventDetailSettings)

	return err
}

type NullableEventDetailSettings struct {
	value *EventDetailSettings
	isSet bool
}

func (v NullableEventDetailSettings) Get() *EventDetailSettings {
	return v.value
}

func (v *NullableEventDetailSettings) Set(val *EventDetailSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDetailSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDetailSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDetailSettings(val *EventDetailSettings) *NullableEventDetailSettings {
	return &NullableEventDetailSettings{value: val, isSet: true}
}

func (v NullableEventDetailSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDetailSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


