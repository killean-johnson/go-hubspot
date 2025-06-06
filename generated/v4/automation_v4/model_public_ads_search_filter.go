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

// checks if the PublicAdsSearchFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicAdsSearchFilter{}

// PublicAdsSearchFilter struct for PublicAdsSearchFilter
type PublicAdsSearchFilter struct {
	SearchTerms []string `json:"searchTerms"`
	EntityType string `json:"entityType"`
	AdNetwork string `json:"adNetwork"`
	SearchTermType string `json:"searchTermType"`
	FilterType string `json:"filterType"`
	Operator string `json:"operator"`
}

type _PublicAdsSearchFilter PublicAdsSearchFilter

// NewPublicAdsSearchFilter instantiates a new PublicAdsSearchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAdsSearchFilter(searchTerms []string, entityType string, adNetwork string, searchTermType string, filterType string, operator string) *PublicAdsSearchFilter {
	this := PublicAdsSearchFilter{}
	this.SearchTerms = searchTerms
	this.EntityType = entityType
	this.AdNetwork = adNetwork
	this.SearchTermType = searchTermType
	this.FilterType = filterType
	this.Operator = operator
	return &this
}

// NewPublicAdsSearchFilterWithDefaults instantiates a new PublicAdsSearchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAdsSearchFilterWithDefaults() *PublicAdsSearchFilter {
	this := PublicAdsSearchFilter{}
	var filterType string = "ADS_SEARCH"
	this.FilterType = filterType
	return &this
}

// GetSearchTerms returns the SearchTerms field value
func (o *PublicAdsSearchFilter) GetSearchTerms() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SearchTerms
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value
// and a boolean to check if the value has been set.
func (o *PublicAdsSearchFilter) GetSearchTermsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchTerms, true
}

// SetSearchTerms sets field value
func (o *PublicAdsSearchFilter) SetSearchTerms(v []string) {
	o.SearchTerms = v
}

// GetEntityType returns the EntityType field value
func (o *PublicAdsSearchFilter) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *PublicAdsSearchFilter) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *PublicAdsSearchFilter) SetEntityType(v string) {
	o.EntityType = v
}

// GetAdNetwork returns the AdNetwork field value
func (o *PublicAdsSearchFilter) GetAdNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdNetwork
}

// GetAdNetworkOk returns a tuple with the AdNetwork field value
// and a boolean to check if the value has been set.
func (o *PublicAdsSearchFilter) GetAdNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdNetwork, true
}

// SetAdNetwork sets field value
func (o *PublicAdsSearchFilter) SetAdNetwork(v string) {
	o.AdNetwork = v
}

// GetSearchTermType returns the SearchTermType field value
func (o *PublicAdsSearchFilter) GetSearchTermType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchTermType
}

// GetSearchTermTypeOk returns a tuple with the SearchTermType field value
// and a boolean to check if the value has been set.
func (o *PublicAdsSearchFilter) GetSearchTermTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchTermType, true
}

// SetSearchTermType sets field value
func (o *PublicAdsSearchFilter) SetSearchTermType(v string) {
	o.SearchTermType = v
}

// GetFilterType returns the FilterType field value
func (o *PublicAdsSearchFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicAdsSearchFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicAdsSearchFilter) SetFilterType(v string) {
	o.FilterType = v
}

// GetOperator returns the Operator field value
func (o *PublicAdsSearchFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PublicAdsSearchFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PublicAdsSearchFilter) SetOperator(v string) {
	o.Operator = v
}

func (o PublicAdsSearchFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicAdsSearchFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["searchTerms"] = o.SearchTerms
	toSerialize["entityType"] = o.EntityType
	toSerialize["adNetwork"] = o.AdNetwork
	toSerialize["searchTermType"] = o.SearchTermType
	toSerialize["filterType"] = o.FilterType
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *PublicAdsSearchFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"searchTerms",
		"entityType",
		"adNetwork",
		"searchTermType",
		"filterType",
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

	varPublicAdsSearchFilter := _PublicAdsSearchFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicAdsSearchFilter)

	if err != nil {
		return err
	}

	*o = PublicAdsSearchFilter(varPublicAdsSearchFilter)

	return err
}

type NullablePublicAdsSearchFilter struct {
	value *PublicAdsSearchFilter
	isSet bool
}

func (v NullablePublicAdsSearchFilter) Get() *PublicAdsSearchFilter {
	return v.value
}

func (v *NullablePublicAdsSearchFilter) Set(val *PublicAdsSearchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAdsSearchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAdsSearchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAdsSearchFilter(val *PublicAdsSearchFilter) *NullablePublicAdsSearchFilter {
	return &NullablePublicAdsSearchFilter{value: val, isSet: true}
}

func (v NullablePublicAdsSearchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAdsSearchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


