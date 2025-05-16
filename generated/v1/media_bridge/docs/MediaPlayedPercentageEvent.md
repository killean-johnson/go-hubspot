# MediaPlayedPercentageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaUrl** | Pointer to **string** |  | [optional] 
**ContactId** | **int64** |  | 
**MediaType** | **string** |  | 
**SessionId** | **string** |  | 
**PlayedPercent** | **int32** |  | 
**PageId** | Pointer to **int64** |  | [optional] 
**MediaBridgeObjectTypeId** | **string** |  | 
**PageName** | Pointer to **string** |  | [optional] 
**PageObjectCoordinates** | Pointer to **string** |  | [optional] 
**OccurredTimestamp** | **int64** |  | 
**ProviderId** | **int32** |  | 
**PortalId** | **int32** |  | 
**PageUrl** | Pointer to **string** |  | [optional] 
**MediaBridgeId** | **int64** |  | 
**MediaBridgeObjectCoordinates** | **string** |  | 
**MediaName** | **string** |  | 

## Methods

### NewMediaPlayedPercentageEvent

`func NewMediaPlayedPercentageEvent(contactId int64, mediaType string, sessionId string, playedPercent int32, mediaBridgeObjectTypeId string, occurredTimestamp int64, providerId int32, portalId int32, mediaBridgeId int64, mediaBridgeObjectCoordinates string, mediaName string, ) *MediaPlayedPercentageEvent`

NewMediaPlayedPercentageEvent instantiates a new MediaPlayedPercentageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaPlayedPercentageEventWithDefaults

`func NewMediaPlayedPercentageEventWithDefaults() *MediaPlayedPercentageEvent`

NewMediaPlayedPercentageEventWithDefaults instantiates a new MediaPlayedPercentageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaUrl

`func (o *MediaPlayedPercentageEvent) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *MediaPlayedPercentageEvent) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *MediaPlayedPercentageEvent) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *MediaPlayedPercentageEvent) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetContactId

`func (o *MediaPlayedPercentageEvent) GetContactId() int64`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *MediaPlayedPercentageEvent) GetContactIdOk() (*int64, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *MediaPlayedPercentageEvent) SetContactId(v int64)`

SetContactId sets ContactId field to given value.


### GetMediaType

`func (o *MediaPlayedPercentageEvent) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *MediaPlayedPercentageEvent) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *MediaPlayedPercentageEvent) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetSessionId

`func (o *MediaPlayedPercentageEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *MediaPlayedPercentageEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *MediaPlayedPercentageEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetPlayedPercent

`func (o *MediaPlayedPercentageEvent) GetPlayedPercent() int32`

GetPlayedPercent returns the PlayedPercent field if non-nil, zero value otherwise.

### GetPlayedPercentOk

`func (o *MediaPlayedPercentageEvent) GetPlayedPercentOk() (*int32, bool)`

GetPlayedPercentOk returns a tuple with the PlayedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayedPercent

`func (o *MediaPlayedPercentageEvent) SetPlayedPercent(v int32)`

SetPlayedPercent sets PlayedPercent field to given value.


### GetPageId

`func (o *MediaPlayedPercentageEvent) GetPageId() int64`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *MediaPlayedPercentageEvent) GetPageIdOk() (*int64, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *MediaPlayedPercentageEvent) SetPageId(v int64)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *MediaPlayedPercentageEvent) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetMediaBridgeObjectTypeId

`func (o *MediaPlayedPercentageEvent) GetMediaBridgeObjectTypeId() string`

GetMediaBridgeObjectTypeId returns the MediaBridgeObjectTypeId field if non-nil, zero value otherwise.

### GetMediaBridgeObjectTypeIdOk

`func (o *MediaPlayedPercentageEvent) GetMediaBridgeObjectTypeIdOk() (*string, bool)`

GetMediaBridgeObjectTypeIdOk returns a tuple with the MediaBridgeObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBridgeObjectTypeId

`func (o *MediaPlayedPercentageEvent) SetMediaBridgeObjectTypeId(v string)`

SetMediaBridgeObjectTypeId sets MediaBridgeObjectTypeId field to given value.


### GetPageName

`func (o *MediaPlayedPercentageEvent) GetPageName() string`

