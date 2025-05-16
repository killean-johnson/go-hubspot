# MarketingEventPublicDefaultResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventOrganizer** | Pointer to **string** |  | [optional] 
**EventUrl** | Pointer to **string** |  | [optional] 
**AppInfo** | Pointer to [**AppInfo**](AppInfo.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**EventCompleted** | Pointer to **bool** |  | [optional] 
**EndDateTime** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**StartDateTime** | Pointer to **time.Time** |  | [optional] 
**CustomProperties** | [**[]CrmPropertyWrapper**](CrmPropertyWrapper.md) |  | 
**EventCancelled** | Pointer to **bool** |  | [optional] 
**EventDescription** | Pointer to **string** |  | [optional] 
**EventName** | **string** |  | 
**ObjectId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewMarketingEventPublicDefaultResponseV2

`func NewMarketingEventPublicDefaultResponseV2(createdAt time.Time, customProperties []CrmPropertyWrapper, eventName string, objectId string, updatedAt time.Time, ) *MarketingEventPublicDefaultResponseV2`

NewMarketingEventPublicDefaultResponseV2 instantiates a new MarketingEventPublicDefaultResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEventPublicDefaultResponseV2WithDefaults

`func NewMarketingEventPublicDefaultResponseV2WithDefaults() *MarketingEventPublicDefaultResponseV2`

NewMarketingEventPublicDefaultResponseV2WithDefaults instantiates a new MarketingEventPublicDefaultResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventOrganizer

`func (o *MarketingEventPublicDefaultResponseV2) GetEventOrganizer() string`

GetEventOrganizer returns the EventOrganizer field if non-nil, zero value otherwise.

### GetEventOrganizerOk

`func (o *MarketingEventPublicDefaultResponseV2) GetEventOrganizerOk() (*string, bool)`

GetEventOrganizerOk returns a tuple with the EventOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOrganizer

`func (o *MarketingEventPublicDefaultResponseV2) SetEventOrganizer(v string)`

SetEventOrganizer sets EventOrganizer field to given value.

### HasEventOrganizer

`func (o *MarketingEventPublicDefaultResponseV2) HasEventOrganizer() bool`

HasEventOrganizer returns a boolean if a field has been set.

### GetEventUrl

`func (o *MarketingEventPublicDefaultResponseV2) GetEventUrl() string`

GetEventUrl returns the EventUrl field if non-nil, zero value otherwise.

### GetEventUrlOk

`func (o *MarketingEventPublicDefaultResponseV2) GetEventUrlOk() (*string, bool)`

GetEventUrlOk returns a tuple with the EventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUrl

`func (o *MarketingEventPublicDefaultResponseV2) SetEventUrl(v string)`

SetEventUrl sets EventUrl field to given value.

### HasEventUrl

`func (o *MarketingEventPublicDefaultResponseV2) HasEventUrl() bool`

HasEventUrl returns a boolean if a field has been set.

### GetAppInfo

`func (o *MarketingEventPublicDefaultResponseV2) GetAppInfo() AppInfo`

GetAppInfo returns the AppInfo field if non-nil, zero value otherwise.

### GetAppInfoOk

`func (o *MarketingEventPublicDefaultResponseV2) GetAppInfoOk() (*AppInfo, bool)`

GetAppInfoOk returns a tuple with the AppInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInfo

`func (o *MarketingEventPublicDefaultResponseV2) SetAppInfo(v AppInfo)`

SetAppInfo sets AppInfo field to given value.

### HasAppInfo

`func (o *MarketingEventPublicDefaultResponseV2) HasAppInfo() bool`

HasAppInfo returns a boolean if a field has been set.

### GetEventType

`func (o *MarketingEventPublicDefaultResponseV2) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketingEventPublicDefaultResponseV2) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketingEventPublicDefaultResponseV2) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MarketingEventPublicDefaultResponseV2) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventCompleted

`func (o *MarketingEventPublicDefaultResponseV2) GetEventCompleted() bool`

GetEventCompleted returns the EventCompleted field if non-nil, zero value otherwise.

### GetEventCompletedOk

`func (o *MarketingEventPublicDefaultResponseV2) GetEventCompletedOk() (*bool, bool)`

GetEventCompletedOk returns a tuple with the EventCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCompleted

`func (o *MarketingEventPublicDefaultResponseV2) SetEventCompleted(v bool)`

SetEventCompleted sets EventCompleted field to given value.

### HasEventCompleted

`func (o *MarketingEventPublicDefaultResponseV2) HasEventCompleted() bool`

HasEventCompleted returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MarketingEventPublicDefaultResponseV2) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MarketingEventPublicDefaultResponseV2) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MarketingEventPublicDefaultResponseV2) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MarketingEventPublicDefaultResponseV2) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingEventPublicDefaultResponseV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingEventPublicDefaultResponseV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingEventPublicDefaultResponseV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartDateTime

`func (o *MarketingEventPublicDefaultResponseV2) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MarketingEventPublicDefaultResponseV2) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MarketingEventPublicDefaultResponseV2) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MarketingEventPublicDefaultResponseV2) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCustomProperties

`func (o *MarketingEventPublicDefaultResponseV2) GetCustomProperties() []CrmPropertyWrapper`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *MarketingEventPublicDefaultResponseV2) GetCustomPropertiesOk() (*[]CrmPropertyWrapper, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *MarketingEventPublicDefaultResponseV2) SetCustomProperties(v []CrmPropertyWrapper)`

SetCustomProperties sets CustomProperties field to given value.


### GetEventCancelled

`func (o *MarketingEventPublicDefaultResponseV2) GetEventCancelled() bool`

GetEventCancelled returns the EventCancelled field if non-nil, zero value otherwise.

### GetEventCancelledOk

`func (o *MarketingEventPublicDefaultResponseV2) GetEventCancelledOk() (*bool, bool)`

GetEventCancelledOk returns a tuple with the EventCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCancelled

`func (o *MarketingEventPublicDefaultResponseV2) SetEventCancelled(v bool)`

SetEventCancelled sets EventCancelled field to given value.

### HasEventCancelled

`func (o *MarketingEventPublicDefaultResponseV2) HasEventCancelled() bool`

HasEventCancelled returns a boolean if a field has been set.

### GetEventDescription

`func (o *MarketingEventPublicDefaultResponseV2) GetEventDescription() string`

GetEventDescription returns the EventDescription field if non-nil, zero value otherwise.

### GetEventDescriptionOk

`func (o *MarketingEventPublicDefaultResponseV2) GetEventDescriptionOk() (*string, bool)`

GetEventDescriptionOk returns a tuple with the EventDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDescription

`func (o *MarketingEventPublicDefaultResponseV2) SetEventDescription(v string)`

SetEventDescription sets EventDescription field to given value.

### HasEventDescription

`func (o *MarketingEventPublicDefaultResponseV2) HasEventDescription() bool`

HasEventDescription returns a boolean if a field has been set.

### GetEventName

`func (o *MarketingEventPublicDefaultResponseV2) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *MarketingEventPublicDefaultResponseV2) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *MarketingEventPublicDefaultResponseV2) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetObjectId

`func (o *MarketingEventPublicDefaultResponseV2) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MarketingEventPublicDefaultResponseV2) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MarketingEventPublicDefaultResponseV2) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetUpdatedAt

`func (o *MarketingEventPublicDefaultResponseV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketingEventPublicDefaultResponseV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketingEventPublicDefaultResponseV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


