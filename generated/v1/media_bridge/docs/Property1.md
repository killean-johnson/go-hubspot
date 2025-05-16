# Property1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedUserId** | Pointer to **string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**ModificationMetadata** | Pointer to [**PropertyModificationMetadata**](PropertyModificationMetadata.md) |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  | 
**ShowCurrencySymbol** | Pointer to **bool** |  | [optional] 
**Label** | **string** |  | 
**Type** | **string** |  | 
**HubspotDefined** | Pointer to **bool** |  | [optional] 
**FormField** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**GroupName** | **string** |  | 
**ReferencedObjectType** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Options** | [**[]Option1**](Option1.md) |  | 
**CalculationFormula** | Pointer to **string** |  | [optional] 
**HasUniqueValue** | Pointer to **bool** |  | [optional] 
**FieldType** | **string** |  | 
**UpdatedUserId** | Pointer to **string** |  | [optional] 
**Calculated** | Pointer to **bool** |  | [optional] 
**ExternalOptions** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProperty1

`func NewProperty1(description string, label string, type_ string, groupName string, name string, options []Option1, fieldType string, ) *Property1`

NewProperty1 instantiates a new Property1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProperty1WithDefaults

`func NewProperty1WithDefaults() *Property1`

NewProperty1WithDefaults instantiates a new Property1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedUserId

`func (o *Property1) GetCreatedUserId() string`

GetCreatedUserId returns the CreatedUserId field if non-nil, zero value otherwise.

### GetCreatedUserIdOk

`func (o *Property1) GetCreatedUserIdOk() (*string, bool)`

GetCreatedUserIdOk returns a tuple with the CreatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUserId

`func (o *Property1) SetCreatedUserId(v string)`

SetCreatedUserId sets CreatedUserId field to given value.

### HasCreatedUserId

`func (o *Property1) HasCreatedUserId() bool`

HasCreatedUserId returns a boolean if a field has been set.

### GetHidden

`func (o *Property1) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Property1) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Property1) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Property1) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetModificationMetadata

`func (o *Property1) GetModificationMetadata() PropertyModificationMetadata`

GetModificationMetadata returns the ModificationMetadata field if non-nil, zero value otherwise.

### GetModificationMetadataOk

`func (o *Property1) GetModificationMetadataOk() (*PropertyModificationMetadata, bool)`

GetModificationMetadataOk returns a tuple with the ModificationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationMetadata

`func (o *Property1) SetModificationMetadata(v PropertyModificationMetadata)`

SetModificationMetadata sets ModificationMetadata field to given value.

### HasModificationMetadata

`func (o *Property1) HasModificationMetadata() bool`

HasModificationMetadata returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *Property1) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *Property1) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *Property1) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *Property1) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetDescription

`func (o *Property1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Property1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Property1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetShowCurrencySymbol

`func (o *Property1) GetShowCurrencySymbol() bool`

GetShowCurrencySymbol returns the ShowCurrencySymbol field if non-nil, zero value otherwise.

### GetShowCurrencySymbolOk

`func (o *Property1) GetShowCurrencySymbolOk() (*bool, bool)`

GetShowCurrencySymbolOk returns a tuple with the ShowCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCurrencySymbol

`func (o *Property1) SetShowCurrencySymbol(v bool)`

SetShowCurrencySymbol sets ShowCurrencySymbol field to given value.

### HasShowCurrencySymbol

`func (o *Property1) HasShowCurrencySymbol() bool`

HasShowCurrencySymbol returns a boolean if a field has been set.

### GetLabel

`func (o *Property1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Property1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Property1) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *Property1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Property1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Property1) SetType(v string)`

SetType sets Type field to given value.


### GetHubspotDefined

`func (o *Property1) GetHubspotDefined() bool`

GetHubspotDefined returns the HubspotDefined field if non-nil, zero value otherwise.

### GetHubspotDefinedOk

`func (o *Property1) GetHubspotDefinedOk() (*bool, bool)`

GetHubspotDefinedOk returns a tuple with the HubspotDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubspotDefined

`func (o *Property1) SetHubspotDefined(v bool)`

SetHubspotDefined sets HubspotDefined field to given value.

### HasHubspotDefined

`func (o *Property1) HasHubspotDefined() bool`

HasHubspotDefined returns a boolean if a field has been set.

### GetFormField

`func (o *Property1) GetFormField() bool`

GetFormField returns the FormField field if non-nil, zero value otherwise.

### GetFormFieldOk

`func (o *Property1) GetFormFieldOk() (*bool, bool)`

GetFormFieldOk returns a tuple with the FormField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormField

`func (o *Property1) SetFormField(v bool)`

SetFormField sets FormField field to given value.

### HasFormField

`func (o *Property1) HasFormField() bool`

HasFormField returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Property1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Property1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Property1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Property1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Property1) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Property1) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Property1) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Property1) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetArchived

