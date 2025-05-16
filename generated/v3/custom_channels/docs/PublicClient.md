# PublicClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientType** | **string** |  | 
**IntegrationAppId** | Pointer to **int32** |  | [optional] 

## Methods

### NewPublicClient

`func NewPublicClient(clientType string, ) *PublicClient`

NewPublicClient instantiates a new PublicClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicClientWithDefaults

`func NewPublicClientWithDefaults() *PublicClient`

NewPublicClientWithDefaults instantiates a new PublicClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientType

`func (o *PublicClient) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *PublicClient) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *PublicClient) SetClientType(v string)`

SetClientType sets ClientType field to given value.


### GetIntegrationAppId

`func (o *PublicClient) GetIntegrationAppId() int32`

GetIntegrationAppId returns the IntegrationAppId field if non-nil, zero value otherwise.

### GetIntegrationAppIdOk

`func (o *PublicClient) GetIntegrationAppIdOk() (*int32, bool)`

GetIntegrationAppIdOk returns a tuple with the IntegrationAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationAppId

`func (o *PublicClient) SetIntegrationAppId(v int32)`

SetIntegrationAppId sets IntegrationAppId field to given value.

### HasIntegrationAppId

`func (o *PublicClient) HasIntegrationAppId() bool`

HasIntegrationAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


