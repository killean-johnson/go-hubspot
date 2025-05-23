# PublicImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportTemplate** | Pointer to [**ImportTemplate**](ImportTemplate.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Metadata** | [**PublicImportMetadata**](PublicImportMetadata.md) |  | 
**ImportRequestJson** | Pointer to **map[string]interface{}** |  | [optional] 
**ImportSource** | Pointer to **string** |  | [optional] 
**ImportName** | Pointer to **string** |  | [optional] 
**State** | **string** | The status of the import. | 
**Id** | **string** |  | 
**OptOutImport** | **bool** | Whether or not the import is a list of people disqualified from receiving emails. | 
**MappedObjectTypeIds** | **[]string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPublicImportResponse

`func NewPublicImportResponse(createdAt time.Time, metadata PublicImportMetadata, state string, id string, optOutImport bool, mappedObjectTypeIds []string, updatedAt time.Time, ) *PublicImportResponse`

NewPublicImportResponse instantiates a new PublicImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicImportResponseWithDefaults

`func NewPublicImportResponseWithDefaults() *PublicImportResponse`

NewPublicImportResponseWithDefaults instantiates a new PublicImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportTemplate

`func (o *PublicImportResponse) GetImportTemplate() ImportTemplate`

GetImportTemplate returns the ImportTemplate field if non-nil, zero value otherwise.

### GetImportTemplateOk

`func (o *PublicImportResponse) GetImportTemplateOk() (*ImportTemplate, bool)`

GetImportTemplateOk returns a tuple with the ImportTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTemplate

`func (o *PublicImportResponse) SetImportTemplate(v ImportTemplate)`

SetImportTemplate sets ImportTemplate field to given value.

### HasImportTemplate

`func (o *PublicImportResponse) HasImportTemplate() bool`

HasImportTemplate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicImportResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicImportResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicImportResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMetadata

`func (o *PublicImportResponse) GetMetadata() PublicImportMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublicImportResponse) GetMetadataOk() (*PublicImportMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublicImportResponse) SetMetadata(v PublicImportMetadata)`

SetMetadata sets Metadata field to given value.


### GetImportRequestJson

`func (o *PublicImportResponse) GetImportRequestJson() map[string]interface{}`

GetImportRequestJson returns the ImportRequestJson field if non-nil, zero value otherwise.

### GetImportRequestJsonOk

`func (o *PublicImportResponse) GetImportRequestJsonOk() (*map[string]interface{}, bool)`

GetImportRequestJsonOk returns a tuple with the ImportRequestJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportRequestJson

`func (o *PublicImportResponse) SetImportRequestJson(v map[string]interface{})`

SetImportRequestJson sets ImportRequestJson field to given value.

### HasImportRequestJson

`func (o *PublicImportResponse) HasImportRequestJson() bool`

HasImportRequestJson returns a boolean if a field has been set.

### GetImportSource

`func (o *PublicImportResponse) GetImportSource() string`

GetImportSource returns the ImportSource field if non-nil, zero value otherwise.

### GetImportSourceOk

`func (o *PublicImportResponse) GetImportSourceOk() (*string, bool)`

GetImportSourceOk returns a tuple with the ImportSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSource

`func (o *PublicImportResponse) SetImportSource(v string)`

SetImportSource sets ImportSource field to given value.

### HasImportSource

`func (o *PublicImportResponse) HasImportSource() bool`

HasImportSource returns a boolean if a field has been set.

### GetImportName

`func (o *PublicImportResponse) GetImportName() string`

GetImportName returns the ImportName field if non-nil, zero value otherwise.

### GetImportNameOk

`func (o *PublicImportResponse) GetImportNameOk() (*string, bool)`

GetImportNameOk returns a tuple with the ImportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportName

`func (o *PublicImportResponse) SetImportName(v string)`

SetImportName sets ImportName field to given value.

### HasImportName

`func (o *PublicImportResponse) HasImportName() bool`

HasImportName returns a boolean if a field has been set.

### GetState

`func (o *PublicImportResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PublicImportResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PublicImportResponse) SetState(v string)`

SetState sets State field to given value.


### GetId

`func (o *PublicImportResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicImportResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicImportResponse) SetId(v string)`

SetId sets Id field to given value.


### GetOptOutImport

`func (o *PublicImportResponse) GetOptOutImport() bool`

GetOptOutImport returns the OptOutImport field if non-nil, zero value otherwise.

### GetOptOutImportOk

`func (o *PublicImportResponse) GetOptOutImportOk() (*bool, bool)`

GetOptOutImportOk returns a tuple with the OptOutImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOutImport

`func (o *PublicImportResponse) SetOptOutImport(v bool)`

SetOptOutImport sets OptOutImport field to given value.


### GetMappedObjectTypeIds

`func (o *PublicImportResponse) GetMappedObjectTypeIds() []string`

GetMappedObjectTypeIds returns the MappedObjectTypeIds field if non-nil, zero value otherwise.

### GetMappedObjectTypeIdsOk

`func (o *PublicImportResponse) GetMappedObjectTypeIdsOk() (*[]string, bool)`

GetMappedObjectTypeIdsOk returns a tuple with the MappedObjectTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedObjectTypeIds

`func (o *PublicImportResponse) SetMappedObjectTypeIds(v []string)`

SetMappedObjectTypeIds sets MappedObjectTypeIds field to given value.


### GetUpdatedAt

`func (o *PublicImportResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicImportResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicImportResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


