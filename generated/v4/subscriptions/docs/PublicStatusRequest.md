# PublicStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusState** | **string** | The status of the contact&#39;s subscription. | 
**Channel** | **string** | The type of communication channel. Currently, only &#x60;EMAIL&#x60; is supported. | 
**SubscriberIdString** | **string** | The contact&#39;s email address. | 
**LegalBasis** | Pointer to **string** | The legal basis for communication. | [optional] 
**SubscriptionId** | **int32** | The ID of the subscription to update. | 
**LegalBasisExplanation** | Pointer to **string** | The explanation for the legal basis. | [optional] 

## Methods

### NewPublicStatusRequest

`func NewPublicStatusRequest(statusState string, channel string, subscriberIdString string, subscriptionId int32, ) *PublicStatusRequest`

NewPublicStatusRequest instantiates a new PublicStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicStatusRequestWithDefaults

`func NewPublicStatusRequestWithDefaults() *PublicStatusRequest`

NewPublicStatusRequestWithDefaults instantiates a new PublicStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusState

`func (o *PublicStatusRequest) GetStatusState() string`

GetStatusState returns the StatusState field if non-nil, zero value otherwise.

### GetStatusStateOk

`func (o *PublicStatusRequest) GetStatusStateOk() (*string, bool)`

GetStatusStateOk returns a tuple with the StatusState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusState

`func (o *PublicStatusRequest) SetStatusState(v string)`

SetStatusState sets StatusState field to given value.


### GetChannel

`func (o *PublicStatusRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PublicStatusRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PublicStatusRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetSubscriberIdString

`func (o *PublicStatusRequest) GetSubscriberIdString() string`

GetSubscriberIdString returns the SubscriberIdString field if non-nil, zero value otherwise.

### GetSubscriberIdStringOk

`func (o *PublicStatusRequest) GetSubscriberIdStringOk() (*string, bool)`

GetSubscriberIdStringOk returns a tuple with the SubscriberIdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdString

`func (o *PublicStatusRequest) SetSubscriberIdString(v string)`

SetSubscriberIdString sets SubscriberIdString field to given value.


### GetLegalBasis

`func (o *PublicStatusRequest) GetLegalBasis() string`

GetLegalBasis returns the LegalBasis field if non-nil, zero value otherwise.

### GetLegalBasisOk

`func (o *PublicStatusRequest) GetLegalBasisOk() (*string, bool)`

GetLegalBasisOk returns a tuple with the LegalBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasis

`func (o *PublicStatusRequest) SetLegalBasis(v string)`

SetLegalBasis sets LegalBasis field to given value.

### HasLegalBasis

`func (o *PublicStatusRequest) HasLegalBasis() bool`

HasLegalBasis returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *PublicStatusRequest) GetSubscriptionId() int32`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PublicStatusRequest) GetSubscriptionIdOk() (*int32, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PublicStatusRequest) SetSubscriptionId(v int32)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetLegalBasisExplanation

`func (o *PublicStatusRequest) GetLegalBasisExplanation() string`

GetLegalBasisExplanation returns the LegalBasisExplanation field if non-nil, zero value otherwise.

### GetLegalBasisExplanationOk

`func (o *PublicStatusRequest) GetLegalBasisExplanationOk() (*string, bool)`

GetLegalBasisExplanationOk returns a tuple with the LegalBasisExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasisExplanation

`func (o *PublicStatusRequest) SetLegalBasisExplanation(v string)`

SetLegalBasisExplanation sets LegalBasisExplanation field to given value.

### HasLegalBasisExplanation

`func (o *PublicStatusRequest) HasLegalBasisExplanation() bool`

HasLegalBasisExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


