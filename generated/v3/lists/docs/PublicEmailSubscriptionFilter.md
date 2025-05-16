# PublicEmailSubscriptionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionType** | Pointer to **string** |  | [optional] 
**SubscriptionIds** | **[]string** |  | 
**FilterType** | **string** |  | [default to "EMAIL_SUBSCRIPTION"]
**AcceptedStatuses** | **[]string** |  | 

## Methods

### NewPublicEmailSubscriptionFilter

`func NewPublicEmailSubscriptionFilter(subscriptionIds []string, filterType string, acceptedStatuses []string, ) *PublicEmailSubscriptionFilter`

NewPublicEmailSubscriptionFilter instantiates a new PublicEmailSubscriptionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEmailSubscriptionFilterWithDefaults

`func NewPublicEmailSubscriptionFilterWithDefaults() *PublicEmailSubscriptionFilter`

NewPublicEmailSubscriptionFilterWithDefaults instantiates a new PublicEmailSubscriptionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionType

`func (o *PublicEmailSubscriptionFilter) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *PublicEmailSubscriptionFilter) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *PublicEmailSubscriptionFilter) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *PublicEmailSubscriptionFilter) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *PublicEmailSubscriptionFilter) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *PublicEmailSubscriptionFilter) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *PublicEmailSubscriptionFilter) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.


### GetFilterType

`func (o *PublicEmailSubscriptionFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicEmailSubscriptionFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicEmailSubscriptionFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetAcceptedStatuses

`func (o *PublicEmailSubscriptionFilter) GetAcceptedStatuses() []string`

GetAcceptedStatuses returns the AcceptedStatuses field if non-nil, zero value otherwise.

### GetAcceptedStatusesOk

`func (o *PublicEmailSubscriptionFilter) GetAcceptedStatusesOk() (*[]string, bool)`

GetAcceptedStatusesOk returns a tuple with the AcceptedStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedStatuses

`func (o *PublicEmailSubscriptionFilter) SetAcceptedStatuses(v []string)`

SetAcceptedStatuses sets AcceptedStatuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


