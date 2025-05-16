# Endpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discovery** | **bool** |  | 
**Schemes** | **[]string** |  | 
**Url** | **string** |  | 

## Methods

### NewEndpoints

`func NewEndpoints(discovery bool, schemes []string, url string, ) *Endpoints`

NewEndpoints instantiates a new Endpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointsWithDefaults

`func NewEndpointsWithDefaults() *Endpoints`

NewEndpointsWithDefaults instantiates a new Endpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscovery

`func (o *Endpoints) GetDiscovery() bool`

GetDiscovery returns the Discovery field if non-nil, zero value otherwise.

### GetDiscoveryOk

`func (o *Endpoints) GetDiscoveryOk() (*bool, bool)`

GetDiscoveryOk returns a tuple with the Discovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovery

`func (o *Endpoints) SetDiscovery(v bool)`

SetDiscovery sets Discovery field to given value.


### GetSchemes

`func (o *Endpoints) GetSchemes() []string`

GetSchemes returns the Schemes field if non-nil, zero value otherwise.

### GetSchemesOk

`func (o *Endpoints) GetSchemesOk() (*[]string, bool)`

GetSchemesOk returns a tuple with the Schemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemes

`func (o *Endpoints) SetSchemes(v []string)`

SetSchemes sets Schemes field to given value.


### GetUrl

`func (o *Endpoints) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Endpoints) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Endpoints) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


