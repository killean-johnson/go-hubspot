# AssociationDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InverseLabel** | Pointer to **string** |  | [optional] 
**Hidden** | **bool** |  | 
**AllowsCustomLabels** | **bool** |  | 
**MaxFromObjectIds** | **int32** |  | 
**InverseAllowsCustomLabels** | **bool** |  | 
**IsPrimary** | **bool** |  | 
**HasUserEnforcedMaxToObjectIds** | **bool** |  | 
**HasUserEnforcedMaxFromObjectIds** | **bool** |  | 
**Id** | **int32** |  | 
**ToObjectTypeId** | **string** |  | 
**InverseCardinality** | **string** |  | 
**MaxToObjectIds** | **int32** |  | 
**HasCascadingDeletes** | **bool** |  | 
**FromObjectType** | Pointer to **string** |  | [optional] 
**ToObjectType** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**InverseId** | **int32** |  | 
**Cardinality** | **string** |  | 
**InverseName** | **string** |  | 
**HasAllAssociatedObjects** | **bool** |  | 
**InverseHasAllAssociatedObjects** | **bool** |  | 
**FromObjectTypeId** | **string** |  | 
**PortalUniqueIdentifier** | **string** |  | 
**IsInversePrimary** | **bool** |  | 
**Name** | **string** |  | 
**Category** | **string** |  | 

## Methods

### NewAssociationDefinition

`func NewAssociationDefinition(hidden bool, allowsCustomLabels bool, maxFromObjectIds int32, inverseAllowsCustomLabels bool, isPrimary bool, hasUserEnforcedMaxToObjectIds bool, hasUserEnforcedMaxFromObjectIds bool, id int32, toObjectTypeId string, inverseCardinality string, maxToObjectIds int32, hasCascadingDeletes bool, inverseId int32, cardinality string, inverseName string, hasAllAssociatedObjects bool, inverseHasAllAssociatedObjects bool, fromObjectTypeId string, portalUniqueIdentifier string, isInversePrimary bool, name string, category string, ) *AssociationDefinition`

NewAssociationDefinition instantiates a new AssociationDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationDefinitionWithDefaults

`func NewAssociationDefinitionWithDefaults() *AssociationDefinition`

NewAssociationDefinitionWithDefaults instantiates a new AssociationDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInverseLabel

`func (o *AssociationDefinition) GetInverseLabel() string`

GetInverseLabel returns the InverseLabel field if non-nil, zero value otherwise.

### GetInverseLabelOk

`func (o *AssociationDefinition) GetInverseLabelOk() (*string, bool)`

GetInverseLabelOk returns a tuple with the InverseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseLabel

`func (o *AssociationDefinition) SetInverseLabel(v string)`

SetInverseLabel sets InverseLabel field to given value.

### HasInverseLabel

`func (o *AssociationDefinition) HasInverseLabel() bool`

HasInverseLabel returns a boolean if a field has been set.

### GetHidden

`func (o *AssociationDefinition) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *AssociationDefinition) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *AssociationDefinition) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetAllowsCustomLabels

`func (o *AssociationDefinition) GetAllowsCustomLabels() bool`

GetAllowsCustomLabels returns the AllowsCustomLabels field if non-nil, zero value otherwise.

### GetAllowsCustomLabelsOk

`func (o *AssociationDefinition) GetAllowsCustomLabelsOk() (*bool, bool)`

GetAllowsCustomLabelsOk returns a tuple with the AllowsCustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsCustomLabels

`func (o *AssociationDefinition) SetAllowsCustomLabels(v bool)`

SetAllowsCustomLabels sets AllowsCustomLabels field to given value.


### GetMaxFromObjectIds

`func (o *AssociationDefinition) GetMaxFromObjectIds() int32`

GetMaxFromObjectIds returns the MaxFromObjectIds field if non-nil, zero value otherwise.

### GetMaxFromObjectIdsOk

`func (o *AssociationDefinition) GetMaxFromObjectIdsOk() (*int32, bool)`

GetMaxFromObjectIdsOk returns a tuple with the MaxFromObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFromObjectIds

`func (o *AssociationDefinition) SetMaxFromObjectIds(v int32)`

SetMaxFromObjectIds sets MaxFromObjectIds field to given value.


### GetInverseAllowsCustomLabels

`func (o *AssociationDefinition) GetInverseAllowsCustomLabels() bool`

GetInverseAllowsCustomLabels returns the InverseAllowsCustomLabels field if non-nil, zero value otherwise.

### GetInverseAllowsCustomLabelsOk

`func (o *AssociationDefinition) GetInverseAllowsCustomLabelsOk() (*bool, bool)`

GetInverseAllowsCustomLabelsOk returns a tuple with the InverseAllowsCustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseAllowsCustomLabels

`func (o *AssociationDefinition) SetInverseAllowsCustomLabels(v bool)`

SetInverseAllowsCustomLabels sets InverseAllowsCustomLabels field to given value.


### GetIsPrimary

`func (o *AssociationDefinition) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *AssociationDefinition) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *AssociationDefinition) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### GetHasUserEnforcedMaxToObjectIds

