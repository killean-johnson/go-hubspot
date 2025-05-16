# TimeOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**OffsetDirection** | **string** |  | 
**TimeUnit** | **string** |  | 

## Methods

### NewTimeOffset

`func NewTimeOffset(amount int32, offsetDirection string, timeUnit string, ) *TimeOffset`

NewTimeOffset instantiates a new TimeOffset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffsetWithDefaults

`func NewTimeOffsetWithDefaults() *TimeOffset`

NewTimeOffsetWithDefaults instantiates a new TimeOffset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TimeOffset) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TimeOffset) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TimeOffset) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetOffsetDirection

`func (o *TimeOffset) GetOffsetDirection() string`

GetOffsetDirection returns the OffsetDirection field if non-nil, zero value otherwise.

### GetOffsetDirectionOk

`func (o *TimeOffset) GetOffsetDirectionOk() (*string, bool)`

GetOffsetDirectionOk returns a tuple with the OffsetDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetDirection

`func (o *TimeOffset) SetOffsetDirection(v string)`

SetOffsetDirection sets OffsetDirection field to given value.


### GetTimeUnit

`func (o *TimeOffset) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *TimeOffset) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *TimeOffset) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


