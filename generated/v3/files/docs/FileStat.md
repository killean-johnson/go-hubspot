# FileStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**File**](File.md) |  | [optional] 
**Folder** | Pointer to [**Folder**](Folder.md) |  | [optional] 

## Methods

### NewFileStat

`func NewFileStat() *FileStat`

NewFileStat instantiates a new FileStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileStatWithDefaults

`func NewFileStatWithDefaults() *FileStat`

NewFileStatWithDefaults instantiates a new FileStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *FileStat) GetFile() File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileStat) GetFileOk() (*File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileStat) SetFile(v File)`

SetFile sets File field to given value.

### HasFile

`func (o *FileStat) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFolder

`func (o *FileStat) GetFolder() Folder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *FileStat) GetFolderOk() (*Folder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *FileStat) SetFolder(v Folder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *FileStat) HasFolder() bool`

HasFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


