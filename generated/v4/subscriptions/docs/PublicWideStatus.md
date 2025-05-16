# PublicWideStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WideStatusType** | **string** |  | 
**SubscriberIdString** | **string** |  | 
**Channel** | **string** |  | 
**BusinessUnitId** | Pointer to **int64** |  | [optional] 
**Status** | **string** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewPublicWideStatus

`func NewPublicWideStatus(wideStatusType string, subscriberIdString string, channel string, status string, timestamp time.Time, ) *PublicWideStatus`

NewPublicWideStatus instantiates a new PublicWideStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicWideStatusWithDefaults

`func NewPublicWideStatusWithDefaults() *PublicWideStatus`

NewPublicWideStatusWithDefaults instantiates a new PublicWideStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWideStatusType

`func (o *PublicWideStatus) GetWideStatusType() string`

GetWideStatusType returns the WideStatusType field if non-nil, zero value otherwise.

### GetWideStatusTypeOk

`func (o *PublicWideStatus) GetWideStatusTypeOk() (*string, bool)`

GetWideStatusTypeOk returns a tuple with the WideStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWideStatusType

`func (o *PublicWideStatus) SetWideStatusType(v string)`

SetWideStatusType sets WideStatusType field to given value.


### GetSubscriberIdString

`func (o *PublicWideStatus) GetSubscriberIdString() string`

GetSubscriberIdString returns the SubscriberIdString field if non-nil, zero value otherwise.

### GetSubscriberIdStringOk

`func (o *PublicWideStatus) GetSubscriberIdStringOk() (*string, bool)`

GetSubscriberIdStringOk returns a tuple with the SubscriberIdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdString

`func (o *PublicWideStatus) SetSubscriberIdString(v string)`

SetSubscriberIdString sets SubscriberIdString field to given value.


### GetChannel

`func (o *PublicWideStatus) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PublicWideStatus) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PublicWideStatus) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetBusinessUnitId

`func (o *PublicWideStatus) GetBusinessUnitId() int64`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *PublicWideStatus) GetBusinessUnitIdOk() (*int64, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *PublicWideStatus) SetBusinessUnitId(v int64)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *PublicWideStatus) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetStatus

`func (o *PublicWideStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicWideStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicWideStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *PublicWideStatus) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublicWideStatus) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublicWideStatus) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


