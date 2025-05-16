# ExternalBehavioralEventPropertyDefinitionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**[]OptionInput**](OptionInput.md) | A list of available options for the property if it is an enumeration. NOTE: This field is only applicable for enumerated properties. | [optional] 
**Description** | Pointer to **string** | A description of the property that will be shown as help text in HubSpot. | [optional] 
**Label** | Pointer to **string** | Human readable label for the property. Used in HubSpot UI | [optional] 

## Methods

### NewExternalBehavioralEventPropertyDefinitionPatch

`func NewExternalBehavioralEventPropertyDefinitionPatch() *ExternalBehavioralEventPropertyDefinitionPatch`

NewExternalBehavioralEventPropertyDefinitionPatch instantiates a new ExternalBehavioralEventPropertyDefinitionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalBehavioralEventPropertyDefinitionPatchWithDefaults

`func NewExternalBehavioralEventPropertyDefinitionPatchWithDefaults() *ExternalBehavioralEventPropertyDefinitionPatch`

NewExternalBehavioralEventPropertyDefinitionPatchWithDefaults instantiates a new ExternalBehavioralEventPropertyDefinitionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) GetOptions() []OptionInput`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) GetOptionsOk() (*[]OptionInput, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) SetOptions(v []OptionInput)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ExternalBehavioralEventPropertyDefinitionPatch) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


