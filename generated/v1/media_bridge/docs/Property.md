# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FavoritedOrder** | **int32** |  | 
**ReadOnlyValue** | **bool** |  | 
**Hidden** | **bool** |  | 
**OptionSortStrategy** | **string** |  | 
**DisplayOrder** | **int32** |  | 
**Description** | **string** |  | 
**ShowCurrencySymbol** | **bool** |  | 
**Type** | **string** |  | 
**ReadOnlyDefinition** | **bool** |  | 
**HubspotDefined** | **bool** |  | 
**AllowedObjectTypes** | [**[]ObjectTypeIdProto**](ObjectTypeIdProto.md) |  | 
**CreatedAt** | **int64** |  | 
**SearchTextAnalysisMode** | **string** |  | 
**TextDisplayHint** | **string** |  | 
**PortalId** | **int64** |  | 
**Options** | [**[]Option**](Option.md) |  | 
**IsPartial** | **bool** |  | 
**OptionsAreMutable** | **bool** |  | 
**HasUniqueValue** | **bool** |  | 
**Calculated** | **bool** |  | 
**ExternalOptions** | **bool** |  | 
**Favorited** | **bool** |  | 
**UpdatedAt** | **int64** |  | 
**CreatedUserId** | **int64** |  | 
**FromUserId** | **int64** |  | 
**MutableDefinitionNotDeletable** | **bool** |  | 
**IsMultiValued** | **bool** |  | 
**CurrencyPropertyName** | **string** |  | 
**IsCustomizedDefault** | **bool** |  | 
**Label** | **string** |  | 
**FormField** | **bool** |  | 
**DisplayMode** | **string** |  | 
**GroupName** | **string** |  | 
**Deleted** | **bool** |  | 
**ReferencedObjectType** | **string** |  | 
**Name** | **string** |  | 
**SearchableInGlobalSearch** | **bool** |  | 
**ExternalOptionsReferenceType** | **string** |  | 
**NumberDisplayHint** | **string** |  | 
**FieldType** | **string** |  | 

## Methods

### NewProperty

`func NewProperty(favoritedOrder int32, readOnlyValue bool, hidden bool, optionSortStrategy string, displayOrder int32, description string, showCurrencySymbol bool, type_ string, readOnlyDefinition bool, hubspotDefined bool, allowedObjectTypes []ObjectTypeIdProto, createdAt int64, searchTextAnalysisMode string, textDisplayHint string, portalId int64, options []Option, isPartial bool, optionsAreMutable bool, hasUniqueValue bool, calculated bool, externalOptions bool, favorited bool, updatedAt int64, createdUserId int64, fromUserId int64, mutableDefinitionNotDeletable bool, isMultiValued bool, currencyPropertyName string, isCustomizedDefault bool, label string, formField bool, displayMode string, groupName string, deleted bool, referencedObjectType string, name string, searchableInGlobalSearch bool, externalOptionsReferenceType string, numberDisplayHint string, fieldType string, ) *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavoritedOrder

`func (o *Property) GetFavoritedOrder() int32`

GetFavoritedOrder returns the FavoritedOrder field if non-nil, zero value otherwise.

### GetFavoritedOrderOk

`func (o *Property) GetFavoritedOrderOk() (*int32, bool)`

GetFavoritedOrderOk returns a tuple with the FavoritedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoritedOrder

`func (o *Property) SetFavoritedOrder(v int32)`

SetFavoritedOrder sets FavoritedOrder field to given value.


### GetReadOnlyValue

`func (o *Property) GetReadOnlyValue() bool`

GetReadOnlyValue returns the ReadOnlyValue field if non-nil, zero value otherwise.

### GetReadOnlyValueOk

`func (o *Property) GetReadOnlyValueOk() (*bool, bool)`

GetReadOnlyValueOk returns a tuple with the ReadOnlyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyValue

`func (o *Property) SetReadOnlyValue(v bool)`

SetReadOnlyValue sets ReadOnlyValue field to given value.


### GetHidden

`func (o *Property) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Property) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Property) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetOptionSortStrategy

`func (o *Property) GetOptionSortStrategy() string`

GetOptionSortStrategy returns the OptionSortStrategy field if non-nil, zero value otherwise.

### GetOptionSortStrategyOk

`func (o *Property) GetOptionSortStrategyOk() (*string, bool)`

