# ApiEventBasedEnrollmentCriteriaRefinementCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterBranchType** | **string** |  | [default to "ASSOCIATION"]
**FilterBranches** | [**[]ApiContactFlowPutRequestAllOfGoalFilterBranch**](ApiContactFlowPutRequestAllOfGoalFilterBranch.md) |  | 
**FilterBranchOperator** | **string** |  | 
**Filters** | [**[]PublicPropertyAssociationFilterBranchFiltersInner**](PublicPropertyAssociationFilterBranchFiltersInner.md) |  | 
**EventTypeId** | **string** |  | 
**CoalescingRefineBy** | Pointer to [**PublicFormSubmissionFilterCoalescingRefineBy**](PublicFormSubmissionFilterCoalescingRefineBy.md) |  | [optional] 
**Operator** | **string** |  | 
**ObjectTypeId** | **string** |  | 
**PropertyWithObjectId** | **string** |  | 
**AssociationTypeId** | **int32** |  | 
**AssociationCategory** | **string** |  | 

## Methods

### NewApiEventBasedEnrollmentCriteriaRefinementCriteria

`func NewApiEventBasedEnrollmentCriteriaRefinementCriteria(filterBranchType string, filterBranches []ApiContactFlowPutRequestAllOfGoalFilterBranch, filterBranchOperator string, filters []PublicPropertyAssociationFilterBranchFiltersInner, eventTypeId string, operator string, objectTypeId string, propertyWithObjectId string, associationTypeId int32, associationCategory string, ) *ApiEventBasedEnrollmentCriteriaRefinementCriteria`

NewApiEventBasedEnrollmentCriteriaRefinementCriteria instantiates a new ApiEventBasedEnrollmentCriteriaRefinementCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEventBasedEnrollmentCriteriaRefinementCriteriaWithDefaults

`func NewApiEventBasedEnrollmentCriteriaRefinementCriteriaWithDefaults() *ApiEventBasedEnrollmentCriteriaRefinementCriteria`

NewApiEventBasedEnrollmentCriteriaRefinementCriteriaWithDefaults instantiates a new ApiEventBasedEnrollmentCriteriaRefinementCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterBranchType

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetFilterBranchType() string`

GetFilterBranchType returns the FilterBranchType field if non-nil, zero value otherwise.

### GetFilterBranchTypeOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetFilterBranchTypeOk() (*string, bool)`

GetFilterBranchTypeOk returns a tuple with the FilterBranchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchType

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetFilterBranchType(v string)`

SetFilterBranchType sets FilterBranchType field to given value.


### GetFilterBranches

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetFilterBranches() []ApiContactFlowPutRequestAllOfGoalFilterBranch`

GetFilterBranches returns the FilterBranches field if non-nil, zero value otherwise.

### GetFilterBranchesOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetFilterBranchesOk() (*[]ApiContactFlowPutRequestAllOfGoalFilterBranch, bool)`

GetFilterBranchesOk returns a tuple with the FilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranches

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetFilterBranches(v []ApiContactFlowPutRequestAllOfGoalFilterBranch)`

SetFilterBranches sets FilterBranches field to given value.


### GetFilterBranchOperator

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetFilterBranchOperator() string`

GetFilterBranchOperator returns the FilterBranchOperator field if non-nil, zero value otherwise.

### GetFilterBranchOperatorOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetFilterBranchOperatorOk() (*string, bool)`

GetFilterBranchOperatorOk returns a tuple with the FilterBranchOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchOperator

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetFilterBranchOperator(v string)`

SetFilterBranchOperator sets FilterBranchOperator field to given value.


### GetFilters

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetFilters() []PublicPropertyAssociationFilterBranchFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetFiltersOk() (*[]PublicPropertyAssociationFilterBranchFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetFilters(v []PublicPropertyAssociationFilterBranchFiltersInner)`

SetFilters sets Filters field to given value.


### GetEventTypeId

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetEventTypeId() string`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetEventTypeIdOk() (*string, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetEventTypeId(v string)`

SetEventTypeId sets EventTypeId field to given value.


### GetCoalescingRefineBy

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetCoalescingRefineBy() PublicFormSubmissionFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetCoalescingRefineByOk() (*PublicFormSubmissionFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetCoalescingRefineBy(v PublicFormSubmissionFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.

### HasCoalescingRefineBy

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) HasCoalescingRefineBy() bool`

HasCoalescingRefineBy returns a boolean if a field has been set.

### GetOperator

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetObjectTypeId

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetPropertyWithObjectId

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetPropertyWithObjectId() string`

GetPropertyWithObjectId returns the PropertyWithObjectId field if non-nil, zero value otherwise.

### GetPropertyWithObjectIdOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetPropertyWithObjectIdOk() (*string, bool)`

GetPropertyWithObjectIdOk returns a tuple with the PropertyWithObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyWithObjectId

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetPropertyWithObjectId(v string)`

SetPropertyWithObjectId sets PropertyWithObjectId field to given value.


### GetAssociationTypeId

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *ApiEventBasedEnrollmentCriteriaRefinementCriteria) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


