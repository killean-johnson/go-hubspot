# PublicSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SenderField** | Pointer to **string** |  | [optional] 
**DeliveryIdentifier** | Pointer to [**PublicDeliveryIdentifier**](PublicDeliveryIdentifier.md) |  | [optional] 

## Methods

### NewPublicSender

`func NewPublicSender() *PublicSender`

NewPublicSender instantiates a new PublicSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSenderWithDefaults

`func NewPublicSenderWithDefaults() *PublicSender`

NewPublicSenderWithDefaults instantiates a new PublicSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorId

`func (o *PublicSender) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *PublicSender) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *PublicSender) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *PublicSender) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetName

`func (o *PublicSender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicSender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicSender) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicSender) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSenderField

`func (o *PublicSender) GetSenderField() string`

GetSenderField returns the SenderField field if non-nil, zero value otherwise.

### GetSenderFieldOk

`func (o *PublicSender) GetSenderFieldOk() (*string, bool)`

GetSenderFieldOk returns a tuple with the SenderField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderField

`func (o *PublicSender) SetSenderField(v string)`

SetSenderField sets SenderField field to given value.

### HasSenderField

`func (o *PublicSender) HasSenderField() bool`

HasSenderField returns a boolean if a field has been set.

### GetDeliveryIdentifier

`func (o *PublicSender) GetDeliveryIdentifier() PublicDeliveryIdentifier`

GetDeliveryIdentifier returns the DeliveryIdentifier field if non-nil, zero value otherwise.

### GetDeliveryIdentifierOk

`func (o *PublicSender) GetDeliveryIdentifierOk() (*PublicDeliveryIdentifier, bool)`

GetDeliveryIdentifierOk returns a tuple with the DeliveryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIdentifier

`func (o *PublicSender) SetDeliveryIdentifier(v PublicDeliveryIdentifier)`

SetDeliveryIdentifier sets DeliveryIdentifier field to given value.

### HasDeliveryIdentifier

`func (o *PublicSender) HasDeliveryIdentifier() bool`

HasDeliveryIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


