# ApiFlowCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "CONTACT_FLOW"]
**IsEnabled** | **bool** |  | 
**ObjectTypeId** | **string** |  | 
**FlowType** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**StartActionId** | Pointer to **string** |  | [optional] 
**Actions** | [**[]ApiPlatformFlowCreateRequestAllOfActions**](ApiPlatformFlowCreateRequestAllOfActions.md) |  | 
**EnrollmentCriteria** | Pointer to [**ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria**](ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria.md) |  | [optional] 
**EnrollmentSchedule** | Pointer to [**ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule**](ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule.md) |  | [optional] 
**TimeWindows** | [**[]ApiTimeWindow**](ApiTimeWindow.md) |  | 
**BlockedDates** | [**[]ApiBlockedDate**](ApiBlockedDate.md) |  | 
**CustomProperties** | **map[string]string** |  | 
**DataSources** | [**[]ApiPlatformFlowCreateRequestAllOfDataSources**](ApiPlatformFlowCreateRequestAllOfDataSources.md) |  | 
**SuppressionFilterBranch** | Pointer to [**ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch**](ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch.md) |  | [optional] 
**CanEnrollFromSalesforce** | **bool** |  | 
**SuppressionListIds** | **[]int32** |  | 
**GoalFilterBranch** | Pointer to [**ApiContactFlowCreateRequestAllOfGoalFilterBranch**](ApiContactFlowCreateRequestAllOfGoalFilterBranch.md) |  | [optional] 
**EventAnchor** | Pointer to [**ApiContactFlowCreateRequestAllOfEventAnchor**](ApiContactFlowCreateRequestAllOfEventAnchor.md) |  | [optional] 
**UnEnrollmentSetting** | Pointer to [**ApiUnEnrollmentSetting**](ApiUnEnrollmentSetting.md) |  | [optional] 

## Methods

### NewApiFlowCreateRequest

`func NewApiFlowCreateRequest(type_ string, isEnabled bool, objectTypeId string, flowType string, actions []ApiPlatformFlowCreateRequestAllOfActions, timeWindows []ApiTimeWindow, blockedDates []ApiBlockedDate, customProperties map[string]string, dataSources []ApiPlatformFlowCreateRequestAllOfDataSources, canEnrollFromSalesforce bool, suppressionListIds []int32, ) *ApiFlowCreateRequest`

NewApiFlowCreateRequest instantiates a new ApiFlowCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFlowCreateRequestWithDefaults

`func NewApiFlowCreateRequestWithDefaults() *ApiFlowCreateRequest`

NewApiFlowCreateRequestWithDefaults instantiates a new ApiFlowCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiFlowCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiFlowCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiFlowCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetIsEnabled

`func (o *ApiFlowCreateRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiFlowCreateRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiFlowCreateRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetObjectTypeId

`func (o *ApiFlowCreateRequest) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiFlowCreateRequest) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiFlowCreateRequest) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetFlowType

