# ExternalMeetingBookingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalendarEventId** | **string** |  | 
**ContactId** | **string** |  | 
**BookingTimezone** | **string** |  | 
**LegalConsentResponses** | [**[]ExternalLegalConsentResponse**](ExternalLegalConsentResponse.md) |  | 
**Subject** | **string** |  | 
**Start** | **time.Time** |  | 
**Locale** | Pointer to **string** |  | [optional] 
**WebConferenceUrl** | Pointer to **string** |  | [optional] 
**Duration** | **int64** |  | 
**GuestEmails** | **[]string** |  | 
**WebConferenceMeetingId** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**IsOffline** | **bool** |  | 
**End** | **time.Time** |  | 
**FormFields** | [**[]ExternalValidatedFormField**](ExternalValidatedFormField.md) |  | 

## Methods

### NewExternalMeetingBookingResponse

`func NewExternalMeetingBookingResponse(calendarEventId string, contactId string, bookingTimezone string, legalConsentResponses []ExternalLegalConsentResponse, subject string, start time.Time, duration int64, guestEmails []string, isOffline bool, end time.Time, formFields []ExternalValidatedFormField, ) *ExternalMeetingBookingResponse`

NewExternalMeetingBookingResponse instantiates a new ExternalMeetingBookingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMeetingBookingResponseWithDefaults

`func NewExternalMeetingBookingResponseWithDefaults() *ExternalMeetingBookingResponse`

NewExternalMeetingBookingResponseWithDefaults instantiates a new ExternalMeetingBookingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalendarEventId

`func (o *ExternalMeetingBookingResponse) GetCalendarEventId() string`

GetCalendarEventId returns the CalendarEventId field if non-nil, zero value otherwise.

### GetCalendarEventIdOk

`func (o *ExternalMeetingBookingResponse) GetCalendarEventIdOk() (*string, bool)`

GetCalendarEventIdOk returns a tuple with the CalendarEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarEventId

`func (o *ExternalMeetingBookingResponse) SetCalendarEventId(v string)`

SetCalendarEventId sets CalendarEventId field to given value.


### GetContactId

`func (o *ExternalMeetingBookingResponse) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ExternalMeetingBookingResponse) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ExternalMeetingBookingResponse) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetBookingTimezone

`func (o *ExternalMeetingBookingResponse) GetBookingTimezone() string`

GetBookingTimezone returns the BookingTimezone field if non-nil, zero value otherwise.

### GetBookingTimezoneOk

`func (o *ExternalMeetingBookingResponse) GetBookingTimezoneOk() (*string, bool)`

GetBookingTimezoneOk returns a tuple with the BookingTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingTimezone

`func (o *ExternalMeetingBookingResponse) SetBookingTimezone(v string)`

SetBookingTimezone sets BookingTimezone field to given value.


### GetLegalConsentResponses

`func (o *ExternalMeetingBookingResponse) GetLegalConsentResponses() []ExternalLegalConsentResponse`

GetLegalConsentResponses returns the LegalConsentResponses field if non-nil, zero value otherwise.

### GetLegalConsentResponsesOk

`func (o *ExternalMeetingBookingResponse) GetLegalConsentResponsesOk() (*[]ExternalLegalConsentResponse, bool)`

GetLegalConsentResponsesOk returns a tuple with the LegalConsentResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalConsentResponses

`func (o *ExternalMeetingBookingResponse) SetLegalConsentResponses(v []ExternalLegalConsentResponse)`

SetLegalConsentResponses sets LegalConsentResponses field to given value.


### GetSubject

`func (o *ExternalMeetingBookingResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ExternalMeetingBookingResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ExternalMeetingBookingResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetStart

`func (o *ExternalMeetingBookingResponse) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ExternalMeetingBookingResponse) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ExternalMeetingBookingResponse) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetLocale

`func (o *ExternalMeetingBookingResponse) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ExternalMeetingBookingResponse) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ExternalMeetingBookingResponse) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ExternalMeetingBookingResponse) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetWebConferenceUrl

