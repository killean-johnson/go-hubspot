# PublicInbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | When the inbox was created. | [optional] 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 
**Archived** | **bool** |  | 
**Name** | Pointer to **string** | The name of the inbox. | [optional] 
**Id** | Pointer to **string** | The ID of the inbox. | [optional] 
**Type** | **string** | Specifies whether this refers to a Conversations Inbox or to the Help Desk. Valid values are INBOX or HELP_DESK | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPublicInbox

`func NewPublicInbox(archived bool, type_ string, ) *PublicInbox`

NewPublicInbox instantiates a new PublicInbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicInboxWithDefaults

`func NewPublicInboxWithDefaults() *PublicInbox`

NewPublicInboxWithDefaults instantiates a new PublicInbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicInbox) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicInbox) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicInbox) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicInbox) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *PublicInbox) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *PublicInbox) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *PublicInbox) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *PublicInbox) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetArchived

`func (o *PublicInbox) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PublicInbox) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PublicInbox) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetName

`func (o *PublicInbox) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicInbox) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicInbox) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicInbox) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *PublicInbox) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicInbox) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicInbox) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicInbox) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PublicInbox) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicInbox) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicInbox) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *PublicInbox) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicInbox) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicInbox) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicInbox) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


