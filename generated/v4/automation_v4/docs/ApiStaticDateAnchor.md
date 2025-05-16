# ApiStaticDateAnchor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | The month of the date to anchor on | 
**Year** | Pointer to **int32** | The year of the date to anchor on. If this is not provided then this flow will re-run each year. | [optional] 
**DayOfMonth** | **int32** | The day of the date to anchor on | 
**Type** | **string** | The type of event anchor this is, can be: \&quot;CONTACT_PROPERTY_ANCHOR\&quot; or \&quot;STATIC_DATE_ANCHOR\&quot; | [default to "STATIC_DATE_ANCHOR"]

## Methods

### NewApiStaticDateAnchor

`func NewApiStaticDateAnchor(month string, dayOfMonth int32, type_ string, ) *ApiStaticDateAnchor`

NewApiStaticDateAnchor instantiates a new ApiStaticDateAnchor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaticDateAnchorWithDefaults

`func NewApiStaticDateAnchorWithDefaults() *ApiStaticDateAnchor`

NewApiStaticDateAnchorWithDefaults instantiates a new ApiStaticDateAnchor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *ApiStaticDateAnchor) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ApiStaticDateAnchor) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ApiStaticDateAnchor) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetYear

`func (o *ApiStaticDateAnchor) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ApiStaticDateAnchor) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ApiStaticDateAnchor) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ApiStaticDateAnchor) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *ApiStaticDateAnchor) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *ApiStaticDateAnchor) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *ApiStaticDateAnchor) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.


### GetType

`func (o *ApiStaticDateAnchor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStaticDateAnchor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStaticDateAnchor) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


