# EmailFieldValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseDefaultBlockList** | **bool** | Whether to block the free email providers. | 
**BlockedEmailDomains** | **[]string** | A list of email domains to block. | 

## Methods

### NewEmailFieldValidation

`func NewEmailFieldValidation(useDefaultBlockList bool, blockedEmailDomains []string, ) *EmailFieldValidation`

NewEmailFieldValidation instantiates a new EmailFieldValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailFieldValidationWithDefaults

`func NewEmailFieldValidationWithDefaults() *EmailFieldValidation`

NewEmailFieldValidationWithDefaults instantiates a new EmailFieldValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseDefaultBlockList

`func (o *EmailFieldValidation) GetUseDefaultBlockList() bool`

GetUseDefaultBlockList returns the UseDefaultBlockList field if non-nil, zero value otherwise.

### GetUseDefaultBlockListOk

`func (o *EmailFieldValidation) GetUseDefaultBlockListOk() (*bool, bool)`

GetUseDefaultBlockListOk returns a tuple with the UseDefaultBlockList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultBlockList

`func (o *EmailFieldValidation) SetUseDefaultBlockList(v bool)`

SetUseDefaultBlockList sets UseDefaultBlockList field to given value.


### GetBlockedEmailDomains

`func (o *EmailFieldValidation) GetBlockedEmailDomains() []string`

GetBlockedEmailDomains returns the BlockedEmailDomains field if non-nil, zero value otherwise.

### GetBlockedEmailDomainsOk

`func (o *EmailFieldValidation) GetBlockedEmailDomainsOk() (*[]string, bool)`

GetBlockedEmailDomainsOk returns a tuple with the BlockedEmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedEmailDomains

`func (o *EmailFieldValidation) SetBlockedEmailDomains(v []string)`

SetBlockedEmailDomains sets BlockedEmailDomains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


