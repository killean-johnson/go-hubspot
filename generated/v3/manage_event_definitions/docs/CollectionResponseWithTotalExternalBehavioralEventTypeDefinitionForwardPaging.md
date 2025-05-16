# CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 
**Results** | [**[]ExternalBehavioralEventTypeDefinition**](ExternalBehavioralEventTypeDefinition.md) |  | 

## Methods

### NewCollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging

`func NewCollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging(total int32, results []ExternalBehavioralEventTypeDefinition, ) *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging`

NewCollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging instantiates a new CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPagingWithDefaults

`func NewCollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPagingWithDefaults() *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging`

NewCollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) GetResults() []ExternalBehavioralEventTypeDefinition`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) GetResultsOk() (*[]ExternalBehavioralEventTypeDefinition, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalExternalBehavioralEventTypeDefinitionForwardPaging) SetResults(v []ExternalBehavioralEventTypeDefinition)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


