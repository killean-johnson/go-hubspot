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

// checks if the ApiFlowBatchFetchMigrationFlowIdCoordinate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiFlowBatchFetchMigrationFlowIdCoordinate{}

// ApiFlowBatchFetchMigrationFlowIdCoordinate struct for ApiFlowBatchFetchMigrationFlowIdCoordinate
type ApiFlowBatchFetchMigrationFlowIdCoordinate struct {
	FlowMigrationStatuses string `json:"flowMigrationStatuses"`
	Type string `json:"type"`
}

type _ApiFlowBatchFetchMigrationFlowIdCoordinate ApiFlowBatchFetchMigrationFlowIdCoordinate

// NewApiFlowBatchFetchMigrationFlowIdCoordinate instantiates a new ApiFlowBatchFetchMigrationFlowIdCoordinate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFlowBatchFetchMigrationFlowIdCoordinate(flowMigrationStatuses string, type_ string) *ApiFlowBatchFetchMigrationFlowIdCoordinate {
	this := ApiFlowBatchFetchMigrationFlowIdCoordinate{}
	this.FlowMigrationStatuses = flowMigrationStatuses
	this.Type = type_
	return &this
}

// NewApiFlowBatchFetchMigrationFlowIdCoordinateWithDefaults instantiates a new ApiFlowBatchFetchMigrationFlowIdCoordinate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFlowBatchFetchMigrationFlowIdCoordinateWithDefaults() *ApiFlowBatchFetchMigrationFlowIdCoordinate {
	this := ApiFlowBatchFetchMigrationFlowIdCoordinate{}
	var type_ string = "FLOW_ID"
	this.Type = type_
	return &this
}

// GetFlowMigrationStatuses returns the FlowMigrationStatuses field value
func (o *ApiFlowBatchFetchMigrationFlowIdCoordinate) GetFlowMigrationStatuses() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlowMigrationStatuses
}

// GetFlowMigrationStatusesOk returns a tuple with the FlowMigrationStatuses field value
// and a boolean to check if the value has been set.
func (o *ApiFlowBatchFetchMigrationFlowIdCoordinate) GetFlowMigrationStatusesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlowMigrationStatuses, true
}

// SetFlowMigrationStatuses sets field value
func (o *ApiFlowBatchFetchMigrationFlowIdCoordinate) SetFlowMigrationStatuses(v string) {
	o.FlowMigrationStatuses = v
}

// GetType returns the Type field value
func (o *ApiFlowBatchFetchMigrationFlowIdCoordinate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiFlowBatchFetchMigrationFlowIdCoordinate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiFlowBatchFetchMigrationFlowIdCoordinate) SetType(v string) {
	o.Type = v
}

func (o ApiFlowBatchFetchMigrationFlowIdCoordinate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiFlowBatchFetchMigrationFlowIdCoordinate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["flowMigrationStatuses"] = o.FlowMigrationStatuses
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ApiFlowBatchFetchMigrationFlowIdCoordinate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"flowMigrationStatuses",
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

	varApiFlowBatchFetchMigrationFlowIdCoordinate := _ApiFlowBatchFetchMigrationFlowIdCoordinate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiFlowBatchFetchMigrationFlowIdCoordinate)

	if err != nil {
		return err
	}

	*o = ApiFlowBatchFetchMigrationFlowIdCoordinate(varApiFlowBatchFetchMigrationFlowIdCoordinate)

	return err
}

type NullableApiFlowBatchFetchMigrationFlowIdCoordinate struct {
	value *ApiFlowBatchFetchMigrationFlowIdCoordinate
	isSet bool
}

func (v NullableApiFlowBatchFetchMigrationFlowIdCoordinate) Get() *ApiFlowBatchFetchMigrationFlowIdCoordinate {
	return v.value
}

func (v *NullableApiFlowBatchFetchMigrationFlowIdCoordinate) Set(val *ApiFlowBatchFetchMigrationFlowIdCoordinate) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFlowBatchFetchMigrationFlowIdCoordinate) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFlowBatchFetchMigrationFlowIdCoordinate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFlowBatchFetchMigrationFlowIdCoordinate(val *ApiFlowBatchFetchMigrationFlowIdCoordinate) *NullableApiFlowBatchFetchMigrationFlowIdCoordinate {
	return &NullableApiFlowBatchFetchMigrationFlowIdCoordinate{value: val, isSet: true}
}

func (v NullableApiFlowBatchFetchMigrationFlowIdCoordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFlowBatchFetchMigrationFlowIdCoordinate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


