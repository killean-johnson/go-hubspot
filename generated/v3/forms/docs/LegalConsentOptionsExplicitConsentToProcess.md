# LegalConsentOptionsExplicitConsentToProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationsCheckboxes** | [**[]LegalConsentCheckbox**](LegalConsentCheckbox.md) |  | 
**CommunicationConsentText** | Pointer to **string** |  | [optional] 
**ConsentToProcessCheckboxLabel** | Pointer to **string** |  | [optional] 
**ConsentToProcessFooterText** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "explicit_consent_to_process"]
**PrivacyText** | **string** |  | 
**ConsentToProcessText** | Pointer to **string** |  | [optional] 

## Methods

### NewLegalConsentOptionsExplicitConsentToProcess

`func NewLegalConsentOptionsExplicitConsentToProcess(communicationsCheckboxes []LegalConsentCheckbox, type_ string, privacyText string, ) *LegalConsentOptionsExplicitConsentToProcess`

NewLegalConsentOptionsExplicitConsentToProcess instantiates a new LegalConsentOptionsExplicitConsentToProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalConsentOptionsExplicitConsentToProcessWithDefaults

`func NewLegalConsentOptionsExplicitConsentToProcessWithDefaults() *LegalConsentOptionsExplicitConsentToProcess`

NewLegalConsentOptionsExplicitConsentToProcessWithDefaults instantiates a new LegalConsentOptionsExplicitConsentToProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationsCheckboxes

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetCommunicationsCheckboxes() []LegalConsentCheckbox`

GetCommunicationsCheckboxes returns the CommunicationsCheckboxes field if non-nil, zero value otherwise.

### GetCommunicationsCheckboxesOk

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetCommunicationsCheckboxesOk() (*[]LegalConsentCheckbox, bool)`

GetCommunicationsCheckboxesOk returns a tuple with the CommunicationsCheckboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationsCheckboxes

`func (o *LegalConsentOptionsExplicitConsentToProcess) SetCommunicationsCheckboxes(v []LegalConsentCheckbox)`

SetCommunicationsCheckboxes sets CommunicationsCheckboxes field to given value.


### GetCommunicationConsentText

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetCommunicationConsentText() string`

GetCommunicationConsentText returns the CommunicationConsentText field if non-nil, zero value otherwise.

### GetCommunicationConsentTextOk

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetCommunicationConsentTextOk() (*string, bool)`

GetCommunicationConsentTextOk returns a tuple with the CommunicationConsentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationConsentText

`func (o *LegalConsentOptionsExplicitConsentToProcess) SetCommunicationConsentText(v string)`

SetCommunicationConsentText sets CommunicationConsentText field to given value.

### HasCommunicationConsentText

`func (o *LegalConsentOptionsExplicitConsentToProcess) HasCommunicationConsentText() bool`

HasCommunicationConsentText returns a boolean if a field has been set.

### GetConsentToProcessCheckboxLabel

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetConsentToProcessCheckboxLabel() string`

GetConsentToProcessCheckboxLabel returns the ConsentToProcessCheckboxLabel field if non-nil, zero value otherwise.

### GetConsentToProcessCheckboxLabelOk

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetConsentToProcessCheckboxLabelOk() (*string, bool)`

GetConsentToProcessCheckboxLabelOk returns a tuple with the ConsentToProcessCheckboxLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentToProcessCheckboxLabel

`func (o *LegalConsentOptionsExplicitConsentToProcess) SetConsentToProcessCheckboxLabel(v string)`

SetConsentToProcessCheckboxLabel sets ConsentToProcessCheckboxLabel field to given value.

### HasConsentToProcessCheckboxLabel

`func (o *LegalConsentOptionsExplicitConsentToProcess) HasConsentToProcessCheckboxLabel() bool`

HasConsentToProcessCheckboxLabel returns a boolean if a field has been set.

### GetConsentToProcessFooterText

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetConsentToProcessFooterText() string`

GetConsentToProcessFooterText returns the ConsentToProcessFooterText field if non-nil, zero value otherwise.

### GetConsentToProcessFooterTextOk

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetConsentToProcessFooterTextOk() (*string, bool)`

GetConsentToProcessFooterTextOk returns a tuple with the ConsentToProcessFooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentToProcessFooterText

`func (o *LegalConsentOptionsExplicitConsentToProcess) SetConsentToProcessFooterText(v string)`

SetConsentToProcessFooterText sets ConsentToProcessFooterText field to given value.

### HasConsentToProcessFooterText

`func (o *LegalConsentOptionsExplicitConsentToProcess) HasConsentToProcessFooterText() bool`

HasConsentToProcessFooterText returns a boolean if a field has been set.

### GetType

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LegalConsentOptionsExplicitConsentToProcess) SetType(v string)`

SetType sets Type field to given value.


### GetPrivacyText

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetPrivacyText() string`

GetPrivacyText returns the PrivacyText field if non-nil, zero value otherwise.

### GetPrivacyTextOk

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetPrivacyTextOk() (*string, bool)`

GetPrivacyTextOk returns a tuple with the PrivacyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyText

`func (o *LegalConsentOptionsExplicitConsentToProcess) SetPrivacyText(v string)`

SetPrivacyText sets PrivacyText field to given value.


### GetConsentToProcessText

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetConsentToProcessText() string`

GetConsentToProcessText returns the ConsentToProcessText field if non-nil, zero value otherwise.

### GetConsentToProcessTextOk

`func (o *LegalConsentOptionsExplicitConsentToProcess) GetConsentToProcessTextOk() (*string, bool)`

GetConsentToProcessTextOk returns a tuple with the ConsentToProcessText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentToProcessText

`func (o *LegalConsentOptionsExplicitConsentToProcess) SetConsentToProcessText(v string)`

SetConsentToProcessText sets ConsentToProcessText field to given value.

### HasConsentToProcessText

`func (o *LegalConsentOptionsExplicitConsentToProcess) HasConsentToProcessText() bool`

HasConsentToProcessText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