`func (o *AssociationDefinition) GetHasUserEnforcedMaxToObjectIds() bool`

GetHasUserEnforcedMaxToObjectIds returns the HasUserEnforcedMaxToObjectIds field if non-nil, zero value otherwise.

### GetHasUserEnforcedMaxToObjectIdsOk

`func (o *AssociationDefinition) GetHasUserEnforcedMaxToObjectIdsOk() (*bool, bool)`

GetHasUserEnforcedMaxToObjectIdsOk returns a tuple with the HasUserEnforcedMaxToObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUserEnforcedMaxToObjectIds

`func (o *AssociationDefinition) SetHasUserEnforcedMaxToObjectIds(v bool)`

SetHasUserEnforcedMaxToObjectIds sets HasUserEnforcedMaxToObjectIds field to given value.


### GetHasUserEnforcedMaxFromObjectIds

`func (o *AssociationDefinition) GetHasUserEnforcedMaxFromObjectIds() bool`

GetHasUserEnforcedMaxFromObjectIds returns the HasUserEnforcedMaxFromObjectIds field if non-nil, zero value otherwise.

### GetHasUserEnforcedMaxFromObjectIdsOk

`func (o *AssociationDefinition) GetHasUserEnforcedMaxFromObjectIdsOk() (*bool, bool)`

GetHasUserEnforcedMaxFromObjectIdsOk returns a tuple with the HasUserEnforcedMaxFromObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUserEnforcedMaxFromObjectIds

`func (o *AssociationDefinition) SetHasUserEnforcedMaxFromObjectIds(v bool)`

SetHasUserEnforcedMaxFromObjectIds sets HasUserEnforcedMaxFromObjectIds field to given value.


### GetId

`func (o *AssociationDefinition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssociationDefinition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssociationDefinition) SetId(v int32)`

SetId sets Id field to given value.


### GetToObjectTypeId

`func (o *AssociationDefinition) GetToObjectTypeId() string`

GetToObjectTypeId returns the ToObjectTypeId field if non-nil, zero value otherwise.

### GetToObjectTypeIdOk

`func (o *AssociationDefinition) GetToObjectTypeIdOk() (*string, bool)`

GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectTypeId

`func (o *AssociationDefinition) SetToObjectTypeId(v string)`

SetToObjectTypeId sets ToObjectTypeId field to given value.


### GetInverseCardinality

`func (o *AssociationDefinition) GetInverseCardinality() string`

GetInverseCardinality returns the InverseCardinality field if non-nil, zero value otherwise.

### GetInverseCardinalityOk

`func (o *AssociationDefinition) GetInverseCardinalityOk() (*string, bool)`

GetInverseCardinalityOk returns a tuple with the InverseCardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseCardinality

`func (o *AssociationDefinition) SetInverseCardinality(v string)`

SetInverseCardinality sets InverseCardinality field to given value.


### GetMaxToObjectIds

`func (o *AssociationDefinition) GetMaxToObjectIds() int32`

GetMaxToObjectIds returns the MaxToObjectIds field if non-nil, zero value otherwise.

### GetMaxToObjectIdsOk

`func (o *AssociationDefinition) GetMaxToObjectIdsOk() (*int32, bool)`

GetMaxToObjectIdsOk returns a tuple with the MaxToObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxToObjectIds

`func (o *AssociationDefinition) SetMaxToObjectIds(v int32)`

SetMaxToObjectIds sets MaxToObjectIds field to given value.


### GetHasCascadingDeletes

`func (o *AssociationDefinition) GetHasCascadingDeletes() bool`

GetHasCascadingDeletes returns the HasCascadingDeletes field if non-nil, zero value otherwise.

### GetHasCascadingDeletesOk

`func (o *AssociationDefinition) GetHasCascadingDeletesOk() (*bool, bool)`

GetHasCascadingDeletesOk returns a tuple with the HasCascadingDeletes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCascadingDeletes

`func (o *AssociationDefinition) SetHasCascadingDeletes(v bool)`

SetHasCascadingDeletes sets HasCascadingDeletes field to given value.


### GetFromObjectType

`func (o *AssociationDefinition) GetFromObjectType() string`

GetFromObjectType returns the FromObjectType field if non-nil, zero value otherwise.

### GetFromObjectTypeOk

`func (o *AssociationDefinition) GetFromObjectTypeOk() (*string, bool)`

GetFromObjectTypeOk returns a tuple with the FromObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromObjectType

`func (o *AssociationDefinition) SetFromObjectType(v string)`

SetFromObjectType sets FromObjectType field to given value.

### HasFromObjectType

`func (o *AssociationDefinition) HasFromObjectType() bool`

HasFromObjectType returns a boolean if a field has been set.

### GetToObjectType

`func (o *AssociationDefinition) GetToObjectType() string`

GetToObjectType returns the ToObjectType field if non-nil, zero value otherwise.

### GetToObjectTypeOk

