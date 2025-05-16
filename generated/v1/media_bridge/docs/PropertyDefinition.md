# PropertyDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulcrumPortalId** | Pointer to **int32** |  | [optional] 
**ObjectTypeId** | **string** |  | 
**JanusGroup** | Pointer to **string** |  | [optional] 
**FulcrumTimestamp** | Pointer to **int64** |  | [optional] 
**Property** | [**Property**](Property.md) |  | 
**CalculationFormula** | Pointer to **string** |  | [optional] 
**PropertyRequirements** | Pointer to [**DefaultRequirements**](DefaultRequirements.md) |  | [optional] 
**DefinitionSource** | Pointer to [**PropertyDefinitionSource**](PropertyDefinitionSource.md) |  | [optional] 
**ExternalOptionsMetaData** | Pointer to [**ExternalOptionsMetaData**](ExternalOptionsMetaData.md) |  | [optional] 
**Permission** | Pointer to [**FieldLevelPermission**](FieldLevelPermission.md) |  | [optional] 
**CalculationExpression** | Pointer to [**OrInputsInner**](OrInputsInner.md) |  | [optional] 
**RollupExpression** | Pointer to [**RollupExpression**](RollupExpression.md) |  | [optional] 

## Methods

### NewPropertyDefinition

`func NewPropertyDefinition(objectTypeId string, property Property, ) *PropertyDefinition`

NewPropertyDefinition instantiates a new PropertyDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyDefinitionWithDefaults

`func NewPropertyDefinitionWithDefaults() *PropertyDefinition`

NewPropertyDefinitionWithDefaults instantiates a new PropertyDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulcrumPortalId

`func (o *PropertyDefinition) GetFulcrumPortalId() int32`

GetFulcrumPortalId returns the FulcrumPortalId field if non-nil, zero value otherwise.

### GetFulcrumPortalIdOk

`func (o *PropertyDefinition) GetFulcrumPortalIdOk() (*int32, bool)`

GetFulcrumPortalIdOk returns a tuple with the FulcrumPortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulcrumPortalId

`func (o *PropertyDefinition) SetFulcrumPortalId(v int32)`

SetFulcrumPortalId sets FulcrumPortalId field to given value.

### HasFulcrumPortalId

`func (o *PropertyDefinition) HasFulcrumPortalId() bool`

HasFulcrumPortalId returns a boolean if a field has been set.

### GetObjectTypeId

`func (o *PropertyDefinition) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *PropertyDefinition) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *PropertyDefinition) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetJanusGroup

`func (o *PropertyDefinition) GetJanusGroup() string`

GetJanusGroup returns the JanusGroup field if non-nil, zero value otherwise.

### GetJanusGroupOk

`func (o *PropertyDefinition) GetJanusGroupOk() (*string, bool)`

GetJanusGroupOk returns a tuple with the JanusGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJanusGroup

`func (o *PropertyDefinition) SetJanusGroup(v string)`

SetJanusGroup sets JanusGroup field to given value.

### HasJanusGroup

`func (o *PropertyDefinition) HasJanusGroup() bool`

HasJanusGroup returns a boolean if a field has been set.

### GetFulcrumTimestamp

`func (o *PropertyDefinition) GetFulcrumTimestamp() int64`

GetFulcrumTimestamp returns the FulcrumTimestamp field if non-nil, zero value otherwise.

### GetFulcrumTimestampOk

`func (o *PropertyDefinition) GetFulcrumTimestampOk() (*int64, bool)`

GetFulcrumTimestampOk returns a tuple with the FulcrumTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulcrumTimestamp

`func (o *PropertyDefinition) SetFulcrumTimestamp(v int64)`

SetFulcrumTimestamp sets FulcrumTimestamp field to given value.

### HasFulcrumTimestamp

`func (o *PropertyDefinition) HasFulcrumTimestamp() bool`

HasFulcrumTimestamp returns a boolean if a field has been set.

### GetProperty

`func (o *PropertyDefinition) GetProperty() Property`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PropertyDefinition) GetPropertyOk() (*Property, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PropertyDefinition) SetProperty(v Property)`

SetProperty sets Property field to given value.


### GetCalculationFormula

`func (o *PropertyDefinition) GetCalculationFormula() string`

GetCalculationFormula returns the CalculationFormula field if non-nil, zero value otherwise.

### GetCalculationFormulaOk

`func (o *PropertyDefinition) GetCalculationFormulaOk() (*string, bool)`

