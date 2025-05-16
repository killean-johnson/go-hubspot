# ExternalUserBusyTimes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeetingsUser** | [**ExternalMeetingsUser**](ExternalMeetingsUser.md) |  | 
**BusyTimes** | [**[]ExternalTimeRange**](ExternalTimeRange.md) |  | 
**IsOffline** | **bool** |  | 

## Methods

### NewExternalUserBusyTimes

`func NewExternalUserBusyTimes(meetingsUser ExternalMeetingsUser, busyTimes []ExternalTimeRange, isOffline bool, ) *ExternalUserBusyTimes`

NewExternalUserBusyTimes instantiates a new ExternalUserBusyTimes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalUserBusyTimesWithDefaults

`func NewExternalUserBusyTimesWithDefaults() *ExternalUserBusyTimes`

NewExternalUserBusyTimesWithDefaults instantiates a new ExternalUserBusyTimes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeetingsUser

`func (o *ExternalUserBusyTimes) GetMeetingsUser() ExternalMeetingsUser`

GetMeetingsUser returns the MeetingsUser field if non-nil, zero value otherwise.

### GetMeetingsUserOk

`func (o *ExternalUserBusyTimes) GetMeetingsUserOk() (*ExternalMeetingsUser, bool)`

GetMeetingsUserOk returns a tuple with the MeetingsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingsUser

`func (o *ExternalUserBusyTimes) SetMeetingsUser(v ExternalMeetingsUser)`

SetMeetingsUser sets MeetingsUser field to given value.


### GetBusyTimes

`func (o *ExternalUserBusyTimes) GetBusyTimes() []ExternalTimeRange`

GetBusyTimes returns the BusyTimes field if non-nil, zero value otherwise.

### GetBusyTimesOk

`func (o *ExternalUserBusyTimes) GetBusyTimesOk() (*[]ExternalTimeRange, bool)`

GetBusyTimesOk returns a tuple with the BusyTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyTimes

`func (o *ExternalUserBusyTimes) SetBusyTimes(v []ExternalTimeRange)`

SetBusyTimes sets BusyTimes field to given value.


### GetIsOffline

`func (o *ExternalUserBusyTimes) GetIsOffline() bool`

GetIsOffline returns the IsOffline field if non-nil, zero value otherwise.

### GetIsOfflineOk

`func (o *ExternalUserBusyTimes) GetIsOfflineOk() (*bool, bool)`

GetIsOfflineOk returns a tuple with the IsOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOffline

`func (o *ExternalUserBusyTimes) SetIsOffline(v bool)`

SetIsOffline sets IsOffline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


