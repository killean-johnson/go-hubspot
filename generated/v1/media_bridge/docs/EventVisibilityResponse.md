# EventVisibilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**VisibilitySettings** | [**[]EventVisibilityChange**](EventVisibilityChange.md) |  | 

## Methods

### NewEventVisibilityResponse

`func NewEventVisibilityResponse(createdAt time.Time, visibilitySettings []EventVisibilityChange, ) *EventVisibilityResponse`

NewEventVisibilityResponse instantiates a new EventVisibilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventVisibilityResponseWithDefaults

`func NewEventVisibilityResponseWithDefaults() *EventVisibilityResponse`

NewEventVisibilityResponseWithDefaults instantiates a new EventVisibilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EventVisibilityResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventVisibilityResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventVisibilityResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetVisibilitySettings

`func (o *EventVisibilityResponse) GetVisibilitySettings() []EventVisibilityChange`

GetVisibilitySettings returns the VisibilitySettings field if non-nil, zero value otherwise.

### GetVisibilitySettingsOk

`func (o *EventVisibilityResponse) GetVisibilitySettingsOk() (*[]EventVisibilityChange, bool)`

GetVisibilitySettingsOk returns a tuple with the VisibilitySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilitySettings

`func (o *EventVisibilityResponse) SetVisibilitySettings(v []EventVisibilityChange)`

SetVisibilitySettings sets VisibilitySettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


