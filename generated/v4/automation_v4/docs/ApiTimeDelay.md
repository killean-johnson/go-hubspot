# ApiTimeDelay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeZoneStrategy** | Pointer to [**ApiTimeDelayTimeZoneStrategy**](ApiTimeDelayTimeZoneStrategy.md) |  | [optional] 
**Delta** | **int32** |  | 
**DaysOfWeek** | **[]string** |  | 
**TimeUnit** | **string** |  | 
**TimeOfDay** | Pointer to [**ApiTimeOfDay**](ApiTimeOfDay.md) |  | [optional] 

## Methods

### NewApiTimeDelay

`func NewApiTimeDelay(delta int32, daysOfWeek []string, timeUnit string, ) *ApiTimeDelay`

NewApiTimeDelay instantiates a new ApiTimeDelay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTimeDelayWithDefaults

`func NewApiTimeDelayWithDefaults() *ApiTimeDelay`

NewApiTimeDelayWithDefaults instantiates a new ApiTimeDelay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeZoneStrategy

`func (o *ApiTimeDelay) GetTimeZoneStrategy() ApiTimeDelayTimeZoneStrategy`

GetTimeZoneStrategy returns the TimeZoneStrategy field if non-nil, zero value otherwise.

### GetTimeZoneStrategyOk

`func (o *ApiTimeDelay) GetTimeZoneStrategyOk() (*ApiTimeDelayTimeZoneStrategy, bool)`

GetTimeZoneStrategyOk returns a tuple with the TimeZoneStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneStrategy

`func (o *ApiTimeDelay) SetTimeZoneStrategy(v ApiTimeDelayTimeZoneStrategy)`

SetTimeZoneStrategy sets TimeZoneStrategy field to given value.

### HasTimeZoneStrategy

`func (o *ApiTimeDelay) HasTimeZoneStrategy() bool`

HasTimeZoneStrategy returns a boolean if a field has been set.

### GetDelta

`func (o *ApiTimeDelay) GetDelta() int32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *ApiTimeDelay) GetDeltaOk() (*int32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *ApiTimeDelay) SetDelta(v int32)`

SetDelta sets Delta field to given value.


### GetDaysOfWeek

`func (o *ApiTimeDelay) GetDaysOfWeek() []string`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *ApiTimeDelay) GetDaysOfWeekOk() (*[]string, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *ApiTimeDelay) SetDaysOfWeek(v []string)`

SetDaysOfWeek sets DaysOfWeek field to given value.


### GetTimeUnit

`func (o *ApiTimeDelay) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *ApiTimeDelay) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *ApiTimeDelay) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.


### GetTimeOfDay

`func (o *ApiTimeDelay) GetTimeOfDay() ApiTimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *ApiTimeDelay) GetTimeOfDayOk() (*ApiTimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *ApiTimeDelay) SetTimeOfDay(v ApiTimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *ApiTimeDelay) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


