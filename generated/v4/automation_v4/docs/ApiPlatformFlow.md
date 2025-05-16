# ApiPlatformFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "PLATFORM_FLOW"]
**Id** | **string** |  | 
**IsEnabled** | **bool** |  | 
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
**SuppressionListIds** | **[]int32** |  | 
**GoalFilterBranch** | Pointer to [**ApiContactFlowAllOfGoalFilterBranch**](ApiContactFlowAllOfGoalFilterBranch.md) |  | [optional] 
**CanEnrollFromSalesforce** | **bool** |  | 
**EventAnchor** | Pointer to [**ApiContactFlowAllOfEventAnchor**](ApiContactFlowAllOfEventAnchor.md) |  | [optional] 
**UnEnrollmentSetting** | Pointer to [**ApiUnEnrollmentSetting**](ApiUnEnrollmentSetting.md) |  | [optional] 
**ObjectTypeId** | **string** |  | 
**SuppressionFilterBranch** | Pointer to [**ApiPlatformFlowAllOfSuppressionFilterBranch**](ApiPlatformFlowAllOfSuppressionFilterBranch.md) |  | [optional] 

## Methods

### NewApiPlatformFlow

`func NewApiPlatformFlow(type_ string, id string, isEnabled bool, flowType string, revisionId string, createdAt time.Time, updatedAt time.Time, nextAvailableActionId string, actions []ApiPlatformFlowAllOfActions, timeWindows []ApiTimeWindow, blockedDates []ApiBlockedDate, customProperties map[string]string, dataSources []ApiPlatformFlowAllOfDataSources, crmObjectCreationStatus string, suppressionListIds []int32, canEnrollFromSalesforce bool, objectTypeId string, ) *ApiPlatformFlow`

NewApiPlatformFlow instantiates a new ApiPlatformFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPlatformFlowWithDefaults

`func NewApiPlatformFlowWithDefaults() *ApiPlatformFlow`

NewApiPlatformFlowWithDefaults instantiates a new ApiPlatformFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiPlatformFlow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPlatformFlow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPlatformFlow) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ApiPlatformFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPlatformFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPlatformFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIsEnabled

`func (o *ApiPlatformFlow) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiPlatformFlow) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiPlatformFlow) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetFlowType

`func (o *ApiPlatformFlow) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *ApiPlatformFlow) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *ApiPlatformFlow) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetRevisionId

`func (o *ApiPlatformFlow) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ApiPlatformFlow) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ApiPlatformFlow) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetName

`func (o *ApiPlatformFlow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPlatformFlow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPlatformFlow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPlatformFlow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiPlatformFlow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPlatformFlow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPlatformFlow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPlatformFlow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *ApiPlatformFlow) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiPlatformFlow) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiPlatformFlow) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiPlatformFlow) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiPlatformFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiPlatformFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiPlatformFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ApiPlatformFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiPlatformFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiPlatformFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartActionId

`func (o *ApiPlatformFlow) GetStartActionId() string`

GetStartActionId returns the StartActionId field if non-nil, zero value otherwise.

### GetStartActionIdOk

`func (o *ApiPlatformFlow) GetStartActionIdOk() (*string, bool)`

GetStartActionIdOk returns a tuple with the StartActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActionId

`func (o *ApiPlatformFlow) SetStartActionId(v string)`

SetStartActionId sets StartActionId field to given value.

### HasStartActionId

`func (o *ApiPlatformFlow) HasStartActionId() bool`

HasStartActionId returns a boolean if a field has been set.

### GetNextAvailableActionId

`func (o *ApiPlatformFlow) GetNextAvailableActionId() string`

GetNextAvailableActionId returns the NextAvailableActionId field if non-nil, zero value otherwise.

### GetNextAvailableActionIdOk

`func (o *ApiPlatformFlow) GetNextAvailableActionIdOk() (*string, bool)`

GetNextAvailableActionIdOk returns a tuple with the NextAvailableActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableActionId

`func (o *ApiPlatformFlow) SetNextAvailableActionId(v string)`

SetNextAvailableActionId sets NextAvailableActionId field to given value.


### GetActions

`func (o *ApiPlatformFlow) GetActions() []ApiPlatformFlowAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPlatformFlow) GetActionsOk() (*[]ApiPlatformFlowAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPlatformFlow) SetActions(v []ApiPlatformFlowAllOfActions)`

SetActions sets Actions field to given value.


### GetEnrollmentCriteria

