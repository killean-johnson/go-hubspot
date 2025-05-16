# CollectionResponseWithTotalPublicInboxForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 
**Results** | [**[]PublicInbox**](PublicInbox.md) |  | 

## Methods

### NewCollectionResponseWithTotalPublicInboxForwardPaging

`func NewCollectionResponseWithTotalPublicInboxForwardPaging(total int32, results []PublicInbox, ) *CollectionResponseWithTotalPublicInboxForwardPaging`

NewCollectionResponseWithTotalPublicInboxForwardPaging instantiates a new CollectionResponseWithTotalPublicInboxForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalPublicInboxForwardPagingWithDefaults

`func NewCollectionResponseWithTotalPublicInboxForwardPagingWithDefaults() *CollectionResponseWithTotalPublicInboxForwardPaging`

NewCollectionResponseWithTotalPublicInboxForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalPublicInboxForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) GetResults() []PublicInbox`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) GetResultsOk() (*[]PublicInbox, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalPublicInboxForwardPaging) SetResults(v []PublicInbox)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


