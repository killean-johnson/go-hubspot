# PublicBudgetTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpendItems** | [**[]PublicSpendItem**](PublicSpendItem.md) |  | 
**BudgetTotal** | Pointer to **float32** |  | [optional] 
**RemainingBudget** | Pointer to **float32** |  | [optional] 
**SpendTotal** | Pointer to **float32** |  | [optional] 
**CurrencyCode** | **string** |  | 
**BudgetItems** | [**[]PublicBudgetItem**](PublicBudgetItem.md) |  | 

## Methods

### NewPublicBudgetTotals

`func NewPublicBudgetTotals(spendItems []PublicSpendItem, currencyCode string, budgetItems []PublicBudgetItem, ) *PublicBudgetTotals`

NewPublicBudgetTotals instantiates a new PublicBudgetTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicBudgetTotalsWithDefaults

`func NewPublicBudgetTotalsWithDefaults() *PublicBudgetTotals`

NewPublicBudgetTotalsWithDefaults instantiates a new PublicBudgetTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpendItems

`func (o *PublicBudgetTotals) GetSpendItems() []PublicSpendItem`

GetSpendItems returns the SpendItems field if non-nil, zero value otherwise.

### GetSpendItemsOk

`func (o *PublicBudgetTotals) GetSpendItemsOk() (*[]PublicSpendItem, bool)`

GetSpendItemsOk returns a tuple with the SpendItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendItems

`func (o *PublicBudgetTotals) SetSpendItems(v []PublicSpendItem)`

SetSpendItems sets SpendItems field to given value.


### GetBudgetTotal

`func (o *PublicBudgetTotals) GetBudgetTotal() float32`

GetBudgetTotal returns the BudgetTotal field if non-nil, zero value otherwise.

### GetBudgetTotalOk

`func (o *PublicBudgetTotals) GetBudgetTotalOk() (*float32, bool)`

GetBudgetTotalOk returns a tuple with the BudgetTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetTotal

`func (o *PublicBudgetTotals) SetBudgetTotal(v float32)`

SetBudgetTotal sets BudgetTotal field to given value.

### HasBudgetTotal

`func (o *PublicBudgetTotals) HasBudgetTotal() bool`

HasBudgetTotal returns a boolean if a field has been set.

### GetRemainingBudget

`func (o *PublicBudgetTotals) GetRemainingBudget() float32`

GetRemainingBudget returns the RemainingBudget field if non-nil, zero value otherwise.

### GetRemainingBudgetOk

`func (o *PublicBudgetTotals) GetRemainingBudgetOk() (*float32, bool)`

GetRemainingBudgetOk returns a tuple with the RemainingBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBudget

`func (o *PublicBudgetTotals) SetRemainingBudget(v float32)`

SetRemainingBudget sets RemainingBudget field to given value.

### HasRemainingBudget

`func (o *PublicBudgetTotals) HasRemainingBudget() bool`

HasRemainingBudget returns a boolean if a field has been set.

### GetSpendTotal

`func (o *PublicBudgetTotals) GetSpendTotal() float32`

GetSpendTotal returns the SpendTotal field if non-nil, zero value otherwise.

### GetSpendTotalOk

`func (o *PublicBudgetTotals) GetSpendTotalOk() (*float32, bool)`

GetSpendTotalOk returns a tuple with the SpendTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendTotal

`func (o *PublicBudgetTotals) SetSpendTotal(v float32)`

SetSpendTotal sets SpendTotal field to given value.

### HasSpendTotal

`func (o *PublicBudgetTotals) HasSpendTotal() bool`

HasSpendTotal returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PublicBudgetTotals) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PublicBudgetTotals) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PublicBudgetTotals) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetBudgetItems

`func (o *PublicBudgetTotals) GetBudgetItems() []PublicBudgetItem`

GetBudgetItems returns the BudgetItems field if non-nil, zero value otherwise.

### GetBudgetItemsOk

`func (o *PublicBudgetTotals) GetBudgetItemsOk() (*[]PublicBudgetItem, bool)`

GetBudgetItemsOk returns a tuple with the BudgetItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetItems

`func (o *PublicBudgetTotals) SetBudgetItems(v []PublicBudgetItem)`

SetBudgetItems sets BudgetItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


