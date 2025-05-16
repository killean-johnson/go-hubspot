# ApiAuthKeyWebhookAuthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretName** | **string** | The secret to pass through in this auth key. | 
**Name** | **string** | The name to use for this auth key. | 
**Location** | **string** | Where in the request this auth key should be located: \&quot;HEADER\&quot; or \&quot;QUERY_PARAM\&quot; | 
**Type** | **string** | The type of webhook auth settings this is, can be: \&quot;AUTH_KEY\&quot; or \&quot;SIGNATURE\&quot; | [default to "AUTH_KEY"]

## Methods

### NewApiAuthKeyWebhookAuthSettings

`func NewApiAuthKeyWebhookAuthSettings(secretName string, name string, location string, type_ string, ) *ApiAuthKeyWebhookAuthSettings`

NewApiAuthKeyWebhookAuthSettings instantiates a new ApiAuthKeyWebhookAuthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAuthKeyWebhookAuthSettingsWithDefaults

`func NewApiAuthKeyWebhookAuthSettingsWithDefaults() *ApiAuthKeyWebhookAuthSettings`

NewApiAuthKeyWebhookAuthSettingsWithDefaults instantiates a new ApiAuthKeyWebhookAuthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretName

`func (o *ApiAuthKeyWebhookAuthSettings) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *ApiAuthKeyWebhookAuthSettings) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *ApiAuthKeyWebhookAuthSettings) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetName

`func (o *ApiAuthKeyWebhookAuthSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAuthKeyWebhookAuthSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAuthKeyWebhookAuthSettings) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *ApiAuthKeyWebhookAuthSettings) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApiAuthKeyWebhookAuthSettings) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApiAuthKeyWebhookAuthSettings) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetType

`func (o *ApiAuthKeyWebhookAuthSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAuthKeyWebhookAuthSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAuthKeyWebhookAuthSettings) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


