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

// checks if the ChannelIntegrationMessageEgg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelIntegrationMessageEgg{}

// ChannelIntegrationMessageEgg struct for ChannelIntegrationMessageEgg
type ChannelIntegrationMessageEgg struct {
	MessageDirection string `json:"messageDirection"`
	Attachments []ChannelIntegrationMessageEggAttachmentsInner `json:"attachments"`
	Recipients []ChannelIntegrationParticipant `json:"recipients"`
	IntegrationThreadId string `json:"integrationThreadId"`
	IntegrationIdempotencyId *string `json:"integrationIdempotencyId,omitempty"`
	Text string `json:"text"`
	RichText *string `json:"richText,omitempty"`
	ChannelAccountId string `json:"channelAccountId"`
	Senders []ChannelIntegrationParticipant `json:"senders"`
	InReplyToId *string `json:"inReplyToId,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type _ChannelIntegrationMessageEgg ChannelIntegrationMessageEgg

// NewChannelIntegrationMessageEgg instantiates a new ChannelIntegrationMessageEgg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelIntegrationMessageEgg(messageDirection string, attachments []ChannelIntegrationMessageEggAttachmentsInner, recipients []ChannelIntegrationParticipant, integrationThreadId string, text string, channelAccountId string, senders []ChannelIntegrationParticipant, timestamp time.Time) *ChannelIntegrationMessageEgg {
	this := ChannelIntegrationMessageEgg{}
	this.MessageDirection = messageDirection
	this.Attachments = attachments
	this.Recipients = recipients
	this.IntegrationThreadId = integrationThreadId
	this.Text = text
	this.ChannelAccountId = channelAccountId
	this.Senders = senders
	this.Timestamp = timestamp
	return &this
}

// NewChannelIntegrationMessageEggWithDefaults instantiates a new ChannelIntegrationMessageEgg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelIntegrationMessageEggWithDefaults() *ChannelIntegrationMessageEgg {
	this := ChannelIntegrationMessageEgg{}
	return &this
}

// GetMessageDirection returns the MessageDirection field value
func (o *ChannelIntegrationMessageEgg) GetMessageDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageDirection
}

// GetMessageDirectionOk returns a tuple with the MessageDirection field value
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetMessageDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageDirection, true
}

// SetMessageDirection sets field value
func (o *ChannelIntegrationMessageEgg) SetMessageDirection(v string) {
	o.MessageDirection = v
}

// GetAttachments returns the Attachments field value
func (o *ChannelIntegrationMessageEgg) GetAttachments() []ChannelIntegrationMessageEggAttachmentsInner {
	if o == nil {
		var ret []ChannelIntegrationMessageEggAttachmentsInner
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetAttachmentsOk() ([]ChannelIntegrationMessageEggAttachmentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// SetAttachments sets field value
func (o *ChannelIntegrationMessageEgg) SetAttachments(v []ChannelIntegrationMessageEggAttachmentsInner) {
	o.Attachments = v
}

// GetRecipients returns the Recipients field value
func (o *ChannelIntegrationMessageEgg) GetRecipients() []ChannelIntegrationParticipant {
	if o == nil {
		var ret []ChannelIntegrationParticipant
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetRecipientsOk() ([]ChannelIntegrationParticipant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ChannelIntegrationMessageEgg) SetRecipients(v []ChannelIntegrationParticipant) {
	o.Recipients = v
}

// GetIntegrationThreadId returns the IntegrationThreadId field value
func (o *ChannelIntegrationMessageEgg) GetIntegrationThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationThreadId
}

// GetIntegrationThreadIdOk returns a tuple with the IntegrationThreadId field value
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetIntegrationThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationThreadId, true
}

// SetIntegrationThreadId sets field value
func (o *ChannelIntegrationMessageEgg) SetIntegrationThreadId(v string) {
	o.IntegrationThreadId = v
}

// GetIntegrationIdempotencyId returns the IntegrationIdempotencyId field value if set, zero value otherwise.
func (o *ChannelIntegrationMessageEgg) GetIntegrationIdempotencyId() string {
	if o == nil || IsNil(o.IntegrationIdempotencyId) {
		var ret string
		return ret
	}
	return *o.IntegrationIdempotencyId
}

// GetIntegrationIdempotencyIdOk returns a tuple with the IntegrationIdempotencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetIntegrationIdempotencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationIdempotencyId) {
		return nil, false
	}
	return o.IntegrationIdempotencyId, true
}

// HasIntegrationIdempotencyId returns a boolean if a field has been set.
func (o *ChannelIntegrationMessageEgg) HasIntegrationIdempotencyId() bool {
	if o != nil && !IsNil(o.IntegrationIdempotencyId) {
		return true
	}

	return false
}

// SetIntegrationIdempotencyId gets a reference to the given string and assigns it to the IntegrationIdempotencyId field.
func (o *ChannelIntegrationMessageEgg) SetIntegrationIdempotencyId(v string) {
	o.IntegrationIdempotencyId = &v
}

// GetText returns the Text field value
func (o *ChannelIntegrationMessageEgg) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *ChannelIntegrationMessageEgg) SetText(v string) {
	o.Text = v
}

// GetRichText returns the RichText field value if set, zero value otherwise.
func (o *ChannelIntegrationMessageEgg) GetRichText() string {
	if o == nil || IsNil(o.RichText) {
		var ret string
		return ret
	}
	return *o.RichText
}

// GetRichTextOk returns a tuple with the RichText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetRichTextOk() (*string, bool) {
	if o == nil || IsNil(o.RichText) {
		return nil, false
	}
	return o.RichText, true
}

// HasRichText returns a boolean if a field has been set.
func (o *ChannelIntegrationMessageEgg) HasRichText() bool {
	if o != nil && !IsNil(o.RichText) {
		return true
	}

	return false
}

// SetRichText gets a reference to the given string and assigns it to the RichText field.
func (o *ChannelIntegrationMessageEgg) SetRichText(v string) {
	o.RichText = &v
}

// GetChannelAccountId returns the ChannelAccountId field value
func (o *ChannelIntegrationMessageEgg) GetChannelAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelAccountId
}

// GetChannelAccountIdOk returns a tuple with the ChannelAccountId field value
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetChannelAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelAccountId, true
}

// SetChannelAccountId sets field value
func (o *ChannelIntegrationMessageEgg) SetChannelAccountId(v string) {
	o.ChannelAccountId = v
}

// GetSenders returns the Senders field value
func (o *ChannelIntegrationMessageEgg) GetSenders() []ChannelIntegrationParticipant {
	if o == nil {
		var ret []ChannelIntegrationParticipant
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetSendersOk() ([]ChannelIntegrationParticipant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Senders, true
}

// SetSenders sets field value
func (o *ChannelIntegrationMessageEgg) SetSenders(v []ChannelIntegrationParticipant) {
	o.Senders = v
}

// GetInReplyToId returns the InReplyToId field value if set, zero value otherwise.
func (o *ChannelIntegrationMessageEgg) GetInReplyToId() string {
	if o == nil || IsNil(o.InReplyToId) {
		var ret string
		return ret
	}
	return *o.InReplyToId
}

// GetInReplyToIdOk returns a tuple with the InReplyToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetInReplyToIdOk() (*string, bool) {
	if o == nil || IsNil(o.InReplyToId) {
		return nil, false
	}
	return o.InReplyToId, true
}

// HasInReplyToId returns a boolean if a field has been set.
func (o *ChannelIntegrationMessageEgg) HasInReplyToId() bool {
	if o != nil && !IsNil(o.InReplyToId) {
		return true
	}

	return false
}

// SetInReplyToId gets a reference to the given string and assigns it to the InReplyToId field.
func (o *ChannelIntegrationMessageEgg) SetInReplyToId(v string) {
	o.InReplyToId = &v
}

// GetTimestamp returns the Timestamp field value
func (o *ChannelIntegrationMessageEgg) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ChannelIntegrationMessageEgg) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ChannelIntegrationMessageEgg) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o ChannelIntegrationMessageEgg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelIntegrationMessageEgg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messageDirection"] = o.MessageDirection
	toSerialize["attachments"] = o.Attachments
	toSerialize["recipients"] = o.Recipients
	toSerialize["integrationThreadId"] = o.IntegrationThreadId
	if !IsNil(o.IntegrationIdempotencyId) {
		toSerialize["integrationIdempotencyId"] = o.IntegrationIdempotencyId
	}
	toSerialize["text"] = o.Text
	if !IsNil(o.RichText) {
		toSerialize["richText"] = o.RichText
	}
	toSerialize["channelAccountId"] = o.ChannelAccountId
	toSerialize["senders"] = o.Senders
	if !IsNil(o.InReplyToId) {
		toSerialize["inReplyToId"] = o.InReplyToId
	}
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *ChannelIntegrationMessageEgg) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messageDirection",
		"attachments",
		"recipients",
		"integrationThreadId",
		"text",
		"channelAccountId",
		"senders",
		"timestamp",
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

	varChannelIntegrationMessageEgg := _ChannelIntegrationMessageEgg{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelIntegrationMessageEgg)

	if err != nil {
		return err
	}

	*o = ChannelIntegrationMessageEgg(varChannelIntegrationMessageEgg)

	return err
}

type NullableChannelIntegrationMessageEgg struct {
	value *ChannelIntegrationMessageEgg
	isSet bool
}

func (v NullableChannelIntegrationMessageEgg) Get() *ChannelIntegrationMessageEgg {
	return v.value
}

func (v *NullableChannelIntegrationMessageEgg) Set(val *ChannelIntegrationMessageEgg) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelIntegrationMessageEgg) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelIntegrationMessageEgg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelIntegrationMessageEgg(val *ChannelIntegrationMessageEgg) *NullableChannelIntegrationMessageEgg {
	return &NullableChannelIntegrationMessageEgg{value: val, isSet: true}
}

func (v NullableChannelIntegrationMessageEgg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelIntegrationMessageEgg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


