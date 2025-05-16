# ApiMonthlySpecificDaysEnrollmentSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfMonth** | **[]int32** | Which days of the month to run this workflow on. | 
**Type** | **string** | The type of enrollment schedule this is, can be: \&quot;DAILY\&quot;, \&quot;WEEKLY\&quot;, \&quot;MONTHLY_SPECIFIC_DAYS\&quot;, \&quot;MONTHLY_RELATIVE_DAYS\&quot;, \&quot;YEARLY\&quot; | [default to "MONTHLY_SPECIFIC_DAYS"]
**TimeOfDay** | [**ApiTimeOfDay**](ApiTimeOfDay.md) |  | 

## Methods

### NewApiMonthlySpecificDaysEnrollmentSchedule

`func NewApiMonthlySpecificDaysEnrollmentSchedule(daysOfMonth []int32, type_ string, timeOfDay ApiTimeOfDay, ) *ApiMonthlySpecificDaysEnrollmentSchedule`

NewApiMonthlySpecificDaysEnrollmentSchedule instantiates a new ApiMonthlySpecificDaysEnrollmentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonthlySpecificDaysEnrollmentScheduleWithDefaults

`func NewApiMonthlySpecificDaysEnrollmentScheduleWithDefaults() *ApiMonthlySpecificDaysEnrollmentSchedule`

NewApiMonthlySpecificDaysEnrollmentScheduleWithDefaults instantiates a new ApiMonthlySpecificDaysEnrollmentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysOfMonth

`func (o *ApiMonthlySpecificDaysEnrollmentSchedule) GetDaysOfMonth() []int32`

GetDaysOfMonth returns the DaysOfMonth field if non-nil, zero value otherwise.

### GetDaysOfMonthOk

`func (o *ApiMonthlySpecificDaysEnrollmentSchedule) GetDaysOfMonthOk() (*[]int32, bool)`

GetDaysOfMonthOk returns a tuple with the DaysOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfMonth

`func (o *ApiMonthlySpecificDaysEnrollmentSchedule) SetDaysOfMonth(v []int32)`

SetDaysOfMonth sets DaysOfMonth field to given value.


### GetType

`func (o *ApiMonthlySpecificDaysEnrollmentSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiMonthlySpecificDaysEnrollmentSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiMonthlySpecificDaysEnrollmentSchedule) SetType(v string)`

SetType sets Type field to given value.


### GetTimeOfDay

`func (o *ApiMonthlySpecificDaysEnrollmentSchedule) GetTimeOfDay() ApiTimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *ApiMonthlySpecificDaysEnrollmentSchedule) GetTimeOfDayOk() (*ApiTimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *ApiMonthlySpecificDaysEnrollmentSchedule) SetTimeOfDay(v ApiTimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


