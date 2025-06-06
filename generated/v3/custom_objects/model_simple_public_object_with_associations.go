/*
Custom Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package custom_objects

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the SimplePublicObjectWithAssociations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplePublicObjectWithAssociations{}

// SimplePublicObjectWithAssociations struct for SimplePublicObjectWithAssociations
type SimplePublicObjectWithAssociations struct {
	Associations *map[string]CollectionResponseAssociatedId `json:"associations,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	Archived *bool `json:"archived,omitempty"`
	ArchivedAt *time.Time `json:"archivedAt,omitempty"`
	PropertiesWithHistory *map[string][]ValueWithTimestamp `json:"propertiesWithHistory,omitempty"`
	Id string `json:"id"`
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`
	Properties map[string]string `json:"properties"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type _SimplePublicObjectWithAssociations SimplePublicObjectWithAssociations

// NewSimplePublicObjectWithAssociations instantiates a new SimplePublicObjectWithAssociations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectWithAssociations(createdAt time.Time, id string, properties map[string]string, updatedAt time.Time) *SimplePublicObjectWithAssociations {
	this := SimplePublicObjectWithAssociations{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Properties = properties
	this.UpdatedAt = updatedAt
	return &this
}

// NewSimplePublicObjectWithAssociationsWithDefaults instantiates a new SimplePublicObjectWithAssociations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectWithAssociationsWithDefaults() *SimplePublicObjectWithAssociations {
	this := SimplePublicObjectWithAssociations{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *SimplePublicObjectWithAssociations) GetAssociations() map[string]CollectionResponseAssociatedId {
	if o == nil || IsNil(o.Associations) {
		var ret map[string]CollectionResponseAssociatedId
		return ret
	}
	return *o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetAssociationsOk() (*map[string]CollectionResponseAssociatedId, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *SimplePublicObjectWithAssociations) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given map[string]CollectionResponseAssociatedId and assigns it to the Associations field.
func (o *SimplePublicObjectWithAssociations) SetAssociations(v map[string]CollectionResponseAssociatedId) {
	o.Associations = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SimplePublicObjectWithAssociations) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SimplePublicObjectWithAssociations) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *SimplePublicObjectWithAssociations) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *SimplePublicObjectWithAssociations) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *SimplePublicObjectWithAssociations) SetArchived(v bool) {
	o.Archived = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *SimplePublicObjectWithAssociations) GetArchivedAt() time.Time {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *SimplePublicObjectWithAssociations) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *SimplePublicObjectWithAssociations) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetPropertiesWithHistory returns the PropertiesWithHistory field value if set, zero value otherwise.
func (o *SimplePublicObjectWithAssociations) GetPropertiesWithHistory() map[string][]ValueWithTimestamp {
	if o == nil || IsNil(o.PropertiesWithHistory) {
		var ret map[string][]ValueWithTimestamp
		return ret
	}
	return *o.PropertiesWithHistory
}

// GetPropertiesWithHistoryOk returns a tuple with the PropertiesWithHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetPropertiesWithHistoryOk() (*map[string][]ValueWithTimestamp, bool) {
	if o == nil || IsNil(o.PropertiesWithHistory) {
		return nil, false
	}
	return o.PropertiesWithHistory, true
}

// HasPropertiesWithHistory returns a boolean if a field has been set.
func (o *SimplePublicObjectWithAssociations) HasPropertiesWithHistory() bool {
	if o != nil && !IsNil(o.PropertiesWithHistory) {
		return true
	}

	return false
}

// SetPropertiesWithHistory gets a reference to the given map[string][]ValueWithTimestamp and assigns it to the PropertiesWithHistory field.
func (o *SimplePublicObjectWithAssociations) SetPropertiesWithHistory(v map[string][]ValueWithTimestamp) {
	o.PropertiesWithHistory = &v
}

// GetId returns the Id field value
func (o *SimplePublicObjectWithAssociations) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplePublicObjectWithAssociations) SetId(v string) {
	o.Id = v
}

// GetObjectWriteTraceId returns the ObjectWriteTraceId field value if set, zero value otherwise.
func (o *SimplePublicObjectWithAssociations) GetObjectWriteTraceId() string {
	if o == nil || IsNil(o.ObjectWriteTraceId) {
		var ret string
		return ret
	}
	return *o.ObjectWriteTraceId
}

// GetObjectWriteTraceIdOk returns a tuple with the ObjectWriteTraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetObjectWriteTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectWriteTraceId) {
		return nil, false
	}
	return o.ObjectWriteTraceId, true
}

// HasObjectWriteTraceId returns a boolean if a field has been set.
func (o *SimplePublicObjectWithAssociations) HasObjectWriteTraceId() bool {
	if o != nil && !IsNil(o.ObjectWriteTraceId) {
		return true
	}

	return false
}

// SetObjectWriteTraceId gets a reference to the given string and assigns it to the ObjectWriteTraceId field.
func (o *SimplePublicObjectWithAssociations) SetObjectWriteTraceId(v string) {
	o.ObjectWriteTraceId = &v
}

// GetProperties returns the Properties field value
func (o *SimplePublicObjectWithAssociations) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SimplePublicObjectWithAssociations) SetProperties(v map[string]string) {
	o.Properties = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SimplePublicObjectWithAssociations) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SimplePublicObjectWithAssociations) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o SimplePublicObjectWithAssociations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplePublicObjectWithAssociations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if !IsNil(o.PropertiesWithHistory) {
		toSerialize["propertiesWithHistory"] = o.PropertiesWithHistory
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.ObjectWriteTraceId) {
		toSerialize["objectWriteTraceId"] = o.ObjectWriteTraceId
	}
	toSerialize["properties"] = o.Properties
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *SimplePublicObjectWithAssociations) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"id",
		"properties",
		"updatedAt",
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

	varSimplePublicObjectWithAssociations := _SimplePublicObjectWithAssociations{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimplePublicObjectWithAssociations)

	if err != nil {
		return err
	}

	*o = SimplePublicObjectWithAssociations(varSimplePublicObjectWithAssociations)

	return err
}

type NullableSimplePublicObjectWithAssociations struct {
	value *SimplePublicObjectWithAssociations
	isSet bool
}

func (v NullableSimplePublicObjectWithAssociations) Get() *SimplePublicObjectWithAssociations {
	return v.value
}

func (v *NullableSimplePublicObjectWithAssociations) Set(val *SimplePublicObjectWithAssociations) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectWithAssociations) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectWithAssociations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectWithAssociations(val *SimplePublicObjectWithAssociations) *NullableSimplePublicObjectWithAssociations {
	return &NullableSimplePublicObjectWithAssociations{value: val, isSet: true}
}

func (v NullableSimplePublicObjectWithAssociations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectWithAssociations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


