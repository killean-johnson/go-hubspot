# HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "none"]
**SubscriptionTypeIds** | **[]int64** |  | 
**LawfulBasis** | **string** |  | 
**PrivacyText** | **string** |  | 
**CommunicationsCheckboxes** | [**[]LegalConsentCheckbox**](LegalConsentCheckbox.md) |  | 
**CommunicationConsentText** | Pointer to **string** |  | [optional] 
**ConsentToProcessCheckboxLabel** | Pointer to **string** |  | [optional] 
**ConsentToProcessFooterText** | Pointer to **string** |  | [optional] 
**ConsentToProcessText** | Pointer to **string** |  | [optional] 

## Methods

### NewHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions

`func NewHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions(type_ string, subscriptionTypeIds []int64, lawfulBasis string, privacyText string, communicationsCheckboxes []LegalConsentCheckbox, ) *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions`

NewHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions instantiates a new HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptionsWithDefaults

`func NewHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptionsWithDefaults() *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions`

NewHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptionsWithDefaults instantiates a new HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) SetType(v string)`

SetType sets Type field to given value.


### GetSubscriptionTypeIds

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetSubscriptionTypeIds() []int64`

GetSubscriptionTypeIds returns the SubscriptionTypeIds field if non-nil, zero value otherwise.

### GetSubscriptionTypeIdsOk

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetSubscriptionTypeIdsOk() (*[]int64, bool)`

GetSubscriptionTypeIdsOk returns a tuple with the SubscriptionTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTypeIds

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) SetSubscriptionTypeIds(v []int64)`

SetSubscriptionTypeIds sets SubscriptionTypeIds field to given value.


### GetLawfulBasis

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetLawfulBasis() string`

GetLawfulBasis returns the LawfulBasis field if non-nil, zero value otherwise.

### GetLawfulBasisOk

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetLawfulBasisOk() (*string, bool)`

GetLawfulBasisOk returns a tuple with the LawfulBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLawfulBasis

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) SetLawfulBasis(v string)`

SetLawfulBasis sets LawfulBasis field to given value.


### GetPrivacyText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetPrivacyText() string`

GetPrivacyText returns the PrivacyText field if non-nil, zero value otherwise.

### GetPrivacyTextOk

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetPrivacyTextOk() (*string, bool)`

GetPrivacyTextOk returns a tuple with the PrivacyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) SetPrivacyText(v string)`

SetPrivacyText sets PrivacyText field to given value.


### GetCommunicationsCheckboxes

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetCommunicationsCheckboxes() []LegalConsentCheckbox`

GetCommunicationsCheckboxes returns the CommunicationsCheckboxes field if non-nil, zero value otherwise.

### GetCommunicationsCheckboxesOk

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetCommunicationsCheckboxesOk() (*[]LegalConsentCheckbox, bool)`

GetCommunicationsCheckboxesOk returns a tuple with the CommunicationsCheckboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationsCheckboxes

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) SetCommunicationsCheckboxes(v []LegalConsentCheckbox)`

SetCommunicationsCheckboxes sets CommunicationsCheckboxes field to given value.


### GetCommunicationConsentText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetCommunicationConsentText() string`

GetCommunicationConsentText returns the CommunicationConsentText field if non-nil, zero value otherwise.

### GetCommunicationConsentTextOk

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetCommunicationConsentTextOk() (*string, bool)`

GetCommunicationConsentTextOk returns a tuple with the CommunicationConsentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationConsentText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) SetCommunicationConsentText(v string)`

SetCommunicationConsentText sets CommunicationConsentText field to given value.

### HasCommunicationConsentText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) HasCommunicationConsentText() bool`

HasCommunicationConsentText returns a boolean if a field has been set.

### GetConsentToProcessCheckboxLabel

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetConsentToProcessCheckboxLabel() string`

GetConsentToProcessCheckboxLabel returns the ConsentToProcessCheckboxLabel field if non-nil, zero value otherwise.

### GetConsentToProcessCheckboxLabelOk

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetConsentToProcessCheckboxLabelOk() (*string, bool)`

GetConsentToProcessCheckboxLabelOk returns a tuple with the ConsentToProcessCheckboxLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentToProcessCheckboxLabel

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) SetConsentToProcessCheckboxLabel(v string)`

SetConsentToProcessCheckboxLabel sets ConsentToProcessCheckboxLabel field to given value.

### HasConsentToProcessCheckboxLabel

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) HasConsentToProcessCheckboxLabel() bool`

HasConsentToProcessCheckboxLabel returns a boolean if a field has been set.

### GetConsentToProcessFooterText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetConsentToProcessFooterText() string`

GetConsentToProcessFooterText returns the ConsentToProcessFooterText field if non-nil, zero value otherwise.

### GetConsentToProcessFooterTextOk

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetConsentToProcessFooterTextOk() (*string, bool)`

GetConsentToProcessFooterTextOk returns a tuple with the ConsentToProcessFooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentToProcessFooterText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) SetConsentToProcessFooterText(v string)`

SetConsentToProcessFooterText sets ConsentToProcessFooterText field to given value.

### HasConsentToProcessFooterText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) HasConsentToProcessFooterText() bool`

HasConsentToProcessFooterText returns a boolean if a field has been set.

### GetConsentToProcessText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetConsentToProcessText() string`

GetConsentToProcessText returns the ConsentToProcessText field if non-nil, zero value otherwise.

### GetConsentToProcessTextOk

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetConsentToProcessTextOk() (*string, bool)`

GetConsentToProcessTextOk returns a tuple with the ConsentToProcessText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentToProcessText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) SetConsentToProcessText(v string)`

SetConsentToProcessText sets ConsentToProcessText field to given value.

### HasConsentToProcessText

`func (o *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) HasConsentToProcessText() bool`

HasConsentToProcessText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


