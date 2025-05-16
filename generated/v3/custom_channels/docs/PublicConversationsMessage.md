# PublicConversationsMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]PublicConversationsMessageAttachmentsInner**](PublicConversationsMessageAttachmentsInner.md) |  | 
**Subject** | Pointer to **string** |  | [optional] 
**ConversationsThreadId** | **string** |  | 
**Type** | **string** |  | [default to "MESSAGE"]
**RichText** | Pointer to **string** |  | [optional] 
**InReplyToId** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Archived** | **bool** |  | 
**CreatedBy** | **string** |  | 
**Recipients** | [**[]PublicRecipient**](PublicRecipient.md) |  | 
**TruncationStatus** | **string** |  | 
**Client** | [**PublicClient**](PublicClient.md) |  | 
**Id** | **string** |  | 
**Text** | **string** |  | 
**ChannelAccountId** | **string** |  | 
**Senders** | [**[]PublicSender**](PublicSender.md) |  | 
**ChannelId** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**PublicMessageStatus**](PublicMessageStatus.md) |  | [optional] 
**Direction** | **string** |  | 

## Methods

### NewPublicConversationsMessage

`func NewPublicConversationsMessage(attachments []PublicConversationsMessageAttachmentsInner, conversationsThreadId string, type_ string, createdAt time.Time, archived bool, createdBy string, recipients []PublicRecipient, truncationStatus string, client PublicClient, id string, text string, channelAccountId string, senders []PublicSender, channelId string, direction string, ) *PublicConversationsMessage`

NewPublicConversationsMessage instantiates a new PublicConversationsMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicConversationsMessageWithDefaults

`func NewPublicConversationsMessageWithDefaults() *PublicConversationsMessage`

NewPublicConversationsMessageWithDefaults instantiates a new PublicConversationsMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *PublicConversationsMessage) GetAttachments() []PublicConversationsMessageAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PublicConversationsMessage) GetAttachmentsOk() (*[]PublicConversationsMessageAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PublicConversationsMessage) SetAttachments(v []PublicConversationsMessageAttachmentsInner)`

SetAttachments sets Attachments field to given value.


### GetSubject

`func (o *PublicConversationsMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PublicConversationsMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PublicConversationsMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PublicConversationsMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetConversationsThreadId

`func (o *PublicConversationsMessage) GetConversationsThreadId() string`

GetConversationsThreadId returns the ConversationsThreadId field if non-nil, zero value otherwise.

### GetConversationsThreadIdOk

`func (o *PublicConversationsMessage) GetConversationsThreadIdOk() (*string, bool)`

GetConversationsThreadIdOk returns a tuple with the ConversationsThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationsThreadId

`func (o *PublicConversationsMessage) SetConversationsThreadId(v string)`

SetConversationsThreadId sets ConversationsThreadId field to given value.


### GetType

`func (o *PublicConversationsMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicConversationsMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicConversationsMessage) SetType(v string)`

SetType sets Type field to given value.


### GetRichText

`func (o *PublicConversationsMessage) GetRichText() string`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *PublicConversationsMessage) GetRichTextOk() (*string, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *PublicConversationsMessage) SetRichText(v string)`

SetRichText sets RichText field to given value.

### HasRichText

`func (o *PublicConversationsMessage) HasRichText() bool`

HasRichText returns a boolean if a field has been set.

### GetInReplyToId

`func (o *PublicConversationsMessage) GetInReplyToId() string`

GetInReplyToId returns the InReplyToId field if non-nil, zero value otherwise.

### GetInReplyToIdOk

`func (o *PublicConversationsMessage) GetInReplyToIdOk() (*string, bool)`

GetInReplyToIdOk returns a tuple with the InReplyToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToId

`func (o *PublicConversationsMessage) SetInReplyToId(v string)`

SetInReplyToId sets InReplyToId field to given value.

### HasInReplyToId

`func (o *PublicConversationsMessage) HasInReplyToId() bool`

HasInReplyToId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicConversationsMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicConversationsMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicConversationsMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetArchived

`func (o *PublicConversationsMessage) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PublicConversationsMessage) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PublicConversationsMessage) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCreatedBy

`func (o *PublicConversationsMessage) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PublicConversationsMessage) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PublicConversationsMessage) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetRecipients

`func (o *PublicConversationsMessage) GetRecipients() []PublicRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PublicConversationsMessage) GetRecipientsOk() (*[]PublicRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PublicConversationsMessage) SetRecipients(v []PublicRecipient)`

SetRecipients sets Recipients field to given value.


### GetTruncationStatus

`func (o *PublicConversationsMessage) GetTruncationStatus() string`

GetTruncationStatus returns the TruncationStatus field if non-nil, zero value otherwise.

### GetTruncationStatusOk

`func (o *PublicConversationsMessage) GetTruncationStatusOk() (*string, bool)`

GetTruncationStatusOk returns a tuple with the TruncationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncationStatus

`func (o *PublicConversationsMessage) SetTruncationStatus(v string)`

SetTruncationStatus sets TruncationStatus field to given value.


### GetClient

`func (o *PublicConversationsMessage) GetClient() PublicClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *PublicConversationsMessage) GetClientOk() (*PublicClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *PublicConversationsMessage) SetClient(v PublicClient)`

SetClient sets Client field to given value.


### GetId

`func (o *PublicConversationsMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicConversationsMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicConversationsMessage) SetId(v string)`

SetId sets Id field to given value.


### GetText

`func (o *PublicConversationsMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PublicConversationsMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PublicConversationsMessage) SetText(v string)`

SetText sets Text field to given value.


### GetChannelAccountId

`func (o *PublicConversationsMessage) GetChannelAccountId() string`

GetChannelAccountId returns the ChannelAccountId field if non-nil, zero value otherwise.

### GetChannelAccountIdOk

`func (o *PublicConversationsMessage) GetChannelAccountIdOk() (*string, bool)`

GetChannelAccountIdOk returns a tuple with the ChannelAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAccountId

`func (o *PublicConversationsMessage) SetChannelAccountId(v string)`

SetChannelAccountId sets ChannelAccountId field to given value.


### GetSenders

`func (o *PublicConversationsMessage) GetSenders() []PublicSender`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *PublicConversationsMessage) GetSendersOk() (*[]PublicSender, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *PublicConversationsMessage) SetSenders(v []PublicSender)`

SetSenders sets Senders field to given value.


### GetChannelId

`func (o *PublicConversationsMessage) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PublicConversationsMessage) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PublicConversationsMessage) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetUpdatedAt

`func (o *PublicConversationsMessage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicConversationsMessage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicConversationsMessage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicConversationsMessage) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PublicConversationsMessage) GetStatus() PublicMessageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicConversationsMessage) GetStatusOk() (*PublicMessageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicConversationsMessage) SetStatus(v PublicMessageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicConversationsMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDirection

`func (o *PublicConversationsMessage) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PublicConversationsMessage) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PublicConversationsMessage) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