`func (o *ApiPlatformFlow) GetEnrollmentCriteria() ApiPlatformFlowAllOfEnrollmentCriteria`

GetEnrollmentCriteria returns the EnrollmentCriteria field if non-nil, zero value otherwise.

### GetEnrollmentCriteriaOk

`func (o *ApiPlatformFlow) GetEnrollmentCriteriaOk() (*ApiPlatformFlowAllOfEnrollmentCriteria, bool)`

GetEnrollmentCriteriaOk returns a tuple with the EnrollmentCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCriteria

`func (o *ApiPlatformFlow) SetEnrollmentCriteria(v ApiPlatformFlowAllOfEnrollmentCriteria)`

SetEnrollmentCriteria sets EnrollmentCriteria field to given value.

### HasEnrollmentCriteria

`func (o *ApiPlatformFlow) HasEnrollmentCriteria() bool`

HasEnrollmentCriteria returns a boolean if a field has been set.

### GetEnrollmentSchedule

`func (o *ApiPlatformFlow) GetEnrollmentSchedule() ApiPlatformFlowAllOfEnrollmentSchedule`

GetEnrollmentSchedule returns the EnrollmentSchedule field if non-nil, zero value otherwise.

### GetEnrollmentScheduleOk

`func (o *ApiPlatformFlow) GetEnrollmentScheduleOk() (*ApiPlatformFlowAllOfEnrollmentSchedule, bool)`

GetEnrollmentScheduleOk returns a tuple with the EnrollmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSchedule

`func (o *ApiPlatformFlow) SetEnrollmentSchedule(v ApiPlatformFlowAllOfEnrollmentSchedule)`

SetEnrollmentSchedule sets EnrollmentSchedule field to given value.

### HasEnrollmentSchedule

`func (o *ApiPlatformFlow) HasEnrollmentSchedule() bool`

HasEnrollmentSchedule returns a boolean if a field has been set.

### GetTimeWindows

`func (o *ApiPlatformFlow) GetTimeWindows() []ApiTimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *ApiPlatformFlow) GetTimeWindowsOk() (*[]ApiTimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *ApiPlatformFlow) SetTimeWindows(v []ApiTimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetBlockedDates

`func (o *ApiPlatformFlow) GetBlockedDates() []ApiBlockedDate`

GetBlockedDates returns the BlockedDates field if non-nil, zero value otherwise.

### GetBlockedDatesOk

`func (o *ApiPlatformFlow) GetBlockedDatesOk() (*[]ApiBlockedDate, bool)`

GetBlockedDatesOk returns a tuple with the BlockedDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDates

`func (o *ApiPlatformFlow) SetBlockedDates(v []ApiBlockedDate)`

SetBlockedDates sets BlockedDates field to given value.


### GetCustomProperties

`func (o *ApiPlatformFlow) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ApiPlatformFlow) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ApiPlatformFlow) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.


### GetDataSources

