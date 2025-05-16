# CollectionResponsePublicCampaignAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Results** | [**[]PublicCampaignAsset**](PublicCampaignAsset.md) |  | 

## Methods

### NewCollectionResponsePublicCampaignAsset

`func NewCollectionResponsePublicCampaignAsset(results []PublicCampaignAsset, ) *CollectionResponsePublicCampaignAsset`

NewCollectionResponsePublicCampaignAsset instantiates a new CollectionResponsePublicCampaignAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponsePublicCampaignAssetWithDefaults

`func NewCollectionResponsePublicCampaignAssetWithDefaults() *CollectionResponsePublicCampaignAsset`

NewCollectionResponsePublicCampaignAssetWithDefaults instantiates a new CollectionResponsePublicCampaignAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *CollectionResponsePublicCampaignAsset) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponsePublicCampaignAsset) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponsePublicCampaignAsset) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponsePublicCampaignAsset) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponsePublicCampaignAsset) GetResults() []PublicCampaignAsset`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponsePublicCampaignAsset) GetResultsOk() (*[]PublicCampaignAsset, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponsePublicCampaignAsset) SetResults(v []PublicCampaignAsset)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


