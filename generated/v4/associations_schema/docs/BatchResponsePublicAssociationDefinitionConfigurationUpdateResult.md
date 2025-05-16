# BatchResponsePublicAssociationDefinitionConfigurationUpdateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** |  | 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 
**Results** | [**[]PublicAssociationDefinitionConfigurationUpdateResult**](PublicAssociationDefinitionConfigurationUpdateResult.md) |  | 
**Status** | **string** |  | 

## Methods

### NewBatchResponsePublicAssociationDefinitionConfigurationUpdateResult

`func NewBatchResponsePublicAssociationDefinitionConfigurationUpdateResult(completedAt time.Time, startedAt time.Time, results []PublicAssociationDefinitionConfigurationUpdateResult, status string, ) *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult`

NewBatchResponsePublicAssociationDefinitionConfigurationUpdateResult instantiates a new BatchResponsePublicAssociationDefinitionConfigurationUpdateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponsePublicAssociationDefinitionConfigurationUpdateResultWithDefaults

`func NewBatchResponsePublicAssociationDefinitionConfigurationUpdateResultWithDefaults() *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult`

NewBatchResponsePublicAssociationDefinitionConfigurationUpdateResultWithDefaults instantiates a new BatchResponsePublicAssociationDefinitionConfigurationUpdateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetRequestedAt

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetResults() []PublicAssociationDefinitionConfigurationUpdateResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetResultsOk() (*[]PublicAssociationDefinitionConfigurationUpdateResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) SetResults(v []PublicAssociationDefinitionConfigurationUpdateResult)`

SetResults sets Results field to given value.


### GetStatus

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


