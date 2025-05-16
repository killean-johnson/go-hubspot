# ApiContactFlowAllOfEventAnchor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of event anchor this is, can be: \&quot;CONTACT_PROPERTY_ANCHOR\&quot; or \&quot;STATIC_DATE_ANCHOR\&quot; | [default to "CONTACT_PROPERTY_ANCHOR"]
**ContactProperty** | **string** | A date property on the contact to use as the anchor point of this workflow. | 
**Month** | **string** | The month of the date to anchor on | 
**Year** | Pointer to **int32** | The year of the date to anchor on. If this is not provided then this flow will re-run each year. | [optional] 
**DayOfMonth** | **int32** | The day of the date to anchor on | 

## Methods

### NewApiContactFlowAllOfEventAnchor

`func NewApiContactFlowAllOfEventAnchor(type_ string, contactProperty string, month string, dayOfMonth int32, ) *ApiContactFlowAllOfEventAnchor`

NewApiContactFlowAllOfEventAnchor instantiates a new ApiContactFlowAllOfEventAnchor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactFlowAllOfEventAnchorWithDefaults

`func NewApiContactFlowAllOfEventAnchorWithDefaults() *ApiContactFlowAllOfEventAnchor`

NewApiContactFlowAllOfEventAnchorWithDefaults instantiates a new ApiContactFlowAllOfEventAnchor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiContactFlowAllOfEventAnchor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactFlowAllOfEventAnchor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactFlowAllOfEventAnchor) SetType(v string)`

SetType sets Type field to given value.


### GetContactProperty

`func (o *ApiContactFlowAllOfEventAnchor) GetContactProperty() string`

GetContactProperty returns the ContactProperty field if non-nil, zero value otherwise.

### GetContactPropertyOk

`func (o *ApiContactFlowAllOfEventAnchor) GetContactPropertyOk() (*string, bool)`

GetContactPropertyOk returns a tuple with the ContactProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactProperty

`func (o *ApiContactFlowAllOfEventAnchor) SetContactProperty(v string)`

SetContactProperty sets ContactProperty field to given value.


### GetMonth

`func (o *ApiContactFlowAllOfEventAnchor) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ApiContactFlowAllOfEventAnchor) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ApiContactFlowAllOfEventAnchor) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetYear

`func (o *ApiContactFlowAllOfEventAnchor) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ApiContactFlowAllOfEventAnchor) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ApiContactFlowAllOfEventAnchor) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ApiContactFlowAllOfEventAnchor) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *ApiContactFlowAllOfEventAnchor) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *ApiContactFlowAllOfEventAnchor) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *ApiContactFlowAllOfEventAnchor) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


