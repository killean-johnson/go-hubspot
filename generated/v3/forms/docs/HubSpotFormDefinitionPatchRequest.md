# HubSpotFormDefinitionPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldGroups** | Pointer to [**[]FieldGroup**](FieldGroup.md) | The fields in the form, grouped in rows. | [optional] 
**Archived** | Pointer to **bool** | Whether this form is archived. | [optional] 
**Configuration** | Pointer to [**HubSpotFormConfiguration**](HubSpotFormConfiguration.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the form. Expected to be unique for a hub. | [optional] 
**LegalConsentOptions** | Pointer to [**HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions**](HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions.md) |  | [optional] 
**DisplayOptions** | Pointer to [**FormDisplayOptions**](FormDisplayOptions.md) |  | [optional] 

## Methods

### NewHubSpotFormDefinitionPatchRequest

`func NewHubSpotFormDefinitionPatchRequest() *HubSpotFormDefinitionPatchRequest`

NewHubSpotFormDefinitionPatchRequest instantiates a new HubSpotFormDefinitionPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubSpotFormDefinitionPatchRequestWithDefaults

`func NewHubSpotFormDefinitionPatchRequestWithDefaults() *HubSpotFormDefinitionPatchRequest`

NewHubSpotFormDefinitionPatchRequestWithDefaults instantiates a new HubSpotFormDefinitionPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldGroups

`func (o *HubSpotFormDefinitionPatchRequest) GetFieldGroups() []FieldGroup`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *HubSpotFormDefinitionPatchRequest) GetFieldGroupsOk() (*[]FieldGroup, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *HubSpotFormDefinitionPatchRequest) SetFieldGroups(v []FieldGroup)`

SetFieldGroups sets FieldGroups field to given value.

### HasFieldGroups

`func (o *HubSpotFormDefinitionPatchRequest) HasFieldGroups() bool`

HasFieldGroups returns a boolean if a field has been set.

### GetArchived

`func (o *HubSpotFormDefinitionPatchRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *HubSpotFormDefinitionPatchRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *HubSpotFormDefinitionPatchRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *HubSpotFormDefinitionPatchRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetConfiguration

`func (o *HubSpotFormDefinitionPatchRequest) GetConfiguration() HubSpotFormConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *HubSpotFormDefinitionPatchRequest) GetConfigurationOk() (*HubSpotFormConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *HubSpotFormDefinitionPatchRequest) SetConfiguration(v HubSpotFormConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *HubSpotFormDefinitionPatchRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetName

`func (o *HubSpotFormDefinitionPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HubSpotFormDefinitionPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HubSpotFormDefinitionPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HubSpotFormDefinitionPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLegalConsentOptions

`func (o *HubSpotFormDefinitionPatchRequest) GetLegalConsentOptions() HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions`

GetLegalConsentOptions returns the LegalConsentOptions field if non-nil, zero value otherwise.

### GetLegalConsentOptionsOk

`func (o *HubSpotFormDefinitionPatchRequest) GetLegalConsentOptionsOk() (*HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions, bool)`

GetLegalConsentOptionsOk returns a tuple with the LegalConsentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalConsentOptions

`func (o *HubSpotFormDefinitionPatchRequest) SetLegalConsentOptions(v HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions)`

SetLegalConsentOptions sets LegalConsentOptions field to given value.

### HasLegalConsentOptions

`func (o *HubSpotFormDefinitionPatchRequest) HasLegalConsentOptions() bool`

HasLegalConsentOptions returns a boolean if a field has been set.

### GetDisplayOptions

`func (o *HubSpotFormDefinitionPatchRequest) GetDisplayOptions() FormDisplayOptions`

GetDisplayOptions returns the DisplayOptions field if non-nil, zero value otherwise.

### GetDisplayOptionsOk

`func (o *HubSpotFormDefinitionPatchRequest) GetDisplayOptionsOk() (*FormDisplayOptions, bool)`

GetDisplayOptionsOk returns a tuple with the DisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOptions

`func (o *HubSpotFormDefinitionPatchRequest) SetDisplayOptions(v FormDisplayOptions)`

SetDisplayOptions sets DisplayOptions field to given value.

### HasDisplayOptions

`func (o *HubSpotFormDefinitionPatchRequest) HasDisplayOptions() bool`

HasDisplayOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