`func (o *ExternalMeetingBookingResponse) GetWebConferenceUrl() string`

GetWebConferenceUrl returns the WebConferenceUrl field if non-nil, zero value otherwise.

### GetWebConferenceUrlOk

`func (o *ExternalMeetingBookingResponse) GetWebConferenceUrlOk() (*string, bool)`

GetWebConferenceUrlOk returns a tuple with the WebConferenceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebConferenceUrl

`func (o *ExternalMeetingBookingResponse) SetWebConferenceUrl(v string)`

SetWebConferenceUrl sets WebConferenceUrl field to given value.

### HasWebConferenceUrl

`func (o *ExternalMeetingBookingResponse) HasWebConferenceUrl() bool`

HasWebConferenceUrl returns a boolean if a field has been set.

### GetDuration

`func (o *ExternalMeetingBookingResponse) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ExternalMeetingBookingResponse) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ExternalMeetingBookingResponse) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetGuestEmails

`func (o *ExternalMeetingBookingResponse) GetGuestEmails() []string`

GetGuestEmails returns the GuestEmails field if non-nil, zero value otherwise.

### GetGuestEmailsOk

`func (o *ExternalMeetingBookingResponse) GetGuestEmailsOk() (*[]string, bool)`

GetGuestEmailsOk returns a tuple with the GuestEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestEmails

`func (o *ExternalMeetingBookingResponse) SetGuestEmails(v []string)`

SetGuestEmails sets GuestEmails field to given value.


### GetWebConferenceMeetingId

`func (o *ExternalMeetingBookingResponse) GetWebConferenceMeetingId() string`

GetWebConferenceMeetingId returns the WebConferenceMeetingId field if non-nil, zero value otherwise.

### GetWebConferenceMeetingIdOk

`func (o *ExternalMeetingBookingResponse) GetWebConferenceMeetingIdOk() (*string, bool)`

GetWebConferenceMeetingIdOk returns a tuple with the WebConferenceMeetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebConferenceMeetingId

`func (o *ExternalMeetingBookingResponse) SetWebConferenceMeetingId(v string)`

SetWebConferenceMeetingId sets WebConferenceMeetingId field to given value.

### HasWebConferenceMeetingId

`func (o *ExternalMeetingBookingResponse) HasWebConferenceMeetingId() bool`

HasWebConferenceMeetingId returns a boolean if a field has been set.

### GetLocation

`func (o *ExternalMeetingBookingResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExternalMeetingBookingResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExternalMeetingBookingResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExternalMeetingBookingResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetIsOffline

`func (o *ExternalMeetingBookingResponse) GetIsOffline() bool`

GetIsOffline returns the IsOffline field if non-nil, zero value otherwise.

### GetIsOfflineOk

`func (o *ExternalMeetingBookingResponse) GetIsOfflineOk() (*bool, bool)`

GetIsOfflineOk returns a tuple with the IsOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOffline

`func (o *ExternalMeetingBookingResponse) SetIsOffline(v bool)`

SetIsOffline sets IsOffline field to given value.


### GetEnd

`func (o *ExternalMeetingBookingResponse) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ExternalMeetingBookingResponse) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ExternalMeetingBookingResponse) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetFormFields

`func (o *ExternalMeetingBookingResponse) GetFormFields() []ExternalValidatedFormField`

GetFormFields returns the FormFields field if non-nil, zero value otherwise.

### GetFormFieldsOk

`func (o *ExternalMeetingBookingResponse) GetFormFieldsOk() (*[]ExternalValidatedFormField, bool)`

GetFormFieldsOk returns a tuple with the FormFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFields

`func (o *ExternalMeetingBookingResponse) SetFormFields(v []ExternalValidatedFormField)`

SetFormFields sets FormFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


