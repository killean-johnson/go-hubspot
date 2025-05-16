# InboundDbObjectType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasOwners** | **bool** |  | 
**PermissioningType** | **string** |  | 
**PrimaryDisplayLabelPropertyName** | Pointer to **string** |  | [optional] 
**ObjectTypeId** | **string** |  | 
**LastModifiedPropertyName** | **string** |  | 
**IntegrationAppId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PipelinePropertyName** | **string** |  | 
**PipelineTimeToClosePropertyName** | Pointer to **string** |  | [optional] 
**SingularForm** | Pointer to **string** |  | [optional] 
**HasPipelines** | **bool** |  | 
**HasExternalObjectIds** | **bool** |  | 
**ScopeMappings** | [**[]ScopeMapping**](ScopeMapping.md) |  | 
**HasDefaultProperties** | **bool** |  | 
**IndexedForFiltersAndReports** | **bool** |  | 
**CreateDatePropertyName** | **string** |  | 
**Id** | **int32** |  | 
**OwnerPortalId** | Pointer to **int32** |  | [optional] 
**Restorable** | **bool** |  | 
**PluralForm** | Pointer to **string** |  | [optional] 
**MetaTypeId** | **int32** |  | 
**DefaultSearchPropertyNames** | **[]string** |  | 
**JanusGroup** | Pointer to **string** |  | [optional] 
**PipelineCloseDatePropertyName** | Pointer to **string** |  | [optional] 
**FullyQualifiedName** | **string** |  | 
**Deleted** | **bool** |  | 
**RequiredProperties** | **[]string** |  | 
**WriteScopeName** | Pointer to **string** |  | [optional] 
**AccessScopeName** | Pointer to **string** |  | [optional] 
**HasCustomProperties** | **bool** |  | 
**Name** | **string** |  | 
**MetaType** | **string** |  | 
**PipelineStagePropertyName** | **string** |  | 
**ReadScopeName** | Pointer to **string** |  | [optional] 
**SecondaryDisplayLabelPropertyNames** | **[]string** |  | 

## Methods

### NewInboundDbObjectType

`func NewInboundDbObjectType(hasOwners bool, permissioningType string, objectTypeId string, lastModifiedPropertyName string, pipelinePropertyName string, hasPipelines bool, hasExternalObjectIds bool, scopeMappings []ScopeMapping, hasDefaultProperties bool, indexedForFiltersAndReports bool, createDatePropertyName string, id int32, restorable bool, metaTypeId int32, defaultSearchPropertyNames []string, fullyQualifiedName string, deleted bool, requiredProperties []string, hasCustomProperties bool, name string, metaType string, pipelineStagePropertyName string, secondaryDisplayLabelPropertyNames []string, ) *InboundDbObjectType`

NewInboundDbObjectType instantiates a new InboundDbObjectType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundDbObjectTypeWithDefaults

`func NewInboundDbObjectTypeWithDefaults() *InboundDbObjectType`

NewInboundDbObjectTypeWithDefaults instantiates a new InboundDbObjectType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasOwners

`func (o *InboundDbObjectType) GetHasOwners() bool`

GetHasOwners returns the HasOwners field if non-nil, zero value otherwise.

### GetHasOwnersOk

`func (o *InboundDbObjectType) GetHasOwnersOk() (*bool, bool)`

GetHasOwnersOk returns a tuple with the HasOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOwners

`func (o *InboundDbObjectType) SetHasOwners(v bool)`

SetHasOwners sets HasOwners field to given value.


### GetPermissioningType

`func (o *InboundDbObjectType) GetPermissioningType() string`

GetPermissioningType returns the PermissioningType field if non-nil, zero value otherwise.

### GetPermissioningTypeOk

`func (o *InboundDbObjectType) GetPermissioningTypeOk() (*string, bool)`

GetPermissioningTypeOk returns a tuple with the PermissioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissioningType

`func (o *InboundDbObjectType) SetPermissioningType(v string)`

SetPermissioningType sets PermissioningType field to given value.


### GetPrimaryDisplayLabelPropertyName

`func (o *InboundDbObjectType) GetPrimaryDisplayLabelPropertyName() string`

GetPrimaryDisplayLabelPropertyName returns the PrimaryDisplayLabelPropertyName field if non-nil, zero value otherwise.

### GetPrimaryDisplayLabelPropertyNameOk

`func (o *InboundDbObjectType) GetPrimaryDisplayLabelPropertyNameOk() (*string, bool)`

