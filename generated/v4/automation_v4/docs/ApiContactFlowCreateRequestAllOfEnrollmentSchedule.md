# ApiContactFlowCreateRequestAllOfEnrollmentSchedule

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

### NewApiContactFlowCreateRequestAllOfEnrollmentSchedule

`func NewApiContactFlowCreateRequestAllOfEnrollmentSchedule(type_ string, timeOfDay ApiTimeOfDay, daysOfWeek []string, daysOfMonth []int32, monthlyRelativeDays string, month string, dayOfMonth int32, dateProperty string, daysDelta int32, yearly bool, ) *ApiContactFlowCreateRequestAllOfEnrollmentSchedule`

NewApiContactFlowCreateRequestAllOfEnrollmentSchedule instantiates a new ApiContactFlowCreateRequestAllOfEnrollmentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactFlowCreateRequestAllOfEnrollmentScheduleWithDefaults

`func NewApiContactFlowCreateRequestAllOfEnrollmentScheduleWithDefaults() *ApiContactFlowCreateRequestAllOfEnrollmentSchedule`

NewApiContactFlowCreateRequestAllOfEnrollmentScheduleWithDefaults instantiates a new ApiContactFlowCreateRequestAllOfEnrollmentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetType(v string)`

SetType sets Type field to given value.


### GetTimeOfDay

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetTimeOfDay() ApiTimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetTimeOfDayOk() (*ApiTimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetTimeOfDay(v ApiTimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.


### GetDaysOfWeek

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDaysOfWeek() []string`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDaysOfWeekOk() (*[]string, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetDaysOfWeek(v []string)`

SetDaysOfWeek sets DaysOfWeek field to given value.


### GetDaysOfMonth

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDaysOfMonth() []int32`

GetDaysOfMonth returns the DaysOfMonth field if non-nil, zero value otherwise.

### GetDaysOfMonthOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDaysOfMonthOk() (*[]int32, bool)`

GetDaysOfMonthOk returns a tuple with the DaysOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfMonth

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetDaysOfMonth(v []int32)`

SetDaysOfMonth sets DaysOfMonth field to given value.


### GetMonthlyRelativeDays

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetMonthlyRelativeDays() string`

GetMonthlyRelativeDays returns the MonthlyRelativeDays field if non-nil, zero value otherwise.

### GetMonthlyRelativeDaysOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetMonthlyRelativeDaysOk() (*string, bool)`

GetMonthlyRelativeDaysOk returns a tuple with the MonthlyRelativeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyRelativeDays

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetMonthlyRelativeDays(v string)`

SetMonthlyRelativeDays sets MonthlyRelativeDays field to given value.


### GetMonth

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetDayOfMonth

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.


### GetDateProperty

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDateProperty() string`

GetDateProperty returns the DateProperty field if non-nil, zero value otherwise.

### GetDatePropertyOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDatePropertyOk() (*string, bool)`

GetDatePropertyOk returns a tuple with the DateProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProperty

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetDateProperty(v string)`

SetDateProperty sets DateProperty field to given value.


### GetDaysDelta

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDaysDelta() int32`

GetDaysDelta returns the DaysDelta field if non-nil, zero value otherwise.

### GetDaysDeltaOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetDaysDeltaOk() (*int32, bool)`

GetDaysDeltaOk returns a tuple with the DaysDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysDelta

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetDaysDelta(v int32)`

SetDaysDelta sets DaysDelta field to given value.


### GetYearly

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetYearly() bool`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) GetYearlyOk() (*bool, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *ApiContactFlowCreateRequestAllOfEnrollmentSchedule) SetYearly(v bool)`

SetYearly sets Yearly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


