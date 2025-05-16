# PublicChannelAccountStagingToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**AccountName** | Pointer to **string** |  | [optional] 
**GenericChannelId** | **int32** |  | 
**DeliveryIdentifier** | Pointer to [**PublicDeliveryIdentifier**](PublicDeliveryIdentifier.md) |  | [optional] 
**AccountToken** | **string** |  | 
**UserId** | **int32** |  | 
**InboxId** | **int32** |  | 

## Methods

### NewPublicChannelAccountStagingToken

`func NewPublicChannelAccountStagingToken(createdAt time.Time, genericChannelId int32, accountToken string, userId int32, inboxId int32, ) *PublicChannelAccountStagingToken`

NewPublicChannelAccountStagingToken instantiates a new PublicChannelAccountStagingToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicChannelAccountStagingTokenWithDefaults

`func NewPublicChannelAccountStagingTokenWithDefaults() *PublicChannelAccountStagingToken`

NewPublicChannelAccountStagingTokenWithDefaults instantiates a new PublicChannelAccountStagingToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicChannelAccountStagingToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicChannelAccountStagingToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicChannelAccountStagingToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAccountName

`func (o *PublicChannelAccountStagingToken) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PublicChannelAccountStagingToken) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PublicChannelAccountStagingToken) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PublicChannelAccountStagingToken) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetGenericChannelId

`func (o *PublicChannelAccountStagingToken) GetGenericChannelId() int32`

GetGenericChannelId returns the GenericChannelId field if non-nil, zero value otherwise.

### GetGenericChannelIdOk

`func (o *PublicChannelAccountStagingToken) GetGenericChannelIdOk() (*int32, bool)`

GetGenericChannelIdOk returns a tuple with the GenericChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericChannelId

`func (o *PublicChannelAccountStagingToken) SetGenericChannelId(v int32)`

SetGenericChannelId sets GenericChannelId field to given value.


### GetDeliveryIdentifier

`func (o *PublicChannelAccountStagingToken) GetDeliveryIdentifier() PublicDeliveryIdentifier`

GetDeliveryIdentifier returns the DeliveryIdentifier field if non-nil, zero value otherwise.

### GetDeliveryIdentifierOk

`func (o *PublicChannelAccountStagingToken) GetDeliveryIdentifierOk() (*PublicDeliveryIdentifier, bool)`

GetDeliveryIdentifierOk returns a tuple with the DeliveryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIdentifier

`func (o *PublicChannelAccountStagingToken) SetDeliveryIdentifier(v PublicDeliveryIdentifier)`

SetDeliveryIdentifier sets DeliveryIdentifier field to given value.

### HasDeliveryIdentifier

`func (o *PublicChannelAccountStagingToken) HasDeliveryIdentifier() bool`

HasDeliveryIdentifier returns a boolean if a field has been set.

### GetAccountToken

`func (o *PublicChannelAccountStagingToken) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *PublicChannelAccountStagingToken) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *PublicChannelAccountStagingToken) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetUserId

`func (o *PublicChannelAccountStagingToken) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicChannelAccountStagingToken) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicChannelAccountStagingToken) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetInboxId

`func (o *PublicChannelAccountStagingToken) GetInboxId() int32`

GetInboxId returns the InboxId field if non-nil, zero value otherwise.

### GetInboxIdOk

`func (o *PublicChannelAccountStagingToken) GetInboxIdOk() (*int32, bool)`

GetInboxIdOk returns a tuple with the InboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboxId

`func (o *PublicChannelAccountStagingToken) SetInboxId(v int32)`

SetInboxId sets InboxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


