# ApiFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "CONTACT_FLOW"]
**Id** | **string** |  | 
**IsEnabled** | **bool** |  | 
**ObjectTypeId** | **string** |  | 
**FlowType** | **string** |  | 
**RevisionId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StartActionId** | Pointer to **string** |  | [optional] 
**NextAvailableActionId** | **string** |  | 
**Actions** | [**[]ApiPlatformFlowAllOfActions**](ApiPlatformFlowAllOfActions.md) |  | 
**EnrollmentCriteria** | Pointer to [**ApiPlatformFlowAllOfEnrollmentCriteria**](ApiPlatformFlowAllOfEnrollmentCriteria.md) |  | [optional] 
**EnrollmentSchedule** | Pointer to [**ApiPlatformFlowAllOfEnrollmentSchedule**](ApiPlatformFlowAllOfEnrollmentSchedule.md) |  | [optional] 
**TimeWindows** | [**[]ApiTimeWindow**](ApiTimeWindow.md) |  | 
**BlockedDates** | [**[]ApiBlockedDate**](ApiBlockedDate.md) |  | 
**CustomProperties** | **map[string]string** |  | 
**DataSources** | [**[]ApiPlatformFlowAllOfDataSources**](ApiPlatformFlowAllOfDataSources.md) |  | 
**CrmObjectCreationStatus** | **string** |  | 
**SuppressionFilterBranch** | Pointer to [**ApiPlatformFlowAllOfSuppressionFilterBranch**](ApiPlatformFlowAllOfSuppressionFilterBranch.md) |  | [optional] 
**SuppressionListIds** | **[]int32** |  | 
**GoalFilterBranch** | Pointer to [**ApiContactFlowAllOfGoalFilterBranch**](ApiContactFlowAllOfGoalFilterBranch.md) |  | [optional] 
**CanEnrollFromSalesforce** | **bool** |  | 
**EventAnchor** | Pointer to [**ApiContactFlowAllOfEventAnchor**](ApiContactFlowAllOfEventAnchor.md) |  | [optional] 
**UnEnrollmentSetting** | Pointer to [**ApiUnEnrollmentSetting**](ApiUnEnrollmentSetting.md) |  | [optional] 

## Methods

### NewApiFlow

`func NewApiFlow(type_ string, id string, isEnabled bool, objectTypeId string, flowType string, revisionId string, createdAt time.Time, updatedAt time.Time, nextAvailableActionId string, actions []ApiPlatformFlowAllOfActions, timeWindows []ApiTimeWindow, blockedDates []ApiBlockedDate, customProperties map[string]string, dataSources []ApiPlatformFlowAllOfDataSources, crmObjectCreationStatus string, suppressionListIds []int32, canEnrollFromSalesforce bool, ) *ApiFlow`

NewApiFlow instantiates a new ApiFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFlowWithDefaults

`func NewApiFlowWithDefaults() *ApiFlow`

NewApiFlowWithDefaults instantiates a new ApiFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiFlow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiFlow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiFlow) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ApiFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIsEnabled

`func (o *ApiFlow) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiFlow) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiFlow) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetObjectTypeId

`func (o *ApiFlow) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiFlow) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiFlow) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetFlowType

`func (o *ApiFlow) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *ApiFlow) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *ApiFlow) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetRevisionId

`func (o *ApiFlow) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ApiFlow) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ApiFlow) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetName

