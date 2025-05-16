# BatchResponseFlowIdWorkflowIdMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** |  | 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 
**Results** | [**[]FlowIdWorkflowIdMappingResponse**](FlowIdWorkflowIdMappingResponse.md) |  | 
**Status** | **string** |  | 

## Methods

### NewBatchResponseFlowIdWorkflowIdMappingResponse

`func NewBatchResponseFlowIdWorkflowIdMappingResponse(completedAt time.Time, startedAt time.Time, results []FlowIdWorkflowIdMappingResponse, status string, ) *BatchResponseFlowIdWorkflowIdMappingResponse`

NewBatchResponseFlowIdWorkflowIdMappingResponse instantiates a new BatchResponseFlowIdWorkflowIdMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseFlowIdWorkflowIdMappingResponseWithDefaults

`func NewBatchResponseFlowIdWorkflowIdMappingResponseWithDefaults() *BatchResponseFlowIdWorkflowIdMappingResponse`

NewBatchResponseFlowIdWorkflowIdMappingResponseWithDefaults instantiates a new BatchResponseFlowIdWorkflowIdMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetRequestedAt

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetResults() []FlowIdWorkflowIdMappingResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetResultsOk() (*[]FlowIdWorkflowIdMappingResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) SetResults(v []FlowIdWorkflowIdMappingResponse)`

SetResults sets Results field to given value.


### GetStatus

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseFlowIdWorkflowIdMappingResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


