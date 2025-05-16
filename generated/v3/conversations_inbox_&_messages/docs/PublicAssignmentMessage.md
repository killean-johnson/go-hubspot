# PublicAssignmentMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "ASSIGNMENT"]
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
**Attachments** | [**[]PublicCommentAllOfAttachments**](PublicCommentAllOfAttachments.md) |  | 
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

### NewPublicAssignmentMessage

`func NewPublicAssignmentMessage(type_ string, id string, conversationsThreadId string, createdAt time.Time, createdBy string, client PublicClient, senders []PublicSender, recipients []PublicRecipient, archived bool, text string, richText string, attachments []PublicCommentAllOfAttachments, truncationStatus string, direction string, channelId string, channelAccountId string, newStatus string, fromInboxId string, toInboxId string, ) *PublicAssignmentMessage`

NewPublicAssignmentMessage instantiates a new PublicAssignmentMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssignmentMessageWithDefaults

`func NewPublicAssignmentMessageWithDefaults() *PublicAssignmentMessage`

NewPublicAssignmentMessageWithDefaults instantiates a new PublicAssignmentMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicAssignmentMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicAssignmentMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicAssignmentMessage) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PublicAssignmentMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicAssignmentMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicAssignmentMessage) SetId(v string)`

SetId sets Id field to given value.


### GetConversationsThreadId

`func (o *PublicAssignmentMessage) GetConversationsThreadId() string`

GetConversationsThreadId returns the ConversationsThreadId field if non-nil, zero value otherwise.

### GetConversationsThreadIdOk

`func (o *PublicAssignmentMessage) GetConversationsThreadIdOk() (*string, bool)`

GetConversationsThreadIdOk returns a tuple with the ConversationsThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationsThreadId

`func (o *PublicAssignmentMessage) SetConversationsThreadId(v string)`

SetConversationsThreadId sets ConversationsThreadId field to given value.


### GetCreatedAt

`func (o *PublicAssignmentMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicAssignmentMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicAssignmentMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PublicAssignmentMessage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicAssignmentMessage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicAssignmentMessage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicAssignmentMessage) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PublicAssignmentMessage) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PublicAssignmentMessage) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PublicAssignmentMessage) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetClient

`func (o *PublicAssignmentMessage) GetClient() PublicClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *PublicAssignmentMessage) GetClientOk() (*PublicClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *PublicAssignmentMessage) SetClient(v PublicClient)`

SetClient sets Client field to given value.


### GetSenders

`func (o *PublicAssignmentMessage) GetSenders() []PublicSender`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *PublicAssignmentMessage) GetSendersOk() (*[]PublicSender, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *PublicAssignmentMessage) SetSenders(v []PublicSender)`

SetSenders sets Senders field to given value.


### GetRecipients

`func (o *PublicAssignmentMessage) GetRecipients() []PublicRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PublicAssignmentMessage) GetRecipientsOk() (*[]PublicRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PublicAssignmentMessage) SetRecipients(v []PublicRecipient)`

SetRecipients sets Recipients field to given value.


### GetArchived

`func (o *PublicAssignmentMessage) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PublicAssignmentMessage) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PublicAssignmentMessage) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetText

`func (o *PublicAssignmentMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PublicAssignmentMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PublicAssignmentMessage) SetText(v string)`

SetText sets Text field to given value.


### GetRichText

