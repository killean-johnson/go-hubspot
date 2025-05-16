# ThrottlingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConcurrentRequests** | **int32** | The maximum number of concurrent HTTP requests HubSpot will attempt to make to your app. | 

## Methods

### NewThrottlingSettings

`func NewThrottlingSettings(maxConcurrentRequests int32, ) *ThrottlingSettings`

NewThrottlingSettings instantiates a new ThrottlingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThrottlingSettingsWithDefaults

`func NewThrottlingSettingsWithDefaults() *ThrottlingSettings`

NewThrottlingSettingsWithDefaults instantiates a new ThrottlingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxConcurrentRequests

`func (o *ThrottlingSettings) GetMaxConcurrentRequests() int32`

GetMaxConcurrentRequests returns the MaxConcurrentRequests field if non-nil, zero value otherwise.

### GetMaxConcurrentRequestsOk

`func (o *ThrottlingSettings) GetMaxConcurrentRequestsOk() (*int32, bool)`

GetMaxConcurrentRequestsOk returns a tuple with the MaxConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentRequests

`func (o *ThrottlingSettings) SetMaxConcurrentRequests(v int32)`

SetMaxConcurrentRequests sets MaxConcurrentRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


