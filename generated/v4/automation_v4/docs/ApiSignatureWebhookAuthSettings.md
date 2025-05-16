# ApiSignatureWebhookAuthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int32** | The appId that this signature will be generated for. | 
**Type** | **string** | The type of webhook auth settings this is, can be: \&quot;AUTH_KEY\&quot; or \&quot;SIGNATURE\&quot; | [default to "SIGNATURE"]

## Methods

### NewApiSignatureWebhookAuthSettings

`func NewApiSignatureWebhookAuthSettings(appId int32, type_ string, ) *ApiSignatureWebhookAuthSettings`

NewApiSignatureWebhookAuthSettings instantiates a new ApiSignatureWebhookAuthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSignatureWebhookAuthSettingsWithDefaults

`func NewApiSignatureWebhookAuthSettingsWithDefaults() *ApiSignatureWebhookAuthSettings`

NewApiSignatureWebhookAuthSettingsWithDefaults instantiates a new ApiSignatureWebhookAuthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ApiSignatureWebhookAuthSettings) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ApiSignatureWebhookAuthSettings) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ApiSignatureWebhookAuthSettings) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetType

`func (o *ApiSignatureWebhookAuthSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiSignatureWebhookAuthSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiSignatureWebhookAuthSettings) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


