# ExternalLinkMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIdsOfLinkMembers** | **[]string** |  | 
**CreatedAt** | **time.Time** |  | 
**Link** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**OrganizerUserId** | **string** |  | 
**Id** | **string** |  | 
**DefaultLink** | **bool** |  | 
**Type** | **string** |  | 
**Slug** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewExternalLinkMetadata

`func NewExternalLinkMetadata(userIdsOfLinkMembers []string, createdAt time.Time, link string, organizerUserId string, id string, defaultLink bool, type_ string, slug string, ) *ExternalLinkMetadata`

NewExternalLinkMetadata instantiates a new ExternalLinkMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalLinkMetadataWithDefaults

`func NewExternalLinkMetadataWithDefaults() *ExternalLinkMetadata`

NewExternalLinkMetadataWithDefaults instantiates a new ExternalLinkMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIdsOfLinkMembers

`func (o *ExternalLinkMetadata) GetUserIdsOfLinkMembers() []string`

GetUserIdsOfLinkMembers returns the UserIdsOfLinkMembers field if non-nil, zero value otherwise.

### GetUserIdsOfLinkMembersOk

`func (o *ExternalLinkMetadata) GetUserIdsOfLinkMembersOk() (*[]string, bool)`

GetUserIdsOfLinkMembersOk returns a tuple with the UserIdsOfLinkMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdsOfLinkMembers

`func (o *ExternalLinkMetadata) SetUserIdsOfLinkMembers(v []string)`

SetUserIdsOfLinkMembers sets UserIdsOfLinkMembers field to given value.


### GetCreatedAt

`func (o *ExternalLinkMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalLinkMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalLinkMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLink

`func (o *ExternalLinkMetadata) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ExternalLinkMetadata) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ExternalLinkMetadata) SetLink(v string)`

SetLink sets Link field to given value.


### GetName

`func (o *ExternalLinkMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalLinkMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalLinkMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalLinkMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizerUserId

`func (o *ExternalLinkMetadata) GetOrganizerUserId() string`

GetOrganizerUserId returns the OrganizerUserId field if non-nil, zero value otherwise.

### GetOrganizerUserIdOk

`func (o *ExternalLinkMetadata) GetOrganizerUserIdOk() (*string, bool)`

GetOrganizerUserIdOk returns a tuple with the OrganizerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerUserId

`func (o *ExternalLinkMetadata) SetOrganizerUserId(v string)`

SetOrganizerUserId sets OrganizerUserId field to given value.


### GetId

`func (o *ExternalLinkMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalLinkMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalLinkMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetDefaultLink

`func (o *ExternalLinkMetadata) GetDefaultLink() bool`

GetDefaultLink returns the DefaultLink field if non-nil, zero value otherwise.

### GetDefaultLinkOk

`func (o *ExternalLinkMetadata) GetDefaultLinkOk() (*bool, bool)`

GetDefaultLinkOk returns a tuple with the DefaultLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLink

`func (o *ExternalLinkMetadata) SetDefaultLink(v bool)`

SetDefaultLink sets DefaultLink field to given value.


### GetType

`func (o *ExternalLinkMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalLinkMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalLinkMetadata) SetType(v string)`

SetType sets Type field to given value.


### GetSlug

`func (o *ExternalLinkMetadata) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ExternalLinkMetadata) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ExternalLinkMetadata) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetUpdatedAt

`func (o *ExternalLinkMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExternalLinkMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExternalLinkMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExternalLinkMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


