# CollectionResponseWithTotalPublicCampaignForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 
**Results** | [**[]PublicCampaign**](PublicCampaign.md) |  | 

## Methods

### NewCollectionResponseWithTotalPublicCampaignForwardPaging

`func NewCollectionResponseWithTotalPublicCampaignForwardPaging(total int32, results []PublicCampaign, ) *CollectionResponseWithTotalPublicCampaignForwardPaging`

NewCollectionResponseWithTotalPublicCampaignForwardPaging instantiates a new CollectionResponseWithTotalPublicCampaignForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalPublicCampaignForwardPagingWithDefaults

`func NewCollectionResponseWithTotalPublicCampaignForwardPagingWithDefaults() *CollectionResponseWithTotalPublicCampaignForwardPaging`

NewCollectionResponseWithTotalPublicCampaignForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalPublicCampaignForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) GetResults() []PublicCampaign`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) GetResultsOk() (*[]PublicCampaign, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalPublicCampaignForwardPaging) SetResults(v []PublicCampaign)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


