# VisitorActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "VISITOR"]
**Id** | **string** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 
**Avatar** | Pointer to **string** |  | [optional] 

## Methods

### NewVisitorActor

`func NewVisitorActor(type_ string, id string, name string, email string, ) *VisitorActor`

NewVisitorActor instantiates a new VisitorActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisitorActorWithDefaults

`func NewVisitorActorWithDefaults() *VisitorActor`

NewVisitorActorWithDefaults instantiates a new VisitorActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VisitorActor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VisitorActor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VisitorActor) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *VisitorActor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VisitorActor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VisitorActor) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VisitorActor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VisitorActor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VisitorActor) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *VisitorActor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VisitorActor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VisitorActor) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAvatar

`func (o *VisitorActor) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *VisitorActor) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *VisitorActor) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *VisitorActor) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


