# PropertyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrameworkFilterId** | Pointer to **int64** |  | [optional] 
**Property** | **string** |  | 
**FilterType** | **string** |  | [default to "PROPERTY"]
**Operation** | [**PropertyFilterOperation**](PropertyFilterOperation.md) |  | 

## Methods

### NewPropertyFilter

`func NewPropertyFilter(property string, filterType string, operation PropertyFilterOperation, ) *PropertyFilter`

NewPropertyFilter instantiates a new PropertyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyFilterWithDefaults

`func NewPropertyFilterWithDefaults() *PropertyFilter`

NewPropertyFilterWithDefaults instantiates a new PropertyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrameworkFilterId

`func (o *PropertyFilter) GetFrameworkFilterId() int64`

GetFrameworkFilterId returns the FrameworkFilterId field if non-nil, zero value otherwise.

### GetFrameworkFilterIdOk

`func (o *PropertyFilter) GetFrameworkFilterIdOk() (*int64, bool)`

GetFrameworkFilterIdOk returns a tuple with the FrameworkFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkFilterId

`func (o *PropertyFilter) SetFrameworkFilterId(v int64)`

SetFrameworkFilterId sets FrameworkFilterId field to given value.

### HasFrameworkFilterId

`func (o *PropertyFilter) HasFrameworkFilterId() bool`

HasFrameworkFilterId returns a boolean if a field has been set.

### GetProperty

`func (o *PropertyFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PropertyFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PropertyFilter) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetFilterType

`func (o *PropertyFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PropertyFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PropertyFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetOperation

`func (o *PropertyFilter) GetOperation() PropertyFilterOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PropertyFilter) GetOperationOk() (*PropertyFilterOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PropertyFilter) SetOperation(v PropertyFilterOperation)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


