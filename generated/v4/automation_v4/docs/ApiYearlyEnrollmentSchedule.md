# ApiYearlyEnrollmentSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | The month of the date each year to run this flow. | 
**DayOfMonth** | **int32** | The day of the date each year to run this flow. | 
**Type** | **string** | The type of enrollment schedule this is, can be: \&quot;DAILY\&quot;, \&quot;WEEKLY\&quot;, \&quot;MONTHLY_SPECIFIC_DAYS\&quot;, \&quot;MONTHLY_RELATIVE_DAYS\&quot;, \&quot;YEARLY\&quot; | [default to "YEARLY"]
**TimeOfDay** | [**ApiTimeOfDay**](ApiTimeOfDay.md) |  | 

## Methods

### NewApiYearlyEnrollmentSchedule

`func NewApiYearlyEnrollmentSchedule(month string, dayOfMonth int32, type_ string, timeOfDay ApiTimeOfDay, ) *ApiYearlyEnrollmentSchedule`

NewApiYearlyEnrollmentSchedule instantiates a new ApiYearlyEnrollmentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiYearlyEnrollmentScheduleWithDefaults

`func NewApiYearlyEnrollmentScheduleWithDefaults() *ApiYearlyEnrollmentSchedule`

NewApiYearlyEnrollmentScheduleWithDefaults instantiates a new ApiYearlyEnrollmentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *ApiYearlyEnrollmentSchedule) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ApiYearlyEnrollmentSchedule) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ApiYearlyEnrollmentSchedule) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetDayOfMonth

`func (o *ApiYearlyEnrollmentSchedule) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *ApiYearlyEnrollmentSchedule) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *ApiYearlyEnrollmentSchedule) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.


### GetType

`func (o *ApiYearlyEnrollmentSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiYearlyEnrollmentSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiYearlyEnrollmentSchedule) SetType(v string)`

SetType sets Type field to given value.


### GetTimeOfDay

`func (o *ApiYearlyEnrollmentSchedule) GetTimeOfDay() ApiTimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *ApiYearlyEnrollmentSchedule) GetTimeOfDayOk() (*ApiTimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *ApiYearlyEnrollmentSchedule) SetTimeOfDay(v ApiTimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


