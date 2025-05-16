# EventVisibilityChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowInWorkflows** | Pointer to **bool** |  | [optional] 
**ShowInReporting** | Pointer to **bool** |  | [optional] 
**EventType** | **string** |  | 
**ShowInTimeline** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | **int64** |  | 

## Methods

### NewEventVisibilityChange

`func NewEventVisibilityChange(eventType string, updatedAt int64, ) *EventVisibilityChange`

NewEventVisibilityChange instantiates a new EventVisibilityChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventVisibilityChangeWithDefaults

`func NewEventVisibilityChangeWithDefaults() *EventVisibilityChange`

NewEventVisibilityChangeWithDefaults instantiates a new EventVisibilityChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowInWorkflows

`func (o *EventVisibilityChange) GetShowInWorkflows() bool`

GetShowInWorkflows returns the ShowInWorkflows field if non-nil, zero value otherwise.

### GetShowInWorkflowsOk

`func (o *EventVisibilityChange) GetShowInWorkflowsOk() (*bool, bool)`

GetShowInWorkflowsOk returns a tuple with the ShowInWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInWorkflows

`func (o *EventVisibilityChange) SetShowInWorkflows(v bool)`

SetShowInWorkflows sets ShowInWorkflows field to given value.

### HasShowInWorkflows

`func (o *EventVisibilityChange) HasShowInWorkflows() bool`

HasShowInWorkflows returns a boolean if a field has been set.

### GetShowInReporting

`func (o *EventVisibilityChange) GetShowInReporting() bool`

GetShowInReporting returns the ShowInReporting field if non-nil, zero value otherwise.

### GetShowInReportingOk

`func (o *EventVisibilityChange) GetShowInReportingOk() (*bool, bool)`

GetShowInReportingOk returns a tuple with the ShowInReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInReporting

`func (o *EventVisibilityChange) SetShowInReporting(v bool)`

SetShowInReporting sets ShowInReporting field to given value.

### HasShowInReporting

`func (o *EventVisibilityChange) HasShowInReporting() bool`

HasShowInReporting returns a boolean if a field has been set.

### GetEventType

`func (o *EventVisibilityChange) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventVisibilityChange) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventVisibilityChange) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetShowInTimeline

`func (o *EventVisibilityChange) GetShowInTimeline() bool`

GetShowInTimeline returns the ShowInTimeline field if non-nil, zero value otherwise.

### GetShowInTimelineOk

`func (o *EventVisibilityChange) GetShowInTimelineOk() (*bool, bool)`

GetShowInTimelineOk returns a tuple with the ShowInTimeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInTimeline

`func (o *EventVisibilityChange) SetShowInTimeline(v bool)`

SetShowInTimeline sets ShowInTimeline field to given value.

### HasShowInTimeline

`func (o *EventVisibilityChange) HasShowInTimeline() bool`

HasShowInTimeline returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EventVisibilityChange) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventVisibilityChange) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventVisibilityChange) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