GetOptionSortStrategyOk returns a tuple with the OptionSortStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSortStrategy

`func (o *Property) SetOptionSortStrategy(v string)`

SetOptionSortStrategy sets OptionSortStrategy field to given value.


### GetDisplayOrder

`func (o *Property) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *Property) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *Property) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetDescription

`func (o *Property) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Property) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Property) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetShowCurrencySymbol

`func (o *Property) GetShowCurrencySymbol() bool`

GetShowCurrencySymbol returns the ShowCurrencySymbol field if non-nil, zero value otherwise.

### GetShowCurrencySymbolOk

`func (o *Property) GetShowCurrencySymbolOk() (*bool, bool)`

GetShowCurrencySymbolOk returns a tuple with the ShowCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCurrencySymbol

`func (o *Property) SetShowCurrencySymbol(v bool)`

SetShowCurrencySymbol sets ShowCurrencySymbol field to given value.


### GetType

`func (o *Property) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Property) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Property) SetType(v string)`

SetType sets Type field to given value.


### GetReadOnlyDefinition

`func (o *Property) GetReadOnlyDefinition() bool`

GetReadOnlyDefinition returns the ReadOnlyDefinition field if non-nil, zero value otherwise.

### GetReadOnlyDefinitionOk

`func (o *Property) GetReadOnlyDefinitionOk() (*bool, bool)`

GetReadOnlyDefinitionOk returns a tuple with the ReadOnlyDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyDefinition

`func (o *Property) SetReadOnlyDefinition(v bool)`

SetReadOnlyDefinition sets ReadOnlyDefinition field to given value.


### GetHubspotDefined

`func (o *Property) GetHubspotDefined() bool`

GetHubspotDefined returns the HubspotDefined field if non-nil, zero value otherwise.

### GetHubspotDefinedOk

`func (o *Property) GetHubspotDefinedOk() (*bool, bool)`

GetHubspotDefinedOk returns a tuple with the HubspotDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubspotDefined

`func (o *Property) SetHubspotDefined(v bool)`

SetHubspotDefined sets HubspotDefined field to given value.


### GetAllowedObjectTypes

`func (o *Property) GetAllowedObjectTypes() []ObjectTypeIdProto`

GetAllowedObjectTypes returns the AllowedObjectTypes field if non-nil, zero value otherwise.

### GetAllowedObjectTypesOk

`func (o *Property) GetAllowedObjectTypesOk() (*[]ObjectTypeIdProto, bool)`

GetAllowedObjectTypesOk returns a tuple with the AllowedObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedObjectTypes

`func (o *Property) SetAllowedObjectTypes(v []ObjectTypeIdProto)`

SetAllowedObjectTypes sets AllowedObjectTypes field to given value.


### GetCreatedAt

`func (o *Property) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Property) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Property) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetSearchTextAnalysisMode

`func (o *Property) GetSearchTextAnalysisMode() string`

GetSearchTextAnalysisMode returns the SearchTextAnalysisMode field if non-nil, zero value otherwise.

### GetSearchTextAnalysisModeOk

`func (o *Property) GetSearchTextAnalysisModeOk() (*string, bool)`

GetSearchTextAnalysisModeOk returns a tuple with the SearchTextAnalysisMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTextAnalysisMode

`func (o *Property) SetSearchTextAnalysisMode(v string)`

SetSearchTextAnalysisMode sets SearchTextAnalysisMode field to given value.


### GetTextDisplayHint

`func (o *Property) GetTextDisplayHint() string`

GetTextDisplayHint returns the TextDisplayHint field if non-nil, zero value otherwise.

### GetTextDisplayHintOk

`func (o *Property) GetTextDisplayHintOk() (*string, bool)`

GetTextDisplayHintOk returns a tuple with the TextDisplayHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDisplayHint

`func (o *Property) SetTextDisplayHint(v string)`

SetTextDisplayHint sets TextDisplayHint field to given value.


### GetPortalId

`func (o *Property) GetPortalId() int64`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *Property) GetPortalIdOk() (*int64, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *Property) SetPortalId(v int64)`

SetPortalId sets PortalId field to given value.


### GetOptions

