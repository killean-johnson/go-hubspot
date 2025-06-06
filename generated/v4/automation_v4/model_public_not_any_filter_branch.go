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

// checks if the PublicNotAnyFilterBranch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNotAnyFilterBranch{}

// PublicNotAnyFilterBranch struct for PublicNotAnyFilterBranch
type PublicNotAnyFilterBranch struct {
	FilterBranchType string `json:"filterBranchType"`
	FilterBranches []ApiContactFlowPutRequestAllOfGoalFilterBranch `json:"filterBranches"`
	FilterBranchOperator string `json:"filterBranchOperator"`
	Filters []PublicPropertyAssociationFilterBranchFiltersInner `json:"filters"`
}

type _PublicNotAnyFilterBranch PublicNotAnyFilterBranch

// NewPublicNotAnyFilterBranch instantiates a new PublicNotAnyFilterBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNotAnyFilterBranch(filterBranchType string, filterBranches []ApiContactFlowPutRequestAllOfGoalFilterBranch, filterBranchOperator string, filters []PublicPropertyAssociationFilterBranchFiltersInner) *PublicNotAnyFilterBranch {
	this := PublicNotAnyFilterBranch{}
	this.FilterBranchType = filterBranchType
	this.FilterBranches = filterBranches
	this.FilterBranchOperator = filterBranchOperator
	this.Filters = filters
	return &this
}

// NewPublicNotAnyFilterBranchWithDefaults instantiates a new PublicNotAnyFilterBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNotAnyFilterBranchWithDefaults() *PublicNotAnyFilterBranch {
	this := PublicNotAnyFilterBranch{}
	var filterBranchType string = "NOT_ANY"
	this.FilterBranchType = filterBranchType
	return &this
}

// GetFilterBranchType returns the FilterBranchType field value
func (o *PublicNotAnyFilterBranch) GetFilterBranchType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterBranchType
}

// GetFilterBranchTypeOk returns a tuple with the FilterBranchType field value
// and a boolean to check if the value has been set.
func (o *PublicNotAnyFilterBranch) GetFilterBranchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterBranchType, true
}

// SetFilterBranchType sets field value
func (o *PublicNotAnyFilterBranch) SetFilterBranchType(v string) {
	o.FilterBranchType = v
}

// GetFilterBranches returns the FilterBranches field value
func (o *PublicNotAnyFilterBranch) GetFilterBranches() []ApiContactFlowPutRequestAllOfGoalFilterBranch {
	if o == nil {
		var ret []ApiContactFlowPutRequestAllOfGoalFilterBranch
		return ret
	}

	return o.FilterBranches
}

// GetFilterBranchesOk returns a tuple with the FilterBranches field value
// and a boolean to check if the value has been set.
func (o *PublicNotAnyFilterBranch) GetFilterBranchesOk() ([]ApiContactFlowPutRequestAllOfGoalFilterBranch, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterBranches, true
}

// SetFilterBranches sets field value
func (o *PublicNotAnyFilterBranch) SetFilterBranches(v []ApiContactFlowPutRequestAllOfGoalFilterBranch) {
	o.FilterBranches = v
}

// GetFilterBranchOperator returns the FilterBranchOperator field value
func (o *PublicNotAnyFilterBranch) GetFilterBranchOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterBranchOperator
}

// GetFilterBranchOperatorOk returns a tuple with the FilterBranchOperator field value
// and a boolean to check if the value has been set.
func (o *PublicNotAnyFilterBranch) GetFilterBranchOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterBranchOperator, true
}

// SetFilterBranchOperator sets field value
func (o *PublicNotAnyFilterBranch) SetFilterBranchOperator(v string) {
	o.FilterBranchOperator = v
}

// GetFilters returns the Filters field value
func (o *PublicNotAnyFilterBranch) GetFilters() []PublicPropertyAssociationFilterBranchFiltersInner {
	if o == nil {
		var ret []PublicPropertyAssociationFilterBranchFiltersInner
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *PublicNotAnyFilterBranch) GetFiltersOk() ([]PublicPropertyAssociationFilterBranchFiltersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *PublicNotAnyFilterBranch) SetFilters(v []PublicPropertyAssociationFilterBranchFiltersInner) {
	o.Filters = v
}

func (o PublicNotAnyFilterBranch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicNotAnyFilterBranch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterBranchType"] = o.FilterBranchType
	toSerialize["filterBranches"] = o.FilterBranches
	toSerialize["filterBranchOperator"] = o.FilterBranchOperator
	toSerialize["filters"] = o.Filters
	return toSerialize, nil
}

func (o *PublicNotAnyFilterBranch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterBranchType",
		"filterBranches",
		"filterBranchOperator",
		"filters",
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

	varPublicNotAnyFilterBranch := _PublicNotAnyFilterBranch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicNotAnyFilterBranch)

	if err != nil {
		return err
	}

	*o = PublicNotAnyFilterBranch(varPublicNotAnyFilterBranch)

	return err
}

type NullablePublicNotAnyFilterBranch struct {
	value *PublicNotAnyFilterBranch
	isSet bool
}

func (v NullablePublicNotAnyFilterBranch) Get() *PublicNotAnyFilterBranch {
	return v.value
}

func (v *NullablePublicNotAnyFilterBranch) Set(val *PublicNotAnyFilterBranch) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNotAnyFilterBranch) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNotAnyFilterBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNotAnyFilterBranch(val *PublicNotAnyFilterBranch) *NullablePublicNotAnyFilterBranch {
	return &NullablePublicNotAnyFilterBranch{value: val, isSet: true}
}

func (v NullablePublicNotAnyFilterBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNotAnyFilterBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


