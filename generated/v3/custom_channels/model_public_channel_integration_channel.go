/*
Conversations Custom Channels

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package custom_channels

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the PublicChannelIntegrationChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicChannelIntegrationChannel{}

// PublicChannelIntegrationChannel struct for PublicChannelIntegrationChannel
type PublicChannelIntegrationChannel struct {
	ChannelDescription *string `json:"channelDescription,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	// 
	Capabilities map[string]map[string]interface{} `json:"capabilities"`
	ChannelAccountConnectionRedirectUrl *string `json:"channelAccountConnectionRedirectUrl,omitempty"`
	ChannelLogoUrl *string `json:"channelLogoUrl,omitempty"`
	Name string `json:"name"`
	Id string `json:"id"`
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}

type _PublicChannelIntegrationChannel PublicChannelIntegrationChannel

// NewPublicChannelIntegrationChannel instantiates a new PublicChannelIntegrationChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicChannelIntegrationChannel(createdAt time.Time, capabilities map[string]map[string]interface{}, name string, id string) *PublicChannelIntegrationChannel {
	this := PublicChannelIntegrationChannel{}
	this.CreatedAt = createdAt
	this.Capabilities = capabilities
	this.Name = name
	this.Id = id
	return &this
}

// NewPublicChannelIntegrationChannelWithDefaults instantiates a new PublicChannelIntegrationChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicChannelIntegrationChannelWithDefaults() *PublicChannelIntegrationChannel {
	this := PublicChannelIntegrationChannel{}
	return &this
}

// GetChannelDescription returns the ChannelDescription field value if set, zero value otherwise.
func (o *PublicChannelIntegrationChannel) GetChannelDescription() string {
	if o == nil || IsNil(o.ChannelDescription) {
		var ret string
		return ret
	}
	return *o.ChannelDescription
}

// GetChannelDescriptionOk returns a tuple with the ChannelDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelIntegrationChannel) GetChannelDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelDescription) {
		return nil, false
	}
	return o.ChannelDescription, true
}

// HasChannelDescription returns a boolean if a field has been set.
func (o *PublicChannelIntegrationChannel) HasChannelDescription() bool {
	if o != nil && !IsNil(o.ChannelDescription) {
		return true
	}

	return false
}

// SetChannelDescription gets a reference to the given string and assigns it to the ChannelDescription field.
func (o *PublicChannelIntegrationChannel) SetChannelDescription(v string) {
	o.ChannelDescription = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PublicChannelIntegrationChannel) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PublicChannelIntegrationChannel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PublicChannelIntegrationChannel) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCapabilities returns the Capabilities field value
func (o *PublicChannelIntegrationChannel) GetCapabilities() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *PublicChannelIntegrationChannel) GetCapabilitiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil {
		return map[string]map[string]interface{}{}, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *PublicChannelIntegrationChannel) SetCapabilities(v map[string]map[string]interface{}) {
	o.Capabilities = v
}

// GetChannelAccountConnectionRedirectUrl returns the ChannelAccountConnectionRedirectUrl field value if set, zero value otherwise.
func (o *PublicChannelIntegrationChannel) GetChannelAccountConnectionRedirectUrl() string {
	if o == nil || IsNil(o.ChannelAccountConnectionRedirectUrl) {
		var ret string
		return ret
	}
	return *o.ChannelAccountConnectionRedirectUrl
}

// GetChannelAccountConnectionRedirectUrlOk returns a tuple with the ChannelAccountConnectionRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelIntegrationChannel) GetChannelAccountConnectionRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelAccountConnectionRedirectUrl) {
		return nil, false
	}
	return o.ChannelAccountConnectionRedirectUrl, true
}

// HasChannelAccountConnectionRedirectUrl returns a boolean if a field has been set.
func (o *PublicChannelIntegrationChannel) HasChannelAccountConnectionRedirectUrl() bool {
	if o != nil && !IsNil(o.ChannelAccountConnectionRedirectUrl) {
		return true
	}

	return false
}

// SetChannelAccountConnectionRedirectUrl gets a reference to the given string and assigns it to the ChannelAccountConnectionRedirectUrl field.
func (o *PublicChannelIntegrationChannel) SetChannelAccountConnectionRedirectUrl(v string) {
	o.ChannelAccountConnectionRedirectUrl = &v
}

// GetChannelLogoUrl returns the ChannelLogoUrl field value if set, zero value otherwise.
func (o *PublicChannelIntegrationChannel) GetChannelLogoUrl() string {
	if o == nil || IsNil(o.ChannelLogoUrl) {
		var ret string
		return ret
	}
	return *o.ChannelLogoUrl
}

// GetChannelLogoUrlOk returns a tuple with the ChannelLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelIntegrationChannel) GetChannelLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelLogoUrl) {
		return nil, false
	}
	return o.ChannelLogoUrl, true
}

// HasChannelLogoUrl returns a boolean if a field has been set.
func (o *PublicChannelIntegrationChannel) HasChannelLogoUrl() bool {
	if o != nil && !IsNil(o.ChannelLogoUrl) {
		return true
	}

	return false
}

// SetChannelLogoUrl gets a reference to the given string and assigns it to the ChannelLogoUrl field.
func (o *PublicChannelIntegrationChannel) SetChannelLogoUrl(v string) {
	o.ChannelLogoUrl = &v
}

// GetName returns the Name field value
func (o *PublicChannelIntegrationChannel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicChannelIntegrationChannel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicChannelIntegrationChannel) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *PublicChannelIntegrationChannel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicChannelIntegrationChannel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicChannelIntegrationChannel) SetId(v string) {
	o.Id = v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *PublicChannelIntegrationChannel) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelIntegrationChannel) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *PublicChannelIntegrationChannel) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *PublicChannelIntegrationChannel) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o PublicChannelIntegrationChannel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicChannelIntegrationChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelDescription) {
		toSerialize["channelDescription"] = o.ChannelDescription
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["capabilities"] = o.Capabilities
	if !IsNil(o.ChannelAccountConnectionRedirectUrl) {
		toSerialize["channelAccountConnectionRedirectUrl"] = o.ChannelAccountConnectionRedirectUrl
	}
	if !IsNil(o.ChannelLogoUrl) {
		toSerialize["channelLogoUrl"] = o.ChannelLogoUrl
	}
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	return toSerialize, nil
}

func (o *PublicChannelIntegrationChannel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"capabilities",
		"name",
		"id",
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

	varPublicChannelIntegrationChannel := _PublicChannelIntegrationChannel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicChannelIntegrationChannel)

	if err != nil {
		return err
	}

	*o = PublicChannelIntegrationChannel(varPublicChannelIntegrationChannel)

	return err
}

type NullablePublicChannelIntegrationChannel struct {
	value *PublicChannelIntegrationChannel
	isSet bool
}

func (v NullablePublicChannelIntegrationChannel) Get() *PublicChannelIntegrationChannel {
	return v.value
}

func (v *NullablePublicChannelIntegrationChannel) Set(val *PublicChannelIntegrationChannel) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicChannelIntegrationChannel) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicChannelIntegrationChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicChannelIntegrationChannel(val *PublicChannelIntegrationChannel) *NullablePublicChannelIntegrationChannel {
	return &NullablePublicChannelIntegrationChannel{value: val, isSet: true}
}

func (v NullablePublicChannelIntegrationChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicChannelIntegrationChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


