# ObjectTypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecondaryDisplayProperties** | **[]string** |  | 
**ObjectTypeId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**FullyQualifiedName** | **string** |  | 
**Labels** | [**ObjectTypeDefinitionLabels**](ObjectTypeDefinitionLabels.md) |  | 
**Archived** | **bool** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**RequiredProperties** | **[]string** |  | 
**SearchableProperties** | **[]string** |  | 
**PortalId** | Pointer to **int32** |  | [optional] 
**PrimaryDisplayProperty** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Id** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewObjectTypeDefinition

`func NewObjectTypeDefinition(secondaryDisplayProperties []string, objectTypeId string, fullyQualifiedName string, labels ObjectTypeDefinitionLabels, archived bool, requiredProperties []string, searchableProperties []string, name string, id string, ) *ObjectTypeDefinition`

NewObjectTypeDefinition instantiates a new ObjectTypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectTypeDefinitionWithDefaults

`func NewObjectTypeDefinitionWithDefaults() *ObjectTypeDefinition`

NewObjectTypeDefinitionWithDefaults instantiates a new ObjectTypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecondaryDisplayProperties

`func (o *ObjectTypeDefinition) GetSecondaryDisplayProperties() []string`

GetSecondaryDisplayProperties returns the SecondaryDisplayProperties field if non-nil, zero value otherwise.

### GetSecondaryDisplayPropertiesOk

`func (o *ObjectTypeDefinition) GetSecondaryDisplayPropertiesOk() (*[]string, bool)`

GetSecondaryDisplayPropertiesOk returns a tuple with the SecondaryDisplayProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDisplayProperties

`func (o *ObjectTypeDefinition) SetSecondaryDisplayProperties(v []string)`

SetSecondaryDisplayProperties sets SecondaryDisplayProperties field to given value.


### GetObjectTypeId

`func (o *ObjectTypeDefinition) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ObjectTypeDefinition) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ObjectTypeDefinition) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetDescription

`func (o *ObjectTypeDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectTypeDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectTypeDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectTypeDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *ObjectTypeDefinition) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *ObjectTypeDefinition) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *ObjectTypeDefinition) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.


### GetLabels

`func (o *ObjectTypeDefinition) GetLabels() ObjectTypeDefinitionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ObjectTypeDefinition) GetLabelsOk() (*ObjectTypeDefinitionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ObjectTypeDefinition) SetLabels(v ObjectTypeDefinitionLabels)`

SetLabels sets Labels field to given value.


### GetArchived

`func (o *ObjectTypeDefinition) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ObjectTypeDefinition) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ObjectTypeDefinition) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCreatedAt

`func (o *ObjectTypeDefinition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ObjectTypeDefinition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ObjectTypeDefinition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ObjectTypeDefinition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRequiredProperties

`func (o *ObjectTypeDefinition) GetRequiredProperties() []string`

GetRequiredProperties returns the RequiredProperties field if non-nil, zero value otherwise.

### GetRequiredPropertiesOk

`func (o *ObjectTypeDefinition) GetRequiredPropertiesOk() (*[]string, bool)`

GetRequiredPropertiesOk returns a tuple with the RequiredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredProperties

`func (o *ObjectTypeDefinition) SetRequiredProperties(v []string)`

SetRequiredProperties sets RequiredProperties field to given value.


### GetSearchableProperties

`func (o *ObjectTypeDefinition) GetSearchableProperties() []string`

GetSearchableProperties returns the SearchableProperties field if non-nil, zero value otherwise.

### GetSearchablePropertiesOk

`func (o *ObjectTypeDefinition) GetSearchablePropertiesOk() (*[]string, bool)`

GetSearchablePropertiesOk returns a tuple with the SearchableProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableProperties

`func (o *ObjectTypeDefinition) SetSearchableProperties(v []string)`

SetSearchableProperties sets SearchableProperties field to given value.


### GetPortalId

`func (o *ObjectTypeDefinition) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *ObjectTypeDefinition) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *ObjectTypeDefinition) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *ObjectTypeDefinition) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetPrimaryDisplayProperty

`func (o *ObjectTypeDefinition) GetPrimaryDisplayProperty() string`

GetPrimaryDisplayProperty returns the PrimaryDisplayProperty field if non-nil, zero value otherwise.

### GetPrimaryDisplayPropertyOk

`func (o *ObjectTypeDefinition) GetPrimaryDisplayPropertyOk() (*string, bool)`

GetPrimaryDisplayPropertyOk returns a tuple with the PrimaryDisplayProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDisplayProperty

`func (o *ObjectTypeDefinition) SetPrimaryDisplayProperty(v string)`

SetPrimaryDisplayProperty sets PrimaryDisplayProperty field to given value.

### HasPrimaryDisplayProperty

`func (o *ObjectTypeDefinition) HasPrimaryDisplayProperty() bool`

HasPrimaryDisplayProperty returns a boolean if a field has been set.

### GetName

`func (o *ObjectTypeDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectTypeDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectTypeDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *ObjectTypeDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectTypeDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectTypeDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *ObjectTypeDefinition) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ObjectTypeDefinition) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ObjectTypeDefinition) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ObjectTypeDefinition) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


