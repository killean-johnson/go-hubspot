# LegalConsentCheckbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionTypeId** | **int32** |  | 
**Label** | **string** | The main label for the form field. | 
**Required** | **bool** | Whether this checkbox is required when submitting the form. | 

## Methods

### NewLegalConsentCheckbox

`func NewLegalConsentCheckbox(subscriptionTypeId int32, label string, required bool, ) *LegalConsentCheckbox`

NewLegalConsentCheckbox instantiates a new LegalConsentCheckbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalConsentCheckboxWithDefaults

`func NewLegalConsentCheckboxWithDefaults() *LegalConsentCheckbox`

NewLegalConsentCheckboxWithDefaults instantiates a new LegalConsentCheckbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionTypeId

`func (o *LegalConsentCheckbox) GetSubscriptionTypeId() int32`

GetSubscriptionTypeId returns the SubscriptionTypeId field if non-nil, zero value otherwise.

### GetSubscriptionTypeIdOk

`func (o *LegalConsentCheckbox) GetSubscriptionTypeIdOk() (*int32, bool)`

GetSubscriptionTypeIdOk returns a tuple with the SubscriptionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTypeId

`func (o *LegalConsentCheckbox) SetSubscriptionTypeId(v int32)`

SetSubscriptionTypeId sets SubscriptionTypeId field to given value.


### GetLabel

`func (o *LegalConsentCheckbox) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LegalConsentCheckbox) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LegalConsentCheckbox) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRequired

`func (o *LegalConsentCheckbox) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *LegalConsentCheckbox) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *LegalConsentCheckbox) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


