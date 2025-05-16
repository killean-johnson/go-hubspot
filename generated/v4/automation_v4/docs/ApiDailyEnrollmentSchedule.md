# ApiDailyEnrollmentSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of enrollment schedule this is, can be: \&quot;DAILY\&quot;, \&quot;WEEKLY\&quot;, \&quot;MONTHLY_SPECIFIC_DAYS\&quot;, \&quot;MONTHLY_RELATIVE_DAYS\&quot;, \&quot;YEARLY\&quot; | [default to "DAILY"]
**TimeOfDay** | [**ApiTimeOfDay**](ApiTimeOfDay.md) |  | 

## Methods

### NewApiDailyEnrollmentSchedule

`func NewApiDailyEnrollmentSchedule(type_ string, timeOfDay ApiTimeOfDay, ) *ApiDailyEnrollmentSchedule`

NewApiDailyEnrollmentSchedule instantiates a new ApiDailyEnrollmentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDailyEnrollmentScheduleWithDefaults

`func NewApiDailyEnrollmentScheduleWithDefaults() *ApiDailyEnrollmentSchedule`

NewApiDailyEnrollmentScheduleWithDefaults instantiates a new ApiDailyEnrollmentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiDailyEnrollmentSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiDailyEnrollmentSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiDailyEnrollmentSchedule) SetType(v string)`

SetType sets Type field to given value.


### GetTimeOfDay

`func (o *ApiDailyEnrollmentSchedule) GetTimeOfDay() ApiTimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *ApiDailyEnrollmentSchedule) GetTimeOfDayOk() (*ApiTimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *ApiDailyEnrollmentSchedule) SetTimeOfDay(v ApiTimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


