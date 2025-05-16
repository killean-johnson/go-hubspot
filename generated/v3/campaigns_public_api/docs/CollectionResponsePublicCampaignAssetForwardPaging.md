# CollectionResponsePublicCampaignAssetForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 
**Results** | [**[]PublicCampaignAsset**](PublicCampaignAsset.md) |  | 

## Methods

### NewCollectionResponsePublicCampaignAssetForwardPaging

`func NewCollectionResponsePublicCampaignAssetForwardPaging(results []PublicCampaignAsset, ) *CollectionResponsePublicCampaignAssetForwardPaging`

NewCollectionResponsePublicCampaignAssetForwardPaging instantiates a new CollectionResponsePublicCampaignAssetForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponsePublicCampaignAssetForwardPagingWithDefaults

`func NewCollectionResponsePublicCampaignAssetForwardPagingWithDefaults() *CollectionResponsePublicCampaignAssetForwardPaging`

NewCollectionResponsePublicCampaignAssetForwardPagingWithDefaults instantiates a new CollectionResponsePublicCampaignAssetForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *CollectionResponsePublicCampaignAssetForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponsePublicCampaignAssetForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponsePublicCampaignAssetForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponsePublicCampaignAssetForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponsePublicCampaignAssetForwardPaging) GetResults() []PublicCampaignAsset`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponsePublicCampaignAssetForwardPaging) GetResultsOk() (*[]PublicCampaignAsset, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponsePublicCampaignAssetForwardPaging) SetResults(v []PublicCampaignAsset)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


