# MediaPlayedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaUrl** | Pointer to **string** |  | [optional] 
**ContactId** | **int64** |  | 
**MediaType** | **string** |  | 
**SessionId** | **string** |  | 
**PageId** | Pointer to **int64** |  | [optional] 
**MediaBridgeObjectTypeId** | **string** |  | 
**PageName** | Pointer to **string** |  | [optional] 
**IframeUrl** | Pointer to **string** |  | [optional] 
**PageObjectCoordinates** | Pointer to **string** |  | [optional] 
**OccurredTimestamp** | **int64** |  | 
**ProviderId** | **int32** |  | 
**PortalId** | **int32** |  | 
**PageUrl** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**MediaBridgeId** | **int64** |  | 
**MediaBridgeObjectCoordinates** | **string** |  | 
**MediaName** | **string** |  | 

## Methods

### NewMediaPlayedEvent

`func NewMediaPlayedEvent(contactId int64, mediaType string, sessionId string, mediaBridgeObjectTypeId string, occurredTimestamp int64, providerId int32, portalId int32, state string, mediaBridgeId int64, mediaBridgeObjectCoordinates string, mediaName string, ) *MediaPlayedEvent`

NewMediaPlayedEvent instantiates a new MediaPlayedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaPlayedEventWithDefaults

`func NewMediaPlayedEventWithDefaults() *MediaPlayedEvent`

NewMediaPlayedEventWithDefaults instantiates a new MediaPlayedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaUrl

`func (o *MediaPlayedEvent) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *MediaPlayedEvent) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *MediaPlayedEvent) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *MediaPlayedEvent) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetContactId

`func (o *MediaPlayedEvent) GetContactId() int64`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *MediaPlayedEvent) GetContactIdOk() (*int64, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *MediaPlayedEvent) SetContactId(v int64)`

SetContactId sets ContactId field to given value.


### GetMediaType

`func (o *MediaPlayedEvent) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *MediaPlayedEvent) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *MediaPlayedEvent) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetSessionId

`func (o *MediaPlayedEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *MediaPlayedEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *MediaPlayedEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetPageId

`func (o *MediaPlayedEvent) GetPageId() int64`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *MediaPlayedEvent) GetPageIdOk() (*int64, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *MediaPlayedEvent) SetPageId(v int64)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *MediaPlayedEvent) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetMediaBridgeObjectTypeId

`func (o *MediaPlayedEvent) GetMediaBridgeObjectTypeId() string`

GetMediaBridgeObjectTypeId returns the MediaBridgeObjectTypeId field if non-nil, zero value otherwise.

### GetMediaBridgeObjectTypeIdOk

`func (o *MediaPlayedEvent) GetMediaBridgeObjectTypeIdOk() (*string, bool)`

GetMediaBridgeObjectTypeIdOk returns a tuple with the MediaBridgeObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBridgeObjectTypeId

`func (o *MediaPlayedEvent) SetMediaBridgeObjectTypeId(v string)`

SetMediaBridgeObjectTypeId sets MediaBridgeObjectTypeId field to given value.


### GetPageName

`func (o *MediaPlayedEvent) GetPageName() string`

GetPageName returns the PageName field if non-nil, zero value otherwise.

### GetPageNameOk

`func (o *MediaPlayedEvent) GetPageNameOk() (*string, bool)`

GetPageNameOk returns a tuple with the PageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageName

`func (o *MediaPlayedEvent) SetPageName(v string)`

SetPageName sets PageName field to given value.

### HasPageName

`func (o *MediaPlayedEvent) HasPageName() bool`

HasPageName returns a boolean if a field has been set.

### GetIframeUrl

`func (o *MediaPlayedEvent) GetIframeUrl() string`

GetIframeUrl returns the IframeUrl field if non-nil, zero value otherwise.

### GetIframeUrlOk

`func (o *MediaPlayedEvent) GetIframeUrlOk() (*string, bool)`

