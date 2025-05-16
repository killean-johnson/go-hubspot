# ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule

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

### NewApiPlatformFlowCreateRequestAllOfEnrollmentSchedule

`func NewApiPlatformFlowCreateRequestAllOfEnrollmentSchedule(type_ string, timeOfDay ApiTimeOfDay, daysOfWeek []string, daysOfMonth []int32, monthlyRelativeDays string, month string, dayOfMonth int32, dateProperty string, daysDelta int32, yearly bool, ) *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule`

NewApiPlatformFlowCreateRequestAllOfEnrollmentSchedule instantiates a new ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPlatformFlowCreateRequestAllOfEnrollmentScheduleWithDefaults

`func NewApiPlatformFlowCreateRequestAllOfEnrollmentScheduleWithDefaults() *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule`

NewApiPlatformFlowCreateRequestAllOfEnrollmentScheduleWithDefaults instantiates a new ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetType(v string)`

SetType sets Type field to given value.


### GetTimeOfDay

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetTimeOfDay() ApiTimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetTimeOfDayOk() (*ApiTimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetTimeOfDay(v ApiTimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.


### GetDaysOfWeek

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDaysOfWeek() []string`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDaysOfWeekOk() (*[]string, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetDaysOfWeek(v []string)`

SetDaysOfWeek sets DaysOfWeek field to given value.


### GetDaysOfMonth

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDaysOfMonth() []int32`

GetDaysOfMonth returns the DaysOfMonth field if non-nil, zero value otherwise.

### GetDaysOfMonthOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDaysOfMonthOk() (*[]int32, bool)`

GetDaysOfMonthOk returns a tuple with the DaysOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfMonth

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetDaysOfMonth(v []int32)`

SetDaysOfMonth sets DaysOfMonth field to given value.


### GetMonthlyRelativeDays

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetMonthlyRelativeDays() string`

GetMonthlyRelativeDays returns the MonthlyRelativeDays field if non-nil, zero value otherwise.

### GetMonthlyRelativeDaysOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetMonthlyRelativeDaysOk() (*string, bool)`

GetMonthlyRelativeDaysOk returns a tuple with the MonthlyRelativeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyRelativeDays

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetMonthlyRelativeDays(v string)`

SetMonthlyRelativeDays sets MonthlyRelativeDays field to given value.


### GetMonth

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetDayOfMonth

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.


### GetDateProperty

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDateProperty() string`

GetDateProperty returns the DateProperty field if non-nil, zero value otherwise.

### GetDatePropertyOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDatePropertyOk() (*string, bool)`

GetDatePropertyOk returns a tuple with the DateProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProperty

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetDateProperty(v string)`

SetDateProperty sets DateProperty field to given value.


### GetDaysDelta

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDaysDelta() int32`

GetDaysDelta returns the DaysDelta field if non-nil, zero value otherwise.

### GetDaysDeltaOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetDaysDeltaOk() (*int32, bool)`

GetDaysDeltaOk returns a tuple with the DaysDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysDelta

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetDaysDelta(v int32)`

SetDaysDelta sets DaysDelta field to given value.


### GetYearly

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetYearly() bool`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) GetYearlyOk() (*bool, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule) SetYearly(v bool)`

SetYearly sets Yearly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


