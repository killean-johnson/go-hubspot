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

// checks if the ExternalMeetingsUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalMeetingsUser{}

// ExternalMeetingsUser struct for ExternalMeetingsUser
type ExternalMeetingsUser struct {
	IsSalesStarter bool `json:"isSalesStarter"`
	CalendarProvider string `json:"calendarProvider"`
	Id string `json:"id"`
	UserId string `json:"userId"`
	UserProfile ExternalUserProfile `json:"userProfile"`
}

type _ExternalMeetingsUser ExternalMeetingsUser

// NewExternalMeetingsUser instantiates a new ExternalMeetingsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalMeetingsUser(isSalesStarter bool, calendarProvider string, id string, userId string, userProfile ExternalUserProfile) *ExternalMeetingsUser {
	this := ExternalMeetingsUser{}
	this.IsSalesStarter = isSalesStarter
	this.CalendarProvider = calendarProvider
	this.Id = id
	this.UserId = userId
	this.UserProfile = userProfile
	return &this
}

// NewExternalMeetingsUserWithDefaults instantiates a new ExternalMeetingsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalMeetingsUserWithDefaults() *ExternalMeetingsUser {
	this := ExternalMeetingsUser{}
	return &this
}

// GetIsSalesStarter returns the IsSalesStarter field value
func (o *ExternalMeetingsUser) GetIsSalesStarter() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSalesStarter
}

// GetIsSalesStarterOk returns a tuple with the IsSalesStarter field value
// and a boolean to check if the value has been set.
func (o *ExternalMeetingsUser) GetIsSalesStarterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSalesStarter, true
}

// SetIsSalesStarter sets field value
func (o *ExternalMeetingsUser) SetIsSalesStarter(v bool) {
	o.IsSalesStarter = v
}

// GetCalendarProvider returns the CalendarProvider field value
func (o *ExternalMeetingsUser) GetCalendarProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CalendarProvider
}

// GetCalendarProviderOk returns a tuple with the CalendarProvider field value
// and a boolean to check if the value has been set.
func (o *ExternalMeetingsUser) GetCalendarProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CalendarProvider, true
}

// SetCalendarProvider sets field value
func (o *ExternalMeetingsUser) SetCalendarProvider(v string) {
	o.CalendarProvider = v
}

// GetId returns the Id field value
func (o *ExternalMeetingsUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalMeetingsUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalMeetingsUser) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *ExternalMeetingsUser) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ExternalMeetingsUser) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ExternalMeetingsUser) SetUserId(v string) {
	o.UserId = v
}

// GetUserProfile returns the UserProfile field value
func (o *ExternalMeetingsUser) GetUserProfile() ExternalUserProfile {
	if o == nil {
		var ret ExternalUserProfile
		return ret
	}

	return o.UserProfile
}

// GetUserProfileOk returns a tuple with the UserProfile field value
// and a boolean to check if the value has been set.
func (o *ExternalMeetingsUser) GetUserProfileOk() (*ExternalUserProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserProfile, true
}

// SetUserProfile sets field value
func (o *ExternalMeetingsUser) SetUserProfile(v ExternalUserProfile) {
	o.UserProfile = v
}

func (o ExternalMeetingsUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalMeetingsUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isSalesStarter"] = o.IsSalesStarter
	toSerialize["calendarProvider"] = o.CalendarProvider
	toSerialize["id"] = o.Id
	toSerialize["userId"] = o.UserId
	toSerialize["userProfile"] = o.UserProfile
	return toSerialize, nil
}

func (o *ExternalMeetingsUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isSalesStarter",
		"calendarProvider",
		"id",
		"userId",
		"userProfile",
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

	varExternalMeetingsUser := _ExternalMeetingsUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalMeetingsUser)

	if err != nil {
		return err
	}

	*o = ExternalMeetingsUser(varExternalMeetingsUser)

	return err
}

type NullableExternalMeetingsUser struct {
	value *ExternalMeetingsUser
	isSet bool
}

func (v NullableExternalMeetingsUser) Get() *ExternalMeetingsUser {
	return v.value
}

func (v *NullableExternalMeetingsUser) Set(val *ExternalMeetingsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalMeetingsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalMeetingsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalMeetingsUser(val *ExternalMeetingsUser) *NullableExternalMeetingsUser {
	return &NullableExternalMeetingsUser{value: val, isSet: true}
}

func (v NullableExternalMeetingsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalMeetingsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


