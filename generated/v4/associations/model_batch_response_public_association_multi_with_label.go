/*
Associations

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package associations

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the BatchResponsePublicAssociationMultiWithLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchResponsePublicAssociationMultiWithLabel{}

// BatchResponsePublicAssociationMultiWithLabel struct for BatchResponsePublicAssociationMultiWithLabel
type BatchResponsePublicAssociationMultiWithLabel struct {
	CompletedAt time.Time `json:"completedAt"`
	RequestedAt *time.Time `json:"requestedAt,omitempty"`
	StartedAt time.Time `json:"startedAt"`
	Links *map[string]string `json:"links,omitempty"`
	Results []PublicAssociationMultiWithLabel `json:"results"`
	Status string `json:"status"`
}

type _BatchResponsePublicAssociationMultiWithLabel BatchResponsePublicAssociationMultiWithLabel

// NewBatchResponsePublicAssociationMultiWithLabel instantiates a new BatchResponsePublicAssociationMultiWithLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponsePublicAssociationMultiWithLabel(completedAt time.Time, startedAt time.Time, results []PublicAssociationMultiWithLabel, status string) *BatchResponsePublicAssociationMultiWithLabel {
	this := BatchResponsePublicAssociationMultiWithLabel{}
	this.CompletedAt = completedAt
	this.StartedAt = startedAt
	this.Results = results
	this.Status = status
	return &this
}

// NewBatchResponsePublicAssociationMultiWithLabelWithDefaults instantiates a new BatchResponsePublicAssociationMultiWithLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponsePublicAssociationMultiWithLabelWithDefaults() *BatchResponsePublicAssociationMultiWithLabel {
	this := BatchResponsePublicAssociationMultiWithLabel{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponsePublicAssociationMultiWithLabel) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponsePublicAssociationMultiWithLabel) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponsePublicAssociationMultiWithLabel) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponsePublicAssociationMultiWithLabel) GetRequestedAt() time.Time {
	if o == nil || IsNil(o.RequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponsePublicAssociationMultiWithLabel) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestedAt) {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponsePublicAssociationMultiWithLabel) HasRequestedAt() bool {
	if o != nil && !IsNil(o.RequestedAt) {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponsePublicAssociationMultiWithLabel) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponsePublicAssociationMultiWithLabel) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponsePublicAssociationMultiWithLabel) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponsePublicAssociationMultiWithLabel) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponsePublicAssociationMultiWithLabel) GetLinks() map[string]string {
	if o == nil || IsNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponsePublicAssociationMultiWithLabel) GetLinksOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponsePublicAssociationMultiWithLabel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponsePublicAssociationMultiWithLabel) SetLinks(v map[string]string) {
	o.Links = &v
}

// GetResults returns the Results field value
func (o *BatchResponsePublicAssociationMultiWithLabel) GetResults() []PublicAssociationMultiWithLabel {
	if o == nil {
		var ret []PublicAssociationMultiWithLabel
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponsePublicAssociationMultiWithLabel) GetResultsOk() ([]PublicAssociationMultiWithLabel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponsePublicAssociationMultiWithLabel) SetResults(v []PublicAssociationMultiWithLabel) {
	o.Results = v
}

// GetStatus returns the Status field value
func (o *BatchResponsePublicAssociationMultiWithLabel) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponsePublicAssociationMultiWithLabel) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponsePublicAssociationMultiWithLabel) SetStatus(v string) {
	o.Status = v
}

func (o BatchResponsePublicAssociationMultiWithLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchResponsePublicAssociationMultiWithLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completedAt"] = o.CompletedAt
	if !IsNil(o.RequestedAt) {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	toSerialize["startedAt"] = o.StartedAt
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["results"] = o.Results
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *BatchResponsePublicAssociationMultiWithLabel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completedAt",
		"startedAt",
		"results",
		"status",
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

	varBatchResponsePublicAssociationMultiWithLabel := _BatchResponsePublicAssociationMultiWithLabel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchResponsePublicAssociationMultiWithLabel)

	if err != nil {
		return err
	}

	*o = BatchResponsePublicAssociationMultiWithLabel(varBatchResponsePublicAssociationMultiWithLabel)

	return err
}

type NullableBatchResponsePublicAssociationMultiWithLabel struct {
	value *BatchResponsePublicAssociationMultiWithLabel
	isSet bool
}

func (v NullableBatchResponsePublicAssociationMultiWithLabel) Get() *BatchResponsePublicAssociationMultiWithLabel {
	return v.value
}

func (v *NullableBatchResponsePublicAssociationMultiWithLabel) Set(val *BatchResponsePublicAssociationMultiWithLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponsePublicAssociationMultiWithLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponsePublicAssociationMultiWithLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponsePublicAssociationMultiWithLabel(val *BatchResponsePublicAssociationMultiWithLabel) *NullableBatchResponsePublicAssociationMultiWithLabel {
	return &NullableBatchResponsePublicAssociationMultiWithLabel{value: val, isSet: true}
}

func (v NullableBatchResponsePublicAssociationMultiWithLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponsePublicAssociationMultiWithLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


