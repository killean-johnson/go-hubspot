/*
Marketing Emails

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_emails

import (
	"encoding/json"
)

// checks if the PublicEmailToDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicEmailToDetails{}

// PublicEmailToDetails Data structure representing the to fields of the email.
type PublicEmailToDetails struct {
	ContactIlsLists *PublicEmailRecipients `json:"contactIlsLists,omitempty"`
	// 
	LimitSendFrequency *bool `json:"limitSendFrequency,omitempty"`
	// 
	SuppressGraymail *bool `json:"suppressGraymail,omitempty"`
	ContactIds *PublicEmailRecipients `json:"contactIds,omitempty"`
	ContactLists *PublicEmailRecipients `json:"contactLists,omitempty"`
}

// NewPublicEmailToDetails instantiates a new PublicEmailToDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicEmailToDetails() *PublicEmailToDetails {
	this := PublicEmailToDetails{}
	return &this
}

// NewPublicEmailToDetailsWithDefaults instantiates a new PublicEmailToDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicEmailToDetailsWithDefaults() *PublicEmailToDetails {
	this := PublicEmailToDetails{}
	return &this
}

// GetContactIlsLists returns the ContactIlsLists field value if set, zero value otherwise.
func (o *PublicEmailToDetails) GetContactIlsLists() PublicEmailRecipients {
	if o == nil || IsNil(o.ContactIlsLists) {
		var ret PublicEmailRecipients
		return ret
	}
	return *o.ContactIlsLists
}

// GetContactIlsListsOk returns a tuple with the ContactIlsLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEmailToDetails) GetContactIlsListsOk() (*PublicEmailRecipients, bool) {
	if o == nil || IsNil(o.ContactIlsLists) {
		return nil, false
	}
	return o.ContactIlsLists, true
}

// HasContactIlsLists returns a boolean if a field has been set.
func (o *PublicEmailToDetails) HasContactIlsLists() bool {
	if o != nil && !IsNil(o.ContactIlsLists) {
		return true
	}

	return false
}

// SetContactIlsLists gets a reference to the given PublicEmailRecipients and assigns it to the ContactIlsLists field.
func (o *PublicEmailToDetails) SetContactIlsLists(v PublicEmailRecipients) {
	o.ContactIlsLists = &v
}

// GetLimitSendFrequency returns the LimitSendFrequency field value if set, zero value otherwise.
func (o *PublicEmailToDetails) GetLimitSendFrequency() bool {
	if o == nil || IsNil(o.LimitSendFrequency) {
		var ret bool
		return ret
	}
	return *o.LimitSendFrequency
}

// GetLimitSendFrequencyOk returns a tuple with the LimitSendFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEmailToDetails) GetLimitSendFrequencyOk() (*bool, bool) {
	if o == nil || IsNil(o.LimitSendFrequency) {
		return nil, false
	}
	return o.LimitSendFrequency, true
}

// HasLimitSendFrequency returns a boolean if a field has been set.
func (o *PublicEmailToDetails) HasLimitSendFrequency() bool {
	if o != nil && !IsNil(o.LimitSendFrequency) {
		return true
	}

	return false
}

// SetLimitSendFrequency gets a reference to the given bool and assigns it to the LimitSendFrequency field.
func (o *PublicEmailToDetails) SetLimitSendFrequency(v bool) {
	o.LimitSendFrequency = &v
}

// GetSuppressGraymail returns the SuppressGraymail field value if set, zero value otherwise.
func (o *PublicEmailToDetails) GetSuppressGraymail() bool {
	if o == nil || IsNil(o.SuppressGraymail) {
		var ret bool
		return ret
	}
	return *o.SuppressGraymail
}

// GetSuppressGraymailOk returns a tuple with the SuppressGraymail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEmailToDetails) GetSuppressGraymailOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressGraymail) {
		return nil, false
	}
	return o.SuppressGraymail, true
}

// HasSuppressGraymail returns a boolean if a field has been set.
func (o *PublicEmailToDetails) HasSuppressGraymail() bool {
	if o != nil && !IsNil(o.SuppressGraymail) {
		return true
	}

	return false
}

// SetSuppressGraymail gets a reference to the given bool and assigns it to the SuppressGraymail field.
func (o *PublicEmailToDetails) SetSuppressGraymail(v bool) {
	o.SuppressGraymail = &v
}

// GetContactIds returns the ContactIds field value if set, zero value otherwise.
func (o *PublicEmailToDetails) GetContactIds() PublicEmailRecipients {
	if o == nil || IsNil(o.ContactIds) {
		var ret PublicEmailRecipients
		return ret
	}
	return *o.ContactIds
}

// GetContactIdsOk returns a tuple with the ContactIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEmailToDetails) GetContactIdsOk() (*PublicEmailRecipients, bool) {
	if o == nil || IsNil(o.ContactIds) {
		return nil, false
	}
	return o.ContactIds, true
}

// HasContactIds returns a boolean if a field has been set.
func (o *PublicEmailToDetails) HasContactIds() bool {
	if o != nil && !IsNil(o.ContactIds) {
		return true
	}

	return false
}

// SetContactIds gets a reference to the given PublicEmailRecipients and assigns it to the ContactIds field.
func (o *PublicEmailToDetails) SetContactIds(v PublicEmailRecipients) {
	o.ContactIds = &v
}

// GetContactLists returns the ContactLists field value if set, zero value otherwise.
func (o *PublicEmailToDetails) GetContactLists() PublicEmailRecipients {
	if o == nil || IsNil(o.ContactLists) {
		var ret PublicEmailRecipients
		return ret
	}
	return *o.ContactLists
}

// GetContactListsOk returns a tuple with the ContactLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEmailToDetails) GetContactListsOk() (*PublicEmailRecipients, bool) {
	if o == nil || IsNil(o.ContactLists) {
		return nil, false
	}
	return o.ContactLists, true
}

// HasContactLists returns a boolean if a field has been set.
func (o *PublicEmailToDetails) HasContactLists() bool {
	if o != nil && !IsNil(o.ContactLists) {
		return true
	}

	return false
}

// SetContactLists gets a reference to the given PublicEmailRecipients and assigns it to the ContactLists field.
func (o *PublicEmailToDetails) SetContactLists(v PublicEmailRecipients) {
	o.ContactLists = &v
}

func (o PublicEmailToDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicEmailToDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactIlsLists) {
		toSerialize["contactIlsLists"] = o.ContactIlsLists
	}
	if !IsNil(o.LimitSendFrequency) {
		toSerialize["limitSendFrequency"] = o.LimitSendFrequency
	}
	if !IsNil(o.SuppressGraymail) {
		toSerialize["suppressGraymail"] = o.SuppressGraymail
	}
	if !IsNil(o.ContactIds) {
		toSerialize["contactIds"] = o.ContactIds
	}
	if !IsNil(o.ContactLists) {
		toSerialize["contactLists"] = o.ContactLists
	}
	return toSerialize, nil
}

type NullablePublicEmailToDetails struct {
	value *PublicEmailToDetails
	isSet bool
}

func (v NullablePublicEmailToDetails) Get() *PublicEmailToDetails {
	return v.value
}

func (v *NullablePublicEmailToDetails) Set(val *PublicEmailToDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicEmailToDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicEmailToDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicEmailToDetails(val *PublicEmailToDetails) *NullablePublicEmailToDetails {
	return &NullablePublicEmailToDetails{value: val, isSet: true}
}

func (v NullablePublicEmailToDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicEmailToDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


