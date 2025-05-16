# ApiPlatformFlowCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "PLATFORM_FLOW"]
**ObjectTypeId** | **string** |  | 
**CanEnrollFromSalesforce** | **bool** |  | 
**IsEnabled** | **bool** |  | 
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
**SuppressionListIds** | **[]int32** |  | 
**GoalFilterBranch** | Pointer to [**ApiContactFlowCreateRequestAllOfGoalFilterBranch**](ApiContactFlowCreateRequestAllOfGoalFilterBranch.md) |  | [optional] 
**EventAnchor** | Pointer to [**ApiContactFlowCreateRequestAllOfEventAnchor**](ApiContactFlowCreateRequestAllOfEventAnchor.md) |  | [optional] 
**UnEnrollmentSetting** | Pointer to [**ApiUnEnrollmentSetting**](ApiUnEnrollmentSetting.md) |  | [optional] 
**SuppressionFilterBranch** | Pointer to [**ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch**](ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch.md) |  | [optional] 

## Methods

### NewApiPlatformFlowCreateRequest

`func NewApiPlatformFlowCreateRequest(type_ string, objectTypeId string, canEnrollFromSalesforce bool, isEnabled bool, flowType string, actions []ApiPlatformFlowCreateRequestAllOfActions, timeWindows []ApiTimeWindow, blockedDates []ApiBlockedDate, customProperties map[string]string, dataSources []ApiPlatformFlowCreateRequestAllOfDataSources, suppressionListIds []int32, ) *ApiPlatformFlowCreateRequest`

NewApiPlatformFlowCreateRequest instantiates a new ApiPlatformFlowCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPlatformFlowCreateRequestWithDefaults

`func NewApiPlatformFlowCreateRequestWithDefaults() *ApiPlatformFlowCreateRequest`

NewApiPlatformFlowCreateRequestWithDefaults instantiates a new ApiPlatformFlowCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiPlatformFlowCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPlatformFlowCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPlatformFlowCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetObjectTypeId

