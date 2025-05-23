/*
Campaigns Public Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package campaigns_public_api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicSpendItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicSpendItem{}

// PublicSpendItem struct for PublicSpendItem
type PublicSpendItem struct {
	CreatedAt int32 `json:"createdAt"`
	Amount float32 `json:"amount"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Id string `json:"id"`
	Order int32 `json:"order"`
	UpdatedAt int32 `json:"updatedAt"`
}

type _PublicSpendItem PublicSpendItem

// NewPublicSpendItem instantiates a new PublicSpendItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSpendItem(createdAt int32, amount float32, name string, id string, order int32, updatedAt int32) *PublicSpendItem {
	this := PublicSpendItem{}
	this.CreatedAt = createdAt
	this.Amount = amount
	this.Name = name
	this.Id = id
	this.Order = order
	this.UpdatedAt = updatedAt
	return &this
}

// NewPublicSpendItemWithDefaults instantiates a new PublicSpendItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSpendItemWithDefaults() *PublicSpendItem {
	this := PublicSpendItem{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *PublicSpendItem) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PublicSpendItem) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PublicSpendItem) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetAmount returns the Amount field value
func (o *PublicSpendItem) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PublicSpendItem) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PublicSpendItem) SetAmount(v float32) {
	o.Amount = v
}

// GetName returns the Name field value
func (o *PublicSpendItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicSpendItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicSpendItem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PublicSpendItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSpendItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PublicSpendItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PublicSpendItem) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *PublicSpendItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicSpendItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicSpendItem) SetId(v string) {
	o.Id = v
}

// GetOrder returns the Order field value
func (o *PublicSpendItem) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *PublicSpendItem) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *PublicSpendItem) SetOrder(v int32) {
	o.Order = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PublicSpendItem) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PublicSpendItem) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PublicSpendItem) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o PublicSpendItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSpendItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["amount"] = o.Amount
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["order"] = o.Order
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *PublicSpendItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"amount",
		"name",
		"id",
		"order",
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

	varPublicSpendItem := _PublicSpendItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicSpendItem)

	if err != nil {
		return err
	}

	*o = PublicSpendItem(varPublicSpendItem)

	return err
}

type NullablePublicSpendItem struct {
	value *PublicSpendItem
	isSet bool
}

func (v NullablePublicSpendItem) Get() *PublicSpendItem {
	return v.value
}

func (v *NullablePublicSpendItem) Set(val *PublicSpendItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSpendItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSpendItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSpendItem(val *PublicSpendItem) *NullablePublicSpendItem {
	return &NullablePublicSpendItem{value: val, isSet: true}
}

func (v NullablePublicSpendItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSpendItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


