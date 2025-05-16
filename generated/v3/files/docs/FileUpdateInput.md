# FileUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | NONE: Do not run any duplicate validation. REJECT: Reject the upload if a duplicate is found. RETURN_EXISTING: If a duplicate file is found, do not upload a new file and return the found duplicate instead.  | [optional] 
**ParentFolderId** | Pointer to **string** | FolderId where the file should be moved to. folderId and folderPath parameters cannot be set at the same time. | [optional] 
**Name** | Pointer to **string** | New name for the file. | [optional] 
**ParentFolderPath** | Pointer to **string** | Folder path where the file should be moved to. folderId and folderPath parameters cannot be set at the same time. | [optional] 
**ClearExpires** | Pointer to **bool** |  | [optional] 
**IsUsableInContent** | Pointer to **bool** | Mark whether the file should be used in new content or not. | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFileUpdateInput

`func NewFileUpdateInput() *FileUpdateInput`

NewFileUpdateInput instantiates a new FileUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUpdateInputWithDefaults

`func NewFileUpdateInputWithDefaults() *FileUpdateInput`

NewFileUpdateInputWithDefaults instantiates a new FileUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *FileUpdateInput) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FileUpdateInput) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FileUpdateInput) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FileUpdateInput) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetParentFolderId

`func (o *FileUpdateInput) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *FileUpdateInput) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *FileUpdateInput) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *FileUpdateInput) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetName

`func (o *FileUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentFolderPath

`func (o *FileUpdateInput) GetParentFolderPath() string`

GetParentFolderPath returns the ParentFolderPath field if non-nil, zero value otherwise.

### GetParentFolderPathOk

`func (o *FileUpdateInput) GetParentFolderPathOk() (*string, bool)`

GetParentFolderPathOk returns a tuple with the ParentFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderPath

`func (o *FileUpdateInput) SetParentFolderPath(v string)`

SetParentFolderPath sets ParentFolderPath field to given value.

### HasParentFolderPath

`func (o *FileUpdateInput) HasParentFolderPath() bool`

HasParentFolderPath returns a boolean if a field has been set.

### GetClearExpires

`func (o *FileUpdateInput) GetClearExpires() bool`

GetClearExpires returns the ClearExpires field if non-nil, zero value otherwise.

### GetClearExpiresOk

`func (o *FileUpdateInput) GetClearExpiresOk() (*bool, bool)`

GetClearExpiresOk returns a tuple with the ClearExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearExpires

`func (o *FileUpdateInput) SetClearExpires(v bool)`

SetClearExpires sets ClearExpires field to given value.

### HasClearExpires

`func (o *FileUpdateInput) HasClearExpires() bool`

HasClearExpires returns a boolean if a field has been set.

### GetIsUsableInContent

`func (o *FileUpdateInput) GetIsUsableInContent() bool`

GetIsUsableInContent returns the IsUsableInContent field if non-nil, zero value otherwise.

### GetIsUsableInContentOk

`func (o *FileUpdateInput) GetIsUsableInContentOk() (*bool, bool)`

GetIsUsableInContentOk returns a tuple with the IsUsableInContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsableInContent

`func (o *FileUpdateInput) SetIsUsableInContent(v bool)`

SetIsUsableInContent sets IsUsableInContent field to given value.

### HasIsUsableInContent

`func (o *FileUpdateInput) HasIsUsableInContent() bool`

HasIsUsableInContent returns a boolean if a field has been set.

### GetExpiresAt

`func (o *FileUpdateInput) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FileUpdateInput) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FileUpdateInput) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FileUpdateInput) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


