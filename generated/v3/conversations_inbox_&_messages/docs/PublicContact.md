# PublicContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "CONTACT"]
**ContactProfile** | [**ContactProfile**](ContactProfile.md) |  | 

## Methods

### NewPublicContact

`func NewPublicContact(type_ string, contactProfile ContactProfile, ) *PublicContact`

NewPublicContact instantiates a new PublicContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicContactWithDefaults

`func NewPublicContactWithDefaults() *PublicContact`

NewPublicContactWithDefaults instantiates a new PublicContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicContact) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicContact) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicContact) SetType(v string)`

SetType sets Type field to given value.


### GetContactProfile

`func (o *PublicContact) GetContactProfile() ContactProfile`

GetContactProfile returns the ContactProfile field if non-nil, zero value otherwise.

### GetContactProfileOk

`func (o *PublicContact) GetContactProfileOk() (*ContactProfile, bool)`

GetContactProfileOk returns a tuple with the ContactProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactProfile

`func (o *PublicContact) SetContactProfile(v ContactProfile)`

SetContactProfile sets ContactProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


