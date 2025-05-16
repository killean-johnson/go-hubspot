# ApiEventBasedEnrollmentCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefinementCriteria** | Pointer to [**ApiEventBasedEnrollmentCriteriaRefinementCriteria**](ApiEventBasedEnrollmentCriteriaRefinementCriteria.md) |  | [optional] 
**ListMembershipFilterBranches** | [**[]ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch**](ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch.md) | If you want to listen to list-membership events (an object was added to a list, an object was removed from a list) you need to use this &#x60;listMembershipFilterBranches&#x60; property instead of &#x60;eventFilterBranches&#x60;, because list membership events work differently. | 
**EventFilterBranches** | [**[]PublicUnifiedEventsFilterBranch**](PublicUnifiedEventsFilterBranch.md) |  | 
**ShouldReEnroll** | **bool** | Whether or not the same object can enroll in this workflow twice. | 
**Type** | **string** | The type of enrollment criteria this is, this can be \&quot;LIST_BASED\&quot;, \&quot;EVENT_BASED\&quot;, or \&quot;MANUAL\&quot;. | [default to "EVENT_BASED"]

## Methods

### NewApiEventBasedEnrollmentCriteria

`func NewApiEventBasedEnrollmentCriteria(listMembershipFilterBranches []ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch, eventFilterBranches []PublicUnifiedEventsFilterBranch, shouldReEnroll bool, type_ string, ) *ApiEventBasedEnrollmentCriteria`

NewApiEventBasedEnrollmentCriteria instantiates a new ApiEventBasedEnrollmentCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEventBasedEnrollmentCriteriaWithDefaults

`func NewApiEventBasedEnrollmentCriteriaWithDefaults() *ApiEventBasedEnrollmentCriteria`

NewApiEventBasedEnrollmentCriteriaWithDefaults instantiates a new ApiEventBasedEnrollmentCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefinementCriteria

`func (o *ApiEventBasedEnrollmentCriteria) GetRefinementCriteria() ApiEventBasedEnrollmentCriteriaRefinementCriteria`

GetRefinementCriteria returns the RefinementCriteria field if non-nil, zero value otherwise.

### GetRefinementCriteriaOk

`func (o *ApiEventBasedEnrollmentCriteria) GetRefinementCriteriaOk() (*ApiEventBasedEnrollmentCriteriaRefinementCriteria, bool)`

GetRefinementCriteriaOk returns a tuple with the RefinementCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefinementCriteria

`func (o *ApiEventBasedEnrollmentCriteria) SetRefinementCriteria(v ApiEventBasedEnrollmentCriteriaRefinementCriteria)`

SetRefinementCriteria sets RefinementCriteria field to given value.

### HasRefinementCriteria

`func (o *ApiEventBasedEnrollmentCriteria) HasRefinementCriteria() bool`

HasRefinementCriteria returns a boolean if a field has been set.

### GetListMembershipFilterBranches

`func (o *ApiEventBasedEnrollmentCriteria) GetListMembershipFilterBranches() []ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch`

GetListMembershipFilterBranches returns the ListMembershipFilterBranches field if non-nil, zero value otherwise.

### GetListMembershipFilterBranchesOk

`func (o *ApiEventBasedEnrollmentCriteria) GetListMembershipFilterBranchesOk() (*[]ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch, bool)`

GetListMembershipFilterBranchesOk returns a tuple with the ListMembershipFilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMembershipFilterBranches

`func (o *ApiEventBasedEnrollmentCriteria) SetListMembershipFilterBranches(v []ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch)`

SetListMembershipFilterBranches sets ListMembershipFilterBranches field to given value.


### GetEventFilterBranches

`func (o *ApiEventBasedEnrollmentCriteria) GetEventFilterBranches() []PublicUnifiedEventsFilterBranch`

GetEventFilterBranches returns the EventFilterBranches field if non-nil, zero value otherwise.

### GetEventFilterBranchesOk

`func (o *ApiEventBasedEnrollmentCriteria) GetEventFilterBranchesOk() (*[]PublicUnifiedEventsFilterBranch, bool)`

GetEventFilterBranchesOk returns a tuple with the EventFilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilterBranches

`func (o *ApiEventBasedEnrollmentCriteria) SetEventFilterBranches(v []PublicUnifiedEventsFilterBranch)`

SetEventFilterBranches sets EventFilterBranches field to given value.


### GetShouldReEnroll

`func (o *ApiEventBasedEnrollmentCriteria) GetShouldReEnroll() bool`

GetShouldReEnroll returns the ShouldReEnroll field if non-nil, zero value otherwise.

### GetShouldReEnrollOk

`func (o *ApiEventBasedEnrollmentCriteria) GetShouldReEnrollOk() (*bool, bool)`

GetShouldReEnrollOk returns a tuple with the ShouldReEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldReEnroll

`func (o *ApiEventBasedEnrollmentCriteria) SetShouldReEnroll(v bool)`

SetShouldReEnroll sets ShouldReEnroll field to given value.


### GetType

`func (o *ApiEventBasedEnrollmentCriteria) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiEventBasedEnrollmentCriteria) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiEventBasedEnrollmentCriteria) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


