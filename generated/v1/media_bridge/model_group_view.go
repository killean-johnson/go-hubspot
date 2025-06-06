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

// checks if the GroupView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupView{}

// GroupView struct for GroupView
type GroupView struct {
	FulcrumPortalId int32 `json:"fulcrumPortalId"`
	FulcrumTimestamp int64 `json:"fulcrumTimestamp"`
	DisplayName string `json:"displayName"`
	Name string `json:"name"`
	DisplayOrder int32 `json:"displayOrder"`
	HubspotDefined bool `json:"hubspotDefined"`
}

type _GroupView GroupView

// NewGroupView instantiates a new GroupView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupView(fulcrumPortalId int32, fulcrumTimestamp int64, displayName string, name string, displayOrder int32, hubspotDefined bool) *GroupView {
	this := GroupView{}
	this.FulcrumPortalId = fulcrumPortalId
	this.FulcrumTimestamp = fulcrumTimestamp
	this.DisplayName = displayName
	this.Name = name
	this.DisplayOrder = displayOrder
	this.HubspotDefined = hubspotDefined
	return &this
}

// NewGroupViewWithDefaults instantiates a new GroupView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupViewWithDefaults() *GroupView {
	this := GroupView{}
	return &this
}

// GetFulcrumPortalId returns the FulcrumPortalId field value
func (o *GroupView) GetFulcrumPortalId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FulcrumPortalId
}

// GetFulcrumPortalIdOk returns a tuple with the FulcrumPortalId field value
// and a boolean to check if the value has been set.
func (o *GroupView) GetFulcrumPortalIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulcrumPortalId, true
}

// SetFulcrumPortalId sets field value
func (o *GroupView) SetFulcrumPortalId(v int32) {
	o.FulcrumPortalId = v
}

// GetFulcrumTimestamp returns the FulcrumTimestamp field value
func (o *GroupView) GetFulcrumTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FulcrumTimestamp
}

// GetFulcrumTimestampOk returns a tuple with the FulcrumTimestamp field value
// and a boolean to check if the value has been set.
func (o *GroupView) GetFulcrumTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulcrumTimestamp, true
}

// SetFulcrumTimestamp sets field value
func (o *GroupView) SetFulcrumTimestamp(v int64) {
	o.FulcrumTimestamp = v
}

// GetDisplayName returns the DisplayName field value
func (o *GroupView) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GroupView) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GroupView) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetName returns the Name field value
func (o *GroupView) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupView) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupView) SetName(v string) {
	o.Name = v
}

// GetDisplayOrder returns the DisplayOrder field value
func (o *GroupView) GetDisplayOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
func (o *GroupView) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrder, true
}

// SetDisplayOrder sets field value
func (o *GroupView) SetDisplayOrder(v int32) {
	o.DisplayOrder = v
}

// GetHubspotDefined returns the HubspotDefined field value
func (o *GroupView) GetHubspotDefined() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HubspotDefined
}

// GetHubspotDefinedOk returns a tuple with the HubspotDefined field value
// and a boolean to check if the value has been set.
func (o *GroupView) GetHubspotDefinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubspotDefined, true
}

// SetHubspotDefined sets field value
func (o *GroupView) SetHubspotDefined(v bool) {
	o.HubspotDefined = v
}

func (o GroupView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fulcrumPortalId"] = o.FulcrumPortalId
	toSerialize["fulcrumTimestamp"] = o.FulcrumTimestamp
	toSerialize["displayName"] = o.DisplayName
	toSerialize["name"] = o.Name
	toSerialize["displayOrder"] = o.DisplayOrder
	toSerialize["hubspotDefined"] = o.HubspotDefined
	return toSerialize, nil
}

func (o *GroupView) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fulcrumPortalId",
		"fulcrumTimestamp",
		"displayName",
		"name",
		"displayOrder",
		"hubspotDefined",
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

	varGroupView := _GroupView{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupView)

	if err != nil {
		return err
	}

	*o = GroupView(varGroupView)

	return err
}

type NullableGroupView struct {
	value *GroupView
	isSet bool
}

func (v NullableGroupView) Get() *GroupView {
	return v.value
}

func (v *NullableGroupView) Set(val *GroupView) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupView) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupView(val *GroupView) *NullableGroupView {
	return &NullableGroupView{value: val, isSet: true}
}

func (v NullableGroupView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


