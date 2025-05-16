# ExternalBehavioralEventTypeDefinitionEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyDefinitions** | [**[]ExternalBehavioralEventPropertyCreate**](ExternalBehavioralEventPropertyCreate.md) | List of custom properties on event | 
**Name** | Pointer to **string** | Internal event name, which must be used when referencing the event from this event definitions API. If a name is not supplied, one will be generated based on the label. The &#x60;name&#x60; value will also be used to automatically generate a &#x60;fullyQualifiedName&#x60; for the event definition, which you&#39;ll use when sending event completions to this event.  | [optional] 
**Description** | Pointer to **string** | A description of the event that will be shown as help text in HubSpot. | [optional] 
**Label** | **string** | Human readable label for the event. Used in HubSpot UI | 
**PrimaryObject** | Pointer to **string** | The object type to associate this event to. Can be one of CONTACT, COMPANY, DEAL, TICKET. If no primaryObject is supplied, we will default to associating the event to CONTACT objects. | [optional] 

## Methods

### NewExternalBehavioralEventTypeDefinitionEgg

`func NewExternalBehavioralEventTypeDefinitionEgg(propertyDefinitions []ExternalBehavioralEventPropertyCreate, label string, ) *ExternalBehavioralEventTypeDefinitionEgg`

NewExternalBehavioralEventTypeDefinitionEgg instantiates a new ExternalBehavioralEventTypeDefinitionEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalBehavioralEventTypeDefinitionEggWithDefaults

`func NewExternalBehavioralEventTypeDefinitionEggWithDefaults() *ExternalBehavioralEventTypeDefinitionEgg`

NewExternalBehavioralEventTypeDefinitionEggWithDefaults instantiates a new ExternalBehavioralEventTypeDefinitionEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyDefinitions

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetPropertyDefinitions() []ExternalBehavioralEventPropertyCreate`

GetPropertyDefinitions returns the PropertyDefinitions field if non-nil, zero value otherwise.

### GetPropertyDefinitionsOk

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetPropertyDefinitionsOk() (*[]ExternalBehavioralEventPropertyCreate, bool)`

GetPropertyDefinitionsOk returns a tuple with the PropertyDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDefinitions

`func (o *ExternalBehavioralEventTypeDefinitionEgg) SetPropertyDefinitions(v []ExternalBehavioralEventPropertyCreate)`

SetPropertyDefinitions sets PropertyDefinitions field to given value.


### GetName

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalBehavioralEventTypeDefinitionEgg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalBehavioralEventTypeDefinitionEgg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalBehavioralEventTypeDefinitionEgg) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalBehavioralEventTypeDefinitionEgg) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExternalBehavioralEventTypeDefinitionEgg) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPrimaryObject

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetPrimaryObject() string`

GetPrimaryObject returns the PrimaryObject field if non-nil, zero value otherwise.

### GetPrimaryObjectOk

`func (o *ExternalBehavioralEventTypeDefinitionEgg) GetPrimaryObjectOk() (*string, bool)`

GetPrimaryObjectOk returns a tuple with the PrimaryObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryObject

`func (o *ExternalBehavioralEventTypeDefinitionEgg) SetPrimaryObject(v string)`

SetPrimaryObject sets PrimaryObject field to given value.

### HasPrimaryObject

`func (o *ExternalBehavioralEventTypeDefinitionEgg) HasPrimaryObject() bool`

HasPrimaryObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


