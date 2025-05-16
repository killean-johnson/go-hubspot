# CollectionResponseWithTotalVersionPublicEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of content emails. | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Results** | [**[]VersionPublicEmail**](VersionPublicEmail.md) | Collection of emails. | 

## Methods

### NewCollectionResponseWithTotalVersionPublicEmail

`func NewCollectionResponseWithTotalVersionPublicEmail(total int32, results []VersionPublicEmail, ) *CollectionResponseWithTotalVersionPublicEmail`

NewCollectionResponseWithTotalVersionPublicEmail instantiates a new CollectionResponseWithTotalVersionPublicEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalVersionPublicEmailWithDefaults

`func NewCollectionResponseWithTotalVersionPublicEmailWithDefaults() *CollectionResponseWithTotalVersionPublicEmail`

NewCollectionResponseWithTotalVersionPublicEmailWithDefaults instantiates a new CollectionResponseWithTotalVersionPublicEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalVersionPublicEmail) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalVersionPublicEmail) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalVersionPublicEmail) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalVersionPublicEmail) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalVersionPublicEmail) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalVersionPublicEmail) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalVersionPublicEmail) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalVersionPublicEmail) GetResults() []VersionPublicEmail`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalVersionPublicEmail) GetResultsOk() (*[]VersionPublicEmail, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalVersionPublicEmail) SetResults(v []VersionPublicEmail)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