GetCalculationFormulaOk returns a tuple with the CalculationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationFormula

`func (o *PropertyDefinition) SetCalculationFormula(v string)`

SetCalculationFormula sets CalculationFormula field to given value.

### HasCalculationFormula

`func (o *PropertyDefinition) HasCalculationFormula() bool`

HasCalculationFormula returns a boolean if a field has been set.

### GetPropertyRequirements

`func (o *PropertyDefinition) GetPropertyRequirements() DefaultRequirements`

GetPropertyRequirements returns the PropertyRequirements field if non-nil, zero value otherwise.

### GetPropertyRequirementsOk

`func (o *PropertyDefinition) GetPropertyRequirementsOk() (*DefaultRequirements, bool)`

GetPropertyRequirementsOk returns a tuple with the PropertyRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRequirements

`func (o *PropertyDefinition) SetPropertyRequirements(v DefaultRequirements)`

SetPropertyRequirements sets PropertyRequirements field to given value.

### HasPropertyRequirements

`func (o *PropertyDefinition) HasPropertyRequirements() bool`

HasPropertyRequirements returns a boolean if a field has been set.

### GetDefinitionSource

`func (o *PropertyDefinition) GetDefinitionSource() PropertyDefinitionSource`

GetDefinitionSource returns the DefinitionSource field if non-nil, zero value otherwise.

### GetDefinitionSourceOk

`func (o *PropertyDefinition) GetDefinitionSourceOk() (*PropertyDefinitionSource, bool)`

GetDefinitionSourceOk returns a tuple with the DefinitionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionSource

`func (o *PropertyDefinition) SetDefinitionSource(v PropertyDefinitionSource)`

SetDefinitionSource sets DefinitionSource field to given value.

### HasDefinitionSource

`func (o *PropertyDefinition) HasDefinitionSource() bool`

HasDefinitionSource returns a boolean if a field has been set.

### GetExternalOptionsMetaData

`func (o *PropertyDefinition) GetExternalOptionsMetaData() ExternalOptionsMetaData`

GetExternalOptionsMetaData returns the ExternalOptionsMetaData field if non-nil, zero value otherwise.

### GetExternalOptionsMetaDataOk

`func (o *PropertyDefinition) GetExternalOptionsMetaDataOk() (*ExternalOptionsMetaData, bool)`

GetExternalOptionsMetaDataOk returns a tuple with the ExternalOptionsMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOptionsMetaData

`func (o *PropertyDefinition) SetExternalOptionsMetaData(v ExternalOptionsMetaData)`

SetExternalOptionsMetaData sets ExternalOptionsMetaData field to given value.

### HasExternalOptionsMetaData

`func (o *PropertyDefinition) HasExternalOptionsMetaData() bool`

HasExternalOptionsMetaData returns a boolean if a field has been set.

### GetPermission

`func (o *PropertyDefinition) GetPermission() FieldLevelPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PropertyDefinition) GetPermissionOk() (*FieldLevelPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PropertyDefinition) SetPermission(v FieldLevelPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *PropertyDefinition) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetCalculationExpression

`func (o *PropertyDefinition) GetCalculationExpression() OrInputsInner`

GetCalculationExpression returns the CalculationExpression field if non-nil, zero value otherwise.

### GetCalculationExpressionOk

`func (o *PropertyDefinition) GetCalculationExpressionOk() (*OrInputsInner, bool)`

GetCalculationExpressionOk returns a tuple with the CalculationExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationExpression

`func (o *PropertyDefinition) SetCalculationExpression(v OrInputsInner)`

SetCalculationExpression sets CalculationExpression field to given value.

### HasCalculationExpression

`func (o *PropertyDefinition) HasCalculationExpression() bool`

HasCalculationExpression returns a boolean if a field has been set.

### GetRollupExpression

`func (o *PropertyDefinition) GetRollupExpression() RollupExpression`

GetRollupExpression returns the RollupExpression field if non-nil, zero value otherwise.

### GetRollupExpressionOk

`func (o *PropertyDefinition) GetRollupExpressionOk() (*RollupExpression, bool)`

GetRollupExpressionOk returns a tuple with the RollupExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupExpression

`func (o *PropertyDefinition) SetRollupExpression(v RollupExpression)`

SetRollupExpression sets RollupExpression field to given value.

### HasRollupExpression

`func (o *PropertyDefinition) HasRollupExpression() bool`

HasRollupExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


