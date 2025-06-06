/*
Lists

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the PublicObjectListSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicObjectListSearchResult{}

// PublicObjectListSearchResult struct for PublicObjectListSearchResult
type PublicObjectListSearchResult struct {
	// The processing type of the list.
	ProcessingType string `json:"processingType"`
	// The object type of the list.
	ObjectTypeId string `json:"objectTypeId"`
	// The ID of the user that last updated the list.
	UpdatedById *string `json:"updatedById,omitempty"`
	// The time when the filters for this list were last updated.
	FiltersUpdatedAt *time.Time `json:"filtersUpdatedAt,omitempty"`
	// The **ILS ID** of the list.
	ListId string `json:"listId"`
	// The time when the list was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The processing status of the list.
	ProcessingStatus string `json:"processingStatus"`
	// The time when the list was deleted.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The version of the list.
	ListVersion int32 `json:"listVersion"`
	// The name of the list.
	Name string `json:"name"`
	// The name and value of any additional properties that exist for this list and that were included in the search request.
	AdditionalPropertiesField map[string]string `json:"additionalProperties"`
	// The ID of the user that created the list.
	CreatedById *string `json:"createdById,omitempty"`
	// The time the list was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _PublicObjectListSearchResult PublicObjectListSearchResult

// NewPublicObjectListSearchResult instantiates a new PublicObjectListSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicObjectListSearchResult(processingType string, objectTypeId string, listId string, processingStatus string, listVersion int32, name string, additionalPropertiesField map[string]string) *PublicObjectListSearchResult {
	this := PublicObjectListSearchResult{}
	this.ProcessingType = processingType
	this.ObjectTypeId = objectTypeId
	this.ListId = listId
	this.ProcessingStatus = processingStatus
	this.ListVersion = listVersion
	this.Name = name
	this.AdditionalPropertiesField = additionalPropertiesField
	return &this
}

// NewPublicObjectListSearchResultWithDefaults instantiates a new PublicObjectListSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicObjectListSearchResultWithDefaults() *PublicObjectListSearchResult {
	this := PublicObjectListSearchResult{}
	return &this
}

// GetProcessingType returns the ProcessingType field value
func (o *PublicObjectListSearchResult) GetProcessingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessingType
}

// GetProcessingTypeOk returns a tuple with the ProcessingType field value
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetProcessingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingType, true
}

// SetProcessingType sets field value
func (o *PublicObjectListSearchResult) SetProcessingType(v string) {
	o.ProcessingType = v
}

// GetObjectTypeId returns the ObjectTypeId field value
func (o *PublicObjectListSearchResult) GetObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectTypeId
}

// GetObjectTypeIdOk returns a tuple with the ObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectTypeId, true
}

// SetObjectTypeId sets field value
func (o *PublicObjectListSearchResult) SetObjectTypeId(v string) {
	o.ObjectTypeId = v
}

// GetUpdatedById returns the UpdatedById field value if set, zero value otherwise.
func (o *PublicObjectListSearchResult) GetUpdatedById() string {
	if o == nil || IsNil(o.UpdatedById) {
		var ret string
		return ret
	}
	return *o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetUpdatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedById) {
		return nil, false
	}
	return o.UpdatedById, true
}

// HasUpdatedById returns a boolean if a field has been set.
func (o *PublicObjectListSearchResult) HasUpdatedById() bool {
	if o != nil && !IsNil(o.UpdatedById) {
		return true
	}

	return false
}

// SetUpdatedById gets a reference to the given string and assigns it to the UpdatedById field.
func (o *PublicObjectListSearchResult) SetUpdatedById(v string) {
	o.UpdatedById = &v
}

// GetFiltersUpdatedAt returns the FiltersUpdatedAt field value if set, zero value otherwise.
func (o *PublicObjectListSearchResult) GetFiltersUpdatedAt() time.Time {
	if o == nil || IsNil(o.FiltersUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.FiltersUpdatedAt
}

// GetFiltersUpdatedAtOk returns a tuple with the FiltersUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetFiltersUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FiltersUpdatedAt) {
		return nil, false
	}
	return o.FiltersUpdatedAt, true
}

// HasFiltersUpdatedAt returns a boolean if a field has been set.
func (o *PublicObjectListSearchResult) HasFiltersUpdatedAt() bool {
	if o != nil && !IsNil(o.FiltersUpdatedAt) {
		return true
	}

	return false
}

// SetFiltersUpdatedAt gets a reference to the given time.Time and assigns it to the FiltersUpdatedAt field.
func (o *PublicObjectListSearchResult) SetFiltersUpdatedAt(v time.Time) {
	o.FiltersUpdatedAt = &v
}

// GetListId returns the ListId field value
func (o *PublicObjectListSearchResult) GetListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *PublicObjectListSearchResult) SetListId(v string) {
	o.ListId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PublicObjectListSearchResult) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PublicObjectListSearchResult) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PublicObjectListSearchResult) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetProcessingStatus returns the ProcessingStatus field value
func (o *PublicObjectListSearchResult) GetProcessingStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessingStatus
}

// GetProcessingStatusOk returns a tuple with the ProcessingStatus field value
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetProcessingStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingStatus, true
}

// SetProcessingStatus sets field value
func (o *PublicObjectListSearchResult) SetProcessingStatus(v string) {
	o.ProcessingStatus = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *PublicObjectListSearchResult) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *PublicObjectListSearchResult) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *PublicObjectListSearchResult) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetListVersion returns the ListVersion field value
func (o *PublicObjectListSearchResult) GetListVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ListVersion
}

// GetListVersionOk returns a tuple with the ListVersion field value
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetListVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListVersion, true
}

// SetListVersion sets field value
func (o *PublicObjectListSearchResult) SetListVersion(v int32) {
	o.ListVersion = v
}

// GetName returns the Name field value
func (o *PublicObjectListSearchResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicObjectListSearchResult) SetName(v string) {
	o.Name = v
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value
func (o *PublicObjectListSearchResult) GetAdditionalPropertiesField() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetAdditionalPropertiesFieldOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalPropertiesField, true
}

// SetAdditionalPropertiesField sets field value
func (o *PublicObjectListSearchResult) SetAdditionalPropertiesField(v map[string]string) {
	o.AdditionalPropertiesField = v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *PublicObjectListSearchResult) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *PublicObjectListSearchResult) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *PublicObjectListSearchResult) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PublicObjectListSearchResult) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectListSearchResult) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PublicObjectListSearchResult) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PublicObjectListSearchResult) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PublicObjectListSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicObjectListSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["processingType"] = o.ProcessingType
	toSerialize["objectTypeId"] = o.ObjectTypeId
	if !IsNil(o.UpdatedById) {
		toSerialize["updatedById"] = o.UpdatedById
	}
	if !IsNil(o.FiltersUpdatedAt) {
		toSerialize["filtersUpdatedAt"] = o.FiltersUpdatedAt
	}
	toSerialize["listId"] = o.ListId
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["processingStatus"] = o.ProcessingStatus
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	toSerialize["listVersion"] = o.ListVersion
	toSerialize["name"] = o.Name
	toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *PublicObjectListSearchResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"processingType",
		"objectTypeId",
		"listId",
		"processingStatus",
		"listVersion",
		"name",
		"additionalProperties",
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

	varPublicObjectListSearchResult := _PublicObjectListSearchResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicObjectListSearchResult)

	if err != nil {
		return err
	}

	*o = PublicObjectListSearchResult(varPublicObjectListSearchResult)

	return err
}

type NullablePublicObjectListSearchResult struct {
	value *PublicObjectListSearchResult
	isSet bool
}

func (v NullablePublicObjectListSearchResult) Get() *PublicObjectListSearchResult {
	return v.value
}

func (v *NullablePublicObjectListSearchResult) Set(val *PublicObjectListSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicObjectListSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicObjectListSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicObjectListSearchResult(val *PublicObjectListSearchResult) *NullablePublicObjectListSearchResult {
	return &NullablePublicObjectListSearchResult{value: val, isSet: true}
}

func (v NullablePublicObjectListSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicObjectListSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


