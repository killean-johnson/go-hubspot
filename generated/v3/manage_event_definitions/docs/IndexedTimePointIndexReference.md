# IndexedTimePointIndexReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to **int32** |  | [optional] 
**Millisecond** | Pointer to **int32** |  | [optional] 
**ReferenceType** | **string** |  | [default to "FISCAL_YEAR"]
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 
**DayOfWeek** | **string** |  | 
**Day** | **int32** |  | 
**Month** | **int32** |  | 

## Methods

### NewIndexedTimePointIndexReference

`func NewIndexedTimePointIndexReference(referenceType string, dayOfWeek string, day int32, month int32, ) *IndexedTimePointIndexReference`

NewIndexedTimePointIndexReference instantiates a new IndexedTimePointIndexReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexedTimePointIndexReferenceWithDefaults

`func NewIndexedTimePointIndexReferenceWithDefaults() *IndexedTimePointIndexReference`

NewIndexedTimePointIndexReferenceWithDefaults instantiates a new IndexedTimePointIndexReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *IndexedTimePointIndexReference) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *IndexedTimePointIndexReference) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *IndexedTimePointIndexReference) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *IndexedTimePointIndexReference) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMillisecond

`func (o *IndexedTimePointIndexReference) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *IndexedTimePointIndexReference) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *IndexedTimePointIndexReference) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *IndexedTimePointIndexReference) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetReferenceType

`func (o *IndexedTimePointIndexReference) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *IndexedTimePointIndexReference) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *IndexedTimePointIndexReference) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.


### GetMinute

`func (o *IndexedTimePointIndexReference) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *IndexedTimePointIndexReference) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *IndexedTimePointIndexReference) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *IndexedTimePointIndexReference) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *IndexedTimePointIndexReference) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *IndexedTimePointIndexReference) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *IndexedTimePointIndexReference) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *IndexedTimePointIndexReference) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *IndexedTimePointIndexReference) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *IndexedTimePointIndexReference) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *IndexedTimePointIndexReference) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.


### GetDay

`func (o *IndexedTimePointIndexReference) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *IndexedTimePointIndexReference) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *IndexedTimePointIndexReference) SetDay(v int32)`

SetDay sets Day field to given value.


### GetMonth

`func (o *IndexedTimePointIndexReference) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *IndexedTimePointIndexReference) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *IndexedTimePointIndexReference) SetMonth(v int32)`

SetMonth sets Month field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


