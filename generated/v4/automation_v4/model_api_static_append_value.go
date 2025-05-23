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

// checks if the ApiStaticAppendValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStaticAppendValue{}

// ApiStaticAppendValue struct for ApiStaticAppendValue
type ApiStaticAppendValue struct {
	// The value to append
	StaticAppendValue string `json:"staticAppendValue"`
	// This is the type of input value. This can be one of: \"FIELD_DATA\", \"OBJECT_PROPERTY\", \"STATIC_VALUE\", \"RELATIVE_DATETIME\", \"TIMESTAMP\", \"INCREMENT\", \"FETCHED_OBJECT_PROPERTY\", \"APPEND_OBJECT_PROPERTY\", \"STATIC_APPEND_VALUE\", \"ENROLLMENT_EVENT_PROPERTY\"
	Type string `json:"type"`
}

type _ApiStaticAppendValue ApiStaticAppendValue

// NewApiStaticAppendValue instantiates a new ApiStaticAppendValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStaticAppendValue(staticAppendValue string, type_ string) *ApiStaticAppendValue {
	this := ApiStaticAppendValue{}
	this.StaticAppendValue = staticAppendValue
	this.Type = type_
	return &this
}

// NewApiStaticAppendValueWithDefaults instantiates a new ApiStaticAppendValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStaticAppendValueWithDefaults() *ApiStaticAppendValue {
	this := ApiStaticAppendValue{}
	var type_ string = "STATIC_APPEND_VALUE"
	this.Type = type_
	return &this
}

// GetStaticAppendValue returns the StaticAppendValue field value
func (o *ApiStaticAppendValue) GetStaticAppendValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StaticAppendValue
}

// GetStaticAppendValueOk returns a tuple with the StaticAppendValue field value
// and a boolean to check if the value has been set.
func (o *ApiStaticAppendValue) GetStaticAppendValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StaticAppendValue, true
}

// SetStaticAppendValue sets field value
func (o *ApiStaticAppendValue) SetStaticAppendValue(v string) {
	o.StaticAppendValue = v
}

// GetType returns the Type field value
func (o *ApiStaticAppendValue) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiStaticAppendValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiStaticAppendValue) SetType(v string) {
	o.Type = v
}

func (o ApiStaticAppendValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStaticAppendValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["staticAppendValue"] = o.StaticAppendValue
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ApiStaticAppendValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"staticAppendValue",
		"type",
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

	varApiStaticAppendValue := _ApiStaticAppendValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiStaticAppendValue)

	if err != nil {
		return err
	}

	*o = ApiStaticAppendValue(varApiStaticAppendValue)

	return err
}

type NullableApiStaticAppendValue struct {
	value *ApiStaticAppendValue
	isSet bool
}

func (v NullableApiStaticAppendValue) Get() *ApiStaticAppendValue {
	return v.value
}

func (v *NullableApiStaticAppendValue) Set(val *ApiStaticAppendValue) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStaticAppendValue) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStaticAppendValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStaticAppendValue(val *ApiStaticAppendValue) *NullableApiStaticAppendValue {
	return &NullableApiStaticAppendValue{value: val, isSet: true}
}

func (v NullableApiStaticAppendValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStaticAppendValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


