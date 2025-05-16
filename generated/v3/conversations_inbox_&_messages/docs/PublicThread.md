# PublicThread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedContactId** | **string** | The ID of the associated Contact in the CRM. If the Contact for the thread has not yet been added or created, the &#x60;associatedContactId&#x60; returned will be a visitorID and cannot be used to search for the Contact in the CRM. | 
**ThreadAssociations** | Pointer to [**PublicThreadAssociations**](PublicThreadAssociations.md) |  | [optional] 
**AssignedTo** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** | When the thread was created. | 
**Archived** | Pointer to **bool** | Whether this thread is archived. | [optional] 
**OriginalChannelId** | **string** |  | 
**LatestMessageTimestamp** | Pointer to **time.Time** | The time that the latest message was sent or received on the thread. | [optional] 
**LatestMessageSentTimestamp** | Pointer to **time.Time** | The time that the latest message was sent on the thread. | [optional] 
**OriginalChannelAccountId** | **string** |  | 
**Id** | **string** | The unique ID of the thread. | 
**ClosedAt** | Pointer to **time.Time** | When the thread was closed. Only set if the thread is closed. | [optional] 
**Spam** | **bool** | Whether the thread is marked as spam. | 
**InboxId** | **string** | The ID of the conversations inbox containing the thread. | 
**Status** | **string** | The thread&#39;s status: &#x60;OPEN&#x60; or &#x60;CLOSED&#x60;. | 
**LatestMessageReceivedTimestamp** | Pointer to **time.Time** | The time that the latest message was sent on the thread. | [optional] 

## Methods

### NewPublicThread

`func NewPublicThread(associatedContactId string, createdAt time.Time, originalChannelId string, originalChannelAccountId string, id string, spam bool, inboxId string, status string, ) *PublicThread`

NewPublicThread instantiates a new PublicThread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicThreadWithDefaults

`func NewPublicThreadWithDefaults() *PublicThread`

NewPublicThreadWithDefaults instantiates a new PublicThread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedContactId

`func (o *PublicThread) GetAssociatedContactId() string`

GetAssociatedContactId returns the AssociatedContactId field if non-nil, zero value otherwise.

### GetAssociatedContactIdOk

`func (o *PublicThread) GetAssociatedContactIdOk() (*string, bool)`

GetAssociatedContactIdOk returns a tuple with the AssociatedContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedContactId

`func (o *PublicThread) SetAssociatedContactId(v string)`

SetAssociatedContactId sets AssociatedContactId field to given value.


### GetThreadAssociations

`func (o *PublicThread) GetThreadAssociations() PublicThreadAssociations`

GetThreadAssociations returns the ThreadAssociations field if non-nil, zero value otherwise.

### GetThreadAssociationsOk

`func (o *PublicThread) GetThreadAssociationsOk() (*PublicThreadAssociations, bool)`

GetThreadAssociationsOk returns a tuple with the ThreadAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadAssociations

`func (o *PublicThread) SetThreadAssociations(v PublicThreadAssociations)`

SetThreadAssociations sets ThreadAssociations field to given value.

### HasThreadAssociations

`func (o *PublicThread) HasThreadAssociations() bool`

HasThreadAssociations returns a boolean if a field has been set.

### GetAssignedTo

`func (o *PublicThread) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *PublicThread) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *PublicThread) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *PublicThread) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicThread) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicThread) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicThread) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetArchived

`func (o *PublicThread) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PublicThread) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PublicThread) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PublicThread) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetOriginalChannelId

`func (o *PublicThread) GetOriginalChannelId() string`

GetOriginalChannelId returns the OriginalChannelId field if non-nil, zero value otherwise.

### GetOriginalChannelIdOk

`func (o *PublicThread) GetOriginalChannelIdOk() (*string, bool)`

GetOriginalChannelIdOk returns a tuple with the OriginalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalChannelId

`func (o *PublicThread) SetOriginalChannelId(v string)`

SetOriginalChannelId sets OriginalChannelId field to given value.


### GetLatestMessageTimestamp

