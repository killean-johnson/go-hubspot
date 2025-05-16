# CollectionResponseWithTotalVersionBlog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of blog versions. | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Results** | [**[]VersionBlog**](VersionBlog.md) | Collection of blog versions. | 

## Methods

### NewCollectionResponseWithTotalVersionBlog

`func NewCollectionResponseWithTotalVersionBlog(total int32, results []VersionBlog, ) *CollectionResponseWithTotalVersionBlog`

NewCollectionResponseWithTotalVersionBlog instantiates a new CollectionResponseWithTotalVersionBlog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalVersionBlogWithDefaults

`func NewCollectionResponseWithTotalVersionBlogWithDefaults() *CollectionResponseWithTotalVersionBlog`

NewCollectionResponseWithTotalVersionBlogWithDefaults instantiates a new CollectionResponseWithTotalVersionBlog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalVersionBlog) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalVersionBlog) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalVersionBlog) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalVersionBlog) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalVersionBlog) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalVersionBlog) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalVersionBlog) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalVersionBlog) GetResults() []VersionBlog`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalVersionBlog) GetResultsOk() (*[]VersionBlog, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalVersionBlog) SetResults(v []VersionBlog)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


