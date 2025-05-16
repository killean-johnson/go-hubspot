# PublicChannelAccountStagingTokenUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** |  | 
**DeliveryIdentifier** | [**PublicDeliveryIdentifier**](PublicDeliveryIdentifier.md) |  | 

## Methods

### NewPublicChannelAccountStagingTokenUpdateRequest

`func NewPublicChannelAccountStagingTokenUpdateRequest(accountName string, deliveryIdentifier PublicDeliveryIdentifier, ) *PublicChannelAccountStagingTokenUpdateRequest`

NewPublicChannelAccountStagingTokenUpdateRequest instantiates a new PublicChannelAccountStagingTokenUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicChannelAccountStagingTokenUpdateRequestWithDefaults

`func NewPublicChannelAccountStagingTokenUpdateRequestWithDefaults() *PublicChannelAccountStagingTokenUpdateRequest`

NewPublicChannelAccountStagingTokenUpdateRequestWithDefaults instantiates a new PublicChannelAccountStagingTokenUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *PublicChannelAccountStagingTokenUpdateRequest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PublicChannelAccountStagingTokenUpdateRequest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PublicChannelAccountStagingTokenUpdateRequest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetDeliveryIdentifier

`func (o *PublicChannelAccountStagingTokenUpdateRequest) GetDeliveryIdentifier() PublicDeliveryIdentifier`

GetDeliveryIdentifier returns the DeliveryIdentifier field if non-nil, zero value otherwise.

### GetDeliveryIdentifierOk

`func (o *PublicChannelAccountStagingTokenUpdateRequest) GetDeliveryIdentifierOk() (*PublicDeliveryIdentifier, bool)`

GetDeliveryIdentifierOk returns a tuple with the DeliveryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIdentifier

`func (o *PublicChannelAccountStagingTokenUpdateRequest) SetDeliveryIdentifier(v PublicDeliveryIdentifier)`

SetDeliveryIdentifier sets DeliveryIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


