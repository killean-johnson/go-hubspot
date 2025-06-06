# CollectionResponsePublicMessageForwardPagingResultsInner

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
**Attachments** | [**[]PublicCommentAllOfAttachments**](PublicCommentAllOfAttachments.md) |  | 
**ChannelId** | **string** |  | 
**ChannelAccountId** | **string** |  | 
**AssignedTo** | Pointer to **string** |  | [optional] 
**AssignedFrom** | Pointer to **string** |  | [optional] 
**NewStatus** | **string** |  | 
**FromInboxId** | **string** |  | 
**ToInboxId** | **string** |  | 
**Subject** | Pointer to **string** |  | [optional] 
**TruncationStatus** | **string** |  | 
**InReplyToId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PublicMessageStatus**](PublicMessageStatus.md) |  | [optional] 
**Direction** | **string** |  | 

## Methods

### NewCollectionResponsePublicMessageForwardPagingResultsInner

`func NewCollectionResponsePublicMessageForwardPagingResultsInner(type_ string, id string, conversationsThreadId string, createdAt time.Time, createdBy string, client PublicClient, senders []PublicSender, recipients []PublicRecipient, archived bool, text string, richText string, attachments []PublicCommentAllOfAttachments, channelId string, channelAccountId string, newStatus string, fromInboxId string, toInboxId string, truncationStatus string, direction string, ) *CollectionResponsePublicMessageForwardPagingResultsInner`

NewCollectionResponsePublicMessageForwardPagingResultsInner instantiates a new CollectionResponsePublicMessageForwardPagingResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponsePublicMessageForwardPagingResultsInnerWithDefaults

`func NewCollectionResponsePublicMessageForwardPagingResultsInnerWithDefaults() *CollectionResponsePublicMessageForwardPagingResultsInner`

NewCollectionResponsePublicMessageForwardPagingResultsInnerWithDefaults instantiates a new CollectionResponsePublicMessageForwardPagingResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetConversationsThreadId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetConversationsThreadId() string`

GetConversationsThreadId returns the ConversationsThreadId field if non-nil, zero value otherwise.

### GetConversationsThreadIdOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetConversationsThreadIdOk() (*string, bool)`

GetConversationsThreadIdOk returns a tuple with the ConversationsThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationsThreadId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetConversationsThreadId(v string)`

SetConversationsThreadId sets ConversationsThreadId field to given value.


### GetCreatedAt

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetClient

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetClient() PublicClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetClientOk() (*PublicClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetClient(v PublicClient)`

SetClient sets Client field to given value.


### GetSenders

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetSenders() []PublicSender`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetSendersOk() (*[]PublicSender, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetSenders(v []PublicSender)`

SetSenders sets Senders field to given value.


### GetRecipients

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetRecipients() []PublicRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetRecipientsOk() (*[]PublicRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetRecipients(v []PublicRecipient)`

SetRecipients sets Recipients field to given value.


### GetArchived

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetText

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetText(v string)`

SetText sets Text field to given value.


### GetRichText

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetRichText() string`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetRichTextOk() (*string, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetRichText(v string)`

SetRichText sets RichText field to given value.


### GetAttachments

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetAttachments() []PublicCommentAllOfAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetAttachmentsOk() (*[]PublicCommentAllOfAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetAttachments(v []PublicCommentAllOfAttachments)`

SetAttachments sets Attachments field to given value.


### GetChannelId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetChannelAccountId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetChannelAccountId() string`

GetChannelAccountId returns the ChannelAccountId field if non-nil, zero value otherwise.

### GetChannelAccountIdOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetChannelAccountIdOk() (*string, bool)`

GetChannelAccountIdOk returns a tuple with the ChannelAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAccountId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetChannelAccountId(v string)`

SetChannelAccountId sets ChannelAccountId field to given value.


### GetAssignedTo

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetAssignedFrom

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetAssignedFrom() string`

GetAssignedFrom returns the AssignedFrom field if non-nil, zero value otherwise.

### GetAssignedFromOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetAssignedFromOk() (*string, bool)`

GetAssignedFromOk returns a tuple with the AssignedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedFrom

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetAssignedFrom(v string)`

SetAssignedFrom sets AssignedFrom field to given value.

### HasAssignedFrom

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) HasAssignedFrom() bool`

HasAssignedFrom returns a boolean if a field has been set.

### GetNewStatus

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.


### GetFromInboxId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetFromInboxId() string`

GetFromInboxId returns the FromInboxId field if non-nil, zero value otherwise.

### GetFromInboxIdOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetFromInboxIdOk() (*string, bool)`

GetFromInboxIdOk returns a tuple with the FromInboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromInboxId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetFromInboxId(v string)`

SetFromInboxId sets FromInboxId field to given value.


### GetToInboxId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetToInboxId() string`

GetToInboxId returns the ToInboxId field if non-nil, zero value otherwise.

### GetToInboxIdOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetToInboxIdOk() (*string, bool)`

GetToInboxIdOk returns a tuple with the ToInboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInboxId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetToInboxId(v string)`

SetToInboxId sets ToInboxId field to given value.


### GetSubject

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTruncationStatus

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetTruncationStatus() string`

GetTruncationStatus returns the TruncationStatus field if non-nil, zero value otherwise.

### GetTruncationStatusOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetTruncationStatusOk() (*string, bool)`

GetTruncationStatusOk returns a tuple with the TruncationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncationStatus

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetTruncationStatus(v string)`

SetTruncationStatus sets TruncationStatus field to given value.


### GetInReplyToId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetInReplyToId() string`

GetInReplyToId returns the InReplyToId field if non-nil, zero value otherwise.

### GetInReplyToIdOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetInReplyToIdOk() (*string, bool)`

GetInReplyToIdOk returns a tuple with the InReplyToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetInReplyToId(v string)`

SetInReplyToId sets InReplyToId field to given value.

### HasInReplyToId

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) HasInReplyToId() bool`

HasInReplyToId returns a boolean if a field has been set.

### GetStatus

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetStatus() PublicMessageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetStatusOk() (*PublicMessageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetStatus(v PublicMessageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDirection

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CollectionResponsePublicMessageForwardPagingResultsInner) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


