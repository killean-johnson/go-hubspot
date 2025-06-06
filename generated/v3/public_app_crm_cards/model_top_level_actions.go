/*
Public App Crm Cards

Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public_app_crm_cards

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TopLevelActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopLevelActions{}

// TopLevelActions struct for TopLevelActions
type TopLevelActions struct {
	Secondary []IntegratorObjectResultActionsInner `json:"secondary"`
	Settings *IFrameActionBody `json:"settings,omitempty"`
	Primary *IntegratorObjectResultActionsInner `json:"primary,omitempty"`
}

type _TopLevelActions TopLevelActions

// NewTopLevelActions instantiates a new TopLevelActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopLevelActions(secondary []IntegratorObjectResultActionsInner) *TopLevelActions {
	this := TopLevelActions{}
	this.Secondary = secondary
	return &this
}

// NewTopLevelActionsWithDefaults instantiates a new TopLevelActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopLevelActionsWithDefaults() *TopLevelActions {
	this := TopLevelActions{}
	return &this
}

// GetSecondary returns the Secondary field value
func (o *TopLevelActions) GetSecondary() []IntegratorObjectResultActionsInner {
	if o == nil {
		var ret []IntegratorObjectResultActionsInner
		return ret
	}

	return o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value
// and a boolean to check if the value has been set.
func (o *TopLevelActions) GetSecondaryOk() ([]IntegratorObjectResultActionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secondary, true
}

// SetSecondary sets field value
func (o *TopLevelActions) SetSecondary(v []IntegratorObjectResultActionsInner) {
	o.Secondary = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *TopLevelActions) GetSettings() IFrameActionBody {
	if o == nil || IsNil(o.Settings) {
		var ret IFrameActionBody
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopLevelActions) GetSettingsOk() (*IFrameActionBody, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *TopLevelActions) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given IFrameActionBody and assigns it to the Settings field.
func (o *TopLevelActions) SetSettings(v IFrameActionBody) {
	o.Settings = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *TopLevelActions) GetPrimary() IntegratorObjectResultActionsInner {
	if o == nil || IsNil(o.Primary) {
		var ret IntegratorObjectResultActionsInner
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopLevelActions) GetPrimaryOk() (*IntegratorObjectResultActionsInner, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *TopLevelActions) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given IntegratorObjectResultActionsInner and assigns it to the Primary field.
func (o *TopLevelActions) SetPrimary(v IntegratorObjectResultActionsInner) {
	o.Primary = &v
}

func (o TopLevelActions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopLevelActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secondary"] = o.Secondary
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	return toSerialize, nil
}

func (o *TopLevelActions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"secondary",
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

	varTopLevelActions := _TopLevelActions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTopLevelActions)

	if err != nil {
		return err
	}

	*o = TopLevelActions(varTopLevelActions)

	return err
}

type NullableTopLevelActions struct {
	value *TopLevelActions
	isSet bool
}

func (v NullableTopLevelActions) Get() *TopLevelActions {
	return v.value
}

func (v *NullableTopLevelActions) Set(val *TopLevelActions) {
	v.value = val
	v.isSet = true
}

func (v NullableTopLevelActions) IsSet() bool {
	return v.isSet
}

func (v *NullableTopLevelActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopLevelActions(val *TopLevelActions) *NullableTopLevelActions {
	return &NullableTopLevelActions{value: val, isSet: true}
}

func (v NullableTopLevelActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopLevelActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


