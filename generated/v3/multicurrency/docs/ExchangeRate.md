# ExchangeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**ToCurrencyCode** | **string** |  | 
**VisibleInUI** | **bool** |  | 
**EffectiveAt** | **time.Time** |  | 
**Id** | **string** |  | 
**ConversionRate** | **float32** |  | 
**FromCurrencyCode** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewExchangeRate

`func NewExchangeRate(createdAt time.Time, toCurrencyCode string, visibleInUI bool, effectiveAt time.Time, id string, conversionRate float32, fromCurrencyCode string, updatedAt time.Time, ) *ExchangeRate`

NewExchangeRate instantiates a new ExchangeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRateWithDefaults

`func NewExchangeRateWithDefaults() *ExchangeRate`

NewExchangeRateWithDefaults instantiates a new ExchangeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ExchangeRate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExchangeRate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExchangeRate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetToCurrencyCode

`func (o *ExchangeRate) GetToCurrencyCode() string`

GetToCurrencyCode returns the ToCurrencyCode field if non-nil, zero value otherwise.

### GetToCurrencyCodeOk

`func (o *ExchangeRate) GetToCurrencyCodeOk() (*string, bool)`

GetToCurrencyCodeOk returns a tuple with the ToCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrencyCode

`func (o *ExchangeRate) SetToCurrencyCode(v string)`

SetToCurrencyCode sets ToCurrencyCode field to given value.


### GetVisibleInUI

`func (o *ExchangeRate) GetVisibleInUI() bool`

GetVisibleInUI returns the VisibleInUI field if non-nil, zero value otherwise.

### GetVisibleInUIOk

`func (o *ExchangeRate) GetVisibleInUIOk() (*bool, bool)`

GetVisibleInUIOk returns a tuple with the VisibleInUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleInUI

`func (o *ExchangeRate) SetVisibleInUI(v bool)`

SetVisibleInUI sets VisibleInUI field to given value.


### GetEffectiveAt

`func (o *ExchangeRate) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *ExchangeRate) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *ExchangeRate) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetId

`func (o *ExchangeRate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExchangeRate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExchangeRate) SetId(v string)`

SetId sets Id field to given value.


### GetConversionRate

`func (o *ExchangeRate) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *ExchangeRate) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *ExchangeRate) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.


### GetFromCurrencyCode

`func (o *ExchangeRate) GetFromCurrencyCode() string`

GetFromCurrencyCode returns the FromCurrencyCode field if non-nil, zero value otherwise.

### GetFromCurrencyCodeOk

`func (o *ExchangeRate) GetFromCurrencyCodeOk() (*string, bool)`

GetFromCurrencyCodeOk returns a tuple with the FromCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrencyCode

`func (o *ExchangeRate) SetFromCurrencyCode(v string)`

SetFromCurrencyCode sets FromCurrencyCode field to given value.


### GetUpdatedAt

`func (o *ExchangeRate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExchangeRate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExchangeRate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


