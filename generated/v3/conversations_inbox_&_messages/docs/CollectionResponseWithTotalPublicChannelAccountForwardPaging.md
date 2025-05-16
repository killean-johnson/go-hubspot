# CollectionResponseWithTotalPublicChannelAccountForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 
**Results** | [**[]PublicChannelAccount**](PublicChannelAccount.md) |  | 

## Methods

### NewCollectionResponseWithTotalPublicChannelAccountForwardPaging

`func NewCollectionResponseWithTotalPublicChannelAccountForwardPaging(total int32, results []PublicChannelAccount, ) *CollectionResponseWithTotalPublicChannelAccountForwardPaging`

NewCollectionResponseWithTotalPublicChannelAccountForwardPaging instantiates a new CollectionResponseWithTotalPublicChannelAccountForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalPublicChannelAccountForwardPagingWithDefaults

`func NewCollectionResponseWithTotalPublicChannelAccountForwardPagingWithDefaults() *CollectionResponseWithTotalPublicChannelAccountForwardPaging`

NewCollectionResponseWithTotalPublicChannelAccountForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalPublicChannelAccountForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) GetResults() []PublicChannelAccount`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) GetResultsOk() (*[]PublicChannelAccount, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalPublicChannelAccountForwardPaging) SetResults(v []PublicChannelAccount)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


