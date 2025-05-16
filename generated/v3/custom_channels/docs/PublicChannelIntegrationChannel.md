# PublicChannelIntegrationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelDescription** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Capabilities** | **map[string]map[string]interface{}** |  | 
**ChannelAccountConnectionRedirectUrl** | Pointer to **string** |  | [optional] 
**ChannelLogoUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Id** | **string** |  | 
**WebhookUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicChannelIntegrationChannel

`func NewPublicChannelIntegrationChannel(createdAt time.Time, capabilities map[string]map[string]interface{}, name string, id string, ) *PublicChannelIntegrationChannel`

NewPublicChannelIntegrationChannel instantiates a new PublicChannelIntegrationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicChannelIntegrationChannelWithDefaults

`func NewPublicChannelIntegrationChannelWithDefaults() *PublicChannelIntegrationChannel`

NewPublicChannelIntegrationChannelWithDefaults instantiates a new PublicChannelIntegrationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelDescription

`func (o *PublicChannelIntegrationChannel) GetChannelDescription() string`

GetChannelDescription returns the ChannelDescription field if non-nil, zero value otherwise.

### GetChannelDescriptionOk

`func (o *PublicChannelIntegrationChannel) GetChannelDescriptionOk() (*string, bool)`

GetChannelDescriptionOk returns a tuple with the ChannelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDescription

`func (o *PublicChannelIntegrationChannel) SetChannelDescription(v string)`

SetChannelDescription sets ChannelDescription field to given value.

### HasChannelDescription

`func (o *PublicChannelIntegrationChannel) HasChannelDescription() bool`

HasChannelDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicChannelIntegrationChannel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicChannelIntegrationChannel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicChannelIntegrationChannel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCapabilities

`func (o *PublicChannelIntegrationChannel) GetCapabilities() map[string]map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PublicChannelIntegrationChannel) GetCapabilitiesOk() (*map[string]map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PublicChannelIntegrationChannel) SetCapabilities(v map[string]map[string]interface{})`

SetCapabilities sets Capabilities field to given value.


### GetChannelAccountConnectionRedirectUrl

`func (o *PublicChannelIntegrationChannel) GetChannelAccountConnectionRedirectUrl() string`

GetChannelAccountConnectionRedirectUrl returns the ChannelAccountConnectionRedirectUrl field if non-nil, zero value otherwise.

### GetChannelAccountConnectionRedirectUrlOk

`func (o *PublicChannelIntegrationChannel) GetChannelAccountConnectionRedirectUrlOk() (*string, bool)`

GetChannelAccountConnectionRedirectUrlOk returns a tuple with the ChannelAccountConnectionRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAccountConnectionRedirectUrl

`func (o *PublicChannelIntegrationChannel) SetChannelAccountConnectionRedirectUrl(v string)`

SetChannelAccountConnectionRedirectUrl sets ChannelAccountConnectionRedirectUrl field to given value.

### HasChannelAccountConnectionRedirectUrl

`func (o *PublicChannelIntegrationChannel) HasChannelAccountConnectionRedirectUrl() bool`

HasChannelAccountConnectionRedirectUrl returns a boolean if a field has been set.

### GetChannelLogoUrl

`func (o *PublicChannelIntegrationChannel) GetChannelLogoUrl() string`

GetChannelLogoUrl returns the ChannelLogoUrl field if non-nil, zero value otherwise.

### GetChannelLogoUrlOk

`func (o *PublicChannelIntegrationChannel) GetChannelLogoUrlOk() (*string, bool)`

GetChannelLogoUrlOk returns a tuple with the ChannelLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLogoUrl

`func (o *PublicChannelIntegrationChannel) SetChannelLogoUrl(v string)`

SetChannelLogoUrl sets ChannelLogoUrl field to given value.

### HasChannelLogoUrl

`func (o *PublicChannelIntegrationChannel) HasChannelLogoUrl() bool`

HasChannelLogoUrl returns a boolean if a field has been set.

### GetName

`func (o *PublicChannelIntegrationChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicChannelIntegrationChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicChannelIntegrationChannel) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *PublicChannelIntegrationChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicChannelIntegrationChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicChannelIntegrationChannel) SetId(v string)`

SetId sets Id field to given value.


### GetWebhookUrl

`func (o *PublicChannelIntegrationChannel) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *PublicChannelIntegrationChannel) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *PublicChannelIntegrationChannel) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *PublicChannelIntegrationChannel) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