`func (o *PublicThread) GetLatestMessageTimestamp() time.Time`

GetLatestMessageTimestamp returns the LatestMessageTimestamp field if non-nil, zero value otherwise.

### GetLatestMessageTimestampOk

`func (o *PublicThread) GetLatestMessageTimestampOk() (*time.Time, bool)`

GetLatestMessageTimestampOk returns a tuple with the LatestMessageTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMessageTimestamp

`func (o *PublicThread) SetLatestMessageTimestamp(v time.Time)`

SetLatestMessageTimestamp sets LatestMessageTimestamp field to given value.

### HasLatestMessageTimestamp

`func (o *PublicThread) HasLatestMessageTimestamp() bool`

HasLatestMessageTimestamp returns a boolean if a field has been set.

### GetLatestMessageSentTimestamp

`func (o *PublicThread) GetLatestMessageSentTimestamp() time.Time`

GetLatestMessageSentTimestamp returns the LatestMessageSentTimestamp field if non-nil, zero value otherwise.

### GetLatestMessageSentTimestampOk

`func (o *PublicThread) GetLatestMessageSentTimestampOk() (*time.Time, bool)`

GetLatestMessageSentTimestampOk returns a tuple with the LatestMessageSentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMessageSentTimestamp

`func (o *PublicThread) SetLatestMessageSentTimestamp(v time.Time)`

SetLatestMessageSentTimestamp sets LatestMessageSentTimestamp field to given value.

### HasLatestMessageSentTimestamp

`func (o *PublicThread) HasLatestMessageSentTimestamp() bool`

HasLatestMessageSentTimestamp returns a boolean if a field has been set.

### GetOriginalChannelAccountId

`func (o *PublicThread) GetOriginalChannelAccountId() string`

GetOriginalChannelAccountId returns the OriginalChannelAccountId field if non-nil, zero value otherwise.

### GetOriginalChannelAccountIdOk

`func (o *PublicThread) GetOriginalChannelAccountIdOk() (*string, bool)`

GetOriginalChannelAccountIdOk returns a tuple with the OriginalChannelAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalChannelAccountId

`func (o *PublicThread) SetOriginalChannelAccountId(v string)`

SetOriginalChannelAccountId sets OriginalChannelAccountId field to given value.


### GetId

`func (o *PublicThread) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicThread) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicThread) SetId(v string)`

SetId sets Id field to given value.


### GetClosedAt

`func (o *PublicThread) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *PublicThread) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *PublicThread) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *PublicThread) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetSpam

`func (o *PublicThread) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *PublicThread) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *PublicThread) SetSpam(v bool)`

SetSpam sets Spam field to given value.


### GetInboxId

`func (o *PublicThread) GetInboxId() string`

GetInboxId returns the InboxId field if non-nil, zero value otherwise.

### GetInboxIdOk

`func (o *PublicThread) GetInboxIdOk() (*string, bool)`

GetInboxIdOk returns a tuple with the InboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboxId

`func (o *PublicThread) SetInboxId(v string)`

SetInboxId sets InboxId field to given value.


### GetStatus

`func (o *PublicThread) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicThread) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicThread) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLatestMessageReceivedTimestamp

`func (o *PublicThread) GetLatestMessageReceivedTimestamp() time.Time`

GetLatestMessageReceivedTimestamp returns the LatestMessageReceivedTimestamp field if non-nil, zero value otherwise.

### GetLatestMessageReceivedTimestampOk

`func (o *PublicThread) GetLatestMessageReceivedTimestampOk() (*time.Time, bool)`

GetLatestMessageReceivedTimestampOk returns a tuple with the LatestMessageReceivedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMessageReceivedTimestamp

`func (o *PublicThread) SetLatestMessageReceivedTimestamp(v time.Time)`

SetLatestMessageReceivedTimestamp sets LatestMessageReceivedTimestamp field to given value.

### HasLatestMessageReceivedTimestamp

`func (o *PublicThread) HasLatestMessageReceivedTimestamp() bool`

HasLatestMessageReceivedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


