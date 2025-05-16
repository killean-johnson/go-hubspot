# CollectionResponseWithTotalContentFolderForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of content folders. | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 
**Results** | [**[]ContentFolder**](ContentFolder.md) | Collection of content folders. | 

## Methods

### NewCollectionResponseWithTotalContentFolderForwardPaging

`func NewCollectionResponseWithTotalContentFolderForwardPaging(total int32, results []ContentFolder, ) *CollectionResponseWithTotalContentFolderForwardPaging`

NewCollectionResponseWithTotalContentFolderForwardPaging instantiates a new CollectionResponseWithTotalContentFolderForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalContentFolderForwardPagingWithDefaults

`func NewCollectionResponseWithTotalContentFolderForwardPagingWithDefaults() *CollectionResponseWithTotalContentFolderForwardPaging`

NewCollectionResponseWithTotalContentFolderForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalContentFolderForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) GetResults() []ContentFolder`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) GetResultsOk() (*[]ContentFolder, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalContentFolderForwardPaging) SetResults(v []ContentFolder)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


