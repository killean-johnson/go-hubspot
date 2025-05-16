# ApiContactFlowCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "CONTACT_FLOW"]
**ObjectTypeId** | **string** |  | 
**CanEnrollFromSalesforce** | **bool** |  | 
**IsEnabled** | **bool** |  | 
**FlowType** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**StartActionId** | Pointer to **string** |  | [optional] 
**Actions** | [**[]ApiContactFlowCreateRequestAllOfActions**](ApiContactFlowCreateRequestAllOfActions.md) |  | 
**EnrollmentCriteria** | Pointer to [**ApiContactFlowCreateRequestAllOfEnrollmentCriteria**](ApiContactFlowCreateRequestAllOfEnrollmentCriteria.md) |  | [optional] 
**EnrollmentSchedule** | Pointer to [**ApiContactFlowCreateRequestAllOfEnrollmentSchedule**](ApiContactFlowCreateRequestAllOfEnrollmentSchedule.md) |  | [optional] 
**TimeWindows** | [**[]ApiTimeWindow**](ApiTimeWindow.md) |  | 
**BlockedDates** | [**[]ApiBlockedDate**](ApiBlockedDate.md) |  | 
**CustomProperties** | **map[string]string** |  | 
**DataSources** | [**[]ApiContactFlowCreateRequestAllOfDataSources**](ApiContactFlowCreateRequestAllOfDataSources.md) |  | 
**SuppressionListIds** | **[]int32** |  | 
**GoalFilterBranch** | Pointer to [**ApiContactFlowCreateRequestAllOfGoalFilterBranch**](ApiContactFlowCreateRequestAllOfGoalFilterBranch.md) |  | [optional] 
**EventAnchor** | Pointer to [**ApiContactFlowCreateRequestAllOfEventAnchor**](ApiContactFlowCreateRequestAllOfEventAnchor.md) |  | [optional] 
**UnEnrollmentSetting** | Pointer to [**ApiUnEnrollmentSetting**](ApiUnEnrollmentSetting.md) |  | [optional] 
**SuppressionFilterBranch** | Pointer to [**ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch**](ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch.md) |  | [optional] 

## Methods

### NewApiContactFlowCreateRequest

`func NewApiContactFlowCreateRequest(type_ string, objectTypeId string, canEnrollFromSalesforce bool, isEnabled bool, flowType string, actions []ApiContactFlowCreateRequestAllOfActions, timeWindows []ApiTimeWindow, blockedDates []ApiBlockedDate, customProperties map[string]string, dataSources []ApiContactFlowCreateRequestAllOfDataSources, suppressionListIds []int32, ) *ApiContactFlowCreateRequest`

NewApiContactFlowCreateRequest instantiates a new ApiContactFlowCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactFlowCreateRequestWithDefaults

`func NewApiContactFlowCreateRequestWithDefaults() *ApiContactFlowCreateRequest`

NewApiContactFlowCreateRequestWithDefaults instantiates a new ApiContactFlowCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiContactFlowCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactFlowCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactFlowCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetObjectTypeId

