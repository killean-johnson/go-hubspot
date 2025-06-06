/*
CRM Calling Extensions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calling_extensions

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RecordingSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingSettingsRequest{}

// RecordingSettingsRequest struct for RecordingSettingsRequest
type RecordingSettingsRequest struct {
	UrlToRetrieveAuthedRecording string `json:"urlToRetrieveAuthedRecording"`
}

type _RecordingSettingsRequest RecordingSettingsRequest

// NewRecordingSettingsRequest instantiates a new RecordingSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingSettingsRequest(urlToRetrieveAuthedRecording string) *RecordingSettingsRequest {
	this := RecordingSettingsRequest{}
	this.UrlToRetrieveAuthedRecording = urlToRetrieveAuthedRecording
	return &this
}

// NewRecordingSettingsRequestWithDefaults instantiates a new RecordingSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingSettingsRequestWithDefaults() *RecordingSettingsRequest {
	this := RecordingSettingsRequest{}
	return &this
}

// GetUrlToRetrieveAuthedRecording returns the UrlToRetrieveAuthedRecording field value
func (o *RecordingSettingsRequest) GetUrlToRetrieveAuthedRecording() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlToRetrieveAuthedRecording
}

// GetUrlToRetrieveAuthedRecordingOk returns a tuple with the UrlToRetrieveAuthedRecording field value
// and a boolean to check if the value has been set.
func (o *RecordingSettingsRequest) GetUrlToRetrieveAuthedRecordingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlToRetrieveAuthedRecording, true
}

// SetUrlToRetrieveAuthedRecording sets field value
func (o *RecordingSettingsRequest) SetUrlToRetrieveAuthedRecording(v string) {
	o.UrlToRetrieveAuthedRecording = v
}

func (o RecordingSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["urlToRetrieveAuthedRecording"] = o.UrlToRetrieveAuthedRecording
	return toSerialize, nil
}

func (o *RecordingSettingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"urlToRetrieveAuthedRecording",
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

	varRecordingSettingsRequest := _RecordingSettingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecordingSettingsRequest)

	if err != nil {
		return err
	}

	*o = RecordingSettingsRequest(varRecordingSettingsRequest)

	return err
}

type NullableRecordingSettingsRequest struct {
	value *RecordingSettingsRequest
	isSet bool
}

func (v NullableRecordingSettingsRequest) Get() *RecordingSettingsRequest {
	return v.value
}

func (v *NullableRecordingSettingsRequest) Set(val *RecordingSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingSettingsRequest(val *RecordingSettingsRequest) *NullableRecordingSettingsRequest {
	return &NullableRecordingSettingsRequest{value: val, isSet: true}
}

func (v NullableRecordingSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


