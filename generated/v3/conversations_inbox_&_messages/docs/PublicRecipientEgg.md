# PublicRecipientEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryIdentifiers** | [**[]PublicDeliveryIdentifier**](PublicDeliveryIdentifier.md) |  | 
**ActorId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DeliveryIdentifier** | Pointer to [**PublicDeliveryIdentifier**](PublicDeliveryIdentifier.md) |  | [optional] 
**RecipientField** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicRecipientEgg

`func NewPublicRecipientEgg(deliveryIdentifiers []PublicDeliveryIdentifier, ) *PublicRecipientEgg`

NewPublicRecipientEgg instantiates a new PublicRecipientEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicRecipientEggWithDefaults

`func NewPublicRecipientEggWithDefaults() *PublicRecipientEgg`

NewPublicRecipientEggWithDefaults instantiates a new PublicRecipientEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryIdentifiers

`func (o *PublicRecipientEgg) GetDeliveryIdentifiers() []PublicDeliveryIdentifier`

GetDeliveryIdentifiers returns the DeliveryIdentifiers field if non-nil, zero value otherwise.

### GetDeliveryIdentifiersOk

`func (o *PublicRecipientEgg) GetDeliveryIdentifiersOk() (*[]PublicDeliveryIdentifier, bool)`

GetDeliveryIdentifiersOk returns a tuple with the DeliveryIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIdentifiers

`func (o *PublicRecipientEgg) SetDeliveryIdentifiers(v []PublicDeliveryIdentifier)`

SetDeliveryIdentifiers sets DeliveryIdentifiers field to given value.


### GetActorId

`func (o *PublicRecipientEgg) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *PublicRecipientEgg) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *PublicRecipientEgg) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *PublicRecipientEgg) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetName

`func (o *PublicRecipientEgg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicRecipientEgg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicRecipientEgg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicRecipientEgg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeliveryIdentifier

`func (o *PublicRecipientEgg) GetDeliveryIdentifier() PublicDeliveryIdentifier`

GetDeliveryIdentifier returns the DeliveryIdentifier field if non-nil, zero value otherwise.

### GetDeliveryIdentifierOk

`func (o *PublicRecipientEgg) GetDeliveryIdentifierOk() (*PublicDeliveryIdentifier, bool)`

GetDeliveryIdentifierOk returns a tuple with the DeliveryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIdentifier

`func (o *PublicRecipientEgg) SetDeliveryIdentifier(v PublicDeliveryIdentifier)`

SetDeliveryIdentifier sets DeliveryIdentifier field to given value.

### HasDeliveryIdentifier

`func (o *PublicRecipientEgg) HasDeliveryIdentifier() bool`

HasDeliveryIdentifier returns a boolean if a field has been set.

### GetRecipientField

`func (o *PublicRecipientEgg) GetRecipientField() string`

GetRecipientField returns the RecipientField field if non-nil, zero value otherwise.

### GetRecipientFieldOk

`func (o *PublicRecipientEgg) GetRecipientFieldOk() (*string, bool)`

GetRecipientFieldOk returns a tuple with the RecipientField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientField

`func (o *PublicRecipientEgg) SetRecipientField(v string)`

SetRecipientField sets RecipientField field to given value.

### HasRecipientField

`func (o *PublicRecipientEgg) HasRecipientField() bool`

HasRecipientField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


