# PublicListFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The time the folder was created at. | [optional] 
**ParentFolderId** | **int32** | The Id of the folder this folder is in, the root folder is represented as 0. | 
**ChildNodes** | [**[]PublicListFolder**](PublicListFolder.md) |  | 
**Name** | Pointer to **string** | The name of the folder. | [optional] 
**Id** | **int32** | The Id of the folder. | 
**ChildLists** | **[]int32** | An array of list Id&#39;s contained in this folder. | 
**UpdatedContentsAt** | Pointer to **time.Time** | The time that the contents of the folder was last updated at. | [optional] 
**UserId** | Pointer to **int32** | The user Id of the owner of the folder. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the folder was last updated at. | [optional] 

## Methods

### NewPublicListFolder

`func NewPublicListFolder(parentFolderId int32, childNodes []PublicListFolder, id int32, childLists []int32, ) *PublicListFolder`

NewPublicListFolder instantiates a new PublicListFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicListFolderWithDefaults

`func NewPublicListFolderWithDefaults() *PublicListFolder`

NewPublicListFolderWithDefaults instantiates a new PublicListFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicListFolder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicListFolder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicListFolder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicListFolder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetParentFolderId

`func (o *PublicListFolder) GetParentFolderId() int32`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *PublicListFolder) GetParentFolderIdOk() (*int32, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *PublicListFolder) SetParentFolderId(v int32)`

SetParentFolderId sets ParentFolderId field to given value.


### GetChildNodes

`func (o *PublicListFolder) GetChildNodes() []PublicListFolder`

GetChildNodes returns the ChildNodes field if non-nil, zero value otherwise.

### GetChildNodesOk

`func (o *PublicListFolder) GetChildNodesOk() (*[]PublicListFolder, bool)`

GetChildNodesOk returns a tuple with the ChildNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildNodes

`func (o *PublicListFolder) SetChildNodes(v []PublicListFolder)`

SetChildNodes sets ChildNodes field to given value.


### GetName

`func (o *PublicListFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicListFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicListFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicListFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *PublicListFolder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicListFolder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicListFolder) SetId(v int32)`

SetId sets Id field to given value.


### GetChildLists

`func (o *PublicListFolder) GetChildLists() []int32`

GetChildLists returns the ChildLists field if non-nil, zero value otherwise.

### GetChildListsOk

`func (o *PublicListFolder) GetChildListsOk() (*[]int32, bool)`

GetChildListsOk returns a tuple with the ChildLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildLists

`func (o *PublicListFolder) SetChildLists(v []int32)`

SetChildLists sets ChildLists field to given value.


### GetUpdatedContentsAt

`func (o *PublicListFolder) GetUpdatedContentsAt() time.Time`

GetUpdatedContentsAt returns the UpdatedContentsAt field if non-nil, zero value otherwise.

### GetUpdatedContentsAtOk

`func (o *PublicListFolder) GetUpdatedContentsAtOk() (*time.Time, bool)`

GetUpdatedContentsAtOk returns a tuple with the UpdatedContentsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedContentsAt

`func (o *PublicListFolder) SetUpdatedContentsAt(v time.Time)`

SetUpdatedContentsAt sets UpdatedContentsAt field to given value.

### HasUpdatedContentsAt

`func (o *PublicListFolder) HasUpdatedContentsAt() bool`

HasUpdatedContentsAt returns a boolean if a field has been set.

### GetUserId

`func (o *PublicListFolder) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicListFolder) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicListFolder) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PublicListFolder) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicListFolder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicListFolder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicListFolder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicListFolder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