GetPrimaryDisplayLabelPropertyNameOk returns a tuple with the PrimaryDisplayLabelPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDisplayLabelPropertyName

`func (o *InboundDbObjectType) SetPrimaryDisplayLabelPropertyName(v string)`

SetPrimaryDisplayLabelPropertyName sets PrimaryDisplayLabelPropertyName field to given value.

### HasPrimaryDisplayLabelPropertyName

`func (o *InboundDbObjectType) HasPrimaryDisplayLabelPropertyName() bool`

HasPrimaryDisplayLabelPropertyName returns a boolean if a field has been set.

### GetObjectTypeId

`func (o *InboundDbObjectType) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *InboundDbObjectType) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *InboundDbObjectType) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetLastModifiedPropertyName

`func (o *InboundDbObjectType) GetLastModifiedPropertyName() string`

GetLastModifiedPropertyName returns the LastModifiedPropertyName field if non-nil, zero value otherwise.

### GetLastModifiedPropertyNameOk

`func (o *InboundDbObjectType) GetLastModifiedPropertyNameOk() (*string, bool)`

GetLastModifiedPropertyNameOk returns a tuple with the LastModifiedPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedPropertyName

`func (o *InboundDbObjectType) SetLastModifiedPropertyName(v string)`

SetLastModifiedPropertyName sets LastModifiedPropertyName field to given value.


### GetIntegrationAppId

`func (o *InboundDbObjectType) GetIntegrationAppId() int32`

GetIntegrationAppId returns the IntegrationAppId field if non-nil, zero value otherwise.

### GetIntegrationAppIdOk

`func (o *InboundDbObjectType) GetIntegrationAppIdOk() (*int32, bool)`

GetIntegrationAppIdOk returns a tuple with the IntegrationAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationAppId

`func (o *InboundDbObjectType) SetIntegrationAppId(v int32)`

SetIntegrationAppId sets IntegrationAppId field to given value.

### HasIntegrationAppId

`func (o *InboundDbObjectType) HasIntegrationAppId() bool`

HasIntegrationAppId returns a boolean if a field has been set.

### GetDescription

`func (o *InboundDbObjectType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InboundDbObjectType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InboundDbObjectType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InboundDbObjectType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPipelinePropertyName

`func (o *InboundDbObjectType) GetPipelinePropertyName() string`

GetPipelinePropertyName returns the PipelinePropertyName field if non-nil, zero value otherwise.

### GetPipelinePropertyNameOk

`func (o *InboundDbObjectType) GetPipelinePropertyNameOk() (*string, bool)`

GetPipelinePropertyNameOk returns a tuple with the PipelinePropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelinePropertyName

`func (o *InboundDbObjectType) SetPipelinePropertyName(v string)`

SetPipelinePropertyName sets PipelinePropertyName field to given value.


### GetPipelineTimeToClosePropertyName

`func (o *InboundDbObjectType) GetPipelineTimeToClosePropertyName() string`

GetPipelineTimeToClosePropertyName returns the PipelineTimeToClosePropertyName field if non-nil, zero value otherwise.

### GetPipelineTimeToClosePropertyNameOk

`func (o *InboundDbObjectType) GetPipelineTimeToClosePropertyNameOk() (*string, bool)`

GetPipelineTimeToClosePropertyNameOk returns a tuple with the PipelineTimeToClosePropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineTimeToClosePropertyName

`func (o *InboundDbObjectType) SetPipelineTimeToClosePropertyName(v string)`

SetPipelineTimeToClosePropertyName sets PipelineTimeToClosePropertyName field to given value.

### HasPipelineTimeToClosePropertyName

`func (o *InboundDbObjectType) HasPipelineTimeToClosePropertyName() bool`

HasPipelineTimeToClosePropertyName returns a boolean if a field has been set.

### GetSingularForm

`func (o *InboundDbObjectType) GetSingularForm() string`

GetSingularForm returns the SingularForm field if non-nil, zero value otherwise.

### GetSingularFormOk

`func (o *InboundDbObjectType) GetSingularFormOk() (*string, bool)`

GetSingularFormOk returns a tuple with the SingularForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingularForm

`func (o *InboundDbObjectType) SetSingularForm(v string)`

SetSingularForm sets SingularForm field to given value.

### HasSingularForm

`func (o *InboundDbObjectType) HasSingularForm() bool`

HasSingularForm returns a boolean if a field has been set.

### GetHasPipelines

`func (o *InboundDbObjectType) GetHasPipelines() bool`

GetHasPipelines returns the HasPipelines field if non-nil, zero value otherwise.

### GetHasPipelinesOk

`func (o *InboundDbObjectType) GetHasPipelinesOk() (*bool, bool)`

GetHasPipelinesOk returns a tuple with the HasPipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPipelines

`func (o *InboundDbObjectType) SetHasPipelines(v bool)`

SetHasPipelines sets HasPipelines field to given value.


### GetHasExternalObjectIds

`func (o *InboundDbObjectType) GetHasExternalObjectIds() bool`

GetHasExternalObjectIds returns the HasExternalObjectIds field if non-nil, zero value otherwise.

### GetHasExternalObjectIdsOk

`func (o *InboundDbObjectType) GetHasExternalObjectIdsOk() (*bool, bool)`

GetHasExternalObjectIdsOk returns a tuple with the HasExternalObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExternalObjectIds

`func (o *InboundDbObjectType) SetHasExternalObjectIds(v bool)`

SetHasExternalObjectIds sets HasExternalObjectIds field to given value.


### GetScopeMappings

`func (o *InboundDbObjectType) GetScopeMappings() []ScopeMapping`

GetScopeMappings returns the ScopeMappings field if non-nil, zero value otherwise.

### GetScopeMappingsOk

`func (o *InboundDbObjectType) GetScopeMappingsOk() (*[]ScopeMapping, bool)`

GetScopeMappingsOk returns a tuple with the ScopeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeMappings

`func (o *InboundDbObjectType) SetScopeMappings(v []ScopeMapping)`

SetScopeMappings sets ScopeMappings field to given value.


### GetHasDefaultProperties

`func (o *InboundDbObjectType) GetHasDefaultProperties() bool`

GetHasDefaultProperties returns the HasDefaultProperties field if non-nil, zero value otherwise.

### GetHasDefaultPropertiesOk

`func (o *InboundDbObjectType) GetHasDefaultPropertiesOk() (*bool, bool)`

GetHasDefaultPropertiesOk returns a tuple with the HasDefaultProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefaultProperties

`func (o *InboundDbObjectType) SetHasDefaultProperties(v bool)`

SetHasDefaultProperties sets HasDefaultProperties field to given value.


### GetIndexedForFiltersAndReports

`func (o *InboundDbObjectType) GetIndexedForFiltersAndReports() bool`

GetIndexedForFiltersAndReports returns the IndexedForFiltersAndReports field if non-nil, zero value otherwise.

### GetIndexedForFiltersAndReportsOk

`func (o *InboundDbObjectType) GetIndexedForFiltersAndReportsOk() (*bool, bool)`

GetIndexedForFiltersAndReportsOk returns a tuple with the IndexedForFiltersAndReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedForFiltersAndReports

`func (o *InboundDbObjectType) SetIndexedForFiltersAndReports(v bool)`

SetIndexedForFiltersAndReports sets IndexedForFiltersAndReports field to given value.


### GetCreateDatePropertyName

`func (o *InboundDbObjectType) GetCreateDatePropertyName() string`

GetCreateDatePropertyName returns the CreateDatePropertyName field if non-nil, zero value otherwise.

### GetCreateDatePropertyNameOk

`func (o *InboundDbObjectType) GetCreateDatePropertyNameOk() (*string, bool)`

GetCreateDatePropertyNameOk returns a tuple with the CreateDatePropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDatePropertyName

`func (o *InboundDbObjectType) SetCreateDatePropertyName(v string)`

SetCreateDatePropertyName sets CreateDatePropertyName field to given value.


### GetId

`func (o *InboundDbObjectType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboundDbObjectType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboundDbObjectType) SetId(v int32)`

SetId sets Id field to given value.


### GetOwnerPortalId

`func (o *InboundDbObjectType) GetOwnerPortalId() int32`

GetOwnerPortalId returns the OwnerPortalId field if non-nil, zero value otherwise.

### GetOwnerPortalIdOk

`func (o *InboundDbObjectType) GetOwnerPortalIdOk() (*int32, bool)`

GetOwnerPortalIdOk returns a tuple with the OwnerPortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPortalId

`func (o *InboundDbObjectType) SetOwnerPortalId(v int32)`

SetOwnerPortalId sets OwnerPortalId field to given value.

### HasOwnerPortalId

`func (o *InboundDbObjectType) HasOwnerPortalId() bool`

HasOwnerPortalId returns a boolean if a field has been set.

### GetRestorable

`func (o *InboundDbObjectType) GetRestorable() bool`

GetRestorable returns the Restorable field if non-nil, zero value otherwise.

### GetRestorableOk

`func (o *InboundDbObjectType) GetRestorableOk() (*bool, bool)`

GetRestorableOk returns a tuple with the Restorable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorable

`func (o *InboundDbObjectType) SetRestorable(v bool)`

SetRestorable sets Restorable field to given value.


### GetPluralForm

`func (o *InboundDbObjectType) GetPluralForm() string`

GetPluralForm returns the PluralForm field if non-nil, zero value otherwise.

### GetPluralFormOk

`func (o *InboundDbObjectType) GetPluralFormOk() (*string, bool)`

GetPluralFormOk returns a tuple with the PluralForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluralForm

`func (o *InboundDbObjectType) SetPluralForm(v string)`

SetPluralForm sets PluralForm field to given value.

### HasPluralForm

`func (o *InboundDbObjectType) HasPluralForm() bool`

HasPluralForm returns a boolean if a field has been set.

### GetMetaTypeId

`func (o *InboundDbObjectType) GetMetaTypeId() int32`

GetMetaTypeId returns the MetaTypeId field if non-nil, zero value otherwise.

### GetMetaTypeIdOk

`func (o *InboundDbObjectType) GetMetaTypeIdOk() (*int32, bool)`

GetMetaTypeIdOk returns a tuple with the MetaTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTypeId

`func (o *InboundDbObjectType) SetMetaTypeId(v int32)`

SetMetaTypeId sets MetaTypeId field to given value.


### GetDefaultSearchPropertyNames

`func (o *InboundDbObjectType) GetDefaultSearchPropertyNames() []string`

GetDefaultSearchPropertyNames returns the DefaultSearchPropertyNames field if non-nil, zero value otherwise.

### GetDefaultSearchPropertyNamesOk

`func (o *InboundDbObjectType) GetDefaultSearchPropertyNamesOk() (*[]string, bool)`

GetDefaultSearchPropertyNamesOk returns a tuple with the DefaultSearchPropertyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSearchPropertyNames

`func (o *InboundDbObjectType) SetDefaultSearchPropertyNames(v []string)`

SetDefaultSearchPropertyNames sets DefaultSearchPropertyNames field to given value.


### GetJanusGroup

`func (o *InboundDbObjectType) GetJanusGroup() string`

GetJanusGroup returns the JanusGroup field if non-nil, zero value otherwise.

### GetJanusGroupOk

`func (o *InboundDbObjectType) GetJanusGroupOk() (*string, bool)`

GetJanusGroupOk returns a tuple with the JanusGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJanusGroup

`func (o *InboundDbObjectType) SetJanusGroup(v string)`

SetJanusGroup sets JanusGroup field to given value.

### HasJanusGroup

`func (o *InboundDbObjectType) HasJanusGroup() bool`

HasJanusGroup returns a boolean if a field has been set.

### GetPipelineCloseDatePropertyName

`func (o *InboundDbObjectType) GetPipelineCloseDatePropertyName() string`

GetPipelineCloseDatePropertyName returns the PipelineCloseDatePropertyName field if non-nil, zero value otherwise.

### GetPipelineCloseDatePropertyNameOk

`func (o *InboundDbObjectType) GetPipelineCloseDatePropertyNameOk() (*string, bool)`

GetPipelineCloseDatePropertyNameOk returns a tuple with the PipelineCloseDatePropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineCloseDatePropertyName

`func (o *InboundDbObjectType) SetPipelineCloseDatePropertyName(v string)`

SetPipelineCloseDatePropertyName sets PipelineCloseDatePropertyName field to given value.

### HasPipelineCloseDatePropertyName

`func (o *InboundDbObjectType) HasPipelineCloseDatePropertyName() bool`

HasPipelineCloseDatePropertyName returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *InboundDbObjectType) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *InboundDbObjectType) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *InboundDbObjectType) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.


