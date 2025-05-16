# ExternalLegalConsentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegitimateInterestSubscriptionTypes** | **[]int64** |  | 
**ProcessingConsentCheckboxLabel** | **string** |  | 
**CommunicationConsentText** | **string** |  | 
**LegitimateInterestLegalBasis** | Pointer to **string** |  | [optional] 
**ProcessingConsentType** | **string** |  | 
**PrivacyPolicyText** | **string** |  | 
**ProcessingConsentText** | **string** |  | 
**CommunicationConsentCheckboxes** | [**[]ExternalCommunicationConsentCheckbox**](ExternalCommunicationConsentCheckbox.md) |  | 
**ProcessingConsentFooterText** | **string** |  | 
**IsLegitimateInterest** | **bool** |  | 

## Methods

### NewExternalLegalConsentOptions

`func NewExternalLegalConsentOptions(legitimateInterestSubscriptionTypes []int64, processingConsentCheckboxLabel string, communicationConsentText string, processingConsentType string, privacyPolicyText string, processingConsentText string, communicationConsentCheckboxes []ExternalCommunicationConsentCheckbox, processingConsentFooterText string, isLegitimateInterest bool, ) *ExternalLegalConsentOptions`

NewExternalLegalConsentOptions instantiates a new ExternalLegalConsentOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalLegalConsentOptionsWithDefaults

`func NewExternalLegalConsentOptionsWithDefaults() *ExternalLegalConsentOptions`

NewExternalLegalConsentOptionsWithDefaults instantiates a new ExternalLegalConsentOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegitimateInterestSubscriptionTypes

`func (o *ExternalLegalConsentOptions) GetLegitimateInterestSubscriptionTypes() []int64`

GetLegitimateInterestSubscriptionTypes returns the LegitimateInterestSubscriptionTypes field if non-nil, zero value otherwise.

### GetLegitimateInterestSubscriptionTypesOk

`func (o *ExternalLegalConsentOptions) GetLegitimateInterestSubscriptionTypesOk() (*[]int64, bool)`

GetLegitimateInterestSubscriptionTypesOk returns a tuple with the LegitimateInterestSubscriptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegitimateInterestSubscriptionTypes

`func (o *ExternalLegalConsentOptions) SetLegitimateInterestSubscriptionTypes(v []int64)`

SetLegitimateInterestSubscriptionTypes sets LegitimateInterestSubscriptionTypes field to given value.


### GetProcessingConsentCheckboxLabel

`func (o *ExternalLegalConsentOptions) GetProcessingConsentCheckboxLabel() string`

GetProcessingConsentCheckboxLabel returns the ProcessingConsentCheckboxLabel field if non-nil, zero value otherwise.

### GetProcessingConsentCheckboxLabelOk

`func (o *ExternalLegalConsentOptions) GetProcessingConsentCheckboxLabelOk() (*string, bool)`

GetProcessingConsentCheckboxLabelOk returns a tuple with the ProcessingConsentCheckboxLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingConsentCheckboxLabel

`func (o *ExternalLegalConsentOptions) SetProcessingConsentCheckboxLabel(v string)`

SetProcessingConsentCheckboxLabel sets ProcessingConsentCheckboxLabel field to given value.


### GetCommunicationConsentText

`func (o *ExternalLegalConsentOptions) GetCommunicationConsentText() string`

GetCommunicationConsentText returns the CommunicationConsentText field if non-nil, zero value otherwise.

### GetCommunicationConsentTextOk

`func (o *ExternalLegalConsentOptions) GetCommunicationConsentTextOk() (*string, bool)`

GetCommunicationConsentTextOk returns a tuple with the CommunicationConsentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationConsentText

`func (o *ExternalLegalConsentOptions) SetCommunicationConsentText(v string)`

SetCommunicationConsentText sets CommunicationConsentText field to given value.


### GetLegitimateInterestLegalBasis

`func (o *ExternalLegalConsentOptions) GetLegitimateInterestLegalBasis() string`

GetLegitimateInterestLegalBasis returns the LegitimateInterestLegalBasis field if non-nil, zero value otherwise.

### GetLegitimateInterestLegalBasisOk

`func (o *ExternalLegalConsentOptions) GetLegitimateInterestLegalBasisOk() (*string, bool)`

GetLegitimateInterestLegalBasisOk returns a tuple with the LegitimateInterestLegalBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegitimateInterestLegalBasis

`func (o *ExternalLegalConsentOptions) SetLegitimateInterestLegalBasis(v string)`

