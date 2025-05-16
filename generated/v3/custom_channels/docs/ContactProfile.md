# ContactProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | [**[]ContactEmail**](ContactEmail.md) |  | 
**Addresses** | [**[]ContactAddress**](ContactAddress.md) |  | 
**Urls** | [**[]ContactUrl**](ContactUrl.md) |  | 
**Org** | Pointer to [**ContactOrg**](ContactOrg.md) |  | [optional] 
**Name** | Pointer to [**ContactName**](ContactName.md) |  | [optional] 
**Phones** | [**[]ContactPhone**](ContactPhone.md) |  | 

## Methods

### NewContactProfile

`func NewContactProfile(emails []ContactEmail, addresses []ContactAddress, urls []ContactUrl, phones []ContactPhone, ) *ContactProfile`

NewContactProfile instantiates a new ContactProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactProfileWithDefaults

`func NewContactProfileWithDefaults() *ContactProfile`

NewContactProfileWithDefaults instantiates a new ContactProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *ContactProfile) GetEmails() []ContactEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ContactProfile) GetEmailsOk() (*[]ContactEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ContactProfile) SetEmails(v []ContactEmail)`

SetEmails sets Emails field to given value.


### GetAddresses

`func (o *ContactProfile) GetAddresses() []ContactAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ContactProfile) GetAddressesOk() (*[]ContactAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ContactProfile) SetAddresses(v []ContactAddress)`

SetAddresses sets Addresses field to given value.


### GetUrls

`func (o *ContactProfile) GetUrls() []ContactUrl`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *ContactProfile) GetUrlsOk() (*[]ContactUrl, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *ContactProfile) SetUrls(v []ContactUrl)`

SetUrls sets Urls field to given value.


### GetOrg

`func (o *ContactProfile) GetOrg() ContactOrg`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ContactProfile) GetOrgOk() (*ContactOrg, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ContactProfile) SetOrg(v ContactOrg)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ContactProfile) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetName

`func (o *ContactProfile) GetName() ContactName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactProfile) GetNameOk() (*ContactName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactProfile) SetName(v ContactName)`

SetName sets Name field to given value.

### HasName

`func (o *ContactProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhones

`func (o *ContactProfile) GetPhones() []ContactPhone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *ContactProfile) GetPhonesOk() (*[]ContactPhone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *ContactProfile) SetPhones(v []ContactPhone)`

SetPhones sets Phones field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