`func (o *Property) GetOptions() []Option`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Property) GetOptionsOk() (*[]Option, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Property) SetOptions(v []Option)`

SetOptions sets Options field to given value.


### GetIsPartial

`func (o *Property) GetIsPartial() bool`

GetIsPartial returns the IsPartial field if non-nil, zero value otherwise.

### GetIsPartialOk

`func (o *Property) GetIsPartialOk() (*bool, bool)`

GetIsPartialOk returns a tuple with the IsPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartial

`func (o *Property) SetIsPartial(v bool)`

SetIsPartial sets IsPartial field to given value.


### GetOptionsAreMutable

`func (o *Property) GetOptionsAreMutable() bool`

GetOptionsAreMutable returns the OptionsAreMutable field if non-nil, zero value otherwise.

### GetOptionsAreMutableOk

`func (o *Property) GetOptionsAreMutableOk() (*bool, bool)`

GetOptionsAreMutableOk returns a tuple with the OptionsAreMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsAreMutable

`func (o *Property) SetOptionsAreMutable(v bool)`

SetOptionsAreMutable sets OptionsAreMutable field to given value.


### GetHasUniqueValue

`func (o *Property) GetHasUniqueValue() bool`

GetHasUniqueValue returns the HasUniqueValue field if non-nil, zero value otherwise.

### GetHasUniqueValueOk

`func (o *Property) GetHasUniqueValueOk() (*bool, bool)`

GetHasUniqueValueOk returns a tuple with the HasUniqueValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUniqueValue

`func (o *Property) SetHasUniqueValue(v bool)`

SetHasUniqueValue sets HasUniqueValue field to given value.


### GetCalculated

`func (o *Property) GetCalculated() bool`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *Property) GetCalculatedOk() (*bool, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *Property) SetCalculated(v bool)`

SetCalculated sets Calculated field to given value.


### GetExternalOptions

`func (o *Property) GetExternalOptions() bool`

GetExternalOptions returns the ExternalOptions field if non-nil, zero value otherwise.

### GetExternalOptionsOk

`func (o *Property) GetExternalOptionsOk() (*bool, bool)`

GetExternalOptionsOk returns a tuple with the ExternalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOptions

`func (o *Property) SetExternalOptions(v bool)`

SetExternalOptions sets ExternalOptions field to given value.


### GetFavorited

`func (o *Property) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *Property) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *Property) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetUpdatedAt

`func (o *Property) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Property) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Property) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedUserId

`func (o *Property) GetCreatedUserId() int64`

GetCreatedUserId returns the CreatedUserId field if non-nil, zero value otherwise.

### GetCreatedUserIdOk

`func (o *Property) GetCreatedUserIdOk() (*int64, bool)`

GetCreatedUserIdOk returns a tuple with the CreatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUserId

`func (o *Property) SetCreatedUserId(v int64)`

SetCreatedUserId sets CreatedUserId field to given value.


### GetFromUserId

`func (o *Property) GetFromUserId() int64`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *Property) GetFromUserIdOk() (*int64, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *Property) SetFromUserId(v int64)`

SetFromUserId sets FromUserId field to given value.


### GetMutableDefinitionNotDeletable

`func (o *Property) GetMutableDefinitionNotDeletable() bool`

GetMutableDefinitionNotDeletable returns the MutableDefinitionNotDeletable field if non-nil, zero value otherwise.

### GetMutableDefinitionNotDeletableOk

`func (o *Property) GetMutableDefinitionNotDeletableOk() (*bool, bool)`

GetMutableDefinitionNotDeletableOk returns a tuple with the MutableDefinitionNotDeletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutableDefinitionNotDeletable

`func (o *Property) SetMutableDefinitionNotDeletable(v bool)`

SetMutableDefinitionNotDeletable sets MutableDefinitionNotDeletable field to given value.


### GetIsMultiValued

`func (o *Property) GetIsMultiValued() bool`

GetIsMultiValued returns the IsMultiValued field if non-nil, zero value otherwise.

### GetIsMultiValuedOk

`func (o *Property) GetIsMultiValuedOk() (*bool, bool)`

GetIsMultiValuedOk returns a tuple with the IsMultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiValued

`func (o *Property) SetIsMultiValued(v bool)`

SetIsMultiValued sets IsMultiValued field to given value.


### GetCurrencyPropertyName

`func (o *Property) GetCurrencyPropertyName() string`

GetCurrencyPropertyName returns the CurrencyPropertyName field if non-nil, zero value otherwise.

### GetCurrencyPropertyNameOk

`func (o *Property) GetCurrencyPropertyNameOk() (*string, bool)`

GetCurrencyPropertyNameOk returns a tuple with the CurrencyPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyPropertyName

`func (o *Property) SetCurrencyPropertyName(v string)`

SetCurrencyPropertyName sets CurrencyPropertyName field to given value.


### GetIsCustomizedDefault

`func (o *Property) GetIsCustomizedDefault() bool`

GetIsCustomizedDefault returns the IsCustomizedDefault field if non-nil, zero value otherwise.

### GetIsCustomizedDefaultOk

`func (o *Property) GetIsCustomizedDefaultOk() (*bool, bool)`

GetIsCustomizedDefaultOk returns a tuple with the IsCustomizedDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomizedDefault

`func (o *Property) SetIsCustomizedDefault(v bool)`

SetIsCustomizedDefault sets IsCustomizedDefault field to given value.


### GetLabel

`func (o *Property) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Property) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Property) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFormField

