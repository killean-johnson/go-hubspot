# PipelineLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomObjectTypes** | [**CustomObjectRecordLimitResponse**](CustomObjectRecordLimitResponse.md) |  | 
**HubspotDefinedObjectTypes** | [**[]LimitAndUsageForObjectType**](LimitAndUsageForObjectType.md) |  | 

## Methods

### NewPipelineLimitResponse

`func NewPipelineLimitResponse(customObjectTypes CustomObjectRecordLimitResponse, hubspotDefinedObjectTypes []LimitAndUsageForObjectType, ) *PipelineLimitResponse`

NewPipelineLimitResponse instantiates a new PipelineLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineLimitResponseWithDefaults

`func NewPipelineLimitResponseWithDefaults() *PipelineLimitResponse`

NewPipelineLimitResponseWithDefaults instantiates a new PipelineLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomObjectTypes

`func (o *PipelineLimitResponse) GetCustomObjectTypes() CustomObjectRecordLimitResponse`

GetCustomObjectTypes returns the CustomObjectTypes field if non-nil, zero value otherwise.

### GetCustomObjectTypesOk

`func (o *PipelineLimitResponse) GetCustomObjectTypesOk() (*CustomObjectRecordLimitResponse, bool)`

GetCustomObjectTypesOk returns a tuple with the CustomObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomObjectTypes

`func (o *PipelineLimitResponse) SetCustomObjectTypes(v CustomObjectRecordLimitResponse)`

SetCustomObjectTypes sets CustomObjectTypes field to given value.


### GetHubspotDefinedObjectTypes

`func (o *PipelineLimitResponse) GetHubspotDefinedObjectTypes() []LimitAndUsageForObjectType`

GetHubspotDefinedObjectTypes returns the HubspotDefinedObjectTypes field if non-nil, zero value otherwise.

### GetHubspotDefinedObjectTypesOk

`func (o *PipelineLimitResponse) GetHubspotDefinedObjectTypesOk() (*[]LimitAndUsageForObjectType, bool)`

GetHubspotDefinedObjectTypesOk returns a tuple with the HubspotDefinedObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubspotDefinedObjectTypes

`func (o *PipelineLimitResponse) SetHubspotDefinedObjectTypes(v []LimitAndUsageForObjectType)`

SetHubspotDefinedObjectTypes sets HubspotDefinedObjectTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


