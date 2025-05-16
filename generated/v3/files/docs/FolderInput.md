# FolderInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentFolderId** | Pointer to **string** | FolderId of the parent of the created folder. If not specified, the folder will be created at the root level. parentFolderId and parentFolderPath cannot be set at the same time. | [optional] 
**ParentPath** | Pointer to **string** | Path of the parent of the created folder. If not specified the folder will be created at the root level. parentFolderPath and parentFolderId cannot be set at the same time. | [optional] 
**Name** | **string** | Desired name for the folder. | 

## Methods

### NewFolderInput

`func NewFolderInput(name string, ) *FolderInput`

NewFolderInput instantiates a new FolderInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderInputWithDefaults

`func NewFolderInputWithDefaults() *FolderInput`

NewFolderInputWithDefaults instantiates a new FolderInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentFolderId

`func (o *FolderInput) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *FolderInput) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *FolderInput) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *FolderInput) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetParentPath

`func (o *FolderInput) GetParentPath() string`

GetParentPath returns the ParentPath field if non-nil, zero value otherwise.

### GetParentPathOk

`func (o *FolderInput) GetParentPathOk() (*string, bool)`

GetParentPathOk returns a tuple with the ParentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPath

`func (o *FolderInput) SetParentPath(v string)`

SetParentPath sets ParentPath field to given value.

### HasParentPath

`func (o *FolderInput) HasParentPath() bool`

HasParentPath returns a boolean if a field has been set.

### GetName

`func (o *FolderInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FolderInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FolderInput) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


