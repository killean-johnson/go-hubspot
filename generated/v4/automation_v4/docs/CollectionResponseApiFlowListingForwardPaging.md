# CollectionResponseApiFlowListingForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 
**Results** | [**[]ApiFlowListing**](ApiFlowListing.md) |  | 

## Methods

### NewCollectionResponseApiFlowListingForwardPaging

`func NewCollectionResponseApiFlowListingForwardPaging(results []ApiFlowListing, ) *CollectionResponseApiFlowListingForwardPaging`

NewCollectionResponseApiFlowListingForwardPaging instantiates a new CollectionResponseApiFlowListingForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseApiFlowListingForwardPagingWithDefaults

`func NewCollectionResponseApiFlowListingForwardPagingWithDefaults() *CollectionResponseApiFlowListingForwardPaging`

NewCollectionResponseApiFlowListingForwardPagingWithDefaults instantiates a new CollectionResponseApiFlowListingForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *CollectionResponseApiFlowListingForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseApiFlowListingForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseApiFlowListingForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseApiFlowListingForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseApiFlowListingForwardPaging) GetResults() []ApiFlowListing`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseApiFlowListingForwardPaging) GetResultsOk() (*[]ApiFlowListing, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseApiFlowListingForwardPaging) SetResults(v []ApiFlowListing)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


