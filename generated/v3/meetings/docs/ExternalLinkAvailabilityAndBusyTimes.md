# ExternalLinkAvailabilityAndBusyTimes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkAvailability** | Pointer to [**ExternalLinkAvailability**](ExternalLinkAvailability.md) |  | [optional] 
**AllUsersBusyTimes** | [**[]ExternalUserBusyTimes**](ExternalUserBusyTimes.md) |  | 

## Methods

### NewExternalLinkAvailabilityAndBusyTimes

`func NewExternalLinkAvailabilityAndBusyTimes(allUsersBusyTimes []ExternalUserBusyTimes, ) *ExternalLinkAvailabilityAndBusyTimes`

NewExternalLinkAvailabilityAndBusyTimes instantiates a new ExternalLinkAvailabilityAndBusyTimes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalLinkAvailabilityAndBusyTimesWithDefaults

`func NewExternalLinkAvailabilityAndBusyTimesWithDefaults() *ExternalLinkAvailabilityAndBusyTimes`

NewExternalLinkAvailabilityAndBusyTimesWithDefaults instantiates a new ExternalLinkAvailabilityAndBusyTimes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkAvailability

`func (o *ExternalLinkAvailabilityAndBusyTimes) GetLinkAvailability() ExternalLinkAvailability`

GetLinkAvailability returns the LinkAvailability field if non-nil, zero value otherwise.

### GetLinkAvailabilityOk

`func (o *ExternalLinkAvailabilityAndBusyTimes) GetLinkAvailabilityOk() (*ExternalLinkAvailability, bool)`

GetLinkAvailabilityOk returns a tuple with the LinkAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAvailability

`func (o *ExternalLinkAvailabilityAndBusyTimes) SetLinkAvailability(v ExternalLinkAvailability)`

SetLinkAvailability sets LinkAvailability field to given value.

### HasLinkAvailability

`func (o *ExternalLinkAvailabilityAndBusyTimes) HasLinkAvailability() bool`

HasLinkAvailability returns a boolean if a field has been set.

### GetAllUsersBusyTimes

`func (o *ExternalLinkAvailabilityAndBusyTimes) GetAllUsersBusyTimes() []ExternalUserBusyTimes`

GetAllUsersBusyTimes returns the AllUsersBusyTimes field if non-nil, zero value otherwise.

### GetAllUsersBusyTimesOk

`func (o *ExternalLinkAvailabilityAndBusyTimes) GetAllUsersBusyTimesOk() (*[]ExternalUserBusyTimes, bool)`

GetAllUsersBusyTimesOk returns a tuple with the AllUsersBusyTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUsersBusyTimes

`func (o *ExternalLinkAvailabilityAndBusyTimes) SetAllUsersBusyTimes(v []ExternalUserBusyTimes)`

SetAllUsersBusyTimes sets AllUsersBusyTimes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


