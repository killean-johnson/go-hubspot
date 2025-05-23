/*
Marketing Emails

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_emails

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CollectionResponseWithTotalPublicEmailForwardPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseWithTotalPublicEmailForwardPaging{}

// CollectionResponseWithTotalPublicEmailForwardPaging Response object for collections of marketing emails with pagination information.
type CollectionResponseWithTotalPublicEmailForwardPaging struct {
	// Total number of content emails.
	Total int32 `json:"total"`
	Paging *ForwardPaging `json:"paging,omitempty"`
	// Collection of emails.
	Results []PublicEmail `json:"results"`
}

type _CollectionResponseWithTotalPublicEmailForwardPaging CollectionResponseWithTotalPublicEmailForwardPaging

// NewCollectionResponseWithTotalPublicEmailForwardPaging instantiates a new CollectionResponseWithTotalPublicEmailForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseWithTotalPublicEmailForwardPaging(total int32, results []PublicEmail) *CollectionResponseWithTotalPublicEmailForwardPaging {
	this := CollectionResponseWithTotalPublicEmailForwardPaging{}
	this.Total = total
	this.Results = results
	return &this
}

// NewCollectionResponseWithTotalPublicEmailForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalPublicEmailForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseWithTotalPublicEmailForwardPagingWithDefaults() *CollectionResponseWithTotalPublicEmailForwardPaging {
	this := CollectionResponseWithTotalPublicEmailForwardPaging{}
	return &this
}

// GetTotal returns the Total field value
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) SetTotal(v int32) {
	o.Total = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) GetPaging() ForwardPaging {
	if o == nil || IsNil(o.Paging) {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

// GetResults returns the Results field value
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) GetResults() []PublicEmail {
	if o == nil {
		var ret []PublicEmail
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) GetResultsOk() ([]PublicEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseWithTotalPublicEmailForwardPaging) SetResults(v []PublicEmail) {
	o.Results = v
}

func (o CollectionResponseWithTotalPublicEmailForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseWithTotalPublicEmailForwardPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *CollectionResponseWithTotalPublicEmailForwardPaging) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
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

	varCollectionResponseWithTotalPublicEmailForwardPaging := _CollectionResponseWithTotalPublicEmailForwardPaging{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionResponseWithTotalPublicEmailForwardPaging)

	if err != nil {
		return err
	}

	*o = CollectionResponseWithTotalPublicEmailForwardPaging(varCollectionResponseWithTotalPublicEmailForwardPaging)

	return err
}

type NullableCollectionResponseWithTotalPublicEmailForwardPaging struct {
	value *CollectionResponseWithTotalPublicEmailForwardPaging
	isSet bool
}

func (v NullableCollectionResponseWithTotalPublicEmailForwardPaging) Get() *CollectionResponseWithTotalPublicEmailForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseWithTotalPublicEmailForwardPaging) Set(val *CollectionResponseWithTotalPublicEmailForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseWithTotalPublicEmailForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseWithTotalPublicEmailForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseWithTotalPublicEmailForwardPaging(val *CollectionResponseWithTotalPublicEmailForwardPaging) *NullableCollectionResponseWithTotalPublicEmailForwardPaging {
	return &NullableCollectionResponseWithTotalPublicEmailForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseWithTotalPublicEmailForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseWithTotalPublicEmailForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


