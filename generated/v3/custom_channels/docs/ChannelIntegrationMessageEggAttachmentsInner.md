# ChannelIntegrationMessageEggAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUsageType** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "FILE"]
**FileId** | **int64** |  | 
**Address** | Pointer to **string** |  | [optional] 
**Latitude** | **float32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Longitude** | **float32** |  | 
**ContactProfile** | [**ContactProfile**](ContactProfile.md) |  | 
**Text** | Pointer to **string** |  | [optional] 
**QuickReplies** | [**[]QuickReply**](QuickReply.md) |  | 

## Methods

### NewChannelIntegrationMessageEggAttachmentsInner

`func NewChannelIntegrationMessageEggAttachmentsInner(type_ string, fileId int64, latitude float32, longitude float32, contactProfile ContactProfile, quickReplies []QuickReply, ) *ChannelIntegrationMessageEggAttachmentsInner`

NewChannelIntegrationMessageEggAttachmentsInner instantiates a new ChannelIntegrationMessageEggAttachmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelIntegrationMessageEggAttachmentsInnerWithDefaults

`func NewChannelIntegrationMessageEggAttachmentsInnerWithDefaults() *ChannelIntegrationMessageEggAttachmentsInner`

NewChannelIntegrationMessageEggAttachmentsInnerWithDefaults instantiates a new ChannelIntegrationMessageEggAttachmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUsageType

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetFileUsageType() string`

GetFileUsageType returns the FileUsageType field if non-nil, zero value otherwise.

### GetFileUsageTypeOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetFileUsageTypeOk() (*string, bool)`

GetFileUsageTypeOk returns a tuple with the FileUsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUsageType

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetFileUsageType(v string)`

SetFileUsageType sets FileUsageType field to given value.

### HasFileUsageType

`func (o *ChannelIntegrationMessageEggAttachmentsInner) HasFileUsageType() bool`

HasFileUsageType returns a boolean if a field has been set.

### GetType

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetType(v string)`

SetType sets Type field to given value.


### GetFileId

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetFileId(v int64)`

SetFileId sets FileId field to given value.


### GetAddress

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ChannelIntegrationMessageEggAttachmentsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetName

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelIntegrationMessageEggAttachmentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChannelIntegrationMessageEggAttachmentsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLongitude

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetContactProfile

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetContactProfile() ContactProfile`

GetContactProfile returns the ContactProfile field if non-nil, zero value otherwise.

### GetContactProfileOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetContactProfileOk() (*ContactProfile, bool)`

GetContactProfileOk returns a tuple with the ContactProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactProfile

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetContactProfile(v ContactProfile)`

SetContactProfile sets ContactProfile field to given value.


### GetText

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ChannelIntegrationMessageEggAttachmentsInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetQuickReplies

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetQuickReplies() []QuickReply`

GetQuickReplies returns the QuickReplies field if non-nil, zero value otherwise.

### GetQuickRepliesOk

`func (o *ChannelIntegrationMessageEggAttachmentsInner) GetQuickRepliesOk() (*[]QuickReply, bool)`

GetQuickRepliesOk returns a tuple with the QuickReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickReplies

`func (o *ChannelIntegrationMessageEggAttachmentsInner) SetQuickReplies(v []QuickReply)`

SetQuickReplies sets QuickReplies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


