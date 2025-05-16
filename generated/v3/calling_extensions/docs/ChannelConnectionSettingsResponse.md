# ChannelConnectionSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The timestamp this setting was created | 
**IsReady** | **bool** | If true, this app will be considered to support channel connection | 
**Url** | **string** | The URL to fetch phone numbers available for channel connection | 
**UpdatedAt** | **time.Time** | The timestamp this setting was last updated | 

## Methods

### NewChannelConnectionSettingsResponse

`func NewChannelConnectionSettingsResponse(createdAt time.Time, isReady bool, url string, updatedAt time.Time, ) *ChannelConnectionSettingsResponse`

NewChannelConnectionSettingsResponse instantiates a new ChannelConnectionSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelConnectionSettingsResponseWithDefaults

`func NewChannelConnectionSettingsResponseWithDefaults() *ChannelConnectionSettingsResponse`

NewChannelConnectionSettingsResponseWithDefaults instantiates a new ChannelConnectionSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ChannelConnectionSettingsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChannelConnectionSettingsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChannelConnectionSettingsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsReady

`func (o *ChannelConnectionSettingsResponse) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *ChannelConnectionSettingsResponse) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *ChannelConnectionSettingsResponse) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.


### GetUrl

`func (o *ChannelConnectionSettingsResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChannelConnectionSettingsResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChannelConnectionSettingsResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUpdatedAt

`func (o *ChannelConnectionSettingsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChannelConnectionSettingsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChannelConnectionSettingsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


