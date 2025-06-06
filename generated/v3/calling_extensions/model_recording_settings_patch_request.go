/*
CRM Calling Extensions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calling_extensions

import (
	"encoding/json"
)

// checks if the RecordingSettingsPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingSettingsPatchRequest{}

// RecordingSettingsPatchRequest struct for RecordingSettingsPatchRequest
type RecordingSettingsPatchRequest struct {
	UrlToRetrieveAuthedRecording *string `json:"urlToRetrieveAuthedRecording,omitempty"`
}

// NewRecordingSettingsPatchRequest instantiates a new RecordingSettingsPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingSettingsPatchRequest() *RecordingSettingsPatchRequest {
	this := RecordingSettingsPatchRequest{}
	return &this
}

// NewRecordingSettingsPatchRequestWithDefaults instantiates a new RecordingSettingsPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingSettingsPatchRequestWithDefaults() *RecordingSettingsPatchRequest {
	this := RecordingSettingsPatchRequest{}
	return &this
}

// GetUrlToRetrieveAuthedRecording returns the UrlToRetrieveAuthedRecording field value if set, zero value otherwise.
func (o *RecordingSettingsPatchRequest) GetUrlToRetrieveAuthedRecording() string {
	if o == nil || IsNil(o.UrlToRetrieveAuthedRecording) {
		var ret string
		return ret
	}
	return *o.UrlToRetrieveAuthedRecording
}

// GetUrlToRetrieveAuthedRecordingOk returns a tuple with the UrlToRetrieveAuthedRecording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingSettingsPatchRequest) GetUrlToRetrieveAuthedRecordingOk() (*string, bool) {
	if o == nil || IsNil(o.UrlToRetrieveAuthedRecording) {
		return nil, false
	}
	return o.UrlToRetrieveAuthedRecording, true
}

// HasUrlToRetrieveAuthedRecording returns a boolean if a field has been set.
func (o *RecordingSettingsPatchRequest) HasUrlToRetrieveAuthedRecording() bool {
	if o != nil && !IsNil(o.UrlToRetrieveAuthedRecording) {
		return true
	}

	return false
}

// SetUrlToRetrieveAuthedRecording gets a reference to the given string and assigns it to the UrlToRetrieveAuthedRecording field.
func (o *RecordingSettingsPatchRequest) SetUrlToRetrieveAuthedRecording(v string) {
	o.UrlToRetrieveAuthedRecording = &v
}

func (o RecordingSettingsPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingSettingsPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UrlToRetrieveAuthedRecording) {
		toSerialize["urlToRetrieveAuthedRecording"] = o.UrlToRetrieveAuthedRecording
	}
	return toSerialize, nil
}

type NullableRecordingSettingsPatchRequest struct {
	value *RecordingSettingsPatchRequest
	isSet bool
}

func (v NullableRecordingSettingsPatchRequest) Get() *RecordingSettingsPatchRequest {
	return v.value
}

func (v *NullableRecordingSettingsPatchRequest) Set(val *RecordingSettingsPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingSettingsPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingSettingsPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingSettingsPatchRequest(val *RecordingSettingsPatchRequest) *NullableRecordingSettingsPatchRequest {
	return &NullableRecordingSettingsPatchRequest{value: val, isSet: true}
}

func (v NullableRecordingSettingsPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingSettingsPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


