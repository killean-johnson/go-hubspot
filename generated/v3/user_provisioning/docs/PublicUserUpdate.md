# PublicUserUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**PrimaryTeamId** | Pointer to **string** | The user&#39;s primary team | [optional] 
**RoleId** | Pointer to **string** | The user&#39;s role | [optional] 
**SecondaryTeamIds** | Pointer to **[]string** | The user&#39;s additional teams | [optional] 

## Methods

### NewPublicUserUpdate

`func NewPublicUserUpdate() *PublicUserUpdate`

NewPublicUserUpdate instantiates a new PublicUserUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUserUpdateWithDefaults

`func NewPublicUserUpdateWithDefaults() *PublicUserUpdate`

NewPublicUserUpdateWithDefaults instantiates a new PublicUserUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *PublicUserUpdate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PublicUserUpdate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PublicUserUpdate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PublicUserUpdate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PublicUserUpdate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PublicUserUpdate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PublicUserUpdate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PublicUserUpdate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPrimaryTeamId

`func (o *PublicUserUpdate) GetPrimaryTeamId() string`

GetPrimaryTeamId returns the PrimaryTeamId field if non-nil, zero value otherwise.

### GetPrimaryTeamIdOk

`func (o *PublicUserUpdate) GetPrimaryTeamIdOk() (*string, bool)`

GetPrimaryTeamIdOk returns a tuple with the PrimaryTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTeamId

`func (o *PublicUserUpdate) SetPrimaryTeamId(v string)`

SetPrimaryTeamId sets PrimaryTeamId field to given value.

### HasPrimaryTeamId

`func (o *PublicUserUpdate) HasPrimaryTeamId() bool`

HasPrimaryTeamId returns a boolean if a field has been set.

### GetRoleId

`func (o *PublicUserUpdate) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *PublicUserUpdate) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *PublicUserUpdate) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *PublicUserUpdate) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSecondaryTeamIds

`func (o *PublicUserUpdate) GetSecondaryTeamIds() []string`

GetSecondaryTeamIds returns the SecondaryTeamIds field if non-nil, zero value otherwise.

### GetSecondaryTeamIdsOk

`func (o *PublicUserUpdate) GetSecondaryTeamIdsOk() (*[]string, bool)`

GetSecondaryTeamIdsOk returns a tuple with the SecondaryTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTeamIds

`func (o *PublicUserUpdate) SetSecondaryTeamIds(v []string)`

SetSecondaryTeamIds sets SecondaryTeamIds field to given value.

### HasSecondaryTeamIds

`func (o *PublicUserUpdate) HasSecondaryTeamIds() bool`

HasSecondaryTeamIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


