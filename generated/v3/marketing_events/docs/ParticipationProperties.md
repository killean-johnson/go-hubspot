# ParticipationProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OccurredAt** | **int64** |  | 
**AttendancePercentage** | Pointer to **string** |  | [optional] 
**AttendanceState** | **string** |  | 
**AttendanceDurationSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewParticipationProperties

`func NewParticipationProperties(occurredAt int64, attendanceState string, ) *ParticipationProperties`

NewParticipationProperties instantiates a new ParticipationProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipationPropertiesWithDefaults

`func NewParticipationPropertiesWithDefaults() *ParticipationProperties`

NewParticipationPropertiesWithDefaults instantiates a new ParticipationProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurredAt

`func (o *ParticipationProperties) GetOccurredAt() int64`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *ParticipationProperties) GetOccurredAtOk() (*int64, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *ParticipationProperties) SetOccurredAt(v int64)`

SetOccurredAt sets OccurredAt field to given value.


### GetAttendancePercentage

`func (o *ParticipationProperties) GetAttendancePercentage() string`

GetAttendancePercentage returns the AttendancePercentage field if non-nil, zero value otherwise.

### GetAttendancePercentageOk

`func (o *ParticipationProperties) GetAttendancePercentageOk() (*string, bool)`

GetAttendancePercentageOk returns a tuple with the AttendancePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendancePercentage

`func (o *ParticipationProperties) SetAttendancePercentage(v string)`

SetAttendancePercentage sets AttendancePercentage field to given value.

### HasAttendancePercentage

`func (o *ParticipationProperties) HasAttendancePercentage() bool`

HasAttendancePercentage returns a boolean if a field has been set.

### GetAttendanceState

`func (o *ParticipationProperties) GetAttendanceState() string`

GetAttendanceState returns the AttendanceState field if non-nil, zero value otherwise.

### GetAttendanceStateOk

`func (o *ParticipationProperties) GetAttendanceStateOk() (*string, bool)`

GetAttendanceStateOk returns a tuple with the AttendanceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendanceState

`func (o *ParticipationProperties) SetAttendanceState(v string)`

SetAttendanceState sets AttendanceState field to given value.


### GetAttendanceDurationSeconds

`func (o *ParticipationProperties) GetAttendanceDurationSeconds() int32`

GetAttendanceDurationSeconds returns the AttendanceDurationSeconds field if non-nil, zero value otherwise.

### GetAttendanceDurationSecondsOk

`func (o *ParticipationProperties) GetAttendanceDurationSecondsOk() (*int32, bool)`

GetAttendanceDurationSecondsOk returns a tuple with the AttendanceDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendanceDurationSeconds

`func (o *ParticipationProperties) SetAttendanceDurationSeconds(v int32)`

SetAttendanceDurationSeconds sets AttendanceDurationSeconds field to given value.

### HasAttendanceDurationSeconds

`func (o *ParticipationProperties) HasAttendanceDurationSeconds() bool`

HasAttendanceDurationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