### GetDeleted

`func (o *InboundDbObjectType) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *InboundDbObjectType) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *InboundDbObjectType) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetRequiredProperties

`func (o *InboundDbObjectType) GetRequiredProperties() []string`

GetRequiredProperties returns the RequiredProperties field if non-nil, zero value otherwise.

### GetRequiredPropertiesOk

`func (o *InboundDbObjectType) GetRequiredPropertiesOk() (*[]string, bool)`

GetRequiredPropertiesOk returns a tuple with the RequiredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredProperties

`func (o *InboundDbObjectType) SetRequiredProperties(v []string)`

SetRequiredProperties sets RequiredProperties field to given value.


### GetWriteScopeName

`func (o *InboundDbObjectType) GetWriteScopeName() string`

GetWriteScopeName returns the WriteScopeName field if non-nil, zero value otherwise.

### GetWriteScopeNameOk

`func (o *InboundDbObjectType) GetWriteScopeNameOk() (*string, bool)`

GetWriteScopeNameOk returns a tuple with the WriteScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteScopeName

`func (o *InboundDbObjectType) SetWriteScopeName(v string)`

SetWriteScopeName sets WriteScopeName field to given value.

### HasWriteScopeName

`func (o *InboundDbObjectType) HasWriteScopeName() bool`

HasWriteScopeName returns a boolean if a field has been set.

### GetAccessScopeName

`func (o *InboundDbObjectType) GetAccessScopeName() string`

GetAccessScopeName returns the AccessScopeName field if non-nil, zero value otherwise.

### GetAccessScopeNameOk

`func (o *InboundDbObjectType) GetAccessScopeNameOk() (*string, bool)`

GetAccessScopeNameOk returns a tuple with the AccessScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScopeName

`func (o *InboundDbObjectType) SetAccessScopeName(v string)`

SetAccessScopeName sets AccessScopeName field to given value.

### HasAccessScopeName

`func (o *InboundDbObjectType) HasAccessScopeName() bool`

HasAccessScopeName returns a boolean if a field has been set.

### GetHasCustomProperties

`func (o *InboundDbObjectType) GetHasCustomProperties() bool`

GetHasCustomProperties returns the HasCustomProperties field if non-nil, zero value otherwise.

### GetHasCustomPropertiesOk

`func (o *InboundDbObjectType) GetHasCustomPropertiesOk() (*bool, bool)`

GetHasCustomPropertiesOk returns a tuple with the HasCustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomProperties

`func (o *InboundDbObjectType) SetHasCustomProperties(v bool)`

SetHasCustomProperties sets HasCustomProperties field to given value.


### GetName

`func (o *InboundDbObjectType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InboundDbObjectType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InboundDbObjectType) SetName(v string)`

SetName sets Name field to given value.


### GetMetaType

`func (o *InboundDbObjectType) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *InboundDbObjectType) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *InboundDbObjectType) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetPipelineStagePropertyName

