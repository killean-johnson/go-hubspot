/*
Lists

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCreateRequest{}

// ListCreateRequest struct for ListCreateRequest
type ListCreateRequest struct {
	// The object type ID of the type of objects that the list will store.
	ObjectTypeId string `json:"objectTypeId"`
	// The processing type of the list. One of: `SNAPSHOT`, `MANUAL`, or `DYNAMIC`.
	ProcessingType string `json:"processingType"`
	// The list of custom properties to tie to the list. Custom property name is the key, the value is the value.
	CustomProperties *map[string]string `json:"customProperties,omitempty"`
	// The ID of the folder that the list should be created in. If left blank, then the list will be created in the root of the list folder structure.
	ListFolderId *int32 `json:"listFolderId,omitempty"`
	// The name of the list, which must be globally unique across all public lists in the portal.
	Name string `json:"name"`
	FilterBranch *PublicPropertyAssociationFilterBranchFilterBranchesInner `json:"filterBranch,omitempty"`
}

type _ListCreateRequest ListCreateRequest

// NewListCreateRequest instantiates a new ListCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCreateRequest(objectTypeId string, processingType string, name string) *ListCreateRequest {
	this := ListCreateRequest{}
	this.ObjectTypeId = objectTypeId
	this.ProcessingType = processingType
	this.Name = name
	return &this
}

// NewListCreateRequestWithDefaults instantiates a new ListCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCreateRequestWithDefaults() *ListCreateRequest {
	this := ListCreateRequest{}
	return &this
}

// GetObjectTypeId returns the ObjectTypeId field value
func (o *ListCreateRequest) GetObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectTypeId
}

// GetObjectTypeIdOk returns a tuple with the ObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *ListCreateRequest) GetObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectTypeId, true
}

// SetObjectTypeId sets field value
func (o *ListCreateRequest) SetObjectTypeId(v string) {
	o.ObjectTypeId = v
}

// GetProcessingType returns the ProcessingType field value
func (o *ListCreateRequest) GetProcessingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessingType
}

// GetProcessingTypeOk returns a tuple with the ProcessingType field value
// and a boolean to check if the value has been set.
func (o *ListCreateRequest) GetProcessingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingType, true
}

// SetProcessingType sets field value
func (o *ListCreateRequest) SetProcessingType(v string) {
	o.ProcessingType = v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *ListCreateRequest) GetCustomProperties() map[string]string {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]string
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreateRequest) GetCustomPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *ListCreateRequest) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]string and assigns it to the CustomProperties field.
func (o *ListCreateRequest) SetCustomProperties(v map[string]string) {
	o.CustomProperties = &v
}

// GetListFolderId returns the ListFolderId field value if set, zero value otherwise.
func (o *ListCreateRequest) GetListFolderId() int32 {
	if o == nil || IsNil(o.ListFolderId) {
		var ret int32
		return ret
	}
	return *o.ListFolderId
}

// GetListFolderIdOk returns a tuple with the ListFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreateRequest) GetListFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ListFolderId) {
		return nil, false
	}
	return o.ListFolderId, true
}

// HasListFolderId returns a boolean if a field has been set.
func (o *ListCreateRequest) HasListFolderId() bool {
	if o != nil && !IsNil(o.ListFolderId) {
		return true
	}

	return false
}

// SetListFolderId gets a reference to the given int32 and assigns it to the ListFolderId field.
func (o *ListCreateRequest) SetListFolderId(v int32) {
	o.ListFolderId = &v
}

// GetName returns the Name field value
func (o *ListCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListCreateRequest) SetName(v string) {
	o.Name = v
}

// GetFilterBranch returns the FilterBranch field value if set, zero value otherwise.
func (o *ListCreateRequest) GetFilterBranch() PublicPropertyAssociationFilterBranchFilterBranchesInner {
	if o == nil || IsNil(o.FilterBranch) {
		var ret PublicPropertyAssociationFilterBranchFilterBranchesInner
		return ret
	}
	return *o.FilterBranch
}

// GetFilterBranchOk returns a tuple with the FilterBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreateRequest) GetFilterBranchOk() (*PublicPropertyAssociationFilterBranchFilterBranchesInner, bool) {
	if o == nil || IsNil(o.FilterBranch) {
		return nil, false
	}
	return o.FilterBranch, true
}

// HasFilterBranch returns a boolean if a field has been set.
func (o *ListCreateRequest) HasFilterBranch() bool {
	if o != nil && !IsNil(o.FilterBranch) {
		return true
	}

	return false
}

// SetFilterBranch gets a reference to the given PublicPropertyAssociationFilterBranchFilterBranchesInner and assigns it to the FilterBranch field.
func (o *ListCreateRequest) SetFilterBranch(v PublicPropertyAssociationFilterBranchFilterBranchesInner) {
	o.FilterBranch = &v
}

func (o ListCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objectTypeId"] = o.ObjectTypeId
	toSerialize["processingType"] = o.ProcessingType
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.ListFolderId) {
		toSerialize["listFolderId"] = o.ListFolderId
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.FilterBranch) {
		toSerialize["filterBranch"] = o.FilterBranch
	}
	return toSerialize, nil
}

func (o *ListCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectTypeId",
		"processingType",
		"name",
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

	varListCreateRequest := _ListCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListCreateRequest)

	if err != nil {
		return err
	}

	*o = ListCreateRequest(varListCreateRequest)

	return err
}

type NullableListCreateRequest struct {
	value *ListCreateRequest
	isSet bool
}

func (v NullableListCreateRequest) Get() *ListCreateRequest {
	return v.value
}

func (v *NullableListCreateRequest) Set(val *ListCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCreateRequest(val *ListCreateRequest) *NullableListCreateRequest {
	return &NullableListCreateRequest{value: val, isSet: true}
}

func (v NullableListCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


