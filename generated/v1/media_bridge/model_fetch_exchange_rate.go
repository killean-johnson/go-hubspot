/*
CMS Media Bridge

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package media_bridge

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FetchExchangeRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchExchangeRate{}

// FetchExchangeRate struct for FetchExchangeRate
type FetchExchangeRate struct {
	PropertyName *string `json:"propertyName,omitempty"`
	Inputs []OrInputsInner `json:"inputs,omitempty"`
	Value *float32 `json:"value,omitempty"`
	Operator string `json:"operator"`
}

type _FetchExchangeRate FetchExchangeRate

// NewFetchExchangeRate instantiates a new FetchExchangeRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchExchangeRate(operator string) *FetchExchangeRate {
	this := FetchExchangeRate{}
	this.Operator = operator
	return &this
}

// NewFetchExchangeRateWithDefaults instantiates a new FetchExchangeRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchExchangeRateWithDefaults() *FetchExchangeRate {
	this := FetchExchangeRate{}
	var operator string = "FETCH_EXCHANGE_RATE"
	this.Operator = operator
	return &this
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *FetchExchangeRate) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchExchangeRate) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *FetchExchangeRate) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *FetchExchangeRate) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *FetchExchangeRate) GetInputs() []OrInputsInner {
	if o == nil || IsNil(o.Inputs) {
		var ret []OrInputsInner
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchExchangeRate) GetInputsOk() ([]OrInputsInner, bool) {
	if o == nil || IsNil(o.Inputs) {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *FetchExchangeRate) HasInputs() bool {
	if o != nil && !IsNil(o.Inputs) {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []OrInputsInner and assigns it to the Inputs field.
func (o *FetchExchangeRate) SetInputs(v []OrInputsInner) {
	o.Inputs = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FetchExchangeRate) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchExchangeRate) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FetchExchangeRate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *FetchExchangeRate) SetValue(v float32) {
	o.Value = &v
}

// GetOperator returns the Operator field value
func (o *FetchExchangeRate) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *FetchExchangeRate) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *FetchExchangeRate) SetOperator(v string) {
	o.Operator = v
}

func (o FetchExchangeRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchExchangeRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyName) {
		toSerialize["propertyName"] = o.PropertyName
	}
	if !IsNil(o.Inputs) {
		toSerialize["inputs"] = o.Inputs
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *FetchExchangeRate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operator",
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

	varFetchExchangeRate := _FetchExchangeRate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFetchExchangeRate)

	if err != nil {
		return err
	}

	*o = FetchExchangeRate(varFetchExchangeRate)

	return err
}

type NullableFetchExchangeRate struct {
	value *FetchExchangeRate
	isSet bool
}

func (v NullableFetchExchangeRate) Get() *FetchExchangeRate {
	return v.value
}

func (v *NullableFetchExchangeRate) Set(val *FetchExchangeRate) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchExchangeRate) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchExchangeRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchExchangeRate(val *FetchExchangeRate) *NullableFetchExchangeRate {
	return &NullableFetchExchangeRate{value: val, isSet: true}
}

func (v NullableFetchExchangeRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchExchangeRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


