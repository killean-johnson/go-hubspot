# MarketingEventPublicReadResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registrants** | Pointer to **int32** |  | [optional] 
**EventOrganizer** | Pointer to **string** |  | [optional] 
**EventUrl** | Pointer to **string** |  | [optional] 
**Attendees** | Pointer to **int32** |  | [optional] 
**AppInfo** | Pointer to [**AppInfo**](AppInfo.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**EventCompleted** | Pointer to **bool** |  | [optional] 
**EndDateTime** | Pointer to **time.Time** |  | [optional] 
**NoShows** | Pointer to **int32** |  | [optional] 
**Cancellations** | Pointer to **int32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**StartDateTime** | Pointer to **time.Time** |  | [optional] 
**CustomProperties** | [**[]CrmPropertyWrapper**](CrmPropertyWrapper.md) |  | 
**EventCancelled** | Pointer to **bool** |  | [optional] 
**ExternalEventId** | Pointer to **string** |  | [optional] 
**EventStatus** | Pointer to **string** |  | [optional] 
**EventDescription** | Pointer to **string** |  | [optional] 
**EventName** | **string** |  | 
**ObjectId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewMarketingEventPublicReadResponseV2

`func NewMarketingEventPublicReadResponseV2(createdAt time.Time, customProperties []CrmPropertyWrapper, eventName string, objectId string, updatedAt time.Time, ) *MarketingEventPublicReadResponseV2`

NewMarketingEventPublicReadResponseV2 instantiates a new MarketingEventPublicReadResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEventPublicReadResponseV2WithDefaults

`func NewMarketingEventPublicReadResponseV2WithDefaults() *MarketingEventPublicReadResponseV2`

NewMarketingEventPublicReadResponseV2WithDefaults instantiates a new MarketingEventPublicReadResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrants

`func (o *MarketingEventPublicReadResponseV2) GetRegistrants() int32`

GetRegistrants returns the Registrants field if non-nil, zero value otherwise.

### GetRegistrantsOk

`func (o *MarketingEventPublicReadResponseV2) GetRegistrantsOk() (*int32, bool)`

GetRegistrantsOk returns a tuple with the Registrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrants

`func (o *MarketingEventPublicReadResponseV2) SetRegistrants(v int32)`

SetRegistrants sets Registrants field to given value.

### HasRegistrants

`func (o *MarketingEventPublicReadResponseV2) HasRegistrants() bool`

HasRegistrants returns a boolean if a field has been set.

### GetEventOrganizer

`func (o *MarketingEventPublicReadResponseV2) GetEventOrganizer() string`

GetEventOrganizer returns the EventOrganizer field if non-nil, zero value otherwise.

### GetEventOrganizerOk

`func (o *MarketingEventPublicReadResponseV2) GetEventOrganizerOk() (*string, bool)`

GetEventOrganizerOk returns a tuple with the EventOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOrganizer

`func (o *MarketingEventPublicReadResponseV2) SetEventOrganizer(v string)`

SetEventOrganizer sets EventOrganizer field to given value.

### HasEventOrganizer

`func (o *MarketingEventPublicReadResponseV2) HasEventOrganizer() bool`

HasEventOrganizer returns a boolean if a field has been set.

### GetEventUrl

`func (o *MarketingEventPublicReadResponseV2) GetEventUrl() string`

GetEventUrl returns the EventUrl field if non-nil, zero value otherwise.

### GetEventUrlOk

`func (o *MarketingEventPublicReadResponseV2) GetEventUrlOk() (*string, bool)`

GetEventUrlOk returns a tuple with the EventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUrl

`func (o *MarketingEventPublicReadResponseV2) SetEventUrl(v string)`

SetEventUrl sets EventUrl field to given value.

### HasEventUrl

`func (o *MarketingEventPublicReadResponseV2) HasEventUrl() bool`

HasEventUrl returns a boolean if a field has been set.

### GetAttendees

`func (o *MarketingEventPublicReadResponseV2) GetAttendees() int32`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *MarketingEventPublicReadResponseV2) GetAttendeesOk() (*int32, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *MarketingEventPublicReadResponseV2) SetAttendees(v int32)`

SetAttendees sets Attendees field to given value.

### HasAttendees

`func (o *MarketingEventPublicReadResponseV2) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### GetAppInfo

`func (o *MarketingEventPublicReadResponseV2) GetAppInfo() AppInfo`

GetAppInfo returns the AppInfo field if non-nil, zero value otherwise.

### GetAppInfoOk

`func (o *MarketingEventPublicReadResponseV2) GetAppInfoOk() (*AppInfo, bool)`

GetAppInfoOk returns a tuple with the AppInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInfo

`func (o *MarketingEventPublicReadResponseV2) SetAppInfo(v AppInfo)`

SetAppInfo sets AppInfo field to given value.

### HasAppInfo

`func (o *MarketingEventPublicReadResponseV2) HasAppInfo() bool`

HasAppInfo returns a boolean if a field has been set.

### GetEventType

