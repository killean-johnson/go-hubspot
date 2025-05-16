# PublicBulkOptOutFromAllResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberIdString** | **string** |  | 
**Statuses** | Pointer to [**[]PublicStatus**](PublicStatus.md) |  | [optional] 

## Methods

### NewPublicBulkOptOutFromAllResponse

`func NewPublicBulkOptOutFromAllResponse(subscriberIdString string, ) *PublicBulkOptOutFromAllResponse`

NewPublicBulkOptOutFromAllResponse instantiates a new PublicBulkOptOutFromAllResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicBulkOptOutFromAllResponseWithDefaults

`func NewPublicBulkOptOutFromAllResponseWithDefaults() *PublicBulkOptOutFromAllResponse`

NewPublicBulkOptOutFromAllResponseWithDefaults instantiates a new PublicBulkOptOutFromAllResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberIdString

`func (o *PublicBulkOptOutFromAllResponse) GetSubscriberIdString() string`

GetSubscriberIdString returns the SubscriberIdString field if non-nil, zero value otherwise.

### GetSubscriberIdStringOk

`func (o *PublicBulkOptOutFromAllResponse) GetSubscriberIdStringOk() (*string, bool)`

GetSubscriberIdStringOk returns a tuple with the SubscriberIdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdString

`func (o *PublicBulkOptOutFromAllResponse) SetSubscriberIdString(v string)`

SetSubscriberIdString sets SubscriberIdString field to given value.


### GetStatuses

`func (o *PublicBulkOptOutFromAllResponse) GetStatuses() []PublicStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *PublicBulkOptOutFromAllResponse) GetStatusesOk() (*[]PublicStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *PublicBulkOptOutFromAllResponse) SetStatuses(v []PublicStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *PublicBulkOptOutFromAllResponse) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


