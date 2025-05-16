# ApiContactFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "CONTACT_FLOW"]
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
**Actions** | [**[]ApiContactFlowAllOfActions**](ApiContactFlowAllOfActions.md) |  | 
**EnrollmentCriteria** | Pointer to [**ApiContactFlowAllOfEnrollmentCriteria**](ApiContactFlowAllOfEnrollmentCriteria.md) |  | [optional] 
**EnrollmentSchedule** | Pointer to [**ApiContactFlowAllOfEnrollmentSchedule**](ApiContactFlowAllOfEnrollmentSchedule.md) |  | [optional] 
**TimeWindows** | [**[]ApiTimeWindow**](ApiTimeWindow.md) |  | 
**BlockedDates** | [**[]ApiBlockedDate**](ApiBlockedDate.md) |  | 
**CustomProperties** | **map[string]string** |  | 
**DataSources** | [**[]ApiContactFlowAllOfDataSources**](ApiContactFlowAllOfDataSources.md) |  | 
**CrmObjectCreationStatus** | **string** |  | 
**SuppressionListIds** | **[]int32** |  | 
**GoalFilterBranch** | Pointer to [**ApiContactFlowAllOfGoalFilterBranch**](ApiContactFlowAllOfGoalFilterBranch.md) |  | [optional] 
**CanEnrollFromSalesforce** | **bool** |  | 
**EventAnchor** | Pointer to [**ApiContactFlowAllOfEventAnchor**](ApiContactFlowAllOfEventAnchor.md) |  | [optional] 
**UnEnrollmentSetting** | Pointer to [**ApiUnEnrollmentSetting**](ApiUnEnrollmentSetting.md) |  | [optional] 
**ObjectTypeId** | **string** |  | 
**SuppressionFilterBranch** | Pointer to [**ApiPlatformFlowAllOfSuppressionFilterBranch**](ApiPlatformFlowAllOfSuppressionFilterBranch.md) |  | [optional] 

## Methods

### NewApiContactFlow

`func NewApiContactFlow(type_ string, id string, isEnabled bool, flowType string, revisionId string, createdAt time.Time, updatedAt time.Time, nextAvailableActionId string, actions []ApiContactFlowAllOfActions, timeWindows []ApiTimeWindow, blockedDates []ApiBlockedDate, customProperties map[string]string, dataSources []ApiContactFlowAllOfDataSources, crmObjectCreationStatus string, suppressionListIds []int32, canEnrollFromSalesforce bool, objectTypeId string, ) *ApiContactFlow`

NewApiContactFlow instantiates a new ApiContactFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactFlowWithDefaults

`func NewApiContactFlowWithDefaults() *ApiContactFlow`

NewApiContactFlowWithDefaults instantiates a new ApiContactFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiContactFlow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactFlow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactFlow) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ApiContactFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiContactFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiContactFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIsEnabled

`func (o *ApiContactFlow) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiContactFlow) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiContactFlow) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetFlowType

`func (o *ApiContactFlow) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *ApiContactFlow) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *ApiContactFlow) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetRevisionId

`func (o *ApiContactFlow) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ApiContactFlow) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ApiContactFlow) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetName

`func (o *ApiContactFlow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiContactFlow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiContactFlow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiContactFlow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiContactFlow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiContactFlow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiContactFlow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiContactFlow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *ApiContactFlow) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiContactFlow) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiContactFlow) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiContactFlow) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiContactFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiContactFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiContactFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ApiContactFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiContactFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiContactFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartActionId

`func (o *ApiContactFlow) GetStartActionId() string`

GetStartActionId returns the StartActionId field if non-nil, zero value otherwise.

### GetStartActionIdOk

`func (o *ApiContactFlow) GetStartActionIdOk() (*string, bool)`

GetStartActionIdOk returns a tuple with the StartActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActionId

`func (o *ApiContactFlow) SetStartActionId(v string)`

SetStartActionId sets StartActionId field to given value.

### HasStartActionId

`func (o *ApiContactFlow) HasStartActionId() bool`

HasStartActionId returns a boolean if a field has been set.

### GetNextAvailableActionId

`func (o *ApiContactFlow) GetNextAvailableActionId() string`

GetNextAvailableActionId returns the NextAvailableActionId field if non-nil, zero value otherwise.

### GetNextAvailableActionIdOk

`func (o *ApiContactFlow) GetNextAvailableActionIdOk() (*string, bool)`

GetNextAvailableActionIdOk returns a tuple with the NextAvailableActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableActionId

`func (o *ApiContactFlow) SetNextAvailableActionId(v string)`

