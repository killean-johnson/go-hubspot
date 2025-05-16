# PublicChannelIntegrationMessageUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusType** | **string** | Valid status are SENT, FAILED, and READ | 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicChannelIntegrationMessageUpdateRequest

`func NewPublicChannelIntegrationMessageUpdateRequest(statusType string, ) *PublicChannelIntegrationMessageUpdateRequest`

NewPublicChannelIntegrationMessageUpdateRequest instantiates a new PublicChannelIntegrationMessageUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicChannelIntegrationMessageUpdateRequestWithDefaults

`func NewPublicChannelIntegrationMessageUpdateRequestWithDefaults() *PublicChannelIntegrationMessageUpdateRequest`

NewPublicChannelIntegrationMessageUpdateRequestWithDefaults instantiates a new PublicChannelIntegrationMessageUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusType

`func (o *PublicChannelIntegrationMessageUpdateRequest) GetStatusType() string`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *PublicChannelIntegrationMessageUpdateRequest) GetStatusTypeOk() (*string, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *PublicChannelIntegrationMessageUpdateRequest) SetStatusType(v string)`

SetStatusType sets StatusType field to given value.


### GetErrorMessage

`func (o *PublicChannelIntegrationMessageUpdateRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PublicChannelIntegrationMessageUpdateRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PublicChannelIntegrationMessageUpdateRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PublicChannelIntegrationMessageUpdateRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


