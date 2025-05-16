# ApiTimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | [**ApiTimeOfDay**](ApiTimeOfDay.md) |  | 
**EndTime** | [**ApiTimeOfDay**](ApiTimeOfDay.md) |  | 
**Day** | **string** |  | 

## Methods

### NewApiTimeWindow

`func NewApiTimeWindow(startTime ApiTimeOfDay, endTime ApiTimeOfDay, day string, ) *ApiTimeWindow`

NewApiTimeWindow instantiates a new ApiTimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTimeWindowWithDefaults

`func NewApiTimeWindowWithDefaults() *ApiTimeWindow`

NewApiTimeWindowWithDefaults instantiates a new ApiTimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *ApiTimeWindow) GetStartTime() ApiTimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApiTimeWindow) GetStartTimeOk() (*ApiTimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApiTimeWindow) SetStartTime(v ApiTimeOfDay)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ApiTimeWindow) GetEndTime() ApiTimeOfDay`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApiTimeWindow) GetEndTimeOk() (*ApiTimeOfDay, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApiTimeWindow) SetEndTime(v ApiTimeOfDay)`

SetEndTime sets EndTime field to given value.


### GetDay

`func (o *ApiTimeWindow) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *ApiTimeWindow) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *ApiTimeWindow) SetDay(v string)`

SetDay sets Day field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


