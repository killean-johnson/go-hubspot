/*
Authors

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authors

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BatchInputBlogAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputBlogAuthor{}

// BatchInputBlogAuthor Wrapper for providing an array of blog authors as inputs.
type BatchInputBlogAuthor struct {
	// Blog authors to input.
	Inputs []BlogAuthor `json:"inputs"`
}

type _BatchInputBlogAuthor BatchInputBlogAuthor

// NewBatchInputBlogAuthor instantiates a new BatchInputBlogAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputBlogAuthor(inputs []BlogAuthor) *BatchInputBlogAuthor {
	this := BatchInputBlogAuthor{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputBlogAuthorWithDefaults instantiates a new BatchInputBlogAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputBlogAuthorWithDefaults() *BatchInputBlogAuthor {
	this := BatchInputBlogAuthor{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputBlogAuthor) GetInputs() []BlogAuthor {
	if o == nil {
		var ret []BlogAuthor
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputBlogAuthor) GetInputsOk() ([]BlogAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputBlogAuthor) SetInputs(v []BlogAuthor) {
	o.Inputs = v
}

func (o BatchInputBlogAuthor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputBlogAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchInputBlogAuthor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inputs",
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

	varBatchInputBlogAuthor := _BatchInputBlogAuthor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchInputBlogAuthor)

	if err != nil {
		return err
	}

	*o = BatchInputBlogAuthor(varBatchInputBlogAuthor)

	return err
}

type NullableBatchInputBlogAuthor struct {
	value *BatchInputBlogAuthor
	isSet bool
}

func (v NullableBatchInputBlogAuthor) Get() *BatchInputBlogAuthor {
	return v.value
}

func (v *NullableBatchInputBlogAuthor) Set(val *BatchInputBlogAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputBlogAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputBlogAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputBlogAuthor(val *BatchInputBlogAuthor) *NullableBatchInputBlogAuthor {
	return &NullableBatchInputBlogAuthor{value: val, isSet: true}
}

func (v NullableBatchInputBlogAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputBlogAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


