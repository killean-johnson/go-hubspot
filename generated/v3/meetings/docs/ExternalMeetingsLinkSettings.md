# ExternalMeetingsLinkSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUrl** | Pointer to **string** |  | [optional] 
**StartTimeIncrementMinutes** | **string** |  | 
**GuestSettings** | Pointer to [**ExternalGuestSettings**](ExternalGuestSettings.md) |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Availability** | [**map[string]ExternalClosedRange**](ExternalClosedRange.md) |  | 
**Locale** | Pointer to **string** |  | [optional] 
**OwnerPrioritized** | **bool** |  | 
**LegalConsentOptions** | Pointer to [**ExternalLegalConsentOptions**](ExternalLegalConsentOptions.md) |  | [optional] 
**WelcomeScreenInfo** | Pointer to [**ExternalMeetingsWelcomeScreenInfo**](ExternalMeetingsWelcomeScreenInfo.md) |  | [optional] 
**MeetingBufferTime** | **int32** |  | 
**WeeksToAdvertise** | **int32** |  | 
**DisplayInfo** | Pointer to [**ExternalLinkDisplayInfo**](ExternalLinkDisplayInfo.md) |  | [optional] 
**Durations** | **[]int64** |  | 
**Location** | Pointer to **string** |  | [optional] 
**FormFields** | [**[]ExternalLinkFormField**](ExternalLinkFormField.md) |  | 
**LegalConsentEnabled** | **bool** |  | 
**CustomAvailabilityEndDate** | Pointer to **int64** |  | [optional] 
**CustomAvailabilityStartDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewExternalMeetingsLinkSettings

`func NewExternalMeetingsLinkSettings(startTimeIncrementMinutes string, availability map[string]ExternalClosedRange, ownerPrioritized bool, meetingBufferTime int32, weeksToAdvertise int32, durations []int64, formFields []ExternalLinkFormField, legalConsentEnabled bool, ) *ExternalMeetingsLinkSettings`

NewExternalMeetingsLinkSettings instantiates a new ExternalMeetingsLinkSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMeetingsLinkSettingsWithDefaults

`func NewExternalMeetingsLinkSettingsWithDefaults() *ExternalMeetingsLinkSettings`

NewExternalMeetingsLinkSettingsWithDefaults instantiates a new ExternalMeetingsLinkSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUrl

`func (o *ExternalMeetingsLinkSettings) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ExternalMeetingsLinkSettings) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ExternalMeetingsLinkSettings) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ExternalMeetingsLinkSettings) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetStartTimeIncrementMinutes

`func (o *ExternalMeetingsLinkSettings) GetStartTimeIncrementMinutes() string`

GetStartTimeIncrementMinutes returns the StartTimeIncrementMinutes field if non-nil, zero value otherwise.

### GetStartTimeIncrementMinutesOk

`func (o *ExternalMeetingsLinkSettings) GetStartTimeIncrementMinutesOk() (*string, bool)`

GetStartTimeIncrementMinutesOk returns a tuple with the StartTimeIncrementMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeIncrementMinutes

`func (o *ExternalMeetingsLinkSettings) SetStartTimeIncrementMinutes(v string)`

SetStartTimeIncrementMinutes sets StartTimeIncrementMinutes field to given value.


### GetGuestSettings

`func (o *ExternalMeetingsLinkSettings) GetGuestSettings() ExternalGuestSettings`

GetGuestSettings returns the GuestSettings field if non-nil, zero value otherwise.

### GetGuestSettingsOk

`func (o *ExternalMeetingsLinkSettings) GetGuestSettingsOk() (*ExternalGuestSettings, bool)`

GetGuestSettingsOk returns a tuple with the GuestSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSettings

`func (o *ExternalMeetingsLinkSettings) SetGuestSettings(v ExternalGuestSettings)`

SetGuestSettings sets GuestSettings field to given value.

### HasGuestSettings

`func (o *ExternalMeetingsLinkSettings) HasGuestSettings() bool`

HasGuestSettings returns a boolean if a field has been set.

### GetLanguage

`func (o *ExternalMeetingsLinkSettings) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ExternalMeetingsLinkSettings) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ExternalMeetingsLinkSettings) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ExternalMeetingsLinkSettings) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetAvailability

