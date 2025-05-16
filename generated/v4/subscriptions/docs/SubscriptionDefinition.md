# SubscriptionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsInternal** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**IsDefault** | **bool** |  | 
**CommunicationMethod** | Pointer to **string** |  | [optional] 
**Purpose** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Id** | **string** |  | 
**IsActive** | **bool** |  | 
**BusinessUnitId** | Pointer to **int64** |  | [optional] 
**SubscriptionTranslations** | Pointer to [**[]PublicSubscriptionTranslation**](PublicSubscriptionTranslation.md) |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewSubscriptionDefinition

`func NewSubscriptionDefinition(isInternal bool, createdAt time.Time, isDefault bool, name string, description string, id string, isActive bool, updatedAt time.Time, ) *SubscriptionDefinition`

NewSubscriptionDefinition instantiates a new SubscriptionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDefinitionWithDefaults

`func NewSubscriptionDefinitionWithDefaults() *SubscriptionDefinition`

NewSubscriptionDefinitionWithDefaults instantiates a new SubscriptionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsInternal

`func (o *SubscriptionDefinition) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *SubscriptionDefinition) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *SubscriptionDefinition) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.


### GetCreatedAt

`func (o *SubscriptionDefinition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionDefinition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionDefinition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsDefault

`func (o *SubscriptionDefinition) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SubscriptionDefinition) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SubscriptionDefinition) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCommunicationMethod

`func (o *SubscriptionDefinition) GetCommunicationMethod() string`

GetCommunicationMethod returns the CommunicationMethod field if non-nil, zero value otherwise.

### GetCommunicationMethodOk

`func (o *SubscriptionDefinition) GetCommunicationMethodOk() (*string, bool)`

GetCommunicationMethodOk returns a tuple with the CommunicationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationMethod

`func (o *SubscriptionDefinition) SetCommunicationMethod(v string)`

SetCommunicationMethod sets CommunicationMethod field to given value.

### HasCommunicationMethod

`func (o *SubscriptionDefinition) HasCommunicationMethod() bool`

HasCommunicationMethod returns a boolean if a field has been set.

### GetPurpose

`func (o *SubscriptionDefinition) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *SubscriptionDefinition) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *SubscriptionDefinition) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *SubscriptionDefinition) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetName

`func (o *SubscriptionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SubscriptionDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *SubscriptionDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *SubscriptionDefinition) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SubscriptionDefinition) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SubscriptionDefinition) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetBusinessUnitId

`func (o *SubscriptionDefinition) GetBusinessUnitId() int64`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *SubscriptionDefinition) GetBusinessUnitIdOk() (*int64, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *SubscriptionDefinition) SetBusinessUnitId(v int64)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *SubscriptionDefinition) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetSubscriptionTranslations

`func (o *SubscriptionDefinition) GetSubscriptionTranslations() []PublicSubscriptionTranslation`

GetSubscriptionTranslations returns the SubscriptionTranslations field if non-nil, zero value otherwise.

### GetSubscriptionTranslationsOk

`func (o *SubscriptionDefinition) GetSubscriptionTranslationsOk() (*[]PublicSubscriptionTranslation, bool)`

GetSubscriptionTranslationsOk returns a tuple with the SubscriptionTranslations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTranslations

`func (o *SubscriptionDefinition) SetSubscriptionTranslations(v []PublicSubscriptionTranslation)`

SetSubscriptionTranslations sets SubscriptionTranslations field to given value.

### HasSubscriptionTranslations

`func (o *SubscriptionDefinition) HasSubscriptionTranslations() bool`

HasSubscriptionTranslations returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionDefinition) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionDefinition) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionDefinition) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


