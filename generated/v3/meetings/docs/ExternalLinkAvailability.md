# ExternalLinkAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**LinkAvailabilityByDuration** | [**map[string]ExternalLinkAvailabilityForDuration**](ExternalLinkAvailabilityForDuration.md) |  | 

## Methods

### NewExternalLinkAvailability

`func NewExternalLinkAvailability(hasMore bool, linkAvailabilityByDuration map[string]ExternalLinkAvailabilityForDuration, ) *ExternalLinkAvailability`

NewExternalLinkAvailability instantiates a new ExternalLinkAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalLinkAvailabilityWithDefaults

`func NewExternalLinkAvailabilityWithDefaults() *ExternalLinkAvailability`

NewExternalLinkAvailabilityWithDefaults instantiates a new ExternalLinkAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ExternalLinkAvailability) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ExternalLinkAvailability) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ExternalLinkAvailability) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetLinkAvailabilityByDuration

`func (o *ExternalLinkAvailability) GetLinkAvailabilityByDuration() map[string]ExternalLinkAvailabilityForDuration`

GetLinkAvailabilityByDuration returns the LinkAvailabilityByDuration field if non-nil, zero value otherwise.

### GetLinkAvailabilityByDurationOk

`func (o *ExternalLinkAvailability) GetLinkAvailabilityByDurationOk() (*map[string]ExternalLinkAvailabilityForDuration, bool)`

GetLinkAvailabilityByDurationOk returns a tuple with the LinkAvailabilityByDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAvailabilityByDuration

`func (o *ExternalLinkAvailability) SetLinkAvailabilityByDuration(v map[string]ExternalLinkAvailabilityForDuration)`

SetLinkAvailabilityByDuration sets LinkAvailabilityByDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


