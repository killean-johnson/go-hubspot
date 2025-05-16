# ApiFlowPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "CONTACT_FLOW"]
**IsEnabled** | **bool** |  | 
**RevisionId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**StartActionId** | Pointer to **string** |  | [optional] 
**Actions** | [**[]ApiPlatformFlowPutRequestAllOfActions**](ApiPlatformFlowPutRequestAllOfActions.md) |  | 
**EnrollmentCriteria** | Pointer to [**ApiPlatformFlowPutRequestAllOfEnrollmentCriteria**](ApiPlatformFlowPutRequestAllOfEnrollmentCriteria.md) |  | [optional] 
**EnrollmentSchedule** | Pointer to [**ApiPlatformFlowPutRequestAllOfEnrollmentSchedule**](ApiPlatformFlowPutRequestAllOfEnrollmentSchedule.md) |  | [optional] 
**TimeWindows** | [**[]ApiTimeWindow**](ApiTimeWindow.md) |  | 
**BlockedDates** | [**[]ApiBlockedDate**](ApiBlockedDate.md) |  | 
**CustomProperties** | **map[string]string** |  | 
**SuppressionFilterBranch** | Pointer to [**ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch**](ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch.md) |  | [optional] 
**SuppressionListIds** | **[]int32** |  | 
**GoalFilterBranch** | Pointer to [**ApiContactFlowPutRequestAllOfGoalFilterBranch**](ApiContactFlowPutRequestAllOfGoalFilterBranch.md) |  | [optional] 
**EventAnchor** | Pointer to [**ApiContactFlowPutRequestAllOfEventAnchor**](ApiContactFlowPutRequestAllOfEventAnchor.md) |  | [optional] 
**UnEnrollmentSetting** | Pointer to [**ApiUnEnrollmentSetting**](ApiUnEnrollmentSetting.md) |  | [optional] 
**CanEnrollFromSalesforce** | **bool** |  | 

## Methods

### NewApiFlowPutRequest

`func NewApiFlowPutRequest(type_ string, isEnabled bool, revisionId string, actions []ApiPlatformFlowPutRequestAllOfActions, timeWindows []ApiTimeWindow, blockedDates []ApiBlockedDate, customProperties map[string]string, suppressionListIds []int32, canEnrollFromSalesforce bool, ) *ApiFlowPutRequest`

NewApiFlowPutRequest instantiates a new ApiFlowPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFlowPutRequestWithDefaults

`func NewApiFlowPutRequestWithDefaults() *ApiFlowPutRequest`

