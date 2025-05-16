# PublicSequenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Settings** | Pointer to [**PublicSequenceSettingsResponse**](PublicSequenceSettingsResponse.md) |  | [optional] 
**Name** | **string** |  | 
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Steps** | [**[]PublicSequenceStepResponse**](PublicSequenceStepResponse.md) |  | 
**FolderId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**Dependencies** | [**[]PublicSequenceStepDependencyResponse**](PublicSequenceStepDependencyResponse.md) |  | 

## Methods

### NewPublicSequenceResponse

`func NewPublicSequenceResponse(createdAt time.Time, name string, id string, userId string, steps []PublicSequenceStepResponse, updatedAt time.Time, dependencies []PublicSequenceStepDependencyResponse, ) *PublicSequenceResponse`

NewPublicSequenceResponse instantiates a new PublicSequenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSequenceResponseWithDefaults

`func NewPublicSequenceResponseWithDefaults() *PublicSequenceResponse`

NewPublicSequenceResponseWithDefaults instantiates a new PublicSequenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicSequenceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicSequenceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicSequenceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSettings

`func (o *PublicSequenceResponse) GetSettings() PublicSequenceSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PublicSequenceResponse) GetSettingsOk() (*PublicSequenceSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PublicSequenceResponse) SetSettings(v PublicSequenceSettingsResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PublicSequenceResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetName

`func (o *PublicSequenceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicSequenceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicSequenceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *PublicSequenceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicSequenceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicSequenceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *PublicSequenceResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicSequenceResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicSequenceResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSteps

`func (o *PublicSequenceResponse) GetSteps() []PublicSequenceStepResponse`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *PublicSequenceResponse) GetStepsOk() (*[]PublicSequenceStepResponse, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *PublicSequenceResponse) SetSteps(v []PublicSequenceStepResponse)`

SetSteps sets Steps field to given value.


### GetFolderId

`func (o *PublicSequenceResponse) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *PublicSequenceResponse) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *PublicSequenceResponse) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *PublicSequenceResponse) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicSequenceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicSequenceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicSequenceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDependencies

`func (o *PublicSequenceResponse) GetDependencies() []PublicSequenceStepDependencyResponse`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *PublicSequenceResponse) GetDependenciesOk() (*[]PublicSequenceStepDependencyResponse, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *PublicSequenceResponse) SetDependencies(v []PublicSequenceStepDependencyResponse)`

SetDependencies sets Dependencies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