`func (o *PublicAssignmentMessage) GetRichText() string`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *PublicAssignmentMessage) GetRichTextOk() (*string, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *PublicAssignmentMessage) SetRichText(v string)`

SetRichText sets RichText field to given value.


### GetAttachments

`func (o *PublicAssignmentMessage) GetAttachments() []PublicCommentAllOfAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PublicAssignmentMessage) GetAttachmentsOk() (*[]PublicCommentAllOfAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PublicAssignmentMessage) SetAttachments(v []PublicCommentAllOfAttachments)`

SetAttachments sets Attachments field to given value.


### GetSubject

`func (o *PublicAssignmentMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PublicAssignmentMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PublicAssignmentMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PublicAssignmentMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTruncationStatus

`func (o *PublicAssignmentMessage) GetTruncationStatus() string`

GetTruncationStatus returns the TruncationStatus field if non-nil, zero value otherwise.

### GetTruncationStatusOk

`func (o *PublicAssignmentMessage) GetTruncationStatusOk() (*string, bool)`

GetTruncationStatusOk returns a tuple with the TruncationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncationStatus

`func (o *PublicAssignmentMessage) SetTruncationStatus(v string)`

SetTruncationStatus sets TruncationStatus field to given value.


### GetInReplyToId

`func (o *PublicAssignmentMessage) GetInReplyToId() string`

GetInReplyToId returns the InReplyToId field if non-nil, zero value otherwise.

### GetInReplyToIdOk

`func (o *PublicAssignmentMessage) GetInReplyToIdOk() (*string, bool)`

GetInReplyToIdOk returns a tuple with the InReplyToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToId

`func (o *PublicAssignmentMessage) SetInReplyToId(v string)`

SetInReplyToId sets InReplyToId field to given value.

### HasInReplyToId

`func (o *PublicAssignmentMessage) HasInReplyToId() bool`

HasInReplyToId returns a boolean if a field has been set.

### GetStatus

`func (o *PublicAssignmentMessage) GetStatus() PublicMessageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicAssignmentMessage) GetStatusOk() (*PublicMessageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicAssignmentMessage) SetStatus(v PublicMessageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicAssignmentMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDirection

`func (o *PublicAssignmentMessage) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PublicAssignmentMessage) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PublicAssignmentMessage) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetChannelId

`func (o *PublicAssignmentMessage) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PublicAssignmentMessage) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PublicAssignmentMessage) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetChannelAccountId

`func (o *PublicAssignmentMessage) GetChannelAccountId() string`

GetChannelAccountId returns the ChannelAccountId field if non-nil, zero value otherwise.

### GetChannelAccountIdOk

`func (o *PublicAssignmentMessage) GetChannelAccountIdOk() (*string, bool)`

GetChannelAccountIdOk returns a tuple with the ChannelAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAccountId

`func (o *PublicAssignmentMessage) SetChannelAccountId(v string)`

SetChannelAccountId sets ChannelAccountId field to given value.


### GetAssignedTo

`func (o *PublicAssignmentMessage) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *PublicAssignmentMessage) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *PublicAssignmentMessage) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *PublicAssignmentMessage) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetAssignedFrom

`func (o *PublicAssignmentMessage) GetAssignedFrom() string`

GetAssignedFrom returns the AssignedFrom field if non-nil, zero value otherwise.

### GetAssignedFromOk

`func (o *PublicAssignmentMessage) GetAssignedFromOk() (*string, bool)`

GetAssignedFromOk returns a tuple with the AssignedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedFrom

`func (o *PublicAssignmentMessage) SetAssignedFrom(v string)`

SetAssignedFrom sets AssignedFrom field to given value.

### HasAssignedFrom

`func (o *PublicAssignmentMessage) HasAssignedFrom() bool`

HasAssignedFrom returns a boolean if a field has been set.

### GetNewStatus

`func (o *PublicAssignmentMessage) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *PublicAssignmentMessage) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *PublicAssignmentMessage) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.


### GetFromInboxId

`func (o *PublicAssignmentMessage) GetFromInboxId() string`

GetFromInboxId returns the FromInboxId field if non-nil, zero value otherwise.

### GetFromInboxIdOk

`func (o *PublicAssignmentMessage) GetFromInboxIdOk() (*string, bool)`

GetFromInboxIdOk returns a tuple with the FromInboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromInboxId

`func (o *PublicAssignmentMessage) SetFromInboxId(v string)`

SetFromInboxId sets FromInboxId field to given value.


### GetToInboxId

`func (o *PublicAssignmentMessage) GetToInboxId() string`

GetToInboxId returns the ToInboxId field if non-nil, zero value otherwise.

### GetToInboxIdOk

`func (o *PublicAssignmentMessage) GetToInboxIdOk() (*string, bool)`

GetToInboxIdOk returns a tuple with the ToInboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInboxId

`func (o *PublicAssignmentMessage) SetToInboxId(v string)`

SetToInboxId sets ToInboxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


