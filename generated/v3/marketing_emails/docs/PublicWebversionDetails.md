# PublicWebversionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**RedirectToPageId** | Pointer to **string** |  | [optional] 
**IsPageRedirected** | Pointer to **bool** |  | [optional] 
**RedirectToUrl** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**PageExpiryEnabled** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPublicWebversionDetails

`func NewPublicWebversionDetails() *PublicWebversionDetails`

NewPublicWebversionDetails instantiates a new PublicWebversionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicWebversionDetailsWithDefaults

`func NewPublicWebversionDetailsWithDefaults() *PublicWebversionDetails`

NewPublicWebversionDetailsWithDefaults instantiates a new PublicWebversionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *PublicWebversionDetails) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PublicWebversionDetails) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PublicWebversionDetails) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PublicWebversionDetails) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRedirectToPageId

`func (o *PublicWebversionDetails) GetRedirectToPageId() string`

GetRedirectToPageId returns the RedirectToPageId field if non-nil, zero value otherwise.

### GetRedirectToPageIdOk

`func (o *PublicWebversionDetails) GetRedirectToPageIdOk() (*string, bool)`

GetRedirectToPageIdOk returns a tuple with the RedirectToPageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToPageId

`func (o *PublicWebversionDetails) SetRedirectToPageId(v string)`

SetRedirectToPageId sets RedirectToPageId field to given value.

### HasRedirectToPageId

`func (o *PublicWebversionDetails) HasRedirectToPageId() bool`

HasRedirectToPageId returns a boolean if a field has been set.

### GetIsPageRedirected

`func (o *PublicWebversionDetails) GetIsPageRedirected() bool`

GetIsPageRedirected returns the IsPageRedirected field if non-nil, zero value otherwise.

### GetIsPageRedirectedOk

`func (o *PublicWebversionDetails) GetIsPageRedirectedOk() (*bool, bool)`

GetIsPageRedirectedOk returns a tuple with the IsPageRedirected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPageRedirected

`func (o *PublicWebversionDetails) SetIsPageRedirected(v bool)`

SetIsPageRedirected sets IsPageRedirected field to given value.

### HasIsPageRedirected

`func (o *PublicWebversionDetails) HasIsPageRedirected() bool`

HasIsPageRedirected returns a boolean if a field has been set.

### GetRedirectToUrl

`func (o *PublicWebversionDetails) GetRedirectToUrl() string`

GetRedirectToUrl returns the RedirectToUrl field if non-nil, zero value otherwise.

### GetRedirectToUrlOk

`func (o *PublicWebversionDetails) GetRedirectToUrlOk() (*string, bool)`

GetRedirectToUrlOk returns a tuple with the RedirectToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToUrl

`func (o *PublicWebversionDetails) SetRedirectToUrl(v string)`

SetRedirectToUrl sets RedirectToUrl field to given value.

### HasRedirectToUrl

`func (o *PublicWebversionDetails) HasRedirectToUrl() bool`

HasRedirectToUrl returns a boolean if a field has been set.

### GetTitle

`func (o *PublicWebversionDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PublicWebversionDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PublicWebversionDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PublicWebversionDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetaDescription

`func (o *PublicWebversionDetails) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *PublicWebversionDetails) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *PublicWebversionDetails) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *PublicWebversionDetails) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetPageExpiryEnabled

`func (o *PublicWebversionDetails) GetPageExpiryEnabled() bool`

GetPageExpiryEnabled returns the PageExpiryEnabled field if non-nil, zero value otherwise.

### GetPageExpiryEnabledOk

`func (o *PublicWebversionDetails) GetPageExpiryEnabledOk() (*bool, bool)`

GetPageExpiryEnabledOk returns a tuple with the PageExpiryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiryEnabled

`func (o *PublicWebversionDetails) SetPageExpiryEnabled(v bool)`

SetPageExpiryEnabled sets PageExpiryEnabled field to given value.

### HasPageExpiryEnabled

`func (o *PublicWebversionDetails) HasPageExpiryEnabled() bool`

HasPageExpiryEnabled returns a boolean if a field has been set.

### GetEnabled

`func (o *PublicWebversionDetails) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PublicWebversionDetails) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PublicWebversionDetails) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PublicWebversionDetails) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSlug

`func (o *PublicWebversionDetails) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PublicWebversionDetails) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PublicWebversionDetails) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PublicWebversionDetails) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUrl

`func (o *PublicWebversionDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PublicWebversionDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PublicWebversionDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PublicWebversionDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PublicWebversionDetails) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PublicWebversionDetails) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PublicWebversionDetails) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PublicWebversionDetails) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


