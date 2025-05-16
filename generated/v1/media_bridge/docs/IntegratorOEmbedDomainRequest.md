# IntegratorOEmbedDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | [**Endpoints**](Endpoints.md) |  | 
**PortalId** | Pointer to **int32** |  | [optional] 

## Methods

### NewIntegratorOEmbedDomainRequest

`func NewIntegratorOEmbedDomainRequest(endpoints Endpoints, ) *IntegratorOEmbedDomainRequest`

NewIntegratorOEmbedDomainRequest instantiates a new IntegratorOEmbedDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegratorOEmbedDomainRequestWithDefaults

`func NewIntegratorOEmbedDomainRequestWithDefaults() *IntegratorOEmbedDomainRequest`

NewIntegratorOEmbedDomainRequestWithDefaults instantiates a new IntegratorOEmbedDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *IntegratorOEmbedDomainRequest) GetEndpoints() Endpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *IntegratorOEmbedDomainRequest) GetEndpointsOk() (*Endpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *IntegratorOEmbedDomainRequest) SetEndpoints(v Endpoints)`

SetEndpoints sets Endpoints field to given value.


### GetPortalId

`func (o *IntegratorOEmbedDomainRequest) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *IntegratorOEmbedDomainRequest) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *IntegratorOEmbedDomainRequest) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *IntegratorOEmbedDomainRequest) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


