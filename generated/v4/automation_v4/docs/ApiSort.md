# ApiSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | **string** |  | 
**Missing** | Pointer to **string** |  | [optional] 
**Order** | **string** |  | 

## Methods

### NewApiSort

`func NewApiSort(property string, order string, ) *ApiSort`

NewApiSort instantiates a new ApiSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSortWithDefaults

`func NewApiSortWithDefaults() *ApiSort`

NewApiSortWithDefaults instantiates a new ApiSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ApiSort) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ApiSort) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ApiSort) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetMissing

`func (o *ApiSort) GetMissing() string`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *ApiSort) GetMissingOk() (*string, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *ApiSort) SetMissing(v string)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *ApiSort) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetOrder

`func (o *ApiSort) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApiSort) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApiSort) SetOrder(v string)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


