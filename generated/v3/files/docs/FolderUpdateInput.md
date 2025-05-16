# FolderUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentFolderId** | Pointer to **int64** | New parent folderId. If changed, the folder and all it&#39;s children will be moved into the specified folder. parentFolderId and parentFolderPath cannot be specified at the same time. | [optional] 
**Name** | Pointer to **string** | New name. If specified the folder&#39;s name and fullPath will change. All children of the folder will be updated accordingly. | [optional] 

## Methods

### NewFolderUpdateInput

`func NewFolderUpdateInput() *FolderUpdateInput`

NewFolderUpdateInput instantiates a new FolderUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderUpdateInputWithDefaults

`func NewFolderUpdateInputWithDefaults() *FolderUpdateInput`

NewFolderUpdateInputWithDefaults instantiates a new FolderUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentFolderId

`func (o *FolderUpdateInput) GetParentFolderId() int64`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *FolderUpdateInput) GetParentFolderIdOk() (*int64, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *FolderUpdateInput) SetParentFolderId(v int64)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *FolderUpdateInput) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetName

`func (o *FolderUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FolderUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FolderUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FolderUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


