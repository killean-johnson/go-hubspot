# ApiContactFlowPutRequestAllOfEnrollmentCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListFilterBranch** | [**ApiListBasedEnrollmentCriteriaListFilterBranch**](ApiListBasedEnrollmentCriteriaListFilterBranch.md) |  | 
**ReEnrollmentTriggersFilterBranches** | [**[]ApiContactFlowPutRequestAllOfGoalFilterBranch**](ApiContactFlowPutRequestAllOfGoalFilterBranch.md) | A list of filter branches to listen for in order to re-enroll objects into this workflow. | 
**UnEnrollObjectsNotMeetingCriteria** | **bool** | Whether or not to remove objects from this workflow if they stop meeting the enrollment criteria. | 
**ShouldReEnroll** | **bool** | Whether or not the same object can enroll in this workflow twice. | 
**Type** | **string** | The type of enrollment criteria this is, this can be \&quot;LIST_BASED\&quot;, \&quot;EVENT_BASED\&quot;, or \&quot;MANUAL\&quot;. | [default to "LIST_BASED"]
**RefinementCriteria** | Pointer to [**ApiEventBasedEnrollmentCriteriaRefinementCriteria**](ApiEventBasedEnrollmentCriteriaRefinementCriteria.md) |  | [optional] 
**ListMembershipFilterBranches** | [**[]ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch**](ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch.md) | If you want to listen to list-membership events (an object was added to a list, an object was removed from a list) you need to use this &#x60;listMembershipFilterBranches&#x60; property instead of &#x60;eventFilterBranches&#x60;, because list membership events work differently. | 
**EventFilterBranches** | [**[]PublicUnifiedEventsFilterBranch**](PublicUnifiedEventsFilterBranch.md) |  | 

## Methods

### NewApiContactFlowPutRequestAllOfEnrollmentCriteria

`func NewApiContactFlowPutRequestAllOfEnrollmentCriteria(listFilterBranch ApiListBasedEnrollmentCriteriaListFilterBranch, reEnrollmentTriggersFilterBranches []ApiContactFlowPutRequestAllOfGoalFilterBranch, unEnrollObjectsNotMeetingCriteria bool, shouldReEnroll bool, type_ string, listMembershipFilterBranches []ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch, eventFilterBranches []PublicUnifiedEventsFilterBranch, ) *ApiContactFlowPutRequestAllOfEnrollmentCriteria`

NewApiContactFlowPutRequestAllOfEnrollmentCriteria instantiates a new ApiContactFlowPutRequestAllOfEnrollmentCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactFlowPutRequestAllOfEnrollmentCriteriaWithDefaults

`func NewApiContactFlowPutRequestAllOfEnrollmentCriteriaWithDefaults() *ApiContactFlowPutRequestAllOfEnrollmentCriteria`

NewApiContactFlowPutRequestAllOfEnrollmentCriteriaWithDefaults instantiates a new ApiContactFlowPutRequestAllOfEnrollmentCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListFilterBranch

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetListFilterBranch() ApiListBasedEnrollmentCriteriaListFilterBranch`

GetListFilterBranch returns the ListFilterBranch field if non-nil, zero value otherwise.

### GetListFilterBranchOk

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetListFilterBranchOk() (*ApiListBasedEnrollmentCriteriaListFilterBranch, bool)`

GetListFilterBranchOk returns a tuple with the ListFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListFilterBranch

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) SetListFilterBranch(v ApiListBasedEnrollmentCriteriaListFilterBranch)`

SetListFilterBranch sets ListFilterBranch field to given value.


### GetReEnrollmentTriggersFilterBranches

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetReEnrollmentTriggersFilterBranches() []ApiContactFlowPutRequestAllOfGoalFilterBranch`

GetReEnrollmentTriggersFilterBranches returns the ReEnrollmentTriggersFilterBranches field if non-nil, zero value otherwise.

