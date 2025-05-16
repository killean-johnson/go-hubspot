# DatePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **int32** |  | 
**Hour** | Pointer to **int32** |  | [optional] 
**Year** | **int32** |  | 
**TimezoneSource** | **string** |  | 
**Millisecond** | Pointer to **int32** |  | [optional] 
**TimeType** | **string** |  | [default to "DATE"]
**ZoneId** | **string** |  | 
**Day** | **int32** |  | 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 

## Methods

### NewDatePoint

`func NewDatePoint(month int32, year int32, timezoneSource string, timeType string, zoneId string, day int32, ) *DatePoint`

NewDatePoint instantiates a new DatePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatePointWithDefaults

`func NewDatePointWithDefaults() *DatePoint`

NewDatePointWithDefaults instantiates a new DatePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *DatePoint) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *DatePoint) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *DatePoint) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetHour

`func (o *DatePoint) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *DatePoint) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *DatePoint) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *DatePoint) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetYear

`func (o *DatePoint) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *DatePoint) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *DatePoint) SetYear(v int32)`

SetYear sets Year field to given value.


### GetTimezoneSource

`func (o *DatePoint) GetTimezoneSource() string`

GetTimezoneSource returns the TimezoneSource field if non-nil, zero value otherwise.

### GetTimezoneSourceOk

`func (o *DatePoint) GetTimezoneSourceOk() (*string, bool)`

GetTimezoneSourceOk returns a tuple with the TimezoneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneSource

`func (o *DatePoint) SetTimezoneSource(v string)`

SetTimezoneSource sets TimezoneSource field to given value.


### GetMillisecond

`func (o *DatePoint) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *DatePoint) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *DatePoint) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *DatePoint) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetTimeType

`func (o *DatePoint) GetTimeType() string`

GetTimeType returns the TimeType field if non-nil, zero value otherwise.

### GetTimeTypeOk

`func (o *DatePoint) GetTimeTypeOk() (*string, bool)`

GetTimeTypeOk returns a tuple with the TimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeType

`func (o *DatePoint) SetTimeType(v string)`

SetTimeType sets TimeType field to given value.


### GetZoneId

`func (o *DatePoint) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DatePoint) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DatePoint) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetDay

`func (o *DatePoint) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DatePoint) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DatePoint) SetDay(v int32)`

SetDay sets Day field to given value.


### GetMinute

`func (o *DatePoint) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *DatePoint) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *DatePoint) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *DatePoint) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *DatePoint) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *DatePoint) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *DatePoint) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *DatePoint) HasSecond() bool`

HasSecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


