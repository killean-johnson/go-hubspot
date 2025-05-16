# ApiPlatformFlowPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "PLATFORM_FLOW"]
**RevisionId** | **string** |  | 
**IsEnabled** | **bool** |  | 
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
**SuppressionListIds** | **[]int32** |  | 
**GoalFilterBranch** | Pointer to [**ApiContactFlowPutRequestAllOfGoalFilterBranch**](ApiContactFlowPutRequestAllOfGoalFilterBranch.md) |  | [optional] 
**EventAnchor** | Pointer to [**ApiContactFlowPutRequestAllOfEventAnchor**](ApiContactFlowPutRequestAllOfEventAnchor.md) |  | [optional] 
**UnEnrollmentSetting** | Pointer to [**ApiUnEnrollmentSetting**](ApiUnEnrollmentSetting.md) |  | [optional] 
**CanEnrollFromSalesforce** | **bool** |  | 
**SuppressionFilterBranch** | Pointer to [**ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch**](ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch.md) |  | [optional] 

## Methods

### NewApiPlatformFlowPutRequest

`func NewApiPlatformFlowPutRequest(type_ string, revisionId string, isEnabled bool, actions []ApiPlatformFlowPutRequestAllOfActions, timeWindows []ApiTimeWindow, blockedDates []ApiBlockedDate, customProperties map[string]string, suppressionListIds []int32, canEnrollFromSalesforce bool, ) *ApiPlatformFlowPutRequest`

NewApiPlatformFlowPutRequest instantiates a new ApiPlatformFlowPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPlatformFlowPutRequestWithDefaults

`func NewApiPlatformFlowPutRequestWithDefaults() *ApiPlatformFlowPutRequest`

NewApiPlatformFlowPutRequestWithDefaults instantiates a new ApiPlatformFlowPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiPlatformFlowPutRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPlatformFlowPutRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPlatformFlowPutRequest) SetType(v string)`

SetType sets Type field to given value.


### GetRevisionId

`func (o *ApiPlatformFlowPutRequest) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ApiPlatformFlowPutRequest) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ApiPlatformFlowPutRequest) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetIsEnabled

`func (o *ApiPlatformFlowPutRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiPlatformFlowPutRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiPlatformFlowPutRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetName

`func (o *ApiPlatformFlowPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPlatformFlowPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPlatformFlowPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPlatformFlowPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiPlatformFlowPutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPlatformFlowPutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPlatformFlowPutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPlatformFlowPutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *ApiPlatformFlowPutRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiPlatformFlowPutRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiPlatformFlowPutRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiPlatformFlowPutRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStartActionId

`func (o *ApiPlatformFlowPutRequest) GetStartActionId() string`

GetStartActionId returns the StartActionId field if non-nil, zero value otherwise.

### GetStartActionIdOk

`func (o *ApiPlatformFlowPutRequest) GetStartActionIdOk() (*string, bool)`

GetStartActionIdOk returns a tuple with the StartActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActionId

`func (o *ApiPlatformFlowPutRequest) SetStartActionId(v string)`

SetStartActionId sets StartActionId field to given value.

### HasStartActionId

`func (o *ApiPlatformFlowPutRequest) HasStartActionId() bool`

HasStartActionId returns a boolean if a field has been set.

### GetActions

`func (o *ApiPlatformFlowPutRequest) GetActions() []ApiPlatformFlowPutRequestAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPlatformFlowPutRequest) GetActionsOk() (*[]ApiPlatformFlowPutRequestAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPlatformFlowPutRequest) SetActions(v []ApiPlatformFlowPutRequestAllOfActions)`

SetActions sets Actions field to given value.


### GetEnrollmentCriteria

`func (o *ApiPlatformFlowPutRequest) GetEnrollmentCriteria() ApiPlatformFlowPutRequestAllOfEnrollmentCriteria`

GetEnrollmentCriteria returns the EnrollmentCriteria field if non-nil, zero value otherwise.

### GetEnrollmentCriteriaOk

`func (o *ApiPlatformFlowPutRequest) GetEnrollmentCriteriaOk() (*ApiPlatformFlowPutRequestAllOfEnrollmentCriteria, bool)`

GetEnrollmentCriteriaOk returns a tuple with the EnrollmentCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCriteria

`func (o *ApiPlatformFlowPutRequest) SetEnrollmentCriteria(v ApiPlatformFlowPutRequestAllOfEnrollmentCriteria)`

SetEnrollmentCriteria sets EnrollmentCriteria field to given value.

### HasEnrollmentCriteria

`func (o *ApiPlatformFlowPutRequest) HasEnrollmentCriteria() bool`

HasEnrollmentCriteria returns a boolean if a field has been set.

### GetEnrollmentSchedule

`func (o *ApiPlatformFlowPutRequest) GetEnrollmentSchedule() ApiPlatformFlowPutRequestAllOfEnrollmentSchedule`

GetEnrollmentSchedule returns the EnrollmentSchedule field if non-nil, zero value otherwise.

### GetEnrollmentScheduleOk

`func (o *ApiPlatformFlowPutRequest) GetEnrollmentScheduleOk() (*ApiPlatformFlowPutRequestAllOfEnrollmentSchedule, bool)`

GetEnrollmentScheduleOk returns a tuple with the EnrollmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSchedule

`func (o *ApiPlatformFlowPutRequest) SetEnrollmentSchedule(v ApiPlatformFlowPutRequestAllOfEnrollmentSchedule)`

SetEnrollmentSchedule sets EnrollmentSchedule field to given value.

### HasEnrollmentSchedule

`func (o *ApiPlatformFlowPutRequest) HasEnrollmentSchedule() bool`

HasEnrollmentSchedule returns a boolean if a field has been set.

### GetTimeWindows

`func (o *ApiPlatformFlowPutRequest) GetTimeWindows() []ApiTimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *ApiPlatformFlowPutRequest) GetTimeWindowsOk() (*[]ApiTimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *ApiPlatformFlowPutRequest) SetTimeWindows(v []ApiTimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetBlockedDates

`func (o *ApiPlatformFlowPutRequest) GetBlockedDates() []ApiBlockedDate`

GetBlockedDates returns the BlockedDates field if non-nil, zero value otherwise.

### GetBlockedDatesOk

`func (o *ApiPlatformFlowPutRequest) GetBlockedDatesOk() (*[]ApiBlockedDate, bool)`

GetBlockedDatesOk returns a tuple with the BlockedDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDates

`func (o *ApiPlatformFlowPutRequest) SetBlockedDates(v []ApiBlockedDate)`

SetBlockedDates sets BlockedDates field to given value.


### GetCustomProperties

`func (o *ApiPlatformFlowPutRequest) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ApiPlatformFlowPutRequest) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ApiPlatformFlowPutRequest) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.


### GetSuppressionListIds

`func (o *ApiPlatformFlowPutRequest) GetSuppressionListIds() []int32`

GetSuppressionListIds returns the SuppressionListIds field if non-nil, zero value otherwise.

### GetSuppressionListIdsOk

`func (o *ApiPlatformFlowPutRequest) GetSuppressionListIdsOk() (*[]int32, bool)`

GetSuppressionListIdsOk returns a tuple with the SuppressionListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionListIds

`func (o *ApiPlatformFlowPutRequest) SetSuppressionListIds(v []int32)`

SetSuppressionListIds sets SuppressionListIds field to given value.


### GetGoalFilterBranch

`func (o *ApiPlatformFlowPutRequest) GetGoalFilterBranch() ApiContactFlowPutRequestAllOfGoalFilterBranch`

GetGoalFilterBranch returns the GoalFilterBranch field if non-nil, zero value otherwise.

### GetGoalFilterBranchOk

`func (o *ApiPlatformFlowPutRequest) GetGoalFilterBranchOk() (*ApiContactFlowPutRequestAllOfGoalFilterBranch, bool)`

