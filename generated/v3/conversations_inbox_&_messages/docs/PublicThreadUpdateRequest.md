# PublicThreadUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | Whether this thread is archived. Set to false to restore the thread. | [optional] 
**Status** | Pointer to **string** | The thread&#39;s status: &#x60;OPEN&#x60; or &#x60;CLOSED&#x60;. | [optional] 

## Methods

### NewPublicThreadUpdateRequest

`func NewPublicThreadUpdateRequest() *PublicThreadUpdateRequest`

NewPublicThreadUpdateRequest instantiates a new PublicThreadUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicThreadUpdateRequestWithDefaults

`func NewPublicThreadUpdateRequestWithDefaults() *PublicThreadUpdateRequest`

NewPublicThreadUpdateRequestWithDefaults instantiates a new PublicThreadUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *PublicThreadUpdateRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PublicThreadUpdateRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PublicThreadUpdateRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PublicThreadUpdateRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetStatus

`func (o *PublicThreadUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicThreadUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicThreadUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicThreadUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


