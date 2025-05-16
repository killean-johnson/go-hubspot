# PublicCommunicationSubscriptionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionType** | **string** |  | 
**SubscriptionIds** | **[]string** |  | 
**Channel** | **string** |  | 
**AcceptedOptStates** | **[]string** |  | 
**FilterType** | **string** |  | [default to "COMMUNICATION_SUBSCRIPTION"]
**BusinessUnitId** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicCommunicationSubscriptionFilter

`func NewPublicCommunicationSubscriptionFilter(subscriptionType string, subscriptionIds []string, channel string, acceptedOptStates []string, filterType string, ) *PublicCommunicationSubscriptionFilter`

NewPublicCommunicationSubscriptionFilter instantiates a new PublicCommunicationSubscriptionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCommunicationSubscriptionFilterWithDefaults

`func NewPublicCommunicationSubscriptionFilterWithDefaults() *PublicCommunicationSubscriptionFilter`

NewPublicCommunicationSubscriptionFilterWithDefaults instantiates a new PublicCommunicationSubscriptionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionType

`func (o *PublicCommunicationSubscriptionFilter) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *PublicCommunicationSubscriptionFilter) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *PublicCommunicationSubscriptionFilter) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.


### GetSubscriptionIds

`func (o *PublicCommunicationSubscriptionFilter) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *PublicCommunicationSubscriptionFilter) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *PublicCommunicationSubscriptionFilter) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.


### GetChannel

`func (o *PublicCommunicationSubscriptionFilter) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PublicCommunicationSubscriptionFilter) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PublicCommunicationSubscriptionFilter) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetAcceptedOptStates

`func (o *PublicCommunicationSubscriptionFilter) GetAcceptedOptStates() []string`

GetAcceptedOptStates returns the AcceptedOptStates field if non-nil, zero value otherwise.

### GetAcceptedOptStatesOk

`func (o *PublicCommunicationSubscriptionFilter) GetAcceptedOptStatesOk() (*[]string, bool)`

GetAcceptedOptStatesOk returns a tuple with the AcceptedOptStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedOptStates

`func (o *PublicCommunicationSubscriptionFilter) SetAcceptedOptStates(v []string)`

SetAcceptedOptStates sets AcceptedOptStates field to given value.


### GetFilterType

`func (o *PublicCommunicationSubscriptionFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicCommunicationSubscriptionFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicCommunicationSubscriptionFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetBusinessUnitId

`func (o *PublicCommunicationSubscriptionFilter) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *PublicCommunicationSubscriptionFilter) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *PublicCommunicationSubscriptionFilter) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *PublicCommunicationSubscriptionFilter) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