SetLegitimateInterestLegalBasis sets LegitimateInterestLegalBasis field to given value.

### HasLegitimateInterestLegalBasis

`func (o *ExternalLegalConsentOptions) HasLegitimateInterestLegalBasis() bool`

HasLegitimateInterestLegalBasis returns a boolean if a field has been set.

### GetProcessingConsentType

`func (o *ExternalLegalConsentOptions) GetProcessingConsentType() string`

GetProcessingConsentType returns the ProcessingConsentType field if non-nil, zero value otherwise.

### GetProcessingConsentTypeOk

`func (o *ExternalLegalConsentOptions) GetProcessingConsentTypeOk() (*string, bool)`

GetProcessingConsentTypeOk returns a tuple with the ProcessingConsentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingConsentType

`func (o *ExternalLegalConsentOptions) SetProcessingConsentType(v string)`

SetProcessingConsentType sets ProcessingConsentType field to given value.


### GetPrivacyPolicyText

`func (o *ExternalLegalConsentOptions) GetPrivacyPolicyText() string`

GetPrivacyPolicyText returns the PrivacyPolicyText field if non-nil, zero value otherwise.

### GetPrivacyPolicyTextOk

`func (o *ExternalLegalConsentOptions) GetPrivacyPolicyTextOk() (*string, bool)`

GetPrivacyPolicyTextOk returns a tuple with the PrivacyPolicyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyText

`func (o *ExternalLegalConsentOptions) SetPrivacyPolicyText(v string)`

SetPrivacyPolicyText sets PrivacyPolicyText field to given value.


### GetProcessingConsentText

`func (o *ExternalLegalConsentOptions) GetProcessingConsentText() string`

GetProcessingConsentText returns the ProcessingConsentText field if non-nil, zero value otherwise.

### GetProcessingConsentTextOk

`func (o *ExternalLegalConsentOptions) GetProcessingConsentTextOk() (*string, bool)`

GetProcessingConsentTextOk returns a tuple with the ProcessingConsentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingConsentText

`func (o *ExternalLegalConsentOptions) SetProcessingConsentText(v string)`

SetProcessingConsentText sets ProcessingConsentText field to given value.


### GetCommunicationConsentCheckboxes

`func (o *ExternalLegalConsentOptions) GetCommunicationConsentCheckboxes() []ExternalCommunicationConsentCheckbox`

GetCommunicationConsentCheckboxes returns the CommunicationConsentCheckboxes field if non-nil, zero value otherwise.

### GetCommunicationConsentCheckboxesOk

`func (o *ExternalLegalConsentOptions) GetCommunicationConsentCheckboxesOk() (*[]ExternalCommunicationConsentCheckbox, bool)`

GetCommunicationConsentCheckboxesOk returns a tuple with the CommunicationConsentCheckboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationConsentCheckboxes

`func (o *ExternalLegalConsentOptions) SetCommunicationConsentCheckboxes(v []ExternalCommunicationConsentCheckbox)`

SetCommunicationConsentCheckboxes sets CommunicationConsentCheckboxes field to given value.


### GetProcessingConsentFooterText

`func (o *ExternalLegalConsentOptions) GetProcessingConsentFooterText() string`

GetProcessingConsentFooterText returns the ProcessingConsentFooterText field if non-nil, zero value otherwise.

### GetProcessingConsentFooterTextOk

`func (o *ExternalLegalConsentOptions) GetProcessingConsentFooterTextOk() (*string, bool)`

GetProcessingConsentFooterTextOk returns a tuple with the ProcessingConsentFooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingConsentFooterText

`func (o *ExternalLegalConsentOptions) SetProcessingConsentFooterText(v string)`

SetProcessingConsentFooterText sets ProcessingConsentFooterText field to given value.


### GetIsLegitimateInterest

`func (o *ExternalLegalConsentOptions) GetIsLegitimateInterest() bool`

GetIsLegitimateInterest returns the IsLegitimateInterest field if non-nil, zero value otherwise.

### GetIsLegitimateInterestOk

`func (o *ExternalLegalConsentOptions) GetIsLegitimateInterestOk() (*bool, bool)`

GetIsLegitimateInterestOk returns a tuple with the IsLegitimateInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegitimateInterest

`func (o *ExternalLegalConsentOptions) SetIsLegitimateInterest(v bool)`

SetIsLegitimateInterest sets IsLegitimateInterest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


