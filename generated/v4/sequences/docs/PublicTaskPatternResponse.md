# PublicTaskPatternResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**TaskType** | **string** |  | 
**Notes** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**TaskPriority** | **string** |  | 
**Id** | **string** |  | 
**TemplateId** | Pointer to **int64** |  | [optional] 
**ThreadEmailToStepOrder** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPublicTaskPatternResponse

`func NewPublicTaskPatternResponse(createdAt time.Time, taskType string, taskPriority string, id string, updatedAt time.Time, ) *PublicTaskPatternResponse`

NewPublicTaskPatternResponse instantiates a new PublicTaskPatternResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTaskPatternResponseWithDefaults

`func NewPublicTaskPatternResponseWithDefaults() *PublicTaskPatternResponse`

NewPublicTaskPatternResponseWithDefaults instantiates a new PublicTaskPatternResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *PublicTaskPatternResponse) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *PublicTaskPatternResponse) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *PublicTaskPatternResponse) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *PublicTaskPatternResponse) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicTaskPatternResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicTaskPatternResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicTaskPatternResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTaskType

`func (o *PublicTaskPatternResponse) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *PublicTaskPatternResponse) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *PublicTaskPatternResponse) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.


### GetNotes

`func (o *PublicTaskPatternResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PublicTaskPatternResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PublicTaskPatternResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PublicTaskPatternResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSubject

`func (o *PublicTaskPatternResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PublicTaskPatternResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PublicTaskPatternResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PublicTaskPatternResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTaskPriority

`func (o *PublicTaskPatternResponse) GetTaskPriority() string`

GetTaskPriority returns the TaskPriority field if non-nil, zero value otherwise.

### GetTaskPriorityOk

`func (o *PublicTaskPatternResponse) GetTaskPriorityOk() (*string, bool)`

GetTaskPriorityOk returns a tuple with the TaskPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPriority

`func (o *PublicTaskPatternResponse) SetTaskPriority(v string)`

SetTaskPriority sets TaskPriority field to given value.


### GetId

`func (o *PublicTaskPatternResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicTaskPatternResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicTaskPatternResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *PublicTaskPatternResponse) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PublicTaskPatternResponse) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PublicTaskPatternResponse) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *PublicTaskPatternResponse) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetThreadEmailToStepOrder

`func (o *PublicTaskPatternResponse) GetThreadEmailToStepOrder() int32`

GetThreadEmailToStepOrder returns the ThreadEmailToStepOrder field if non-nil, zero value otherwise.

### GetThreadEmailToStepOrderOk

`func (o *PublicTaskPatternResponse) GetThreadEmailToStepOrderOk() (*int32, bool)`

GetThreadEmailToStepOrderOk returns a tuple with the ThreadEmailToStepOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadEmailToStepOrder

`func (o *PublicTaskPatternResponse) SetThreadEmailToStepOrder(v int32)`

SetThreadEmailToStepOrder sets ThreadEmailToStepOrder field to given value.

### HasThreadEmailToStepOrder

`func (o *PublicTaskPatternResponse) HasThreadEmailToStepOrder() bool`

HasThreadEmailToStepOrder returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicTaskPatternResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicTaskPatternResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicTaskPatternResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


