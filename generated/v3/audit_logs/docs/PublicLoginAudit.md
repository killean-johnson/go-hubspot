# PublicLoginAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionCode** | Pointer to **string** | The approximate region code of the login. | [optional] 
**CountryCode** | Pointer to **string** | The approximate country code of the login. | [optional] 
**IpAddress** | Pointer to **string** | IP address where the activity originated. | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** | Information about the device used for logging in. | [optional] 
**Id** | **string** | The login activity&#39;s unique ID. | 
**LoginAt** | **time.Time** | The time the login took place. | 
**UserId** | Pointer to **int32** | The user&#39;s unique ID. | [optional] 
**Email** | Pointer to **string** | Email address of the user associated with the login. | [optional] 
**LoginSucceeded** | **bool** | Whether the login was successful or not. | 

## Methods

### NewPublicLoginAudit

`func NewPublicLoginAudit(id string, loginAt time.Time, loginSucceeded bool, ) *PublicLoginAudit`

NewPublicLoginAudit instantiates a new PublicLoginAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicLoginAuditWithDefaults

`func NewPublicLoginAuditWithDefaults() *PublicLoginAudit`

NewPublicLoginAuditWithDefaults instantiates a new PublicLoginAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionCode

`func (o *PublicLoginAudit) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *PublicLoginAudit) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *PublicLoginAudit) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *PublicLoginAudit) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *PublicLoginAudit) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PublicLoginAudit) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PublicLoginAudit) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PublicLoginAudit) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetIpAddress

`func (o *PublicLoginAudit) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PublicLoginAudit) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PublicLoginAudit) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PublicLoginAudit) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLocation

`func (o *PublicLoginAudit) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PublicLoginAudit) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PublicLoginAudit) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PublicLoginAudit) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUserAgent

`func (o *PublicLoginAudit) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *PublicLoginAudit) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *PublicLoginAudit) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *PublicLoginAudit) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetId

`func (o *PublicLoginAudit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicLoginAudit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicLoginAudit) SetId(v string)`

SetId sets Id field to given value.


### GetLoginAt

`func (o *PublicLoginAudit) GetLoginAt() time.Time`

GetLoginAt returns the LoginAt field if non-nil, zero value otherwise.

### GetLoginAtOk

`func (o *PublicLoginAudit) GetLoginAtOk() (*time.Time, bool)`

GetLoginAtOk returns a tuple with the LoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAt

`func (o *PublicLoginAudit) SetLoginAt(v time.Time)`

SetLoginAt sets LoginAt field to given value.


### GetUserId

`func (o *PublicLoginAudit) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicLoginAudit) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicLoginAudit) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PublicLoginAudit) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmail

`func (o *PublicLoginAudit) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PublicLoginAudit) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PublicLoginAudit) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PublicLoginAudit) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLoginSucceeded

`func (o *PublicLoginAudit) GetLoginSucceeded() bool`

GetLoginSucceeded returns the LoginSucceeded field if non-nil, zero value otherwise.

### GetLoginSucceededOk

`func (o *PublicLoginAudit) GetLoginSucceededOk() (*bool, bool)`

GetLoginSucceededOk returns a tuple with the LoginSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginSucceeded

`func (o *PublicLoginAudit) SetLoginSucceeded(v bool)`

SetLoginSucceeded sets LoginSucceeded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


