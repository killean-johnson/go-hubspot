# ListFolderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentFolderId** | Pointer to **string** | The folder this should be created in, if not specified will be created in the root folder 0. | [optional] 
**Name** | **string** | The name of the folder to be created. | 

## Methods

### NewListFolderCreateRequest

`func NewListFolderCreateRequest(name string, ) *ListFolderCreateRequest`

NewListFolderCreateRequest instantiates a new ListFolderCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFolderCreateRequestWithDefaults

`func NewListFolderCreateRequestWithDefaults() *ListFolderCreateRequest`

NewListFolderCreateRequestWithDefaults instantiates a new ListFolderCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentFolderId

`func (o *ListFolderCreateRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *ListFolderCreateRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *ListFolderCreateRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *ListFolderCreateRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetName

`func (o *ListFolderCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListFolderCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListFolderCreateRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


