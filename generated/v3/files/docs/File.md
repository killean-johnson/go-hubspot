# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extension** | Pointer to **string** | Extension of the file. ex: .jpg, .png, .gif, .pdf, etc. | [optional] 
**Access** | **string** | File access. Can be PUBLIC_INDEXABLE, PUBLIC_NOT_INDEXABLE, PRIVATE. | 
**ParentFolderId** | Pointer to **string** | ID of the folder the file is in. | [optional] 
**SourceGroup** | Pointer to **string** |  | [optional] 
**FileMd5** | Pointer to **string** |  | [optional] 
**Encoding** | Pointer to **string** | Encoding of the file. | [optional] 
**Type** | Pointer to **string** | Type of the file. Can be IMG, DOCUMENT, AUDIO, MOVIE, or OTHER. | [optional] 
**IsUsableInContent** | Pointer to **bool** | Previously \&quot;archied\&quot;. Indicates if the file should be used when creating new content like web pages. | [optional] 
**Url** | Pointer to **string** | URL of the given file. This URL can change depending on the domain settings of the account. Will use the select file hosting domain. | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**CreatedAt** | **time.Time** | Creation time of the file object. | 
**ArchivedAt** | Pointer to **time.Time** | Deletion time of the file object. | [optional] 
**Path** | Pointer to **string** | Path of the file in the file manager. | [optional] 
**Archived** | **bool** | If the file is deleted. | 
**Size** | Pointer to **int64** | Size of the file in bytes. | [optional] 
**Name** | Pointer to **string** | Name of the file. | [optional] 
**Width** | Pointer to **int32** | For image and video files, the width of the content. | [optional] 
**Id** | **string** | File ID. | 
**DefaultHostingUrl** | Pointer to **string** | Default hosting URL of the file. This will use one of HubSpot&#39;s provided URLs to serve the file. | [optional] 
**UpdatedAt** | **time.Time** | Timestamp of the latest update to the file. | 
**Height** | Pointer to **int32** | For image and video files, the height of the content. | [optional] 

## Methods

### NewFile

`func NewFile(access string, createdAt time.Time, archived bool, id string, updatedAt time.Time, ) *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtension

`func (o *File) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *File) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *File) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *File) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetAccess

`func (o *File) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *File) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *File) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetParentFolderId

`func (o *File) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *File) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *File) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *File) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetSourceGroup

`func (o *File) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *File) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *File) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *File) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetFileMd5

`func (o *File) GetFileMd5() string`

GetFileMd5 returns the FileMd5 field if non-nil, zero value otherwise.

### GetFileMd5Ok

`func (o *File) GetFileMd5Ok() (*string, bool)`

GetFileMd5Ok returns a tuple with the FileMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMd5

`func (o *File) SetFileMd5(v string)`

SetFileMd5 sets FileMd5 field to given value.

### HasFileMd5

`func (o *File) HasFileMd5() bool`

HasFileMd5 returns a boolean if a field has been set.

### GetEncoding

`func (o *File) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *File) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *File) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *File) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetType

`func (o *File) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *File) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *File) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *File) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsUsableInContent

`func (o *File) GetIsUsableInContent() bool`

GetIsUsableInContent returns the IsUsableInContent field if non-nil, zero value otherwise.

### GetIsUsableInContentOk

`func (o *File) GetIsUsableInContentOk() (*bool, bool)`

GetIsUsableInContentOk returns a tuple with the IsUsableInContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsableInContent

`func (o *File) SetIsUsableInContent(v bool)`

SetIsUsableInContent sets IsUsableInContent field to given value.

### HasIsUsableInContent

`func (o *File) HasIsUsableInContent() bool`

HasIsUsableInContent returns a boolean if a field has been set.

### GetUrl

`func (o *File) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *File) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *File) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *File) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *File) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *File) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *File) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *File) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *File) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *File) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *File) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetArchivedAt

`func (o *File) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *File) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *File) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *File) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetPath

`func (o *File) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *File) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *File) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *File) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetArchived

`func (o *File) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *File) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *File) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetSize

`func (o *File) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *File) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *File) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *File) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetName

`func (o *File) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *File) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *File) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *File) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWidth

`func (o *File) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *File) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *File) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *File) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetId

`func (o *File) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *File) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *File) SetId(v string)`

SetId sets Id field to given value.


### GetDefaultHostingUrl

`func (o *File) GetDefaultHostingUrl() string`

GetDefaultHostingUrl returns the DefaultHostingUrl field if non-nil, zero value otherwise.

### GetDefaultHostingUrlOk

`func (o *File) GetDefaultHostingUrlOk() (*string, bool)`

GetDefaultHostingUrlOk returns a tuple with the DefaultHostingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHostingUrl

`func (o *File) SetDefaultHostingUrl(v string)`

SetDefaultHostingUrl sets DefaultHostingUrl field to given value.

### HasDefaultHostingUrl

`func (o *File) HasDefaultHostingUrl() bool`

HasDefaultHostingUrl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *File) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *File) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *File) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetHeight

`func (o *File) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *File) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *File) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *File) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


