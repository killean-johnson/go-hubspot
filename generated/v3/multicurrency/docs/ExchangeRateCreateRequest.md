# ExchangeRateCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveAt** | Pointer to **time.Time** |  | [optional] 
**ConversionRate** | **float32** |  | 
**FromCurrencyCode** | **string** |  | 

## Methods

### NewExchangeRateCreateRequest

`func NewExchangeRateCreateRequest(conversionRate float32, fromCurrencyCode string, ) *ExchangeRateCreateRequest`

NewExchangeRateCreateRequest instantiates a new ExchangeRateCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRateCreateRequestWithDefaults

`func NewExchangeRateCreateRequestWithDefaults() *ExchangeRateCreateRequest`

NewExchangeRateCreateRequestWithDefaults instantiates a new ExchangeRateCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveAt

`func (o *ExchangeRateCreateRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *ExchangeRateCreateRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *ExchangeRateCreateRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *ExchangeRateCreateRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### GetConversionRate

`func (o *ExchangeRateCreateRequest) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *ExchangeRateCreateRequest) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *ExchangeRateCreateRequest) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.


### GetFromCurrencyCode

`func (o *ExchangeRateCreateRequest) GetFromCurrencyCode() string`

GetFromCurrencyCode returns the FromCurrencyCode field if non-nil, zero value otherwise.

### GetFromCurrencyCodeOk

`func (o *ExchangeRateCreateRequest) GetFromCurrencyCodeOk() (*string, bool)`

GetFromCurrencyCodeOk returns a tuple with the FromCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrencyCode

`func (o *ExchangeRateCreateRequest) SetFromCurrencyCode(v string)`

SetFromCurrencyCode sets FromCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


