# FileAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUsageType** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "FILE"]
**FileId** | **string** |  | 

## Methods

### NewFileAttachment

`func NewFileAttachment(type_ string, fileId string, ) *FileAttachment`

NewFileAttachment instantiates a new FileAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileAttachmentWithDefaults

`func NewFileAttachmentWithDefaults() *FileAttachment`

NewFileAttachmentWithDefaults instantiates a new FileAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUsageType

`func (o *FileAttachment) GetFileUsageType() string`

GetFileUsageType returns the FileUsageType field if non-nil, zero value otherwise.

### GetFileUsageTypeOk

`func (o *FileAttachment) GetFileUsageTypeOk() (*string, bool)`

GetFileUsageTypeOk returns a tuple with the FileUsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUsageType

`func (o *FileAttachment) SetFileUsageType(v string)`

SetFileUsageType sets FileUsageType field to given value.

### HasFileUsageType

`func (o *FileAttachment) HasFileUsageType() bool`

HasFileUsageType returns a boolean if a field has been set.

### GetType

`func (o *FileAttachment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileAttachment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileAttachment) SetType(v string)`

SetType sets Type field to given value.


### GetFileId

`func (o *FileAttachment) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileAttachment) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileAttachment) SetFileId(v string)`

SetFileId sets FileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


