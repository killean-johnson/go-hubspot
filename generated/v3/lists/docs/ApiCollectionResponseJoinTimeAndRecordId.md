# ApiCollectionResponseJoinTimeAndRecordId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Results** | [**[]JoinTimeAndRecordId**](JoinTimeAndRecordId.md) |  | 

## Methods

### NewApiCollectionResponseJoinTimeAndRecordId

`func NewApiCollectionResponseJoinTimeAndRecordId(results []JoinTimeAndRecordId, ) *ApiCollectionResponseJoinTimeAndRecordId`

NewApiCollectionResponseJoinTimeAndRecordId instantiates a new ApiCollectionResponseJoinTimeAndRecordId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCollectionResponseJoinTimeAndRecordIdWithDefaults

`func NewApiCollectionResponseJoinTimeAndRecordIdWithDefaults() *ApiCollectionResponseJoinTimeAndRecordId`

NewApiCollectionResponseJoinTimeAndRecordIdWithDefaults instantiates a new ApiCollectionResponseJoinTimeAndRecordId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ApiCollectionResponseJoinTimeAndRecordId) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApiCollectionResponseJoinTimeAndRecordId) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApiCollectionResponseJoinTimeAndRecordId) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ApiCollectionResponseJoinTimeAndRecordId) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPaging

`func (o *ApiCollectionResponseJoinTimeAndRecordId) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ApiCollectionResponseJoinTimeAndRecordId) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ApiCollectionResponseJoinTimeAndRecordId) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ApiCollectionResponseJoinTimeAndRecordId) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *ApiCollectionResponseJoinTimeAndRecordId) GetResults() []JoinTimeAndRecordId`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApiCollectionResponseJoinTimeAndRecordId) GetResultsOk() (*[]JoinTimeAndRecordId, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApiCollectionResponseJoinTimeAndRecordId) SetResults(v []JoinTimeAndRecordId)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