`func (o *AssociationDefinition) GetToObjectTypeOk() (*string, bool)`

GetToObjectTypeOk returns a tuple with the ToObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectType

`func (o *AssociationDefinition) SetToObjectType(v string)`

SetToObjectType sets ToObjectType field to given value.

### HasToObjectType

`func (o *AssociationDefinition) HasToObjectType() bool`

HasToObjectType returns a boolean if a field has been set.

### GetLabel

`func (o *AssociationDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AssociationDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AssociationDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AssociationDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetInverseId

`func (o *AssociationDefinition) GetInverseId() int32`

GetInverseId returns the InverseId field if non-nil, zero value otherwise.

### GetInverseIdOk

`func (o *AssociationDefinition) GetInverseIdOk() (*int32, bool)`

GetInverseIdOk returns a tuple with the InverseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseId

`func (o *AssociationDefinition) SetInverseId(v int32)`

SetInverseId sets InverseId field to given value.


### GetCardinality

`func (o *AssociationDefinition) GetCardinality() string`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *AssociationDefinition) GetCardinalityOk() (*string, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *AssociationDefinition) SetCardinality(v string)`

SetCardinality sets Cardinality field to given value.


### GetInverseName

`func (o *AssociationDefinition) GetInverseName() string`

GetInverseName returns the InverseName field if non-nil, zero value otherwise.

### GetInverseNameOk

`func (o *AssociationDefinition) GetInverseNameOk() (*string, bool)`

GetInverseNameOk returns a tuple with the InverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseName

`func (o *AssociationDefinition) SetInverseName(v string)`

SetInverseName sets InverseName field to given value.


### GetHasAllAssociatedObjects

`func (o *AssociationDefinition) GetHasAllAssociatedObjects() bool`

GetHasAllAssociatedObjects returns the HasAllAssociatedObjects field if non-nil, zero value otherwise.

### GetHasAllAssociatedObjectsOk

`func (o *AssociationDefinition) GetHasAllAssociatedObjectsOk() (*bool, bool)`

GetHasAllAssociatedObjectsOk returns a tuple with the HasAllAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAllAssociatedObjects

`func (o *AssociationDefinition) SetHasAllAssociatedObjects(v bool)`

SetHasAllAssociatedObjects sets HasAllAssociatedObjects field to given value.


### GetInverseHasAllAssociatedObjects

`func (o *AssociationDefinition) GetInverseHasAllAssociatedObjects() bool`

GetInverseHasAllAssociatedObjects returns the InverseHasAllAssociatedObjects field if non-nil, zero value otherwise.

### GetInverseHasAllAssociatedObjectsOk

`func (o *AssociationDefinition) GetInverseHasAllAssociatedObjectsOk() (*bool, bool)`

GetInverseHasAllAssociatedObjectsOk returns a tuple with the InverseHasAllAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseHasAllAssociatedObjects

`func (o *AssociationDefinition) SetInverseHasAllAssociatedObjects(v bool)`

SetInverseHasAllAssociatedObjects sets InverseHasAllAssociatedObjects field to given value.


### GetFromObjectTypeId

`func (o *AssociationDefinition) GetFromObjectTypeId() string`

GetFromObjectTypeId returns the FromObjectTypeId field if non-nil, zero value otherwise.

### GetFromObjectTypeIdOk

`func (o *AssociationDefinition) GetFromObjectTypeIdOk() (*string, bool)`

GetFromObjectTypeIdOk returns a tuple with the FromObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromObjectTypeId

`func (o *AssociationDefinition) SetFromObjectTypeId(v string)`

SetFromObjectTypeId sets FromObjectTypeId field to given value.


### GetPortalUniqueIdentifier

`func (o *AssociationDefinition) GetPortalUniqueIdentifier() string`

GetPortalUniqueIdentifier returns the PortalUniqueIdentifier field if non-nil, zero value otherwise.

### GetPortalUniqueIdentifierOk

`func (o *AssociationDefinition) GetPortalUniqueIdentifierOk() (*string, bool)`

GetPortalUniqueIdentifierOk returns a tuple with the PortalUniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUniqueIdentifier

`func (o *AssociationDefinition) SetPortalUniqueIdentifier(v string)`

SetPortalUniqueIdentifier sets PortalUniqueIdentifier field to given value.


### GetIsInversePrimary

`func (o *AssociationDefinition) GetIsInversePrimary() bool`

GetIsInversePrimary returns the IsInversePrimary field if non-nil, zero value otherwise.

### GetIsInversePrimaryOk

`func (o *AssociationDefinition) GetIsInversePrimaryOk() (*bool, bool)`

GetIsInversePrimaryOk returns a tuple with the IsInversePrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInversePrimary

`func (o *AssociationDefinition) SetIsInversePrimary(v bool)`

SetIsInversePrimary sets IsInversePrimary field to given value.


### GetName

`func (o *AssociationDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssociationDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssociationDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *AssociationDefinition) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AssociationDefinition) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AssociationDefinition) SetCategory(v string)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


