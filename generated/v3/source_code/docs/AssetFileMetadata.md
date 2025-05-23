# AssetFileMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** | Timestamp of when the object was first created. | 
**ArchivedAt** | Pointer to **int64** | Timestamp of when the object was archived (deleted). | [optional] 
**Folder** | **bool** | Determines whether or not this path points to a folder. | 
**Children** | Pointer to **[]string** | If the object is a folder, contains the filenames of the files within the folder. | [optional] 
**Name** | **string** | The name of the file. | 
**Id** | **string** | The path of the file in the CMS Developer File System. | 
**Hash** | Pointer to **string** |  | [optional] 
**UpdatedAt** | **int64** | Timestamp of when the object was last updated. | 

## Methods

### NewAssetFileMetadata

`func NewAssetFileMetadata(createdAt int64, folder bool, name string, id string, updatedAt int64, ) *AssetFileMetadata`

NewAssetFileMetadata instantiates a new AssetFileMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetFileMetadataWithDefaults

`func NewAssetFileMetadataWithDefaults() *AssetFileMetadata`

NewAssetFileMetadataWithDefaults instantiates a new AssetFileMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AssetFileMetadata) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssetFileMetadata) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssetFileMetadata) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetArchivedAt

`func (o *AssetFileMetadata) GetArchivedAt() int64`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *AssetFileMetadata) GetArchivedAtOk() (*int64, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *AssetFileMetadata) SetArchivedAt(v int64)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *AssetFileMetadata) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetFolder

`func (o *AssetFileMetadata) GetFolder() bool`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *AssetFileMetadata) GetFolderOk() (*bool, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *AssetFileMetadata) SetFolder(v bool)`

SetFolder sets Folder field to given value.


### GetChildren

`func (o *AssetFileMetadata) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AssetFileMetadata) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AssetFileMetadata) SetChildren(v []string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AssetFileMetadata) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetName

`func (o *AssetFileMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetFileMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetFileMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *AssetFileMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetFileMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetFileMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetHash

`func (o *AssetFileMetadata) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AssetFileMetadata) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AssetFileMetadata) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *AssetFileMetadata) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AssetFileMetadata) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssetFileMetadata) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssetFileMetadata) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


