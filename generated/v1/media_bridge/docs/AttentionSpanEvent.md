# AttentionSpanEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaUrl** | Pointer to **string** |  | [optional] 
**ContactId** | **int64** |  | 
**MediaType** | **string** |  | 
**SessionId** | **string** |  | 
**RawData** | Pointer to **string** |  | [optional] 
**PageId** | Pointer to **int64** |  | [optional] 
**MediaBridgeObjectTypeId** | **string** |  | 
**PageName** | Pointer to **string** |  | [optional] 
**PageObjectCoordinates** | Pointer to **string** |  | [optional] 
**OccurredTimestamp** | **int64** |  | 
**ProviderId** | **int32** |  | 
**PortalId** | **int32** |  | 
**TotalPercentPlayed** | **float32** |  | 
**TotalSecondsPlayed** | Pointer to **int32** |  | [optional] 
**PageUrl** | Pointer to **string** |  | [optional] 
**MediaBridgeId** | **int64** |  | 
**PercentRange** | **string** |  | 
**MediaBridgeObjectCoordinates** | **string** |  | 
**MediaName** | **string** |  | 

## Methods

### NewAttentionSpanEvent

`func NewAttentionSpanEvent(contactId int64, mediaType string, sessionId string, mediaBridgeObjectTypeId string, occurredTimestamp int64, providerId int32, portalId int32, totalPercentPlayed float32, mediaBridgeId int64, percentRange string, mediaBridgeObjectCoordinates string, mediaName string, ) *AttentionSpanEvent`

NewAttentionSpanEvent instantiates a new AttentionSpanEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttentionSpanEventWithDefaults

`func NewAttentionSpanEventWithDefaults() *AttentionSpanEvent`

NewAttentionSpanEventWithDefaults instantiates a new AttentionSpanEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaUrl

`func (o *AttentionSpanEvent) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *AttentionSpanEvent) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *AttentionSpanEvent) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *AttentionSpanEvent) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetContactId

`func (o *AttentionSpanEvent) GetContactId() int64`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *AttentionSpanEvent) GetContactIdOk() (*int64, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *AttentionSpanEvent) SetContactId(v int64)`

SetContactId sets ContactId field to given value.


### GetMediaType

`func (o *AttentionSpanEvent) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *AttentionSpanEvent) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *AttentionSpanEvent) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetSessionId

`func (o *AttentionSpanEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AttentionSpanEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AttentionSpanEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetRawData

`func (o *AttentionSpanEvent) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *AttentionSpanEvent) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *AttentionSpanEvent) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *AttentionSpanEvent) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetPageId

`func (o *AttentionSpanEvent) GetPageId() int64`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *AttentionSpanEvent) GetPageIdOk() (*int64, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *AttentionSpanEvent) SetPageId(v int64)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *AttentionSpanEvent) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetMediaBridgeObjectTypeId

`func (o *AttentionSpanEvent) GetMediaBridgeObjectTypeId() string`

GetMediaBridgeObjectTypeId returns the MediaBridgeObjectTypeId field if non-nil, zero value otherwise.

### GetMediaBridgeObjectTypeIdOk

`func (o *AttentionSpanEvent) GetMediaBridgeObjectTypeIdOk() (*string, bool)`

GetMediaBridgeObjectTypeIdOk returns a tuple with the MediaBridgeObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBridgeObjectTypeId

`func (o *AttentionSpanEvent) SetMediaBridgeObjectTypeId(v string)`

SetMediaBridgeObjectTypeId sets MediaBridgeObjectTypeId field to given value.


### GetPageName

`func (o *AttentionSpanEvent) GetPageName() string`

GetPageName returns the PageName field if non-nil, zero value otherwise.

### GetPageNameOk

`func (o *AttentionSpanEvent) GetPageNameOk() (*string, bool)`

GetPageNameOk returns a tuple with the PageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageName

`func (o *AttentionSpanEvent) SetPageName(v string)`

SetPageName sets PageName field to given value.

### HasPageName

`func (o *AttentionSpanEvent) HasPageName() bool`

HasPageName returns a boolean if a field has been set.

### GetPageObjectCoordinates

`func (o *AttentionSpanEvent) GetPageObjectCoordinates() string`

GetPageObjectCoordinates returns the PageObjectCoordinates field if non-nil, zero value otherwise.

### GetPageObjectCoordinatesOk

`func (o *AttentionSpanEvent) GetPageObjectCoordinatesOk() (*string, bool)`

GetPageObjectCoordinatesOk returns a tuple with the PageObjectCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageObjectCoordinates

`func (o *AttentionSpanEvent) SetPageObjectCoordinates(v string)`

SetPageObjectCoordinates sets PageObjectCoordinates field to given value.

### HasPageObjectCoordinates

`func (o *AttentionSpanEvent) HasPageObjectCoordinates() bool`

HasPageObjectCoordinates returns a boolean if a field has been set.

### GetOccurredTimestamp

`func (o *AttentionSpanEvent) GetOccurredTimestamp() int64`

GetOccurredTimestamp returns the OccurredTimestamp field if non-nil, zero value otherwise.

### GetOccurredTimestampOk

`func (o *AttentionSpanEvent) GetOccurredTimestampOk() (*int64, bool)`

GetOccurredTimestampOk returns a tuple with the OccurredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredTimestamp

`func (o *AttentionSpanEvent) SetOccurredTimestamp(v int64)`

SetOccurredTimestamp sets OccurredTimestamp field to given value.


### GetProviderId

`func (o *AttentionSpanEvent) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *AttentionSpanEvent) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *AttentionSpanEvent) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetPortalId

