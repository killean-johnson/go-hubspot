# ContentFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedAt** | **time.Time** | The timestamp (ISO8601 format) when this content folder was deleted. | 
**ParentFolderId** | **int64** | The ID of the content folder this folder is nested under | 
**Created** | **time.Time** |  | 
**Name** | **string** | The name of the folder which will show up in the app dashboard | 
**Id** | **string** | The unique ID of the content folder. | 
**Category** | **int32** | The type of object this folder applies to. Should always be LANDING_PAGE. | 
**Updated** | **time.Time** |  | 

## Methods

### NewContentFolder

`func NewContentFolder(deletedAt time.Time, parentFolderId int64, created time.Time, name string, id string, category int32, updated time.Time, ) *ContentFolder`

NewContentFolder instantiates a new ContentFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentFolderWithDefaults

`func NewContentFolderWithDefaults() *ContentFolder`

NewContentFolderWithDefaults instantiates a new ContentFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedAt

`func (o *ContentFolder) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ContentFolder) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ContentFolder) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetParentFolderId

`func (o *ContentFolder) GetParentFolderId() int64`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *ContentFolder) GetParentFolderIdOk() (*int64, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *ContentFolder) SetParentFolderId(v int64)`

SetParentFolderId sets ParentFolderId field to given value.


### GetCreated

`func (o *ContentFolder) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ContentFolder) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ContentFolder) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetName

`func (o *ContentFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentFolder) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *ContentFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentFolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentFolder) SetId(v string)`

SetId sets Id field to given value.


### GetCategory

`func (o *ContentFolder) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ContentFolder) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ContentFolder) SetCategory(v int32)`

SetCategory sets Category field to given value.


### GetUpdated

`func (o *ContentFolder) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ContentFolder) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ContentFolder) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