NewApiFlowPutRequestWithDefaults instantiates a new ApiFlowPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiFlowPutRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiFlowPutRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiFlowPutRequest) SetType(v string)`

SetType sets Type field to given value.


### GetIsEnabled

`func (o *ApiFlowPutRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiFlowPutRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiFlowPutRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetRevisionId

`func (o *ApiFlowPutRequest) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ApiFlowPutRequest) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ApiFlowPutRequest) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetName

`func (o *ApiFlowPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiFlowPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiFlowPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiFlowPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiFlowPutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiFlowPutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiFlowPutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiFlowPutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *ApiFlowPutRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiFlowPutRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiFlowPutRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiFlowPutRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStartActionId

`func (o *ApiFlowPutRequest) GetStartActionId() string`

GetStartActionId returns the StartActionId field if non-nil, zero value otherwise.

### GetStartActionIdOk

`func (o *ApiFlowPutRequest) GetStartActionIdOk() (*string, bool)`

GetStartActionIdOk returns a tuple with the StartActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActionId

`func (o *ApiFlowPutRequest) SetStartActionId(v string)`

SetStartActionId sets StartActionId field to given value.

### HasStartActionId

`func (o *ApiFlowPutRequest) HasStartActionId() bool`

HasStartActionId returns a boolean if a field has been set.

### GetActions

`func (o *ApiFlowPutRequest) GetActions() []ApiPlatformFlowPutRequestAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiFlowPutRequest) GetActionsOk() (*[]ApiPlatformFlowPutRequestAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiFlowPutRequest) SetActions(v []ApiPlatformFlowPutRequestAllOfActions)`

SetActions sets Actions field to given value.


### GetEnrollmentCriteria

`func (o *ApiFlowPutRequest) GetEnrollmentCriteria() ApiPlatformFlowPutRequestAllOfEnrollmentCriteria`

GetEnrollmentCriteria returns the EnrollmentCriteria field if non-nil, zero value otherwise.

### GetEnrollmentCriteriaOk

`func (o *ApiFlowPutRequest) GetEnrollmentCriteriaOk() (*ApiPlatformFlowPutRequestAllOfEnrollmentCriteria, bool)`

GetEnrollmentCriteriaOk returns a tuple with the EnrollmentCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCriteria

`func (o *ApiFlowPutRequest) SetEnrollmentCriteria(v ApiPlatformFlowPutRequestAllOfEnrollmentCriteria)`

SetEnrollmentCriteria sets EnrollmentCriteria field to given value.

### HasEnrollmentCriteria

`func (o *ApiFlowPutRequest) HasEnrollmentCriteria() bool`

HasEnrollmentCriteria returns a boolean if a field has been set.

### GetEnrollmentSchedule

`func (o *ApiFlowPutRequest) GetEnrollmentSchedule() ApiPlatformFlowPutRequestAllOfEnrollmentSchedule`

GetEnrollmentSchedule returns the EnrollmentSchedule field if non-nil, zero value otherwise.

### GetEnrollmentScheduleOk

`func (o *ApiFlowPutRequest) GetEnrollmentScheduleOk() (*ApiPlatformFlowPutRequestAllOfEnrollmentSchedule, bool)`

GetEnrollmentScheduleOk returns a tuple with the EnrollmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSchedule

`func (o *ApiFlowPutRequest) SetEnrollmentSchedule(v ApiPlatformFlowPutRequestAllOfEnrollmentSchedule)`

SetEnrollmentSchedule sets EnrollmentSchedule field to given value.

### HasEnrollmentSchedule

`func (o *ApiFlowPutRequest) HasEnrollmentSchedule() bool`

HasEnrollmentSchedule returns a boolean if a field has been set.

### GetTimeWindows

`func (o *ApiFlowPutRequest) GetTimeWindows() []ApiTimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *ApiFlowPutRequest) GetTimeWindowsOk() (*[]ApiTimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *ApiFlowPutRequest) SetTimeWindows(v []ApiTimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetBlockedDates

`func (o *ApiFlowPutRequest) GetBlockedDates() []ApiBlockedDate`

GetBlockedDates returns the BlockedDates field if non-nil, zero value otherwise.

### GetBlockedDatesOk

`func (o *ApiFlowPutRequest) GetBlockedDatesOk() (*[]ApiBlockedDate, bool)`

GetBlockedDatesOk returns a tuple with the BlockedDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDates

`func (o *ApiFlowPutRequest) SetBlockedDates(v []ApiBlockedDate)`

SetBlockedDates sets BlockedDates field to given value.


### GetCustomProperties

`func (o *ApiFlowPutRequest) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ApiFlowPutRequest) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ApiFlowPutRequest) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.


### GetSuppressionFilterBranch

`func (o *ApiFlowPutRequest) GetSuppressionFilterBranch() ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch`

GetSuppressionFilterBranch returns the SuppressionFilterBranch field if non-nil, zero value otherwise.

### GetSuppressionFilterBranchOk

`func (o *ApiFlowPutRequest) GetSuppressionFilterBranchOk() (*ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch, bool)`

GetSuppressionFilterBranchOk returns a tuple with the SuppressionFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionFilterBranch

`func (o *ApiFlowPutRequest) SetSuppressionFilterBranch(v ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch)`

SetSuppressionFilterBranch sets SuppressionFilterBranch field to given value.

### HasSuppressionFilterBranch

`func (o *ApiFlowPutRequest) HasSuppressionFilterBranch() bool`

HasSuppressionFilterBranch returns a boolean if a field has been set.

