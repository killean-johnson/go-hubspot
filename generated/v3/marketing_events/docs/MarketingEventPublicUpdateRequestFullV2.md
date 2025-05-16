# MarketingEventPublicUpdateRequestFullV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | Pointer to **time.Time** |  | [optional] 
**CustomProperties** | [**[]PropertyValue**](PropertyValue.md) |  | 
**EventCancelled** | Pointer to **bool** |  | [optional] 
**EventOrganizer** | Pointer to **string** |  | [optional] 
**EventUrl** | Pointer to **string** |  | [optional] 
**EventDescription** | Pointer to **string** |  | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**EndDateTime** | Pointer to **time.Time** |  | [optional] 
**ObjectId** | **string** |  | 

## Methods

### NewMarketingEventPublicUpdateRequestFullV2

`func NewMarketingEventPublicUpdateRequestFullV2(customProperties []PropertyValue, objectId string, ) *MarketingEventPublicUpdateRequestFullV2`

NewMarketingEventPublicUpdateRequestFullV2 instantiates a new MarketingEventPublicUpdateRequestFullV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEventPublicUpdateRequestFullV2WithDefaults

`func NewMarketingEventPublicUpdateRequestFullV2WithDefaults() *MarketingEventPublicUpdateRequestFullV2`

NewMarketingEventPublicUpdateRequestFullV2WithDefaults instantiates a new MarketingEventPublicUpdateRequestFullV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTime

`func (o *MarketingEventPublicUpdateRequestFullV2) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MarketingEventPublicUpdateRequestFullV2) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MarketingEventPublicUpdateRequestFullV2) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCustomProperties

`func (o *MarketingEventPublicUpdateRequestFullV2) GetCustomProperties() []PropertyValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetCustomPropertiesOk() (*[]PropertyValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *MarketingEventPublicUpdateRequestFullV2) SetCustomProperties(v []PropertyValue)`

SetCustomProperties sets CustomProperties field to given value.


### GetEventCancelled

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventCancelled() bool`

GetEventCancelled returns the EventCancelled field if non-nil, zero value otherwise.

### GetEventCancelledOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventCancelledOk() (*bool, bool)`

GetEventCancelledOk returns a tuple with the EventCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCancelled

`func (o *MarketingEventPublicUpdateRequestFullV2) SetEventCancelled(v bool)`

SetEventCancelled sets EventCancelled field to given value.

### HasEventCancelled

`func (o *MarketingEventPublicUpdateRequestFullV2) HasEventCancelled() bool`

HasEventCancelled returns a boolean if a field has been set.

### GetEventOrganizer

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventOrganizer() string`

GetEventOrganizer returns the EventOrganizer field if non-nil, zero value otherwise.

### GetEventOrganizerOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventOrganizerOk() (*string, bool)`

GetEventOrganizerOk returns a tuple with the EventOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOrganizer

`func (o *MarketingEventPublicUpdateRequestFullV2) SetEventOrganizer(v string)`

SetEventOrganizer sets EventOrganizer field to given value.

### HasEventOrganizer

`func (o *MarketingEventPublicUpdateRequestFullV2) HasEventOrganizer() bool`

HasEventOrganizer returns a boolean if a field has been set.

### GetEventUrl

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventUrl() string`

GetEventUrl returns the EventUrl field if non-nil, zero value otherwise.

### GetEventUrlOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventUrlOk() (*string, bool)`

GetEventUrlOk returns a tuple with the EventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUrl

`func (o *MarketingEventPublicUpdateRequestFullV2) SetEventUrl(v string)`

SetEventUrl sets EventUrl field to given value.

### HasEventUrl

`func (o *MarketingEventPublicUpdateRequestFullV2) HasEventUrl() bool`

HasEventUrl returns a boolean if a field has been set.

### GetEventDescription

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventDescription() string`

GetEventDescription returns the EventDescription field if non-nil, zero value otherwise.

### GetEventDescriptionOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventDescriptionOk() (*string, bool)`

GetEventDescriptionOk returns a tuple with the EventDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDescription

`func (o *MarketingEventPublicUpdateRequestFullV2) SetEventDescription(v string)`

SetEventDescription sets EventDescription field to given value.

### HasEventDescription

`func (o *MarketingEventPublicUpdateRequestFullV2) HasEventDescription() bool`

HasEventDescription returns a boolean if a field has been set.

### GetEventName

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *MarketingEventPublicUpdateRequestFullV2) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *MarketingEventPublicUpdateRequestFullV2) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventType

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketingEventPublicUpdateRequestFullV2) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MarketingEventPublicUpdateRequestFullV2) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MarketingEventPublicUpdateRequestFullV2) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MarketingEventPublicUpdateRequestFullV2) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetObjectId

`func (o *MarketingEventPublicUpdateRequestFullV2) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MarketingEventPublicUpdateRequestFullV2) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MarketingEventPublicUpdateRequestFullV2) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