GetPageName returns the PageName field if non-nil, zero value otherwise.

### GetPageNameOk

`func (o *MediaPlayedPercentageEvent) GetPageNameOk() (*string, bool)`

GetPageNameOk returns a tuple with the PageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageName

`func (o *MediaPlayedPercentageEvent) SetPageName(v string)`

SetPageName sets PageName field to given value.

### HasPageName

`func (o *MediaPlayedPercentageEvent) HasPageName() bool`

HasPageName returns a boolean if a field has been set.

### GetPageObjectCoordinates

`func (o *MediaPlayedPercentageEvent) GetPageObjectCoordinates() string`

GetPageObjectCoordinates returns the PageObjectCoordinates field if non-nil, zero value otherwise.

### GetPageObjectCoordinatesOk

`func (o *MediaPlayedPercentageEvent) GetPageObjectCoordinatesOk() (*string, bool)`

GetPageObjectCoordinatesOk returns a tuple with the PageObjectCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageObjectCoordinates

`func (o *MediaPlayedPercentageEvent) SetPageObjectCoordinates(v string)`

SetPageObjectCoordinates sets PageObjectCoordinates field to given value.

### HasPageObjectCoordinates

`func (o *MediaPlayedPercentageEvent) HasPageObjectCoordinates() bool`

HasPageObjectCoordinates returns a boolean if a field has been set.

### GetOccurredTimestamp

`func (o *MediaPlayedPercentageEvent) GetOccurredTimestamp() int64`

GetOccurredTimestamp returns the OccurredTimestamp field if non-nil, zero value otherwise.

### GetOccurredTimestampOk

`func (o *MediaPlayedPercentageEvent) GetOccurredTimestampOk() (*int64, bool)`

GetOccurredTimestampOk returns a tuple with the OccurredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredTimestamp

`func (o *MediaPlayedPercentageEvent) SetOccurredTimestamp(v int64)`

SetOccurredTimestamp sets OccurredTimestamp field to given value.


### GetProviderId

`func (o *MediaPlayedPercentageEvent) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *MediaPlayedPercentageEvent) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *MediaPlayedPercentageEvent) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetPortalId

`func (o *MediaPlayedPercentageEvent) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *MediaPlayedPercentageEvent) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *MediaPlayedPercentageEvent) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.


### GetPageUrl

`func (o *MediaPlayedPercentageEvent) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *MediaPlayedPercentageEvent) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *MediaPlayedPercentageEvent) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.

### HasPageUrl

`func (o *MediaPlayedPercentageEvent) HasPageUrl() bool`

HasPageUrl returns a boolean if a field has been set.

### GetMediaBridgeId

`func (o *MediaPlayedPercentageEvent) GetMediaBridgeId() int64`

GetMediaBridgeId returns the MediaBridgeId field if non-nil, zero value otherwise.

### GetMediaBridgeIdOk

`func (o *MediaPlayedPercentageEvent) GetMediaBridgeIdOk() (*int64, bool)`

GetMediaBridgeIdOk returns a tuple with the MediaBridgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBridgeId

`func (o *MediaPlayedPercentageEvent) SetMediaBridgeId(v int64)`

SetMediaBridgeId sets MediaBridgeId field to given value.


### GetMediaBridgeObjectCoordinates

`func (o *MediaPlayedPercentageEvent) GetMediaBridgeObjectCoordinates() string`

GetMediaBridgeObjectCoordinates returns the MediaBridgeObjectCoordinates field if non-nil, zero value otherwise.

### GetMediaBridgeObjectCoordinatesOk

`func (o *MediaPlayedPercentageEvent) GetMediaBridgeObjectCoordinatesOk() (*string, bool)`

GetMediaBridgeObjectCoordinatesOk returns a tuple with the MediaBridgeObjectCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBridgeObjectCoordinates

`func (o *MediaPlayedPercentageEvent) SetMediaBridgeObjectCoordinates(v string)`

SetMediaBridgeObjectCoordinates sets MediaBridgeObjectCoordinates field to given value.


### GetMediaName

`func (o *MediaPlayedPercentageEvent) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *MediaPlayedPercentageEvent) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *MediaPlayedPercentageEvent) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


