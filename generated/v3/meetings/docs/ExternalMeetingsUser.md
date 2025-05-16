# ExternalMeetingsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSalesStarter** | **bool** |  | 
**CalendarProvider** | **string** |  | 
**Id** | **string** |  | 
**UserId** | **string** |  | 
**UserProfile** | [**ExternalUserProfile**](ExternalUserProfile.md) |  | 

## Methods

### NewExternalMeetingsUser

`func NewExternalMeetingsUser(isSalesStarter bool, calendarProvider string, id string, userId string, userProfile ExternalUserProfile, ) *ExternalMeetingsUser`

NewExternalMeetingsUser instantiates a new ExternalMeetingsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMeetingsUserWithDefaults

`func NewExternalMeetingsUserWithDefaults() *ExternalMeetingsUser`

NewExternalMeetingsUserWithDefaults instantiates a new ExternalMeetingsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSalesStarter

`func (o *ExternalMeetingsUser) GetIsSalesStarter() bool`

GetIsSalesStarter returns the IsSalesStarter field if non-nil, zero value otherwise.

### GetIsSalesStarterOk

`func (o *ExternalMeetingsUser) GetIsSalesStarterOk() (*bool, bool)`

GetIsSalesStarterOk returns a tuple with the IsSalesStarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSalesStarter

`func (o *ExternalMeetingsUser) SetIsSalesStarter(v bool)`

SetIsSalesStarter sets IsSalesStarter field to given value.


### GetCalendarProvider

`func (o *ExternalMeetingsUser) GetCalendarProvider() string`

GetCalendarProvider returns the CalendarProvider field if non-nil, zero value otherwise.

### GetCalendarProviderOk

`func (o *ExternalMeetingsUser) GetCalendarProviderOk() (*string, bool)`

GetCalendarProviderOk returns a tuple with the CalendarProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarProvider

`func (o *ExternalMeetingsUser) SetCalendarProvider(v string)`

SetCalendarProvider sets CalendarProvider field to given value.


### GetId

`func (o *ExternalMeetingsUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalMeetingsUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalMeetingsUser) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ExternalMeetingsUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalMeetingsUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalMeetingsUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserProfile

`func (o *ExternalMeetingsUser) GetUserProfile() ExternalUserProfile`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *ExternalMeetingsUser) GetUserProfileOk() (*ExternalUserProfile, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *ExternalMeetingsUser) SetUserProfile(v ExternalUserProfile)`

SetUserProfile sets UserProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


