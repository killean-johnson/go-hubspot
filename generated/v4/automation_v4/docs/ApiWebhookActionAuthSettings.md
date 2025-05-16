# ApiWebhookActionAuthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretName** | **string** | The secret to pass through in this auth key. | 
**Name** | **string** | The name to use for this auth key. | 
**Location** | **string** | Where in the request this auth key should be located: \&quot;HEADER\&quot; or \&quot;QUERY_PARAM\&quot; | 
**Type** | **string** | The type of webhook auth settings this is, can be: \&quot;AUTH_KEY\&quot; or \&quot;SIGNATURE\&quot; | [default to "AUTH_KEY"]
**AppId** | **int32** | The appId that this signature will be generated for. | 

## Methods

### NewApiWebhookActionAuthSettings

`func NewApiWebhookActionAuthSettings(secretName string, name string, location string, type_ string, appId int32, ) *ApiWebhookActionAuthSettings`

NewApiWebhookActionAuthSettings instantiates a new ApiWebhookActionAuthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWebhookActionAuthSettingsWithDefaults

`func NewApiWebhookActionAuthSettingsWithDefaults() *ApiWebhookActionAuthSettings`

NewApiWebhookActionAuthSettingsWithDefaults instantiates a new ApiWebhookActionAuthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretName

`func (o *ApiWebhookActionAuthSettings) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *ApiWebhookActionAuthSettings) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *ApiWebhookActionAuthSettings) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetName

`func (o *ApiWebhookActionAuthSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiWebhookActionAuthSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiWebhookActionAuthSettings) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *ApiWebhookActionAuthSettings) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApiWebhookActionAuthSettings) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApiWebhookActionAuthSettings) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetType

`func (o *ApiWebhookActionAuthSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiWebhookActionAuthSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiWebhookActionAuthSettings) SetType(v string)`

SetType sets Type field to given value.


### GetAppId

`func (o *ApiWebhookActionAuthSettings) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ApiWebhookActionAuthSettings) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ApiWebhookActionAuthSettings) SetAppId(v int32)`

SetAppId sets AppId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


