# MonthReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to **int32** |  | [optional] 
**Millisecond** | Pointer to **int32** |  | [optional] 
**ReferenceType** | **string** |  | [default to "MONTH"]
**Day** | **int32** |  | 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 

## Methods

### NewMonthReference

`func NewMonthReference(referenceType string, day int32, ) *MonthReference`

NewMonthReference instantiates a new MonthReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthReferenceWithDefaults

`func NewMonthReferenceWithDefaults() *MonthReference`

NewMonthReferenceWithDefaults instantiates a new MonthReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *MonthReference) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *MonthReference) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *MonthReference) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *MonthReference) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMillisecond

`func (o *MonthReference) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *MonthReference) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *MonthReference) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *MonthReference) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetReferenceType

`func (o *MonthReference) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *MonthReference) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *MonthReference) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.


### GetDay

`func (o *MonthReference) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *MonthReference) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *MonthReference) SetDay(v int32)`

SetDay sets Day field to given value.


### GetMinute

`func (o *MonthReference) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *MonthReference) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *MonthReference) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *MonthReference) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *MonthReference) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *MonthReference) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *MonthReference) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *MonthReference) HasSecond() bool`

HasSecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


