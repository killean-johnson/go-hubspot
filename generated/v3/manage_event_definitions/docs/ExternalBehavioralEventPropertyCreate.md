# ExternalBehavioralEventPropertyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Internal property name, which must be used when referencing the property from the API | [optional] 
**Options** | Pointer to [**[]OptionInput**](OptionInput.md) | A list of available options for the property if it is an enumeration. NOTE: This field is only applicable for enumerated properties. | [optional] 
**Description** | Pointer to **string** | A description of the property that will be shown as help text in HubSpot. | [optional] 
**Label** | **string** | Human readable label for the property. Used in HubSpot UI | 
**Type** | **string** | The data type of the property. Can be one of the following: [string, number, enumeration, datetime] | 

## Methods

### NewExternalBehavioralEventPropertyCreate

`func NewExternalBehavioralEventPropertyCreate(label string, type_ string, ) *ExternalBehavioralEventPropertyCreate`

NewExternalBehavioralEventPropertyCreate instantiates a new ExternalBehavioralEventPropertyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalBehavioralEventPropertyCreateWithDefaults

`func NewExternalBehavioralEventPropertyCreateWithDefaults() *ExternalBehavioralEventPropertyCreate`

NewExternalBehavioralEventPropertyCreateWithDefaults instantiates a new ExternalBehavioralEventPropertyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalBehavioralEventPropertyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalBehavioralEventPropertyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalBehavioralEventPropertyCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalBehavioralEventPropertyCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *ExternalBehavioralEventPropertyCreate) GetOptions() []OptionInput`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExternalBehavioralEventPropertyCreate) GetOptionsOk() (*[]OptionInput, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExternalBehavioralEventPropertyCreate) SetOptions(v []OptionInput)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ExternalBehavioralEventPropertyCreate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalBehavioralEventPropertyCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalBehavioralEventPropertyCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalBehavioralEventPropertyCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalBehavioralEventPropertyCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *ExternalBehavioralEventPropertyCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExternalBehavioralEventPropertyCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExternalBehavioralEventPropertyCreate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *ExternalBehavioralEventPropertyCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalBehavioralEventPropertyCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalBehavioralEventPropertyCreate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