`func (o *Property1) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Property1) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Property1) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Property1) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetGroupName

`func (o *Property1) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *Property1) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *Property1) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetReferencedObjectType

`func (o *Property1) GetReferencedObjectType() string`

GetReferencedObjectType returns the ReferencedObjectType field if non-nil, zero value otherwise.

### GetReferencedObjectTypeOk

`func (o *Property1) GetReferencedObjectTypeOk() (*string, bool)`

GetReferencedObjectTypeOk returns a tuple with the ReferencedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedObjectType

`func (o *Property1) SetReferencedObjectType(v string)`

SetReferencedObjectType sets ReferencedObjectType field to given value.

### HasReferencedObjectType

`func (o *Property1) HasReferencedObjectType() bool`

HasReferencedObjectType returns a boolean if a field has been set.

### GetName

`func (o *Property1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Property1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Property1) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *Property1) GetOptions() []Option1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Property1) GetOptionsOk() (*[]Option1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Property1) SetOptions(v []Option1)`

SetOptions sets Options field to given value.


### GetCalculationFormula

`func (o *Property1) GetCalculationFormula() string`

GetCalculationFormula returns the CalculationFormula field if non-nil, zero value otherwise.

### GetCalculationFormulaOk

`func (o *Property1) GetCalculationFormulaOk() (*string, bool)`

GetCalculationFormulaOk returns a tuple with the CalculationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationFormula

`func (o *Property1) SetCalculationFormula(v string)`

SetCalculationFormula sets CalculationFormula field to given value.

### HasCalculationFormula

`func (o *Property1) HasCalculationFormula() bool`

HasCalculationFormula returns a boolean if a field has been set.

### GetHasUniqueValue

`func (o *Property1) GetHasUniqueValue() bool`

GetHasUniqueValue returns the HasUniqueValue field if non-nil, zero value otherwise.

### GetHasUniqueValueOk

`func (o *Property1) GetHasUniqueValueOk() (*bool, bool)`

GetHasUniqueValueOk returns a tuple with the HasUniqueValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUniqueValue

`func (o *Property1) SetHasUniqueValue(v bool)`

SetHasUniqueValue sets HasUniqueValue field to given value.

### HasHasUniqueValue

`func (o *Property1) HasHasUniqueValue() bool`

HasHasUniqueValue returns a boolean if a field has been set.

### GetFieldType

`func (o *Property1) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *Property1) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *Property1) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetUpdatedUserId

`func (o *Property1) GetUpdatedUserId() string`

GetUpdatedUserId returns the UpdatedUserId field if non-nil, zero value otherwise.

### GetUpdatedUserIdOk

`func (o *Property1) GetUpdatedUserIdOk() (*string, bool)`

GetUpdatedUserIdOk returns a tuple with the UpdatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedUserId

`func (o *Property1) SetUpdatedUserId(v string)`

SetUpdatedUserId sets UpdatedUserId field to given value.

### HasUpdatedUserId

`func (o *Property1) HasUpdatedUserId() bool`

HasUpdatedUserId returns a boolean if a field has been set.

### GetCalculated

`func (o *Property1) GetCalculated() bool`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *Property1) GetCalculatedOk() (*bool, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *Property1) SetCalculated(v bool)`

SetCalculated sets Calculated field to given value.

### HasCalculated

`func (o *Property1) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### GetExternalOptions

`func (o *Property1) GetExternalOptions() bool`

GetExternalOptions returns the ExternalOptions field if non-nil, zero value otherwise.

### GetExternalOptionsOk

`func (o *Property1) GetExternalOptionsOk() (*bool, bool)`

GetExternalOptionsOk returns a tuple with the ExternalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOptions

`func (o *Property1) SetExternalOptions(v bool)`

SetExternalOptions sets ExternalOptions field to given value.

### HasExternalOptions

`func (o *Property1) HasExternalOptions() bool`

HasExternalOptions returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Property1) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Property1) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Property1) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Property1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


