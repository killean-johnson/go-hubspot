# PublicChannelIntegrationChannelCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelDescription** | Pointer to **string** |  | [optional] 
**Capabilities** | **map[string]map[string]interface{}** |  | 
**ChannelAccountConnectionRedirectUrl** | Pointer to **string** |  | [optional] 
**ChannelLogoUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**WebhookUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicChannelIntegrationChannelCreate

`func NewPublicChannelIntegrationChannelCreate(capabilities map[string]map[string]interface{}, name string, ) *PublicChannelIntegrationChannelCreate`

NewPublicChannelIntegrationChannelCreate instantiates a new PublicChannelIntegrationChannelCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicChannelIntegrationChannelCreateWithDefaults

`func NewPublicChannelIntegrationChannelCreateWithDefaults() *PublicChannelIntegrationChannelCreate`

NewPublicChannelIntegrationChannelCreateWithDefaults instantiates a new PublicChannelIntegrationChannelCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelDescription

`func (o *PublicChannelIntegrationChannelCreate) GetChannelDescription() string`

GetChannelDescription returns the ChannelDescription field if non-nil, zero value otherwise.

### GetChannelDescriptionOk

`func (o *PublicChannelIntegrationChannelCreate) GetChannelDescriptionOk() (*string, bool)`

GetChannelDescriptionOk returns a tuple with the ChannelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDescription

`func (o *PublicChannelIntegrationChannelCreate) SetChannelDescription(v string)`

SetChannelDescription sets ChannelDescription field to given value.

### HasChannelDescription

`func (o *PublicChannelIntegrationChannelCreate) HasChannelDescription() bool`

HasChannelDescription returns a boolean if a field has been set.

### GetCapabilities

`func (o *PublicChannelIntegrationChannelCreate) GetCapabilities() map[string]map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PublicChannelIntegrationChannelCreate) GetCapabilitiesOk() (*map[string]map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PublicChannelIntegrationChannelCreate) SetCapabilities(v map[string]map[string]interface{})`

SetCapabilities sets Capabilities field to given value.


### GetChannelAccountConnectionRedirectUrl

`func (o *PublicChannelIntegrationChannelCreate) GetChannelAccountConnectionRedirectUrl() string`

GetChannelAccountConnectionRedirectUrl returns the ChannelAccountConnectionRedirectUrl field if non-nil, zero value otherwise.

### GetChannelAccountConnectionRedirectUrlOk

`func (o *PublicChannelIntegrationChannelCreate) GetChannelAccountConnectionRedirectUrlOk() (*string, bool)`

GetChannelAccountConnectionRedirectUrlOk returns a tuple with the ChannelAccountConnectionRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAccountConnectionRedirectUrl

`func (o *PublicChannelIntegrationChannelCreate) SetChannelAccountConnectionRedirectUrl(v string)`

SetChannelAccountConnectionRedirectUrl sets ChannelAccountConnectionRedirectUrl field to given value.

### HasChannelAccountConnectionRedirectUrl

`func (o *PublicChannelIntegrationChannelCreate) HasChannelAccountConnectionRedirectUrl() bool`

HasChannelAccountConnectionRedirectUrl returns a boolean if a field has been set.

### GetChannelLogoUrl

`func (o *PublicChannelIntegrationChannelCreate) GetChannelLogoUrl() string`

GetChannelLogoUrl returns the ChannelLogoUrl field if non-nil, zero value otherwise.

### GetChannelLogoUrlOk

`func (o *PublicChannelIntegrationChannelCreate) GetChannelLogoUrlOk() (*string, bool)`

GetChannelLogoUrlOk returns a tuple with the ChannelLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLogoUrl

`func (o *PublicChannelIntegrationChannelCreate) SetChannelLogoUrl(v string)`

SetChannelLogoUrl sets ChannelLogoUrl field to given value.

### HasChannelLogoUrl

`func (o *PublicChannelIntegrationChannelCreate) HasChannelLogoUrl() bool`

HasChannelLogoUrl returns a boolean if a field has been set.

### GetName

`func (o *PublicChannelIntegrationChannelCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicChannelIntegrationChannelCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicChannelIntegrationChannelCreate) SetName(v string)`

SetName sets Name field to given value.


### GetWebhookUrl

`func (o *PublicChannelIntegrationChannelCreate) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *PublicChannelIntegrationChannelCreate) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *PublicChannelIntegrationChannelCreate) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *PublicChannelIntegrationChannelCreate) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


