# PublicConversationsMessageEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "MESSAGE"]
**Text** | **string** |  | 
**RichText** | Pointer to **string** |  | [optional] 
**Attachments** | [**[]PublicConversationsMessageEggAllOfAttachments**](PublicConversationsMessageEggAllOfAttachments.md) |  | 
**Recipients** | [**[]PublicRecipientEgg**](PublicRecipientEgg.md) |  | 
**SenderActorId** | **string** |  | 
**ChannelId** | **string** |  | 
**ChannelAccountId** | **string** |  | 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicConversationsMessageEgg

`func NewPublicConversationsMessageEgg(type_ string, text string, attachments []PublicConversationsMessageEggAllOfAttachments, recipients []PublicRecipientEgg, senderActorId string, channelId string, channelAccountId string, ) *PublicConversationsMessageEgg`

NewPublicConversationsMessageEgg instantiates a new PublicConversationsMessageEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicConversationsMessageEggWithDefaults

`func NewPublicConversationsMessageEggWithDefaults() *PublicConversationsMessageEgg`

NewPublicConversationsMessageEggWithDefaults instantiates a new PublicConversationsMessageEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicConversationsMessageEgg) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicConversationsMessageEgg) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicConversationsMessageEgg) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *PublicConversationsMessageEgg) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PublicConversationsMessageEgg) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PublicConversationsMessageEgg) SetText(v string)`

SetText sets Text field to given value.


### GetRichText

`func (o *PublicConversationsMessageEgg) GetRichText() string`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *PublicConversationsMessageEgg) GetRichTextOk() (*string, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *PublicConversationsMessageEgg) SetRichText(v string)`

SetRichText sets RichText field to given value.

### HasRichText

`func (o *PublicConversationsMessageEgg) HasRichText() bool`

HasRichText returns a boolean if a field has been set.

### GetAttachments

`func (o *PublicConversationsMessageEgg) GetAttachments() []PublicConversationsMessageEggAllOfAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PublicConversationsMessageEgg) GetAttachmentsOk() (*[]PublicConversationsMessageEggAllOfAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PublicConversationsMessageEgg) SetAttachments(v []PublicConversationsMessageEggAllOfAttachments)`

SetAttachments sets Attachments field to given value.


### GetRecipients

`func (o *PublicConversationsMessageEgg) GetRecipients() []PublicRecipientEgg`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PublicConversationsMessageEgg) GetRecipientsOk() (*[]PublicRecipientEgg, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PublicConversationsMessageEgg) SetRecipients(v []PublicRecipientEgg)`

SetRecipients sets Recipients field to given value.


### GetSenderActorId

`func (o *PublicConversationsMessageEgg) GetSenderActorId() string`

GetSenderActorId returns the SenderActorId field if non-nil, zero value otherwise.

### GetSenderActorIdOk

`func (o *PublicConversationsMessageEgg) GetSenderActorIdOk() (*string, bool)`

GetSenderActorIdOk returns a tuple with the SenderActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderActorId

`func (o *PublicConversationsMessageEgg) SetSenderActorId(v string)`

SetSenderActorId sets SenderActorId field to given value.


### GetChannelId

`func (o *PublicConversationsMessageEgg) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PublicConversationsMessageEgg) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PublicConversationsMessageEgg) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetChannelAccountId

`func (o *PublicConversationsMessageEgg) GetChannelAccountId() string`

GetChannelAccountId returns the ChannelAccountId field if non-nil, zero value otherwise.

### GetChannelAccountIdOk

`func (o *PublicConversationsMessageEgg) GetChannelAccountIdOk() (*string, bool)`

GetChannelAccountIdOk returns a tuple with the ChannelAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAccountId

`func (o *PublicConversationsMessageEgg) SetChannelAccountId(v string)`

SetChannelAccountId sets ChannelAccountId field to given value.


### GetSubject

`func (o *PublicConversationsMessageEgg) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PublicConversationsMessageEgg) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PublicConversationsMessageEgg) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PublicConversationsMessageEgg) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


