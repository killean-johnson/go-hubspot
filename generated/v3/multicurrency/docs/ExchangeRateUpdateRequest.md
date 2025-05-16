# ExchangeRateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**ConversionRate** | **float32** |  | 

## Methods

### NewExchangeRateUpdateRequest

`func NewExchangeRateUpdateRequest(id string, conversionRate float32, ) *ExchangeRateUpdateRequest`

NewExchangeRateUpdateRequest instantiates a new ExchangeRateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRateUpdateRequestWithDefaults

`func NewExchangeRateUpdateRequestWithDefaults() *ExchangeRateUpdateRequest`

NewExchangeRateUpdateRequestWithDefaults instantiates a new ExchangeRateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveAt

`func (o *ExchangeRateUpdateRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *ExchangeRateUpdateRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *ExchangeRateUpdateRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *ExchangeRateUpdateRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### GetId

`func (o *ExchangeRateUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExchangeRateUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExchangeRateUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetConversionRate

`func (o *ExchangeRateUpdateRequest) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *ExchangeRateUpdateRequest) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *ExchangeRateUpdateRequest) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