`func (o *ApiFlowCreateRequest) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *ApiFlowCreateRequest) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *ApiFlowCreateRequest) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetName

`func (o *ApiFlowCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiFlowCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiFlowCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiFlowCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiFlowCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiFlowCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiFlowCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiFlowCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *ApiFlowCreateRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiFlowCreateRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiFlowCreateRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiFlowCreateRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStartActionId

`func (o *ApiFlowCreateRequest) GetStartActionId() string`

GetStartActionId returns the StartActionId field if non-nil, zero value otherwise.

### GetStartActionIdOk

`func (o *ApiFlowCreateRequest) GetStartActionIdOk() (*string, bool)`

GetStartActionIdOk returns a tuple with the StartActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActionId

`func (o *ApiFlowCreateRequest) SetStartActionId(v string)`

SetStartActionId sets StartActionId field to given value.

### HasStartActionId

`func (o *ApiFlowCreateRequest) HasStartActionId() bool`

HasStartActionId returns a boolean if a field has been set.

### GetActions

`func (o *ApiFlowCreateRequest) GetActions() []ApiPlatformFlowCreateRequestAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiFlowCreateRequest) GetActionsOk() (*[]ApiPlatformFlowCreateRequestAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiFlowCreateRequest) SetActions(v []ApiPlatformFlowCreateRequestAllOfActions)`

SetActions sets Actions field to given value.


### GetEnrollmentCriteria

`func (o *ApiFlowCreateRequest) GetEnrollmentCriteria() ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria`

GetEnrollmentCriteria returns the EnrollmentCriteria field if non-nil, zero value otherwise.

### GetEnrollmentCriteriaOk

`func (o *ApiFlowCreateRequest) GetEnrollmentCriteriaOk() (*ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria, bool)`

GetEnrollmentCriteriaOk returns a tuple with the EnrollmentCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCriteria

`func (o *ApiFlowCreateRequest) SetEnrollmentCriteria(v ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria)`

SetEnrollmentCriteria sets EnrollmentCriteria field to given value.

### HasEnrollmentCriteria

`func (o *ApiFlowCreateRequest) HasEnrollmentCriteria() bool`

HasEnrollmentCriteria returns a boolean if a field has been set.

### GetEnrollmentSchedule

`func (o *ApiFlowCreateRequest) GetEnrollmentSchedule() ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule`

GetEnrollmentSchedule returns the EnrollmentSchedule field if non-nil, zero value otherwise.

### GetEnrollmentScheduleOk

`func (o *ApiFlowCreateRequest) GetEnrollmentScheduleOk() (*ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule, bool)`

GetEnrollmentScheduleOk returns a tuple with the EnrollmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSchedule

`func (o *ApiFlowCreateRequest) SetEnrollmentSchedule(v ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule)`

SetEnrollmentSchedule sets EnrollmentSchedule field to given value.

### HasEnrollmentSchedule

`func (o *ApiFlowCreateRequest) HasEnrollmentSchedule() bool`

HasEnrollmentSchedule returns a boolean if a field has been set.

### GetTimeWindows

`func (o *ApiFlowCreateRequest) GetTimeWindows() []ApiTimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *ApiFlowCreateRequest) GetTimeWindowsOk() (*[]ApiTimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *ApiFlowCreateRequest) SetTimeWindows(v []ApiTimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetBlockedDates

`func (o *ApiFlowCreateRequest) GetBlockedDates() []ApiBlockedDate`

GetBlockedDates returns the BlockedDates field if non-nil, zero value otherwise.

### GetBlockedDatesOk

`func (o *ApiFlowCreateRequest) GetBlockedDatesOk() (*[]ApiBlockedDate, bool)`

GetBlockedDatesOk returns a tuple with the BlockedDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDates

`func (o *ApiFlowCreateRequest) SetBlockedDates(v []ApiBlockedDate)`

SetBlockedDates sets BlockedDates field to given value.


### GetCustomProperties

`func (o *ApiFlowCreateRequest) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ApiFlowCreateRequest) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ApiFlowCreateRequest) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.


### GetDataSources

