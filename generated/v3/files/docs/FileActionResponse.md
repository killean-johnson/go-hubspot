# FileActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**File**](File.md) |  | [optional] 
**CompletedAt** | **time.Time** | Time of completion of task. | 
**NumErrors** | Pointer to **int32** | Number of errors resulting from the task. | [optional] 
**RequestedAt** | Pointer to **time.Time** | Timestamp of when the task was requested. | [optional] 
**StartedAt** | **time.Time** | Timestamp of when the task was started. | 
**Links** | Pointer to **map[string]string** | Link to check the status of the requested task. | [optional] 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) | Descriptive error messages. | [optional] 
**TaskId** | **string** | ID of the requested task. | 
**Status** | **string** | Current status of the task. | 

## Methods

### NewFileActionResponse

`func NewFileActionResponse(completedAt time.Time, startedAt time.Time, taskId string, status string, ) *FileActionResponse`

NewFileActionResponse instantiates a new FileActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileActionResponseWithDefaults

`func NewFileActionResponseWithDefaults() *FileActionResponse`

NewFileActionResponseWithDefaults instantiates a new FileActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *FileActionResponse) GetResult() File`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FileActionResponse) GetResultOk() (*File, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FileActionResponse) SetResult(v File)`

SetResult sets Result field to given value.

### HasResult

`func (o *FileActionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCompletedAt

`func (o *FileActionResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *FileActionResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *FileActionResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetNumErrors

`func (o *FileActionResponse) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *FileActionResponse) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *FileActionResponse) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *FileActionResponse) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *FileActionResponse) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *FileActionResponse) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *FileActionResponse) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *FileActionResponse) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *FileActionResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FileActionResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FileActionResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *FileActionResponse) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FileActionResponse) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FileActionResponse) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FileActionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetErrors

`func (o *FileActionResponse) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FileActionResponse) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FileActionResponse) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FileActionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTaskId

`func (o *FileActionResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *FileActionResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *FileActionResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStatus

`func (o *FileActionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileActionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileActionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


