# PublicUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**PrimaryTeamId** | Pointer to **string** | The user&#39;s primary team | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**SendWelcomeEmail** | Pointer to **bool** |  | [optional] 
**RoleId** | Pointer to **string** | The user&#39;s role | [optional] 
**SecondaryTeamIds** | Pointer to **[]string** | The user&#39;s additional teams | [optional] 
**Id** | **string** | The user&#39;s unique ID | 
**SuperAdmin** | Pointer to **bool** |  | [optional] 
**Email** | **string** | The user&#39;s email | 

## Methods

### NewPublicUser

`func NewPublicUser(id string, email string, ) *PublicUser`

NewPublicUser instantiates a new PublicUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUserWithDefaults

`func NewPublicUserWithDefaults() *PublicUser`

NewPublicUserWithDefaults instantiates a new PublicUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *PublicUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PublicUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PublicUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PublicUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PublicUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PublicUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PublicUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PublicUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPrimaryTeamId

`func (o *PublicUser) GetPrimaryTeamId() string`

GetPrimaryTeamId returns the PrimaryTeamId field if non-nil, zero value otherwise.

### GetPrimaryTeamIdOk

`func (o *PublicUser) GetPrimaryTeamIdOk() (*string, bool)`

GetPrimaryTeamIdOk returns a tuple with the PrimaryTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTeamId

`func (o *PublicUser) SetPrimaryTeamId(v string)`

SetPrimaryTeamId sets PrimaryTeamId field to given value.

### HasPrimaryTeamId

`func (o *PublicUser) HasPrimaryTeamId() bool`

HasPrimaryTeamId returns a boolean if a field has been set.

### GetRoleIds

`func (o *PublicUser) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *PublicUser) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *PublicUser) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *PublicUser) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetSendWelcomeEmail

`func (o *PublicUser) GetSendWelcomeEmail() bool`

GetSendWelcomeEmail returns the SendWelcomeEmail field if non-nil, zero value otherwise.

### GetSendWelcomeEmailOk

`func (o *PublicUser) GetSendWelcomeEmailOk() (*bool, bool)`

GetSendWelcomeEmailOk returns a tuple with the SendWelcomeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWelcomeEmail

`func (o *PublicUser) SetSendWelcomeEmail(v bool)`

SetSendWelcomeEmail sets SendWelcomeEmail field to given value.

### HasSendWelcomeEmail

`func (o *PublicUser) HasSendWelcomeEmail() bool`

HasSendWelcomeEmail returns a boolean if a field has been set.

### GetRoleId

`func (o *PublicUser) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *PublicUser) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *PublicUser) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *PublicUser) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSecondaryTeamIds

`func (o *PublicUser) GetSecondaryTeamIds() []string`

GetSecondaryTeamIds returns the SecondaryTeamIds field if non-nil, zero value otherwise.

### GetSecondaryTeamIdsOk

`func (o *PublicUser) GetSecondaryTeamIdsOk() (*[]string, bool)`

GetSecondaryTeamIdsOk returns a tuple with the SecondaryTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTeamIds

`func (o *PublicUser) SetSecondaryTeamIds(v []string)`

SetSecondaryTeamIds sets SecondaryTeamIds field to given value.

### HasSecondaryTeamIds

`func (o *PublicUser) HasSecondaryTeamIds() bool`

HasSecondaryTeamIds returns a boolean if a field has been set.

### GetId

`func (o *PublicUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicUser) SetId(v string)`

SetId sets Id field to given value.


### GetSuperAdmin

`func (o *PublicUser) GetSuperAdmin() bool`

GetSuperAdmin returns the SuperAdmin field if non-nil, zero value otherwise.

### GetSuperAdminOk

`func (o *PublicUser) GetSuperAdminOk() (*bool, bool)`

GetSuperAdminOk returns a tuple with the SuperAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperAdmin

`func (o *PublicUser) SetSuperAdmin(v bool)`

SetSuperAdmin sets SuperAdmin field to given value.

### HasSuperAdmin

`func (o *PublicUser) HasSuperAdmin() bool`

HasSuperAdmin returns a boolean if a field has been set.

### GetEmail

`func (o *PublicUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PublicUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PublicUser) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


