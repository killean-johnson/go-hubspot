# ChannelConnectionSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsReady** | **bool** | If true, this app will be considered to support channel connection | 
**Url** | **string** | The URL to fetch phone numbers available for channel connection | 

## Methods

### NewChannelConnectionSettingsRequest

`func NewChannelConnectionSettingsRequest(isReady bool, url string, ) *ChannelConnectionSettingsRequest`

NewChannelConnectionSettingsRequest instantiates a new ChannelConnectionSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelConnectionSettingsRequestWithDefaults

`func NewChannelConnectionSettingsRequestWithDefaults() *ChannelConnectionSettingsRequest`

NewChannelConnectionSettingsRequestWithDefaults instantiates a new ChannelConnectionSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsReady

`func (o *ChannelConnectionSettingsRequest) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *ChannelConnectionSettingsRequest) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *ChannelConnectionSettingsRequest) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.


### GetUrl

`func (o *ChannelConnectionSettingsRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChannelConnectionSettingsRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChannelConnectionSettingsRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


