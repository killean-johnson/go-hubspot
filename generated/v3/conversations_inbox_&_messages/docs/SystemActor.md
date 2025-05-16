# SystemActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "SYSTEM"]
**Id** | **string** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 
**Avatar** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemActor

`func NewSystemActor(type_ string, id string, name string, email string, ) *SystemActor`

NewSystemActor instantiates a new SystemActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemActorWithDefaults

`func NewSystemActorWithDefaults() *SystemActor`

NewSystemActorWithDefaults instantiates a new SystemActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SystemActor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemActor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemActor) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *SystemActor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemActor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemActor) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SystemActor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemActor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemActor) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *SystemActor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SystemActor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SystemActor) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAvatar

`func (o *SystemActor) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *SystemActor) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *SystemActor) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *SystemActor) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


