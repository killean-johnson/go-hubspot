# ApiListBasedEnrollmentCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListFilterBranch** | [**ApiListBasedEnrollmentCriteriaListFilterBranch**](ApiListBasedEnrollmentCriteriaListFilterBranch.md) |  | 
**ReEnrollmentTriggersFilterBranches** | [**[]ApiContactFlowPutRequestAllOfGoalFilterBranch**](ApiContactFlowPutRequestAllOfGoalFilterBranch.md) | A list of filter branches to listen for in order to re-enroll objects into this workflow. | 
**UnEnrollObjectsNotMeetingCriteria** | **bool** | Whether or not to remove objects from this workflow if they stop meeting the enrollment criteria. | 
**ShouldReEnroll** | **bool** | Whether or not the same object can enroll in this workflow twice. | 
**Type** | **string** | The type of enrollment criteria this is, this can be \&quot;LIST_BASED\&quot;, \&quot;EVENT_BASED\&quot;, or \&quot;MANUAL\&quot;. | [default to "LIST_BASED"]

## Methods

### NewApiListBasedEnrollmentCriteria

`func NewApiListBasedEnrollmentCriteria(listFilterBranch ApiListBasedEnrollmentCriteriaListFilterBranch, reEnrollmentTriggersFilterBranches []ApiContactFlowPutRequestAllOfGoalFilterBranch, unEnrollObjectsNotMeetingCriteria bool, shouldReEnroll bool, type_ string, ) *ApiListBasedEnrollmentCriteria`

NewApiListBasedEnrollmentCriteria instantiates a new ApiListBasedEnrollmentCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiListBasedEnrollmentCriteriaWithDefaults

`func NewApiListBasedEnrollmentCriteriaWithDefaults() *ApiListBasedEnrollmentCriteria`

NewApiListBasedEnrollmentCriteriaWithDefaults instantiates a new ApiListBasedEnrollmentCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListFilterBranch

`func (o *ApiListBasedEnrollmentCriteria) GetListFilterBranch() ApiListBasedEnrollmentCriteriaListFilterBranch`

GetListFilterBranch returns the ListFilterBranch field if non-nil, zero value otherwise.

### GetListFilterBranchOk

`func (o *ApiListBasedEnrollmentCriteria) GetListFilterBranchOk() (*ApiListBasedEnrollmentCriteriaListFilterBranch, bool)`

GetListFilterBranchOk returns a tuple with the ListFilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListFilterBranch

`func (o *ApiListBasedEnrollmentCriteria) SetListFilterBranch(v ApiListBasedEnrollmentCriteriaListFilterBranch)`

SetListFilterBranch sets ListFilterBranch field to given value.


### GetReEnrollmentTriggersFilterBranches

`func (o *ApiListBasedEnrollmentCriteria) GetReEnrollmentTriggersFilterBranches() []ApiContactFlowPutRequestAllOfGoalFilterBranch`

GetReEnrollmentTriggersFilterBranches returns the ReEnrollmentTriggersFilterBranches field if non-nil, zero value otherwise.

### GetReEnrollmentTriggersFilterBranchesOk

`func (o *ApiListBasedEnrollmentCriteria) GetReEnrollmentTriggersFilterBranchesOk() (*[]ApiContactFlowPutRequestAllOfGoalFilterBranch, bool)`

GetReEnrollmentTriggersFilterBranchesOk returns a tuple with the ReEnrollmentTriggersFilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReEnrollmentTriggersFilterBranches

`func (o *ApiListBasedEnrollmentCriteria) SetReEnrollmentTriggersFilterBranches(v []ApiContactFlowPutRequestAllOfGoalFilterBranch)`

SetReEnrollmentTriggersFilterBranches sets ReEnrollmentTriggersFilterBranches field to given value.


### GetUnEnrollObjectsNotMeetingCriteria

`func (o *ApiListBasedEnrollmentCriteria) GetUnEnrollObjectsNotMeetingCriteria() bool`

GetUnEnrollObjectsNotMeetingCriteria returns the UnEnrollObjectsNotMeetingCriteria field if non-nil, zero value otherwise.

### GetUnEnrollObjectsNotMeetingCriteriaOk

`func (o *ApiListBasedEnrollmentCriteria) GetUnEnrollObjectsNotMeetingCriteriaOk() (*bool, bool)`

GetUnEnrollObjectsNotMeetingCriteriaOk returns a tuple with the UnEnrollObjectsNotMeetingCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnEnrollObjectsNotMeetingCriteria

`func (o *ApiListBasedEnrollmentCriteria) SetUnEnrollObjectsNotMeetingCriteria(v bool)`

SetUnEnrollObjectsNotMeetingCriteria sets UnEnrollObjectsNotMeetingCriteria field to given value.


### GetShouldReEnroll

`func (o *ApiListBasedEnrollmentCriteria) GetShouldReEnroll() bool`

GetShouldReEnroll returns the ShouldReEnroll field if non-nil, zero value otherwise.

### GetShouldReEnrollOk

`func (o *ApiListBasedEnrollmentCriteria) GetShouldReEnrollOk() (*bool, bool)`

GetShouldReEnrollOk returns a tuple with the ShouldReEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldReEnroll

`func (o *ApiListBasedEnrollmentCriteria) SetShouldReEnroll(v bool)`

SetShouldReEnroll sets ShouldReEnroll field to given value.


### GetType

`func (o *ApiListBasedEnrollmentCriteria) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiListBasedEnrollmentCriteria) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiListBasedEnrollmentCriteria) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


