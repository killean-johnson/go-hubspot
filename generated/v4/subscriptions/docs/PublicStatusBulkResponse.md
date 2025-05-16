# PublicStatusBulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberIdString** | **string** |  | 
**Statuses** | [**[]PublicStatus**](PublicStatus.md) |  | 

## Methods

### NewPublicStatusBulkResponse

`func NewPublicStatusBulkResponse(subscriberIdString string, statuses []PublicStatus, ) *PublicStatusBulkResponse`

NewPublicStatusBulkResponse instantiates a new PublicStatusBulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicStatusBulkResponseWithDefaults

`func NewPublicStatusBulkResponseWithDefaults() *PublicStatusBulkResponse`

NewPublicStatusBulkResponseWithDefaults instantiates a new PublicStatusBulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberIdString

`func (o *PublicStatusBulkResponse) GetSubscriberIdString() string`

GetSubscriberIdString returns the SubscriberIdString field if non-nil, zero value otherwise.

### GetSubscriberIdStringOk

`func (o *PublicStatusBulkResponse) GetSubscriberIdStringOk() (*string, bool)`

GetSubscriberIdStringOk returns a tuple with the SubscriberIdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdString

`func (o *PublicStatusBulkResponse) SetSubscriberIdString(v string)`

SetSubscriberIdString sets SubscriberIdString field to given value.


### GetStatuses

`func (o *PublicStatusBulkResponse) GetStatuses() []PublicStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *PublicStatusBulkResponse) GetStatusesOk() (*[]PublicStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *PublicStatusBulkResponse) SetStatuses(v []PublicStatus)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


