# PublicSequenceLiteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Name** | **string** |  | 
**Id** | **string** |  | 
**UserId** | **string** |  | 
**FolderId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPublicSequenceLiteResponse

`func NewPublicSequenceLiteResponse(createdAt time.Time, name string, id string, userId string, updatedAt time.Time, ) *PublicSequenceLiteResponse`

NewPublicSequenceLiteResponse instantiates a new PublicSequenceLiteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSequenceLiteResponseWithDefaults

`func NewPublicSequenceLiteResponseWithDefaults() *PublicSequenceLiteResponse`

NewPublicSequenceLiteResponseWithDefaults instantiates a new PublicSequenceLiteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicSequenceLiteResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicSequenceLiteResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicSequenceLiteResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *PublicSequenceLiteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicSequenceLiteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicSequenceLiteResponse) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *PublicSequenceLiteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicSequenceLiteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicSequenceLiteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *PublicSequenceLiteResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicSequenceLiteResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicSequenceLiteResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFolderId

`func (o *PublicSequenceLiteResponse) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *PublicSequenceLiteResponse) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *PublicSequenceLiteResponse) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *PublicSequenceLiteResponse) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicSequenceLiteResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicSequenceLiteResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicSequenceLiteResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


