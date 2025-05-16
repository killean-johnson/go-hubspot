# BehavioralEventHttpCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OccurredAt** | Pointer to **time.Time** | The time when this event occurred. If this isn&#39;t set, the current time will be used. | [optional] 
**EventName** | **string** | The internal name of the event (&#x60;pe&lt;portalID&gt;_eventName&#x60;). Can be retrieved through the [event definitions API](https://developers.hubspot.com/docs/reference/api/analytics-and-events/custom-events/custom-event-definitions#get-%2Fevents%2Fv3%2Fevent-definitions) or in [HubSpot&#39;s UI](https://knowledge.hubspot.com/reports/create-custom-behavioral-events-with-the-code-wizard#find-internal-name).  | 
**Utk** | Pointer to **string** | The visitor&#39;s usertoken. Used for associating the event data with a CRM record. | [optional] 
**Uuid** | Pointer to **string** | Include a universally unique identifier to assign a unique ID to the event completion. Can be useful for matching data between HubSpot and other external systems. | [optional] 
**Email** | Pointer to **string** | The visitor&#39;s email address. Used for associating the event data with a CRM record. | [optional] 
**Properties** | Pointer to **map[string]string** | The event properties to update. Takes the format of key-value pairs (property internal name and property value). Learn more about [HubSpot&#39;s default event properties](https://developers.hubspot.com/docs/guides/api/analytics-and-events/custom-events/custom-event-definitions#hubspot-s-default-event-properties). | [optional] 
**ObjectId** | Pointer to **string** | The ID of the object that completed the event (e.g., contact ID or visitor ID). | [optional] 

## Methods

### NewBehavioralEventHttpCompletionRequest

`func NewBehavioralEventHttpCompletionRequest(eventName string, ) *BehavioralEventHttpCompletionRequest`

NewBehavioralEventHttpCompletionRequest instantiates a new BehavioralEventHttpCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehavioralEventHttpCompletionRequestWithDefaults

`func NewBehavioralEventHttpCompletionRequestWithDefaults() *BehavioralEventHttpCompletionRequest`

NewBehavioralEventHttpCompletionRequestWithDefaults instantiates a new BehavioralEventHttpCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurredAt

`func (o *BehavioralEventHttpCompletionRequest) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BehavioralEventHttpCompletionRequest) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BehavioralEventHttpCompletionRequest) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *BehavioralEventHttpCompletionRequest) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetEventName

`func (o *BehavioralEventHttpCompletionRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *BehavioralEventHttpCompletionRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *BehavioralEventHttpCompletionRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetUtk

`func (o *BehavioralEventHttpCompletionRequest) GetUtk() string`

GetUtk returns the Utk field if non-nil, zero value otherwise.

### GetUtkOk

`func (o *BehavioralEventHttpCompletionRequest) GetUtkOk() (*string, bool)`

GetUtkOk returns a tuple with the Utk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtk

`func (o *BehavioralEventHttpCompletionRequest) SetUtk(v string)`

SetUtk sets Utk field to given value.

### HasUtk

`func (o *BehavioralEventHttpCompletionRequest) HasUtk() bool`

HasUtk returns a boolean if a field has been set.

### GetUuid

`func (o *BehavioralEventHttpCompletionRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BehavioralEventHttpCompletionRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BehavioralEventHttpCompletionRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BehavioralEventHttpCompletionRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetEmail

`func (o *BehavioralEventHttpCompletionRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BehavioralEventHttpCompletionRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BehavioralEventHttpCompletionRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BehavioralEventHttpCompletionRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetProperties

`func (o *BehavioralEventHttpCompletionRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BehavioralEventHttpCompletionRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BehavioralEventHttpCompletionRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BehavioralEventHttpCompletionRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetObjectId

`func (o *BehavioralEventHttpCompletionRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BehavioralEventHttpCompletionRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BehavioralEventHttpCompletionRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BehavioralEventHttpCompletionRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


