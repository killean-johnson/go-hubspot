# PublicChannelAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 
**Archived** | **bool** |  | 
**Authorized** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | The name of the channel account. | [optional] 
**Active** | Pointer to **bool** | Whether the channel account is turned on. | [optional] 
**DeliveryIdentifier** | Pointer to [**PublicDeliveryIdentifier**](PublicDeliveryIdentifier.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the channel account. | [optional] 
**InboxId** | Pointer to **string** | The ID of the conversations inbox that contains the channel account. | [optional] 
**ChannelId** | Pointer to **string** | The ID of the channel that the channel account is an instance of. | [optional] 

## Methods

### NewPublicChannelAccount

`func NewPublicChannelAccount(archived bool, ) *PublicChannelAccount`

NewPublicChannelAccount instantiates a new PublicChannelAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicChannelAccountWithDefaults

`func NewPublicChannelAccountWithDefaults() *PublicChannelAccount`

NewPublicChannelAccountWithDefaults instantiates a new PublicChannelAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicChannelAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicChannelAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicChannelAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicChannelAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *PublicChannelAccount) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *PublicChannelAccount) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *PublicChannelAccount) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *PublicChannelAccount) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetArchived

`func (o *PublicChannelAccount) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PublicChannelAccount) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PublicChannelAccount) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetAuthorized

`func (o *PublicChannelAccount) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *PublicChannelAccount) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *PublicChannelAccount) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *PublicChannelAccount) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### GetName

`func (o *PublicChannelAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicChannelAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicChannelAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicChannelAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PublicChannelAccount) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PublicChannelAccount) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PublicChannelAccount) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PublicChannelAccount) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeliveryIdentifier

`func (o *PublicChannelAccount) GetDeliveryIdentifier() PublicDeliveryIdentifier`

GetDeliveryIdentifier returns the DeliveryIdentifier field if non-nil, zero value otherwise.

### GetDeliveryIdentifierOk

`func (o *PublicChannelAccount) GetDeliveryIdentifierOk() (*PublicDeliveryIdentifier, bool)`

GetDeliveryIdentifierOk returns a tuple with the DeliveryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIdentifier

`func (o *PublicChannelAccount) SetDeliveryIdentifier(v PublicDeliveryIdentifier)`

SetDeliveryIdentifier sets DeliveryIdentifier field to given value.

### HasDeliveryIdentifier

`func (o *PublicChannelAccount) HasDeliveryIdentifier() bool`

HasDeliveryIdentifier returns a boolean if a field has been set.

### GetId

`func (o *PublicChannelAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicChannelAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicChannelAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicChannelAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInboxId

`func (o *PublicChannelAccount) GetInboxId() string`

GetInboxId returns the InboxId field if non-nil, zero value otherwise.

### GetInboxIdOk

`func (o *PublicChannelAccount) GetInboxIdOk() (*string, bool)`

GetInboxIdOk returns a tuple with the InboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboxId

`func (o *PublicChannelAccount) SetInboxId(v string)`

SetInboxId sets InboxId field to given value.

### HasInboxId

`func (o *PublicChannelAccount) HasInboxId() bool`

HasInboxId returns a boolean if a field has been set.

### GetChannelId

`func (o *PublicChannelAccount) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PublicChannelAccount) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PublicChannelAccount) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *PublicChannelAccount) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


