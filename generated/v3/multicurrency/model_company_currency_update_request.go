/*
Settings Multicurrency

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package multicurrency

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CompanyCurrencyUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyCurrencyUpdateRequest{}

// CompanyCurrencyUpdateRequest struct for CompanyCurrencyUpdateRequest
type CompanyCurrencyUpdateRequest struct {
	CurrencyCode string `json:"currencyCode"`
}

type _CompanyCurrencyUpdateRequest CompanyCurrencyUpdateRequest

// NewCompanyCurrencyUpdateRequest instantiates a new CompanyCurrencyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyCurrencyUpdateRequest(currencyCode string) *CompanyCurrencyUpdateRequest {
	this := CompanyCurrencyUpdateRequest{}
	this.CurrencyCode = currencyCode
	return &this
}

// NewCompanyCurrencyUpdateRequestWithDefaults instantiates a new CompanyCurrencyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyCurrencyUpdateRequestWithDefaults() *CompanyCurrencyUpdateRequest {
	this := CompanyCurrencyUpdateRequest{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *CompanyCurrencyUpdateRequest) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *CompanyCurrencyUpdateRequest) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *CompanyCurrencyUpdateRequest) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

func (o CompanyCurrencyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyCurrencyUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currencyCode"] = o.CurrencyCode
	return toSerialize, nil
}

func (o *CompanyCurrencyUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currencyCode",
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

	varCompanyCurrencyUpdateRequest := _CompanyCurrencyUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompanyCurrencyUpdateRequest)

	if err != nil {
		return err
	}

	*o = CompanyCurrencyUpdateRequest(varCompanyCurrencyUpdateRequest)

	return err
}

type NullableCompanyCurrencyUpdateRequest struct {
	value *CompanyCurrencyUpdateRequest
	isSet bool
}

func (v NullableCompanyCurrencyUpdateRequest) Get() *CompanyCurrencyUpdateRequest {
	return v.value
}

func (v *NullableCompanyCurrencyUpdateRequest) Set(val *CompanyCurrencyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyCurrencyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyCurrencyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyCurrencyUpdateRequest(val *CompanyCurrencyUpdateRequest) *NullableCompanyCurrencyUpdateRequest {
	return &NullableCompanyCurrencyUpdateRequest{value: val, isSet: true}
}

func (v NullableCompanyCurrencyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyCurrencyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