`func (o *AttentionSpanEvent) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *AttentionSpanEvent) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *AttentionSpanEvent) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.


### GetTotalPercentPlayed

`func (o *AttentionSpanEvent) GetTotalPercentPlayed() float32`

GetTotalPercentPlayed returns the TotalPercentPlayed field if non-nil, zero value otherwise.

### GetTotalPercentPlayedOk

`func (o *AttentionSpanEvent) GetTotalPercentPlayedOk() (*float32, bool)`

GetTotalPercentPlayedOk returns a tuple with the TotalPercentPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPercentPlayed

`func (o *AttentionSpanEvent) SetTotalPercentPlayed(v float32)`

SetTotalPercentPlayed sets TotalPercentPlayed field to given value.


### GetTotalSecondsPlayed

`func (o *AttentionSpanEvent) GetTotalSecondsPlayed() int32`

GetTotalSecondsPlayed returns the TotalSecondsPlayed field if non-nil, zero value otherwise.

### GetTotalSecondsPlayedOk

`func (o *AttentionSpanEvent) GetTotalSecondsPlayedOk() (*int32, bool)`

GetTotalSecondsPlayedOk returns a tuple with the TotalSecondsPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSecondsPlayed

`func (o *AttentionSpanEvent) SetTotalSecondsPlayed(v int32)`

SetTotalSecondsPlayed sets TotalSecondsPlayed field to given value.

### HasTotalSecondsPlayed

`func (o *AttentionSpanEvent) HasTotalSecondsPlayed() bool`

HasTotalSecondsPlayed returns a boolean if a field has been set.

### GetPageUrl

`func (o *AttentionSpanEvent) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *AttentionSpanEvent) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *AttentionSpanEvent) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.

### HasPageUrl

`func (o *AttentionSpanEvent) HasPageUrl() bool`

HasPageUrl returns a boolean if a field has been set.

### GetMediaBridgeId

`func (o *AttentionSpanEvent) GetMediaBridgeId() int64`

GetMediaBridgeId returns the MediaBridgeId field if non-nil, zero value otherwise.

### GetMediaBridgeIdOk

`func (o *AttentionSpanEvent) GetMediaBridgeIdOk() (*int64, bool)`

GetMediaBridgeIdOk returns a tuple with the MediaBridgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBridgeId

`func (o *AttentionSpanEvent) SetMediaBridgeId(v int64)`

SetMediaBridgeId sets MediaBridgeId field to given value.


### GetPercentRange

`func (o *AttentionSpanEvent) GetPercentRange() string`

GetPercentRange returns the PercentRange field if non-nil, zero value otherwise.

### GetPercentRangeOk

`func (o *AttentionSpanEvent) GetPercentRangeOk() (*string, bool)`

GetPercentRangeOk returns a tuple with the PercentRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentRange

`func (o *AttentionSpanEvent) SetPercentRange(v string)`

SetPercentRange sets PercentRange field to given value.


### GetMediaBridgeObjectCoordinates

`func (o *AttentionSpanEvent) GetMediaBridgeObjectCoordinates() string`

GetMediaBridgeObjectCoordinates returns the MediaBridgeObjectCoordinates field if non-nil, zero value otherwise.

### GetMediaBridgeObjectCoordinatesOk

`func (o *AttentionSpanEvent) GetMediaBridgeObjectCoordinatesOk() (*string, bool)`

GetMediaBridgeObjectCoordinatesOk returns a tuple with the MediaBridgeObjectCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBridgeObjectCoordinates

`func (o *AttentionSpanEvent) SetMediaBridgeObjectCoordinates(v string)`

SetMediaBridgeObjectCoordinates sets MediaBridgeObjectCoordinates field to given value.


### GetMediaName

`func (o *AttentionSpanEvent) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *AttentionSpanEvent) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *AttentionSpanEvent) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


