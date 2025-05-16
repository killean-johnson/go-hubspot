# PublicWideStatusBulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WideStatuses** | [**[]PublicWideStatus**](PublicWideStatus.md) |  | 
**SubscriberIdString** | **string** |  | 

## Methods

### NewPublicWideStatusBulkResponse

`func NewPublicWideStatusBulkResponse(wideStatuses []PublicWideStatus, subscriberIdString string, ) *PublicWideStatusBulkResponse`

NewPublicWideStatusBulkResponse instantiates a new PublicWideStatusBulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicWideStatusBulkResponseWithDefaults

`func NewPublicWideStatusBulkResponseWithDefaults() *PublicWideStatusBulkResponse`

NewPublicWideStatusBulkResponseWithDefaults instantiates a new PublicWideStatusBulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWideStatuses

`func (o *PublicWideStatusBulkResponse) GetWideStatuses() []PublicWideStatus`

GetWideStatuses returns the WideStatuses field if non-nil, zero value otherwise.

### GetWideStatusesOk

`func (o *PublicWideStatusBulkResponse) GetWideStatusesOk() (*[]PublicWideStatus, bool)`

GetWideStatusesOk returns a tuple with the WideStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWideStatuses

`func (o *PublicWideStatusBulkResponse) SetWideStatuses(v []PublicWideStatus)`

SetWideStatuses sets WideStatuses field to given value.


### GetSubscriberIdString

`func (o *PublicWideStatusBulkResponse) GetSubscriberIdString() string`

GetSubscriberIdString returns the SubscriberIdString field if non-nil, zero value otherwise.

### GetSubscriberIdStringOk

`func (o *PublicWideStatusBulkResponse) GetSubscriberIdStringOk() (*string, bool)`

GetSubscriberIdStringOk returns a tuple with the SubscriberIdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdString

`func (o *PublicWideStatusBulkResponse) SetSubscriberIdString(v string)`

SetSubscriberIdString sets SubscriberIdString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