`func (o *ExternalMeetingsLinkSettings) GetAvailability() map[string]ExternalClosedRange`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ExternalMeetingsLinkSettings) GetAvailabilityOk() (*map[string]ExternalClosedRange, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ExternalMeetingsLinkSettings) SetAvailability(v map[string]ExternalClosedRange)`

SetAvailability sets Availability field to given value.


### GetLocale

`func (o *ExternalMeetingsLinkSettings) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ExternalMeetingsLinkSettings) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ExternalMeetingsLinkSettings) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ExternalMeetingsLinkSettings) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetOwnerPrioritized

`func (o *ExternalMeetingsLinkSettings) GetOwnerPrioritized() bool`

GetOwnerPrioritized returns the OwnerPrioritized field if non-nil, zero value otherwise.

### GetOwnerPrioritizedOk

`func (o *ExternalMeetingsLinkSettings) GetOwnerPrioritizedOk() (*bool, bool)`

GetOwnerPrioritizedOk returns a tuple with the OwnerPrioritized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPrioritized

`func (o *ExternalMeetingsLinkSettings) SetOwnerPrioritized(v bool)`

SetOwnerPrioritized sets OwnerPrioritized field to given value.


### GetLegalConsentOptions

`func (o *ExternalMeetingsLinkSettings) GetLegalConsentOptions() ExternalLegalConsentOptions`

GetLegalConsentOptions returns the LegalConsentOptions field if non-nil, zero value otherwise.

### GetLegalConsentOptionsOk

`func (o *ExternalMeetingsLinkSettings) GetLegalConsentOptionsOk() (*ExternalLegalConsentOptions, bool)`

GetLegalConsentOptionsOk returns a tuple with the LegalConsentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalConsentOptions

`func (o *ExternalMeetingsLinkSettings) SetLegalConsentOptions(v ExternalLegalConsentOptions)`

SetLegalConsentOptions sets LegalConsentOptions field to given value.

### HasLegalConsentOptions

`func (o *ExternalMeetingsLinkSettings) HasLegalConsentOptions() bool`

HasLegalConsentOptions returns a boolean if a field has been set.

### GetWelcomeScreenInfo

`func (o *ExternalMeetingsLinkSettings) GetWelcomeScreenInfo() ExternalMeetingsWelcomeScreenInfo`

GetWelcomeScreenInfo returns the WelcomeScreenInfo field if non-nil, zero value otherwise.

### GetWelcomeScreenInfoOk

`func (o *ExternalMeetingsLinkSettings) GetWelcomeScreenInfoOk() (*ExternalMeetingsWelcomeScreenInfo, bool)`

GetWelcomeScreenInfoOk returns a tuple with the WelcomeScreenInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeScreenInfo

`func (o *ExternalMeetingsLinkSettings) SetWelcomeScreenInfo(v ExternalMeetingsWelcomeScreenInfo)`

SetWelcomeScreenInfo sets WelcomeScreenInfo field to given value.

### HasWelcomeScreenInfo

`func (o *ExternalMeetingsLinkSettings) HasWelcomeScreenInfo() bool`

HasWelcomeScreenInfo returns a boolean if a field has been set.

### GetMeetingBufferTime

`func (o *ExternalMeetingsLinkSettings) GetMeetingBufferTime() int32`

GetMeetingBufferTime returns the MeetingBufferTime field if non-nil, zero value otherwise.

### GetMeetingBufferTimeOk

`func (o *ExternalMeetingsLinkSettings) GetMeetingBufferTimeOk() (*int32, bool)`

GetMeetingBufferTimeOk returns a tuple with the MeetingBufferTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingBufferTime

`func (o *ExternalMeetingsLinkSettings) SetMeetingBufferTime(v int32)`

SetMeetingBufferTime sets MeetingBufferTime field to given value.


### GetWeeksToAdvertise

`func (o *ExternalMeetingsLinkSettings) GetWeeksToAdvertise() int32`

GetWeeksToAdvertise returns the WeeksToAdvertise field if non-nil, zero value otherwise.

### GetWeeksToAdvertiseOk

`func (o *ExternalMeetingsLinkSettings) GetWeeksToAdvertiseOk() (*int32, bool)`

GetWeeksToAdvertiseOk returns a tuple with the WeeksToAdvertise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeksToAdvertise

`func (o *ExternalMeetingsLinkSettings) SetWeeksToAdvertise(v int32)`