`func (o *ApiFlowCreateRequest) GetDataSources() []ApiPlatformFlowCreateRequestAllOfDataSources`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ApiFlowCreateRequest) GetDataSourcesOk() (*[]ApiPlatformFlowCreateRequestAllOfDataSources, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ApiFlowCreateRequest) SetDataSources(v []ApiPlatformFlowCreateRequestAllOfDataSources)`

SetDataSources sets DataSources field to given value.


### GetSuppressionFilterBranch

`func (o *ApiFlowCreateRequest) GetSuppressionFilterBranch() ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch`

GetSuppressionFilterBranch returns the SuppressionFilterBranch field if non-nil, zero value otherwise.

### GetSuppressionFilterBranchOk

`func (o *ApiFlowCreateRequest) GetSuppressionFilterBranchOk() (*ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch, bool)`

GetSuppressionFilterBranchOk returns a tuple with the SuppressionFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionFilterBranch

`func (o *ApiFlowCreateRequest) SetSuppressionFilterBranch(v ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch)`

SetSuppressionFilterBranch sets SuppressionFilterBranch field to given value.

### HasSuppressionFilterBranch

`func (o *ApiFlowCreateRequest) HasSuppressionFilterBranch() bool`

HasSuppressionFilterBranch returns a boolean if a field has been set.

### GetCanEnrollFromSalesforce

`func (o *ApiFlowCreateRequest) GetCanEnrollFromSalesforce() bool`

GetCanEnrollFromSalesforce returns the CanEnrollFromSalesforce field if non-nil, zero value otherwise.

### GetCanEnrollFromSalesforceOk

`func (o *ApiFlowCreateRequest) GetCanEnrollFromSalesforceOk() (*bool, bool)`

GetCanEnrollFromSalesforceOk returns a tuple with the CanEnrollFromSalesforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnrollFromSalesforce

`func (o *ApiFlowCreateRequest) SetCanEnrollFromSalesforce(v bool)`

SetCanEnrollFromSalesforce sets CanEnrollFromSalesforce field to given value.


### GetSuppressionListIds

`func (o *ApiFlowCreateRequest) GetSuppressionListIds() []int32`

GetSuppressionListIds returns the SuppressionListIds field if non-nil, zero value otherwise.

### GetSuppressionListIdsOk

`func (o *ApiFlowCreateRequest) GetSuppressionListIdsOk() (*[]int32, bool)`

GetSuppressionListIdsOk returns a tuple with the SuppressionListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionListIds

`func (o *ApiFlowCreateRequest) SetSuppressionListIds(v []int32)`

SetSuppressionListIds sets SuppressionListIds field to given value.


### GetGoalFilterBranch

`func (o *ApiFlowCreateRequest) GetGoalFilterBranch() ApiContactFlowCreateRequestAllOfGoalFilterBranch`

GetGoalFilterBranch returns the GoalFilterBranch field if non-nil, zero value otherwise.

### GetGoalFilterBranchOk

`func (o *ApiFlowCreateRequest) GetGoalFilterBranchOk() (*ApiContactFlowCreateRequestAllOfGoalFilterBranch, bool)`

GetGoalFilterBranchOk returns a tuple with the GoalFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilterBranch

`func (o *ApiFlowCreateRequest) SetGoalFilterBranch(v ApiContactFlowCreateRequestAllOfGoalFilterBranch)`

SetGoalFilterBranch sets GoalFilterBranch field to given value.

### HasGoalFilterBranch

`func (o *ApiFlowCreateRequest) HasGoalFilterBranch() bool`

HasGoalFilterBranch returns a boolean if a field has been set.

### GetEventAnchor

`func (o *ApiFlowCreateRequest) GetEventAnchor() ApiContactFlowCreateRequestAllOfEventAnchor`

GetEventAnchor returns the EventAnchor field if non-nil, zero value otherwise.

### GetEventAnchorOk

`func (o *ApiFlowCreateRequest) GetEventAnchorOk() (*ApiContactFlowCreateRequestAllOfEventAnchor, bool)`

GetEventAnchorOk returns a tuple with the EventAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAnchor

`func (o *ApiFlowCreateRequest) SetEventAnchor(v ApiContactFlowCreateRequestAllOfEventAnchor)`

SetEventAnchor sets EventAnchor field to given value.

### HasEventAnchor

`func (o *ApiFlowCreateRequest) HasEventAnchor() bool`

HasEventAnchor returns a boolean if a field has been set.

### GetUnEnrollmentSetting

`func (o *ApiFlowCreateRequest) GetUnEnrollmentSetting() ApiUnEnrollmentSetting`

GetUnEnrollmentSetting returns the UnEnrollmentSetting field if non-nil, zero value otherwise.

### GetUnEnrollmentSettingOk

`func (o *ApiFlowCreateRequest) GetUnEnrollmentSettingOk() (*ApiUnEnrollmentSetting, bool)`

GetUnEnrollmentSettingOk returns a tuple with the UnEnrollmentSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollmentSetting

`func (o *ApiFlowCreateRequest) SetUnEnrollmentSetting(v ApiUnEnrollmentSetting)`

SetUnEnrollmentSetting sets UnEnrollmentSetting field to given value.

### HasUnEnrollmentSetting

`func (o *ApiFlowCreateRequest) HasUnEnrollmentSetting() bool`

HasUnEnrollmentSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


