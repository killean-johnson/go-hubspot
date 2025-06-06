# CollectionResponseWithTotalVersionContentFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of content folder versions. | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Results** | [**[]VersionContentFolder**](VersionContentFolder.md) | Collection of content folder versions. | 

## Methods

### NewCollectionResponseWithTotalVersionContentFolder

`func NewCollectionResponseWithTotalVersionContentFolder(total int32, results []VersionContentFolder, ) *CollectionResponseWithTotalVersionContentFolder`

NewCollectionResponseWithTotalVersionContentFolder instantiates a new CollectionResponseWithTotalVersionContentFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalVersionContentFolderWithDefaults

`func NewCollectionResponseWithTotalVersionContentFolderWithDefaults() *CollectionResponseWithTotalVersionContentFolder`

NewCollectionResponseWithTotalVersionContentFolderWithDefaults instantiates a new CollectionResponseWithTotalVersionContentFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalVersionContentFolder) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalVersionContentFolder) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalVersionContentFolder) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalVersionContentFolder) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalVersionContentFolder) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalVersionContentFolder) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalVersionContentFolder) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalVersionContentFolder) GetResults() []VersionContentFolder`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalVersionContentFolder) GetResultsOk() (*[]VersionContentFolder, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalVersionContentFolder) SetResults(v []VersionContentFolder)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


