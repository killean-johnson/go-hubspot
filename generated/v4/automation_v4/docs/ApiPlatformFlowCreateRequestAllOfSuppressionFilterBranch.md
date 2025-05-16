# ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch

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

### NewApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch

`func NewApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch(filterBranchType string, filterBranches []ApiContactFlowPutRequestAllOfGoalFilterBranch, filterBranchOperator string, filters []PublicPropertyAssociationFilterBranchFiltersInner, eventTypeId string, operator string, objectTypeId string, propertyWithObjectId string, associationTypeId int32, associationCategory string, ) *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch`

NewApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch instantiates a new ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPlatformFlowCreateRequestAllOfSuppressionFilterBranchWithDefaults

`func NewApiPlatformFlowCreateRequestAllOfSuppressionFilterBranchWithDefaults() *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch`

NewApiPlatformFlowCreateRequestAllOfSuppressionFilterBranchWithDefaults instantiates a new ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterBranchType

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetFilterBranchType() string`

GetFilterBranchType returns the FilterBranchType field if non-nil, zero value otherwise.

### GetFilterBranchTypeOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetFilterBranchTypeOk() (*string, bool)`

GetFilterBranchTypeOk returns a tuple with the FilterBranchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchType

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetFilterBranchType(v string)`

SetFilterBranchType sets FilterBranchType field to given value.


### GetFilterBranches

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetFilterBranches() []ApiContactFlowPutRequestAllOfGoalFilterBranch`

GetFilterBranches returns the FilterBranches field if non-nil, zero value otherwise.

### GetFilterBranchesOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetFilterBranchesOk() (*[]ApiContactFlowPutRequestAllOfGoalFilterBranch, bool)`

GetFilterBranchesOk returns a tuple with the FilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranches

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetFilterBranches(v []ApiContactFlowPutRequestAllOfGoalFilterBranch)`

SetFilterBranches sets FilterBranches field to given value.


### GetFilterBranchOperator

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetFilterBranchOperator() string`

GetFilterBranchOperator returns the FilterBranchOperator field if non-nil, zero value otherwise.

### GetFilterBranchOperatorOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetFilterBranchOperatorOk() (*string, bool)`

GetFilterBranchOperatorOk returns a tuple with the FilterBranchOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchOperator

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetFilterBranchOperator(v string)`

SetFilterBranchOperator sets FilterBranchOperator field to given value.


### GetFilters

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetFilters() []PublicPropertyAssociationFilterBranchFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetFiltersOk() (*[]PublicPropertyAssociationFilterBranchFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetFilters(v []PublicPropertyAssociationFilterBranchFiltersInner)`

SetFilters sets Filters field to given value.


### GetEventTypeId

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetEventTypeId() string`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetEventTypeIdOk() (*string, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetEventTypeId(v string)`

SetEventTypeId sets EventTypeId field to given value.


### GetCoalescingRefineBy

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetCoalescingRefineBy() PublicFormSubmissionFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetCoalescingRefineByOk() (*PublicFormSubmissionFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetCoalescingRefineBy(v PublicFormSubmissionFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.

### HasCoalescingRefineBy

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) HasCoalescingRefineBy() bool`

HasCoalescingRefineBy returns a boolean if a field has been set.

### GetOperator

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetObjectTypeId

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetPropertyWithObjectId

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetPropertyWithObjectId() string`

GetPropertyWithObjectId returns the PropertyWithObjectId field if non-nil, zero value otherwise.

### GetPropertyWithObjectIdOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetPropertyWithObjectIdOk() (*string, bool)`

GetPropertyWithObjectIdOk returns a tuple with the PropertyWithObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyWithObjectId

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetPropertyWithObjectId(v string)`

SetPropertyWithObjectId sets PropertyWithObjectId field to given value.


### GetAssociationTypeId

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


