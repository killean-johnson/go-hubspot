# PublicFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUsageType** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "FILE"]
**Url** | Pointer to **string** |  | [optional] 
**FileId** | **string** |  | 

## Methods

### NewPublicFile

`func NewPublicFile(fileUsageType string, type_ string, fileId string, ) *PublicFile`

NewPublicFile instantiates a new PublicFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFileWithDefaults

`func NewPublicFileWithDefaults() *PublicFile`

NewPublicFileWithDefaults instantiates a new PublicFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUsageType

`func (o *PublicFile) GetFileUsageType() string`

GetFileUsageType returns the FileUsageType field if non-nil, zero value otherwise.

### GetFileUsageTypeOk

`func (o *PublicFile) GetFileUsageTypeOk() (*string, bool)`

GetFileUsageTypeOk returns a tuple with the FileUsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUsageType

`func (o *PublicFile) SetFileUsageType(v string)`

SetFileUsageType sets FileUsageType field to given value.


### GetName

`func (o *PublicFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PublicFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicFile) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *PublicFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PublicFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PublicFile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PublicFile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFileId

`func (o *PublicFile) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *PublicFile) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *PublicFile) SetFileId(v string)`

SetFileId sets FileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