`func (o *MarketingEventPublicReadResponseV2) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketingEventPublicReadResponseV2) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketingEventPublicReadResponseV2) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MarketingEventPublicReadResponseV2) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventCompleted

`func (o *MarketingEventPublicReadResponseV2) GetEventCompleted() bool`

GetEventCompleted returns the EventCompleted field if non-nil, zero value otherwise.

### GetEventCompletedOk

`func (o *MarketingEventPublicReadResponseV2) GetEventCompletedOk() (*bool, bool)`

GetEventCompletedOk returns a tuple with the EventCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCompleted

`func (o *MarketingEventPublicReadResponseV2) SetEventCompleted(v bool)`

SetEventCompleted sets EventCompleted field to given value.

### HasEventCompleted

`func (o *MarketingEventPublicReadResponseV2) HasEventCompleted() bool`

HasEventCompleted returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MarketingEventPublicReadResponseV2) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MarketingEventPublicReadResponseV2) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MarketingEventPublicReadResponseV2) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MarketingEventPublicReadResponseV2) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetNoShows

`func (o *MarketingEventPublicReadResponseV2) GetNoShows() int32`

GetNoShows returns the NoShows field if non-nil, zero value otherwise.

### GetNoShowsOk

`func (o *MarketingEventPublicReadResponseV2) GetNoShowsOk() (*int32, bool)`

GetNoShowsOk returns a tuple with the NoShows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShows

`func (o *MarketingEventPublicReadResponseV2) SetNoShows(v int32)`

SetNoShows sets NoShows field to given value.

### HasNoShows

`func (o *MarketingEventPublicReadResponseV2) HasNoShows() bool`

HasNoShows returns a boolean if a field has been set.

### GetCancellations

`func (o *MarketingEventPublicReadResponseV2) GetCancellations() int32`

GetCancellations returns the Cancellations field if non-nil, zero value otherwise.

### GetCancellationsOk

`func (o *MarketingEventPublicReadResponseV2) GetCancellationsOk() (*int32, bool)`

GetCancellationsOk returns a tuple with the Cancellations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellations

`func (o *MarketingEventPublicReadResponseV2) SetCancellations(v int32)`

SetCancellations sets Cancellations field to given value.

### HasCancellations

`func (o *MarketingEventPublicReadResponseV2) HasCancellations() bool`

HasCancellations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingEventPublicReadResponseV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingEventPublicReadResponseV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingEventPublicReadResponseV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartDateTime

`func (o *MarketingEventPublicReadResponseV2) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MarketingEventPublicReadResponseV2) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MarketingEventPublicReadResponseV2) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MarketingEventPublicReadResponseV2) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCustomProperties

`func (o *MarketingEventPublicReadResponseV2) GetCustomProperties() []CrmPropertyWrapper`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *MarketingEventPublicReadResponseV2) GetCustomPropertiesOk() (*[]CrmPropertyWrapper, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *MarketingEventPublicReadResponseV2) SetCustomProperties(v []CrmPropertyWrapper)`

SetCustomProperties sets CustomProperties field to given value.


### GetEventCancelled

`func (o *MarketingEventPublicReadResponseV2) GetEventCancelled() bool`

GetEventCancelled returns the EventCancelled field if non-nil, zero value otherwise.

### GetEventCancelledOk

`func (o *MarketingEventPublicReadResponseV2) GetEventCancelledOk() (*bool, bool)`

GetEventCancelledOk returns a tuple with the EventCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCancelled

`func (o *MarketingEventPublicReadResponseV2) SetEventCancelled(v bool)`

SetEventCancelled sets EventCancelled field to given value.

### HasEventCancelled

`func (o *MarketingEventPublicReadResponseV2) HasEventCancelled() bool`

HasEventCancelled returns a boolean if a field has been set.

### GetExternalEventId

`func (o *MarketingEventPublicReadResponseV2) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *MarketingEventPublicReadResponseV2) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *MarketingEventPublicReadResponseV2) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.

### HasExternalEventId

`func (o *MarketingEventPublicReadResponseV2) HasExternalEventId() bool`

HasExternalEventId returns a boolean if a field has been set.

### GetEventStatus

`func (o *MarketingEventPublicReadResponseV2) GetEventStatus() string`

GetEventStatus returns the EventStatus field if non-nil, zero value otherwise.

### GetEventStatusOk

`func (o *MarketingEventPublicReadResponseV2) GetEventStatusOk() (*string, bool)`

GetEventStatusOk returns a tuple with the EventStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStatus

`func (o *MarketingEventPublicReadResponseV2) SetEventStatus(v string)`

SetEventStatus sets EventStatus field to given value.

### HasEventStatus

`func (o *MarketingEventPublicReadResponseV2) HasEventStatus() bool`

HasEventStatus returns a boolean if a field has been set.

### GetEventDescription

`func (o *MarketingEventPublicReadResponseV2) GetEventDescription() string`

GetEventDescription returns the EventDescription field if non-nil, zero value otherwise.

### GetEventDescriptionOk

`func (o *MarketingEventPublicReadResponseV2) GetEventDescriptionOk() (*string, bool)`

GetEventDescriptionOk returns a tuple with the EventDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDescription

`func (o *MarketingEventPublicReadResponseV2) SetEventDescription(v string)`

SetEventDescription sets EventDescription field to given value.

### HasEventDescription

`func (o *MarketingEventPublicReadResponseV2) HasEventDescription() bool`

HasEventDescription returns a boolean if a field has been set.

### GetEventName

`func (o *MarketingEventPublicReadResponseV2) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *MarketingEventPublicReadResponseV2) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *MarketingEventPublicReadResponseV2) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetObjectId

`func (o *MarketingEventPublicReadResponseV2) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MarketingEventPublicReadResponseV2) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MarketingEventPublicReadResponseV2) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetUpdatedAt

`func (o *MarketingEventPublicReadResponseV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketingEventPublicReadResponseV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketingEventPublicReadResponseV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


