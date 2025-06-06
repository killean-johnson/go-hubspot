/*
Marketing Events

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BatchInputMarketingEventPublicObjectIdDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputMarketingEventPublicObjectIdDeleteRequest{}

// BatchInputMarketingEventPublicObjectIdDeleteRequest struct for BatchInputMarketingEventPublicObjectIdDeleteRequest
type BatchInputMarketingEventPublicObjectIdDeleteRequest struct {
	Inputs []MarketingEventPublicObjectIdDeleteRequest `json:"inputs"`
}

type _BatchInputMarketingEventPublicObjectIdDeleteRequest BatchInputMarketingEventPublicObjectIdDeleteRequest

// NewBatchInputMarketingEventPublicObjectIdDeleteRequest instantiates a new BatchInputMarketingEventPublicObjectIdDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputMarketingEventPublicObjectIdDeleteRequest(inputs []MarketingEventPublicObjectIdDeleteRequest) *BatchInputMarketingEventPublicObjectIdDeleteRequest {
	this := BatchInputMarketingEventPublicObjectIdDeleteRequest{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputMarketingEventPublicObjectIdDeleteRequestWithDefaults instantiates a new BatchInputMarketingEventPublicObjectIdDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputMarketingEventPublicObjectIdDeleteRequestWithDefaults() *BatchInputMarketingEventPublicObjectIdDeleteRequest {
	this := BatchInputMarketingEventPublicObjectIdDeleteRequest{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputMarketingEventPublicObjectIdDeleteRequest) GetInputs() []MarketingEventPublicObjectIdDeleteRequest {
	if o == nil {
		var ret []MarketingEventPublicObjectIdDeleteRequest
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputMarketingEventPublicObjectIdDeleteRequest) GetInputsOk() ([]MarketingEventPublicObjectIdDeleteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputMarketingEventPublicObjectIdDeleteRequest) SetInputs(v []MarketingEventPublicObjectIdDeleteRequest) {
	o.Inputs = v
}

func (o BatchInputMarketingEventPublicObjectIdDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputMarketingEventPublicObjectIdDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchInputMarketingEventPublicObjectIdDeleteRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBatchInputMarketingEventPublicObjectIdDeleteRequest := _BatchInputMarketingEventPublicObjectIdDeleteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchInputMarketingEventPublicObjectIdDeleteRequest)

	if err != nil {
		return err
	}

	*o = BatchInputMarketingEventPublicObjectIdDeleteRequest(varBatchInputMarketingEventPublicObjectIdDeleteRequest)

	return err
}

type NullableBatchInputMarketingEventPublicObjectIdDeleteRequest struct {
	value *BatchInputMarketingEventPublicObjectIdDeleteRequest
	isSet bool
}

func (v NullableBatchInputMarketingEventPublicObjectIdDeleteRequest) Get() *BatchInputMarketingEventPublicObjectIdDeleteRequest {
	return v.value
}

func (v *NullableBatchInputMarketingEventPublicObjectIdDeleteRequest) Set(val *BatchInputMarketingEventPublicObjectIdDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputMarketingEventPublicObjectIdDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputMarketingEventPublicObjectIdDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputMarketingEventPublicObjectIdDeleteRequest(val *BatchInputMarketingEventPublicObjectIdDeleteRequest) *NullableBatchInputMarketingEventPublicObjectIdDeleteRequest {
	return &NullableBatchInputMarketingEventPublicObjectIdDeleteRequest{value: val, isSet: true}
}

func (v NullableBatchInputMarketingEventPublicObjectIdDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputMarketingEventPublicObjectIdDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