GetIframeUrlOk returns a tuple with the IframeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIframeUrl

`func (o *MediaPlayedEvent) SetIframeUrl(v string)`

SetIframeUrl sets IframeUrl field to given value.

### HasIframeUrl

`func (o *MediaPlayedEvent) HasIframeUrl() bool`

HasIframeUrl returns a boolean if a field has been set.

### GetPageObjectCoordinates

`func (o *MediaPlayedEvent) GetPageObjectCoordinates() string`

GetPageObjectCoordinates returns the PageObjectCoordinates field if non-nil, zero value otherwise.

### GetPageObjectCoordinatesOk

`func (o *MediaPlayedEvent) GetPageObjectCoordinatesOk() (*string, bool)`

GetPageObjectCoordinatesOk returns a tuple with the PageObjectCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageObjectCoordinates

`func (o *MediaPlayedEvent) SetPageObjectCoordinates(v string)`

SetPageObjectCoordinates sets PageObjectCoordinates field to given value.

### HasPageObjectCoordinates

`func (o *MediaPlayedEvent) HasPageObjectCoordinates() bool`

HasPageObjectCoordinates returns a boolean if a field has been set.

### GetOccurredTimestamp

`func (o *MediaPlayedEvent) GetOccurredTimestamp() int64`

GetOccurredTimestamp returns the OccurredTimestamp field if non-nil, zero value otherwise.

### GetOccurredTimestampOk

`func (o *MediaPlayedEvent) GetOccurredTimestampOk() (*int64, bool)`

GetOccurredTimestampOk returns a tuple with the OccurredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredTimestamp

`func (o *MediaPlayedEvent) SetOccurredTimestamp(v int64)`

SetOccurredTimestamp sets OccurredTimestamp field to given value.


### GetProviderId

`func (o *MediaPlayedEvent) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *MediaPlayedEvent) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *MediaPlayedEvent) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetPortalId

`func (o *MediaPlayedEvent) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *MediaPlayedEvent) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *MediaPlayedEvent) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.


### GetPageUrl

`func (o *MediaPlayedEvent) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *MediaPlayedEvent) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *MediaPlayedEvent) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.

### HasPageUrl

`func (o *MediaPlayedEvent) HasPageUrl() bool`

HasPageUrl returns a boolean if a field has been set.

### GetState

`func (o *MediaPlayedEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MediaPlayedEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MediaPlayedEvent) SetState(v string)`

SetState sets State field to given value.


### GetMediaBridgeId

`func (o *MediaPlayedEvent) GetMediaBridgeId() int64`

GetMediaBridgeId returns the MediaBridgeId field if non-nil, zero value otherwise.

### GetMediaBridgeIdOk

`func (o *MediaPlayedEvent) GetMediaBridgeIdOk() (*int64, bool)`

GetMediaBridgeIdOk returns a tuple with the MediaBridgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBridgeId

`func (o *MediaPlayedEvent) SetMediaBridgeId(v int64)`

SetMediaBridgeId sets MediaBridgeId field to given value.


### GetMediaBridgeObjectCoordinates

`func (o *MediaPlayedEvent) GetMediaBridgeObjectCoordinates() string`

GetMediaBridgeObjectCoordinates returns the MediaBridgeObjectCoordinates field if non-nil, zero value otherwise.

### GetMediaBridgeObjectCoordinatesOk

`func (o *MediaPlayedEvent) GetMediaBridgeObjectCoordinatesOk() (*string, bool)`

GetMediaBridgeObjectCoordinatesOk returns a tuple with the MediaBridgeObjectCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBridgeObjectCoordinates

`func (o *MediaPlayedEvent) SetMediaBridgeObjectCoordinates(v string)`

SetMediaBridgeObjectCoordinates sets MediaBridgeObjectCoordinates field to given value.


### GetMediaName

`func (o *MediaPlayedEvent) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *MediaPlayedEvent) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *MediaPlayedEvent) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


