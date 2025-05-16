# ExternalBehavioralEventTypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | [**[]AssociationDefinition**](AssociationDefinition.md) |  | 
**ObjectTypeId** | **string** |  | 
**CreatedUserId** | Pointer to **int32** |  | [optional] 
**TrackingType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PrimaryObjectId** | Pointer to **string** |  | [optional] 
**ComboEventRules** | Pointer to [**ComboEventRuleBranch**](ComboEventRuleBranch.md) |  | [optional] 
**FullyQualifiedName** | **string** |  | 
**PrimaryObject** | Pointer to **string** |  | [optional] 
**Labels** | [**BehavioralEventTypeDefinitionLabels**](BehavioralEventTypeDefinitionLabels.md) |  | 
**Archived** | **bool** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Id** | **string** |  | 
**Properties** | [**[]Property**](Property.md) |  | 

## Methods

### NewExternalBehavioralEventTypeDefinition

`func NewExternalBehavioralEventTypeDefinition(associations []AssociationDefinition, objectTypeId string, fullyQualifiedName string, labels BehavioralEventTypeDefinitionLabels, archived bool, name string, id string, properties []Property, ) *ExternalBehavioralEventTypeDefinition`

NewExternalBehavioralEventTypeDefinition instantiates a new ExternalBehavioralEventTypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalBehavioralEventTypeDefinitionWithDefaults

`func NewExternalBehavioralEventTypeDefinitionWithDefaults() *ExternalBehavioralEventTypeDefinition`

NewExternalBehavioralEventTypeDefinitionWithDefaults instantiates a new ExternalBehavioralEventTypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *ExternalBehavioralEventTypeDefinition) GetAssociations() []AssociationDefinition`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *ExternalBehavioralEventTypeDefinition) GetAssociationsOk() (*[]AssociationDefinition, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *ExternalBehavioralEventTypeDefinition) SetAssociations(v []AssociationDefinition)`

SetAssociations sets Associations field to given value.


### GetObjectTypeId

`func (o *ExternalBehavioralEventTypeDefinition) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ExternalBehavioralEventTypeDefinition) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ExternalBehavioralEventTypeDefinition) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetCreatedUserId

`func (o *ExternalBehavioralEventTypeDefinition) GetCreatedUserId() int32`

GetCreatedUserId returns the CreatedUserId field if non-nil, zero value otherwise.

### GetCreatedUserIdOk

`func (o *ExternalBehavioralEventTypeDefinition) GetCreatedUserIdOk() (*int32, bool)`

GetCreatedUserIdOk returns a tuple with the CreatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUserId

`func (o *ExternalBehavioralEventTypeDefinition) SetCreatedUserId(v int32)`

SetCreatedUserId sets CreatedUserId field to given value.

### HasCreatedUserId

`func (o *ExternalBehavioralEventTypeDefinition) HasCreatedUserId() bool`

HasCreatedUserId returns a boolean if a field has been set.

### GetTrackingType

`func (o *ExternalBehavioralEventTypeDefinition) GetTrackingType() string`

GetTrackingType returns the TrackingType field if non-nil, zero value otherwise.

### GetTrackingTypeOk

`func (o *ExternalBehavioralEventTypeDefinition) GetTrackingTypeOk() (*string, bool)`

GetTrackingTypeOk returns a tuple with the TrackingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingType

`func (o *ExternalBehavioralEventTypeDefinition) SetTrackingType(v string)`

SetTrackingType sets TrackingType field to given value.

### HasTrackingType

`func (o *ExternalBehavioralEventTypeDefinition) HasTrackingType() bool`

HasTrackingType returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalBehavioralEventTypeDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalBehavioralEventTypeDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalBehavioralEventTypeDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalBehavioralEventTypeDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrimaryObjectId

`func (o *ExternalBehavioralEventTypeDefinition) GetPrimaryObjectId() string`

GetPrimaryObjectId returns the PrimaryObjectId field if non-nil, zero value otherwise.

### GetPrimaryObjectIdOk

`func (o *ExternalBehavioralEventTypeDefinition) GetPrimaryObjectIdOk() (*string, bool)`

GetPrimaryObjectIdOk returns a tuple with the PrimaryObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryObjectId

`func (o *ExternalBehavioralEventTypeDefinition) SetPrimaryObjectId(v string)`

SetPrimaryObjectId sets PrimaryObjectId field to given value.

### HasPrimaryObjectId

`func (o *ExternalBehavioralEventTypeDefinition) HasPrimaryObjectId() bool`

HasPrimaryObjectId returns a boolean if a field has been set.

### GetComboEventRules

`func (o *ExternalBehavioralEventTypeDefinition) GetComboEventRules() ComboEventRuleBranch`

GetComboEventRules returns the ComboEventRules field if non-nil, zero value otherwise.

### GetComboEventRulesOk

`func (o *ExternalBehavioralEventTypeDefinition) GetComboEventRulesOk() (*ComboEventRuleBranch, bool)`

GetComboEventRulesOk returns a tuple with the ComboEventRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboEventRules

`func (o *ExternalBehavioralEventTypeDefinition) SetComboEventRules(v ComboEventRuleBranch)`

SetComboEventRules sets ComboEventRules field to given value.

### HasComboEventRules

`func (o *ExternalBehavioralEventTypeDefinition) HasComboEventRules() bool`

HasComboEventRules returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *ExternalBehavioralEventTypeDefinition) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *ExternalBehavioralEventTypeDefinition) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *ExternalBehavioralEventTypeDefinition) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.


### GetPrimaryObject

`func (o *ExternalBehavioralEventTypeDefinition) GetPrimaryObject() string`

GetPrimaryObject returns the PrimaryObject field if non-nil, zero value otherwise.

### GetPrimaryObjectOk

`func (o *ExternalBehavioralEventTypeDefinition) GetPrimaryObjectOk() (*string, bool)`

GetPrimaryObjectOk returns a tuple with the PrimaryObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryObject

`func (o *ExternalBehavioralEventTypeDefinition) SetPrimaryObject(v string)`

SetPrimaryObject sets PrimaryObject field to given value.

### HasPrimaryObject

`func (o *ExternalBehavioralEventTypeDefinition) HasPrimaryObject() bool`

HasPrimaryObject returns a boolean if a field has been set.

### GetLabels

`func (o *ExternalBehavioralEventTypeDefinition) GetLabels() BehavioralEventTypeDefinitionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ExternalBehavioralEventTypeDefinition) GetLabelsOk() (*BehavioralEventTypeDefinitionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ExternalBehavioralEventTypeDefinition) SetLabels(v BehavioralEventTypeDefinitionLabels)`

SetLabels sets Labels field to given value.


### GetArchived

`func (o *ExternalBehavioralEventTypeDefinition) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ExternalBehavioralEventTypeDefinition) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ExternalBehavioralEventTypeDefinition) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCreatedAt

`func (o *ExternalBehavioralEventTypeDefinition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalBehavioralEventTypeDefinition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalBehavioralEventTypeDefinition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExternalBehavioralEventTypeDefinition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *ExternalBehavioralEventTypeDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalBehavioralEventTypeDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalBehavioralEventTypeDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *ExternalBehavioralEventTypeDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalBehavioralEventTypeDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalBehavioralEventTypeDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *ExternalBehavioralEventTypeDefinition) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ExternalBehavioralEventTypeDefinition) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ExternalBehavioralEventTypeDefinition) SetProperties(v []Property)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