`func (o *Property) GetFormField() bool`

GetFormField returns the FormField field if non-nil, zero value otherwise.

### GetFormFieldOk

`func (o *Property) GetFormFieldOk() (*bool, bool)`

GetFormFieldOk returns a tuple with the FormField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormField

`func (o *Property) SetFormField(v bool)`

SetFormField sets FormField field to given value.


### GetDisplayMode

`func (o *Property) GetDisplayMode() string`

GetDisplayMode returns the DisplayMode field if non-nil, zero value otherwise.

### GetDisplayModeOk

`func (o *Property) GetDisplayModeOk() (*string, bool)`

GetDisplayModeOk returns a tuple with the DisplayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMode

`func (o *Property) SetDisplayMode(v string)`

SetDisplayMode sets DisplayMode field to given value.


### GetGroupName

`func (o *Property) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *Property) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *Property) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetDeleted

`func (o *Property) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Property) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Property) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetReferencedObjectType

`func (o *Property) GetReferencedObjectType() string`

GetReferencedObjectType returns the ReferencedObjectType field if non-nil, zero value otherwise.

### GetReferencedObjectTypeOk

`func (o *Property) GetReferencedObjectTypeOk() (*string, bool)`

GetReferencedObjectTypeOk returns a tuple with the ReferencedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedObjectType

`func (o *Property) SetReferencedObjectType(v string)`

SetReferencedObjectType sets ReferencedObjectType field to given value.


### GetName

`func (o *Property) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Property) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Property) SetName(v string)`

SetName sets Name field to given value.


### GetSearchableInGlobalSearch

`func (o *Property) GetSearchableInGlobalSearch() bool`

GetSearchableInGlobalSearch returns the SearchableInGlobalSearch field if non-nil, zero value otherwise.

### GetSearchableInGlobalSearchOk

`func (o *Property) GetSearchableInGlobalSearchOk() (*bool, bool)`

GetSearchableInGlobalSearchOk returns a tuple with the SearchableInGlobalSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableInGlobalSearch

`func (o *Property) SetSearchableInGlobalSearch(v bool)`

SetSearchableInGlobalSearch sets SearchableInGlobalSearch field to given value.


### GetExternalOptionsReferenceType

`func (o *Property) GetExternalOptionsReferenceType() string`

GetExternalOptionsReferenceType returns the ExternalOptionsReferenceType field if non-nil, zero value otherwise.

### GetExternalOptionsReferenceTypeOk

`func (o *Property) GetExternalOptionsReferenceTypeOk() (*string, bool)`

GetExternalOptionsReferenceTypeOk returns a tuple with the ExternalOptionsReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOptionsReferenceType

`func (o *Property) SetExternalOptionsReferenceType(v string)`

SetExternalOptionsReferenceType sets ExternalOptionsReferenceType field to given value.


### GetNumberDisplayHint

`func (o *Property) GetNumberDisplayHint() string`

GetNumberDisplayHint returns the NumberDisplayHint field if non-nil, zero value otherwise.

### GetNumberDisplayHintOk

`func (o *Property) GetNumberDisplayHintOk() (*string, bool)`

GetNumberDisplayHintOk returns a tuple with the NumberDisplayHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberDisplayHint

`func (o *Property) SetNumberDisplayHint(v string)`

SetNumberDisplayHint sets NumberDisplayHint field to given value.


### GetFieldType

`func (o *Property) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *Property) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *Property) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


