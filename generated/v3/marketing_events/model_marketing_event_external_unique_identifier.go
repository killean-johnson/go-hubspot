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

// checks if the MarketingEventExternalUniqueIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketingEventExternalUniqueIdentifier{}

// MarketingEventExternalUniqueIdentifier struct for MarketingEventExternalUniqueIdentifier
type MarketingEventExternalUniqueIdentifier struct {
	// The accountId that is associated with this marketing event in the external event application.
	ExternalAccountId string `json:"externalAccountId"`
	// The id of the marketing event in the external event application.
	ExternalEventId string `json:"externalEventId"`
	// The id of the application that created the marketing event in HubSpot.
	AppId int32 `json:"appId"`
}

type _MarketingEventExternalUniqueIdentifier MarketingEventExternalUniqueIdentifier

// NewMarketingEventExternalUniqueIdentifier instantiates a new MarketingEventExternalUniqueIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingEventExternalUniqueIdentifier(externalAccountId string, externalEventId string, appId int32) *MarketingEventExternalUniqueIdentifier {
	this := MarketingEventExternalUniqueIdentifier{}
	this.ExternalAccountId = externalAccountId
	this.ExternalEventId = externalEventId
	this.AppId = appId
	return &this
}

// NewMarketingEventExternalUniqueIdentifierWithDefaults instantiates a new MarketingEventExternalUniqueIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingEventExternalUniqueIdentifierWithDefaults() *MarketingEventExternalUniqueIdentifier {
	this := MarketingEventExternalUniqueIdentifier{}
	return &this
}

// GetExternalAccountId returns the ExternalAccountId field value
func (o *MarketingEventExternalUniqueIdentifier) GetExternalAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalAccountId
}

// GetExternalAccountIdOk returns a tuple with the ExternalAccountId field value
// and a boolean to check if the value has been set.
func (o *MarketingEventExternalUniqueIdentifier) GetExternalAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalAccountId, true
}

// SetExternalAccountId sets field value
func (o *MarketingEventExternalUniqueIdentifier) SetExternalAccountId(v string) {
	o.ExternalAccountId = v
}

// GetExternalEventId returns the ExternalEventId field value
func (o *MarketingEventExternalUniqueIdentifier) GetExternalEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalEventId
}

// GetExternalEventIdOk returns a tuple with the ExternalEventId field value
// and a boolean to check if the value has been set.
func (o *MarketingEventExternalUniqueIdentifier) GetExternalEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalEventId, true
}

// SetExternalEventId sets field value
func (o *MarketingEventExternalUniqueIdentifier) SetExternalEventId(v string) {
	o.ExternalEventId = v
}

// GetAppId returns the AppId field value
func (o *MarketingEventExternalUniqueIdentifier) GetAppId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MarketingEventExternalUniqueIdentifier) GetAppIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *MarketingEventExternalUniqueIdentifier) SetAppId(v int32) {
	o.AppId = v
}

func (o MarketingEventExternalUniqueIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketingEventExternalUniqueIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalAccountId"] = o.ExternalAccountId
	toSerialize["externalEventId"] = o.ExternalEventId
	toSerialize["appId"] = o.AppId
	return toSerialize, nil
}

func (o *MarketingEventExternalUniqueIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalAccountId",
		"externalEventId",
		"appId",
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

	varMarketingEventExternalUniqueIdentifier := _MarketingEventExternalUniqueIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketingEventExternalUniqueIdentifier)

	if err != nil {
		return err
	}

	*o = MarketingEventExternalUniqueIdentifier(varMarketingEventExternalUniqueIdentifier)

	return err
}

type NullableMarketingEventExternalUniqueIdentifier struct {
	value *MarketingEventExternalUniqueIdentifier
	isSet bool
}

func (v NullableMarketingEventExternalUniqueIdentifier) Get() *MarketingEventExternalUniqueIdentifier {
	return v.value
}

func (v *NullableMarketingEventExternalUniqueIdentifier) Set(val *MarketingEventExternalUniqueIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingEventExternalUniqueIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingEventExternalUniqueIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingEventExternalUniqueIdentifier(val *MarketingEventExternalUniqueIdentifier) *NullableMarketingEventExternalUniqueIdentifier {
	return &NullableMarketingEventExternalUniqueIdentifier{value: val, isSet: true}
}

func (v NullableMarketingEventExternalUniqueIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingEventExternalUniqueIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


