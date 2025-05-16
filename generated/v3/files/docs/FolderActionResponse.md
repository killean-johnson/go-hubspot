# FolderActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**Folder**](Folder.md) |  | [optional] 
**CompletedAt** | **time.Time** | When the requested changes have been completed. | 
**NumErrors** | Pointer to **int32** | Number of errors resulting from the requested changes. | [optional] 
**RequestedAt** | Pointer to **time.Time** | Timestamp representing when the task was requested. | [optional] 
**StartedAt** | **time.Time** | Timestamp representing when the task was started at. | 
**Links** | Pointer to **map[string]string** | Link to check the status of the task. | [optional] 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) | Detailed errors resulting from the task. | [optional] 
**TaskId** | **string** | ID of the task. | 
**Status** | **string** | Current status of the task. | 

## Methods

### NewFolderActionResponse

`func NewFolderActionResponse(completedAt time.Time, startedAt time.Time, taskId string, status string, ) *FolderActionResponse`

NewFolderActionResponse instantiates a new FolderActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderActionResponseWithDefaults

`func NewFolderActionResponseWithDefaults() *FolderActionResponse`

NewFolderActionResponseWithDefaults instantiates a new FolderActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *FolderActionResponse) GetResult() Folder`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FolderActionResponse) GetResultOk() (*Folder, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FolderActionResponse) SetResult(v Folder)`

SetResult sets Result field to given value.

### HasResult

`func (o *FolderActionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCompletedAt

`func (o *FolderActionResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *FolderActionResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *FolderActionResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetNumErrors

`func (o *FolderActionResponse) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *FolderActionResponse) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *FolderActionResponse) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *FolderActionResponse) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *FolderActionResponse) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *FolderActionResponse) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *FolderActionResponse) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *FolderActionResponse) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *FolderActionResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FolderActionResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FolderActionResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *FolderActionResponse) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FolderActionResponse) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FolderActionResponse) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FolderActionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetErrors

`func (o *FolderActionResponse) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FolderActionResponse) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FolderActionResponse) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FolderActionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTaskId

`func (o *FolderActionResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *FolderActionResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *FolderActionResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStatus

`func (o *FolderActionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FolderActionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FolderActionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


