/*
Scheduler Meetings

Meetings Service For HubSpot Sales

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meetings

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ExternalBookingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalBookingInfo{}

// ExternalBookingInfo struct for ExternalBookingInfo
type ExternalBookingInfo struct {
	LinkId string `json:"linkId"`
	CustomParams ExternalMeetingsLinkSettings `json:"customParams"`
	LinkAvailability *ExternalLinkAvailability `json:"linkAvailability,omitempty"`
	BrandingMetadata *ExternalBrandingMetadata `json:"brandingMetadata,omitempty"`
	IsOffline bool `json:"isOffline"`
	LinkType string `json:"linkType"`
	AllUsersBusyTimes []ExternalUserBusyTimes `json:"allUsersBusyTimes"`
}

type _ExternalBookingInfo ExternalBookingInfo

// NewExternalBookingInfo instantiates a new ExternalBookingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalBookingInfo(linkId string, customParams ExternalMeetingsLinkSettings, isOffline bool, linkType string, allUsersBusyTimes []ExternalUserBusyTimes) *ExternalBookingInfo {
	this := ExternalBookingInfo{}
	this.LinkId = linkId
	this.CustomParams = customParams
	this.IsOffline = isOffline
	this.LinkType = linkType
	this.AllUsersBusyTimes = allUsersBusyTimes
	return &this
}

// NewExternalBookingInfoWithDefaults instantiates a new ExternalBookingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalBookingInfoWithDefaults() *ExternalBookingInfo {
	this := ExternalBookingInfo{}
	return &this
}

// GetLinkId returns the LinkId field value
func (o *ExternalBookingInfo) GetLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value
// and a boolean to check if the value has been set.
func (o *ExternalBookingInfo) GetLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkId, true
}

// SetLinkId sets field value
func (o *ExternalBookingInfo) SetLinkId(v string) {
	o.LinkId = v
}

// GetCustomParams returns the CustomParams field value
func (o *ExternalBookingInfo) GetCustomParams() ExternalMeetingsLinkSettings {
	if o == nil {
		var ret ExternalMeetingsLinkSettings
		return ret
	}

	return o.CustomParams
}

// GetCustomParamsOk returns a tuple with the CustomParams field value
// and a boolean to check if the value has been set.
func (o *ExternalBookingInfo) GetCustomParamsOk() (*ExternalMeetingsLinkSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomParams, true
}

// SetCustomParams sets field value
func (o *ExternalBookingInfo) SetCustomParams(v ExternalMeetingsLinkSettings) {
	o.CustomParams = v
}

// GetLinkAvailability returns the LinkAvailability field value if set, zero value otherwise.
func (o *ExternalBookingInfo) GetLinkAvailability() ExternalLinkAvailability {
	if o == nil || IsNil(o.LinkAvailability) {
		var ret ExternalLinkAvailability
		return ret
	}
	return *o.LinkAvailability
}

// GetLinkAvailabilityOk returns a tuple with the LinkAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalBookingInfo) GetLinkAvailabilityOk() (*ExternalLinkAvailability, bool) {
	if o == nil || IsNil(o.LinkAvailability) {
		return nil, false
	}
	return o.LinkAvailability, true
}

// HasLinkAvailability returns a boolean if a field has been set.
func (o *ExternalBookingInfo) HasLinkAvailability() bool {
	if o != nil && !IsNil(o.LinkAvailability) {
		return true
	}

	return false
}

// SetLinkAvailability gets a reference to the given ExternalLinkAvailability and assigns it to the LinkAvailability field.
func (o *ExternalBookingInfo) SetLinkAvailability(v ExternalLinkAvailability) {
	o.LinkAvailability = &v
}

// GetBrandingMetadata returns the BrandingMetadata field value if set, zero value otherwise.
func (o *ExternalBookingInfo) GetBrandingMetadata() ExternalBrandingMetadata {
	if o == nil || IsNil(o.BrandingMetadata) {
		var ret ExternalBrandingMetadata
		return ret
	}
	return *o.BrandingMetadata
}

// GetBrandingMetadataOk returns a tuple with the BrandingMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalBookingInfo) GetBrandingMetadataOk() (*ExternalBrandingMetadata, bool) {
	if o == nil || IsNil(o.BrandingMetadata) {
		return nil, false
	}
	return o.BrandingMetadata, true
}

// HasBrandingMetadata returns a boolean if a field has been set.
func (o *ExternalBookingInfo) HasBrandingMetadata() bool {
	if o != nil && !IsNil(o.BrandingMetadata) {
		return true
	}

	return false
}

// SetBrandingMetadata gets a reference to the given ExternalBrandingMetadata and assigns it to the BrandingMetadata field.
func (o *ExternalBookingInfo) SetBrandingMetadata(v ExternalBrandingMetadata) {
	o.BrandingMetadata = &v
}

// GetIsOffline returns the IsOffline field value
func (o *ExternalBookingInfo) GetIsOffline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOffline
}

// GetIsOfflineOk returns a tuple with the IsOffline field value
// and a boolean to check if the value has been set.
func (o *ExternalBookingInfo) GetIsOfflineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsOffline, true
}

// SetIsOffline sets field value
func (o *ExternalBookingInfo) SetIsOffline(v bool) {
	o.IsOffline = v
}

// GetLinkType returns the LinkType field value
func (o *ExternalBookingInfo) GetLinkType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value
// and a boolean to check if the value has been set.
func (o *ExternalBookingInfo) GetLinkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkType, true
}

// SetLinkType sets field value
func (o *ExternalBookingInfo) SetLinkType(v string) {
	o.LinkType = v
}

// GetAllUsersBusyTimes returns the AllUsersBusyTimes field value
func (o *ExternalBookingInfo) GetAllUsersBusyTimes() []ExternalUserBusyTimes {
	if o == nil {
		var ret []ExternalUserBusyTimes
		return ret
	}

	return o.AllUsersBusyTimes
}

// GetAllUsersBusyTimesOk returns a tuple with the AllUsersBusyTimes field value
// and a boolean to check if the value has been set.
func (o *ExternalBookingInfo) GetAllUsersBusyTimesOk() ([]ExternalUserBusyTimes, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllUsersBusyTimes, true
}

// SetAllUsersBusyTimes sets field value
func (o *ExternalBookingInfo) SetAllUsersBusyTimes(v []ExternalUserBusyTimes) {
	o.AllUsersBusyTimes = v
}

func (o ExternalBookingInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalBookingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["linkId"] = o.LinkId
	toSerialize["customParams"] = o.CustomParams
	if !IsNil(o.LinkAvailability) {
		toSerialize["linkAvailability"] = o.LinkAvailability
	}
	if !IsNil(o.BrandingMetadata) {
		toSerialize["brandingMetadata"] = o.BrandingMetadata
	}
	toSerialize["isOffline"] = o.IsOffline
	toSerialize["linkType"] = o.LinkType
	toSerialize["allUsersBusyTimes"] = o.AllUsersBusyTimes
	return toSerialize, nil
}

func (o *ExternalBookingInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"linkId",
		"customParams",
		"isOffline",
		"linkType",
		"allUsersBusyTimes",
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

	varExternalBookingInfo := _ExternalBookingInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalBookingInfo)

	if err != nil {
		return err
	}

	*o = ExternalBookingInfo(varExternalBookingInfo)

	return err
}

type NullableExternalBookingInfo struct {
	value *ExternalBookingInfo
	isSet bool
}

func (v NullableExternalBookingInfo) Get() *ExternalBookingInfo {
	return v.value
}

func (v *NullableExternalBookingInfo) Set(val *ExternalBookingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalBookingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalBookingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalBookingInfo(val *ExternalBookingInfo) *NullableExternalBookingInfo {
	return &NullableExternalBookingInfo{value: val, isSet: true}
}

func (v NullableExternalBookingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalBookingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


