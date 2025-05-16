# ListMoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListId** | **string** | The Id of the list to move. | 
**NewFolderId** | **string** | The Id of folder to move the list to, the root folder is Id 0. | 

## Methods

### NewListMoveRequest

`func NewListMoveRequest(listId string, newFolderId string, ) *ListMoveRequest`

NewListMoveRequest instantiates a new ListMoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMoveRequestWithDefaults

`func NewListMoveRequestWithDefaults() *ListMoveRequest`

NewListMoveRequestWithDefaults instantiates a new ListMoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListId

`func (o *ListMoveRequest) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *ListMoveRequest) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *ListMoveRequest) SetListId(v string)`

SetListId sets ListId field to given value.


### GetNewFolderId

`func (o *ListMoveRequest) GetNewFolderId() string`

GetNewFolderId returns the NewFolderId field if non-nil, zero value otherwise.

### GetNewFolderIdOk

`func (o *ListMoveRequest) GetNewFolderIdOk() (*string, bool)`

GetNewFolderIdOk returns a tuple with the NewFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFolderId

`func (o *ListMoveRequest) SetNewFolderId(v string)`

SetNewFolderId sets NewFolderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