`func (o *ApiContactFlowCreateRequest) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiContactFlowCreateRequest) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiContactFlowCreateRequest) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetCanEnrollFromSalesforce

`func (o *ApiContactFlowCreateRequest) GetCanEnrollFromSalesforce() bool`

GetCanEnrollFromSalesforce returns the CanEnrollFromSalesforce field if non-nil, zero value otherwise.

### GetCanEnrollFromSalesforceOk

`func (o *ApiContactFlowCreateRequest) GetCanEnrollFromSalesforceOk() (*bool, bool)`

GetCanEnrollFromSalesforceOk returns a tuple with the CanEnrollFromSalesforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnrollFromSalesforce

`func (o *ApiContactFlowCreateRequest) SetCanEnrollFromSalesforce(v bool)`

SetCanEnrollFromSalesforce sets CanEnrollFromSalesforce field to given value.


### GetIsEnabled

`func (o *ApiContactFlowCreateRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiContactFlowCreateRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiContactFlowCreateRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetFlowType

`func (o *ApiContactFlowCreateRequest) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *ApiContactFlowCreateRequest) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *ApiContactFlowCreateRequest) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetName

`func (o *ApiContactFlowCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiContactFlowCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiContactFlowCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiContactFlowCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiContactFlowCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiContactFlowCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiContactFlowCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiContactFlowCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *ApiContactFlowCreateRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiContactFlowCreateRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiContactFlowCreateRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiContactFlowCreateRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStartActionId

`func (o *ApiContactFlowCreateRequest) GetStartActionId() string`

GetStartActionId returns the StartActionId field if non-nil, zero value otherwise.

### GetStartActionIdOk

`func (o *ApiContactFlowCreateRequest) GetStartActionIdOk() (*string, bool)`

GetStartActionIdOk returns a tuple with the StartActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActionId

`func (o *ApiContactFlowCreateRequest) SetStartActionId(v string)`

SetStartActionId sets StartActionId field to given value.

### HasStartActionId

`func (o *ApiContactFlowCreateRequest) HasStartActionId() bool`

HasStartActionId returns a boolean if a field has been set.

### GetActions

`func (o *ApiContactFlowCreateRequest) GetActions() []ApiContactFlowCreateRequestAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiContactFlowCreateRequest) GetActionsOk() (*[]ApiContactFlowCreateRequestAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiContactFlowCreateRequest) SetActions(v []ApiContactFlowCreateRequestAllOfActions)`

SetActions sets Actions field to given value.


### GetEnrollmentCriteria

`func (o *ApiContactFlowCreateRequest) GetEnrollmentCriteria() ApiContactFlowCreateRequestAllOfEnrollmentCriteria`

GetEnrollmentCriteria returns the EnrollmentCriteria field if non-nil, zero value otherwise.

### GetEnrollmentCriteriaOk

`func (o *ApiContactFlowCreateRequest) GetEnrollmentCriteriaOk() (*ApiContactFlowCreateRequestAllOfEnrollmentCriteria, bool)`

GetEnrollmentCriteriaOk returns a tuple with the EnrollmentCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCriteria

`func (o *ApiContactFlowCreateRequest) SetEnrollmentCriteria(v ApiContactFlowCreateRequestAllOfEnrollmentCriteria)`

SetEnrollmentCriteria sets EnrollmentCriteria field to given value.

### HasEnrollmentCriteria

`func (o *ApiContactFlowCreateRequest) HasEnrollmentCriteria() bool`

HasEnrollmentCriteria returns a boolean if a field has been set.

### GetEnrollmentSchedule

`func (o *ApiContactFlowCreateRequest) GetEnrollmentSchedule() ApiContactFlowCreateRequestAllOfEnrollmentSchedule`

GetEnrollmentSchedule returns the EnrollmentSchedule field if non-nil, zero value otherwise.

### GetEnrollmentScheduleOk

`func (o *ApiContactFlowCreateRequest) GetEnrollmentScheduleOk() (*ApiContactFlowCreateRequestAllOfEnrollmentSchedule, bool)`

GetEnrollmentScheduleOk returns a tuple with the EnrollmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSchedule

`func (o *ApiContactFlowCreateRequest) SetEnrollmentSchedule(v ApiContactFlowCreateRequestAllOfEnrollmentSchedule)`

SetEnrollmentSchedule sets EnrollmentSchedule field to given value.

### HasEnrollmentSchedule

`func (o *ApiContactFlowCreateRequest) HasEnrollmentSchedule() bool`

HasEnrollmentSchedule returns a boolean if a field has been set.

### GetTimeWindows

`func (o *ApiContactFlowCreateRequest) GetTimeWindows() []ApiTimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *ApiContactFlowCreateRequest) GetTimeWindowsOk() (*[]ApiTimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *ApiContactFlowCreateRequest) SetTimeWindows(v []ApiTimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetBlockedDates

`func (o *ApiContactFlowCreateRequest) GetBlockedDates() []ApiBlockedDate`

GetBlockedDates returns the BlockedDates field if non-nil, zero value otherwise.

### GetBlockedDatesOk

`func (o *ApiContactFlowCreateRequest) GetBlockedDatesOk() (*[]ApiBlockedDate, bool)`

GetBlockedDatesOk returns a tuple with the BlockedDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDates

`func (o *ApiContactFlowCreateRequest) SetBlockedDates(v []ApiBlockedDate)`

SetBlockedDates sets BlockedDates field to given value.


### GetCustomProperties

`func (o *ApiContactFlowCreateRequest) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ApiContactFlowCreateRequest) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ApiContactFlowCreateRequest) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.


### GetDataSources

