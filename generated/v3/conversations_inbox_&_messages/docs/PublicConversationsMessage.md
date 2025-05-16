# PublicConversationsMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "MESSAGE"]
**Id** | **string** |  | 
**ConversationsThreadId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | **string** |  | 
**Client** | [**PublicClient**](PublicClient.md) |  | 
**Senders** | [**[]PublicSender**](PublicSender.md) |  | 
**Recipients** | [**[]PublicRecipient**](PublicRecipient.md) |  | 
**Archived** | **bool** |  | 
**Text** | **string** |  | 
**RichText** | **string** |  | 
**Attachments** | [**[]PublicConversationsMessageAllOfAttachments**](PublicConversationsMessageAllOfAttachments.md) |  | 
**Subject** | Pointer to **string** |  | [optional] 
**TruncationStatus** | **string** |  | 
**InReplyToId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PublicMessageStatus**](PublicMessageStatus.md) |  | [optional] 
**Direction** | **string** |  | 
**ChannelId** | **string** |  | 
**ChannelAccountId** | **string** |  | 
**AssignedTo** | Pointer to **string** |  | [optional] 
**AssignedFrom** | Pointer to **string** |  | [optional] 
**NewStatus** | **string** |  | 
**FromInboxId** | **string** |  | 
**ToInboxId** | **string** |  | 

## Methods

### NewPublicConversationsMessage

`func NewPublicConversationsMessage(type_ string, id string, conversationsThreadId string, createdAt time.Time, createdBy string, client PublicClient, senders []PublicSender, recipients []PublicRecipient, archived bool, text string, richText string, attachments []PublicConversationsMessageAllOfAttachments, truncationStatus string, direction string, channelId string, channelAccountId string, newStatus string, fromInboxId string, toInboxId string, ) *PublicConversationsMessage`

NewPublicConversationsMessage instantiates a new PublicConversationsMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicConversationsMessageWithDefaults

`func NewPublicConversationsMessageWithDefaults() *PublicConversationsMessage`

NewPublicConversationsMessageWithDefaults instantiates a new PublicConversationsMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetAttachments

`func (o *PublicConversationsMessage) GetAttachments() []PublicConversationsMessageAllOfAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PublicConversationsMessage) GetAttachmentsOk() (*[]PublicConversationsMessageAllOfAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PublicConversationsMessage) SetAttachments(v []PublicConversationsMessageAllOfAttachments)`

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


### GetAssignedTo

`func (o *PublicConversationsMessage) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *PublicConversationsMessage) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *PublicConversationsMessage) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *PublicConversationsMessage) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetAssignedFrom

`func (o *PublicConversationsMessage) GetAssignedFrom() string`

GetAssignedFrom returns the AssignedFrom field if non-nil, zero value otherwise.

### GetAssignedFromOk

`func (o *PublicConversationsMessage) GetAssignedFromOk() (*string, bool)`

GetAssignedFromOk returns a tuple with the AssignedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedFrom

`func (o *PublicConversationsMessage) SetAssignedFrom(v string)`

SetAssignedFrom sets AssignedFrom field to given value.

### HasAssignedFrom

`func (o *PublicConversationsMessage) HasAssignedFrom() bool`

HasAssignedFrom returns a boolean if a field has been set.

### GetNewStatus

`func (o *PublicConversationsMessage) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *PublicConversationsMessage) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *PublicConversationsMessage) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.


### GetFromInboxId

`func (o *PublicConversationsMessage) GetFromInboxId() string`

GetFromInboxId returns the FromInboxId field if non-nil, zero value otherwise.

### GetFromInboxIdOk

`func (o *PublicConversationsMessage) GetFromInboxIdOk() (*string, bool)`

GetFromInboxIdOk returns a tuple with the FromInboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromInboxId

`func (o *PublicConversationsMessage) SetFromInboxId(v string)`

SetFromInboxId sets FromInboxId field to given value.


### GetToInboxId

`func (o *PublicConversationsMessage) GetToInboxId() string`

GetToInboxId returns the ToInboxId field if non-nil, zero value otherwise.

### GetToInboxIdOk

`func (o *PublicConversationsMessage) GetToInboxIdOk() (*string, bool)`

GetToInboxIdOk returns a tuple with the ToInboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInboxId

`func (o *PublicConversationsMessage) SetToInboxId(v string)`

SetToInboxId sets ToInboxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


