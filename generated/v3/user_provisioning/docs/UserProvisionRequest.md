# UserProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**PrimaryTeamId** | Pointer to **string** | The user&#39;s primary team | [optional] 
**SendWelcomeEmail** | Pointer to **bool** | Whether to send a welcome email | [optional] 
**RoleId** | Pointer to **string** | The user&#39;s role | [optional] 
**SecondaryTeamIds** | Pointer to **[]string** | The user&#39;s additional teams | [optional] 
**Email** | **string** | The created user&#39;s email | 

## Methods

### NewUserProvisionRequest

`func NewUserProvisionRequest(email string, ) *UserProvisionRequest`

NewUserProvisionRequest instantiates a new UserProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProvisionRequestWithDefaults

`func NewUserProvisionRequestWithDefaults() *UserProvisionRequest`

NewUserProvisionRequestWithDefaults instantiates a new UserProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UserProvisionRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserProvisionRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserProvisionRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserProvisionRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserProvisionRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserProvisionRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserProvisionRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserProvisionRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPrimaryTeamId

`func (o *UserProvisionRequest) GetPrimaryTeamId() string`

GetPrimaryTeamId returns the PrimaryTeamId field if non-nil, zero value otherwise.

### GetPrimaryTeamIdOk

`func (o *UserProvisionRequest) GetPrimaryTeamIdOk() (*string, bool)`

GetPrimaryTeamIdOk returns a tuple with the PrimaryTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTeamId

`func (o *UserProvisionRequest) SetPrimaryTeamId(v string)`

SetPrimaryTeamId sets PrimaryTeamId field to given value.

### HasPrimaryTeamId

`func (o *UserProvisionRequest) HasPrimaryTeamId() bool`

HasPrimaryTeamId returns a boolean if a field has been set.

### GetSendWelcomeEmail

`func (o *UserProvisionRequest) GetSendWelcomeEmail() bool`

GetSendWelcomeEmail returns the SendWelcomeEmail field if non-nil, zero value otherwise.

### GetSendWelcomeEmailOk

`func (o *UserProvisionRequest) GetSendWelcomeEmailOk() (*bool, bool)`

GetSendWelcomeEmailOk returns a tuple with the SendWelcomeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWelcomeEmail

`func (o *UserProvisionRequest) SetSendWelcomeEmail(v bool)`

SetSendWelcomeEmail sets SendWelcomeEmail field to given value.

### HasSendWelcomeEmail

`func (o *UserProvisionRequest) HasSendWelcomeEmail() bool`

HasSendWelcomeEmail returns a boolean if a field has been set.

### GetRoleId

`func (o *UserProvisionRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UserProvisionRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UserProvisionRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UserProvisionRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSecondaryTeamIds

`func (o *UserProvisionRequest) GetSecondaryTeamIds() []string`

GetSecondaryTeamIds returns the SecondaryTeamIds field if non-nil, zero value otherwise.

### GetSecondaryTeamIdsOk

`func (o *UserProvisionRequest) GetSecondaryTeamIdsOk() (*[]string, bool)`

GetSecondaryTeamIdsOk returns a tuple with the SecondaryTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTeamIds

`func (o *UserProvisionRequest) SetSecondaryTeamIds(v []string)`

SetSecondaryTeamIds sets SecondaryTeamIds field to given value.

### HasSecondaryTeamIds

`func (o *UserProvisionRequest) HasSecondaryTeamIds() bool`

HasSecondaryTeamIds returns a boolean if a field has been set.

### GetEmail

`func (o *UserProvisionRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserProvisionRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserProvisionRequest) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


