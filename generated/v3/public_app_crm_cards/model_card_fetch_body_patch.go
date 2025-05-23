/*
Public App Crm Cards

Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public_app_crm_cards

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CardFetchBodyPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardFetchBodyPatch{}

// CardFetchBodyPatch Variant of CardFetchBody with fields as optional for patches
type CardFetchBodyPatch struct {
	ServerlessFunction *string `json:"serverlessFunction,omitempty"`
	CardType *string `json:"cardType,omitempty"`
	// An array of CRM object types where this card should be displayed. HubSpot will call your target URL whenever a user visits a record page of the types defined here.
	ObjectTypes []CardObjectTypeBody `json:"objectTypes"`
	// URL to a service endpoint that will respond with details for this card. HubSpot will call this endpoint each time a user visits a CRM record page where this card should be displayed.
	TargetUrl *string `json:"targetUrl,omitempty"`
}

type _CardFetchBodyPatch CardFetchBodyPatch

// NewCardFetchBodyPatch instantiates a new CardFetchBodyPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardFetchBodyPatch(objectTypes []CardObjectTypeBody) *CardFetchBodyPatch {
	this := CardFetchBodyPatch{}
	this.ObjectTypes = objectTypes
	return &this
}

// NewCardFetchBodyPatchWithDefaults instantiates a new CardFetchBodyPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardFetchBodyPatchWithDefaults() *CardFetchBodyPatch {
	this := CardFetchBodyPatch{}
	return &this
}

// GetServerlessFunction returns the ServerlessFunction field value if set, zero value otherwise.
func (o *CardFetchBodyPatch) GetServerlessFunction() string {
	if o == nil || IsNil(o.ServerlessFunction) {
		var ret string
		return ret
	}
	return *o.ServerlessFunction
}

// GetServerlessFunctionOk returns a tuple with the ServerlessFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardFetchBodyPatch) GetServerlessFunctionOk() (*string, bool) {
	if o == nil || IsNil(o.ServerlessFunction) {
		return nil, false
	}
	return o.ServerlessFunction, true
}

// HasServerlessFunction returns a boolean if a field has been set.
func (o *CardFetchBodyPatch) HasServerlessFunction() bool {
	if o != nil && !IsNil(o.ServerlessFunction) {
		return true
	}

	return false
}

// SetServerlessFunction gets a reference to the given string and assigns it to the ServerlessFunction field.
func (o *CardFetchBodyPatch) SetServerlessFunction(v string) {
	o.ServerlessFunction = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *CardFetchBodyPatch) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardFetchBodyPatch) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *CardFetchBodyPatch) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *CardFetchBodyPatch) SetCardType(v string) {
	o.CardType = &v
}

// GetObjectTypes returns the ObjectTypes field value
func (o *CardFetchBodyPatch) GetObjectTypes() []CardObjectTypeBody {
	if o == nil {
		var ret []CardObjectTypeBody
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *CardFetchBodyPatch) GetObjectTypesOk() ([]CardObjectTypeBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *CardFetchBodyPatch) SetObjectTypes(v []CardObjectTypeBody) {
	o.ObjectTypes = v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *CardFetchBodyPatch) GetTargetUrl() string {
	if o == nil || IsNil(o.TargetUrl) {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardFetchBodyPatch) GetTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUrl) {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *CardFetchBodyPatch) HasTargetUrl() bool {
	if o != nil && !IsNil(o.TargetUrl) {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *CardFetchBodyPatch) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

func (o CardFetchBodyPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardFetchBodyPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerlessFunction) {
		toSerialize["serverlessFunction"] = o.ServerlessFunction
	}
	if !IsNil(o.CardType) {
		toSerialize["cardType"] = o.CardType
	}
	toSerialize["objectTypes"] = o.ObjectTypes
	if !IsNil(o.TargetUrl) {
		toSerialize["targetUrl"] = o.TargetUrl
	}
	return toSerialize, nil
}

func (o *CardFetchBodyPatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectTypes",
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

	varCardFetchBodyPatch := _CardFetchBodyPatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCardFetchBodyPatch)

	if err != nil {
		return err
	}

	*o = CardFetchBodyPatch(varCardFetchBodyPatch)

	return err
}

type NullableCardFetchBodyPatch struct {
	value *CardFetchBodyPatch
	isSet bool
}

func (v NullableCardFetchBodyPatch) Get() *CardFetchBodyPatch {
	return v.value
}

func (v *NullableCardFetchBodyPatch) Set(val *CardFetchBodyPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableCardFetchBodyPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableCardFetchBodyPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardFetchBodyPatch(val *CardFetchBodyPatch) *NullableCardFetchBodyPatch {
	return &NullableCardFetchBodyPatch{value: val, isSet: true}
}

func (v NullableCardFetchBodyPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardFetchBodyPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


