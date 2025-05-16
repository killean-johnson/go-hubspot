# PublicCommentEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "COMMENT"]
**Text** | **string** |  | 
**RichText** | Pointer to **string** |  | [optional] 
**Attachments** | [**[]PublicCommentEggAllOfAttachments**](PublicCommentEggAllOfAttachments.md) |  | 
**Recipients** | [**[]PublicRecipientEgg**](PublicRecipientEgg.md) |  | 
**SenderActorId** | **string** |  | 
**ChannelId** | **string** |  | 
**ChannelAccountId** | **string** |  | 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicCommentEgg

`func NewPublicCommentEgg(type_ string, text string, attachments []PublicCommentEggAllOfAttachments, recipients []PublicRecipientEgg, senderActorId string, channelId string, channelAccountId string, ) *PublicCommentEgg`

NewPublicCommentEgg instantiates a new PublicCommentEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCommentEggWithDefaults

`func NewPublicCommentEggWithDefaults() *PublicCommentEgg`

NewPublicCommentEggWithDefaults instantiates a new PublicCommentEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicCommentEgg) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicCommentEgg) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicCommentEgg) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *PublicCommentEgg) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PublicCommentEgg) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PublicCommentEgg) SetText(v string)`

SetText sets Text field to given value.


### GetRichText

`func (o *PublicCommentEgg) GetRichText() string`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *PublicCommentEgg) GetRichTextOk() (*string, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *PublicCommentEgg) SetRichText(v string)`

SetRichText sets RichText field to given value.

### HasRichText

`func (o *PublicCommentEgg) HasRichText() bool`

HasRichText returns a boolean if a field has been set.

### GetAttachments

`func (o *PublicCommentEgg) GetAttachments() []PublicCommentEggAllOfAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PublicCommentEgg) GetAttachmentsOk() (*[]PublicCommentEggAllOfAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PublicCommentEgg) SetAttachments(v []PublicCommentEggAllOfAttachments)`

SetAttachments sets Attachments field to given value.


### GetRecipients

`func (o *PublicCommentEgg) GetRecipients() []PublicRecipientEgg`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PublicCommentEgg) GetRecipientsOk() (*[]PublicRecipientEgg, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PublicCommentEgg) SetRecipients(v []PublicRecipientEgg)`

SetRecipients sets Recipients field to given value.


### GetSenderActorId

`func (o *PublicCommentEgg) GetSenderActorId() string`

GetSenderActorId returns the SenderActorId field if non-nil, zero value otherwise.

### GetSenderActorIdOk

`func (o *PublicCommentEgg) GetSenderActorIdOk() (*string, bool)`

GetSenderActorIdOk returns a tuple with the SenderActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderActorId

`func (o *PublicCommentEgg) SetSenderActorId(v string)`

SetSenderActorId sets SenderActorId field to given value.


### GetChannelId

`func (o *PublicCommentEgg) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PublicCommentEgg) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PublicCommentEgg) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetChannelAccountId

`func (o *PublicCommentEgg) GetChannelAccountId() string`

GetChannelAccountId returns the ChannelAccountId field if non-nil, zero value otherwise.

### GetChannelAccountIdOk

`func (o *PublicCommentEgg) GetChannelAccountIdOk() (*string, bool)`

GetChannelAccountIdOk returns a tuple with the ChannelAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAccountId

`func (o *PublicCommentEgg) SetChannelAccountId(v string)`

SetChannelAccountId sets ChannelAccountId field to given value.


### GetSubject

`func (o *PublicCommentEgg) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PublicCommentEgg) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PublicCommentEgg) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PublicCommentEgg) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


