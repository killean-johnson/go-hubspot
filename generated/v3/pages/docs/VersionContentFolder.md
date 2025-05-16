# VersionContentFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of this folder version. | 
**User** | [**VersionUser**](VersionUser.md) |  | 
**Object** | [**ContentFolder**](ContentFolder.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewVersionContentFolder

`func NewVersionContentFolder(id string, user VersionUser, object ContentFolder, updatedAt time.Time, ) *VersionContentFolder`

NewVersionContentFolder instantiates a new VersionContentFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionContentFolderWithDefaults

`func NewVersionContentFolderWithDefaults() *VersionContentFolder`

NewVersionContentFolderWithDefaults instantiates a new VersionContentFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionContentFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionContentFolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionContentFolder) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *VersionContentFolder) GetUser() VersionUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VersionContentFolder) GetUserOk() (*VersionUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VersionContentFolder) SetUser(v VersionUser)`

SetUser sets User field to given value.


### GetObject

`func (o *VersionContentFolder) GetObject() ContentFolder`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VersionContentFolder) GetObjectOk() (*ContentFolder, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VersionContentFolder) SetObject(v ContentFolder)`

SetObject sets Object field to given value.


### GetUpdatedAt

`func (o *VersionContentFolder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VersionContentFolder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VersionContentFolder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


