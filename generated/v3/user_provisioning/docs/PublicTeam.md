# PublicTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | **[]string** | Primary members of this team | 
**Name** | **string** | The team&#39;s name | 
**Id** | **string** | The team&#39;s unique ID | 
**SecondaryUserIds** | **[]string** | Secondary or additional members of this team | 

## Methods

### NewPublicTeam

`func NewPublicTeam(userIds []string, name string, id string, secondaryUserIds []string, ) *PublicTeam`

NewPublicTeam instantiates a new PublicTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTeamWithDefaults

`func NewPublicTeamWithDefaults() *PublicTeam`

NewPublicTeamWithDefaults instantiates a new PublicTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *PublicTeam) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *PublicTeam) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *PublicTeam) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.


### GetName

`func (o *PublicTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicTeam) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *PublicTeam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicTeam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicTeam) SetId(v string)`

SetId sets Id field to given value.


### GetSecondaryUserIds

`func (o *PublicTeam) GetSecondaryUserIds() []string`

GetSecondaryUserIds returns the SecondaryUserIds field if non-nil, zero value otherwise.

### GetSecondaryUserIdsOk

`func (o *PublicTeam) GetSecondaryUserIdsOk() (*[]string, bool)`

GetSecondaryUserIdsOk returns a tuple with the SecondaryUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryUserIds

`func (o *PublicTeam) SetSecondaryUserIds(v []string)`

SetSecondaryUserIds sets SecondaryUserIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


