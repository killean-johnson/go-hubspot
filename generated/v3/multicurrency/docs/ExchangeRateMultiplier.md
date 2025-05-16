# ExchangeRateMultiplier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveAt** | Pointer to **time.Time** |  | [optional] 
**ConversionRate** | **float32** |  | 

## Methods

### NewExchangeRateMultiplier

`func NewExchangeRateMultiplier(conversionRate float32, ) *ExchangeRateMultiplier`

NewExchangeRateMultiplier instantiates a new ExchangeRateMultiplier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRateMultiplierWithDefaults

`func NewExchangeRateMultiplierWithDefaults() *ExchangeRateMultiplier`

NewExchangeRateMultiplierWithDefaults instantiates a new ExchangeRateMultiplier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveAt

`func (o *ExchangeRateMultiplier) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *ExchangeRateMultiplier) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *ExchangeRateMultiplier) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *ExchangeRateMultiplier) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### GetConversionRate

`func (o *ExchangeRateMultiplier) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *ExchangeRateMultiplier) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *ExchangeRateMultiplier) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


