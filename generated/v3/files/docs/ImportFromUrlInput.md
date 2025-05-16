# ImportFromUrlInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderPath** | Pointer to **string** | One of folderPath or folderId is required. Destination folder path for the uploaded file. If the folder path does not exist, there will be an attempt to create the folder path. | [optional] 
**Access** | **string** | PUBLIC_INDEXABLE: File is publicly accessible by anyone who has the URL. Search engines can index the file. PUBLIC_NOT_INDEXABLE: File is publicly accessible by anyone who has the URL. Search engines *can&#39;t* index the file. PRIVATE: File is NOT publicly accessible. Requires a signed URL to see content. Search engines *can&#39;t* index the file.  | 
**DuplicateValidationScope** | Pointer to **string** | ENTIRE_PORTAL: Look for a duplicate file in the entire account. EXACT_FOLDER: Look for a duplicate file in the provided folder.  | [optional] 
**Name** | Pointer to **string** | Name to give the resulting file in the file manager. | [optional] 
**DuplicateValidationStrategy** | Pointer to **string** | NONE: Do not run any duplicate validation. REJECT: Reject the upload if a duplicate is found. RETURN_EXISTING: If a duplicate file is found, do not upload a new file and return the found duplicate instead.  | [optional] 
**Ttl** | Pointer to **string** | Time to live. If specified the file will be deleted after the given time frame. If left unset, the file will exist indefinitely | [optional] 
**Overwrite** | Pointer to **bool** | If true, will overwrite existing file if one with the same name and extension exists in the given folder. The overwritten file will be deleted and the uploaded file will take its place with a new ID. If unset or set as false, the new file&#39;s name will be updated to prevent colliding with existing file if one exists with the same path, name, and extension | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Url** | **string** | URL to download the new file from. | 
**FolderId** | Pointer to **string** | One of folderId or folderPath is required. Destination folderId for the uploaded file. | [optional] 

## Methods

### NewImportFromUrlInput

`func NewImportFromUrlInput(access string, url string, ) *ImportFromUrlInput`

NewImportFromUrlInput instantiates a new ImportFromUrlInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportFromUrlInputWithDefaults

`func NewImportFromUrlInputWithDefaults() *ImportFromUrlInput`

NewImportFromUrlInputWithDefaults instantiates a new ImportFromUrlInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderPath

`func (o *ImportFromUrlInput) GetFolderPath() string`

GetFolderPath returns the FolderPath field if non-nil, zero value otherwise.

### GetFolderPathOk

`func (o *ImportFromUrlInput) GetFolderPathOk() (*string, bool)`

GetFolderPathOk returns a tuple with the FolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderPath

`func (o *ImportFromUrlInput) SetFolderPath(v string)`

SetFolderPath sets FolderPath field to given value.

### HasFolderPath

`func (o *ImportFromUrlInput) HasFolderPath() bool`

HasFolderPath returns a boolean if a field has been set.

### GetAccess

`func (o *ImportFromUrlInput) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ImportFromUrlInput) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ImportFromUrlInput) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetDuplicateValidationScope

`func (o *ImportFromUrlInput) GetDuplicateValidationScope() string`

GetDuplicateValidationScope returns the DuplicateValidationScope field if non-nil, zero value otherwise.

### GetDuplicateValidationScopeOk

`func (o *ImportFromUrlInput) GetDuplicateValidationScopeOk() (*string, bool)`

GetDuplicateValidationScopeOk returns a tuple with the DuplicateValidationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateValidationScope

`func (o *ImportFromUrlInput) SetDuplicateValidationScope(v string)`

SetDuplicateValidationScope sets DuplicateValidationScope field to given value.

### HasDuplicateValidationScope

`func (o *ImportFromUrlInput) HasDuplicateValidationScope() bool`

HasDuplicateValidationScope returns a boolean if a field has been set.

### GetName

`func (o *ImportFromUrlInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportFromUrlInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportFromUrlInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportFromUrlInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDuplicateValidationStrategy

`func (o *ImportFromUrlInput) GetDuplicateValidationStrategy() string`

GetDuplicateValidationStrategy returns the DuplicateValidationStrategy field if non-nil, zero value otherwise.

### GetDuplicateValidationStrategyOk

`func (o *ImportFromUrlInput) GetDuplicateValidationStrategyOk() (*string, bool)`

GetDuplicateValidationStrategyOk returns a tuple with the DuplicateValidationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateValidationStrategy

`func (o *ImportFromUrlInput) SetDuplicateValidationStrategy(v string)`

SetDuplicateValidationStrategy sets DuplicateValidationStrategy field to given value.

### HasDuplicateValidationStrategy

`func (o *ImportFromUrlInput) HasDuplicateValidationStrategy() bool`

HasDuplicateValidationStrategy returns a boolean if a field has been set.

### GetTtl

`func (o *ImportFromUrlInput) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ImportFromUrlInput) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ImportFromUrlInput) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ImportFromUrlInput) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetOverwrite

`func (o *ImportFromUrlInput) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *ImportFromUrlInput) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *ImportFromUrlInput) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *ImportFromUrlInput) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ImportFromUrlInput) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ImportFromUrlInput) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ImportFromUrlInput) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ImportFromUrlInput) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUrl

`func (o *ImportFromUrlInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImportFromUrlInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImportFromUrlInput) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFolderId

`func (o *ImportFromUrlInput) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ImportFromUrlInput) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ImportFromUrlInput) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ImportFromUrlInput) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


