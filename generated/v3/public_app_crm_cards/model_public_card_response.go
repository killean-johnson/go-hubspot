/*
Public App Crm Cards

Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public_app_crm_cards

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the PublicCardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicCardResponse{}

// PublicCardResponse struct for PublicCardResponse
type PublicCardResponse struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Fetch PublicCardFetchBody `json:"fetch"`
	Display CardDisplayBody `json:"display"`
	Id string `json:"id"`
	Title string `json:"title"`
	Actions CardActions `json:"actions"`
	AuditHistory []CardAuditResponse `json:"auditHistory"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _PublicCardResponse PublicCardResponse

// NewPublicCardResponse instantiates a new PublicCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicCardResponse(fetch PublicCardFetchBody, display CardDisplayBody, id string, title string, actions CardActions, auditHistory []CardAuditResponse) *PublicCardResponse {
	this := PublicCardResponse{}
	this.Fetch = fetch
	this.Display = display
	this.Id = id
	this.Title = title
	this.Actions = actions
	this.AuditHistory = auditHistory
	return &this
}

// NewPublicCardResponseWithDefaults instantiates a new PublicCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicCardResponseWithDefaults() *PublicCardResponse {
	this := PublicCardResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PublicCardResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCardResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PublicCardResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PublicCardResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFetch returns the Fetch field value
func (o *PublicCardResponse) GetFetch() PublicCardFetchBody {
	if o == nil {
		var ret PublicCardFetchBody
		return ret
	}

	return o.Fetch
}

// GetFetchOk returns a tuple with the Fetch field value
// and a boolean to check if the value has been set.
func (o *PublicCardResponse) GetFetchOk() (*PublicCardFetchBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fetch, true
}

// SetFetch sets field value
func (o *PublicCardResponse) SetFetch(v PublicCardFetchBody) {
	o.Fetch = v
}

// GetDisplay returns the Display field value
func (o *PublicCardResponse) GetDisplay() CardDisplayBody {
	if o == nil {
		var ret CardDisplayBody
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *PublicCardResponse) GetDisplayOk() (*CardDisplayBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *PublicCardResponse) SetDisplay(v CardDisplayBody) {
	o.Display = v
}

// GetId returns the Id field value
func (o *PublicCardResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicCardResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicCardResponse) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *PublicCardResponse) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PublicCardResponse) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PublicCardResponse) SetTitle(v string) {
	o.Title = v
}

// GetActions returns the Actions field value
func (o *PublicCardResponse) GetActions() CardActions {
	if o == nil {
		var ret CardActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *PublicCardResponse) GetActionsOk() (*CardActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *PublicCardResponse) SetActions(v CardActions) {
	o.Actions = v
}

// GetAuditHistory returns the AuditHistory field value
func (o *PublicCardResponse) GetAuditHistory() []CardAuditResponse {
	if o == nil {
		var ret []CardAuditResponse
		return ret
	}

	return o.AuditHistory
}

// GetAuditHistoryOk returns a tuple with the AuditHistory field value
// and a boolean to check if the value has been set.
func (o *PublicCardResponse) GetAuditHistoryOk() ([]CardAuditResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditHistory, true
}

// SetAuditHistory sets field value
func (o *PublicCardResponse) SetAuditHistory(v []CardAuditResponse) {
	o.AuditHistory = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PublicCardResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCardResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PublicCardResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PublicCardResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PublicCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicCardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["fetch"] = o.Fetch
	toSerialize["display"] = o.Display
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["actions"] = o.Actions
	toSerialize["auditHistory"] = o.AuditHistory
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *PublicCardResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fetch",
		"display",
		"id",
		"title",
		"actions",
		"auditHistory",
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

	varPublicCardResponse := _PublicCardResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicCardResponse)

	if err != nil {
		return err
	}

	*o = PublicCardResponse(varPublicCardResponse)

	return err
}

type NullablePublicCardResponse struct {
	value *PublicCardResponse
	isSet bool
}

func (v NullablePublicCardResponse) Get() *PublicCardResponse {
	return v.value
}

func (v *NullablePublicCardResponse) Set(val *PublicCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicCardResponse(val *PublicCardResponse) *NullablePublicCardResponse {
	return &NullablePublicCardResponse{value: val, isSet: true}
}

func (v NullablePublicCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


