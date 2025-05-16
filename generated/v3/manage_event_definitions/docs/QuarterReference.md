# QuarterReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to **int32** |  | [optional] 
**Month** | **int32** |  | 
**Millisecond** | Pointer to **int32** |  | [optional] 
**ReferenceType** | **string** |  | [default to "QUARTER"]
**Day** | **int32** |  | 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 

## Methods

### NewQuarterReference

`func NewQuarterReference(month int32, referenceType string, day int32, ) *QuarterReference`

NewQuarterReference instantiates a new QuarterReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuarterReferenceWithDefaults

`func NewQuarterReferenceWithDefaults() *QuarterReference`

NewQuarterReferenceWithDefaults instantiates a new QuarterReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *QuarterReference) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *QuarterReference) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *QuarterReference) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *QuarterReference) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMonth

`func (o *QuarterReference) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *QuarterReference) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *QuarterReference) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetMillisecond

`func (o *QuarterReference) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *QuarterReference) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *QuarterReference) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *QuarterReference) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetReferenceType

`func (o *QuarterReference) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *QuarterReference) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *QuarterReference) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.


### GetDay

`func (o *QuarterReference) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *QuarterReference) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *QuarterReference) SetDay(v int32)`

SetDay sets Day field to given value.


### GetMinute

`func (o *QuarterReference) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *QuarterReference) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *QuarterReference) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *QuarterReference) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *QuarterReference) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *QuarterReference) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *QuarterReference) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *QuarterReference) HasSecond() bool`

HasSecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


