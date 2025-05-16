# PublicChannelAccountEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | **bool** |  | 
**Name** | **string** |  | 
**DeliveryIdentifier** | Pointer to [**PublicDeliveryIdentifier**](PublicDeliveryIdentifier.md) |  | [optional] 
**InboxId** | **string** |  | 

## Methods

### NewPublicChannelAccountEgg

`func NewPublicChannelAccountEgg(authorized bool, name string, inboxId string, ) *PublicChannelAccountEgg`

NewPublicChannelAccountEgg instantiates a new PublicChannelAccountEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicChannelAccountEggWithDefaults

`func NewPublicChannelAccountEggWithDefaults() *PublicChannelAccountEgg`

NewPublicChannelAccountEggWithDefaults instantiates a new PublicChannelAccountEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *PublicChannelAccountEgg) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *PublicChannelAccountEgg) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *PublicChannelAccountEgg) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.


### GetName

`func (o *PublicChannelAccountEgg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicChannelAccountEgg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicChannelAccountEgg) SetName(v string)`

SetName sets Name field to given value.


### GetDeliveryIdentifier

`func (o *PublicChannelAccountEgg) GetDeliveryIdentifier() PublicDeliveryIdentifier`

GetDeliveryIdentifier returns the DeliveryIdentifier field if non-nil, zero value otherwise.

### GetDeliveryIdentifierOk

`func (o *PublicChannelAccountEgg) GetDeliveryIdentifierOk() (*PublicDeliveryIdentifier, bool)`

GetDeliveryIdentifierOk returns a tuple with the DeliveryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIdentifier

`func (o *PublicChannelAccountEgg) SetDeliveryIdentifier(v PublicDeliveryIdentifier)`

SetDeliveryIdentifier sets DeliveryIdentifier field to given value.

### HasDeliveryIdentifier

`func (o *PublicChannelAccountEgg) HasDeliveryIdentifier() bool`

HasDeliveryIdentifier returns a boolean if a field has been set.

### GetInboxId

`func (o *PublicChannelAccountEgg) GetInboxId() string`

GetInboxId returns the InboxId field if non-nil, zero value otherwise.

### GetInboxIdOk

`func (o *PublicChannelAccountEgg) GetInboxIdOk() (*string, bool)`

GetInboxIdOk returns a tuple with the InboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboxId

`func (o *PublicChannelAccountEgg) SetInboxId(v string)`

SetInboxId sets InboxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