SetWeeksToAdvertise sets WeeksToAdvertise field to given value.


### GetDisplayInfo

`func (o *ExternalMeetingsLinkSettings) GetDisplayInfo() ExternalLinkDisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *ExternalMeetingsLinkSettings) GetDisplayInfoOk() (*ExternalLinkDisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *ExternalMeetingsLinkSettings) SetDisplayInfo(v ExternalLinkDisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *ExternalMeetingsLinkSettings) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetDurations

`func (o *ExternalMeetingsLinkSettings) GetDurations() []int64`

GetDurations returns the Durations field if non-nil, zero value otherwise.

### GetDurationsOk

`func (o *ExternalMeetingsLinkSettings) GetDurationsOk() (*[]int64, bool)`

GetDurationsOk returns a tuple with the Durations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurations

`func (o *ExternalMeetingsLinkSettings) SetDurations(v []int64)`

SetDurations sets Durations field to given value.


### GetLocation

`func (o *ExternalMeetingsLinkSettings) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExternalMeetingsLinkSettings) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExternalMeetingsLinkSettings) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExternalMeetingsLinkSettings) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetFormFields

`func (o *ExternalMeetingsLinkSettings) GetFormFields() []ExternalLinkFormField`

GetFormFields returns the FormFields field if non-nil, zero value otherwise.

### GetFormFieldsOk

`func (o *ExternalMeetingsLinkSettings) GetFormFieldsOk() (*[]ExternalLinkFormField, bool)`

GetFormFieldsOk returns a tuple with the FormFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFields

`func (o *ExternalMeetingsLinkSettings) SetFormFields(v []ExternalLinkFormField)`

SetFormFields sets FormFields field to given value.


### GetLegalConsentEnabled

`func (o *ExternalMeetingsLinkSettings) GetLegalConsentEnabled() bool`

GetLegalConsentEnabled returns the LegalConsentEnabled field if non-nil, zero value otherwise.

### GetLegalConsentEnabledOk

`func (o *ExternalMeetingsLinkSettings) GetLegalConsentEnabledOk() (*bool, bool)`

GetLegalConsentEnabledOk returns a tuple with the LegalConsentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalConsentEnabled

`func (o *ExternalMeetingsLinkSettings) SetLegalConsentEnabled(v bool)`

SetLegalConsentEnabled sets LegalConsentEnabled field to given value.


### GetCustomAvailabilityEndDate

`func (o *ExternalMeetingsLinkSettings) GetCustomAvailabilityEndDate() int64`

GetCustomAvailabilityEndDate returns the CustomAvailabilityEndDate field if non-nil, zero value otherwise.

### GetCustomAvailabilityEndDateOk

`func (o *ExternalMeetingsLinkSettings) GetCustomAvailabilityEndDateOk() (*int64, bool)`

GetCustomAvailabilityEndDateOk returns a tuple with the CustomAvailabilityEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAvailabilityEndDate

`func (o *ExternalMeetingsLinkSettings) SetCustomAvailabilityEndDate(v int64)`

SetCustomAvailabilityEndDate sets CustomAvailabilityEndDate field to given value.

### HasCustomAvailabilityEndDate

`func (o *ExternalMeetingsLinkSettings) HasCustomAvailabilityEndDate() bool`

HasCustomAvailabilityEndDate returns a boolean if a field has been set.

### GetCustomAvailabilityStartDate

`func (o *ExternalMeetingsLinkSettings) GetCustomAvailabilityStartDate() int64`

GetCustomAvailabilityStartDate returns the CustomAvailabilityStartDate field if non-nil, zero value otherwise.

### GetCustomAvailabilityStartDateOk

`func (o *ExternalMeetingsLinkSettings) GetCustomAvailabilityStartDateOk() (*int64, bool)`

GetCustomAvailabilityStartDateOk returns a tuple with the CustomAvailabilityStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAvailabilityStartDate

`func (o *ExternalMeetingsLinkSettings) SetCustomAvailabilityStartDate(v int64)`

SetCustomAvailabilityStartDate sets CustomAvailabilityStartDate field to given value.

### HasCustomAvailabilityStartDate

`func (o *ExternalMeetingsLinkSettings) HasCustomAvailabilityStartDate() bool`

HasCustomAvailabilityStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


