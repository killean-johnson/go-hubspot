# VersionBlog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the version. | 
**User** | [**VersionUser**](VersionUser.md) |  | 
**Object** | [**Blog**](Blog.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewVersionBlog

`func NewVersionBlog(id string, user VersionUser, object Blog, updatedAt time.Time, ) *VersionBlog`

NewVersionBlog instantiates a new VersionBlog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionBlogWithDefaults

`func NewVersionBlogWithDefaults() *VersionBlog`

NewVersionBlogWithDefaults instantiates a new VersionBlog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionBlog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionBlog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionBlog) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *VersionBlog) GetUser() VersionUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VersionBlog) GetUserOk() (*VersionUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VersionBlog) SetUser(v VersionUser)`

SetUser sets User field to given value.


### GetObject

`func (o *VersionBlog) GetObject() Blog`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VersionBlog) GetObjectOk() (*Blog, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VersionBlog) SetObject(v Blog)`

SetObject sets Object field to given value.


### GetUpdatedAt

`func (o *VersionBlog) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VersionBlog) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VersionBlog) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