SetNextAvailableActionId sets NextAvailableActionId field to given value.


### GetActions

`func (o *ApiContactFlow) GetActions() []ApiContactFlowAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiContactFlow) GetActionsOk() (*[]ApiContactFlowAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiContactFlow) SetActions(v []ApiContactFlowAllOfActions)`

SetActions sets Actions field to given value.


### GetEnrollmentCriteria

`func (o *ApiContactFlow) GetEnrollmentCriteria() ApiContactFlowAllOfEnrollmentCriteria`

GetEnrollmentCriteria returns the EnrollmentCriteria field if non-nil, zero value otherwise.

### GetEnrollmentCriteriaOk

`func (o *ApiContactFlow) GetEnrollmentCriteriaOk() (*ApiContactFlowAllOfEnrollmentCriteria, bool)`

GetEnrollmentCriteriaOk returns a tuple with the EnrollmentCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCriteria

`func (o *ApiContactFlow) SetEnrollmentCriteria(v ApiContactFlowAllOfEnrollmentCriteria)`

SetEnrollmentCriteria sets EnrollmentCriteria field to given value.

### HasEnrollmentCriteria

`func (o *ApiContactFlow) HasEnrollmentCriteria() bool`

HasEnrollmentCriteria returns a boolean if a field has been set.

### GetEnrollmentSchedule

`func (o *ApiContactFlow) GetEnrollmentSchedule() ApiContactFlowAllOfEnrollmentSchedule`

GetEnrollmentSchedule returns the EnrollmentSchedule field if non-nil, zero value otherwise.

### GetEnrollmentScheduleOk

`func (o *ApiContactFlow) GetEnrollmentScheduleOk() (*ApiContactFlowAllOfEnrollmentSchedule, bool)`

GetEnrollmentScheduleOk returns a tuple with the EnrollmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSchedule

`func (o *ApiContactFlow) SetEnrollmentSchedule(v ApiContactFlowAllOfEnrollmentSchedule)`

SetEnrollmentSchedule sets EnrollmentSchedule field to given value.

### HasEnrollmentSchedule

`func (o *ApiContactFlow) HasEnrollmentSchedule() bool`

HasEnrollmentSchedule returns a boolean if a field has been set.

### GetTimeWindows

`func (o *ApiContactFlow) GetTimeWindows() []ApiTimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *ApiContactFlow) GetTimeWindowsOk() (*[]ApiTimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *ApiContactFlow) SetTimeWindows(v []ApiTimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetBlockedDates

`func (o *ApiContactFlow) GetBlockedDates() []ApiBlockedDate`

GetBlockedDates returns the BlockedDates field if non-nil, zero value otherwise.

### GetBlockedDatesOk

`func (o *ApiContactFlow) GetBlockedDatesOk() (*[]ApiBlockedDate, bool)`

GetBlockedDatesOk returns a tuple with the BlockedDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDates

`func (o *ApiContactFlow) SetBlockedDates(v []ApiBlockedDate)`

SetBlockedDates sets BlockedDates field to given value.


### GetCustomProperties

`func (o *ApiContactFlow) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ApiContactFlow) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ApiContactFlow) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.


### GetDataSources

