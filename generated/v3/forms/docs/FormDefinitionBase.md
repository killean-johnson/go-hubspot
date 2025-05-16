# FormDefinitionBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormType** | **string** |  | [default to "hubspot"]
**Id** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Archived** | **bool** |  | 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 
**FieldGroups** | [**[]FieldGroup**](FieldGroup.md) |  | 
**Configuration** | [**HubSpotFormConfiguration**](HubSpotFormConfiguration.md) |  | 
**DisplayOptions** | [**FormDisplayOptions**](FormDisplayOptions.md) |  | 
**LegalConsentOptions** | [**HubSpotFormDefinitionAllOfLegalConsentOptions**](HubSpotFormDefinitionAllOfLegalConsentOptions.md) |  | 

## Methods

### NewFormDefinitionBase

`func NewFormDefinitionBase(formType string, id string, name string, createdAt time.Time, updatedAt time.Time, archived bool, fieldGroups []FieldGroup, configuration HubSpotFormConfiguration, displayOptions FormDisplayOptions, legalConsentOptions HubSpotFormDefinitionAllOfLegalConsentOptions, ) *FormDefinitionBase`

NewFormDefinitionBase instantiates a new FormDefinitionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDefinitionBaseWithDefaults

`func NewFormDefinitionBaseWithDefaults() *FormDefinitionBase`

NewFormDefinitionBaseWithDefaults instantiates a new FormDefinitionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormType

`func (o *FormDefinitionBase) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *FormDefinitionBase) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *FormDefinitionBase) SetFormType(v string)`

SetFormType sets FormType field to given value.


### GetId

`func (o *FormDefinitionBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormDefinitionBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormDefinitionBase) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FormDefinitionBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormDefinitionBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormDefinitionBase) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *FormDefinitionBase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormDefinitionBase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormDefinitionBase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FormDefinitionBase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormDefinitionBase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormDefinitionBase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetArchived

`func (o *FormDefinitionBase) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FormDefinitionBase) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FormDefinitionBase) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchivedAt

`func (o *FormDefinitionBase) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *FormDefinitionBase) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *FormDefinitionBase) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *FormDefinitionBase) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetFieldGroups

`func (o *FormDefinitionBase) GetFieldGroups() []FieldGroup`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *FormDefinitionBase) GetFieldGroupsOk() (*[]FieldGroup, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *FormDefinitionBase) SetFieldGroups(v []FieldGroup)`

SetFieldGroups sets FieldGroups field to given value.


### GetConfiguration

`func (o *FormDefinitionBase) GetConfiguration() HubSpotFormConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *FormDefinitionBase) GetConfigurationOk() (*HubSpotFormConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *FormDefinitionBase) SetConfiguration(v HubSpotFormConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetDisplayOptions

`func (o *FormDefinitionBase) GetDisplayOptions() FormDisplayOptions`

GetDisplayOptions returns the DisplayOptions field if non-nil, zero value otherwise.

### GetDisplayOptionsOk

`func (o *FormDefinitionBase) GetDisplayOptionsOk() (*FormDisplayOptions, bool)`

GetDisplayOptionsOk returns a tuple with the DisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOptions

`func (o *FormDefinitionBase) SetDisplayOptions(v FormDisplayOptions)`

SetDisplayOptions sets DisplayOptions field to given value.


### GetLegalConsentOptions

`func (o *FormDefinitionBase) GetLegalConsentOptions() HubSpotFormDefinitionAllOfLegalConsentOptions`

GetLegalConsentOptions returns the LegalConsentOptions field if non-nil, zero value otherwise.

### GetLegalConsentOptionsOk

`func (o *FormDefinitionBase) GetLegalConsentOptionsOk() (*HubSpotFormDefinitionAllOfLegalConsentOptions, bool)`

GetLegalConsentOptionsOk returns a tuple with the LegalConsentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalConsentOptions

`func (o *FormDefinitionBase) SetLegalConsentOptions(v HubSpotFormDefinitionAllOfLegalConsentOptions)`

SetLegalConsentOptions sets LegalConsentOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