`func (o *ApiFlow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiFlow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiFlow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiFlow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiFlow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiFlow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiFlow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiFlow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *ApiFlow) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiFlow) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiFlow) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiFlow) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ApiFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartActionId

`func (o *ApiFlow) GetStartActionId() string`

GetStartActionId returns the StartActionId field if non-nil, zero value otherwise.

### GetStartActionIdOk

`func (o *ApiFlow) GetStartActionIdOk() (*string, bool)`

GetStartActionIdOk returns a tuple with the StartActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActionId

`func (o *ApiFlow) SetStartActionId(v string)`

SetStartActionId sets StartActionId field to given value.

### HasStartActionId

`func (o *ApiFlow) HasStartActionId() bool`

HasStartActionId returns a boolean if a field has been set.

### GetNextAvailableActionId

`func (o *ApiFlow) GetNextAvailableActionId() string`

GetNextAvailableActionId returns the NextAvailableActionId field if non-nil, zero value otherwise.

### GetNextAvailableActionIdOk

`func (o *ApiFlow) GetNextAvailableActionIdOk() (*string, bool)`

GetNextAvailableActionIdOk returns a tuple with the NextAvailableActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableActionId

`func (o *ApiFlow) SetNextAvailableActionId(v string)`

SetNextAvailableActionId sets NextAvailableActionId field to given value.


### GetActions

`func (o *ApiFlow) GetActions() []ApiPlatformFlowAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiFlow) GetActionsOk() (*[]ApiPlatformFlowAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiFlow) SetActions(v []ApiPlatformFlowAllOfActions)`

SetActions sets Actions field to given value.


### GetEnrollmentCriteria

`func (o *ApiFlow) GetEnrollmentCriteria() ApiPlatformFlowAllOfEnrollmentCriteria`

GetEnrollmentCriteria returns the EnrollmentCriteria field if non-nil, zero value otherwise.

### GetEnrollmentCriteriaOk

`func (o *ApiFlow) GetEnrollmentCriteriaOk() (*ApiPlatformFlowAllOfEnrollmentCriteria, bool)`

GetEnrollmentCriteriaOk returns a tuple with the EnrollmentCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCriteria

`func (o *ApiFlow) SetEnrollmentCriteria(v ApiPlatformFlowAllOfEnrollmentCriteria)`

SetEnrollmentCriteria sets EnrollmentCriteria field to given value.

### HasEnrollmentCriteria

`func (o *ApiFlow) HasEnrollmentCriteria() bool`

HasEnrollmentCriteria returns a boolean if a field has been set.

### GetEnrollmentSchedule

`func (o *ApiFlow) GetEnrollmentSchedule() ApiPlatformFlowAllOfEnrollmentSchedule`

GetEnrollmentSchedule returns the EnrollmentSchedule field if non-nil, zero value otherwise.

### GetEnrollmentScheduleOk

`func (o *ApiFlow) GetEnrollmentScheduleOk() (*ApiPlatformFlowAllOfEnrollmentSchedule, bool)`

GetEnrollmentScheduleOk returns a tuple with the EnrollmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSchedule

`func (o *ApiFlow) SetEnrollmentSchedule(v ApiPlatformFlowAllOfEnrollmentSchedule)`

SetEnrollmentSchedule sets EnrollmentSchedule field to given value.

### HasEnrollmentSchedule

`func (o *ApiFlow) HasEnrollmentSchedule() bool`

HasEnrollmentSchedule returns a boolean if a field has been set.

### GetTimeWindows

`func (o *ApiFlow) GetTimeWindows() []ApiTimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *ApiFlow) GetTimeWindowsOk() (*[]ApiTimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *ApiFlow) SetTimeWindows(v []ApiTimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetBlockedDates

`func (o *ApiFlow) GetBlockedDates() []ApiBlockedDate`

GetBlockedDates returns the BlockedDates field if non-nil, zero value otherwise.

### GetBlockedDatesOk

`func (o *ApiFlow) GetBlockedDatesOk() (*[]ApiBlockedDate, bool)`

GetBlockedDatesOk returns a tuple with the BlockedDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDates

`func (o *ApiFlow) SetBlockedDates(v []ApiBlockedDate)`

SetBlockedDates sets BlockedDates field to given value.


### GetCustomProperties

`func (o *ApiFlow) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ApiFlow) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ApiFlow) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.


### GetDataSources

`func (o *ApiFlow) GetDataSources() []ApiPlatformFlowAllOfDataSources`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ApiFlow) GetDataSourcesOk() (*[]ApiPlatformFlowAllOfDataSources, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ApiFlow) SetDataSources(v []ApiPlatformFlowAllOfDataSources)`

SetDataSources sets DataSources field to given value.


### GetCrmObjectCreationStatus

`func (o *ApiFlow) GetCrmObjectCreationStatus() string`

GetCrmObjectCreationStatus returns the CrmObjectCreationStatus field if non-nil, zero value otherwise.

### GetCrmObjectCreationStatusOk

`func (o *ApiFlow) GetCrmObjectCreationStatusOk() (*string, bool)`

GetCrmObjectCreationStatusOk returns a tuple with the CrmObjectCreationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmObjectCreationStatus

`func (o *ApiFlow) SetCrmObjectCreationStatus(v string)`

SetCrmObjectCreationStatus sets CrmObjectCreationStatus field to given value.


### GetSuppressionFilterBranch

`func (o *ApiFlow) GetSuppressionFilterBranch() ApiPlatformFlowAllOfSuppressionFilterBranch`

GetSuppressionFilterBranch returns the SuppressionFilterBranch field if non-nil, zero value otherwise.

### GetSuppressionFilterBranchOk

`func (o *ApiFlow) GetSuppressionFilterBranchOk() (*ApiPlatformFlowAllOfSuppressionFilterBranch, bool)`

GetSuppressionFilterBranchOk returns a tuple with the SuppressionFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionFilterBranch

`func (o *ApiFlow) SetSuppressionFilterBranch(v ApiPlatformFlowAllOfSuppressionFilterBranch)`

SetSuppressionFilterBranch sets SuppressionFilterBranch field to given value.

### HasSuppressionFilterBranch

`func (o *ApiFlow) HasSuppressionFilterBranch() bool`

HasSuppressionFilterBranch returns a boolean if a field has been set.

### GetSuppressionListIds

`func (o *ApiFlow) GetSuppressionListIds() []int32`

GetSuppressionListIds returns the SuppressionListIds field if non-nil, zero value otherwise.

### GetSuppressionListIdsOk

`func (o *ApiFlow) GetSuppressionListIdsOk() (*[]int32, bool)`

GetSuppressionListIdsOk returns a tuple with the SuppressionListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionListIds

`func (o *ApiFlow) SetSuppressionListIds(v []int32)`

SetSuppressionListIds sets SuppressionListIds field to given value.


### GetGoalFilterBranch

`func (o *ApiFlow) GetGoalFilterBranch() ApiContactFlowAllOfGoalFilterBranch`

GetGoalFilterBranch returns the GoalFilterBranch field if non-nil, zero value otherwise.

### GetGoalFilterBranchOk

`func (o *ApiFlow) GetGoalFilterBranchOk() (*ApiContactFlowAllOfGoalFilterBranch, bool)`

GetGoalFilterBranchOk returns a tuple with the GoalFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilterBranch

`func (o *ApiFlow) SetGoalFilterBranch(v ApiContactFlowAllOfGoalFilterBranch)`

SetGoalFilterBranch sets GoalFilterBranch field to given value.

### HasGoalFilterBranch

`func (o *ApiFlow) HasGoalFilterBranch() bool`

HasGoalFilterBranch returns a boolean if a field has been set.

### GetCanEnrollFromSalesforce

`func (o *ApiFlow) GetCanEnrollFromSalesforce() bool`

GetCanEnrollFromSalesforce returns the CanEnrollFromSalesforce field if non-nil, zero value otherwise.

### GetCanEnrollFromSalesforceOk

`func (o *ApiFlow) GetCanEnrollFromSalesforceOk() (*bool, bool)`

GetCanEnrollFromSalesforceOk returns a tuple with the CanEnrollFromSalesforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnrollFromSalesforce

`func (o *ApiFlow) SetCanEnrollFromSalesforce(v bool)`

SetCanEnrollFromSalesforce sets CanEnrollFromSalesforce field to given value.


### GetEventAnchor

`func (o *ApiFlow) GetEventAnchor() ApiContactFlowAllOfEventAnchor`

GetEventAnchor returns the EventAnchor field if non-nil, zero value otherwise.

### GetEventAnchorOk

`func (o *ApiFlow) GetEventAnchorOk() (*ApiContactFlowAllOfEventAnchor, bool)`

GetEventAnchorOk returns a tuple with the EventAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAnchor

`func (o *ApiFlow) SetEventAnchor(v ApiContactFlowAllOfEventAnchor)`

SetEventAnchor sets EventAnchor field to given value.

### HasEventAnchor

`func (o *ApiFlow) HasEventAnchor() bool`

HasEventAnchor returns a boolean if a field has been set.

### GetUnEnrollmentSetting

`func (o *ApiFlow) GetUnEnrollmentSetting() ApiUnEnrollmentSetting`

GetUnEnrollmentSetting returns the UnEnrollmentSetting field if non-nil, zero value otherwise.

### GetUnEnrollmentSettingOk

`func (o *ApiFlow) GetUnEnrollmentSettingOk() (*ApiUnEnrollmentSetting, bool)`

GetUnEnrollmentSettingOk returns a tuple with the UnEnrollmentSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollmentSetting

`func (o *ApiFlow) SetUnEnrollmentSetting(v ApiUnEnrollmentSetting)`

SetUnEnrollmentSetting sets UnEnrollmentSetting field to given value.

### HasUnEnrollmentSetting

`func (o *ApiFlow) HasUnEnrollmentSetting() bool`

HasUnEnrollmentSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


