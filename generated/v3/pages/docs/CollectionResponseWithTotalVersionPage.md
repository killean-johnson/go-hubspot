# CollectionResponseWithTotalVersionPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of page versions. | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Results** | [**[]VersionPage**](VersionPage.md) | Collection of page versions. | 

## Methods

### NewCollectionResponseWithTotalVersionPage

`func NewCollectionResponseWithTotalVersionPage(total int32, results []VersionPage, ) *CollectionResponseWithTotalVersionPage`

NewCollectionResponseWithTotalVersionPage instantiates a new CollectionResponseWithTotalVersionPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalVersionPageWithDefaults

`func NewCollectionResponseWithTotalVersionPageWithDefaults() *CollectionResponseWithTotalVersionPage`

NewCollectionResponseWithTotalVersionPageWithDefaults instantiates a new CollectionResponseWithTotalVersionPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalVersionPage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalVersionPage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalVersionPage) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalVersionPage) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalVersionPage) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalVersionPage) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalVersionPage) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalVersionPage) GetResults() []VersionPage`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalVersionPage) GetResultsOk() (*[]VersionPage, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalVersionPage) SetResults(v []VersionPage)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