`func (o *ApiContactFlowCreateRequest) GetDataSources() []ApiContactFlowCreateRequestAllOfDataSources`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ApiContactFlowCreateRequest) GetDataSourcesOk() (*[]ApiContactFlowCreateRequestAllOfDataSources, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ApiContactFlowCreateRequest) SetDataSources(v []ApiContactFlowCreateRequestAllOfDataSources)`

SetDataSources sets DataSources field to given value.


### GetSuppressionListIds

`func (o *ApiContactFlowCreateRequest) GetSuppressionListIds() []int32`

GetSuppressionListIds returns the SuppressionListIds field if non-nil, zero value otherwise.

### GetSuppressionListIdsOk

`func (o *ApiContactFlowCreateRequest) GetSuppressionListIdsOk() (*[]int32, bool)`

GetSuppressionListIdsOk returns a tuple with the SuppressionListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionListIds

`func (o *ApiContactFlowCreateRequest) SetSuppressionListIds(v []int32)`

SetSuppressionListIds sets SuppressionListIds field to given value.


### GetGoalFilterBranch

`func (o *ApiContactFlowCreateRequest) GetGoalFilterBranch() ApiContactFlowCreateRequestAllOfGoalFilterBranch`

GetGoalFilterBranch returns the GoalFilterBranch field if non-nil, zero value otherwise.

### GetGoalFilterBranchOk

`func (o *ApiContactFlowCreateRequest) GetGoalFilterBranchOk() (*ApiContactFlowCreateRequestAllOfGoalFilterBranch, bool)`

GetGoalFilterBranchOk returns a tuple with the GoalFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilterBranch

`func (o *ApiContactFlowCreateRequest) SetGoalFilterBranch(v ApiContactFlowCreateRequestAllOfGoalFilterBranch)`

SetGoalFilterBranch sets GoalFilterBranch field to given value.

### HasGoalFilterBranch

`func (o *ApiContactFlowCreateRequest) HasGoalFilterBranch() bool`

HasGoalFilterBranch returns a boolean if a field has been set.

### GetEventAnchor

`func (o *ApiContactFlowCreateRequest) GetEventAnchor() ApiContactFlowCreateRequestAllOfEventAnchor`

GetEventAnchor returns the EventAnchor field if non-nil, zero value otherwise.

### GetEventAnchorOk

`func (o *ApiContactFlowCreateRequest) GetEventAnchorOk() (*ApiContactFlowCreateRequestAllOfEventAnchor, bool)`

GetEventAnchorOk returns a tuple with the EventAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAnchor

`func (o *ApiContactFlowCreateRequest) SetEventAnchor(v ApiContactFlowCreateRequestAllOfEventAnchor)`

SetEventAnchor sets EventAnchor field to given value.

### HasEventAnchor

`func (o *ApiContactFlowCreateRequest) HasEventAnchor() bool`

HasEventAnchor returns a boolean if a field has been set.

### GetUnEnrollmentSetting

`func (o *ApiContactFlowCreateRequest) GetUnEnrollmentSetting() ApiUnEnrollmentSetting`

GetUnEnrollmentSetting returns the UnEnrollmentSetting field if non-nil, zero value otherwise.

### GetUnEnrollmentSettingOk

`func (o *ApiContactFlowCreateRequest) GetUnEnrollmentSettingOk() (*ApiUnEnrollmentSetting, bool)`

GetUnEnrollmentSettingOk returns a tuple with the UnEnrollmentSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollmentSetting

`func (o *ApiContactFlowCreateRequest) SetUnEnrollmentSetting(v ApiUnEnrollmentSetting)`

SetUnEnrollmentSetting sets UnEnrollmentSetting field to given value.

### HasUnEnrollmentSetting

`func (o *ApiContactFlowCreateRequest) HasUnEnrollmentSetting() bool`

HasUnEnrollmentSetting returns a boolean if a field has been set.

### GetSuppressionFilterBranch

`func (o *ApiContactFlowCreateRequest) GetSuppressionFilterBranch() ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch`

GetSuppressionFilterBranch returns the SuppressionFilterBranch field if non-nil, zero value otherwise.

### GetSuppressionFilterBranchOk

`func (o *ApiContactFlowCreateRequest) GetSuppressionFilterBranchOk() (*ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch, bool)`

GetSuppressionFilterBranchOk returns a tuple with the SuppressionFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionFilterBranch

`func (o *ApiContactFlowCreateRequest) SetSuppressionFilterBranch(v ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch)`

SetSuppressionFilterBranch sets SuppressionFilterBranch field to given value.

### HasSuppressionFilterBranch

`func (o *ApiContactFlowCreateRequest) HasSuppressionFilterBranch() bool`

HasSuppressionFilterBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


