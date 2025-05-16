# PublicCampaignWithAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Assets** | [**map[string]CollectionResponsePublicCampaignAsset**](CollectionResponsePublicCampaignAsset.md) |  | 
**Id** | **string** |  | 
**Properties** | **map[string]string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPublicCampaignWithAssets

`func NewPublicCampaignWithAssets(createdAt time.Time, assets map[string]CollectionResponsePublicCampaignAsset, id string, properties map[string]string, updatedAt time.Time, ) *PublicCampaignWithAssets`

NewPublicCampaignWithAssets instantiates a new PublicCampaignWithAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCampaignWithAssetsWithDefaults

`func NewPublicCampaignWithAssetsWithDefaults() *PublicCampaignWithAssets`

NewPublicCampaignWithAssetsWithDefaults instantiates a new PublicCampaignWithAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicCampaignWithAssets) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicCampaignWithAssets) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicCampaignWithAssets) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAssets

`func (o *PublicCampaignWithAssets) GetAssets() map[string]CollectionResponsePublicCampaignAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *PublicCampaignWithAssets) GetAssetsOk() (*map[string]CollectionResponsePublicCampaignAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *PublicCampaignWithAssets) SetAssets(v map[string]CollectionResponsePublicCampaignAsset)`

SetAssets sets Assets field to given value.


### GetId

`func (o *PublicCampaignWithAssets) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicCampaignWithAssets) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicCampaignWithAssets) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *PublicCampaignWithAssets) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PublicCampaignWithAssets) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PublicCampaignWithAssets) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.


### GetUpdatedAt

`func (o *PublicCampaignWithAssets) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicCampaignWithAssets) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicCampaignWithAssets) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


