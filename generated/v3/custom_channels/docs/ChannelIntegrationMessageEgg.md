# ChannelIntegrationMessageEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageDirection** | **string** |  | 
**Attachments** | [**[]ChannelIntegrationMessageEggAttachmentsInner**](ChannelIntegrationMessageEggAttachmentsInner.md) |  | 
**Recipients** | [**[]ChannelIntegrationParticipant**](ChannelIntegrationParticipant.md) |  | 
**IntegrationThreadId** | **string** |  | 
**IntegrationIdempotencyId** | Pointer to **string** |  | [optional] 
**Text** | **string** |  | 
**RichText** | Pointer to **string** |  | [optional] 
**ChannelAccountId** | **string** |  | 
**Senders** | [**[]ChannelIntegrationParticipant**](ChannelIntegrationParticipant.md) |  | 
**InReplyToId** | Pointer to **string** |  | [optional] 
**Timestamp** | **time.Time** |  | 

## Methods

### NewChannelIntegrationMessageEgg

`func NewChannelIntegrationMessageEgg(messageDirection string, attachments []ChannelIntegrationMessageEggAttachmentsInner, recipients []ChannelIntegrationParticipant, integrationThreadId string, text string, channelAccountId string, senders []ChannelIntegrationParticipant, timestamp time.Time, ) *ChannelIntegrationMessageEgg`

NewChannelIntegrationMessageEgg instantiates a new ChannelIntegrationMessageEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelIntegrationMessageEggWithDefaults

`func NewChannelIntegrationMessageEggWithDefaults() *ChannelIntegrationMessageEgg`

NewChannelIntegrationMessageEggWithDefaults instantiates a new ChannelIntegrationMessageEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageDirection

`func (o *ChannelIntegrationMessageEgg) GetMessageDirection() string`

GetMessageDirection returns the MessageDirection field if non-nil, zero value otherwise.

### GetMessageDirectionOk

`func (o *ChannelIntegrationMessageEgg) GetMessageDirectionOk() (*string, bool)`

GetMessageDirectionOk returns a tuple with the MessageDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDirection

`func (o *ChannelIntegrationMessageEgg) SetMessageDirection(v string)`

SetMessageDirection sets MessageDirection field to given value.


### GetAttachments

`func (o *ChannelIntegrationMessageEgg) GetAttachments() []ChannelIntegrationMessageEggAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ChannelIntegrationMessageEgg) GetAttachmentsOk() (*[]ChannelIntegrationMessageEggAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ChannelIntegrationMessageEgg) SetAttachments(v []ChannelIntegrationMessageEggAttachmentsInner)`

SetAttachments sets Attachments field to given value.


### GetRecipients

`func (o *ChannelIntegrationMessageEgg) GetRecipients() []ChannelIntegrationParticipant`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ChannelIntegrationMessageEgg) GetRecipientsOk() (*[]ChannelIntegrationParticipant, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ChannelIntegrationMessageEgg) SetRecipients(v []ChannelIntegrationParticipant)`

SetRecipients sets Recipients field to given value.


### GetIntegrationThreadId

`func (o *ChannelIntegrationMessageEgg) GetIntegrationThreadId() string`

GetIntegrationThreadId returns the IntegrationThreadId field if non-nil, zero value otherwise.

### GetIntegrationThreadIdOk

`func (o *ChannelIntegrationMessageEgg) GetIntegrationThreadIdOk() (*string, bool)`

GetIntegrationThreadIdOk returns a tuple with the IntegrationThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationThreadId

`func (o *ChannelIntegrationMessageEgg) SetIntegrationThreadId(v string)`

SetIntegrationThreadId sets IntegrationThreadId field to given value.


### GetIntegrationIdempotencyId

`func (o *ChannelIntegrationMessageEgg) GetIntegrationIdempotencyId() string`

GetIntegrationIdempotencyId returns the IntegrationIdempotencyId field if non-nil, zero value otherwise.

### GetIntegrationIdempotencyIdOk

`func (o *ChannelIntegrationMessageEgg) GetIntegrationIdempotencyIdOk() (*string, bool)`

GetIntegrationIdempotencyIdOk returns a tuple with the IntegrationIdempotencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIdempotencyId

`func (o *ChannelIntegrationMessageEgg) SetIntegrationIdempotencyId(v string)`

SetIntegrationIdempotencyId sets IntegrationIdempotencyId field to given value.

### HasIntegrationIdempotencyId

`func (o *ChannelIntegrationMessageEgg) HasIntegrationIdempotencyId() bool`

HasIntegrationIdempotencyId returns a boolean if a field has been set.

### GetText

`func (o *ChannelIntegrationMessageEgg) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChannelIntegrationMessageEgg) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChannelIntegrationMessageEgg) SetText(v string)`

SetText sets Text field to given value.


### GetRichText

`func (o *ChannelIntegrationMessageEgg) GetRichText() string`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *ChannelIntegrationMessageEgg) GetRichTextOk() (*string, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *ChannelIntegrationMessageEgg) SetRichText(v string)`

SetRichText sets RichText field to given value.

### HasRichText

`func (o *ChannelIntegrationMessageEgg) HasRichText() bool`

HasRichText returns a boolean if a field has been set.

### GetChannelAccountId

`func (o *ChannelIntegrationMessageEgg) GetChannelAccountId() string`

GetChannelAccountId returns the ChannelAccountId field if non-nil, zero value otherwise.

### GetChannelAccountIdOk

`func (o *ChannelIntegrationMessageEgg) GetChannelAccountIdOk() (*string, bool)`

GetChannelAccountIdOk returns a tuple with the ChannelAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAccountId

`func (o *ChannelIntegrationMessageEgg) SetChannelAccountId(v string)`

SetChannelAccountId sets ChannelAccountId field to given value.


### GetSenders

`func (o *ChannelIntegrationMessageEgg) GetSenders() []ChannelIntegrationParticipant`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ChannelIntegrationMessageEgg) GetSendersOk() (*[]ChannelIntegrationParticipant, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ChannelIntegrationMessageEgg) SetSenders(v []ChannelIntegrationParticipant)`

SetSenders sets Senders field to given value.


### GetInReplyToId

`func (o *ChannelIntegrationMessageEgg) GetInReplyToId() string`

GetInReplyToId returns the InReplyToId field if non-nil, zero value otherwise.

### GetInReplyToIdOk

`func (o *ChannelIntegrationMessageEgg) GetInReplyToIdOk() (*string, bool)`

GetInReplyToIdOk returns a tuple with the InReplyToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToId

`func (o *ChannelIntegrationMessageEgg) SetInReplyToId(v string)`

SetInReplyToId sets InReplyToId field to given value.

### HasInReplyToId

`func (o *ChannelIntegrationMessageEgg) HasInReplyToId() bool`

HasInReplyToId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ChannelIntegrationMessageEgg) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ChannelIntegrationMessageEgg) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ChannelIntegrationMessageEgg) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


