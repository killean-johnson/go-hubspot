# VersionPublicEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of this marketing email version. | 
**User** | [**VersionUser**](VersionUser.md) |  | 
**Object** | [**PublicEmail**](PublicEmail.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewVersionPublicEmail

`func NewVersionPublicEmail(id string, user VersionUser, object PublicEmail, updatedAt time.Time, ) *VersionPublicEmail`

NewVersionPublicEmail instantiates a new VersionPublicEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionPublicEmailWithDefaults

`func NewVersionPublicEmailWithDefaults() *VersionPublicEmail`

NewVersionPublicEmailWithDefaults instantiates a new VersionPublicEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionPublicEmail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionPublicEmail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionPublicEmail) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *VersionPublicEmail) GetUser() VersionUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VersionPublicEmail) GetUserOk() (*VersionUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VersionPublicEmail) SetUser(v VersionUser)`

SetUser sets User field to given value.


### GetObject

`func (o *VersionPublicEmail) GetObject() PublicEmail`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VersionPublicEmail) GetObjectOk() (*PublicEmail, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VersionPublicEmail) SetObject(v PublicEmail)`

SetObject sets Object field to given value.


### GetUpdatedAt

`func (o *VersionPublicEmail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VersionPublicEmail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VersionPublicEmail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


