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

// checks if the OEmbedDomainsCollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OEmbedDomainsCollectionResponse{}

// OEmbedDomainsCollectionResponse struct for OEmbedDomainsCollectionResponse
type OEmbedDomainsCollectionResponse struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []IntegratorOEmbedDomainModel `json:"results"`
}

type _OEmbedDomainsCollectionResponse OEmbedDomainsCollectionResponse

// NewOEmbedDomainsCollectionResponse instantiates a new OEmbedDomainsCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOEmbedDomainsCollectionResponse(results []IntegratorOEmbedDomainModel) *OEmbedDomainsCollectionResponse {
	this := OEmbedDomainsCollectionResponse{}
	this.Results = results
	return &this
}

// NewOEmbedDomainsCollectionResponseWithDefaults instantiates a new OEmbedDomainsCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOEmbedDomainsCollectionResponseWithDefaults() *OEmbedDomainsCollectionResponse {
	this := OEmbedDomainsCollectionResponse{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *OEmbedDomainsCollectionResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OEmbedDomainsCollectionResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *OEmbedDomainsCollectionResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *OEmbedDomainsCollectionResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value
func (o *OEmbedDomainsCollectionResponse) GetResults() []IntegratorOEmbedDomainModel {
	if o == nil {
		var ret []IntegratorOEmbedDomainModel
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *OEmbedDomainsCollectionResponse) GetResultsOk() ([]IntegratorOEmbedDomainModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *OEmbedDomainsCollectionResponse) SetResults(v []IntegratorOEmbedDomainModel) {
	o.Results = v
}

func (o OEmbedDomainsCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OEmbedDomainsCollectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *OEmbedDomainsCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varOEmbedDomainsCollectionResponse := _OEmbedDomainsCollectionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOEmbedDomainsCollectionResponse)

	if err != nil {
		return err
	}

	*o = OEmbedDomainsCollectionResponse(varOEmbedDomainsCollectionResponse)

	return err
}

type NullableOEmbedDomainsCollectionResponse struct {
	value *OEmbedDomainsCollectionResponse
	isSet bool
}

func (v NullableOEmbedDomainsCollectionResponse) Get() *OEmbedDomainsCollectionResponse {
	return v.value
}

func (v *NullableOEmbedDomainsCollectionResponse) Set(val *OEmbedDomainsCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOEmbedDomainsCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOEmbedDomainsCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOEmbedDomainsCollectionResponse(val *OEmbedDomainsCollectionResponse) *NullableOEmbedDomainsCollectionResponse {
	return &NullableOEmbedDomainsCollectionResponse{value: val, isSet: true}
}

func (v NullableOEmbedDomainsCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOEmbedDomainsCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


