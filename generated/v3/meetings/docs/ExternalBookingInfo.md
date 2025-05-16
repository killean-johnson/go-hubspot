# ExternalBookingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | **string** |  | 
**CustomParams** | [**ExternalMeetingsLinkSettings**](ExternalMeetingsLinkSettings.md) |  | 
**LinkAvailability** | Pointer to [**ExternalLinkAvailability**](ExternalLinkAvailability.md) |  | [optional] 
**BrandingMetadata** | Pointer to [**ExternalBrandingMetadata**](ExternalBrandingMetadata.md) |  | [optional] 
**IsOffline** | **bool** |  | 
**LinkType** | **string** |  | 
**AllUsersBusyTimes** | [**[]ExternalUserBusyTimes**](ExternalUserBusyTimes.md) |  | 

## Methods

### NewExternalBookingInfo

`func NewExternalBookingInfo(linkId string, customParams ExternalMeetingsLinkSettings, isOffline bool, linkType string, allUsersBusyTimes []ExternalUserBusyTimes, ) *ExternalBookingInfo`

NewExternalBookingInfo instantiates a new ExternalBookingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalBookingInfoWithDefaults

`func NewExternalBookingInfoWithDefaults() *ExternalBookingInfo`

NewExternalBookingInfoWithDefaults instantiates a new ExternalBookingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *ExternalBookingInfo) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *ExternalBookingInfo) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *ExternalBookingInfo) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.


### GetCustomParams

`func (o *ExternalBookingInfo) GetCustomParams() ExternalMeetingsLinkSettings`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *ExternalBookingInfo) GetCustomParamsOk() (*ExternalMeetingsLinkSettings, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *ExternalBookingInfo) SetCustomParams(v ExternalMeetingsLinkSettings)`

SetCustomParams sets CustomParams field to given value.


### GetLinkAvailability

`func (o *ExternalBookingInfo) GetLinkAvailability() ExternalLinkAvailability`

GetLinkAvailability returns the LinkAvailability field if non-nil, zero value otherwise.

### GetLinkAvailabilityOk

`func (o *ExternalBookingInfo) GetLinkAvailabilityOk() (*ExternalLinkAvailability, bool)`

GetLinkAvailabilityOk returns a tuple with the LinkAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAvailability

`func (o *ExternalBookingInfo) SetLinkAvailability(v ExternalLinkAvailability)`

SetLinkAvailability sets LinkAvailability field to given value.

### HasLinkAvailability

`func (o *ExternalBookingInfo) HasLinkAvailability() bool`

HasLinkAvailability returns a boolean if a field has been set.

### GetBrandingMetadata

`func (o *ExternalBookingInfo) GetBrandingMetadata() ExternalBrandingMetadata`

GetBrandingMetadata returns the BrandingMetadata field if non-nil, zero value otherwise.

### GetBrandingMetadataOk

`func (o *ExternalBookingInfo) GetBrandingMetadataOk() (*ExternalBrandingMetadata, bool)`

GetBrandingMetadataOk returns a tuple with the BrandingMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingMetadata

`func (o *ExternalBookingInfo) SetBrandingMetadata(v ExternalBrandingMetadata)`

SetBrandingMetadata sets BrandingMetadata field to given value.

### HasBrandingMetadata

`func (o *ExternalBookingInfo) HasBrandingMetadata() bool`

HasBrandingMetadata returns a boolean if a field has been set.

### GetIsOffline

`func (o *ExternalBookingInfo) GetIsOffline() bool`

GetIsOffline returns the IsOffline field if non-nil, zero value otherwise.

### GetIsOfflineOk

`func (o *ExternalBookingInfo) GetIsOfflineOk() (*bool, bool)`

GetIsOfflineOk returns a tuple with the IsOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOffline

`func (o *ExternalBookingInfo) SetIsOffline(v bool)`

SetIsOffline sets IsOffline field to given value.


### GetLinkType

`func (o *ExternalBookingInfo) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *ExternalBookingInfo) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *ExternalBookingInfo) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.


### GetAllUsersBusyTimes

`func (o *ExternalBookingInfo) GetAllUsersBusyTimes() []ExternalUserBusyTimes`

GetAllUsersBusyTimes returns the AllUsersBusyTimes field if non-nil, zero value otherwise.

### GetAllUsersBusyTimesOk

`func (o *ExternalBookingInfo) GetAllUsersBusyTimesOk() (*[]ExternalUserBusyTimes, bool)`

GetAllUsersBusyTimesOk returns a tuple with the AllUsersBusyTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUsersBusyTimes

`func (o *ExternalBookingInfo) SetAllUsersBusyTimes(v []ExternalUserBusyTimes)`

SetAllUsersBusyTimes sets AllUsersBusyTimes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


