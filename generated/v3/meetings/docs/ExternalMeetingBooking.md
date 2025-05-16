# ExternalMeetingBooking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int64** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**LikelyAvailableUserIds** | **[]string** |  | 
**Timezone** | Pointer to **string** |  | [optional] 
**LegalConsentResponses** | [**[]ExternalLegalConsentResponse**](ExternalLegalConsentResponse.md) |  | 
**StartTime** | **time.Time** |  | 
**FormFields** | [**[]ExternalBookingFormField**](ExternalBookingFormField.md) |  | 
**Locale** | Pointer to **string** |  | [optional] 
**Slug** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewExternalMeetingBooking

`func NewExternalMeetingBooking(duration int64, firstName string, lastName string, likelyAvailableUserIds []string, legalConsentResponses []ExternalLegalConsentResponse, startTime time.Time, formFields []ExternalBookingFormField, slug string, email string, ) *ExternalMeetingBooking`

NewExternalMeetingBooking instantiates a new ExternalMeetingBooking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMeetingBookingWithDefaults

`func NewExternalMeetingBookingWithDefaults() *ExternalMeetingBooking`

NewExternalMeetingBookingWithDefaults instantiates a new ExternalMeetingBooking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ExternalMeetingBooking) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ExternalMeetingBooking) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ExternalMeetingBooking) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetFirstName

`func (o *ExternalMeetingBooking) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ExternalMeetingBooking) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ExternalMeetingBooking) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ExternalMeetingBooking) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ExternalMeetingBooking) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ExternalMeetingBooking) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetLikelyAvailableUserIds

`func (o *ExternalMeetingBooking) GetLikelyAvailableUserIds() []string`

GetLikelyAvailableUserIds returns the LikelyAvailableUserIds field if non-nil, zero value otherwise.

### GetLikelyAvailableUserIdsOk

`func (o *ExternalMeetingBooking) GetLikelyAvailableUserIdsOk() (*[]string, bool)`

GetLikelyAvailableUserIdsOk returns a tuple with the LikelyAvailableUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikelyAvailableUserIds

`func (o *ExternalMeetingBooking) SetLikelyAvailableUserIds(v []string)`

SetLikelyAvailableUserIds sets LikelyAvailableUserIds field to given value.


### GetTimezone

`func (o *ExternalMeetingBooking) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ExternalMeetingBooking) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ExternalMeetingBooking) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ExternalMeetingBooking) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetLegalConsentResponses

`func (o *ExternalMeetingBooking) GetLegalConsentResponses() []ExternalLegalConsentResponse`

GetLegalConsentResponses returns the LegalConsentResponses field if non-nil, zero value otherwise.

### GetLegalConsentResponsesOk

`func (o *ExternalMeetingBooking) GetLegalConsentResponsesOk() (*[]ExternalLegalConsentResponse, bool)`

GetLegalConsentResponsesOk returns a tuple with the LegalConsentResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalConsentResponses

`func (o *ExternalMeetingBooking) SetLegalConsentResponses(v []ExternalLegalConsentResponse)`

SetLegalConsentResponses sets LegalConsentResponses field to given value.


### GetStartTime

`func (o *ExternalMeetingBooking) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ExternalMeetingBooking) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ExternalMeetingBooking) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetFormFields

`func (o *ExternalMeetingBooking) GetFormFields() []ExternalBookingFormField`

GetFormFields returns the FormFields field if non-nil, zero value otherwise.

### GetFormFieldsOk

`func (o *ExternalMeetingBooking) GetFormFieldsOk() (*[]ExternalBookingFormField, bool)`

GetFormFieldsOk returns a tuple with the FormFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFields

`func (o *ExternalMeetingBooking) SetFormFields(v []ExternalBookingFormField)`

SetFormFields sets FormFields field to given value.


### GetLocale

`func (o *ExternalMeetingBooking) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ExternalMeetingBooking) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ExternalMeetingBooking) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ExternalMeetingBooking) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetSlug

`func (o *ExternalMeetingBooking) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ExternalMeetingBooking) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ExternalMeetingBooking) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEmail

`func (o *ExternalMeetingBooking) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ExternalMeetingBooking) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ExternalMeetingBooking) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