GetGoalFilterBranchOk returns a tuple with the GoalFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilterBranch

`func (o *ApiPlatformFlowPutRequest) SetGoalFilterBranch(v ApiContactFlowPutRequestAllOfGoalFilterBranch)`

SetGoalFilterBranch sets GoalFilterBranch field to given value.

### HasGoalFilterBranch

`func (o *ApiPlatformFlowPutRequest) HasGoalFilterBranch() bool`

HasGoalFilterBranch returns a boolean if a field has been set.

### GetEventAnchor

`func (o *ApiPlatformFlowPutRequest) GetEventAnchor() ApiContactFlowPutRequestAllOfEventAnchor`

GetEventAnchor returns the EventAnchor field if non-nil, zero value otherwise.

### GetEventAnchorOk

`func (o *ApiPlatformFlowPutRequest) GetEventAnchorOk() (*ApiContactFlowPutRequestAllOfEventAnchor, bool)`

GetEventAnchorOk returns a tuple with the EventAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAnchor

`func (o *ApiPlatformFlowPutRequest) SetEventAnchor(v ApiContactFlowPutRequestAllOfEventAnchor)`

SetEventAnchor sets EventAnchor field to given value.

### HasEventAnchor

`func (o *ApiPlatformFlowPutRequest) HasEventAnchor() bool`

HasEventAnchor returns a boolean if a field has been set.

### GetUnEnrollmentSetting

`func (o *ApiPlatformFlowPutRequest) GetUnEnrollmentSetting() ApiUnEnrollmentSetting`

GetUnEnrollmentSetting returns the UnEnrollmentSetting field if non-nil, zero value otherwise.

### GetUnEnrollmentSettingOk

`func (o *ApiPlatformFlowPutRequest) GetUnEnrollmentSettingOk() (*ApiUnEnrollmentSetting, bool)`

GetUnEnrollmentSettingOk returns a tuple with the UnEnrollmentSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollmentSetting

`func (o *ApiPlatformFlowPutRequest) SetUnEnrollmentSetting(v ApiUnEnrollmentSetting)`

SetUnEnrollmentSetting sets UnEnrollmentSetting field to given value.

### HasUnEnrollmentSetting

`func (o *ApiPlatformFlowPutRequest) HasUnEnrollmentSetting() bool`

HasUnEnrollmentSetting returns a boolean if a field has been set.

### GetCanEnrollFromSalesforce

`func (o *ApiPlatformFlowPutRequest) GetCanEnrollFromSalesforce() bool`

GetCanEnrollFromSalesforce returns the CanEnrollFromSalesforce field if non-nil, zero value otherwise.

### GetCanEnrollFromSalesforceOk

`func (o *ApiPlatformFlowPutRequest) GetCanEnrollFromSalesforceOk() (*bool, bool)`

GetCanEnrollFromSalesforceOk returns a tuple with the CanEnrollFromSalesforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnrollFromSalesforce

`func (o *ApiPlatformFlowPutRequest) SetCanEnrollFromSalesforce(v bool)`

SetCanEnrollFromSalesforce sets CanEnrollFromSalesforce field to given value.


### GetSuppressionFilterBranch

`func (o *ApiPlatformFlowPutRequest) GetSuppressionFilterBranch() ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch`

GetSuppressionFilterBranch returns the SuppressionFilterBranch field if non-nil, zero value otherwise.

### GetSuppressionFilterBranchOk

`func (o *ApiPlatformFlowPutRequest) GetSuppressionFilterBranchOk() (*ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch, bool)`

GetSuppressionFilterBranchOk returns a tuple with the SuppressionFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionFilterBranch

`func (o *ApiPlatformFlowPutRequest) SetSuppressionFilterBranch(v ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch)`

SetSuppressionFilterBranch sets SuppressionFilterBranch field to given value.

### HasSuppressionFilterBranch

`func (o *ApiPlatformFlowPutRequest) HasSuppressionFilterBranch() bool`

HasSuppressionFilterBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