`func (o *ApiPlatformFlowCreateRequest) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiPlatformFlowCreateRequest) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiPlatformFlowCreateRequest) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetCanEnrollFromSalesforce

`func (o *ApiPlatformFlowCreateRequest) GetCanEnrollFromSalesforce() bool`

GetCanEnrollFromSalesforce returns the CanEnrollFromSalesforce field if non-nil, zero value otherwise.

### GetCanEnrollFromSalesforceOk

`func (o *ApiPlatformFlowCreateRequest) GetCanEnrollFromSalesforceOk() (*bool, bool)`

GetCanEnrollFromSalesforceOk returns a tuple with the CanEnrollFromSalesforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnrollFromSalesforce

`func (o *ApiPlatformFlowCreateRequest) SetCanEnrollFromSalesforce(v bool)`

SetCanEnrollFromSalesforce sets CanEnrollFromSalesforce field to given value.


### GetIsEnabled

`func (o *ApiPlatformFlowCreateRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiPlatformFlowCreateRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiPlatformFlowCreateRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetFlowType

`func (o *ApiPlatformFlowCreateRequest) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *ApiPlatformFlowCreateRequest) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *ApiPlatformFlowCreateRequest) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetName

`func (o *ApiPlatformFlowCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPlatformFlowCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPlatformFlowCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPlatformFlowCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiPlatformFlowCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPlatformFlowCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPlatformFlowCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPlatformFlowCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *ApiPlatformFlowCreateRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiPlatformFlowCreateRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiPlatformFlowCreateRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiPlatformFlowCreateRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStartActionId

`func (o *ApiPlatformFlowCreateRequest) GetStartActionId() string`

GetStartActionId returns the StartActionId field if non-nil, zero value otherwise.

### GetStartActionIdOk

`func (o *ApiPlatformFlowCreateRequest) GetStartActionIdOk() (*string, bool)`

GetStartActionIdOk returns a tuple with the StartActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActionId

`func (o *ApiPlatformFlowCreateRequest) SetStartActionId(v string)`

SetStartActionId sets StartActionId field to given value.

### HasStartActionId

`func (o *ApiPlatformFlowCreateRequest) HasStartActionId() bool`

HasStartActionId returns a boolean if a field has been set.

### GetActions

`func (o *ApiPlatformFlowCreateRequest) GetActions() []ApiPlatformFlowCreateRequestAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPlatformFlowCreateRequest) GetActionsOk() (*[]ApiPlatformFlowCreateRequestAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPlatformFlowCreateRequest) SetActions(v []ApiPlatformFlowCreateRequestAllOfActions)`

SetActions sets Actions field to given value.


### GetEnrollmentCriteria

`func (o *ApiPlatformFlowCreateRequest) GetEnrollmentCriteria() ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria`

GetEnrollmentCriteria returns the EnrollmentCriteria field if non-nil, zero value otherwise.

### GetEnrollmentCriteriaOk

`func (o *ApiPlatformFlowCreateRequest) GetEnrollmentCriteriaOk() (*ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria, bool)`

GetEnrollmentCriteriaOk returns a tuple with the EnrollmentCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCriteria

`func (o *ApiPlatformFlowCreateRequest) SetEnrollmentCriteria(v ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria)`

SetEnrollmentCriteria sets EnrollmentCriteria field to given value.

### HasEnrollmentCriteria

`func (o *ApiPlatformFlowCreateRequest) HasEnrollmentCriteria() bool`

HasEnrollmentCriteria returns a boolean if a field has been set.

### GetEnrollmentSchedule

`func (o *ApiPlatformFlowCreateRequest) GetEnrollmentSchedule() ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule`

GetEnrollmentSchedule returns the EnrollmentSchedule field if non-nil, zero value otherwise.

### GetEnrollmentScheduleOk

`func (o *ApiPlatformFlowCreateRequest) GetEnrollmentScheduleOk() (*ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule, bool)`

GetEnrollmentScheduleOk returns a tuple with the EnrollmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSchedule

`func (o *ApiPlatformFlowCreateRequest) SetEnrollmentSchedule(v ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule)`

SetEnrollmentSchedule sets EnrollmentSchedule field to given value.

### HasEnrollmentSchedule

`func (o *ApiPlatformFlowCreateRequest) HasEnrollmentSchedule() bool`

HasEnrollmentSchedule returns a boolean if a field has been set.

### GetTimeWindows

`func (o *ApiPlatformFlowCreateRequest) GetTimeWindows() []ApiTimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *ApiPlatformFlowCreateRequest) GetTimeWindowsOk() (*[]ApiTimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *ApiPlatformFlowCreateRequest) SetTimeWindows(v []ApiTimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetBlockedDates

`func (o *ApiPlatformFlowCreateRequest) GetBlockedDates() []ApiBlockedDate`

GetBlockedDates returns the BlockedDates field if non-nil, zero value otherwise.

### GetBlockedDatesOk

`func (o *ApiPlatformFlowCreateRequest) GetBlockedDatesOk() (*[]ApiBlockedDate, bool)`

GetBlockedDatesOk returns a tuple with the BlockedDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDates

`func (o *ApiPlatformFlowCreateRequest) SetBlockedDates(v []ApiBlockedDate)`

SetBlockedDates sets BlockedDates field to given value.


### GetCustomProperties

`func (o *ApiPlatformFlowCreateRequest) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ApiPlatformFlowCreateRequest) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ApiPlatformFlowCreateRequest) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.


### GetDataSources

