/*
Automation V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package automation_v4

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiSort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSort{}

// ApiSort struct for ApiSort
type ApiSort struct {
	Property string `json:"property"`
	Missing *string `json:"missing,omitempty"`
	Order string `json:"order"`
}

type _ApiSort ApiSort

// NewApiSort instantiates a new ApiSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSort(property string, order string) *ApiSort {
	this := ApiSort{}
	this.Property = property
	this.Order = order
	return &this
}

// NewApiSortWithDefaults instantiates a new ApiSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSortWithDefaults() *ApiSort {
	this := ApiSort{}
	return &this
}

// GetProperty returns the Property field value
func (o *ApiSort) GetProperty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *ApiSort) GetPropertyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *ApiSort) SetProperty(v string) {
	o.Property = v
}

// GetMissing returns the Missing field value if set, zero value otherwise.
func (o *ApiSort) GetMissing() string {
	if o == nil || IsNil(o.Missing) {
		var ret string
		return ret
	}
	return *o.Missing
}

// GetMissingOk returns a tuple with the Missing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSort) GetMissingOk() (*string, bool) {
	if o == nil || IsNil(o.Missing) {
		return nil, false
	}
	return o.Missing, true
}

// HasMissing returns a boolean if a field has been set.
func (o *ApiSort) HasMissing() bool {
	if o != nil && !IsNil(o.Missing) {
		return true
	}

	return false
}

// SetMissing gets a reference to the given string and assigns it to the Missing field.
func (o *ApiSort) SetMissing(v string) {
	o.Missing = &v
}

// GetOrder returns the Order field value
func (o *ApiSort) GetOrder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ApiSort) GetOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ApiSort) SetOrder(v string) {
	o.Order = v
}

func (o ApiSort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["property"] = o.Property
	if !IsNil(o.Missing) {
		toSerialize["missing"] = o.Missing
	}
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

func (o *ApiSort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"property",
		"order",
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

	varApiSort := _ApiSort{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiSort)

	if err != nil {
		return err
	}

	*o = ApiSort(varApiSort)

	return err
}

type NullableApiSort struct {
	value *ApiSort
	isSet bool
}

func (v NullableApiSort) Get() *ApiSort {
	return v.value
}

func (v *NullableApiSort) Set(val *ApiSort) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSort) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSort(val *ApiSort) *NullableApiSort {
	return &NullableApiSort{value: val, isSet: true}
}

func (v NullableApiSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


