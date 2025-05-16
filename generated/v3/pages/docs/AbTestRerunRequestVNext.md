# AbTestRerunRequestVNext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariationId** | **string** | ID of the object to reactivate as a test variation. | 
**AbTestId** | **string** | ID of the test to rerun. | 

## Methods

### NewAbTestRerunRequestVNext

`func NewAbTestRerunRequestVNext(variationId string, abTestId string, ) *AbTestRerunRequestVNext`

NewAbTestRerunRequestVNext instantiates a new AbTestRerunRequestVNext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbTestRerunRequestVNextWithDefaults

`func NewAbTestRerunRequestVNextWithDefaults() *AbTestRerunRequestVNext`

NewAbTestRerunRequestVNextWithDefaults instantiates a new AbTestRerunRequestVNext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariationId

`func (o *AbTestRerunRequestVNext) GetVariationId() string`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *AbTestRerunRequestVNext) GetVariationIdOk() (*string, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *AbTestRerunRequestVNext) SetVariationId(v string)`

SetVariationId sets VariationId field to given value.


### GetAbTestId

`func (o *AbTestRerunRequestVNext) GetAbTestId() string`

GetAbTestId returns the AbTestId field if non-nil, zero value otherwise.

### GetAbTestIdOk

`func (o *AbTestRerunRequestVNext) GetAbTestIdOk() (*string, bool)`

GetAbTestIdOk returns a tuple with the AbTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbTestId

`func (o *AbTestRerunRequestVNext) SetAbTestId(v string)`

SetAbTestId sets AbTestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


