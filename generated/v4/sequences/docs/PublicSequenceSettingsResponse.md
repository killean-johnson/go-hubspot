# PublicSequenceSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**TaskReminderMinute** | **int32** |  | 
**IndividualTaskRemindersEnabled** | **bool** |  | 
**Id** | **string** |  | 
**SellingStrategy** | **string** |  | 
**SendWindowStartMinute** | **int32** |  | 
**SendWindowEndMinute** | **int32** |  | 
**EligibleFollowUpDays** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPublicSequenceSettingsResponse

`func NewPublicSequenceSettingsResponse(createdAt time.Time, taskReminderMinute int32, individualTaskRemindersEnabled bool, id string, sellingStrategy string, sendWindowStartMinute int32, sendWindowEndMinute int32, eligibleFollowUpDays string, updatedAt time.Time, ) *PublicSequenceSettingsResponse`

NewPublicSequenceSettingsResponse instantiates a new PublicSequenceSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSequenceSettingsResponseWithDefaults

`func NewPublicSequenceSettingsResponseWithDefaults() *PublicSequenceSettingsResponse`

NewPublicSequenceSettingsResponseWithDefaults instantiates a new PublicSequenceSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicSequenceSettingsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicSequenceSettingsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicSequenceSettingsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTaskReminderMinute

`func (o *PublicSequenceSettingsResponse) GetTaskReminderMinute() int32`

GetTaskReminderMinute returns the TaskReminderMinute field if non-nil, zero value otherwise.

### GetTaskReminderMinuteOk

`func (o *PublicSequenceSettingsResponse) GetTaskReminderMinuteOk() (*int32, bool)`

GetTaskReminderMinuteOk returns a tuple with the TaskReminderMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReminderMinute

`func (o *PublicSequenceSettingsResponse) SetTaskReminderMinute(v int32)`

SetTaskReminderMinute sets TaskReminderMinute field to given value.


### GetIndividualTaskRemindersEnabled

`func (o *PublicSequenceSettingsResponse) GetIndividualTaskRemindersEnabled() bool`

GetIndividualTaskRemindersEnabled returns the IndividualTaskRemindersEnabled field if non-nil, zero value otherwise.

### GetIndividualTaskRemindersEnabledOk

`func (o *PublicSequenceSettingsResponse) GetIndividualTaskRemindersEnabledOk() (*bool, bool)`

GetIndividualTaskRemindersEnabledOk returns a tuple with the IndividualTaskRemindersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualTaskRemindersEnabled

`func (o *PublicSequenceSettingsResponse) SetIndividualTaskRemindersEnabled(v bool)`

SetIndividualTaskRemindersEnabled sets IndividualTaskRemindersEnabled field to given value.


### GetId

`func (o *PublicSequenceSettingsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicSequenceSettingsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicSequenceSettingsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSellingStrategy

`func (o *PublicSequenceSettingsResponse) GetSellingStrategy() string`

GetSellingStrategy returns the SellingStrategy field if non-nil, zero value otherwise.

### GetSellingStrategyOk

`func (o *PublicSequenceSettingsResponse) GetSellingStrategyOk() (*string, bool)`

GetSellingStrategyOk returns a tuple with the SellingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingStrategy

`func (o *PublicSequenceSettingsResponse) SetSellingStrategy(v string)`

SetSellingStrategy sets SellingStrategy field to given value.


### GetSendWindowStartMinute

`func (o *PublicSequenceSettingsResponse) GetSendWindowStartMinute() int32`

GetSendWindowStartMinute returns the SendWindowStartMinute field if non-nil, zero value otherwise.

### GetSendWindowStartMinuteOk

`func (o *PublicSequenceSettingsResponse) GetSendWindowStartMinuteOk() (*int32, bool)`

GetSendWindowStartMinuteOk returns a tuple with the SendWindowStartMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWindowStartMinute

`func (o *PublicSequenceSettingsResponse) SetSendWindowStartMinute(v int32)`

SetSendWindowStartMinute sets SendWindowStartMinute field to given value.


### GetSendWindowEndMinute

`func (o *PublicSequenceSettingsResponse) GetSendWindowEndMinute() int32`

GetSendWindowEndMinute returns the SendWindowEndMinute field if non-nil, zero value otherwise.

### GetSendWindowEndMinuteOk

`func (o *PublicSequenceSettingsResponse) GetSendWindowEndMinuteOk() (*int32, bool)`

GetSendWindowEndMinuteOk returns a tuple with the SendWindowEndMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWindowEndMinute

`func (o *PublicSequenceSettingsResponse) SetSendWindowEndMinute(v int32)`

SetSendWindowEndMinute sets SendWindowEndMinute field to given value.


### GetEligibleFollowUpDays

`func (o *PublicSequenceSettingsResponse) GetEligibleFollowUpDays() string`

GetEligibleFollowUpDays returns the EligibleFollowUpDays field if non-nil, zero value otherwise.

### GetEligibleFollowUpDaysOk

`func (o *PublicSequenceSettingsResponse) GetEligibleFollowUpDaysOk() (*string, bool)`

GetEligibleFollowUpDaysOk returns a tuple with the EligibleFollowUpDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleFollowUpDays

`func (o *PublicSequenceSettingsResponse) SetEligibleFollowUpDays(v string)`

SetEligibleFollowUpDays sets EligibleFollowUpDays field to given value.


### GetUpdatedAt

`func (o *PublicSequenceSettingsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicSequenceSettingsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicSequenceSettingsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


