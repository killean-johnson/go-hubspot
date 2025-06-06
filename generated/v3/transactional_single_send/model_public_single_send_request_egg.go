/*
Transactional Single Send

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactional_single_send

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PublicSingleSendRequestEgg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicSingleSendRequestEgg{}

// PublicSingleSendRequestEgg A request to send a single transactional email asynchronously.
type PublicSingleSendRequestEgg struct {
	// The customProperties field is a map of property values. Each property value contains a name and value property. Each property will be visible in the template under {{ custom.NAME }}. Note: Custom properties do not currently support arrays. To provide a listing in an email, one workaround is to build an HTML list (either with tables or ul) and specify it as a custom property.
	CustomProperties map[string]map[string]interface{} `json:"customProperties,omitempty"`
	// The content ID for the transactional email, which can be found in email tool UI.
	EmailId int32 `json:"emailId"`
	Message PublicSingleSendEmail `json:"message"`
	// The contactProperties field is a map of contact property values. Each contact property value contains a name and value property. Each property will get set on the contact record and will be visible in the template under {{ contact.NAME }}. Use these properties when you want to set a contact property while you’re sending the email. For example, when sending a reciept you may want to set a last_paid_date property, as the sending of the receipt will have information about the last payment.
	ContactProperties *map[string]string `json:"contactProperties,omitempty"`
}

type _PublicSingleSendRequestEgg PublicSingleSendRequestEgg

// NewPublicSingleSendRequestEgg instantiates a new PublicSingleSendRequestEgg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSingleSendRequestEgg(emailId int32, message PublicSingleSendEmail) *PublicSingleSendRequestEgg {
	this := PublicSingleSendRequestEgg{}
	this.EmailId = emailId
	this.Message = message
	return &this
}

// NewPublicSingleSendRequestEggWithDefaults instantiates a new PublicSingleSendRequestEgg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSingleSendRequestEggWithDefaults() *PublicSingleSendRequestEgg {
	this := PublicSingleSendRequestEgg{}
	return &this
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *PublicSingleSendRequestEgg) GetCustomProperties() map[string]map[string]interface{} {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendRequestEgg) GetCustomPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return map[string]map[string]interface{}{}, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *PublicSingleSendRequestEgg) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the CustomProperties field.
func (o *PublicSingleSendRequestEgg) SetCustomProperties(v map[string]map[string]interface{}) {
	o.CustomProperties = v
}

// GetEmailId returns the EmailId field value
func (o *PublicSingleSendRequestEgg) GetEmailId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value
// and a boolean to check if the value has been set.
func (o *PublicSingleSendRequestEgg) GetEmailIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailId, true
}

// SetEmailId sets field value
func (o *PublicSingleSendRequestEgg) SetEmailId(v int32) {
	o.EmailId = v
}

// GetMessage returns the Message field value
func (o *PublicSingleSendRequestEgg) GetMessage() PublicSingleSendEmail {
	if o == nil {
		var ret PublicSingleSendEmail
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PublicSingleSendRequestEgg) GetMessageOk() (*PublicSingleSendEmail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PublicSingleSendRequestEgg) SetMessage(v PublicSingleSendEmail) {
	o.Message = v
}

// GetContactProperties returns the ContactProperties field value if set, zero value otherwise.
func (o *PublicSingleSendRequestEgg) GetContactProperties() map[string]string {
	if o == nil || IsNil(o.ContactProperties) {
		var ret map[string]string
		return ret
	}
	return *o.ContactProperties
}

// GetContactPropertiesOk returns a tuple with the ContactProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendRequestEgg) GetContactPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ContactProperties) {
		return nil, false
	}
	return o.ContactProperties, true
}

// HasContactProperties returns a boolean if a field has been set.
func (o *PublicSingleSendRequestEgg) HasContactProperties() bool {
	if o != nil && !IsNil(o.ContactProperties) {
		return true
	}

	return false
}

// SetContactProperties gets a reference to the given map[string]string and assigns it to the ContactProperties field.
func (o *PublicSingleSendRequestEgg) SetContactProperties(v map[string]string) {
	o.ContactProperties = &v
}

func (o PublicSingleSendRequestEgg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSingleSendRequestEgg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	toSerialize["emailId"] = o.EmailId
	toSerialize["message"] = o.Message
	if !IsNil(o.ContactProperties) {
		toSerialize["contactProperties"] = o.ContactProperties
	}
	return toSerialize, nil
}

func (o *PublicSingleSendRequestEgg) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emailId",
		"message",
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

	varPublicSingleSendRequestEgg := _PublicSingleSendRequestEgg{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicSingleSendRequestEgg)

	if err != nil {
		return err
	}

	*o = PublicSingleSendRequestEgg(varPublicSingleSendRequestEgg)

	return err
}

type NullablePublicSingleSendRequestEgg struct {
	value *PublicSingleSendRequestEgg
	isSet bool
}

func (v NullablePublicSingleSendRequestEgg) Get() *PublicSingleSendRequestEgg {
	return v.value
}

func (v *NullablePublicSingleSendRequestEgg) Set(val *PublicSingleSendRequestEgg) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSingleSendRequestEgg) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSingleSendRequestEgg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSingleSendRequestEgg(val *PublicSingleSendRequestEgg) *NullablePublicSingleSendRequestEgg {
	return &NullablePublicSingleSendRequestEgg{value: val, isSet: true}
}

func (v NullablePublicSingleSendRequestEgg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSingleSendRequestEgg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


