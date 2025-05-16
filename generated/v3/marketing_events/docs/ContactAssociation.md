# ContactAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** |  | [optional] 
**ContactId** | **string** |  | 
**Email** | **string** |  | 
**Lastname** | Pointer to **string** |  | [optional] 

## Methods

### NewContactAssociation

`func NewContactAssociation(contactId string, email string, ) *ContactAssociation`

NewContactAssociation instantiates a new ContactAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactAssociationWithDefaults

`func NewContactAssociationWithDefaults() *ContactAssociation`

NewContactAssociationWithDefaults instantiates a new ContactAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *ContactAssociation) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *ContactAssociation) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *ContactAssociation) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *ContactAssociation) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetContactId

`func (o *ContactAssociation) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ContactAssociation) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ContactAssociation) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetEmail

`func (o *ContactAssociation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactAssociation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactAssociation) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLastname

`func (o *ContactAssociation) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *ContactAssociation) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *ContactAssociation) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *ContactAssociation) HasLastname() bool`

HasLastname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


