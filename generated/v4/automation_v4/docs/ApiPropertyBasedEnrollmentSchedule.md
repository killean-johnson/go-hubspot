# ApiPropertyBasedEnrollmentSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateProperty** | **string** |  | 
**DaysDelta** | **int32** |  | 
**Type** | **string** |  | [default to "PROPERTY_BASED"]
**Yearly** | **bool** |  | 
**TimeOfDay** | [**ApiTimeOfDay**](ApiTimeOfDay.md) |  | 

## Methods

### NewApiPropertyBasedEnrollmentSchedule

`func NewApiPropertyBasedEnrollmentSchedule(dateProperty string, daysDelta int32, type_ string, yearly bool, timeOfDay ApiTimeOfDay, ) *ApiPropertyBasedEnrollmentSchedule`

NewApiPropertyBasedEnrollmentSchedule instantiates a new ApiPropertyBasedEnrollmentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPropertyBasedEnrollmentScheduleWithDefaults

`func NewApiPropertyBasedEnrollmentScheduleWithDefaults() *ApiPropertyBasedEnrollmentSchedule`

NewApiPropertyBasedEnrollmentScheduleWithDefaults instantiates a new ApiPropertyBasedEnrollmentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateProperty

`func (o *ApiPropertyBasedEnrollmentSchedule) GetDateProperty() string`

GetDateProperty returns the DateProperty field if non-nil, zero value otherwise.

### GetDatePropertyOk

`func (o *ApiPropertyBasedEnrollmentSchedule) GetDatePropertyOk() (*string, bool)`

GetDatePropertyOk returns a tuple with the DateProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProperty

`func (o *ApiPropertyBasedEnrollmentSchedule) SetDateProperty(v string)`

SetDateProperty sets DateProperty field to given value.


### GetDaysDelta

`func (o *ApiPropertyBasedEnrollmentSchedule) GetDaysDelta() int32`

GetDaysDelta returns the DaysDelta field if non-nil, zero value otherwise.

### GetDaysDeltaOk

`func (o *ApiPropertyBasedEnrollmentSchedule) GetDaysDeltaOk() (*int32, bool)`

GetDaysDeltaOk returns a tuple with the DaysDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysDelta

`func (o *ApiPropertyBasedEnrollmentSchedule) SetDaysDelta(v int32)`

SetDaysDelta sets DaysDelta field to given value.


### GetType

`func (o *ApiPropertyBasedEnrollmentSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPropertyBasedEnrollmentSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPropertyBasedEnrollmentSchedule) SetType(v string)`

SetType sets Type field to given value.


### GetYearly

`func (o *ApiPropertyBasedEnrollmentSchedule) GetYearly() bool`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *ApiPropertyBasedEnrollmentSchedule) GetYearlyOk() (*bool, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *ApiPropertyBasedEnrollmentSchedule) SetYearly(v bool)`

SetYearly sets Yearly field to given value.


### GetTimeOfDay

`func (o *ApiPropertyBasedEnrollmentSchedule) GetTimeOfDay() ApiTimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *ApiPropertyBasedEnrollmentSchedule) GetTimeOfDayOk() (*ApiTimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *ApiPropertyBasedEnrollmentSchedule) SetTimeOfDay(v ApiTimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


