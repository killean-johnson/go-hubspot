# SubscriptionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypeId** | Pointer to **string** |  | [optional] 
**PropertyName** | Pointer to **string** | The internal name of the property to monitor for changes. Only applies when &#x60;eventType&#x60; is &#x60;propertyChange&#x60;. | [optional] 
**Active** | Pointer to **bool** | Determines if the subscription is active or paused. Defaults to false. | [optional] 
**EventType** | **string** | Type of event to listen for. Can be one of &#x60;create&#x60;, &#x60;delete&#x60;, &#x60;deletedForPrivacy&#x60;, or &#x60;propertyChange&#x60;. | 

## Methods

### NewSubscriptionCreateRequest

`func NewSubscriptionCreateRequest(eventType string, ) *SubscriptionCreateRequest`

NewSubscriptionCreateRequest instantiates a new SubscriptionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateRequestWithDefaults

`func NewSubscriptionCreateRequestWithDefaults() *SubscriptionCreateRequest`

NewSubscriptionCreateRequestWithDefaults instantiates a new SubscriptionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *SubscriptionCreateRequest) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *SubscriptionCreateRequest) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *SubscriptionCreateRequest) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.

### HasObjectTypeId

`func (o *SubscriptionCreateRequest) HasObjectTypeId() bool`

HasObjectTypeId returns a boolean if a field has been set.

### GetPropertyName

`func (o *SubscriptionCreateRequest) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *SubscriptionCreateRequest) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *SubscriptionCreateRequest) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *SubscriptionCreateRequest) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetActive

`func (o *SubscriptionCreateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SubscriptionCreateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SubscriptionCreateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SubscriptionCreateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEventType

`func (o *SubscriptionCreateRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SubscriptionCreateRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SubscriptionCreateRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


