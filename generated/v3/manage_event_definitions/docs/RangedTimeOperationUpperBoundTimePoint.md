# RangedTimeOperationUpperBoundTimePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **int32** |  | 
**Hour** | Pointer to **int32** |  | [optional] 
**Year** | **int32** |  | 
**TimezoneSource** | **string** |  | 
**Millisecond** | Pointer to **int32** |  | [optional] 
**TimeType** | **string** |  | [default to "PROPERTY_REFERENCED"]
**ZoneId** | **string** |  | 
**Day** | **int32** |  | 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to [**IndexOffset**](IndexOffset.md) |  | [optional] 
**IndexReference** | [**IndexedTimePointIndexReference**](IndexedTimePointIndexReference.md) |  | 
**Property** | **string** |  | 
**ReferenceType** | **string** |  | 

## Methods

### NewRangedTimeOperationUpperBoundTimePoint

`func NewRangedTimeOperationUpperBoundTimePoint(month int32, year int32, timezoneSource string, timeType string, zoneId string, day int32, indexReference IndexedTimePointIndexReference, property string, referenceType string, ) *RangedTimeOperationUpperBoundTimePoint`

NewRangedTimeOperationUpperBoundTimePoint instantiates a new RangedTimeOperationUpperBoundTimePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangedTimeOperationUpperBoundTimePointWithDefaults

`func NewRangedTimeOperationUpperBoundTimePointWithDefaults() *RangedTimeOperationUpperBoundTimePoint`

NewRangedTimeOperationUpperBoundTimePointWithDefaults instantiates a new RangedTimeOperationUpperBoundTimePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *RangedTimeOperationUpperBoundTimePoint) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *RangedTimeOperationUpperBoundTimePoint) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetHour

`func (o *RangedTimeOperationUpperBoundTimePoint) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *RangedTimeOperationUpperBoundTimePoint) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *RangedTimeOperationUpperBoundTimePoint) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetYear

`func (o *RangedTimeOperationUpperBoundTimePoint) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *RangedTimeOperationUpperBoundTimePoint) SetYear(v int32)`

SetYear sets Year field to given value.


### GetTimezoneSource

`func (o *RangedTimeOperationUpperBoundTimePoint) GetTimezoneSource() string`

GetTimezoneSource returns the TimezoneSource field if non-nil, zero value otherwise.

### GetTimezoneSourceOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetTimezoneSourceOk() (*string, bool)`

GetTimezoneSourceOk returns a tuple with the TimezoneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneSource

`func (o *RangedTimeOperationUpperBoundTimePoint) SetTimezoneSource(v string)`

SetTimezoneSource sets TimezoneSource field to given value.


### GetMillisecond

`func (o *RangedTimeOperationUpperBoundTimePoint) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *RangedTimeOperationUpperBoundTimePoint) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *RangedTimeOperationUpperBoundTimePoint) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetTimeType

`func (o *RangedTimeOperationUpperBoundTimePoint) GetTimeType() string`

GetTimeType returns the TimeType field if non-nil, zero value otherwise.

### GetTimeTypeOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetTimeTypeOk() (*string, bool)`

GetTimeTypeOk returns a tuple with the TimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeType

`func (o *RangedTimeOperationUpperBoundTimePoint) SetTimeType(v string)`

SetTimeType sets TimeType field to given value.


### GetZoneId

`func (o *RangedTimeOperationUpperBoundTimePoint) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *RangedTimeOperationUpperBoundTimePoint) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetDay

`func (o *RangedTimeOperationUpperBoundTimePoint) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *RangedTimeOperationUpperBoundTimePoint) SetDay(v int32)`

SetDay sets Day field to given value.


### GetMinute

`func (o *RangedTimeOperationUpperBoundTimePoint) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *RangedTimeOperationUpperBoundTimePoint) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *RangedTimeOperationUpperBoundTimePoint) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *RangedTimeOperationUpperBoundTimePoint) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *RangedTimeOperationUpperBoundTimePoint) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *RangedTimeOperationUpperBoundTimePoint) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### GetOffset

`func (o *RangedTimeOperationUpperBoundTimePoint) GetOffset() IndexOffset`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetOffsetOk() (*IndexOffset, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RangedTimeOperationUpperBoundTimePoint) SetOffset(v IndexOffset)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RangedTimeOperationUpperBoundTimePoint) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetIndexReference

`func (o *RangedTimeOperationUpperBoundTimePoint) GetIndexReference() IndexedTimePointIndexReference`

GetIndexReference returns the IndexReference field if non-nil, zero value otherwise.

### GetIndexReferenceOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetIndexReferenceOk() (*IndexedTimePointIndexReference, bool)`

GetIndexReferenceOk returns a tuple with the IndexReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexReference

`func (o *RangedTimeOperationUpperBoundTimePoint) SetIndexReference(v IndexedTimePointIndexReference)`

SetIndexReference sets IndexReference field to given value.


### GetProperty

`func (o *RangedTimeOperationUpperBoundTimePoint) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RangedTimeOperationUpperBoundTimePoint) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetReferenceType

`func (o *RangedTimeOperationUpperBoundTimePoint) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *RangedTimeOperationUpperBoundTimePoint) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *RangedTimeOperationUpperBoundTimePoint) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