`func (o *InboundDbObjectType) GetPipelineStagePropertyName() string`

GetPipelineStagePropertyName returns the PipelineStagePropertyName field if non-nil, zero value otherwise.

### GetPipelineStagePropertyNameOk

`func (o *InboundDbObjectType) GetPipelineStagePropertyNameOk() (*string, bool)`

GetPipelineStagePropertyNameOk returns a tuple with the PipelineStagePropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineStagePropertyName

`func (o *InboundDbObjectType) SetPipelineStagePropertyName(v string)`

SetPipelineStagePropertyName sets PipelineStagePropertyName field to given value.


### GetReadScopeName

`func (o *InboundDbObjectType) GetReadScopeName() string`

GetReadScopeName returns the ReadScopeName field if non-nil, zero value otherwise.

### GetReadScopeNameOk

`func (o *InboundDbObjectType) GetReadScopeNameOk() (*string, bool)`

GetReadScopeNameOk returns a tuple with the ReadScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadScopeName

`func (o *InboundDbObjectType) SetReadScopeName(v string)`

SetReadScopeName sets ReadScopeName field to given value.

### HasReadScopeName

`func (o *InboundDbObjectType) HasReadScopeName() bool`

HasReadScopeName returns a boolean if a field has been set.

### GetSecondaryDisplayLabelPropertyNames

`func (o *InboundDbObjectType) GetSecondaryDisplayLabelPropertyNames() []string`

GetSecondaryDisplayLabelPropertyNames returns the SecondaryDisplayLabelPropertyNames field if non-nil, zero value otherwise.

### GetSecondaryDisplayLabelPropertyNamesOk

`func (o *InboundDbObjectType) GetSecondaryDisplayLabelPropertyNamesOk() (*[]string, bool)`

GetSecondaryDisplayLabelPropertyNamesOk returns a tuple with the SecondaryDisplayLabelPropertyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDisplayLabelPropertyNames

`func (o *InboundDbObjectType) SetSecondaryDisplayLabelPropertyNames(v []string)`

SetSecondaryDisplayLabelPropertyNames sets SecondaryDisplayLabelPropertyNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