### GetReEnrollmentTriggersFilterBranchesOk

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetReEnrollmentTriggersFilterBranchesOk() (*[]ApiContactFlowPutRequestAllOfGoalFilterBranch, bool)`

GetReEnrollmentTriggersFilterBranchesOk returns a tuple with the ReEnrollmentTriggersFilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReEnrollmentTriggersFilterBranches

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) SetReEnrollmentTriggersFilterBranches(v []ApiContactFlowPutRequestAllOfGoalFilterBranch)`

SetReEnrollmentTriggersFilterBranches sets ReEnrollmentTriggersFilterBranches field to given value.


### GetUnEnrollObjectsNotMeetingCriteria

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetUnEnrollObjectsNotMeetingCriteria() bool`

GetUnEnrollObjectsNotMeetingCriteria returns the UnEnrollObjectsNotMeetingCriteria field if non-nil, zero value otherwise.

### GetUnEnrollObjectsNotMeetingCriteriaOk

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetUnEnrollObjectsNotMeetingCriteriaOk() (*bool, bool)`

GetUnEnrollObjectsNotMeetingCriteriaOk returns a tuple with the UnEnrollObjectsNotMeetingCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollObjectsNotMeetingCriteria

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) SetUnEnrollObjectsNotMeetingCriteria(v bool)`

SetUnEnrollObjectsNotMeetingCriteria sets UnEnrollObjectsNotMeetingCriteria field to given value.


### GetShouldReEnroll

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetShouldReEnroll() bool`

GetShouldReEnroll returns the ShouldReEnroll field if non-nil, zero value otherwise.

### GetShouldReEnrollOk

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetShouldReEnrollOk() (*bool, bool)`

GetShouldReEnrollOk returns a tuple with the ShouldReEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldReEnroll

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) SetShouldReEnroll(v bool)`

SetShouldReEnroll sets ShouldReEnroll field to given value.


### GetType

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) SetType(v string)`

SetType sets Type field to given value.


### GetRefinementCriteria

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetRefinementCriteria() ApiEventBasedEnrollmentCriteriaRefinementCriteria`

GetRefinementCriteria returns the RefinementCriteria field if non-nil, zero value otherwise.

### GetRefinementCriteriaOk

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetRefinementCriteriaOk() (*ApiEventBasedEnrollmentCriteriaRefinementCriteria, bool)`

GetRefinementCriteriaOk returns a tuple with the RefinementCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefinementCriteria

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) SetRefinementCriteria(v ApiEventBasedEnrollmentCriteriaRefinementCriteria)`

SetRefinementCriteria sets RefinementCriteria field to given value.

### HasRefinementCriteria

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) HasRefinementCriteria() bool`

HasRefinementCriteria returns a boolean if a field has been set.

### GetListMembershipFilterBranches

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetListMembershipFilterBranches() []ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch`

GetListMembershipFilterBranches returns the ListMembershipFilterBranches field if non-nil, zero value otherwise.

### GetListMembershipFilterBranchesOk

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetListMembershipFilterBranchesOk() (*[]ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch, bool)`

GetListMembershipFilterBranchesOk returns a tuple with the ListMembershipFilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMembershipFilterBranches

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) SetListMembershipFilterBranches(v []ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch)`

SetListMembershipFilterBranches sets ListMembershipFilterBranches field to given value.


### GetEventFilterBranches

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetEventFilterBranches() []PublicUnifiedEventsFilterBranch`

GetEventFilterBranches returns the EventFilterBranches field if non-nil, zero value otherwise.

### GetEventFilterBranchesOk

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) GetEventFilterBranchesOk() (*[]PublicUnifiedEventsFilterBranch, bool)`

GetEventFilterBranchesOk returns a tuple with the EventFilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilterBranches

`func (o *ApiContactFlowPutRequestAllOfEnrollmentCriteria) SetEventFilterBranches(v []PublicUnifiedEventsFilterBranch)`

SetEventFilterBranches sets EventFilterBranches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


