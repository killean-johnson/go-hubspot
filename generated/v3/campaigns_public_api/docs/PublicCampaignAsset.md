# PublicCampaignAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Metrics** | **map[string]float32** |  | 

## Methods

### NewPublicCampaignAsset

`func NewPublicCampaignAsset(id string, metrics map[string]float32, ) *PublicCampaignAsset`

NewPublicCampaignAsset instantiates a new PublicCampaignAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCampaignAssetWithDefaults

`func NewPublicCampaignAssetWithDefaults() *PublicCampaignAsset`

NewPublicCampaignAssetWithDefaults instantiates a new PublicCampaignAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PublicCampaignAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicCampaignAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicCampaignAsset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicCampaignAsset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *PublicCampaignAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicCampaignAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicCampaignAsset) SetId(v string)`

SetId sets Id field to given value.


### GetMetrics

`func (o *PublicCampaignAsset) GetMetrics() map[string]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *PublicCampaignAsset) GetMetricsOk() (*map[string]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *PublicCampaignAsset) SetMetrics(v map[string]float32)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