### GetSuppressionListIds

`func (o *ApiFlowPutRequest) GetSuppressionListIds() []int32`

GetSuppressionListIds returns the SuppressionListIds field if non-nil, zero value otherwise.

### GetSuppressionListIdsOk

`func (o *ApiFlowPutRequest) GetSuppressionListIdsOk() (*[]int32, bool)`

GetSuppressionListIdsOk returns a tuple with the SuppressionListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionListIds

`func (o *ApiFlowPutRequest) SetSuppressionListIds(v []int32)`

SetSuppressionListIds sets SuppressionListIds field to given value.


### GetGoalFilterBranch

`func (o *ApiFlowPutRequest) GetGoalFilterBranch() ApiContactFlowPutRequestAllOfGoalFilterBranch`

GetGoalFilterBranch returns the GoalFilterBranch field if non-nil, zero value otherwise.

### GetGoalFilterBranchOk

`func (o *ApiFlowPutRequest) GetGoalFilterBranchOk() (*ApiContactFlowPutRequestAllOfGoalFilterBranch, bool)`

GetGoalFilterBranchOk returns a tuple with the GoalFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilterBranch

`func (o *ApiFlowPutRequest) SetGoalFilterBranch(v ApiContactFlowPutRequestAllOfGoalFilterBranch)`

SetGoalFilterBranch sets GoalFilterBranch field to given value.

### HasGoalFilterBranch

`func (o *ApiFlowPutRequest) HasGoalFilterBranch() bool`

HasGoalFilterBranch returns a boolean if a field has been set.

### GetEventAnchor

`func (o *ApiFlowPutRequest) GetEventAnchor() ApiContactFlowPutRequestAllOfEventAnchor`

GetEventAnchor returns the EventAnchor field if non-nil, zero value otherwise.

### GetEventAnchorOk

`func (o *ApiFlowPutRequest) GetEventAnchorOk() (*ApiContactFlowPutRequestAllOfEventAnchor, bool)`

GetEventAnchorOk returns a tuple with the EventAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAnchor

`func (o *ApiFlowPutRequest) SetEventAnchor(v ApiContactFlowPutRequestAllOfEventAnchor)`

SetEventAnchor sets EventAnchor field to given value.

### HasEventAnchor

`func (o *ApiFlowPutRequest) HasEventAnchor() bool`

HasEventAnchor returns a boolean if a field has been set.

### GetUnEnrollmentSetting

`func (o *ApiFlowPutRequest) GetUnEnrollmentSetting() ApiUnEnrollmentSetting`

GetUnEnrollmentSetting returns the UnEnrollmentSetting field if non-nil, zero value otherwise.

### GetUnEnrollmentSettingOk

`func (o *ApiFlowPutRequest) GetUnEnrollmentSettingOk() (*ApiUnEnrollmentSetting, bool)`

GetUnEnrollmentSettingOk returns a tuple with the UnEnrollmentSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollmentSetting

`func (o *ApiFlowPutRequest) SetUnEnrollmentSetting(v ApiUnEnrollmentSetting)`

SetUnEnrollmentSetting sets UnEnrollmentSetting field to given value.

### HasUnEnrollmentSetting

`func (o *ApiFlowPutRequest) HasUnEnrollmentSetting() bool`

HasUnEnrollmentSetting returns a boolean if a field has been set.

### GetCanEnrollFromSalesforce

`func (o *ApiFlowPutRequest) GetCanEnrollFromSalesforce() bool`

GetCanEnrollFromSalesforce returns the CanEnrollFromSalesforce field if non-nil, zero value otherwise.

### GetCanEnrollFromSalesforceOk

`func (o *ApiFlowPutRequest) GetCanEnrollFromSalesforceOk() (*bool, bool)`

GetCanEnrollFromSalesforceOk returns a tuple with the CanEnrollFromSalesforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnrollFromSalesforce

`func (o *ApiFlowPutRequest) SetCanEnrollFromSalesforce(v bool)`

SetCanEnrollFromSalesforce sets CanEnrollFromSalesforce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