`func (o *ApiContactFlow) GetDataSources() []ApiContactFlowAllOfDataSources`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ApiContactFlow) GetDataSourcesOk() (*[]ApiContactFlowAllOfDataSources, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ApiContactFlow) SetDataSources(v []ApiContactFlowAllOfDataSources)`

SetDataSources sets DataSources field to given value.


### GetCrmObjectCreationStatus

`func (o *ApiContactFlow) GetCrmObjectCreationStatus() string`

GetCrmObjectCreationStatus returns the CrmObjectCreationStatus field if non-nil, zero value otherwise.

### GetCrmObjectCreationStatusOk

`func (o *ApiContactFlow) GetCrmObjectCreationStatusOk() (*string, bool)`

GetCrmObjectCreationStatusOk returns a tuple with the CrmObjectCreationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmObjectCreationStatus

`func (o *ApiContactFlow) SetCrmObjectCreationStatus(v string)`

SetCrmObjectCreationStatus sets CrmObjectCreationStatus field to given value.


### GetSuppressionListIds

`func (o *ApiContactFlow) GetSuppressionListIds() []int32`

GetSuppressionListIds returns the SuppressionListIds field if non-nil, zero value otherwise.

### GetSuppressionListIdsOk

`func (o *ApiContactFlow) GetSuppressionListIdsOk() (*[]int32, bool)`

GetSuppressionListIdsOk returns a tuple with the SuppressionListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionListIds

`func (o *ApiContactFlow) SetSuppressionListIds(v []int32)`

SetSuppressionListIds sets SuppressionListIds field to given value.


### GetGoalFilterBranch

`func (o *ApiContactFlow) GetGoalFilterBranch() ApiContactFlowAllOfGoalFilterBranch`

GetGoalFilterBranch returns the GoalFilterBranch field if non-nil, zero value otherwise.

### GetGoalFilterBranchOk

`func (o *ApiContactFlow) GetGoalFilterBranchOk() (*ApiContactFlowAllOfGoalFilterBranch, bool)`

GetGoalFilterBranchOk returns a tuple with the GoalFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilterBranch

`func (o *ApiContactFlow) SetGoalFilterBranch(v ApiContactFlowAllOfGoalFilterBranch)`

SetGoalFilterBranch sets GoalFilterBranch field to given value.

### HasGoalFilterBranch

`func (o *ApiContactFlow) HasGoalFilterBranch() bool`

HasGoalFilterBranch returns a boolean if a field has been set.

### GetCanEnrollFromSalesforce

`func (o *ApiContactFlow) GetCanEnrollFromSalesforce() bool`

GetCanEnrollFromSalesforce returns the CanEnrollFromSalesforce field if non-nil, zero value otherwise.

### GetCanEnrollFromSalesforceOk

`func (o *ApiContactFlow) GetCanEnrollFromSalesforceOk() (*bool, bool)`

GetCanEnrollFromSalesforceOk returns a tuple with the CanEnrollFromSalesforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnrollFromSalesforce

`func (o *ApiContactFlow) SetCanEnrollFromSalesforce(v bool)`

SetCanEnrollFromSalesforce sets CanEnrollFromSalesforce field to given value.


### GetEventAnchor

`func (o *ApiContactFlow) GetEventAnchor() ApiContactFlowAllOfEventAnchor`

GetEventAnchor returns the EventAnchor field if non-nil, zero value otherwise.

### GetEventAnchorOk

`func (o *ApiContactFlow) GetEventAnchorOk() (*ApiContactFlowAllOfEventAnchor, bool)`

GetEventAnchorOk returns a tuple with the EventAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAnchor

`func (o *ApiContactFlow) SetEventAnchor(v ApiContactFlowAllOfEventAnchor)`

SetEventAnchor sets EventAnchor field to given value.

### HasEventAnchor

`func (o *ApiContactFlow) HasEventAnchor() bool`

HasEventAnchor returns a boolean if a field has been set.

### GetUnEnrollmentSetting

`func (o *ApiContactFlow) GetUnEnrollmentSetting() ApiUnEnrollmentSetting`

GetUnEnrollmentSetting returns the UnEnrollmentSetting field if non-nil, zero value otherwise.

### GetUnEnrollmentSettingOk

`func (o *ApiContactFlow) GetUnEnrollmentSettingOk() (*ApiUnEnrollmentSetting, bool)`

GetUnEnrollmentSettingOk returns a tuple with the UnEnrollmentSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollmentSetting

`func (o *ApiContactFlow) SetUnEnrollmentSetting(v ApiUnEnrollmentSetting)`

SetUnEnrollmentSetting sets UnEnrollmentSetting field to given value.

### HasUnEnrollmentSetting

`func (o *ApiContactFlow) HasUnEnrollmentSetting() bool`

HasUnEnrollmentSetting returns a boolean if a field has been set.

### GetObjectTypeId

`func (o *ApiContactFlow) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiContactFlow) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiContactFlow) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetSuppressionFilterBranch

`func (o *ApiContactFlow) GetSuppressionFilterBranch() ApiPlatformFlowAllOfSuppressionFilterBranch`

GetSuppressionFilterBranch returns the SuppressionFilterBranch field if non-nil, zero value otherwise.

### GetSuppressionFilterBranchOk

`func (o *ApiContactFlow) GetSuppressionFilterBranchOk() (*ApiPlatformFlowAllOfSuppressionFilterBranch, bool)`

GetSuppressionFilterBranchOk returns a tuple with the SuppressionFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionFilterBranch

`func (o *ApiContactFlow) SetSuppressionFilterBranch(v ApiPlatformFlowAllOfSuppressionFilterBranch)`

SetSuppressionFilterBranch sets SuppressionFilterBranch field to given value.

### HasSuppressionFilterBranch

`func (o *ApiContactFlow) HasSuppressionFilterBranch() bool`

HasSuppressionFilterBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