`func (o *ApiPlatformFlowCreateRequest) GetDataSources() []ApiPlatformFlowCreateRequestAllOfDataSources`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ApiPlatformFlowCreateRequest) GetDataSourcesOk() (*[]ApiPlatformFlowCreateRequestAllOfDataSources, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ApiPlatformFlowCreateRequest) SetDataSources(v []ApiPlatformFlowCreateRequestAllOfDataSources)`

SetDataSources sets DataSources field to given value.


### GetSuppressionListIds

`func (o *ApiPlatformFlowCreateRequest) GetSuppressionListIds() []int32`

GetSuppressionListIds returns the SuppressionListIds field if non-nil, zero value otherwise.

### GetSuppressionListIdsOk

`func (o *ApiPlatformFlowCreateRequest) GetSuppressionListIdsOk() (*[]int32, bool)`

GetSuppressionListIdsOk returns a tuple with the SuppressionListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionListIds

`func (o *ApiPlatformFlowCreateRequest) SetSuppressionListIds(v []int32)`

SetSuppressionListIds sets SuppressionListIds field to given value.


### GetGoalFilterBranch

`func (o *ApiPlatformFlowCreateRequest) GetGoalFilterBranch() ApiContactFlowCreateRequestAllOfGoalFilterBranch`

GetGoalFilterBranch returns the GoalFilterBranch field if non-nil, zero value otherwise.

### GetGoalFilterBranchOk

`func (o *ApiPlatformFlowCreateRequest) GetGoalFilterBranchOk() (*ApiContactFlowCreateRequestAllOfGoalFilterBranch, bool)`

GetGoalFilterBranchOk returns a tuple with the GoalFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilterBranch

`func (o *ApiPlatformFlowCreateRequest) SetGoalFilterBranch(v ApiContactFlowCreateRequestAllOfGoalFilterBranch)`

SetGoalFilterBranch sets GoalFilterBranch field to given value.

### HasGoalFilterBranch

`func (o *ApiPlatformFlowCreateRequest) HasGoalFilterBranch() bool`

HasGoalFilterBranch returns a boolean if a field has been set.

### GetEventAnchor

`func (o *ApiPlatformFlowCreateRequest) GetEventAnchor() ApiContactFlowCreateRequestAllOfEventAnchor`

GetEventAnchor returns the EventAnchor field if non-nil, zero value otherwise.

### GetEventAnchorOk

`func (o *ApiPlatformFlowCreateRequest) GetEventAnchorOk() (*ApiContactFlowCreateRequestAllOfEventAnchor, bool)`

GetEventAnchorOk returns a tuple with the EventAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAnchor

`func (o *ApiPlatformFlowCreateRequest) SetEventAnchor(v ApiContactFlowCreateRequestAllOfEventAnchor)`

SetEventAnchor sets EventAnchor field to given value.

### HasEventAnchor

`func (o *ApiPlatformFlowCreateRequest) HasEventAnchor() bool`

HasEventAnchor returns a boolean if a field has been set.

### GetUnEnrollmentSetting

`func (o *ApiPlatformFlowCreateRequest) GetUnEnrollmentSetting() ApiUnEnrollmentSetting`

GetUnEnrollmentSetting returns the UnEnrollmentSetting field if non-nil, zero value otherwise.

### GetUnEnrollmentSettingOk

`func (o *ApiPlatformFlowCreateRequest) GetUnEnrollmentSettingOk() (*ApiUnEnrollmentSetting, bool)`

GetUnEnrollmentSettingOk returns a tuple with the UnEnrollmentSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollmentSetting

`func (o *ApiPlatformFlowCreateRequest) SetUnEnrollmentSetting(v ApiUnEnrollmentSetting)`

SetUnEnrollmentSetting sets UnEnrollmentSetting field to given value.

### HasUnEnrollmentSetting

`func (o *ApiPlatformFlowCreateRequest) HasUnEnrollmentSetting() bool`

HasUnEnrollmentSetting returns a boolean if a field has been set.

### GetSuppressionFilterBranch

`func (o *ApiPlatformFlowCreateRequest) GetSuppressionFilterBranch() ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch`

GetSuppressionFilterBranch returns the SuppressionFilterBranch field if non-nil, zero value otherwise.

### GetSuppressionFilterBranchOk

`func (o *ApiPlatformFlowCreateRequest) GetSuppressionFilterBranchOk() (*ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch, bool)`

GetSuppressionFilterBranchOk returns a tuple with the SuppressionFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionFilterBranch

`func (o *ApiPlatformFlowCreateRequest) SetSuppressionFilterBranch(v ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch)`

SetSuppressionFilterBranch sets SuppressionFilterBranch field to given value.

### HasSuppressionFilterBranch

`func (o *ApiPlatformFlowCreateRequest) HasSuppressionFilterBranch() bool`

HasSuppressionFilterBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


