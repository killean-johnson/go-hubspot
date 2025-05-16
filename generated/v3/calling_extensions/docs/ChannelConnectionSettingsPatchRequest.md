# ChannelConnectionSettingsPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsReady** | Pointer to **bool** | If true, this app will be considered to support channel connection | [optional] 
**Url** | Pointer to **string** | The URL to fetch phone numbers available for channel connection | [optional] 

## Methods

### NewChannelConnectionSettingsPatchRequest

`func NewChannelConnectionSettingsPatchRequest() *ChannelConnectionSettingsPatchRequest`

NewChannelConnectionSettingsPatchRequest instantiates a new ChannelConnectionSettingsPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelConnectionSettingsPatchRequestWithDefaults

`func NewChannelConnectionSettingsPatchRequestWithDefaults() *ChannelConnectionSettingsPatchRequest`

NewChannelConnectionSettingsPatchRequestWithDefaults instantiates a new ChannelConnectionSettingsPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsReady

`func (o *ChannelConnectionSettingsPatchRequest) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *ChannelConnectionSettingsPatchRequest) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *ChannelConnectionSettingsPatchRequest) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.

### HasIsReady

`func (o *ChannelConnectionSettingsPatchRequest) HasIsReady() bool`

HasIsReady returns a boolean if a field has been set.

### GetUrl

`func (o *ChannelConnectionSettingsPatchRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChannelConnectionSettingsPatchRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChannelConnectionSettingsPatchRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChannelConnectionSettingsPatchRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


