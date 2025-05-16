# CollectionResponseWithTotalParticipationBreakdownForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 
**Results** | [**[]ParticipationBreakdown**](ParticipationBreakdown.md) |  | 

## Methods

### NewCollectionResponseWithTotalParticipationBreakdownForwardPaging

`func NewCollectionResponseWithTotalParticipationBreakdownForwardPaging(total int32, results []ParticipationBreakdown, ) *CollectionResponseWithTotalParticipationBreakdownForwardPaging`

NewCollectionResponseWithTotalParticipationBreakdownForwardPaging instantiates a new CollectionResponseWithTotalParticipationBreakdownForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalParticipationBreakdownForwardPagingWithDefaults

`func NewCollectionResponseWithTotalParticipationBreakdownForwardPagingWithDefaults() *CollectionResponseWithTotalParticipationBreakdownForwardPaging`

NewCollectionResponseWithTotalParticipationBreakdownForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalParticipationBreakdownForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) GetResults() []ParticipationBreakdown`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) GetResultsOk() (*[]ParticipationBreakdown, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalParticipationBreakdownForwardPaging) SetResults(v []ParticipationBreakdown)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


