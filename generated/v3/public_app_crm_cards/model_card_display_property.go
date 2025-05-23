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

// checks if the CardDisplayProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardDisplayProperty{}

// CardDisplayProperty Definition for a card display property.
type CardDisplayProperty struct {
	// Type of data represented by this property.
	DataType string `json:"dataType"`
	// An internal identifier for this property. This value must be unique TODO.
	Name string `json:"name"`
	// An array of available options that can be displayed. Only used in when `dataType` is `STATUS`.
	Options []DisplayOption `json:"options"`
	// The label for this property as you'd like it displayed to users.
	Label string `json:"label"`
}

type _CardDisplayProperty CardDisplayProperty

// NewCardDisplayProperty instantiates a new CardDisplayProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardDisplayProperty(dataType string, name string, options []DisplayOption, label string) *CardDisplayProperty {
	this := CardDisplayProperty{}
	this.DataType = dataType
	this.Name = name
	this.Options = options
	this.Label = label
	return &this
}

// NewCardDisplayPropertyWithDefaults instantiates a new CardDisplayProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardDisplayPropertyWithDefaults() *CardDisplayProperty {
	this := CardDisplayProperty{}
	return &this
}

// GetDataType returns the DataType field value
func (o *CardDisplayProperty) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *CardDisplayProperty) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *CardDisplayProperty) SetDataType(v string) {
	o.DataType = v
}

// GetName returns the Name field value
func (o *CardDisplayProperty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CardDisplayProperty) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CardDisplayProperty) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *CardDisplayProperty) GetOptions() []DisplayOption {
	if o == nil {
		var ret []DisplayOption
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CardDisplayProperty) GetOptionsOk() ([]DisplayOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *CardDisplayProperty) SetOptions(v []DisplayOption) {
	o.Options = v
}

// GetLabel returns the Label field value
func (o *CardDisplayProperty) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CardDisplayProperty) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CardDisplayProperty) SetLabel(v string) {
	o.Label = v
}

func (o CardDisplayProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardDisplayProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataType"] = o.DataType
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

func (o *CardDisplayProperty) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dataType",
		"name",
		"options",
		"label",
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

	varCardDisplayProperty := _CardDisplayProperty{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCardDisplayProperty)

	if err != nil {
		return err
	}

	*o = CardDisplayProperty(varCardDisplayProperty)

	return err
}

type NullableCardDisplayProperty struct {
	value *CardDisplayProperty
	isSet bool
}

func (v NullableCardDisplayProperty) Get() *CardDisplayProperty {
	return v.value
}

func (v *NullableCardDisplayProperty) Set(val *CardDisplayProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableCardDisplayProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableCardDisplayProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardDisplayProperty(val *CardDisplayProperty) *NullableCardDisplayProperty {
	return &NullableCardDisplayProperty{value: val, isSet: true}
}

func (v NullableCardDisplayProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardDisplayProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


