# CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 
**Results** | [**[]PublicChannelIntegrationChannel**](PublicChannelIntegrationChannel.md) |  | 

## Methods

### NewCollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging

`func NewCollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging(total int32, results []PublicChannelIntegrationChannel, ) *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging`

NewCollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging instantiates a new CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalPublicChannelIntegrationChannelForwardPagingWithDefaults

`func NewCollectionResponseWithTotalPublicChannelIntegrationChannelForwardPagingWithDefaults() *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging`

NewCollectionResponseWithTotalPublicChannelIntegrationChannelForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) GetResults() []PublicChannelIntegrationChannel`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) GetResultsOk() (*[]PublicChannelIntegrationChannel, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) SetResults(v []PublicChannelIntegrationChannel)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