`func (o *ApiPlatformFlow) GetDataSources() []ApiPlatformFlowAllOfDataSources`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ApiPlatformFlow) GetDataSourcesOk() (*[]ApiPlatformFlowAllOfDataSources, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ApiPlatformFlow) SetDataSources(v []ApiPlatformFlowAllOfDataSources)`

SetDataSources sets DataSources field to given value.


### GetCrmObjectCreationStatus

`func (o *ApiPlatformFlow) GetCrmObjectCreationStatus() string`

GetCrmObjectCreationStatus returns the CrmObjectCreationStatus field if non-nil, zero value otherwise.

### GetCrmObjectCreationStatusOk

`func (o *ApiPlatformFlow) GetCrmObjectCreationStatusOk() (*string, bool)`

GetCrmObjectCreationStatusOk returns a tuple with the CrmObjectCreationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmObjectCreationStatus

`func (o *ApiPlatformFlow) SetCrmObjectCreationStatus(v string)`

SetCrmObjectCreationStatus sets CrmObjectCreationStatus field to given value.


### GetSuppressionListIds

`func (o *ApiPlatformFlow) GetSuppressionListIds() []int32`

GetSuppressionListIds returns the SuppressionListIds field if non-nil, zero value otherwise.

### GetSuppressionListIdsOk

`func (o *ApiPlatformFlow) GetSuppressionListIdsOk() (*[]int32, bool)`

GetSuppressionListIdsOk returns a tuple with the SuppressionListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionListIds

`func (o *ApiPlatformFlow) SetSuppressionListIds(v []int32)`

SetSuppressionListIds sets SuppressionListIds field to given value.


### GetGoalFilterBranch

`func (o *ApiPlatformFlow) GetGoalFilterBranch() ApiContactFlowAllOfGoalFilterBranch`

GetGoalFilterBranch returns the GoalFilterBranch field if non-nil, zero value otherwise.

### GetGoalFilterBranchOk

`func (o *ApiPlatformFlow) GetGoalFilterBranchOk() (*ApiContactFlowAllOfGoalFilterBranch, bool)`

GetGoalFilterBranchOk returns a tuple with the GoalFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilterBranch

`func (o *ApiPlatformFlow) SetGoalFilterBranch(v ApiContactFlowAllOfGoalFilterBranch)`

SetGoalFilterBranch sets GoalFilterBranch field to given value.

### HasGoalFilterBranch

`func (o *ApiPlatformFlow) HasGoalFilterBranch() bool`

HasGoalFilterBranch returns a boolean if a field has been set.

### GetCanEnrollFromSalesforce

`func (o *ApiPlatformFlow) GetCanEnrollFromSalesforce() bool`

GetCanEnrollFromSalesforce returns the CanEnrollFromSalesforce field if non-nil, zero value otherwise.

### GetCanEnrollFromSalesforceOk

`func (o *ApiPlatformFlow) GetCanEnrollFromSalesforceOk() (*bool, bool)`

GetCanEnrollFromSalesforceOk returns a tuple with the CanEnrollFromSalesforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnrollFromSalesforce

`func (o *ApiPlatformFlow) SetCanEnrollFromSalesforce(v bool)`

SetCanEnrollFromSalesforce sets CanEnrollFromSalesforce field to given value.


### GetEventAnchor

`func (o *ApiPlatformFlow) GetEventAnchor() ApiContactFlowAllOfEventAnchor`

GetEventAnchor returns the EventAnchor field if non-nil, zero value otherwise.

### GetEventAnchorOk

`func (o *ApiPlatformFlow) GetEventAnchorOk() (*ApiContactFlowAllOfEventAnchor, bool)`

GetEventAnchorOk returns a tuple with the EventAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAnchor

`func (o *ApiPlatformFlow) SetEventAnchor(v ApiContactFlowAllOfEventAnchor)`

SetEventAnchor sets EventAnchor field to given value.

### HasEventAnchor

`func (o *ApiPlatformFlow) HasEventAnchor() bool`

HasEventAnchor returns a boolean if a field has been set.

### GetUnEnrollmentSetting

`func (o *ApiPlatformFlow) GetUnEnrollmentSetting() ApiUnEnrollmentSetting`

GetUnEnrollmentSetting returns the UnEnrollmentSetting field if non-nil, zero value otherwise.

### GetUnEnrollmentSettingOk

`func (o *ApiPlatformFlow) GetUnEnrollmentSettingOk() (*ApiUnEnrollmentSetting, bool)`

GetUnEnrollmentSettingOk returns a tuple with the UnEnrollmentSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollmentSetting

`func (o *ApiPlatformFlow) SetUnEnrollmentSetting(v ApiUnEnrollmentSetting)`

SetUnEnrollmentSetting sets UnEnrollmentSetting field to given value.

### HasUnEnrollmentSetting

`func (o *ApiPlatformFlow) HasUnEnrollmentSetting() bool`

HasUnEnrollmentSetting returns a boolean if a field has been set.

### GetObjectTypeId

`func (o *ApiPlatformFlow) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiPlatformFlow) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiPlatformFlow) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetSuppressionFilterBranch

`func (o *ApiPlatformFlow) GetSuppressionFilterBranch() ApiPlatformFlowAllOfSuppressionFilterBranch`

GetSuppressionFilterBranch returns the SuppressionFilterBranch field if non-nil, zero value otherwise.

### GetSuppressionFilterBranchOk

`func (o *ApiPlatformFlow) GetSuppressionFilterBranchOk() (*ApiPlatformFlowAllOfSuppressionFilterBranch, bool)`

GetSuppressionFilterBranchOk returns a tuple with the SuppressionFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionFilterBranch

`func (o *ApiPlatformFlow) SetSuppressionFilterBranch(v ApiPlatformFlowAllOfSuppressionFilterBranch)`

SetSuppressionFilterBranch sets SuppressionFilterBranch field to given value.

### HasSuppressionFilterBranch

`func (o *ApiPlatformFlow) HasSuppressionFilterBranch() bool`

HasSuppressionFilterBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


