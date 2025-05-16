# ApiContactFlowAllOfEnrollmentSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of enrollment schedule this is, can be: \&quot;DAILY\&quot;, \&quot;WEEKLY\&quot;, \&quot;MONTHLY_SPECIFIC_DAYS\&quot;, \&quot;MONTHLY_RELATIVE_DAYS\&quot;, \&quot;YEARLY\&quot; | [default to "DAILY"]
**TimeOfDay** | [**ApiTimeOfDay**](ApiTimeOfDay.md) |  | 
**DaysOfWeek** | **[]string** | Which days of the week to allow enrollments. | 
**DaysOfMonth** | **[]int32** | Which days of the month to run this workflow on. | 
**MonthlyRelativeDays** | **string** | Can be either \&quot;LAST_DAY_OF_MONTH\&quot; or \&quot;FIRST_MONDAY_OF_MONTH\&quot; | 
**Month** | **string** | The month of the date each year to run this flow. | 
**DayOfMonth** | **int32** | The day of the date each year to run this flow. | 
**DateProperty** | **string** |  | 
**DaysDelta** | **int32** |  | 
**Yearly** | **bool** |  | 

## Methods

### NewApiContactFlowAllOfEnrollmentSchedule

`func NewApiContactFlowAllOfEnrollmentSchedule(type_ string, timeOfDay ApiTimeOfDay, daysOfWeek []string, daysOfMonth []int32, monthlyRelativeDays string, month string, dayOfMonth int32, dateProperty string, daysDelta int32, yearly bool, ) *ApiContactFlowAllOfEnrollmentSchedule`

NewApiContactFlowAllOfEnrollmentSchedule instantiates a new ApiContactFlowAllOfEnrollmentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactFlowAllOfEnrollmentScheduleWithDefaults

`func NewApiContactFlowAllOfEnrollmentScheduleWithDefaults() *ApiContactFlowAllOfEnrollmentSchedule`

NewApiContactFlowAllOfEnrollmentScheduleWithDefaults instantiates a new ApiContactFlowAllOfEnrollmentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetType(v string)`

SetType sets Type field to given value.


### GetTimeOfDay

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetTimeOfDay() ApiTimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetTimeOfDayOk() (*ApiTimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetTimeOfDay(v ApiTimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.


### GetDaysOfWeek

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDaysOfWeek() []string`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDaysOfWeekOk() (*[]string, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetDaysOfWeek(v []string)`

SetDaysOfWeek sets DaysOfWeek field to given value.


### GetDaysOfMonth

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDaysOfMonth() []int32`

GetDaysOfMonth returns the DaysOfMonth field if non-nil, zero value otherwise.

### GetDaysOfMonthOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDaysOfMonthOk() (*[]int32, bool)`

GetDaysOfMonthOk returns a tuple with the DaysOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfMonth

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetDaysOfMonth(v []int32)`

SetDaysOfMonth sets DaysOfMonth field to given value.


### GetMonthlyRelativeDays

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetMonthlyRelativeDays() string`

GetMonthlyRelativeDays returns the MonthlyRelativeDays field if non-nil, zero value otherwise.

### GetMonthlyRelativeDaysOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetMonthlyRelativeDaysOk() (*string, bool)`

GetMonthlyRelativeDaysOk returns a tuple with the MonthlyRelativeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyRelativeDays

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetMonthlyRelativeDays(v string)`

SetMonthlyRelativeDays sets MonthlyRelativeDays field to given value.


### GetMonth

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetDayOfMonth

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.


### GetDateProperty

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDateProperty() string`

GetDateProperty returns the DateProperty field if non-nil, zero value otherwise.

### GetDatePropertyOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDatePropertyOk() (*string, bool)`

GetDatePropertyOk returns a tuple with the DateProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProperty

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetDateProperty(v string)`

SetDateProperty sets DateProperty field to given value.


### GetDaysDelta

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDaysDelta() int32`

GetDaysDelta returns the DaysDelta field if non-nil, zero value otherwise.

### GetDaysDeltaOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetDaysDeltaOk() (*int32, bool)`

GetDaysDeltaOk returns a tuple with the DaysDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysDelta

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetDaysDelta(v int32)`

SetDaysDelta sets DaysDelta field to given value.


### GetYearly

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetYearly() bool`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *ApiContactFlowAllOfEnrollmentSchedule) GetYearlyOk() (*bool, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *ApiContactFlowAllOfEnrollmentSchedule) SetYearly(v bool)`

SetYearly sets Yearly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


