# ChannelIntegrationParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DeliveryIdentifier** | [**PublicDeliveryIdentifier**](PublicDeliveryIdentifier.md) |  | 

## Methods

### NewChannelIntegrationParticipant

`func NewChannelIntegrationParticipant(deliveryIdentifier PublicDeliveryIdentifier, ) *ChannelIntegrationParticipant`

NewChannelIntegrationParticipant instantiates a new ChannelIntegrationParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelIntegrationParticipantWithDefaults

`func NewChannelIntegrationParticipantWithDefaults() *ChannelIntegrationParticipant`

NewChannelIntegrationParticipantWithDefaults instantiates a new ChannelIntegrationParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChannelIntegrationParticipant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelIntegrationParticipant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelIntegrationParticipant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelIntegrationParticipant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeliveryIdentifier

`func (o *ChannelIntegrationParticipant) GetDeliveryIdentifier() PublicDeliveryIdentifier`

GetDeliveryIdentifier returns the DeliveryIdentifier field if non-nil, zero value otherwise.

### GetDeliveryIdentifierOk

`func (o *ChannelIntegrationParticipant) GetDeliveryIdentifierOk() (*PublicDeliveryIdentifier, bool)`

GetDeliveryIdentifierOk returns a tuple with the DeliveryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIdentifier

`func (o *ChannelIntegrationParticipant) SetDeliveryIdentifier(v PublicDeliveryIdentifier)`

SetDeliveryIdentifier sets DeliveryIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


