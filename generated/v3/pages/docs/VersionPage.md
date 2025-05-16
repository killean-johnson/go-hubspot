# VersionPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of this page version. | 
**User** | [**VersionUser**](VersionUser.md) |  | 
**Object** | [**Page**](Page.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewVersionPage

`func NewVersionPage(id string, user VersionUser, object Page, updatedAt time.Time, ) *VersionPage`

NewVersionPage instantiates a new VersionPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionPageWithDefaults

`func NewVersionPageWithDefaults() *VersionPage`

NewVersionPageWithDefaults instantiates a new VersionPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionPage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionPage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionPage) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *VersionPage) GetUser() VersionUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VersionPage) GetUserOk() (*VersionUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VersionPage) SetUser(v VersionUser)`

SetUser sets User field to given value.


### GetObject

`func (o *VersionPage) GetObject() Page`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VersionPage) GetObjectOk() (*Page, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VersionPage) SetObject(v Page)`

SetObject sets Object field to given value.


### GetUpdatedAt

`func (o *VersionPage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VersionPage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VersionPage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


